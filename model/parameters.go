package model

import (
	"errors"
	"fmt"
)

type Parameters struct {
	Id              int64  `json:"parameterId"`
	Name            string `json:"name"`            //参数名称
	DataType        string `json:"dataType"`        //数据类型
	DescriptionText string `json:"descriptionText"` //参数描述
	Required        int    `json:"required"`        //是否必选参数，0否，1是
	DocId           int64  `json:"docId"`           //所属文档,doc表外键
	PrmType         int    `json:"prmType"`         //0请求参数1响应参数
	Length          string `json:"length"`          //数据长度
	SerialNumber    string `json:"serialNumber"`    //序号
}

func SaveParameters(ret []Parameters) error {
	_, err := DocDB.Insert(ret)
	return err
}
func FindParametersByDocId(docId int64) ([]Parameters, error) {
	var ret []Parameters
	err := DocDB.Where("doc_id=?", docId).Find(&ret)
	if err != nil {
		return nil, err
	}
	return ret, NoDataMsg(len(ret) > 0, fmt.Sprintf("Parameters Not Found docId: %v", docId))
}
func (p *Parameters) Save() error {
	_, err := DocDB.Insert(p)
	return err
}

func (p *Parameters) Delete() error {
	_, err := DocDB.Where("id=?", p.Id).Delete(Parameters{})
	return err
}
func GetParametersById(id int64) (*Parameters, error) {
	var ret Parameters
	b, err := DocDB.Where("id=?", id).Get(&ret)
	if err != nil {
		return nil, err
	}
	if b == false {

		return nil, errors.New(fmt.Sprintf("Parameters Not Found Value: %v", id))
	}
	return &ret, nil
}

func (p *Parameters) Update() error {
	_, err := DocDB.Where("id=?", p.Id).Update(p)
	return err
}
