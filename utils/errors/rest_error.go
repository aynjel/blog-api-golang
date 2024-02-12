package errors

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
	Level   string `json:"level"`
}

func NewInternalServerError(message string, level string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  500,
		Error:   "internal_server_error",
		Level:   level,
	}
}

func NewBadRequestError(message string, level string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  400,
		Error:   "bad_request",
		Level:   level,
	}
}

func NewNotFoundError(message string, level string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  404,
		Error:   "not_found",
		Level:   level,
	}
}

func NewUnauthorizedError(message string, level string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  401,
		Error:   "unauthorized",
		Level:   level,
	}
}

func NewForbiddenError(message string, level string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  403,
		Error:   "forbidden",
		Level:   level,
	}
}

func NewConflictError(message string, level string) *RestErr {
	return &RestErr{
		Message: message,
		Status:  409,
		Error:   "conflict",
		Level:   level,
	}
}
