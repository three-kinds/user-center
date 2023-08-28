package gin_utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/three-kinds/user-center/initializers"
	"github.com/three-kinds/user-center/utils/generic_utils/case_utils"
	"github.com/three-kinds/user-center/utils/service_utils/se"
	"strings"
)

func trimStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		key := field[strings.Index(field, ".")+1:]
		key = case_utils.CamelCaseToSnakeCase(key)
		res[key] = err
	}

	return res
}

func getValidationErrorsDetail(ctx *gin.Context, errs validator.ValidationErrors) (map[string]string, error) {
	locale := ctx.GetHeader("locale")
	if locale == "" {
		locale = "zh"
	}

	translator, ok := initializers.UT.GetTranslator(locale)
	if !ok {
		return nil, fmt.Errorf(`GetTranslator("%s") error`, locale)
	}

	detail := trimStruct(errs.Translate(translator))
	return detail, nil
}

func AbortWithValidationError(ctx *gin.Context, err error) {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		AbortWithError(ctx, se.ValidationError(err.Error()))
		return
	}

	detail, err := getValidationErrorsDetail(ctx, errs)
	if err != nil {
		AbortWithError(ctx, se.ServerKnownError(err.Error()))
		return
	}

	var firstErrorMessage string
	for _, errorMessage := range detail {
		firstErrorMessage = errorMessage
		break
	}
	serviceError := se.ValidationError(firstErrorMessage, se.Detail(detail))
	_ = ctx.Error(serviceError)
	ctx.AbortWithStatusJSON(serviceError.Code, gin.H{
		"status":  serviceError.Status,
		"message": firstErrorMessage,
		"detail":  detail,
	})
}
