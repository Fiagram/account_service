package grpc

import (
	"context"

	"github.com/Fiagram/account_service/internal/generated/grpc/account_service"
	"github.com/Fiagram/account_service/internal/logic"
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
	request *account_service.CreateAccountRequest,
) (*account_service.CreateAccountResponse, error) {
	output, err := h.accountLogic.CreateAccount(ctx,
		logic.CreateAccountParams{
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

func (h Handler) CheckAccountValid(
	ctx context.Context,
	request *account_service.CheckAccountValidRequest,
) (*account_service.CheckAccountValidResponse, error) {
	output, err := h.accountLogic.CheckAccountValid(ctx,
		logic.CheckAccountValidParams{
			Username: request.Username,
			Password: request.Password,
		})

	if err != nil {
		return nil, err
	}

	return &account_service.CheckAccountValidResponse{
		AccountId: output.AccountId,
	}, nil
}

func (h Handler) IsUsernameTaken(
	ctx context.Context,
	request *account_service.IsUsernameTakenRequest,
) (*account_service.IsUsernameTakenResponse, error) {
	output, err := h.accountLogic.IsUsernameTaken(ctx,
		logic.IsUsernameTakenParams{
			Username: request.Username,
		})
	if err != nil {
		return nil, err
	}
	return &account_service.IsUsernameTakenResponse{
		IsTaken: output.IsTaken,
	}, nil
}

func (h Handler) GetAccount(
	ctx context.Context,
	request *account_service.GetAccountRequest,
) (*account_service.GetAccountResponse, error) {
	output, err := h.accountLogic.GetAccount(ctx,
		logic.GetAccountParams{
			AccountId: request.GetAccountId(),
		},
	)
	if err != nil {
		return nil, err
	}

	return &account_service.GetAccountResponse{
		AccountId: output.AccountId,
		Account: &account_service.AccountInfo{
			Username:    output.AccountInfo.Username,
			Fullname:    output.AccountInfo.Fullname,
			Email:       output.AccountInfo.Email,
			PhoneNumber: output.AccountInfo.PhoneNumber,
			Role:        account_service.AccountInfo_Role(output.AccountInfo.Role),
		},
	}, nil
}
func (h Handler) GetAccountAll(
	ctx context.Context,
	request *account_service.GetAccountAllRequest,
) (*account_service.GetAccountAllResponse, error) {
	output, err := h.accountLogic.GetAccountAll(ctx,
		logic.GetAccountAllParams{},
	)
	if err != nil {
		return nil, err
	}

	accountInfos := make([]*account_service.AccountInfo, 0, len(output.AccountInfos))
	for _, info := range output.AccountInfos {
		accountInfos = append(accountInfos, &account_service.AccountInfo{
			Username:    info.Username,
			Fullname:    info.Fullname,
			Email:       info.Email,
			PhoneNumber: info.PhoneNumber,
			Role:        account_service.AccountInfo_Role(info.Role),
		})
	}

	return &account_service.GetAccountAllResponse{
		AccountIdList:   output.AccountIds,
		AccountInfoList: accountInfos,
	}, nil
}
func (h Handler) GetAccountList(
	ctx context.Context,
	request *account_service.GetAccountListRequest,
) (*account_service.GetAccountListResponse, error) {
	output, err := h.accountLogic.GetAccountList(ctx,
		logic.GetAccountListParams{
			AccountIds: request.GetAccountIdList(),
		},
	)
	if err != nil {
		return nil, err
	}

	accountInfos := make([]*account_service.AccountInfo, 0, len(output.AccountInfos))
	for _, info := range output.AccountInfos {
		accountInfos = append(accountInfos, &account_service.AccountInfo{
			Username:    info.Username,
			Fullname:    info.Fullname,
			Email:       info.Email,
			PhoneNumber: info.PhoneNumber,
			Role:        account_service.AccountInfo_Role(info.Role),
		})
	}

	return &account_service.GetAccountListResponse{
		AccountIdList:   output.AccountIds,
		AccountInfoList: accountInfos,
	}, nil
}
func (h Handler) UpdateAccount(
	ctx context.Context,
	request *account_service.UpdateAccountRequest,
) (*account_service.UpdateAccountResponse, error) {
	output, err := h.accountLogic.UpdateAccount(ctx,
		logic.UpdateAccountParams{
			AccountId: request.GetAccountId(),
			UpdatedAccountInfo: logic.AccountInfo{
				Username:    request.GetUpdatedAccountInfo().GetUsername(),
				Fullname:    request.GetUpdatedAccountInfo().GetFullname(),
				Email:       request.GetUpdatedAccountInfo().GetEmail(),
				PhoneNumber: request.GetUpdatedAccountInfo().GetPhoneNumber(),
				Role:        logic.Role(request.GetUpdatedAccountInfo().GetRole()),
			},
		},
	)
	if err != nil {
		return nil, err
	}

	return &account_service.UpdateAccountResponse{
		AccountId: output.AccountId,
	}, nil
}

func (h Handler) DeleteAccount(
	ctx context.Context,
	request *account_service.DeleteAccountRequest,
) (*account_service.DeleteAccountResponse, error) {
	err := h.accountLogic.DeleteAccount(ctx,
		logic.DeleteAccountParams{
			Username: request.GetUsername(),
		})
	if err != nil {
		return nil, err
	}

	return &account_service.DeleteAccountResponse{
		Username: request.GetUsername(),
	}, nil
}
