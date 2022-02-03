package errorinfo

const (
	//event error 001
	GET_EVENT_FROM_DB_ERROR_CODE = "500001001"

	//http request
	HTTP_REQUEST_PARAMS_JSON_FORMAT_ERROR_CODE        = "500002001"
	HTTP_REQUEST_PARAMS_NULL_ERROR_CODE               = "500002002"
	PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_CODE             = "500002003"
	HTTP_REQUEST_SEND_REQUEST_RETUREN_ERROR_CODE      = "500002004"
	HTTP_REQUEST_PARSER_RESPONSE_TO_STRUCT_ERROR_CODE = "500002005"
	HTTP_REQUEST_PARSER_STRUCT_TO_REQUEST_ERROR_CODE  = "500002006"
	HTTP_REQUEST_GET_RESPONSE_ERROR_CODE              = "500002007"
	HTTP_REQUEST_PARAM_TYPE_ERROR_CODE                = "500002008"

	//database err 003
	GET_RECORD_COUNT_ERROR_CODE  = "500003001"
	GET_RECORD_lIST_ERROR_CODE   = "500003002"
	SAVE_DATA_TO_DB_ERROR_CODE   = "500003003"
	UPDATE_DATA_TO_DB_ERROR_CODE = "500003004"

	//lotus send deal err  005
	SAVE_FILE_ERROR                         = "500005001"
	GET_LATEST_PRICE_OF_FILECOIN_ERROR_CODE = "500005002"
	SENDING_DEAL_GET_NULL_RETURN_VALUE_CODE = "500005003"

	//system command err  006
	GET_HOME_DIR_ERROR_CODE = "500006001"
	CREATE_DIR_ERROR_CODE   = "500006002"

	//type transfer error 007
	TYPE_TRANSFER_ERROR_CODE = "500007001"
)

var errorMap map[string]string

func init() {
	errorMap = map[string]string{
		GET_EVENT_FROM_DB_ERROR_CODE:                      "Get event data from db error",
		HTTP_REQUEST_PARAMS_JSON_FORMAT_ERROR_CODE:        "Request JSON format was error",
		HTTP_REQUEST_PARAMS_NULL_ERROR_CODE:               "Request params value can not be null",
		PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_CODE:             "Page number or page size format was wrong",
		HTTP_REQUEST_SEND_REQUEST_RETUREN_ERROR_CODE:      "Return error when sending http request",
		HTTP_REQUEST_PARSER_RESPONSE_TO_STRUCT_ERROR_CODE: "Parse http request to structure occurred error",
		HTTP_REQUEST_PARSER_STRUCT_TO_REQUEST_ERROR_CODE:  "Parse structure to http request request occurred error",
		HTTP_REQUEST_GET_RESPONSE_ERROR_CODE:              "Get http response occurred error",
		HTTP_REQUEST_PARAM_TYPE_ERROR_CODE:                "Http request param type is error",
		GET_RECORD_COUNT_ERROR_CODE:                       "Get data total count from db occurred error",
		GET_RECORD_lIST_ERROR_CODE:                        "Get data records from db occurred error",
		SAVE_DATA_TO_DB_ERROR_CODE:                        "Saving data to database occurred error",
		UPDATE_DATA_TO_DB_ERROR_CODE:                      "Updating data to database occurred error",
		SAVE_FILE_ERROR:                                   "Failed to save file on server, please contact the administrator",
		GET_LATEST_PRICE_OF_FILECOIN_ERROR_CODE:           "Getting filecion latest price occurred error",
		SENDING_DEAL_GET_NULL_RETURN_VALUE_CODE:           "Get null return value when deal had been sent",
		GET_HOME_DIR_ERROR_CODE:                           "Getting home dir occurred error",
		CREATE_DIR_ERROR_CODE:                             "Creating dir occurred error",
		TYPE_TRANSFER_ERROR_CODE:                          "type transfer occurred error",
	}
}

func GetErrMsg(errCode string) string {
	errMsg := errorMap[errCode]

	if errMsg == "" {
		errMsg = "internal error"
	}

	return errMsg
}
