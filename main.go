package main

import "fmt"

//go:generate stringer -type=HttpErrorCode -linecomment
type HttpErrorCode int

const (
	HTTP_OK HttpErrorCode = 200 // 访问成功
	HTTP_NOT_FOUNT HttpErrorCode = 404 // 404 不存在
	HTTP_SERVICE_UNAVAILABLE HttpErrorCode = 503 // 503 服务端错误
)

func main() {
	fmt.Println(HTTP_OK)
	fmt.Println(HTTP_OK == 201)
}