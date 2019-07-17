package mongodbz

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Manager MongoDB客户端管理
var Manager MongoManager

// MongoManager MongoDB客户端管理
type MongoManager struct {
	curDB  *mongo.Database              // 当前数据库
	curCol *mongo.Collection            //当前集合
	c      *mongo.Client                //数据库客户端
	ctx    context.Context              //上下文内容
	dbs    map[string]*mongo.Database   //数据库
	cols   map[string]*mongo.Collection //数据集合
}

// UpdateResult 更新结果
type UpdateResult struct {
	// The number of documents that matched the filter.
	MatchedCount int64
	// The number of documents that were modified.
	ModifiedCount int64
	// The number of documents that were upserted.
	UpsertedCount int64
	// The identifier of the inserted document if an upsert took place.
	UpsertedID interface{}
}

func init() {
	client, err := newMongoDBClient()
	if err != nil {
		fmt.Println("mongo client err: ", err)
		return
	}
	Manager = MongoManager{
		c:    client,
		ctx:  context.Background(),
		dbs:  make(map[string]*mongo.Database),
		cols: make(map[string]*mongo.Collection),
	}
	Manager.UseDatabase("")
	// Manager.UseCollection("")
}

// newMongoDBClient 初始化并连接mongodb
func newMongoDBClient() (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://127.0.0.1:27017"))
	if err != nil {
		return nil, fmt.Errorf("new mongodb client failed: %v", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, fmt.Errorf("connect mongodb client failed: %v", err)
	}
	return client, nil
}

// UseDatabase 使用指定的数据库 默认数据库znk
func (m MongoManager) UseDatabase(dbName string) error {
	if dbName == "" {
		dbName = "znk"
	}
	db := m.c.Database(dbName)
	if db == nil {
		return errors.New("use database failed")
	}
	m.curDB = db
	m.dbs[dbName] = db
	return nil
}

// UseCollection 使用指定集合，默认集合common
func (m MongoManager) UseCollection(cName string) error {
	if cName == "" {
		cName = "common"
	}
	if m.c == nil {
		return errors.New("use database first")
	}
	col := m.curDB.Collection(cName)
	if col == nil {
		return errors.New("use collection failed")
	}
	m.curCol = col
	m.cols[cName] = col
	return nil
}

// InsertOne 插入一条文档数据
func (m MongoManager) InsertOne(doc interface{}) (interface{}, error) {
	if m.curCol == nil {
		return nil, errors.New("use collection first")
	}

	res, err := m.curCol.InsertOne(m.ctx, doc)
	if err != nil {
		return nil, fmt.Errorf("insert one document failed: %v", err)
	}
	// fmt.Println(res.InsertedID)
	return res.InsertedID, nil
}

// InsertMany 插入多条文档数据
func (m MongoManager) InsertMany(docs []interface{}) ([]interface{}, error) {
	if m.curCol == nil {
		return nil, errors.New("use collection first")
	}

	res, err := m.curCol.InsertMany(m.ctx, docs)
	if err != nil {
		return nil, fmt.Errorf("insert many document failed: %v", err)
	}
	// fmt.Println(res.InsertedID)
	return res.InsertedIDs, nil
}

// FindOneAndDecode 查询一条数据，并解码到模型
func (m MongoManager) FindOneAndDecode(filter interface{}, doc interface{}) error {
	if m.curCol == nil {
		return errors.New("use collection first")
	}
	res := m.curCol.FindOne(m.ctx, filter)
	if res == nil || res.Err() != nil {
		return errors.New("find one document failed")
	}
	if err := res.Decode(doc); err != nil {
		return errors.New("find one invalid document")
	}
	return nil
}

// FindManyAndDecode 查询多条数据，并解码到模型
func (m MongoManager) FindManyAndDecode(filter interface{}, doc interface{}, res func(interface{}, error)) error {
	if m.curCol == nil {
		return errors.New("use collection first")
	}
	cur, err := m.curCol.Find(m.ctx, filter) //m.curCol.FindOne(m.ctx, filter)
	if err != nil {
		return fmt.Errorf("find many document failed: %v", err)
	}
	defer cur.Close(m.ctx)
	for cur.Next(m.ctx) {
		err = cur.Decode(doc)
		if err != nil {
			res(nil, err)
		} else {
			res(doc, nil)
		}
	}
	if err = cur.Err(); err != nil {
		return fmt.Errorf("find many document failed: %v", err)
	}
	return nil
}

// DeleteOne 删除一条数据
func (m MongoManager) DeleteOne(filter interface{}) (int64, error) {
	if m.curCol == nil {
		return 0, errors.New("use collection first")
	}
	res, err := m.curCol.DeleteOne(m.ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("delete one document failed: %v", err)
	}

	return res.DeletedCount, nil
}

// DeleteMany 删除多条数据
func (m MongoManager) DeleteMany(filter interface{}) (int64, error) {
	if m.curCol == nil {
		return 0, errors.New("use collection first")
	}
	res, err := m.curCol.DeleteMany(m.ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("delete one document failed: %v", err)
	}

	return res.DeletedCount, nil
}

// UpdateOne 更新一条数据
func (m MongoManager) UpdateOne(filter interface{}, update interface{}) error {
	if m.curCol == nil {
		return errors.New("use collection first")
	}
	ops := options.Update()
	ops.SetUpsert(true)
	_, err := m.curCol.UpdateOne(m.ctx, filter, update, ops) //m.curCol.DeleteMany(m.ctx, filter)
	if err != nil {
		return fmt.Errorf("delete one document failed: %v", err)
	}

	return nil
}

// UpdateMany 更新多条数据
func (m MongoManager) UpdateMany(filter interface{}, update interface{}) (*UpdateResult, error) {
	if m.curCol == nil {
		return nil, errors.New("use collection first")
	}
	ops := options.Update()
	ops.SetUpsert(true)
	res, err := m.curCol.UpdateMany(m.ctx, filter, update, ops) //m.curCol.DeleteMany(m.ctx, filter)
	if err != nil {
		return nil, fmt.Errorf("delete one document failed: %v", err)
	}

	updateRes := &UpdateResult{
		MatchedCount:  res.MatchedCount,
		ModifiedCount: res.ModifiedCount,
		UpsertedCount: res.UpsertedCount,
		UpsertedID:    res.UpsertedID,
	}

	return updateRes, nil
}

// FindOneAndUpdate 查询一条数据并更新
func (m MongoManager) FindOneAndUpdate(filter interface{}, update interface{}, doc interface{}) error {
	if m.curCol == nil {
		return errors.New("use collection first")
	}
	ops := options.FindOneAndUpdate()
	ops.SetUpsert(true)
	res := m.curCol.FindOneAndUpdate(m.ctx, filter, update, ops)
	if res == nil || res.Err() != nil {
		return fmt.Errorf("find one and update document failed: %v", res.Err())
	}
	err := res.Decode(doc)
	if err != nil {
		return fmt.Errorf("find one and update document decode failed: %v", err)
	}
	return nil
}

// FindOneAndDelete 查询一条并删除
func (m MongoManager) FindOneAndDelete(filter interface{}, doc interface{}) error {
	if m.curCol == nil {
		return errors.New("use collection first")
	}
	res := m.curCol.FindOneAndDelete(m.ctx, filter)
	if res == nil || res.Err() != nil {
		return fmt.Errorf("find one and delete document failed: %v", res.Err())
	}
	err := res.Decode(doc)
	if err != nil {
		return fmt.Errorf("find one and delete document decode failed: %v", res.Err())
	}
	return nil
}

// FindOneAndReplace 查询一条并代替
func (m MongoManager) FindOneAndReplace(filter interface{}, replacement interface{}, doc interface{}) error {
	if m.curCol == nil {
		return errors.New("use collection first")
	}
	ops := options.FindOneAndReplace()
	ops.SetUpsert(true)
	res := m.curCol.FindOneAndReplace(m.ctx, filter, replacement, ops)
	if res == nil || res.Err() != nil {
		return fmt.Errorf("find one and replace document failed: %v", res.Err())
	}
	err := res.Decode(doc)
	if err != nil {
		return fmt.Errorf("find one and replace document decode failed: %v", res.Err())
	}
	return nil
}

// WatchDatabase 监听数据库变化
func (m MongoManager) WatchDatabase(pipeline interface{}, closeNow bool, doc interface{}, docChange func(interface{}, error)) {

	if m.curDB == nil {
		docChange(nil, errors.New("use database first"))
		return
	}
	res, err := m.curDB.Watch(m.ctx, pipeline)
	if err != nil {
		docChange(nil, fmt.Errorf("watch database failed: %v", err))
		return
	}
	if closeNow == true {
		defer res.Close(m.ctx)
	}
	if res.Err() != nil {
		docChange(nil, fmt.Errorf("watch database failed: %v", res.Err()))
		return
	}

	for {
		if res.Next(m.ctx) {
			docChange(res.Decode(doc), nil)
		} else {
			break
		}
	}
}

// WatchCollection 监听集合变化
func (m MongoManager) WatchCollection(pipeline interface{}, closeNow bool, doc interface{}, docChange func(interface{}, error)) {
	if m.curCol == nil {
		docChange(nil, errors.New("use collection first"))
		return
	}
	res, err := m.curCol.Watch(m.ctx, pipeline)
	if err != nil {
		docChange(nil, fmt.Errorf("watch collection failed: %v", err))
		return
	}
	if closeNow == true {
		defer res.Close(m.ctx)
	}
	if res.Err() != nil {
		docChange(nil, fmt.Errorf("watch collection failed: %v", res.Err()))
		return
	}
	for {
		if res.Next(m.ctx) {
			docChange(res.Decode(doc), nil)
		} else {
			break
		}
	}
}
