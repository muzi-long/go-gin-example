package redis

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	ctx := context.TODO()
	cfg := &Config{
		Address:  []string{"127.0.0.1:6379"},
		Password: "root",
		Db:       0,
	}
	rdb, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	rdb.Set(ctx, "aa", "111", time.Hour)
	result := rdb.Get(ctx, "aa").Val()
	fmt.Println("result", result)
	assert.Equal(t, "111", result)

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
