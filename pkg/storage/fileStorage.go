package storage

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson"
)

type FileStorage struct {
	storagePath string
}

func (f *FileStorage) Store(group string, v interface{}) {
	path := filepath.Join(f.storagePath, group)

	bsonV, err := bson.Marshal(v)

	if err != nil {
		log.Fatal(err)
		return
	}
	ioutil.WriteFile(path, bsonV, 0644)
}

func (f *FileStorage) Read(group string, to interface{}) {
	path := filepath.Join(f.storagePath, group)

	bsonV, err := ioutil.ReadFile(path)

	if err != nil {
		log.Println(err)
		return
	}

	err = bson.Unmarshal(bsonV, to)
}

func (f *FileStorage) GetRaw() map[string]interface{} {
	return nil
}

func NewFileStorage(storagePath string) *FileStorage {
	return &FileStorage{
		storagePath: storagePath,
	}
}
