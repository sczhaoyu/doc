package model

import (
	"errors"
	"fmt"
)

type ErrCode struct {
	Id              int64  `json:"errCodeid"`
	Code            string `json:"code"`            //错误代码
	DescriptionText string `json:"descriptionText"` //描述
}

func (e *ErrCode) Save() error {
	_, err := DocDB.Insert(e)
	return err
}

func (e *ErrCode) Delete() error {
	_, err := DocDB.Where("id=?", e.Id).Delete(ErrCode{})
	return err
}
func GetErrCodeById(id int64) (*ErrCode, error) {
	var ret ErrCode
	b, err := DocDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("ErrCode Not Found Value: %v", id))
	}
	return &ret, nil
}

func (e *ErrCode) Update() error {
	_, err := DocDB.Where("id=?", e.Id).Update(e)
	return err
}
func FindErrCode() ([]ErrCode, error) {
	var ret []ErrCode
	err := DocDB.Asc("code").Find(&ret)
	if err != nil {
		return nil, err
	}
	return ret, NoDataMsg(len(ret) > 0, "ErrCode[] null")
}
