package nmongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://root:4124edq34%26r42@lizo.top"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}
	// insert
	collect := client.Database("corn").Collection("jobs")
	record := &RecordLog{
		JobName: "job1",
		Command: "main.go",
		Err:     "",
		Content: "Hello_World",
		Timepoint: TimePoint{
			StartTime: time.Now().Unix(),
			EndTIme:   time.Now().Unix() + 10,
		},
	}
	insertRest, err := collect.InsertOne(ctx, record)
	if err != nil {
		fmt.Println(err)
		return
	}
	insertID := insertRest.InsertedID.(primitive.ObjectID)
	fmt.Printf("%v", insertID)
	// adsaaaaaaaaaaa
	// 创建需要过滤的条件
	logred := &LogRecord{
		JobName: "job1",
	}
	var skip int64 = 0   //从那个开始
	var limit int64 = 20 //炼制几个输出字段
	cursor, err := client.Database("corn").Collection("jobs").Find(context.Background(), logred, &options.FindOptions{
		Skip:  &skip,
		Limit: &limit,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		//创建需要反序列化成什么样子的结构体对象
		records := &RecordLog{}
		//反序列化
		err = cursor.Decode(records)
		if err != nil {
			fmt.Println(err)
			return
		}
		//打印
		fmt.Println(*records)
	}
}

type TimePoint struct {
	StartTime int64 `bson:"startTime"`
	EndTIme   int64 `bson:"endTime"`
}

// 存储在mongodb中的内容
type RecordLog struct {
	JobName   string    `bson:"jobName"`
	Command   string    `bson:"command"`
	Err       string    `bson:"err"`
	Content   string    `bson:"content"`
	Timepoint TimePoint `bson:"timepoint"`
}

type LogRecord struct {
	JobName string `bson:"jobName"`
}
