package se

import "fmt"

type ServiceError struct {
	Code    int
	Status  string
	Message string
	Cause   string
	Detail  map[string]string
}

func (e *ServiceError) Error() string {
	return fmt.Sprintf("%s;%s", e.Message, e.Cause)
}
