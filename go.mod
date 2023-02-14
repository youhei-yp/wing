module github.com/youhei-yp/wing

go 1.14

replace (
	golang.org/x/net => github.com/golang/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190508220229-2d0786266e9c
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190509153222-73554e0f7805
)

require (
	github.com/adamzy/cedar-go v0.0.0-20170805034717-80a9c64b256d // indirect
	github.com/astaxie/beego v1.12.2
	github.com/bwmarrin/snowflake v0.3.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/garyburd/redigo v1.6.2
	github.com/go-playground/validator/v10 v10.9.0
	github.com/huichen/sego v0.0.0-20180617034105-3f3c8a8cfacc
	github.com/issue9/assert v1.4.1 // indirect
	github.com/mozillazg/go-pinyin v0.18.0
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
