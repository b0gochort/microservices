package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/b0gochort/microservices/internal/model"
	"github.com/b0gochort/microservices/pkg/customerrors"
	"github.com/b0gochort/microservices/pkg/utils"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

func makeGetRequest(resServise interface{}, URL string) error {
	resp, err := http.Get(URL)
	if err != nil {
		return errors.Wrap(&customerrors.ErrRequest, "get user cars")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Wrap(&customerrors.ErrRequest, "status code")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(&customerrors.ErrInternal, "readAll")
	}

	if err := json.Unmarshal(data, &resServise); err != nil {
		return errors.Wrap(&customerrors.ErrInternal, "unmarshal response")
	}

	return nil
}

// api/users/id/cars
func GetUserCars(c echo.Context) (err error) {
	var (
		req        model.GetUserCarsReq
		resService model.HTTPResponse
	)
	defer func() {
		c.JSON(utils.HTTPResponse(err, resService.Data))
	}()

	req.UserId, err = strconv.Atoi(c.Param("userID"))
	if err != nil {
		return errors.Wrap(&customerrors.ErrValidate, "convert string to integer")
	}

	if err = req.Validate(); err != nil {
		return errors.Wrap(&customerrors.ErrValidate, "validate user id")
	}

	if err = makeGetRequest(&resService, fmt.Sprintf("http://%s:%s/api/user-service/users/%d/exists", viper.GetString("service.userservice.hostname"), viper.GetString("service.userservice.port"), req.UserId)); err != nil {
		return errors.Wrap(&customerrors.ErrRequest, "get user exists")
	}

	if err = makeGetRequest(&resService, fmt.Sprintf("http://%s:%s/api/cars-service/users/%d/cars", viper.GetString("service.carservice.hostname"), viper.GetString("service.carservice.port"), req.UserId)); err != nil {
		return errors.Wrap(&customerrors.ErrRequest, "get user cars")
	}

	return
}

// api/users/id/engines
func GetUserEngines(c echo.Context) (err error) {
	var (
		req           model.GetUserEnginesReq
		resService    model.HTTPResponse
		resCarService model.GetidsUserCarsRes
	)
	defer func() {
		c.JSON(utils.HTTPResponse(err, resService.Data))
	}()

	req.UserId, err = strconv.Atoi(c.Param("userID"))
	if err != nil {
		return errors.Wrap(&customerrors.ErrValidate, "convert string to integer")
	}

	if err = req.Validate(); err != nil {
		return errors.Wrap(&customerrors.ErrValidate, "validate user id")
	}

	if err = makeGetRequest(&resService, fmt.Sprintf("http://%s:%s/api/user-service/users/%d/exists", viper.GetString("service.userservice.hostname"), viper.GetString("service.userservice.port"), req.UserId)); err != nil {
		return errors.Wrap(&customerrors.ErrRequest, "get user exists")
	}

	if err = makeGetRequest(&resCarService, fmt.Sprintf("http://%s:%s/api/cars-service/users/%d/carsid", viper.GetString("service.carservice.hostname"), viper.GetString("service.carservice.port"), req.UserId)); err != nil {
		return errors.Wrap(&customerrors.ErrRequest, "get user cars")
	}

	engineRequest := model.GetEnginesByidReq{
		Carsid: resCarService.Data.Carsid,
	}

	engineRequestJSON, err := json.Marshal(engineRequest)
	if err != nil {
		return errors.Wrap(&customerrors.ErrUnMarshal, "marshal req id engines")
	}

	if err = makePostReq(&resService, fmt.Sprintf("http://%s:%s/api/engine-service/engines",
		viper.GetString("service.engineservice.hostname"), viper.GetString("service.engineservice.port")),
		engineRequestJSON,
	); err != nil {
		return errors.Wrap(&customerrors.ErrGetEnginesByids, "get engines by id")
	}

	return
}

// /api/engines/brand
func GetEnginesByBrand(c echo.Context) (err error) {
	var (
		req           model.GetBrandCarsidReq
		resService    model.HTTPResponse
		resCarService model.GetidsUserCarsRes
	)

	defer func() {
		c.JSON(utils.HTTPResponse(err, resService.Data))
	}()

	req.Brand = c.Param("brand")

	if err = req.Validate(); err != nil {
		return errors.Wrap(&customerrors.ErrValidate, "validate")
	}

	if err = makeGetRequest(&resCarService,
		fmt.Sprintf("http://%s:%s/api/cars-service/cars/%s", viper.GetString("service.carservice.hostname"), viper.GetString("service.carservice.port"), req.Brand)); err != nil {
		return errors.Wrap(&customerrors.ErrRequest, "get user cars")
	}

	engineRequest := model.GetEnginesByidReq{
		Carsid: resCarService.Data.Carsid,
	}

	engineRequestJSON, err := json.Marshal(engineRequest)
	if err != nil {
		return errors.Wrap(&customerrors.ErrUnMarshal, "marshal req id engines")
	}

	if err = makePostReq(&resService, fmt.Sprintf("http://%s:%s/api/engine-service/engines",
		viper.GetString("service.engineservice.hostname"), viper.GetString("service.engineservice.port")),
		engineRequestJSON,
	); err != nil {
		return errors.Wrap(&customerrors.ErrGetEnginesByids, "get engines by id")
	}

	return
}

// /api/engines/brand/model
func GetEnginesByModel(c echo.Context) (err error) {
	var (
		req           model.GetModelCarsidReq
		resService    model.HTTPResponse
		resCarService model.GetidsUserCarsRes
	)

	defer func() {
		c.JSON(utils.HTTPResponse(err, resService.Data))
	}()

	req.Brand = c.Param("brand")
	req.Model = c.Param("model")

	if err = req.Validate(); err != nil {
		return errors.Wrap(&customerrors.ErrValidate, "validate")
	}

	if err = makeGetRequest(&resCarService, fmt.Sprintf("http://%s:%s/api/cars-service/cars/%s/%s", viper.GetString("service.carservice.hostname"), viper.GetString("service.carservice.port"), req.Brand, req.Model)); err != nil {
		return errors.Wrap(&customerrors.ErrRequest, "get user cars")
	}

	engineRequest := model.GetEnginesByidReq{
		Carsid: resCarService.Data.Carsid,
	}

	engineRequestJSON, err := json.Marshal(engineRequest)
	if err != nil {
		return errors.Wrap(&customerrors.ErrUnMarshal, "marshal req id engines")
	}

	if err = makePostReq(&resService, fmt.Sprintf("http://%s:%s/api/engine-service/engines",
		viper.GetString("service.engineservice.hostname"), viper.GetString("service.engineservice.port")),
		engineRequestJSON,
	); err != nil {
		return errors.Wrap(&customerrors.ErrGetEnginesByids, "get engines by id")
	}

	return
}

func makePostReq(resService *model.HTTPResponse, URL string, req []byte) error {
	client := &http.Client{}

	request, err := http.NewRequest("POST", URL, bytes.NewBuffer(req))
	if err != nil {
		return errors.Wrap(&customerrors.ErrRequest, "new requets to engine services")
	}

	request.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(request)
	if err != nil {
		return errors.Wrap(&customerrors.ErrRequest, "get user engines")
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(&customerrors.ErrInternal, "readAll")
	}

	if err := json.Unmarshal(data, &resService); err != nil {
		return errors.Wrap(&customerrors.ErrInternal, "unmarshal response")
	}

	return nil
}
