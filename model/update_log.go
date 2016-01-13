package model

import (
	"errors"
	"fmt"
)

type UpdateLog struct {
	Id              int64  `json:"id"`
	DescriptionText string `json:"descriptionText"` //更新说明
}

func (u *UpdateLog) Save() error {
	_, err := DocDB.Insert(u)
	return err
}

func (u *UpdateLog) Delete() error {
	_, err := DocDB.Where("id=?", u.Id).Delete(UpdateLog{})
	return err
}
func GetUpdateLogById(id int64) (*UpdateLog, error) {
	var ret UpdateLog
	b, err := DocDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("UpdateLog Not Found Value: %v", id))
	}
	return &ret, nil
}

func (u *UpdateLog) Update() error {
	_, err := DocDB.Where("id=?", u.Id).Update(u)
	return err
}
