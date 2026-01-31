package logic

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrTxCommitFailed = status.Error(codes.Internal, "failed to commit")
	ErrTxBeginFailed  = status.Error(codes.Internal, "failed to take a transaction up")
)
