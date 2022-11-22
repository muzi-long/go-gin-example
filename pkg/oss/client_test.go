package oss

import (
	"strings"
	"testing"
)

func TestClient_PutObject(t *testing.T) {
	cfg := &Config{
		AccessKeyId:     "",
		AccessKeySecret: "",
		Endpoint:        "",
		Bucket:          "",
	}
	client, err := New(cfg)
	if err != nil {
		t.Log(err)
		return
	}
	err = client.PutObject("c.log", strings.NewReader("hello world c"))
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("success")
}

func TestClient_PutObjectFromFile(t *testing.T) {
	cfg := &Config{
		AccessKeyId:     "",
		AccessKeySecret: "",
		Endpoint:        "",
		Bucket:          "",
	}
	client, err := New(cfg)
	if err != nil {
		t.Log(err)
		return
	}
	err = client.PutObjectFromFile("b.log", "./test.log")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("success")
}

func TestClient_GetUrlByFile(t *testing.T) {
	cfg := &Config{
		AccessKeyId:     "",
		AccessKeySecret: "",
		Endpoint:        "",
		Bucket:          "",
	}
	client, err := New(cfg)
	if err != nil {
		t.Log(err)
		return
	}
	remoteFile := "aaaaa.log"
	err = client.PutObject(remoteFile, strings.NewReader("hello world c"))
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(client.GetUrlByFile(remoteFile))
}

func TestClient_PutObjectMulFromFile(t *testing.T) {
	cfg := &Config{
		AccessKeyId:     "",
		AccessKeySecret: "",
		Endpoint:        "",
		Bucket:          "",
	}
	client, err := New(cfg)
	if err != nil {
		t.Log(err)
		return
	}
	err = client.PutObjectMulFromFile("aa.wav", "./aa.wav")
	if err != nil {
		t.Log(err)
		return
	}
	t.Log("success")
}
