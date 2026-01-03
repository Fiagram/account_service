package logic

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrTxCommitFailed = status.Error(codes.Internal, "failed to commit")
	ErrTxInitFailed   = status.Error(codes.Internal, "failed to take a transaction up")
)
