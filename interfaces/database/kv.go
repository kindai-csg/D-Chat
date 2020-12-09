package database

// キーバリュー型を扱いたいときに使うやつ
// 配列を使いたいときは[]interface{}で
type KV struct {
	Key   string
	Value interface{}
}
