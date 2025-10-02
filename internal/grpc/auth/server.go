package auth

import ssov1 "github.com/Deymos01/grpc-auth-service/tree/main/protos/gen/go/sso"

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}
