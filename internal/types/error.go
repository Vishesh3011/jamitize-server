package types

type Layer string

const (
	Application Layer = "application"
	Service     Layer = "service"
	Repository  Layer = "repository"
	Controller  Layer = "controller"
)

type Status int

const (
	NotFound            Status = iota + 404
	BadRequest          Status = 400
	InternalServerError Status = 500
	Unauthorized        Status = 401
	Forbidden           Status = 403
	Conflict            Status = 409
	Unprocessable       Status = 422
	TooManyRequests     Status = 429
	NotImplemented      Status = 501
	BadGateway          Status = 502
)
