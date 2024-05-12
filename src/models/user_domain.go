package model

import (
	"crypto/md5"
	"encoding/hex"

	resterr "github.com/setxpro/crud-go/src/configurations/rest_err"
)

// Constructor
func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &UserDomain{
		email,
		password,
		name,
		age,
	}
}

type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *resterr.RestErr
	UpdateUser(string) *resterr.RestErr
	FindUser(string) (*UserDomain, *resterr.RestErr)
	DeleteUser(string) *resterr.RestErr
}
