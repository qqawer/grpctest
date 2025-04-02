package handler

import (
	"context"
	"grpctest/global"
	"grpctest/model/proto"
	model "grpctest/model/user"
	"grpctest/utils"
)

type UserServer struct {
	proto.UnimplementedUserServer // 关键嵌入
}

func ModelToResponse(user model.User) proto.UserInfoResponse {
	//在grpc的message中字段有默认值，你就不能随便赋值nil进去，容易出错
	//这里要搞清，哪些字段是有默认值的
	UserInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		UserInfoRsp.Birthday = uint64(user.Birthday.Unix())
	}
}

func (s *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	//获取用户列表
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)

	global.DB.Scopes(utils.Paginate(int(req.Pn), int(req.PSize))).Find(&users)

	for _, user := range users {
		userInfoRsp := ModelToResponse(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp,nil
}
