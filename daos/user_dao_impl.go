package daos

import (
	"fmt"
	"github.com/three-kinds/user-center/daos/models"
	"github.com/three-kinds/user-center/services/bo"
	"github.com/three-kinds/user-center/utils/generic_utils/dynamic_utils"
	"github.com/three-kinds/user-center/utils/generic_utils/gorm_addons"
	"github.com/three-kinds/user-center/utils/service_utils/se"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func tCreateUserBO2User(bo *bo.CreateUserBO) *models.User {
	return &models.User{
		Email:       bo.Email,
		Username:    bo.Username,
		Password:    bo.Password,
		IsActive:    true,
		IsSuperuser: bo.IsSuperuser,
	}
}

func tUser2UserDisplayBO(user *models.User) *bo.UserDisplayBO {
	return &bo.UserDisplayBO{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		Nickname:    user.Nickname,
		PhoneNumber: user.PhoneNumber,
		Avatar:      user.Avatar,
		DateJoined:  user.DateJoined,
		LastLogin:   user.LastLogin,
		IsActive:    user.IsActive,
		IsSuperuser: user.IsSuperuser,
	}
}

type UserDAOImpl struct {
	db gorm_addons.IDB
}

func (dao *UserDAOImpl) CreateUser(c *bo.CreateUserBO, id int64, dateJoined time.Time) (*bo.UserDisplayBO, error) {
	newUser := tCreateUserBO2User(c)
	newUser.ID = id
	newUser.DateJoined = dateJoined
	result := dao.db.Create(newUser)
	if result.Error != nil {
		cause := result.Error.Error()
		if strings.Contains(cause, "duplicate key value violates unique") {
			target := regexp.MustCompile("\"(.*)\"").FindStringSubmatch(cause)
			vl := strings.SplitN(target[1], "_", 3)
			return nil, se.ValidationError(fmt.Sprintf("duplicate key: %s", vl[2]))
		} else {
			return nil, se.ServerKnownError(fmt.Sprintf("create user error: %s", cause))
		}
	}
	return tUser2UserDisplayBO(newUser), nil
}

func (dao *UserDAOImpl) UpdatePassword(id int64, password string) error {
	result := dao.db.Model(&models.User{}).Where("id = ?", id).Update("password", password)
	if result.Error != nil {
		return se.ServerKnownError(fmt.Sprintf("update user password error: %s", result.Error))
	}
	return nil
}

func (dao *UserDAOImpl) CheckPassword(id int64, password string) (bool, error) {
	user, err := dao.getRawUserByUniqueField("id", strconv.FormatInt(id, 10))
	if err != nil {
		return false, err
	}
	return user.Password == password, nil
}

func (dao *UserDAOImpl) UpdateUser(id int64, updateUserBO *bo.UpdateUserBO) error {
	updatedFields := dynamic_utils.OptionalStructFieldsToMap(updateUserBO)
	if len(updatedFields) == 0 {
		return nil
	}
	result := dao.db.Model(&models.User{}).Where("id = ?", id).Updates(updatedFields)
	if result.Error != nil {
		return se.ServerKnownError(fmt.Sprintf("update user error: %s", result.Error))
	}
	return nil
}

func (dao *UserDAOImpl) Count(isActive *bool, isSuperuser *bool) (total int64, err error) {
	query := dao.db.Model(&models.User{})
	if isActive != nil {
		query.Where("is_active = ?", *isActive)
	}
	if isSuperuser != nil {
		query.Where("is_superuser = ?", *isSuperuser)
	}
	result := query.Count(&total)
	if result.Error != nil {
		err = se.ServerKnownError(fmt.Sprintf("query user count error: %s", result.Error))
		return
	}
	return
}

func (dao *UserDAOImpl) ListUsers(page int, size int, isActive *bool, isSuperuser *bool) ([]*bo.UserDisplayBO, error) {
	limit := size
	offset := (page - 1) * size
	query := dao.db.Model(&models.User{}).Limit(limit).Offset(offset)
	if isActive != nil {
		query.Where("is_active = ?", *isActive)
	}
	if isSuperuser != nil {
		query.Where("is_superuser = ?", *isSuperuser)
	}

	var users []models.User
	result := query.Find(&users)
	if result.Error != nil {
		err := se.ServerKnownError(fmt.Sprintf("find users error: %s", result.Error))
		return nil, err
	}
	udList := make([]*bo.UserDisplayBO, len(users))
	for i, user := range users {
		udList[i] = tUser2UserDisplayBO(&user)
	}

	return udList, nil
}

func (dao *UserDAOImpl) getRawUserByUniqueField(field, value string) (*models.User, error) {
	user := models.User{}
	result := dao.db.First(&user, fmt.Sprintf("%s =  ?", field), value)
	if result.Error != nil {
		return nil, se.NotFoundError(fmt.Sprintf("not found user: %s", value))
	}
	return &user, nil
}

func (dao *UserDAOImpl) getUserByUniqueField(field, value string) (*bo.UserDisplayBO, error) {
	user, err := dao.getRawUserByUniqueField(field, value)
	if err != nil {
		return nil, err
	}
	return tUser2UserDisplayBO(user), nil
}

func (dao *UserDAOImpl) GetUserByID(id int64) (*bo.UserDisplayBO, error) {
	return dao.getUserByUniqueField("id", strconv.FormatInt(id, 10))
}

func (dao *UserDAOImpl) GetUserByUsername(username string) (*bo.UserDisplayBO, error) {
	return dao.getUserByUniqueField("username", username)
}

func (dao *UserDAOImpl) GetUserByEmail(email string) (*bo.UserDisplayBO, error) {
	return dao.getUserByUniqueField("email", email)
}

func (dao *UserDAOImpl) GetUserByPhoneNumber(number string) (*bo.UserDisplayBO, error) {
	return dao.getUserByUniqueField("phone_number", number)
}

func (dao *UserDAOImpl) DeleteUserByID(id int64) error {
	result := dao.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return se.ServerKnownError(fmt.Sprintf("delete user error: %s", result.Error))
	}
	return nil
}

func NewUserDAOImpl(db gorm_addons.IDB) *UserDAOImpl {
	return &UserDAOImpl{db: db}
}
