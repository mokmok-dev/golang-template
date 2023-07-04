package marshaller

import (
	"github.com/mokmok-dev/golang-template/domain/model"
	proto "github.com/mokmok-dev/golang-template/proto/golang-template/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func UserToProto(model *model.User) *proto.User {
	return &proto.User{
		Id:        model.ID.String(),
		CreatedAt: timestamppb.New(model.CreatedAt),
		UpdatedAt: timestamppb.New(model.UpdatedAt),
	}
}
