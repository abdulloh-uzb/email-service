package postgres

import (
	pbe "email-service/genproto/email"

	"github.com/jmoiron/sqlx"
)

type emailRepo struct {
	db *sqlx.DB
}

func NewEmailRepo(db *sqlx.DB) *emailRepo {
	return &emailRepo{db: db}
}

func (e *emailRepo) RegisterUser(req *pbe.RegUserReq) (*pbe.UserResp, error) {
	user := &pbe.UserResp{}
	err := e.db.QueryRow(`insert into 
	users(username, password, email) 
	values($1,$2,$3)
	returning id, username, password, email, verify`,
		req.Username, req.Password, req.Email).Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.Verify)
	if err != nil {
		return &pbe.UserResp{}, err
	}
	return user, nil
}
