package routers

type CallBack struct {
	Statues int
	Data    interface{}
	Message string
}

func Successful(data interface{}) interface{} {
	var callBack = CallBack{}
	callBack.Data = data
	callBack.Statues = 1
	callBack.Message = "成功"
	return callBack
}

func Error(message string, data interface{}) interface{} {
	var callBack = CallBack{}
	callBack.Data = data
	callBack.Statues = 0
	callBack.Message = message
	return callBack
}
