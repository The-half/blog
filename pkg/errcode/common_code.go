package errcode

var (
	Success = NewError(0, "成功")
	ServerError = NewError(10000000, "服务器内部出现问题")
	InvalidParams = NewError(10000001,"入参测试")
	NotFound = NewError(10000002, "找不到相关数据")

	UnauthorizedAuthNotExist = NewError(10000003, "鉴权失败,找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError = NewError(10000004, "鉴权失败，token错误")
	UnauthorizedTokenTimeout = NewError(10000005, "鉴权失败, token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败,　token生成失败")
	TooManyRequests = NewError(10000007, "请求太多")

	ErrorGetTagListFail = NewError(20010001, "获取标签列表失败")
	ErrorCreateTagFail = NewError(20010002,"创建标签失败")
	ErrorUpdateTagFail = NewError(20010003, "更新标签失败")
	ErrorDeleteTagFail = NewError(20010004, "删除标签失败")
	ErrorCountTagFail = NewError(20010005, "统计标签失败")

	ErrorGetArticleFail = NewError(20020001,"获取文章失败")
	ErrorGetArticlesFail = NewError(20020002, "获取多个文章失败")
	ErrorCreateArticleFail = NewError(20020003,"创建文章失败")
	ErrorUpdateArticleFail = NewError(20020004, "更新文章失败")
	ErrorDeleteArticleFail = NewError(20020005, "删除文章失败")

	ErrorUploadFileFail = NewError(20030001, "上传文件失败")


	//ERROR_GET_TAG_LIST_FAIL = NewError(30010001, "获取tag列表失败")

)