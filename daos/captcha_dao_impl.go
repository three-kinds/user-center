package daos

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/three-kinds/user-center/daos/models"
	"github.com/three-kinds/user-center/services/bo"
	"github.com/three-kinds/user-center/utils/generic_utils/gorm_addons"
	"github.com/three-kinds/user-center/utils/service_utils/se"
	"time"
)

type CaptchaDAOImpl struct {
	db     gorm_addons.IDB
	logger *logrus.Entry
}

func tCaptcha2CaptchaDisplayBo(captcha *models.Captcha) *bo.CaptchaBo {
	return &bo.CaptchaBo{
		Key:        captcha.Key,
		Answer:     captcha.Answer,
		Expiration: captcha.Expiration,
	}
}

func (dao *CaptchaDAOImpl) CreateCaptcha(key string, answer string, expiration time.Time) (*bo.CaptchaBo, error) {
	captcha := &models.Captcha{
		Key:        key,
		Answer:     answer,
		Expiration: expiration,
	}
	result := dao.db.Create(captcha)
	if result.Error != nil {
		return nil, se.ServerKnownError(fmt.Sprintf("create captcha error: %s", result.Error))
	}
	return tCaptcha2CaptchaDisplayBo(captcha), nil
}

func (dao *CaptchaDAOImpl) GetCaptchaByKey(key string) (*bo.CaptchaBo, error) {
	captcha := &models.Captcha{}
	result := dao.db.Where("key = ?", key).First(captcha)
	if result.Error != nil {
		return nil, se.NotFoundError("not found captcha", se.Cause(result.Error.Error()))
	}
	return tCaptcha2CaptchaDisplayBo(captcha), nil
}

func (dao *CaptchaDAOImpl) DeleteExpiredCaptcha() error {
	result := dao.db.Where("expiration <= ?", time.Now()).Delete(&models.Captcha{})
	if result.Error != nil {
		return se.ServerKnownError(fmt.Sprintf("delete expired captcha failed: %s", result.Error))
	}
	if result.RowsAffected > 0 {
		dao.logger.Debug("DeleteExpiredCaptcha, count:", result.RowsAffected)
	}
	return nil

}

func (dao *CaptchaDAOImpl) DeleteCaptchaByKey(key string) error {
	result := dao.db.Where("key = ?", key).Delete(&models.Captcha{})
	if result.Error != nil {
		return se.ServerKnownError(fmt.Sprintf("delete captcha error: %s", result.Error))
	}
	return nil
}

func NewCaptchaDAOImpl(db gorm_addons.IDB) *CaptchaDAOImpl {
	return &CaptchaDAOImpl{
		db: db,
		logger: logrus.WithFields(logrus.Fields{
			"name": "CaptchaDAOImpl",
		}),
	}
}
