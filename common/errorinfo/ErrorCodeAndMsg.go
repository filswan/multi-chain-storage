package errorinfo

const (
	//event error 001
	GET_EVENT_FROM_DB_ERROR_CODE = "500001001"
	GET_EVENT_FROM_DB_ERROR_MSG  = "Get event data from db error"

	//http request
	HTTP_REQUEST_PARAMS_JSON_FORMAT_ERROR_CODE        = "500002001"
	HTTP_REQUEST_PARAMS_JSON_FORMAT_ERROR_MSG         = "Request JSON format was error"
	HTTP_REQUEST_PARAMS_NULL_ERROR_CODE               = "500002002"
	HTTP_REQUEST_PARAMS_NULL_ERROR_MSG                = "Request params value can not be null"
	PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_CODE             = "500002003"
	PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_MSG              = "Page number or page size format was wrong"
	HTTP_REQUEST_SEND_REQUEST_RETUREN_ERROR_CODE      = "500002004"
	HTTP_REQUEST_SEND_REQUEST_RETUREN_ERROR_MSG       = "Return error when sending http request"
	HTTP_REQUEST_PARSER_RESPONSE_TO_STRUCT_ERROR_CODE = "500002005"
	HTTP_REQUEST_PARSER_RESPONSE_TO_STRUCT_ERROR_MSG  = "Parse http request to structure occurred error"
	HTTP_REQUEST_PARSER_STRUCT_TO_REQUEST_ERROR_CODE  = "500002006"
	HTTP_REQUEST_PARSER_STRUCT_TO_REQUEST_ERROR_MSG   = "Parse structure to http request request occurred error"
	HTTP_REQUEST_GET_RESPONSE_ERROR_CODE              = "500002007"
	HTTP_REQUEST_GET_RESPONSE_ERROR_MSG               = "Get http response occurred error"
	HTTP_REQUEST_PARAM_TYPE_ERROR_CODE                = "500002008"
	HTTP_REQUEST_PARAM_TYPE_ERROR_MSG                 = "Http request param type is error"

	//database err 003
	GET_RECORD_COUNT_ERROR_CODE  = "500003001"
	GET_RECORD_COUNT_ERROR_MSG   = "Get data total count from db occurred error"
	GET_RECORD_lIST_ERROR_CODE   = "500003002"
	GET_RECORD_lIST_ERROR_MSG    = "Get data records from db occurred error"
	SAVE_DATA_TO_DB_ERROR_CODE   = "500003003"
	SAVE_DATA_TO_DB_ERROR_MSG    = "Saving data to database occurred error"
	UPDATE_DATA_TO_DB_ERROR_CODE = "500003004"
	UPDATE_DATA_TO_DB_ERROR_MSG  = "Updating data to database occurred error"

	// JWT token err 004
	NO_AUTHORIZATION_TOKEN_ERROR_CODE = "500004001"
	NO_AUTHORIZATION_TOKEN_ERROR_MSG  = "Jwt token is required"

	//lotus send deal err  005
	SENDING_DEAL_ERROR_CODE                 = "500005001"
	SENDING_DEAL_ERROR_MSG                  = "Failed to create a sending deal task, please contact the administrator"
	GET_LATEST_PRICE_OF_FILECOIN_ERROR_CODE = "500005002"
	GET_LATEST_PRICE_OF_FILECOIN_ERROR_MSG  = "Getting filecion latest price occurred error"
	SENDING_DEAL_GET_NULL_RETURN_VALUE_CODE = "500005003"
	SENDING_DEAL_GET_NULL_RETURN_VALUE_MSG  = "Get null return value when deal had been sent"

	//system command err  006
	GET_HOME_DIR_ERROR_CODE = "500006001"
	GET_HOME_DIR_ERROR_MSG  = "Getting home dir occurred error"

	CREATE_DIR_ERROR_CODE = "500006002"
	CREATE_DIR_ERROR_MSG  = "Creating dir occurred error"

	//type transfer error 007
	TYPE_TRANSFER_ERROR_CODE = "500007001"
	TYPE_TRANSFER_ERROR_MSG  = "type transfer occurred error"
)
