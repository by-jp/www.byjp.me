module github.com/by-jp/www.byjp.me/tools/import/webmentionio

go 1.21.6

require (
	github.com/by-jp/www.byjp.me/tools/syndicate v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.5.1
	github.com/kelseyhightower/envconfig v1.4.0
)

replace github.com/by-jp/www.byjp.me/tools/syndicate => ../../syndicate
