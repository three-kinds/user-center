package di

import "github.com/three-kinds/user-center/utils/service_utils/email_util"

func NewEmailUtil() email_util.IEmailUtil {
	return email_util.NewEmailUtilImpl()
}
