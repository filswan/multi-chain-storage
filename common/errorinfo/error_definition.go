package errorinfo

const (
	ERROR_PARAM_NULL            = 10001
	ERROR_PARAM_WRONG_TYPE      = 10002
	ERROR_PARAM_INVALID_VALUE   = 10003
	ERROR_PARAM_PARSE_TO_STRUCT = 10004
	ERROR_INTERNAL              = 20001
)

var errorMap map[int]string

func init() {
	errorMap = map[int]string{
		ERROR_PARAM_NULL:            "param cannot be null",
		ERROR_PARAM_WRONG_TYPE:      "wrong param data type",
		ERROR_PARAM_INVALID_VALUE:   "invalid param value",
		ERROR_PARAM_PARSE_TO_STRUCT: "params parse to structure fail",
		ERROR_INTERNAL:              "Internal error",
	}
}

func GetErrMsg(errCode int) string {
	errMsg := errorMap[errCode]

	if errMsg == "" {
		errMsg = errorMap[ERROR_INTERNAL]
	}

	return errMsg
}
