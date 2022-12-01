package ws

import (
	"encoding/json"
)

type MsgI interface {
	Encode() (string, error)
}

type Msg struct {
	To      string `json:"to"`
	Content string `json:"content"`
}

func (m Msg) Encode() (string, error) {
	marshal, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
