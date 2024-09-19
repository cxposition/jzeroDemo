package main

import (
	"admin/admin-go"
	"github.com/jzero-io/restc"
)

func main() {
	headers := make(map[string][]string)
	headers["Content-Type"] = []string{"application/json"}

	clientset, err := admin_go.NewClientWithOptions(
		restc.WithAddr("127.0.0.1"),
		restc.WithPort("8001"),
		restc.WithProtocol("http"),
		restc.WithHeaders(headers),
	)
	if err != nil {
		panic(err)
	}

	//listResult, err := clientset.User().List(context.Background(), types.ListUserRequest{Page: 1, Size: 10})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(*listResult)

	//addResult, err := clientset.User().Add(context.Background(), types.AddUserRequest{Name: "chenxia"})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(addResult)

	//modifyResult, err := clientset.User().Modify(context.Background(), types.ModifyUserRequest{Id: 1, Name: "chenxia"})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(modifyResult)

	//deleteResult, err := clientset.User().Delete(context.Background(), types.DeleteUserRequest{Id: 1})
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(deleteResult)
	//uploadRequest, err := clientset.File().Upload()
	//if err != nil {
	//	panic(err)
	//}
	//stream, err := uploadRequest.Stream()
	//if err != nil {
	//	panic(err)
	//}
}
