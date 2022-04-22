package constants

const (
	DEFAULT_SELECT_LIMIT    = "100"
	PAGE_SIZE_DEFAULT_VALUE = 10

	URL_BILLING_PREFIX = "billing"
	URL_STORAGE_PREFIX = "storage"

	HTTP_STATUS_SUCCESS = "success"
	HTTP_STATUS_ERROR   = "error"

	HTTP_CODE_200_OK                    = 200 //http.StatusOk
	HTTP_CODE_400_BAD_REQUEST           = 400 //http.StatusBadRequest
	HTTP_CODE_401_UNAUTHORIZED          = 401 //http.StatusUnauthorized
	HTTP_CODE_500_INTERNAL_SERVER_ERROR = 500 //http.StatusInternalServerError

	URL_HOST_GET_COMMON      = "/common"
	URL_HOST_GET_HOST_INFO   = "/host/info"
	URL_SYSTEM_CONFIG_PARAMS = "/system/params"

	SOURCE_ID_OF_PAYMENT = 4

	PROCESS_STATUS_WAITING_PAYMENT      = "Pending" //wait for payment
	PROCESS_STATUS_PROCESSING           = "Processing"
	CAR_FILE_STATUS_TASK_CREATED        = "TaskCreated"
	CAR_FILE_STATUS_DEAL_SENT           = "DealSent"
	CAR_FILE_STATUS_DEAL_SENT_FAILED    = "DealSentFailed"
	CAR_FILE_STATUS_DEAL_SEND_CANCELLED = "DealSendCancelled"
	PROCESS_STATUS_UNLOCK_REFUNDED      = "UnlockRefundSucceeded"
	PROCESS_STATUS_UNLOCK_REFUNDFAILED  = "UnlockRefundFailed"
	PROCESS_STATUS_EXPIRE_REFUNDING     = "Refunding"
	PROCESS_STATUS_EXPIRE_REFUNDED      = "Refunded"

	DEAL_STATUS_ACTIVE = "StorageDealActive"
	DEAL_STATUS_ERROR  = "StorageDealError"

	IPFS_URL_PREFIX_BEFORE_HASH = "/ipfs/"
	IPFS_File_PINNED_STATUS     = "Pinned"

	LOTUS_TASK_TYPE_VERIFIED = "verified"
	LOTUS_TASK_TYPE_REGULAR  = "regular"

	FILE_BLOCK_NUMBER_MAX = 999999999999999

	SOURCE_FILE_UPLOAD_STATUS_CREATED      = "Created"
	SOURCE_FILE_UPLOAD_STATUS_PAID         = "Paid"
	SOURCE_FILE_UPLOAD_STATUS_TASK_CREATED = "TaskCreated"

	OFFLINE_DEAL_STATUS_CREATED              = "Created"
	OFFLINE_DEAL_UNLOCK_STATUS_NOT_UNLOCKED  = "NotUnlocked"
	OFFLINE_DEAL_UNLOCK_STATUS_UNLOCKED      = "Unlocked"
	OFFLINE_DEAL_UNLOCK_STATUS_UNLOCK_FAILED = "UnlockFailed"

	SIGNATURE_SUCCESS_VALUE = "1" //init value,no unlock operation has been performed
	SIGNATURE_FAILED_VALUE  = "2" //init value,no unlock operation has been performed

	DURATION_DAYS_DEFAULT = 525

	SOURCE_FILE_TYPE_NORMAL = 0
	SOURCE_FILE_TYPE_MINT   = 1

	BYTES_1GB     = 1024 * 1024 * 1024
	EPOCH_PER_DAY = 24 * 60 * 2

	PRIVATE_KEY_ON_POLYGON = "privateKeyOnPolygon"

	//0:pay,1:unlock, 2: refund after unlock, 3:refund after expired
	TRANSACTION_TYPE_PAY                  = 0
	TRANSACTION_TYPE_UNLOCK               = 1
	TRANSACTION_TYPE_REFUND_AFTER_UNLOCK  = 2
	TRANSACTION_TYPE_REFUND_AFTER_EXPIRED = 3

	COIN_USDC_NAME   = "USDC"
	COIN_USDC_ADRESS = "USDC_ADDRESS"

	WALLET_TYPE_META_MASK = 0
	WALLET_TYPE_FILE_COIN = 1

	NETWORK_NAME_POLYGON = "polygon"
)
