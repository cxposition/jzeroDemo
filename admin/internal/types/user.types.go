// Code generated by goctl-types plugin. DO NOT EDIT.
package types

import (
	"time"
)

var (
	_ = time.Now()
)

type AddUserRequest struct {
	Name string `json:"name"`
}

type AddUserResponse struct {
	Status int `json:"status"`
}

type DeleteUserRequest struct {
	Id int `json:"id"`
}

type DeleteUserResponse struct {
	Status int `json:"status"`
}

type ListUserRequest struct {
	Page int `form:"page"` // 第几页
	Size int `form:"size"` // 每页数量
}

type ListUserResponse struct {
	List  []User `json:"list"`
	Total int    `json:"total"`
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ModifyUserRequest struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ModifyUserResponse struct {
	Status int `json:"status"`
}
