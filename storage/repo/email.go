package repo

import (
	pbe "email-service/genproto/email"
)

type EmailStorageI interface {
	RegisterUser(*pbe.RegUserReq) (*pbe.UserResp, error)
}
