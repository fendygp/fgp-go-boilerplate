package utils

var (
	ErrInput       = APIError{ResponseCode: "0001", ResponseDescription: "cannot read input data"}
	ErrCreateUsers = APIError{ResponseCode: "0002", ResponseDescription: "error create users"}
)
