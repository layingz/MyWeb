module myweb

go 1.13

require (
	github.com/gin-gonic/gin v1.5.0 // indirect
	github.com/go-sql-driver/mysql v1.4.1 // indirect
	superly.club/web/controllers v0.0.0-00010101000000-000000000000 // indirect
	superly.club/web/models v0.0.0-00010101000000-000000000000 // indirect
	superly.club/web/routers v0.0.0-00010101000000-000000000000 // indirect
)

replace (
	superly.club/web/controllers => ./controllers
	superly.club/web/models => ./models
	superly.club/web/routers => ./routers
)
