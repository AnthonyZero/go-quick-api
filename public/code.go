package public

var (
	OK                  = &CodeNo{Code: 0, Message: "OK"}
	InternalServerError = &CodeNo{Code: 10001, Message: "Internal server error"}
	ErrBind             = &CodeNo{Code: 10002, Message: "Error occurred while binding the request body to the struct"}

	// user errors
	ErrUserNotFound = &CodeNo{Code: 20102, Message: "The user was not found"}

	// validation failed
	ErrValidation = &CodeNo{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &CodeNo{Code: 20002, Message: "Database error."}
	ErrToken      = &CodeNo{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &CodeNo{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrTokenInvalid      = &CodeNo{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &CodeNo{Code: 20104, Message: "The password was incorrect."}
)
