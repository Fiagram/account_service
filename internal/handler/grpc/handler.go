package grpc

import (
	"context"

	"github.com/Fiagram/account_service/internal/generated/grpc/account_service"
	"github.com/Fiagram/account_service/internal/logic"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Handler struct {
	account_service.UnimplementedAccountServiceServer
	accountLogic logic.Account
}

func NewHandler(
	accountLogic logic.Account,
) account_service.AccountServiceServer {
	return &Handler{
		accountLogic: accountLogic,
	}
}

func (h Handler) CreateAccount(
	ctx context.Context,
	request *account_service.CreateAccountRequest) (
	response *account_service.CreateAccountResponse,
	err error,
) {
	output, err := h.accountLogic.CreateAccount(ctx, logic.CreateAccountParams{
		AccountInfo: logic.AccountInfo{
			Username:    request.GetAccountInfo().GetUsername(),
			Fullname:    request.GetAccountInfo().GetFullname(),
			Email:       request.GetAccountInfo().GetEmail(),
			PhoneNumber: request.GetAccountInfo().GetPhoneNumber(),
			Role:        logic.Role(request.GetAccountInfo().GetRole()),
		},
		Password: request.GetPassword(),
	})
	if err != nil {
		return nil, err
	}

	return &account_service.CreateAccountResponse{
		AccountId: output.AccountId,
	}, nil
}

func (h Handler) CheckAccountValid(context.Context,
	*account_service.CheckAccountValidRequest,
) (*account_service.CheckAccountValidResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method CheckAccountValid not implemented")
}
func (h Handler) GetAccount(context.Context,
	*account_service.GetAccountRequest,
) (*account_service.GetAccountResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetAccount not implemented")
}
func (h Handler) GetAccountAll(context.Context,
	*account_service.GetAccountAllRequest,
) (*account_service.GetAccountAllResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetAccountAll not implemented")
}
func (h Handler) GetAccountList(context.Context,
	*account_service.GetAccountListRequest,
) (*account_service.GetAccountListResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method GetAccountList not implemented")
}
func (h Handler) UpdateAccount(context.Context,
	*account_service.UpdateAccountRequest,
) (*account_service.UpdateAccountResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method UpdateAccount not implemented")
}
func (h Handler) DeleteAccount(context.Context,
	*account_service.DeleteAccountRequest,
) (*account_service.DeleteAccountResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method DeleteAccount not implemented")
}
