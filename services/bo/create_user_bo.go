package bo

type CreateUserBO struct {
	Email       string
	Username    string
	Password    string
	IsSuperuser bool
}
