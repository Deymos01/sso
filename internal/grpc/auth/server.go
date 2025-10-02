package auth

import ssov1 "github.com/Deymos01/protos/gen/go/sso"

type serverAPI struct {
	ssov1.UnimplementedAuthServer
}
