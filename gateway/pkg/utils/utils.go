package utils

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/b0gochort/microservices/internal/model"
	"github.com/b0gochort/microservices/pkg/customerrors"
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

func IntToStrSlice(idCars []int) []string {
	idStrings := make([]string, len(idCars))
	for i, id := range idCars {
		idStrings[i] = strconv.Itoa(id)
	}
	return idStrings
}
