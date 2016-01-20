package model

import (
	"errors"
	"fmt"
	"time"
)

type Explain struct {
	Id              int64     `json:"explainId"`
	Title           string    `json:"title"`           //标题
	DescriptionText string    `json:"descriptionText"` //描述
	UserName        string    `json:"userName"`        //发布人
	CreatedAt       time.Time `json:"createdAt"`       //发布时间
	ProjectId       int64     `json:"projectId"`       //项目的ID
	VersionId       int64     `json:"versionId"`       //版本的ID
}

func FindExplain() ([]Explain, error) {
	var ret []Explain
	err := DocDB.Find(&ret)
	if err != nil {
		return nil, err
	}
	return ret, NoDataMsg(len(ret) > 0, "Explain[] null")
}
func (e *Explain) Save() error {
	_, err := DocDB.Insert(e)
	return err
}

func (e *Explain) Delete() error {
	_, err := DocDB.Where("id=?", e.Id).Delete(Explain{})
	return err
}
func GetExplainById(id int64) (*Explain, error) {
	var ret Explain
	b, err := DocDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Explain Not Found Value: %v", id))
	}
	return &ret, nil
}

func (e *Explain) Update() error {
	_, err := DocDB.Where("id=?", e.Id).Update(e)
	return err
}
