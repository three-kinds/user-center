package reset_password_code_service

import (
	"github.com/google/uuid"
	"github.com/three-kinds/user-center/daos"
	"github.com/three-kinds/user-center/services/bo"
	"github.com/three-kinds/user-center/utils/service_utils/se"
	"time"
)

type ResetPasswordCodeServiceImpl struct {
	dao daos.IResetPasswordCodeDAO
}

func (s *ResetPasswordCodeServiceImpl) CreateCode(userID int64) (*bo.ResetPasswordCodeBo, error) {
	err := s.dao.DeleteExpiredCode()
	if err != nil {
		return nil, err
	}

	count, err := s.dao.CountCode(userID)
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, se.ValidationError("the reset password code has send, please check the email")
	}

	return s.dao.CreateCode(uuid.New().String(), userID, time.Now().Add(time.Hour*time.Duration(1)))
}

func (s *ResetPasswordCodeServiceImpl) ValidateCode(codeKey string) (int64, error) {
	err := s.dao.DeleteExpiredCode()
	if err != nil {
		return 0, err
	}

	code, err := s.dao.GetCodeByKey(codeKey)
	if err != nil {
		return 0, err
	}

	err = s.dao.DeleteCodeByKey(codeKey)
	if err != nil {
		return 0, err
	}

	return code.UserID, nil
}

func NewResetPasswordCodeServiceImpl(dao daos.IResetPasswordCodeDAO) *ResetPasswordCodeServiceImpl {
	return &ResetPasswordCodeServiceImpl{dao: dao}
}
