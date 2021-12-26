module picasso

go 1.16

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gabriel-vasile/mimetype v1.4.0 // indirect
	github.com/gin-gonic/gin v1.7.4
	github.com/go-cmd/cmd v1.2.1
	github.com/limoxi/ghost v0.0.1-alpha.0.20200121095608-f52457b4e1c9
)

replace github.com/limoxi/ghost => ../ghost

replace github.com/jinzhu/gorm => github.com/limoxi/gorm v1.9.120
