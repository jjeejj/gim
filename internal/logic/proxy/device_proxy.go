package proxy

import (
	"context"

	"gim/pkg/protocol/pb"
)

type deviceProxy interface {
	ListOnlineByUserId(ctx context.Context, userId string) ([]*pb.Device, error)
}

var DeviceProxy deviceProxy
