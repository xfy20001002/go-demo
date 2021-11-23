package Demo3

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Addr struct {
	Province string `bson:"province"`
	City     string `bson:"city"`
}

type Student struct {
	Name    string `bson:"name"`
	Age     int    `bson:"age"`
	StuAddr Addr   `bson:"stuaddr"`
	//不存入数据库时
	//Age  int    `bson:"-"`

}
type M map[string]interface{}

func Query(con *mongo.Collection, filter interface{}, result interface{}) error {
	resultv := reflect.ValueOf(result)
	//fmt.Printf("%T\n", resultv.Kind())
	if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
		return errors.New("result argument must be a slice address")
	}
	slicev := resultv.Elem()
	slicev = slicev.Slice(0, slicev.Cap())
	elemt := slicev.Type().Elem()

	opt := options.Find()

	cur, err := con.Find(context.TODO(), filter, opt)
	if err != nil {
		return err
	}
	defer func() {
		// Close the cursor once finished
		if err := cur.Close(context.TODO()); err != nil {
			log.Println("cursor close error")
		}
	}()

	i := 0
	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {
		if slicev.Len() == i {
			elemp := reflect.New(elemt)
			err := cur.Decode(elemp.Interface())
			if err != nil {
				return err
			}
			slicev = reflect.Append(slicev, elemp.Elem())
			slicev = slicev.Slice(0, slicev.Cap())
		} else {
			err := cur.Decode(slicev.Index(i).Addr().Interface())
			if err != nil {
				return err
			}
		}
		i++
	}

	if err := cur.Err(); err != nil {
		return err
	}

	resultv.Elem().Set(slicev.Slice(0, i))
	return nil
}
func main() {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("**********************************")
	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// 指定获取要操作的数据集
	collection := client.Database("q1mi").Collection("student")
	//CRUD
	//插入
	s1 := Student{}
	s1.Name = "小红"
	s1.Age = 12
	s1.StuAddr.Province = "江西"
	s1.StuAddr.City = "南昌"

	s2 := Student{}
	s2.Name = "小兰"
	s2.Age = 10
	s2.StuAddr.Province = "广东"
	s2.StuAddr.City = "广州"

	s3 := Student{}
	s3.Name = "小黄"
	s3.Age = 11
	s3.StuAddr.Province = "广东"
	s3.StuAddr.City = "深圳"
	fmt.Println("**********************************")
	//插入一条数据
	insertResult, err := collection.InsertOne(context.TODO(), s1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted a single document: ", insertResult.InsertedID)
	fmt.Println("**********************************")
	//插入多条数据
	students := []interface{}{s2, s3}
	insertManyResult, err := collection.InsertMany(context.TODO(), students)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)
	fmt.Println("**********************************")
	filter := M{}
	//filter["name"] = "小红"
	//查询文档数据
	data := make([]Student, 0)
	Query(collection, filter, &data)

	for _, student := range data {
		fmt.Printf("Found multiple documents: %+v\n", student)
	}

	// 删除所有
	deleteResult2, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult2.DeletedCount)

	// 断开连接
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}
