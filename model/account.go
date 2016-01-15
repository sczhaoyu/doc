package model

import (
	"errors"
	"fmt"
	"time"
)

type Account struct {
	Id        int64     `json:"id"`
	Account   string    `json:"account"`
	Password  string    `json:"password"`
	NickName  string    `json:"nickName"`
	CreatedAt time.Time `json:"createdAt"`
}

func GetAccount(account, password string) (*Account, error) {
	var ret Account
	b, err := DocDB.Where("account=? and password=?", account, password).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New("account or password err")
	}
	return &ret, nil
}
func (a *Account) Save() error {
	_, err := DocDB.Insert(a)
	return err
}

func (a *Account) Delete() error {
	_, err := DocDB.Where("id=?", a.Id).Delete(Account{})
	return err
}
func GetAccountById(id int64) (*Account, error) {
	var ret Account
	b, err := DocDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Account Not Found Value: %v", id))
	}
	return &ret, nil
}

func (a *Account) Update() error {
	_, err := DocDB.Where("id=?", a.Id).Update(a)
	return err
}
