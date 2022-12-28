package models

// response code constants
const (
	INTERNAL_SERVER_ERROR = 500
	SUCCESS               = 1000
	INVALID_TOKEN         = 1001
	INVALID_INPUT         = 1002
)

//  response message constants
const (
	Hashing_ERROR_MESSAGE           = "Hasing error"
	INVALID_TOKEN_MESSAGE           = "Invalid JWT Token error"
	REDIS_SET_ERROR_MESSAGE         = "Redis Key Set Error"
	REDIS_GET_ERROR_MESSAGE         = "Redis Key Get Error"
	REDIS_DEL_ERROR_MESSAGE         = "Redis Key Delete Error"
	DB_OPERATION_ERROR_MESSAGE      = "Invalid Database Operation."
	INVALID_INPUT_MESSAGE           = "Invalid / Empty Input"
	SAMPLE_REQUEST_COMPLETE_MESSAGE = "Successfully Processed Request"
)
