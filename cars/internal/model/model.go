package model

type GetUserCarsReq struct {
	UserID int `json:"user_id"`
}

type GetUserCarsRes struct {
	Cars []Car `json:"cars"`
}

type GetBrandCarsIDReq struct {
	Brand string `json:"brand"`
}

type GetBrandCarsIDRes struct {
	CarsID []int `json:"cars_id"`
}

type GetModelCarsIDReq struct {
	Brand string `json:"brand"`
	Model string `json:"model"`
}

type GetModelCarsIDRes struct {
	CarsID []int `json:"cars_id"`
}

type User struct {
	UserID   int    `json:"id"`
	UserName string `json:"name"`
}

type Car struct {
	Id    int    `json:"id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Year  int    `json:"year"`
}

type HTTPResponse struct {
	Data any    `json:"data"`
	Err  string `json:"err"`
}

type GetIDsUserCarsReq struct {
	UserID int `json:"user_id"`
}

type GetIDsUserCarsRes struct {
	CarsID []int `json:"cars_id"`
}
