package errorinfo

const (
	//http request
	HTTP_REQUEST_PARAM_ERROR_CODE_NULL                = "500002002"
	HTTP_REQUEST_PARAM_ERROR_CODE_WRONG_TYPE          = "500002008"
	HTTP_REQUEST_PARAM_ERROR_CODE_INVALID_VALUE       = "500002009"
	HTTP_REQUEST_PARSER_RESPONSE_TO_STRUCT_ERROR_CODE = "500002005"

	//database err 003
	GET_RECORD_lIST_ERROR_CODE = "500003002"
	SAVE_DATA_TO_DB_ERROR_CODE = "500003003"

	//lotus send deal err  005
	SAVE_FILE_ERROR                         = "500005001"
	GET_LATEST_PRICE_OF_FILECOIN_ERROR_CODE = "500005002"
)

var errorMap map[string]string

func init() {
	errorMap = map[string]string{
		HTTP_REQUEST_PARAM_ERROR_CODE_NULL:                "Request param cannot be null",
		HTTP_REQUEST_PARAM_ERROR_CODE_WRONG_TYPE:          "Request param data type is wrong",
		HTTP_REQUEST_PARAM_ERROR_CODE_INVALID_VALUE:       "Request param value is invalid",
		HTTP_REQUEST_PARSER_RESPONSE_TO_STRUCT_ERROR_CODE: "Parse http request to structure occurred error",
		GET_RECORD_lIST_ERROR_CODE:                        "Get data records from db occurred error",
		SAVE_DATA_TO_DB_ERROR_CODE:                        "Saving data to database occurred error",
		SAVE_FILE_ERROR:                                   "Failed to save file on server, please contact the administrator",
		GET_LATEST_PRICE_OF_FILECOIN_ERROR_CODE:           "Getting filecion latest price occurred error",
	}
}

func GetErrMsg(errCode string) string {
	errMsg := errorMap[errCode]

	if errMsg == "" {
		errMsg = "internal error"
	}

	return errMsg
}
