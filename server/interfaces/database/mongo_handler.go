package database

type MongoHandler interface {
	CreateIndex(string, []KV, []KV) error
	Insert(string, []KV) (string, error)
	Find(string, []KV) ([][]KV, error)
	FindOne(string, []KV) ([]KV, error)
	Update(string, []KV, []KV) error
	Delete(string, []KV) error
}
