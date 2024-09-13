package db_test

import (
	"NewsBack/internal/db"
	"testing"
)

var configDB = "host=localhost user=postgres password=passwordtest dbname=Todos port=5432 sslmode=disable TimeZone=Asia/Shanghai"
var configDBBad = "host=localhost Router=postgres dbname=Todos_not_exists port=1111 sslmode=disable TimeZone=Asia/Shanghai"

func TestConnect_Success(t *testing.T) {
	dbObject, err := db.Connect(configDB)
	if err != nil {
		t.Errorf("expected '%q' but got '%v, %d'", "nil", dbObject, err)
	}
}

func TestConnect_Fail(t *testing.T) {
	dbObject, err := db.Connect(configDBBad)
	if err == nil {
		t.Errorf("expected '%q' but got '%v, %d'", "error", dbObject, err)
	}
}
