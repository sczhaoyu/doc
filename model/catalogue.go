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
	ProjectId    int64  `json:"projectId"`    //项目的ID
	VersionId    int64  `json:"versionId"`    //版本的ID
}

//删除目录
func DeleteCatalogue(id int64) error {
	session := DocDB.NewSession()
	defer session.Close()
	session.Begin() //开启事务
	//删除目录
	_, cErr := session.Where("id=?", id).Delete(Catalogue{})
	if cErr != nil {
		session.Rollback()
		return cErr
	}
	doc, err := FindDoc(id)
	var docId []int64 = make([]int64, 0, len(doc))
	//删除文档
	if err == nil {
		for i := 0; i < len(doc); i++ {
			docId = append(docId, doc[i].Id)
			_, docDelerr := session.Where("id=?", doc[i].Id).Delete(Doc{})
			if docDelerr != nil {
				session.Rollback()
				return docDelerr
			}
		}
	}
	//删除参数
	if len(docId) > 0 {
		_, prmDelErr := session.In("doc_id", docId).Delete(Parameters{})
		if prmDelErr != nil {
			session.Rollback()
			return prmDelErr
		}
	}
	return session.Commit()

}

//获取全部目录
func FindCatalogueAll(projectId, versionId int64) ([]Catalogue, error) {
	var ret []Catalogue
	err := DocDB.Where("project_id=? and version_id=?", projectId, versionId).OrderBy("(serial_number+0)  asc").Find(&ret)
	if err != nil {
		return nil, err
	}
	return ret, NoDataMsg(len(ret) > 0, "catalogue[] null")
}

//获取全部父目录
func FindCatalogueAllParent(projectId, versionId int64) ([]Catalogue, error) {
	var ret []Catalogue
	err := DocDB.Where("parent_id=? and project_id=? and version_id=?", 0, projectId, versionId).OrderBy("(serial_number+0)  asc").Find(&ret)
	if err != nil {
		return nil, err
	}
	return ret, NoDataMsg(len(ret) > 0, "catalogue[] null")
}

//获取子目录
func FindFindCatalogueByParent(projectId, versionId, catalogueId int64) ([]Catalogue, error) {
	var ret []Catalogue
	err := DocDB.Where("parent_id=?  and project_id=? and version_id=?", catalogueId, projectId, versionId).Find(&ret)
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
