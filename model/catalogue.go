package model

import (
	"errors"
	"fmt"
)

type Catalogue struct {
	Id           int64  `json:"catalogueId"`
	ParentId     int64  `json:"parentId"`     //父目录，没有为0
	Name         string `json:"name"`         //目录名称
	SerialNumber string `json:"serialNumber"` //序号
}

//获取全部父目录
func FindCatalogueAllParent() ([]Catalogue, error) {
	var ret []Catalogue
	err := DocDB.Where("parent_id=?", 0).Find(&ret)
	if err != nil {
		return nil, err
	}
	return ret, NoDataMsg(len(ret) > 0, "catalogue[] null")
}

//获取子目录
func FindFindCatalogueByParent(catalogueId int64) ([]Catalogue, error) {
	var ret []Catalogue
	err := DocDB.Where("parent_id=?", catalogueId).Find(&ret)
	if err != nil {
		return nil, err
	}
	return ret, NoDataMsg(len(ret) > 0, "catalogue[] null")
}
func (c *Catalogue) Save() error {
	_, err := DocDB.Insert(c)
	return err
}

func (c *Catalogue) Delete() error {
	_, err := DocDB.Where("id=?", c.Id).Delete(Catalogue{})
	return err
}
func GetCatalogueById(id int64) (*Catalogue, error) {
	var ret Catalogue
	b, err := DocDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Catalogue Not Found Value: %v", id))
	}
	return &ret, nil
}

func (c *Catalogue) Update() error {
	_, err := DocDB.Where("id=?", c.Id).Update(c)
	return err
}
