package model

import (
	"errors"
	"fmt"
	"time"
)

type UpdateLog struct {
	Id              int64     `json:"id"`
	DescriptionText string    `json:"descriptionText"` //更新说明
	CreatedAt       time.Time `json:"createdAt"`       //添加时间
}

func (u *UpdateLog) Save() error {
	_, err := DocDB.Insert(u)
	return err
}
func AddUpdateLog(txt string) {
	var l UpdateLog
	l.CreatedAt = time.Now().Local()
	l.DescriptionText = txt
	DocDB.Insert(&l)
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

//只显示最新的50条
func FindUpdateLog() ([]UpdateLog, error) {
	var ret []UpdateLog
	err := DocDB.Desc("created_at").Limit(50, 0).Find(&ret)
	if err != nil {
		return nil, err
	}
	return ret, NoDataMsg(len(ret) > 0, "UpdateLog[] null")
}
