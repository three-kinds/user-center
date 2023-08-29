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

type ResetPasswordCodeDAOImpl struct {
	db     gorm_addons.IDB
	logger *logrus.Entry
}

func tCode2CodeBo(code *models.ResetPasswordCode) *bo.ResetPasswordCodeBo {
	return &bo.ResetPasswordCodeBo{
		Key:        code.Key,
		UserID:     code.UserID,
		Expiration: code.Expiration,
	}
}

func (dao *ResetPasswordCodeDAOImpl) CreateCode(key string, userID int64, expiration time.Time) (*bo.ResetPasswordCodeBo, error) {
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

func (dao *ResetPasswordCodeDAOImpl) GetCodeByKey(key string) (*bo.ResetPasswordCodeBo, error) {
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

func (dao *ResetPasswordCodeDAOImpl) CountCode(userID int64) (int64, error) {
	var count int64
	result := dao.db.Model(&models.ResetPasswordCode{}).Where("user_id = ?", userID).Count(&count)
	if result.Error != nil {
		return 0, se.ServerKnownError(fmt.Sprintf("count code error: %s", result.Error))
	}
	return count, nil
}

func NewResetPasswordCodeDAOImpl(db gorm_addons.IDB) *ResetPasswordCodeDAOImpl {
	return &ResetPasswordCodeDAOImpl{
		db: db,
		logger: logrus.WithFields(logrus.Fields{
			"name": "CodeDAOImpl",
		}),
	}
}
