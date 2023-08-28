package email_util

type IEmailUtil interface {
	SendResetPasswordEmail(email string, codeKey string) error
}
