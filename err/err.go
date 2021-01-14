package err

var (
	// APIInitErr api 未初始化
	APIInitErr = "the API is not initialized "
	// APIPathErr api路径有问题
	APIPathErr = "there is a problem with the API path, please check the API"
	// APIServerNotExist 服务不存在
	APIServerNotExist = "the API server is not exist, please make sure the server"
	// APIServiceNotExist API服务不存在
	APIServiceNotExist = "the API service does not exist, please make sure the API"
	// APIReturnErr 接口未返回批量接口实现
	APIReturnErr = "the api does not return MVC.IActionResult"
)
