package model

import (
	"errors"
	"fmt"
)

type Doc struct {
	Id              int64  `json:"docId"`           //主键
	CatalogueId     int64  `json:"catalogueId"`     //所属目录
	Path            string `json:"path"`            //接口请求路径
	DescriptionText string `json:"descriptionText"` //接口介绍
	InputDemo       string `json:"inputDemo"`       //输入示例
	OutDemo         string `json:"outDemo"`         //输出示例
	Name            string `json:"name"`            //文档名称
	SerialNumber    string `json:"serialNumber"`    //序号
	ProjectId       int64  `json:"projectId"`       //项目的ID
	VersionId       int64  `json:"versionId"`       //版本的ID
}

func FindDoc(catalogueId int64) ([]Doc, error) {
	var ret []Doc
	err := DocDB.OrderBy("(serial_number+0)  asc").Where("catalogue_id=?", catalogueId).Find(&ret)
	if err != nil {
		return nil, err
	}
	return ret, NoDataMsg(len(ret) > 0, "not []doc")
}
func (d *Doc) Save() error {
	_, err := DocDB.Insert(d)
	return err
}

func (d *Doc) Delete() error {
	_, err := DocDB.Where("id=?", d.Id).Delete(Doc{})
	return err
}
func GetDocById(id int64) (*Doc, error) {
	var ret Doc
	b, err := DocDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Doc Not Found Value: %v", id))
	}
	return &ret, nil
}
func GetDocByCatalogueId(id int64) (*Doc, error) {
	var ret Doc
	b, err := DocDB.Where("catalogue_id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Doc Not Found catalogue_id Value: %v", id))
	}
	return &ret, nil
}

func (d *Doc) Update() error {
	_, err := DocDB.Where("id=?", d.Id).Update(d)
	return err
}
