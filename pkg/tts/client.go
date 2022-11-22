package tts

import (
	"errors"
	"log"
	"os"
	"time"

	nls "github.com/aliyun/alibabacloud-nls-go-sdk"
)

type (
	Client struct {
		Config *Config
	}

	// UserParam SDK
	UserParam struct {
		Content []byte //合成后的数据
	}
)

func (c *Config) New() *Client {
	return &Client{
		Config: c,
	}
}

func onTaskFailed(text string, param interface{}) {}

func onSynthesisResult(data []byte, param interface{}) {
	if p, ok := param.(*UserParam); ok {
		p.Content = append(p.Content, data...)
	}
}

func onCompleted(text string, param interface{}) {}

func onClose(param interface{}) {}

func (c *Client) Run(text string, options ...RequestOption) ([]byte, error) {
	if text == "" {
		return nil, errors.New("tts is space")
	}
	param := nls.DefaultSpeechSynthesisParam()
	param.Voice = "siyue"
	for _, option := range options {
		option(&param)
	}
	config, err := nls.NewConnectionConfigWithAKInfoDefault(nls.DEFAULT_URL, c.Config.AppKey, c.Config.AccessKeyId, c.Config.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	userParam := new(UserParam)
	logger := nls.NewNlsLogger(os.Stderr, "tts", log.LstdFlags|log.Lmicroseconds)
	tClient, err := nls.NewSpeechSynthesis(config, logger,
		onTaskFailed, onSynthesisResult, nil,
		onCompleted, onClose, userParam)
	if err != nil {
		return nil, err
	}
	defer tClient.Shutdown()
	ch, err := tClient.Start(text, param, nil)
	if err != nil {
		return nil, err
	}
	select {
	case done := <-ch:
		if !done {
			return nil, errors.New("tts.wait.failed")
		}
	case <-time.After(60 * time.Second):
		return nil, errors.New("tts.wait.timeout")

	}
	return userParam.Content, nil
}
