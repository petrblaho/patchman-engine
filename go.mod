module app

go 1.18

require (
	github.com/aws/aws-sdk-go v1.44.116
	github.com/ezamriy/gorpm v0.0.0-20160905202458-25f7273cbf51
	github.com/getkin/kin-openapi v0.106.0
	github.com/gin-contrib/gzip v0.0.6
	github.com/gin-contrib/timeout v0.0.3
	github.com/gin-gonic/gin v1.8.1
	github.com/gocarina/gocsv v0.0.0-20200330101823-46266ca37bd3 // newer version skips some columns
	github.com/golang-migrate/migrate/v4 v4.15.2
	github.com/hashicorp/golang-lru v0.5.4
	github.com/jackc/pgconn v1.13.0
	github.com/joho/godotenv v1.4.0
	github.com/lestrrat-go/backoff v1.0.1
	github.com/lib/pq v1.10.7
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0
	github.com/prometheus/client_golang v1.13.0
	github.com/redhatinsights/app-common-go v1.6.4
	github.com/redhatinsights/identity v0.0.0-20220719174832-36a7b1cbeff1
	github.com/redhatinsights/platform-go-middlewares v0.20.0
	github.com/segmentio/kafka-go v0.4.35
	github.com/sirupsen/logrus v1.9.0
	github.com/stretchr/testify v1.8.0
	github.com/swaggo/files v0.0.0-20220728132757-551d4a08d97a
	github.com/swaggo/gin-swagger v1.5.3
	github.com/zsais/go-gin-prometheus v0.1.0
	golang.org/x/net v0.7.0
	gorm.io/driver/postgres v1.4.4
	gorm.io/gorm v1.24.0
	modernc.org/strutil v1.1.3
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.7 // indirect
	github.com/go-openapi/swag v0.22.3 // indirect
	github.com/go-playground/locales v0.14.0 // indirect
	github.com/go-playground/universal-translator v0.18.0 // indirect
	github.com/go-playground/validator/v10 v10.11.1 // indirect
	github.com/goccy/go-json v0.9.11 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/invopop/yaml v0.2.0 // indirect
	github.com/jackc/chunkreader/v2 v2.0.1 // indirect
	github.com/jackc/pgio v1.0.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgproto3/v2 v2.3.1 // indirect
	github.com/jackc/pgservicefile v0.0.0-20200714003250-2b9c44734f2b // indirect
	github.com/jackc/pgtype v1.12.0 // indirect
	github.com/jackc/pgx/v4 v4.17.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/klauspost/compress v1.15.11 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.16 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.2 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/mohae/deepcopy v0.0.0-20170929034955-c48cc78d4826 // indirect
	github.com/pelletier/go-toml/v2 v2.0.5 // indirect
	github.com/pierrec/lz4/v4 v4.1.17 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.37.0 // indirect
	github.com/prometheus/procfs v0.8.0 // indirect
	github.com/swaggo/swag v1.8.6 // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/xdg/scram v1.0.5 // indirect
	github.com/xdg/stringprep v1.0.3 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	golang.org/x/crypto v0.0.0-20221012134737-56aed061732a // indirect
	golang.org/x/sys v0.5.0 // indirect
	golang.org/x/text v0.7.0 // indirect
	golang.org/x/tools v0.1.12 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/ezamriy/gorpm v0.0.0-20160905202458-25f7273cbf51 => github.com/MichaelMraka/gorpm v0.0.0-20210923131407-e21b5950f175
