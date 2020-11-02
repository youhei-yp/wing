module github.com/youhei-yp/wing

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190506204251-e1dfcc566284
	golang.org/x/net => github.com/golang/net v0.0.0-20190503192946-f4e77d36d62c
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190508220229-2d0786266e9c
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190509153222-73554e0f7805
	gopkg.in/gomail.v2 => github.com/go-gomail/gomail v0.0.0-20160411212932-81ebce5c23df
	gopkg.in/yaml.v2 => github.com/go-yaml/yaml v2.1.0+incompatible
)

require (
	github.com/766b/go-outliner v0.0.0-20180511142203-fc6edecdadd7 // indirect
	github.com/BurntSushi/toml v0.3.1
	github.com/asaskevich/govalidator v0.0.0-20190424111038-f61b66f89f4a
	github.com/astaxie/beego v1.11.1
	github.com/bwmarrin/snowflake v0.3.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-sql-driver/mysql v1.4.1
	github.com/googollee/go-socket.io v1.0.1
	github.com/mozillazg/go-pinyin v0.18.0
	github.com/satori/go.uuid v1.2.0
	golang.org/x/crypto v0.0.0-20190308221718-c2843e01d9a2
	gopkg.in/alexcesaro/quotedprintable.v3 v3.0.0-20150716171945-2caba252f4dc // indirect
	gopkg.in/gomail.v2 v2.0.0-20160411212932-81ebce5c23df
	gopkg.in/yaml.v2 v2.2.2 // indirect
)
