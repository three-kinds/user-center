package user_management_service

type CreateUserBO struct {
	Email       string
	Username    string
	Password    string
	IsSuperuser bool
}
