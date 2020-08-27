package test

import (
	"reflect"
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
}
