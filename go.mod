module myweb

go 1.13

require superly.club/web/routers v0.0.0-00010101000000-000000000000 // indirect

replace (
	superly.club/web/controllers => ./controllers
	superly.club/web/models => ./models
	superly.club/web/routers => ./routers
)
