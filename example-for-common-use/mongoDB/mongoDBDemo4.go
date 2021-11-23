package main

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

type Product struct {
	Id    int    `bson:"province"`
	Name  string `bson:"city"`
	Model string `bson:"model"`
}

type Model struct {
	Name        string `bson:"name"`
	ParentModel string `bson:"parentmodel"`
	Info        string `bson:"info"`
}

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

type M map[string]interface{}

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
	collection_product := client.Database("q1mi").Collection("product")
	collection_model := client.Database("q1mi").Collection("model")
	//插入多条数据
	p1 := Product{1, "产品1", "Aa"}
	p2 := Product{2, "产品2", "Ab"}
	p3 := Product{3, "产品3", "Ab"}
	p4 := Product{4, "产品4", "Ba"}
	p5 := Product{5, "产品5", "Bb"}
	products := []interface{}{
		p1, p2, p3, p4, p5,
	}
	collection_product.InsertMany(context.TODO(), products)
	if err != nil {
		log.Fatal(err)
	}
	models := []interface{}{
		Model{"Aa", "A", "型号Aa"},
		Model{"Ab", "A", "型号Ab"},
		Model{"Ba", "B", "型号Ba"},
		Model{"Bb", "B", "型号Bb"},
	}
	collection_model.InsertMany(context.TODO(), models)
	if err != nil {
		log.Fatal(err)
	}
	//查询一级型号下二级型号的数量
	filter := M{}
	filter["parentmodel"] = "A"
	data := make([]Model, 0)
	Query(collection_model, filter, &data)

	for _, value := range data {
		fmt.Printf("Found multiple documents: %+v\n", value)
	}
	fmt.Println(len(data))
	//查询一级型号的数量 得到不重复的parentmodel数组
	results, err := collection_model.Distinct(context.TODO(), "parentmodel", bson.D{})
	if err != nil {
		panic(err)
	}
	fmt.Println(len(results))
	//查询Ab型号的产品有多少个
	filter1 := M{}
	filter1["model"] = "Ab"
	data1 := make([]Product, 0)
	Query(collection_product, filter, &data1)
	//查询A型号下有多少产品(多表连接查询)
	filter2 := M{}
	modelSet := make([]string, 0)
	for _, v := range data {
		modelSet = append(modelSet, v.Name)
	}
	filter2["model"] = M{
		"$in": modelSet,
	}
	product_data := make([]Product, 0)
	Query(collection_product, filter2, &product_data)
	for _, value := range product_data {
		fmt.Printf("Found multiple documents: %+v\n", value)
	}
	//插入已有型号时报错//插入前检测是否有对应数据//如果有则报错

	// 删除所有
	collection_product.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	collection_model.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	// 断开连接
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connection to MongoDB closed.")

}
