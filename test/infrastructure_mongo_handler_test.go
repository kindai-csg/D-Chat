package test

import (
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/kindai-csg/D-Chat/infrastructure"
	"github.com/kindai-csg/D-Chat/interfaces/database"
)

// MongoHandlerの生成
// 接続テストも兼ねてる
func newMongoHandler(t *testing.T) *infrastructure.MongoHandler {
	mongoHandler, err := infrastructure.NewMongoHandler()
	if err != nil {
		t.Errorf("faild connection to mongodb => " + err.Error())
	}
	return mongoHandler
}

// Test
// MongoDBへのInsertテスト
// []database.KV -> bson.Dへの変換テストも兼ねてる
func TestMongoHandlerInsert(t *testing.T) {
	handler := newMongoHandler(t)
	rand.Seed(time.Now().UnixNano())

	// -------------test1-------------
	doc := []database.KV{
		{"user_id", "test_user"},
		{"name", "test_name"},
		{"password", "aiueo"},
		{"mail", "test@example.com"},
		{"bio", "test_bio"},
		{"status", "test_status"},
		{"status_text", "test_status_text"},
		{"auth", 0},
	}
	// collection_name ex)Test_2020_08_26_TestMongoHandlerInsert1
	id, err := handler.Insert("Test_"+time.Now().Format("2006_01_02")+"_TestMongoHandlerInsert1", doc)
	if err != nil {
		t.Errorf("faild insert to mongodb => " + err.Error())
	}
	if reflect.TypeOf(id).Kind() != reflect.String {
		t.Errorf("Expectation: string")
	}
	t.Logf("id => %s", id)
	// -------------ここまで-------------

	// -------------test2-------------
	doc = []database.KV{
		{"msg", database.KV{"body", "test_msg_body"}},
		{"reaction_count", 0},
	}
	// collection_name ex)Test_2020_08_26_TestMongoHandlerInsert2
	id, err = handler.Insert("Test_"+time.Now().Format("2006_01_02")+"_TestMongoHandlerInsert2", doc)
	if err != nil {
		t.Errorf("faild insert to mongodb => " + err.Error())
	}
	if reflect.TypeOf(id).Kind() != reflect.String {
		t.Errorf("Expectation: string")
	}
	t.Logf("id => %s", id)
	// -------------ここまで-------------

	// -------------test3-------------
	doc = []database.KV{
		{"tags", []interface{}{"test1", "test2", "test3", 1, 3}},
	}
	// collection_name ex)Test_2020_08_26_TestMongoHandlerInsert3
	id, err = handler.Insert("Test_"+time.Now().Format("2006_01_02")+"_TestMongoHandlerInsert3", doc)
	if err != nil {
		t.Errorf("faild insert to mongodb => " + err.Error())
	}
	if reflect.TypeOf(id).Kind() != reflect.String {
		t.Errorf("Expectation: string")
	}
	t.Logf("id => %s", id)
	// -------------ここまで-------------

	// -------------test4-------------
	doc = []database.KV{
		{"room", []database.KV{
			{"users", []interface{}{"user1", "user2", "user3", "user4"}},
			{"msgs", []interface{}{[]database.KV{{"body", "text1"}, {"reaction_count", 2}}, []database.KV{{"body", "text2"}, {"reaction_count", 3}}}},
		}},
	}
	// collection_name ex)Test_2020_08_26_TestMongoHandlerInsert4
	id, err = handler.Insert("Test_"+time.Now().Format("2006_01_02")+"_TestMongoHandlerInsert4", doc)
	if err != nil {
		t.Errorf("faild insert to mongodb => " + err.Error())
	}
	if reflect.TypeOf(id).Kind() != reflect.String {
		t.Errorf("Expectation: string")
	}
	t.Logf("id => %s", id)
	// -------------ここまで-------------

	// -------------test5-------------
	doc = []database.KV{{"_id", strconv.Itoa(rand.Intn(1000))}, {"test", "test"}}
	// collection_name ex)Test_2020_08_26_TestMongoHandlerInsert5
	id, err = handler.Insert("Test_"+time.Now().Format("2006_01_02")+"_TestMongoHandlerInsert5", doc)
	if err != nil {
		t.Errorf("faild insert to mongodb => " + err.Error())
	}
	if reflect.TypeOf(id).Kind() != reflect.String {
		t.Errorf("Expectation: string")
	}
	t.Logf("id => %s", id)
	// -------------ここまで-------------

	// -------------test6-------------
	rand.Seed(time.Now().UnixNano())
	doc = []database.KV{{"_id", rand.Intn(1000)}, {"test", "test"}}
	// collection_name ex)Test_2020_08_26_TestMongoHandlerInsert5
	id, err = handler.Insert("Test_"+time.Now().Format("2006_01_02")+"_TestMongoHandlerInsert6", doc)
	if err != nil {
		t.Errorf("faild insert to mongodb => " + err.Error())
	}
	if reflect.TypeOf(id).Kind() != reflect.String {
		t.Errorf("Expectation: string")
	}
	t.Logf("id => %s", id)
	// -------------ここまで-------------
}

// MongoDBのUpdateテスト
func TestMongoHandlerUpdate(t *testing.T) {
	handler := newMongoHandler(t)
	rand.Seed(time.Now().UnixNano())

	// -------------test1-------------
	doc := []database.KV{
		{"_id", strconv.Itoa(rand.Intn(1000))},
		{"user_id", "test_user"},
		{"password", "test_password"},
	}
	// collection_name ex)Test_2020_08_26_TestMongoHandlerUpdate1
	collectionName := "Test_" + time.Now().Format("2006_01_02") + "_TestMongoHandlerUpdate1"
	id, err := handler.Insert(collectionName, doc)
	if err != nil {
		t.Errorf("faild insert to mongodb => " + err.Error())
	}
	if reflect.TypeOf(id).Kind() != reflect.String {
		t.Errorf("Expectation: string")
	}
	t.Logf("id => %s", id)

	query := []database.KV{
		{"_id", id},
	}
	update := []database.KV{
		{"$set", []database.KV{
			{"user_id", "test_user_update"},
			{"password", "test_password_update"},
		}},
	}
	err = handler.Update(collectionName, query, update)
	if err != nil {
		t.Errorf("faild update document => " + err.Error())
	}
	// -------------ここまで-------------

	// -------------test2-------------
	doc = []database.KV{
		{"count", 100},
		{"name", "test_name"},
	}
	// collection_name ex)Test_2020_08_26_TestMongoHandlerUpdate2
	collectionName = "Test_" + time.Now().Format("2006_01_02") + "_TestMongoHandlerUpdate2"
	id, err = handler.Insert(collectionName, doc)
	if err != nil {
		t.Errorf("faild insert to mongodb => " + err.Error())
	}
	if reflect.TypeOf(id).Kind() != reflect.String {
		t.Errorf("Expectation: string")
	}
	t.Logf("id => %s", id)
	doc = []database.KV{
		{"count", 50},
		{"name", "test_name"},
	}
	id, err = handler.Insert(collectionName, doc)
	if err != nil {
		t.Errorf("faild insert to mongodb => " + err.Error())
	}
	if reflect.TypeOf(id).Kind() != reflect.String {
		t.Errorf("Expectation: string")
	}
	t.Logf("id => %s", id)

	query = []database.KV{
		{"count", []database.KV{
			{"$gte", 100},
		}},
	}
	update = []database.KV{
		{"$set", []database.KV{
			{"name", "test_name_update"},
		}},
	}
	err = handler.Update(collectionName, query, update)
	if err != nil {
		t.Errorf("faild update document => " + err.Error())
	}
	// -------------ここまで-------------
}

// MongoDBのFindテスト
// bson.D -> []database.KVへの変換テストも兼ねてる
func TestMongoHandlerFind(t *testing.T) {
	handler := newMongoHandler(t)
	rand.Seed(time.Now().UnixNano())

	docs := [][]database.KV{
		{
			{"_id", strconv.Itoa(rand.Intn(1000))},
			{"user_id", "test_user"},
			{"password", "test_password"},
		},
		{
			{"count", 0},
			{"tags", []interface{}{"test1", "test2", 3}},
		},
		{
			{"name", "test"},
			{"msg", []database.KV{
				{"count", 2},
				{"body", "test_body"},
			}},
		},
		{
			{"msg", database.KV{"body", "test_body"}},
		},
	}
	querys := [][]database.KV{
		{
			{"user_id", "test_user"},
		},
		{
			{"tags", database.KV{"$exists", true}},
		},
		{
			{"msg.count", database.KV{"$gte", 2}},
		},
		{},
	}
	for i, _ := range docs {
		t.Logf("--------start %d--------", i)
		// collection_name ex)Test_2020_08_26_TestMongoHandlerFind${i}
		collectionName := "Test_" + time.Now().Format("2006_01_02") + "_TestMongoHandlerFind" + strconv.Itoa(i)
		id, err := handler.Insert(collectionName, docs[i])
		if err != nil {
			t.Errorf("faild insert to mongodb => " + err.Error())
		}
		t.Logf("id => %s", id)

		result, err := handler.Find(collectionName, querys[i])
		if err != nil {
			t.Errorf("faild find document => " + err.Error())
		}
		t.Log(result)
		t.Logf("--------end %d--------", i)
	}
}

// MongoDBのCreateIndexテスト
func TestMongoHandlerCreateIndex(t *testing.T) {
	handler := newMongoHandler(t)
	rand.Seed(time.Now().UnixNano())

	doc := []database.KV{
		{"user_id", "test1"},
		{"name", "test2"},
	}
	// collection_name ex)Test_2020_08_26_TestMongoHandlerCreateIndex0
	collectionName := "Test_" + time.Now().Format("2006_01_02") + "_TestMongoHandlerIndex0"
	id, err := handler.Insert(collectionName, doc)
	if err != nil {
		t.Errorf("faild insert to mongodb => " + err.Error())
	}
	t.Logf("id => %s", id)

	index := []database.KV{
		{"user_id", 1},
	}
	opt := []database.KV{
		{"unique", true},
	}
	err = handler.CreateIndex(collectionName, index, opt)
	if err != nil {
		t.Errorf("faild create index => " + err.Error())
	}

	// 重複チェック
	doc = []database.KV{
		{"user_id", "test1"},
		{"name", "test2"},
	}
	_, err = handler.Insert(collectionName, doc)
	if err == nil {
		t.Errorf("faild unique index")
	}
}

// MongoDBのFindOneテスト
func TestMongoHandlerFindOne(t *testing.T) {
	handler := newMongoHandler(t)
	rand.Seed(time.Now().UnixNano())

	r := strconv.Itoa(rand.Intn(1000))
	docs := [][]database.KV{
		{
			{"_id", r},
			{"user_id", "test_user"},
			{"password", "test_password"},
		},
	}
	querys := [][]database.KV{
		{
			{"_id", r},
		},
	}
	for i, _ := range docs {
		t.Logf("--------start %d--------", i)
		// collection_name ex)Test_2020_08_26_TestMongoHandlerFindOne${i}
		collectionName := "Test_" + time.Now().Format("2006_01_02") + "_TestMongoHandlerFindOne" + strconv.Itoa(i)
		id, err := handler.Insert(collectionName, docs[i])
		if err != nil {
			t.Errorf("faild insert to mongodb => " + err.Error())
		}
		t.Logf("id => %s", id)

		result, err := handler.FindOne(collectionName, querys[i])
		if err != nil {
			t.Errorf("faild find document => " + err.Error())
		}
		t.Log(result)
		t.Logf("--------end %d--------", i)
	}
}

// MongoDBのDeleteテスト
func TestMongoHandlerDelete(t *testing.T) {
	handler := newMongoHandler(t)
	rand.Seed(time.Now().UnixNano())

	docs := [][]database.KV{
		{
			{"count", rand.Intn(1000)},
		},
	}
	querys := [][]database.KV{
		{
			{"count", database.KV{"$gte", 500}},
		},
	}
	for i, _ := range docs {
		t.Logf("--------start %d--------", i)
		// collection_name ex)Test_2020_08_26_TestMongoHandlerDelete${i}
		collectionName := "Test_" + time.Now().Format("2006_01_02") + "_TestMongoHandlerDelete" + strconv.Itoa(i)
		id, err := handler.Insert(collectionName, docs[i])
		if err != nil {
			t.Errorf("faild insert to mongodb => " + err.Error())
		}
		t.Logf("id => %s", id)

		result, err := handler.Delete(collectionName, querys[i])
		if err != nil {
			t.Errorf("faild delete document => " + err.Error())
		}
		t.Log(result)
		t.Logf("--------end %d--------", i)
	}
}
