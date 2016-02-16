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

func CopyVersion(projectId, oldVersionId, versionId int64) error {
	session := DocDB.NewSession()
	defer session.Close()
	session.Begin() //开启事务
	//===复制文章
	explains, err := FindExplain(projectId, oldVersionId)
	if err == nil {
		//开始复制文章
		for i := 0; i < len(explains); i++ {
			explains[i].VersionId = versionId
			explains[i].Id = 0
			_, explainsErr := session.Insert(explains[i])
			if explainsErr != nil {
				session.Rollback()
				return explainsErr
			}
		}
	}
	//===复制目录
	c, cErr := FindCatalogueAll(projectId, oldVersionId)
	if cErr == nil {
		for i := 0; i < len(c); i++ {
			//===获取需要复制文档
			docs, docErr := FindDoc(c[i].Id) //获取目录的文档
			c[i].VersionId = versionId

			c[i].Id = 0 //将ID设置为0
			_, catalogueErr := session.Insert(&c[i])
			if catalogueErr != nil {
				//出错 回滚
				session.Rollback()
				return catalogueErr
			}
			//---开始复制文档
			if docErr == nil {
				for j := 0; j < len(docs); j++ {
					//获取需要复制的参数
					docId := docs[j].Id
					prms, prmsErr := FindParametersByDocId(docId)
					docs[j].CatalogueId = c[i].Id
					docs[j].Id = 0                //将ID设置为0
					docs[j].VersionId = versionId //设置文档的版本
					_, docInsertErr := session.Insert(docs[j])
					if docInsertErr != nil {
						//出错 回滚
						session.Rollback()
						return docInsertErr
					}
					//开始复制参数
					if prmsErr == nil {

						for k := 0; k < len(prms); k++ {
							prms[k].Id = 0
							prms[k].DocId = docId
							prms[k].VersionId = versionId
							_, prmInsertErr := session.Insert(prms[k])
							if prmInsertErr != nil {
								//出错 回滚
								session.Rollback()
								return prmInsertErr
							}
						}
					}

				}
			}

		}

	}
	return session.Commit()

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
