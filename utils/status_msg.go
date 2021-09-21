package utils

const (
	CodeSuccess = 200
	CodeError   = 500

	/*
		code = 1xxx,zhihu
	*/
	CodeZhihuErr           = 1000
	CodeZhihuFavorsErr     = 1001
	CodeZhihuFavorItemsErr = 1002

	CodeZhihuInvalidParamErr = 1101
	CodeZhihuNotFoundErr     = 1102

	/*
		code=2xxx, bilibili
	*/
	CodeBiliErr           = 2000
	CodeBiliFavorsErr     = 2001
	CodeBiliFavorItemsErr = 2002

	CodeBiliInvalidParamErr = 2101
	CodeBiliNotFoundErr     = 2102

	/*
		code=3xxx, pinbox
	*/
	CodePinboxErr           = 3000
	CodePinboxFavorsErr     = 3001
	CodePinboxFavorItemsErr = 3002

	CodePinboxInvalidParamErr = 3101
	CodePinboxNotFoundErr     = 3102
)

var codeMsg = map[int]string{
	CodeSuccess: "Succ",
	CodeError:   "系统内部错误",

	CodeZhihuErr:           "获取知乎信息错误",
	CodeZhihuFavorsErr:     "获取知乎所有收藏夹信息错误",
	CodeZhihuFavorItemsErr: "获取知乎收藏夹items错误",

	CodeZhihuInvalidParamErr: "知乎-非法参数",
	CodeZhihuNotFoundErr:     "知乎-对应参数资源不存在",

	CodeBiliErr:           "获取bilibili信息错误",
	CodeBiliFavorsErr:     "获取bilibili所有收藏夹信息错误",
	CodeBiliFavorItemsErr: "获取bilibili收藏夹items错误",

	CodeBiliInvalidParamErr: "bilibili-非法参数",
	CodeBiliNotFoundErr:     "bilibili-对应参数资源不存在",

	CodePinboxErr:           "获取pinbox信息错误",
	CodePinboxFavorsErr:     "获取pinbox所有收藏夹信息错误",
	CodePinboxFavorItemsErr: "获取pinbox收藏夹items错误",

	CodePinboxInvalidParamErr: "pinbox-非法参数",
	CodePinboxNotFoundErr:     "pinbox-对应参数资源不存在",
}

func GetStatusMsg(statusCode int) string {
	s, exist := codeMsg[statusCode]
	if !exist {
		return "UNKOWN STATUS CODE!!"
	}
	return s
}
