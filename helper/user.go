package helper

import "errors"

type RelationModel struct {
	CnName string //中文名称。
	EnName string //数据库储存的名称(英文。)
}

func CheckRelationShip(relationStr string) error {
	if relationStr != "father" &&
		relationStr != "mother" &&
		relationStr != "wife" &&
		relationStr != "elder_brother" &&
		relationStr != "sister" &&
		relationStr != "son" &&
		relationStr != "daughter" {
		return errors.New("关系不合法")
	}
	return nil
}
func CheckGender(gender string) error {
	if gender != "0" &&
		gender != "1" {
		return errors.New("性别不合法")
	}
	return nil
}
func GetRelationShips() []RelationModel {

	var rr = []RelationModel{
		RelationModel{CnName: "父亲", EnName: "father"},
		RelationModel{CnName: "母亲", EnName: "mother"},
		RelationModel{CnName: "妻子", EnName: "wife"},
		RelationModel{CnName: "丈夫", EnName: "husband"},
		RelationModel{CnName: "哥哥", EnName: "old_brother"},
		RelationModel{CnName: "弟弟", EnName: "young_brother"},
		RelationModel{CnName: "姐姐", EnName: "old_sister"},
		RelationModel{CnName: "妹妹", EnName: "young_sister"},
		RelationModel{CnName: "儿子", EnName: "son"},
		RelationModel{CnName: "女儿", EnName: "daughter"},
	}
	return rr
}
