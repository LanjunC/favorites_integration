package utils

const (
	CodeSuccess = 200
	CodeError = 500

	/*
	code = 1xxx,zhihu
	*/
	CodeZhihuErr = 1000
	CodeZhihuFavorsErr = 1001
	CodeZhihuFavorItemsErr = 1002


	CodeZhihuInvalidParam = 1101
	CodeZhihuNotFoundErr = 1102
)

	var codeMsg =  map[int]string {
		CodeSuccess : "Succ",
		CodeError: "系统内部错误",

		CodeZhihuErr : "获取知乎信息错误",
		CodeZhihuFavorsErr : "获取知乎所有收藏夹信息错误",
		CodeZhihuFavorItemsErr : "获取知乎收藏夹items错误",

		CodeZhihuInvalidParam : "知乎-非法参数",
		CodeZhihuNotFoundErr : "知乎-对应参数资源不存在",
}

func GetStatusMsg(statusCode int) string{
	s,exist := codeMsg[statusCode]
	if !exist {
		return "UNKOWN STATUS CODE!!"
	}
	return s
}