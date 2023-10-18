package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"jjbot/service/bot/botapi"
	"jjbot/service/logger"
	"time"
)

func (client *api) InsertOne(d string, c string, a any) *mongo.InsertOneResult {
	if client == nil {
		logger.SugarLogger.Error("未链接数据库")
		return nil
	}
	collection := client.c.Database(d).Collection(c)
	insertResult, err := collection.InsertOne(context.TODO(), a)
	if err != nil {
		logger.SugarLogger.Errorf("插入数据错误: %v", err)
		botapi.SendPrivateMsg(345793738, "插入数据错误", false)
	}
	logger.SugarLogger.Debugf("插入单条数据：%s", insertResult.InsertedID)
	return insertResult
}

func (client *api) InsertMany(d string, c string, a []any) *mongo.InsertManyResult {
	if client == nil {
		logger.SugarLogger.Error("未链接数据库")
		return nil
	}
	collection := client.c.Database(d).Collection(c)
	insertManyResult, err := collection.InsertMany(context.TODO(), a)
	if err != nil {
		logger.SugarLogger.Errorf("插入数据错误: %v", err)
		botapi.SendPrivateMsg(345793738, "插入数据错误", false)
	}
	logger.SugarLogger.Debugf("插入多个数据: %d", insertManyResult.InsertedIDs)
	return insertManyResult
}

func (client *api) FindOne(d string, c string, b any, opt *options.FindOneOptions) bson.D {
	if client == nil {
		logger.SugarLogger.Error("未链接数据库")
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var result bson.D
	collection := client.c.Database(d).Collection(c)
	err := collection.FindOne(ctx, b, opt).Decode(&result)
	if err != nil {
		//logger.SugarLogger.Errorf("寻找数据错误: %v", err)
	}
	logger.SugarLogger.Debugf("FindOne: %v", result)
	return result
}

// FindOneUnmarshal 输入struct变量指针，数值会通过指针传入变量，不返回值
func (client *api) FindOneUnmarshal(d string, c string, b any, opt *options.FindOneOptions, r any) {
	if client == nil {
		logger.SugarLogger.Error("未链接数据库")
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.c.Database(d).Collection(c)
	err := collection.FindOne(ctx, b, opt).Decode(r)
	if err != nil {
		//logger.SugarLogger.Errorf("寻找数据错误: %v", err)
	}
	logger.SugarLogger.Debugf("FindOne: %v", r)
}

func (client *api) Find(d string, c string, b any, opt *options.FindOptions) []bson.D {
	if client == nil {
		logger.SugarLogger.Error("未链接数据库")
		return nil
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.c.Database(d).Collection(c)
	cur, err := collection.Find(ctx, b, opt)
	if err != nil {
		//logger.SugarLogger.Errorf("寻找数据错误: %v", err)
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			//logger.SugarLogger.Errorf("寻找数据错误: %v", err)
		}
	}(cur, ctx)
	var result []bson.D
	for cur.Next(ctx) {
		var r bson.D
		err := cur.Decode(&r)
		if err != nil {
			logger.SugarLogger.Errorf("寻找数据错误: %v", err)
		}
		result = append(result, r)
	}
	if err := cur.Err(); err != nil {
		logger.SugarLogger.Errorf("寻找数据错误: %v", err)
	}
	logger.SugarLogger.Debugf("FindAll: 共%d条数据", len(result))
	return result
}

// FindUnmarshal 输入struct变量指针，来foreach该指针的数组
func (client *api) FindUnmarshal(d string, c string, b any, opt *options.FindOptions, t any, f func(int, *any)) int {
	if client == nil {
		logger.SugarLogger.Error("未链接数据库")
		return 0
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	collection := client.c.Database(d).Collection(c)
	cur, err := collection.Find(ctx, b, opt)
	if err != nil {
		//logger.SugarLogger.Errorf("寻找数据错误: %v", err)
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			//logger.SugarLogger.Errorf("寻找数据错误: %v", err)
		}
	}(cur, ctx)
	result := 0
	for cur.Next(ctx) {
		err := cur.Decode(t)
		if err != nil {
			logger.SugarLogger.Errorf("寻找数据错误: %v", err)
		}
		f(result, &t)
		result++
	}
	if err := cur.Err(); err != nil {
		logger.SugarLogger.Errorf("寻找数据错误: %v", err)
	}
	logger.SugarLogger.Debugf("FindAll: 共%d条数据", result)
	return result
}

func (client *api) UpdateOne(d string, c string, b any, a any) *mongo.UpdateResult {
	a = bson.D{{"$set", a}}
	if client == nil {
		logger.SugarLogger.Error("未链接数据库")
	}
	ctx := context.TODO()
	collection := client.c.Database(d).Collection(c)

	ur, err := collection.UpdateOne(ctx, b, a)
	if err != nil {
		logger.SugarLogger.Errorf("更新数据错误: %v", err)
		botapi.SendPrivateMsg(345793738, "更新数据错误", false)
	}
	logger.SugarLogger.Debugf("更新了%v条数据", ur.ModifiedCount)
	return ur
}

func (client *api) UpdateMany(d string, c string, b any, a any) *mongo.UpdateResult {
	a = bson.D{{"$set", a}}
	if client == nil {
		logger.SugarLogger.Error("未链接数据库")
		return nil
	}
	ctx := context.TODO()
	collection := client.c.Database(d).Collection(c)

	ur, err := collection.UpdateMany(ctx, b, a)
	if err != nil {
		logger.SugarLogger.Errorf("更新数据错误: %v", err)
		botapi.SendPrivateMsg(345793738, "更新数据错误", false)
	}
	logger.SugarLogger.Debugf("更新了%v条数据", ur.ModifiedCount)
	return ur
}

func (client *api) DeleteOne(d string, c string, b any) *mongo.DeleteResult {
	if client == nil {
		logger.SugarLogger.Error("未链接数据库")
		return nil
	}
	collection := client.c.Database(d).Collection(c)
	ctx := context.TODO()

	dr, err := collection.DeleteOne(ctx, b)

	if err != nil {
		logger.SugarLogger.Errorf("删除数据错误: %v", err)
	}
	logger.SugarLogger.Debugf("删除了%v条数据", dr.DeletedCount)
	return dr
}

func (client *api) DeleteMany(d string, c string, b any) *mongo.DeleteResult {
	if client == nil {
		logger.SugarLogger.Error("未链接数据库")
		return nil
	}
	collection := client.c.Database(d).Collection(c)
	ctx := context.TODO()

	dr, err := collection.DeleteMany(ctx, b)

	if err != nil {
		logger.SugarLogger.Errorf("删除数据错误: %v", err)
	}
	logger.SugarLogger.Debugf("删除了%v条数据", dr.DeletedCount)
	return dr
}

func (client *api) RunCommand(d string, command bson.D) bool {
	db := client.c.Database(d)
	var result bson.M
	if err := db.RunCommand(context.TODO(), command).Decode(&result); err != nil {
		logger.SugarLogger.Warn(err)
		return false
	}
	logger.SugarLogger.Info(result)
	return true
}
