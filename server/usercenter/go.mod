module github.com/znk_fullstack/server/usercenter

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v7 v7.2.0
	github.com/golang/protobuf v1.3.5
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/jinzhu/gorm v1.9.12
	github.com/oklog/ulid/v2 v2.0.2
	github.com/rs/zerolog v1.18.0
	github.com/segmentio/ksuid v1.0.2
	golang.org/x/net v0.0.0-20191109021931-daa7c04131f5 // indirect
	golang.org/x/sys v0.0.0-20191110163157-d32e6e3b99c4 // indirect
	google.golang.org/genproto v0.0.0-20191108220845-16a3f7862a1a // indirect
	google.golang.org/grpc v1.25.1
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.28.0
