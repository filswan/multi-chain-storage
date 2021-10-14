package errorinfo

const (
	//event error 001
	GET_EVENT_FROM_DB_ERROR_CODE = "500001001"
	GET_EVENT_FROM_DB_ERROR_MSG  = "Get event data from db error"

	//page info err 002
	PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_CODE = "500002001"
	PAGE_NUMBER_OR_SIZE_FORMAT_ERROR_MSG  = "Page number or page size format was wrong"

	//database err 003
	GET_RECORD_COUNT_ERROR_CODE = "500003001"
	GET_RECORD_COUNT_ERROR_MSG  = "Get data total count from db occurred error"
	GET_RECORD_lIST_ERROR_CODE  = "500003002"
	GET_RECORD_lIST_ERROR_MSG   = "Get data records from db occurred error"
)
