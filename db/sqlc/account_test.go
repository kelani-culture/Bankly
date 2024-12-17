package sqlc

import (
	"bankly/util"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func CreateAccount(t *testing.T) Account {

	args := CreateAccountParams{
		Owner:    util.GenerateOwnerName(),
		Balance:  util.GenerateRandomAmount(),
		Currency: util.GenerateRandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Owner, args.Owner)
	require.Equal(t, account.Balance, args.Balance)
	require.Equal(t, account.Currency, args.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	CreateAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := CreateAccount(t)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.Equal(t, account2.Owner, account1.Owner)
	require.Equal(t, account2.Balance, account1.Balance)
	require.Equal(t, account2.Currency, account1.Currency)

	require.NotZero(t, account2.ID)
	require.NotZero(t, account2.CreatedAt)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	account1 := CreateAccount(t)

	args := updateAccountParams{
		ID:      account1.ID,
		Balance: util.GenerateRandomAmount(),
	}

	account2, err := testQueries.updateAccount(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, account2)
	require.NotZero(t, account2.ID)
	require.NotZero(t, account2.CreatedAt)
	// require.Equal(t, account2.Owner, args.Owner)
	require.Equal(t, account2.Balance, args.Balance)
	require.Equal(t, account2.ID, account1.ID)

	// require.Equal(t, account2.Currency, args.Currency)

	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account1 := CreateAccount(t)

	delAccount := testQueries.deleteAccount(context.Background(), account1.ID)

	// require.NoError(t , err)
	require.NoError(t, delAccount)
	require.Empty(t, delAccount)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)
	require.Empty(t, account2)
}

func TestGetListAccount(t *testing.T) {
	for i := 10; i < 11; i++ {
		CreateAccount(t)
	}

	listAccounts, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{
		Limit:  5,
		Offset: 5,
	})

	require.NoError(t, err)
	require.Len(t, listAccounts, 5)

	for _, account := range listAccounts {
		require.NotEmpty(t, account)
		require.NotZero(t, account.ID)
		require.NotZero(t, account.CreatedAt)
	}
}
