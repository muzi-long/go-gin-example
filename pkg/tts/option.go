package tts

import nls "github.com/aliyun/alibabacloud-nls-go-sdk"

type (
	// RequestOption 设置自定义请求参数
	RequestOption func(*nls.SpeechSynthesisStartParam)
)

// SetFormat 设置音频格式，wav ｜ mp3
func SetFormat(format string) RequestOption {
	return func(request *nls.SpeechSynthesisStartParam) {
		request.Format = format
	}
}

// SetSampleRate 设置采样率 16000
func SetSampleRate(rate int) RequestOption {
	return func(request *nls.SpeechSynthesisStartParam) {
		request.SampleRate = rate
	}
}

// SetVoice 设置发音人
func SetVoice(name string) RequestOption {
	return func(request *nls.SpeechSynthesisStartParam) {
		request.Voice = name
	}
}

// SetVolume 设置音量,范围是0~100，可选，默认50
func SetVolume(num int) RequestOption {
	return func(request *nls.SpeechSynthesisStartParam) {
		request.Volume = num
	}
}

// SetSpeechRate 设置语速 范围是-500~500，可选，默认是0。
func SetSpeechRate(num int) RequestOption {
	return func(request *nls.SpeechSynthesisStartParam) {
		request.SpeechRate = num
	}
}

// SetPitchRate 设置语调 范围是-500~500，可选，默认是0。
func SetPitchRate(num int) RequestOption {
	return func(request *nls.SpeechSynthesisStartParam) {
		request.PitchRate = num
	}
}
