package service

import (
	"chat/conf"
	"chat/model/ws"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sort"
	"time"
)


type SendSortMsg struct {
	Content string `json:"content"`
	Read uint `json:"read"`
	CreateAt int64 `json:"create_at"`
}

func InsertMsg(database string, id string, content string, read uint, expire int64) (err error) {
	collection := conf.MongoDBClient.Database(database).Collection(id)
	comment := ws.Trainer{
		Content:   content,
		StartTime: time.Now().Unix(),
		EndTime:   time.Now().Unix() + expire,
		Read:      read,
	}
	_, err = collection.InsertOne(context.TODO(),comment)
	return
}

func FindManyMsg(database string, sendID string,id string,time int64,pageSize int) (results []ws.Result,err error) {
	var resultsMe []ws.Trainer
	var resultsYou []ws.Trainer
	sendIdCollection := conf.MongoDBClient.Database(database).Collection(sendID)
	idCollection := conf.MongoDBClient.Database(database).Collection(id)
	filter := bson.M{"startTime": bson.M{"$lt": time}}
	sendIdTimeCursor, err := sendIdCollection.Find(context.TODO(), filter,
		options.Find().SetSort(bson.D{{"StartTime", -1}}), options.Find().SetLimit(int64(pageSize)))
	idTimeCurcor, err := idCollection.Find(context.TODO(), filter,
		options.Find().SetSort(bson.D{{"StartTime", -1}}), options.Find().SetLimit(int64(pageSize)))
	err = sendIdTimeCursor.All(context.TODO(), &resultsYou) // sendId 对面发过来的
	err = idTimeCurcor.All(context.TODO(), &resultsMe)      // Id 发给对面的
	results, _ = AppendAndSort(resultsMe, resultsYou)
	return
}

func FirstFindMsg(database string,sendId string,id string) (results []ws.Result,err error) {
	// 首次查询(把对方发来的所有未读都取出来)
	var resultsMe []ws.Trainer
	var resultsYou []ws.Trainer
	sendIdCollection := conf.MongoDBClient.Database(database).Collection(sendId)
	idCollection:=conf.MongoDBClient.Database(database).Collection(sendId)
	filter := bson.M{"read":bson.M{
		"&all":[]uint{0},
	}}
	sendIdCursor ,err := sendIdCollection.Find(context.TODO(),filter,options.Find().SetSort(bson.D{{
		"startTime",1}}),options.Find().SetLimit(1))
	if sendIdCursor == nil {
		return
	}
	var unreads []ws.Trainer
	err = sendIdCursor.All(context.TODO(),&unreads)
	if err != nil {
		log.Println("sendIdCursor err",err)
	}
	if len(unreads)>0 {
		timeFilter := bson.M{
			"startTime":bson.M{
				"$gte":unreads[0].StartTime,
			},
		}
		sendIdTimeCursor , _ := sendIdCollection.Find(context.TODO(),timeFilter)
		idtimeCursor ,_ := idCollection.Find(context.TODO(),timeFilter)
		err = sendIdTimeCursor.All(context.TODO(),&resultsYou)
		err = idtimeCursor.All(context.TODO(),&resultsMe)
		results,err = AppendAndSort(resultsMe,resultsYou)
	}else {
		results,err = FindManyMsg(database,sendId,id,99999999,10)
	}
	overTimeFilter := bson.D{
		{"$and",bson.A{
			bson.D{{"endTime",bson.M{"&lt":time.Now().Unix()}}},
			bson.D{{"read",bson.M{"$eq":1}}},
		}},
	}
	_,_ = sendIdCollection.DeleteMany(context.TODO(),overTimeFilter)
	_,_ = idCollection.DeleteMany(context.TODO(),overTimeFilter)
	// 将所有的维度设置为已读
	_,_ = sendIdCollection.UpdateMany(context.TODO(),filter,bson.M{
		"$set":bson.M{"read":1},
	})
	_,_=sendIdCollection.UpdateMany(context.TODO(),filter,bson.M{
		"&set":bson.M{"ebdTime":time.Now().Unix()+int64(3*month)},
	})
	return
}

func AppendAndSort(resultsMe, resultsYou []ws.Trainer) (results []ws.Result, err error) {
	for _,r:=range resultsMe{
		sendSort := SendSortMsg{
			Content:r.Content,
			Read: r.Read,
			CreateAt:r.StartTime,
		}
		result := ws.Result{
			StartTime:r.StartTime,
			Msg: sendSort,
			From:"me",
		}
		results = append(results,result)
	}
	for _,r:=range resultsYou{
		sendSort := SendSortMsg{
			Content:r.Content,
			Read: r.Read,
			CreateAt:r.StartTime,
		}
		result := ws.Result{
			StartTime:r.StartTime,
			Msg: sendSort,
			From:"you",
		}
		results = append(results,result)
	}
	// 最后进行排序
	sort.Slice(results, func(i, j int) bool { return results[i].StartTime < results[j].StartTime })
	return results, nil
}

