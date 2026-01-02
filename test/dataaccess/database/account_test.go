package database_test

import (
	"context"
	"testing"

	"github.com/Fiagram/account_service/internal/dataaccess/database"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	asor := database.NewAccountAccessor(sqlDb, logger)
	input := RandomAccount()
	id, err := asor.CreateAccount(context.Background(), input)

	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.NotZero(t, id)

	errD := asor.DeleteAccountByUsername(context.Background(), input.Username)
	require.NoError(t, errD)
}

func TestGetAccountById(t *testing.T) {
	asor := database.NewAccountAccessor(sqlDb, logger)
	input := RandomAccount()
	id, err := asor.CreateAccount(context.Background(), input)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	acc, err := asor.GetAccountById(context.Background(), id)
	require.NoError(t, err)
	require.NotEmpty(t, acc)
	require.Equal(t, input.Username, acc.Username)
	require.Equal(t, input.Fullname, acc.Fullname)
	require.Equal(t, input.Email, acc.Email)
	require.Equal(t, input.PhoneNumber, acc.PhoneNumber)
	require.Equal(t, input.RoleId, acc.RoleId)

	errD := asor.DeleteAccountByUsername(context.Background(), input.Username)
	require.NoError(t, errD)
}

func TestGetAccountByUsername(t *testing.T) {
	asor := database.NewAccountAccessor(sqlDb, logger)
	input := RandomAccount()
	id, err := asor.CreateAccount(context.Background(), input)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	acc, err := asor.GetAccountByUsername(context.Background(), input.Username)
	require.NoError(t, err)
	require.NotEmpty(t, acc)
	require.Equal(t, input.Username, acc.Username)
	require.Equal(t, input.Fullname, acc.Fullname)
	require.Equal(t, input.Email, acc.Email)
	require.Equal(t, input.PhoneNumber, acc.PhoneNumber)
	require.Equal(t, input.RoleId, acc.RoleId)

	errD := asor.DeleteAccountByUsername(context.Background(), input.Username)
	require.NoError(t, errD)
}

func TestDeleteAccountById(t *testing.T) {
	asor := database.NewAccountAccessor(sqlDb, logger)
	input := RandomAccount()
	id, err := asor.CreateAccount(context.Background(), input)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	errD := asor.DeleteAccountById(context.Background(), id)
	require.NoError(t, errD)
}

func TestDeleteAccountByUsername(t *testing.T) {
	asor := database.NewAccountAccessor(sqlDb, logger)
	input := RandomAccount()
	id, err := asor.CreateAccount(context.Background(), input)
	require.NoError(t, err)
	require.NotEmpty(t, id)

	errD := asor.DeleteAccountByUsername(context.Background(), input.Username)
	require.NoError(t, errD)
}

func TestUpdateAccount(t *testing.T) {
	asor := database.NewAccountAccessor(sqlDb, logger)
	in1 := RandomAccount()
	id1, err1 := asor.CreateAccount(context.Background(), in1)
	require.NoError(t, err1)
	require.NotEmpty(t, id1)
	require.NotZero(t, id1)

	in2 := in1
	in2.Fullname = RandomVnPersonName()
	in2.Email = RandomGmailAddress()
	in2.PhoneNumber = RandomVnPhoneNum()
	in2.RoleId = 1
	err2 := asor.UpdateAccount(context.Background(), in2)
	require.NoError(t, err2)

	updatedAcc, errU := asor.GetAccountByUsername(context.Background(), in1.Username)
	require.NoError(t, errU)
	require.NotEmpty(t, updatedAcc)
	require.Equal(t, in2.Username, updatedAcc.Username)
	require.Equal(t, in2.Fullname, updatedAcc.Fullname)
	require.Equal(t, in2.Email, updatedAcc.Email)
	require.Equal(t, in2.PhoneNumber, updatedAcc.PhoneNumber)
	require.Equal(t, in2.RoleId, updatedAcc.RoleId)

	errD := asor.DeleteAccountByUsername(context.Background(), in1.Username)
	require.NoError(t, errD)
}
