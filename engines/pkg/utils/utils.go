package utils

import (
	"fmt"
	"net/http"

	"github.com/b0gochort/microservices/internal/model"
	"github.com/b0gochort/microservices/pkg/customerrors"
	"github.com/pkg/errors"
)

func HTTPResponse(err error, res any) (int, model.HTTPResponse) {
	var httpRes model.HTTPResponse

	if err == nil {
		httpRes.Data = res
		return http.StatusOK, httpRes
	}

	var e *customerrors.HttpErr

	if errors.As(err, &e) {
		httpRes.Err = e.Message

		return e.Code, httpRes
	}
	httpRes.Err = customerrors.ErrInternal.Message

	return http.StatusInternalServerError, httpRes
}

func IntSliceToString(slice []int) string {
	result := ""
	for i, val := range slice {
		if i > 0 {
			result += ","
		}
		result += fmt.Sprintf("%d", val)
	}
	return result
}
