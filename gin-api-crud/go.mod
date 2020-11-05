module gin-api-crud

go 1.15

require (
	crud v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/gin v1.6.3
)

replace crud => ../manipulate-with-database/crud-mysql/
