module github.com/znk_fullstack/server/usercenter

go 1.13

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v7 v7.2.0
	github.com/golang/protobuf v1.3.5
	github.com/jinzhu/gorm v1.9.12
	github.com/micro/go-micro v1.18.0
	github.com/oklog/ulid/v2 v2.0.2
	github.com/rs/zerolog v1.18.0
	github.com/segmentio/ksuid v1.0.2
	google.golang.org/grpc v1.26.0
	k8s.io/apiserver v0.18.2 // indirect
	k8s.io/client-go v12.0.0+incompatible // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.28.0
