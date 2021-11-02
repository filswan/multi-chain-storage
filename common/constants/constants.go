package constants

const (
	SWAN_PAYMENT_ABI_JSON = "on-chain/contracts/abi/SwanPayment.json"
	DEFAULT_SELECT_LIMIT  = "100"

	URL_EVENT_PREFIX   = "events"
	URL_BILLING_PREFIX = "billing"
	URL_STORAGE_PREFIX = "storage"

	HTTP_STATUS_SUCCESS = "success"
	HTTP_STATUS_FAIL    = "fail"
	HTTP_STATUS_ERROR   = "error"

	HTTP_CODE_200_OK                    = "200" //http.StatusOk
	HTTP_CODE_400_BAD_REQUEST           = "400" //http.StatusBadRequest
	HTTP_CODE_401_UNAUTHORIZED          = "401" //http.StatusUnauthorized
	HTTP_CODE_500_INTERNAL_SERVER_ERROR = "500" //http.StatusInternalServerError

	NETWORK_TYPE_GOERLI  = "goerli"
	NETWORK_TYPE_POLYGON = "polygon"
	NETWORK_TYPE_NBAI    = "nbai"
	NETWORK_TYPE_BSC     = "bsc"
	NETWORK_TYPE_ETH     = "ethereum"

	TRANSACTION_STATUS_SUCCESS = "success"
	TRANSACTION_STATUS_FAIL    = "fail"

	URL_HOST_GET_COMMON      = "/common"
	URL_HOST_GET_HOST_INFO   = "/host/info"
	URL_SYSTEM_CONFIG_PARAMS = "/system/params"

	TABLE_NAME_EVENT_BSC     = "event_bsc"
	TABLE_NAME_EVENT_GOERLI  = "event_goerli"
	TABLE_NAME_EVENT_POLYGON = "event_polygon"

	PAGE_SIZE_DEFAULT_VALUE = "10"

	EVENT_NAME_LOCAK_PAYMENT = "LockPayment"
)
