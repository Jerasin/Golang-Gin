package response

type BaseResponse struct {
	ResponseKey     string `json:"response_key"`
	ResponseMessage string `json:"response_message"`
}

type PaginationResponse struct {
	BaseResponse
	TotalPage int `json:"totalPage"`
	Page      int `json:"page"`
	PageSize  int `json:"pageSize"`
}

type UpdateDataResponse struct {
	BaseResponse
	Message string `json:"message" example:"update success"`
}

type CreateDataResponse struct {
	BaseResponse
	Message string `json:"message" example:"create success"`
}

type DeleteDataResponse struct {
	BaseResponse
	Message string `json:"message" example:"delete success"`
}
