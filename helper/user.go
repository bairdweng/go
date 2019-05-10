package helper

import "errors"

type RelationModel struct {
	CnName string //中文名称。
	EnName string //数据库储存的名称(英文。)
}

// 检查关系
func CheckRelationShip(relationStr string) error {
	var ships = GetRelationShips()
	has := false
	for _, model := range ships {
		if model.EnName == relationStr {
			has = true
		}
	}
	if has == true {
		return nil
	}
	return errors.New("关系不合法")
}

// 检查性别
func CheckGender(gender string) error {
	if gender != "0" &&
		gender != "1" {
		return errors.New("性别不合法")
	}
	return nil
}

// 关系转换
func GetRelationCn(relation string) string {
	relationCn := ""
	var ships = GetRelationShips()
	for _, model := range ships {
		if model.EnName == relation {
			relationCn = model.CnName
		}
	}
	return relationCn
}

func GetRelationEn(relation string) string {
	relationEn := ""
	var ships = GetRelationShips()
	for _, model := range ships {
		if model.CnName == relation {
			relationEn = model.EnName
		}
	}
	return relationEn
}

// 获取所有关系
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

// 关系反向
func ReverseRalationShip(gender int, relation string) string {
	var newRelation string
	// 父母->儿子或女儿
	if relation == "father" || relation == "mother" {
		if gender == 1 {
			newRelation = "son"
		} else {
			newRelation = "daughter"
		}
		// 妻子->丈夫
	} else if relation == "wife" {
		newRelation = "husband"
		// 丈夫->妻子
	} else if relation == "husband" {
		newRelation = "wife"
		// 哥哥，姐姐->弟弟或妹妹
	} else if relation == "old_brother" || relation == "old_sister" {
		if gender == 1 {
			newRelation = "young_brother"
		} else {
			newRelation = "young_sister"
		}
		// 弟弟，妹妹->哥哥，姐姐
	} else if relation == "young_brother" || relation == "young_sister" {
		if gender == 1 {
			newRelation = "old_brother"
		} else {
			newRelation = "old_sister"
		}
		// 儿子，女儿->父亲或母亲
	} else {
		if gender == 1 {
			newRelation = "father"
		} else {
			newRelation = "mother"
		}
	}
	return newRelation
}

//驼峰转下划线
// func HumpConversionUnderline(Field string) string {

// 	return Field
// }

//下划线转驼峰
func UnderlineConversionHump(Field string) string {
	var uppStr string
	vv := []rune(Field)
	for i := 0; i < len(vv); i++ {
		if i == 0 {
			if vv[i] >= 97 && vv[i] <= 122 {
				vv[i] -= 32
				uppStr += string(vv[i])
			} else {
				uppStr += string(vv[i])
			}
		} else {
			if string(vv[i]) != "_" {
				if string(vv[i-1]) == "_" && vv[i] >= 97 && vv[i] <= 122 {
					vv[i] -= 32
					uppStr += string(vv[i])
				} else {
					uppStr += string(vv[i])
				}
			}
		}
	}
	return uppStr
}
