package constants

const (
	SWAN_PAYMENT_ABI_JSON = "on-chain/contracts/abi/SwanPayment.json"
	DEFAULT_SELECT_LIMIT  = "100"

	URL_EVENT_PREFIX = "events"

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

	TRANSACTION_STATUS_SUCCESS = "success"
	TRANSACTION_STATUS_FAIL    = "fail"

	URL_HOST_GET_HOST_INFO = "/host/info"
	URL_HOST_GET_COMMON    = "/common"
)
