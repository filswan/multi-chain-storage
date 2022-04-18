package constants

const (
	DEFAULT_SELECT_LIMIT    = "100"
	PAGE_SIZE_DEFAULT_VALUE = "10"

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

	URL_HOST_GET_COMMON      = "/common"
	URL_HOST_GET_HOST_INFO   = "/host/info"
	URL_SYSTEM_CONFIG_PARAMS = "/system/params"

	SOURCE_ID_OF_PAYMENT = 4

	PROCESS_STATUS_WAITING_PAYMENT     = "Pending" //wait for payment
	PROCESS_STATUS_PAID                = "Paid"
	PROCESS_STATUS_PROCESSING          = "Processing"
	PROCESS_STATUS_TASK_CREATED        = "TaskCreated"
	PROCESS_STATUS_DEAL_SENT           = "DealSent"
	PROCESS_STATUS_DEAL_SENT_FAILED    = "DealSentFailed"
	PROCESS_STATUS_DEAL_SEND_CANCELLED = "DealSendCancelled"
	PROCESS_STATUS_UNLOCK_REFUNDED     = "UnlockRefundSucceeded"
	PROCESS_STATUS_UNLOCK_REFUNDFAILED = "UnlockRefundFailed"
	PROCESS_STATUS_EXPIRE_REFUNDING    = "Refunding"
	PROCESS_STATUS_EXPIRE_REFUNDED     = "Refunded"

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

	OFFLINE_DEAL_UNLOCK_STATUS_NOT_UNLOCKED  = "NotUnlocked"
	OFFLINE_DEAL_UNLOCK_STATUS_UNLOCKED      = "Unlocked"
	OFFLINE_DEAL_UNLOCK_STATUS_UNLOCK_FAILED = "UnlockFailed"

	SIGNATURE_SUCCESS_VALUE = "1" //init value,no unlock operation has been performed
	SIGNATURE_FAILED_VALUE  = "2" //init value,no unlock operation has been performed

	DURATION_DAYS_DEFAULT = 525

	SOURCE_FILE_TYPE_NORMAL = 0

	BYTES_1GB     = 1024 * 1024 * 1024
	EPOCH_PER_DAY = 24 * 60 * 2

	PRIVATE_KEY_ON_POLYGON = "privateKeyOnPolygon"

	//0:payment contract,1:pay admin fee, 2:dao, 2:user,1000:user,2000:file coin
	WALLET_TYPE_PAY_CONTRACT = 0
	WALLET_TYPE_ADMIN_FEE    = 1
	WALLET_TYPE_DAO          = 2
	WALLET_TYPE_USER         = 1000
	WALLET_TYPE_FILE_COIN    = 2000

	//0:pay,1:unlock, 2: refund after unlock, 3:refund after expired
	TRANSACTION_TYPE_PAY                  = 0
	TRANSACTION_TYPE_UNLOCK               = 1
	TRANSACTION_TYPE_REFUND_AFTER_UNLOCK  = 2
	TRANSACTION_TYPE_REFUND_AFTER_EXPIRED = 3

	CONFIG_SWAN_PAYMENT_CONTRACT_ADDRESS = "SWAN_PAYMENT_CONTRACT_ADDRESS"

	COIN_USDC_NAME   = "USDC"
	COIN_USDC_ADRESS = "USDC_ADDRESS"
)
