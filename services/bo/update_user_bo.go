package bo

type UpdateUserBO struct {
	Email       *string
	Username    *string
	Password    *string
	IsActive    *bool
	IsStaff     *bool
	IsSuperuser *bool
	Nickname    *string
	PhoneNumber *string
	Avatar      *string
}
