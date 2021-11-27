module payment-bridge

go 1.16

require (
	github.com/BurntSushi/toml v0.4.1
	github.com/ethereum/go-ethereum v1.10.7
	github.com/go-sql-driver/mysql v1.6.0
	github.com/jinzhu/gorm v1.9.16
)

require (
	github.com/StackExchange/wmi v0.0.0-20190523213315-cbe66965904d // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/filswan/go-swan-client v0.0.39
	github.com/filswan/go-swan-lib v0.2.55
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.3.0
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-ole/go-ole v1.2.4 // indirect
	github.com/itsjamie/gin-cors v0.0.0-20160420130702-97b4a9da7933
	github.com/joho/godotenv v1.4.0
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/lib/pq v1.10.2 // indirect
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mattn/go-isatty v0.0.13 // indirect
	github.com/mattn/go-runewidth v0.0.10 // indirect
	github.com/onsi/ginkgo v1.16.4 // indirect
	github.com/onsi/gomega v1.13.0 // indirect
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/robfig/cron v1.2.0
	github.com/shopspring/decimal v1.3.1
	github.com/sirupsen/logrus v1.8.1
	github.com/ugorji/go v1.2.6 // indirect
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97 // indirect
	golang.org/x/tools v0.1.7 // indirect
	gopkg.in/go-playground/assert.v1 v1.2.1 // indirect
	gopkg.in/go-playground/validator.v8 v8.18.2 // indirect
)

replace github.com/filecoin-project/filecoin-ffi => ./extern/filecoin-ffi
