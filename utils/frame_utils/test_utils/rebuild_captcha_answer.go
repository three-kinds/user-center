package test_utils

import (
	"encoding/json"
	"fmt"
	goCaptcha "github.com/wenlng/go-captcha/captcha"
	"strings"
)

func RebuildCaptchaAnswer(answer *string) *string {
	rightAnswer := map[int]goCaptcha.CharDot{}
	_ = json.Unmarshal([]byte(*answer), &rightAnswer)

	dotList := make([]int, len(rightAnswer)*2)
	for i, dot := range rightAnswer {
		dotList[i*2] = dot.Dx
		dotList[i*2+1] = dot.Dy
	}

	target := strings.Trim(strings.Replace(fmt.Sprint(dotList), " ", ",", -1), "[]")
	return &target
}
