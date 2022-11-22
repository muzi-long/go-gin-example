package mysql

import "testing"

func TestNew(t *testing.T) {
	cfg := &Config{
		Host:            "124.221.59.129",
		Port:            3306,
		Username:        "opensips",
		Password:        "opensipsrw",
		Dbname:          "opensips",
		MaxIdleConn:     0,
		MaxOpenConn:     0,
		ConnMaxLifeTime: 0,
	}

	db, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("success", db)
}
