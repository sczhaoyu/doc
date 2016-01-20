package model

import (
	"errors"
	"fmt"
)

type Project struct {
	Id   int64  `json:"projectId"`
	Name string `json:"name"` //项目名称

}

//保存单条信息
func (p *Project) Save() error {
	_, err := DocDB.Insert(p)
	return err
}

//根据ID删除
func (p *Project) Delete() error {
	_, err := DocDB.Where("id=?", p.Id).Delete(Project{})
	return err
}

//根据ID查询
func GetProjectById(id int64) (*Project, error) {
	var ret Project
	b, err := DocDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Project Not Found Value: %v", id))
	}
	return &ret, nil
}

//根据ID更新数据
func (p *Project) Update() error {
	_, err := DocDB.Where("id=?", p.Id).Update(p)
	return err
}

//Project查询数据分页
func FindProject() ([]Project, error) {
	var ret []Project
	err := DocDB.Find(&ret)
	if err != nil {
		return nil, err
	}
	if len(ret) == 0 {
		return nil, errors.New("Project [] null")
	}
	return ret, nil
}
