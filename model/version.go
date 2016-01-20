package model

import (
	"errors"
	"fmt"
)

type Version struct {
	Id        int64  `json:"versionId"`
	Version   string `json:"version"`   //版本号
	ProjectId int64  `json:"projectId"` //项目的ID
}

//保存单条信息
func (v *Version) Save() error {
	_, err := DocDB.Insert(v)
	return err
}

//根据ID删除
func (v *Version) Delete() error {
	_, err := DocDB.Where("id=?", v.Id).Delete(Version{})
	return err
}

//根据ID查询
func GetVersionById(id int64) (*Version, error) {
	var ret Version
	b, err := DocDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Version Not Found Value: %v", id))
	}
	return &ret, nil
}

//根据ID更新数据
func (v *Version) Update() error {
	_, err := DocDB.Where("id=?", v.Id).Update(v)
	return err
}

//Version查询数据分页
func FindVersion(projectId int64) ([]Version, error) {
	var ret []Version
	err := DocDB.Where("project_id=?", projectId).Find(&ret)
	if err != nil {
		return nil, err
	}
	if len(ret) == 0 {
		return nil, errors.New("Version [] null")
	}
	return ret, nil
}
