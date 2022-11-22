package redis

import "testing"

func TestNew(t *testing.T) {
	cfg := &Config{
		Address:  []string{"127.0.0.1:6379"},
		Password: "xx:xxx@xxx",
		Db:       0,
	}
	rdb, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success", rdb)
}

func TestNewCluster(t *testing.T) {
	cfg := &Config{
		Address:  []string{"192.168.0.1:6379", "192.168.0.2:6379"},
		Password: "xx:xxx@xxx",
		Db:       0,
	}
	rdb, err := NewCluster(cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success", rdb)
}
