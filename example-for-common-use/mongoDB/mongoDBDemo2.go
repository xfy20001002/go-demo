package Demo2

import (
	"context"
	"fmt"
	"log"

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

	filter := bson.D{{"stuaddr.province", "广东"}}

	//查询文档数据
	// 创建一个Student变量用来接收查询的结果
	var result Student
	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found a single document: %+v\n", result)
	//查询多个数据
	findOptions := options.Find()
	findOptions.SetLimit(5)
	fmt.Println("**********************************")
	// 定义一个切片用来存储查询结果
	var results []Student
	// 把bson.D{{}}作为一个filter来匹配所有文档
	cur, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 查找多个文档返回一个光标
	// 遍历游标允许我们一次解码一个文档
	for cur.Next(context.TODO()) {
		// 创建一个值，将单个文档解码为该值
		var elem Student
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	// 完成后关闭游标
	cur.Close(context.TODO())

	for _, value := range results {
		fmt.Printf("Found multiple documents (array of pointers): %+v\n", value)
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
