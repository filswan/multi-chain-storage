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
	HTTP_REQUEST_HEADER_AUTHRORIZATION  = "Authorization"

	NETWORK_TYPE_GOERLI       = "goerli"
	NETWORK_TYPE_GOERLI_UUID  = "927ccb7b-072b-47c1-af43-0c07e362ae23"
	NETWORK_TYPE_POLYGON      = "polygon"
	NETWORK_TYPE_POLYGON_UUID = "42746d02-b407-4bd9-bf2a-38381e009517"
	NETWORK_TYPE_NBAI         = "nbai"
	NETWORK_TYPE_NBAI_UUID    = "05502a3a-22a8-49e4-86bc-539a297f76be"
	NETWORK_TYPE_BSC          = "bsc"
	NETWORK_TYPE_BSC_UUID     = "f74f7f00-6ea3-41b6-85f8-912e3a14f132"
	NETWORK_TYPE_ETH          = "ethereum"
	NETWORK_TYPE_ETH_UUID     = "3835b57a-0941-4d95-9e92-2726bdc013c3"

	COIN_TYPE_USDC_ON_POLYGON_UUID = "0732db61-10c7-4f16-b349-d60afc1d7a34"
	COIN_TYPE_WFIL_ON_POLYGON_UUID = "c4623fc0-2b38-4af1-9bb5-297474187823"

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

	SIGNATURE_DEFAULT_VALUE = "0" //init value,no unlock operation has been performed
	SIGNATURE_SUCCESS_VALUE = "1" //init value,no unlock operation has been performed
	SIGNATURE_FAILED_VALUE  = "2" //init value,no unlock operation has been performed

	SOURCE_ID_OF_PAYMENT = 4
	TASK_STATUS_ASSIGNED = "Assigned"

	LOCK_PAYMENT_STATUS_FREE       = "Free"
	LOCK_PAYMENT_STATUS_WAITING    = "Pending"    //wait for lock payment
	LOCK_PAYMENT_STATUS_PAID       = "Paid"       //lock payment paid
	LOCK_PAYMENT_STATUS_PROCESSING = "Processing" //lock payment fail
	LOCK_PAYMENT_STATUS_SUCCESS    = "Active"     //lock payment active
	LOCK_PAYMENT_STATUS_REFUNDED   = "Refunded"

	SEND_DEAL_STATUS_SUCCESS = "Success"
	SEND_DEAL_STATUS_FAIL    = "Fail"

	DEAL_STATUS_ACTIVE = "StorageDealActive"
	DEAL_STATUS_ERROR  = "StorageDealError"

	IPFS_URL_PREFIX_BEFORE_HASH = "/ipfs/"
	IPFS_File_PINNED_STATUS     = "Pinned"

	LOTUS_TASK_TYPE_VERIFIED = "verified"
	LOTUS_TASK_TYPE_REGULAR  = "regular"

	LOTUS_PRICE_MULTIPLE_1E18 = 1e18 // 10^18
	FILE_BLOCK_NUMBER_MAX     = 999999999999999
	TIME_HALF_AN_HOUR         = 30 * 60 * 1000
)
