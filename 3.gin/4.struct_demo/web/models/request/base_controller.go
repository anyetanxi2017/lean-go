package request

type BaseController struct {
}

type BaseQuery struct {
	PageNo   int `form:"page_no" json:"page_no"`
	PageSize int `form:"page_size" json:"page_size"`
}

func (b *BaseController) DefaultPage(key int) int {
	defaultValue := 1
	if key != 0 {
		return key
	}
	return defaultValue
}

func (b *BaseController) DefaultPageSize(key int) int {
	defaultValue := 10
	if key != 0 {
		return key
	}
	return defaultValue
}
