package database_test

import (
	"context"
	"testing"

	"github.com/Fiagram/account_service/internal/dataaccess/database"
	"github.com/stretchr/testify/require"
)

func TestCreateAndDeleteAccountPassword(t *testing.T) {
	aAsor := database.NewAccountAccessor(sqlDb, logger)
	pAsor := database.NewAccountPasswordAccessor(sqlDb, logger)
	ctx := context.Background()

	acc := RandomAccount()
	id, err := aAsor.CreateAccount(ctx, acc)
	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.NotZero(t, id)

	input := database.AccountPassword{
		OfAccountId:  id,
		HashedString: RandomString(128),
	}
	require.NoError(t, pAsor.CreateAccountPassword(ctx, input))

	require.NoError(t, pAsor.DeleteAccountPassword(ctx, id))
	require.NoError(t, aAsor.DeleteAccountById(ctx, id))
}

func TestGetAccountPassword(t *testing.T) {
	aAsor := database.NewAccountAccessor(sqlDb, logger)
	pAsor := database.NewAccountPasswordAccessor(sqlDb, logger)
	ctx := context.Background()

	acc := RandomAccount()
	id, err := aAsor.CreateAccount(ctx, acc)
	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.NotZero(t, id)

	input := database.AccountPassword{
		OfAccountId:  id,
		HashedString: RandomString(128),
	}
	require.NoError(t, pAsor.CreateAccountPassword(ctx, input))
	output, err := pAsor.GetAccountPassword(ctx, id)
	require.NoError(t, err)
	require.Equal(t, input.OfAccountId, output.OfAccountId)
	require.Equal(t, input.HashedString, output.HashedString)

	require.NoError(t, pAsor.DeleteAccountPassword(ctx, id))
	require.NoError(t, aAsor.DeleteAccountById(ctx, id))
}

func TestUpdateAccountPassword(t *testing.T) {
	aAsor := database.NewAccountAccessor(sqlDb, logger)
	pAsor := database.NewAccountPasswordAccessor(sqlDb, logger)
	ctx := context.Background()

	acc := RandomAccount()
	id, err := aAsor.CreateAccount(ctx, acc)
	require.NoError(t, err)
	require.NotEmpty(t, id)
	require.NotZero(t, id)

	input := database.AccountPassword{
		OfAccountId:  id,
		HashedString: RandomString(128),
	}
	require.NoError(t, pAsor.CreateAccountPassword(ctx, input))

	updatedInput := input
	updatedInput.HashedString = RandomString(128)
	require.NoError(t, pAsor.UpdateAccountPassword(ctx, updatedInput))

	output, err := pAsor.GetAccountPassword(ctx, id)
	require.NoError(t, err)
	require.Equal(t, updatedInput.OfAccountId, output.OfAccountId)
	require.Equal(t, updatedInput.HashedString, output.HashedString)

	require.NoError(t, pAsor.DeleteAccountPassword(ctx, id))
	require.NoError(t, aAsor.DeleteAccountById(ctx, id))
}
