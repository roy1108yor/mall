package e

import "net/http"

func ErrBadRequest() *Error {
	return New(http.StatusBadRequest, "bad request")
}

func ErrUnauthorized() *Error {
	return New(http.StatusUnauthorized, "Unauthorized")
}

func ErrForbidden() *Error {
	return New(http.StatusForbidden, "Forbidden")
}

func ErrInternalServer() *Error {
	return New(http.StatusInternalServerError, "Internal Server Error")
}

func ErrInvalidRequestBody() *Error {
	return New(http.StatusUnprocessableEntity, "字段校验失败, 请查看errs中的提示信息")
}

func ErrNotFound() *Error {
	return New(http.StatusNotFound, "not found")
}
