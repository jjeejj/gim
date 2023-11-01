package api

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"gim/pkg/protocol/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func getBusinessExtClient() pb.BusinessExtClient {
	conn, err := grpc.Dial("127.0.0.1:8020", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return pb.NewBusinessExtClient(conn)
}

func getCtx() context.Context {
	token := "0"
	return metadata.NewOutgoingContext(context.TODO(), metadata.Pairs(
		"user_id", "1",
		"device_id", "1",
		"token", token,
		"request_id", strconv.FormatInt(time.Now().UnixNano(), 10)))
}

func TestUserExtServer_SignIn(t *testing.T) {
	resp, err := getBusinessExtClient().SignIn(getCtx(), &pb.SignInReq{
		PhoneNumber: "22222222222",
		Code:        "0",
		DeviceId:    1,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestUserExtServer_GetUser(t *testing.T) {
	resp, err := getBusinessExtClient().GetUser(getCtx(), &pb.GetUserReq{UserId: 1})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestBusinessExtServer_UpdateUser(t *testing.T) {
	_, err := getBusinessExtClient().UpdateUser(getCtx(), &pb.UpdateUserReq{
		Nickname:  "iyi",
		Sex:       2,
		AvatarUrl: "https://p9-passport.byteacctimg.com/img/user-avatar/73efcc5a1c6617d447257cb997bd4ffc~40x40.awebp",
	})
	if err != nil {
		fmt.Println(err)
	}
}

func TestBusinessExtServer_SearchUser(t *testing.T) {
	resp, err := getBusinessExtClient().SearchUser(getCtx(), &pb.SearchUserReq{
		Key: "i",
	})
	if err != nil {
		fmt.Println(err)
	}
	for i, user := range resp.Users {
		fmt.Printf("i: %d, user: %+v\n", i, user)
	}
}
