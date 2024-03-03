package dto

type ApiResponse[T any] struct {
	ResponseKey     string `json:"response_key"`
	ResponseMessage string `json:"response_message"`
	Data            T      `json:"data"`
}

type ApiPaginationResponse[T any] struct {
	ResponseKey     string `json:"response_key"`
	ResponseMessage string `json:"response_message"`
	Data            T      `json:"data"`
	TotalPage       int64  `json:"totalPage"`
	Page            int    `json:"page"`
	PageSize        int    `json:"pageSize"`
}
