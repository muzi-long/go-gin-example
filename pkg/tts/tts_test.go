package tts

import "testing"

func TestNew(t *testing.T) {
	c := &Config{
		AccessKeyId:     "",
		AccessKeySecret: "",
		AppKey:          "",
	}
	client := c.New()
	b, err := client.Run("今天天气真不错")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(b)
}
