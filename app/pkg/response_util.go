package pkg

import (
	"fmt"

	"github.com/Jerasin/app/constant"
	"github.com/Jerasin/app/dto"
	"github.com/jinzhu/copier"
)

func Null() interface{} {
	return nil
}

func ModelDump(res interface{}, data interface{}) error {
	err := copier.Copy(res, data)
	if err != nil {
		return fmt.Errorf("error copying data: %v", err)
	}

	return nil
}

func CreateResponse() map[string]string {
	return map[string]string{
		"message": "create success",
	}
}

func UpdateResponse() map[string]string {
	return map[string]string{
		"message": "update success",
	}
}

func DeleteResponse() map[string]string {
	return map[string]string{
		"message": "delete success",
	}
}

func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) dto.ApiResponse[T] {
	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](status string, message string, data T) dto.ApiResponse[T] {
	return dto.ApiResponse[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}
}

func BuildPaginationResponse[T any](responseStatus constant.ResponseStatus, data T, totalPage int64, page int, pageSize int) dto.ApiPaginationResponse[T] {
	return BuildPaginationResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data, totalPage, page, pageSize)
}

func BuildPaginationResponse_[T any](status string, message string, data T, totalPage int64, page int, pageSize int) dto.ApiPaginationResponse[T] {
	return dto.ApiPaginationResponse[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
		TotalPage:       totalPage,
		Page:            page,
		PageSize:        pageSize,
	}
}
