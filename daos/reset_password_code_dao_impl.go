package daos

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/three-kinds/user-center/daos/models"
	"github.com/three-kinds/user-center/services/reset_password_code_service"
	"github.com/three-kinds/user-center/utils/generic_utils/gorm_addons"
	"github.com/three-kinds/user-center/utils/service_utils/se"
	"time"
)

type ResetPasswordCodeDAOImpl struct {
	db     gorm_addons.IDB
	logger *logrus.Entry
}

func tCode2CodeBo(code *models.ResetPasswordCode) *reset_password_code_service.ResetPasswordCodeBo {
	return &reset_password_code_service.ResetPasswordCodeBo{
		Key:        code.Key,
		UserID:     code.UserID,
		Expiration: code.Expiration,
	}
}

func (dao *ResetPasswordCodeDAOImpl) CreateCode(key string, userID int64, expiration time.Time) (*reset_password_code_service.ResetPasswordCodeBo, error) {
	code := &models.ResetPasswordCode{
		Key:        key,
		UserID:     userID,
		Expiration: expiration,
	}
	result := dao.db.Create(code)
	if result.Error != nil {
		return nil, se.ServerKnownError(fmt.Sprintf("create code error: %s", result.Error))
	}
	return tCode2CodeBo(code), nil
}

func (dao *ResetPasswordCodeDAOImpl) GetCodeByKey(key string) (*reset_password_code_service.ResetPasswordCodeBo, error) {
	code := &models.ResetPasswordCode{}
	result := dao.db.Where("key = ?", key).First(code)
	if result.Error != nil {
		return nil, se.NotFoundError("not found code", se.Cause(result.Error.Error()))
	}
	return tCode2CodeBo(code), nil
}

func (dao *ResetPasswordCodeDAOImpl) DeleteExpiredCode() error {
	result := dao.db.Where("expiration <= ?", time.Now()).Delete(&models.ResetPasswordCode{})
	if result.Error != nil {
		return se.ServerKnownError(fmt.Sprintf("delete expired code failed: %s", result.Error))
	}
	if result.RowsAffected > 0 {
		dao.logger.Debug("DeleteExpiredCode, count:", result.RowsAffected)
	}
	return nil
}

func (dao *ResetPasswordCodeDAOImpl) DeleteCodeByKey(key string) error {
	result := dao.db.Where("key = ?", key).Delete(&models.ResetPasswordCode{})
	if result.Error != nil {
		return se.ServerKnownError(fmt.Sprintf("delete code error: %s", result.Error))
	}
	return nil
}

func NewResetPasswordCodeDAOImpl(db gorm_addons.IDB) *ResetPasswordCodeDAOImpl {
	return &ResetPasswordCodeDAOImpl{
		db: db,
		logger: logrus.WithFields(logrus.Fields{
			"name": "CodeDAOImpl",
		}),
	}
}