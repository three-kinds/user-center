package captcha_service

import (
	"encoding/json"
	"fmt"
	"github.com/three-kinds/user-center/daos"
	"github.com/three-kinds/user-center/utils/service_utils/se"
	"strconv"
	"strings"
	"time"
)
import goCaptcha "github.com/wenlng/go-captcha/captcha"

type CaptchaServiceImpl struct {
	dao daos.ICaptchaDAO
}

func (s *CaptchaServiceImpl) GetCaptcha() (b64 string, tb64 string, key string, err error) {
	err = s.dao.DeleteExpiredCaptcha()
	if err != nil {
		return
	}

	capt := goCaptcha.GetCaptcha()
	dots, b64, tb64, key, err := capt.Generate()
	if err != nil {
		err = se.ServerKnownError(fmt.Sprintf("go-captcha generate error: %s", err))
		return
	}

	answer, err := json.Marshal(dots)
	if err != nil {
		err = se.ServerKnownError(fmt.Sprintf("go-captcha dots cannot json marshal: %s", err))
		return
	}

	expiration := time.Now().Add(time.Hour * time.Duration(1))
	_, err = s.dao.CreateCaptcha(key, string(answer), expiration)
	if err != nil {
		return
	}

	return
}

func (s *CaptchaServiceImpl) ValidateCaptcha(key string, answer string) (bool, error) {
	err := s.dao.DeleteExpiredCaptcha()
	if err != nil {
		return false, err
	}

	captcha, err := s.dao.GetCaptchaByKey(key)
	if err != nil {
		return false, err
	}

	rightAnswer := map[int]goCaptcha.CharDot{}
	err = json.Unmarshal([]byte(captcha.Answer), &rightAnswer)
	if err != nil {
		return false, se.ServerKnownError("captcha.answer cannot json unmarshal")
	}

	src := strings.Split(answer, ",")
	chkRet := false
	if (len(rightAnswer) * 2) == len(src) {
		for i, dot := range rightAnswer {
			j := i * 2
			k := i*2 + 1
			sx, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[j]), 64)
			sy, _ := strconv.ParseFloat(fmt.Sprintf("%v", src[k]), 64)

			// 检测点位置
			// chkRet = captcha.CheckPointDist(int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height))

			// 校验点的位置,在原有的区域上添加额外边距进行扩张计算区域,不推荐设置过大的padding
			// 例如：文本的宽和高为30，校验范围x为10-40，y为15-45，此时扩充5像素后校验范围宽和高为40，则校验范围x为5-45，位置y为10-50
			chkRet = goCaptcha.CheckPointDistWithPadding(
				int64(sx), int64(sy), int64(dot.Dx), int64(dot.Dy), int64(dot.Width), int64(dot.Height), 5,
			)
			if !chkRet {
				break
			}
		}
	}

	err = s.dao.DeleteCaptchaByKey(key)
	if err != nil {
		return false, err
	}
	return chkRet, nil
}

func NewCaptchaServiceImpl(dao daos.ICaptchaDAO) *CaptchaServiceImpl {
	return &CaptchaServiceImpl{dao: dao}
}
