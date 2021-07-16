package mysql

import (
	"gin-web/config"
	"testing"
)

func setup() {
	cfg := config.MysqlConfig{}
	if err := Init(cfg); err != nil {
		//panic(err)
	}
}

func teardown() {
	//_ = db.Close()
}

func TestMain(m *testing.M) {
	setup()
	m.Run()
	teardown()
}

func TestCreatePost(t *testing.T) {
	//t.Fatal()
}

func TestCreatePost2(t *testing.T) {

}
