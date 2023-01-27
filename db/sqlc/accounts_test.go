package db

import (
	"context"
	"log"
	"testing"
	"time"

	"user-model-db/util"

	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Accounts {
	arg := CreateAccountParams{
		// reference : https://stackoverflow.com/questions/60792313/unable-to-use-type-string-as-sql-nullstring
		Owner:    util.RandomOwner(),
		Balance:  util.RandomInt(0, 3),
    Currency: util.RandomString(1),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)      // check the error must be nil
	require.NotEmpty(t, account) // user should be not an empty object

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

  return account
}

func TestCreateAccount(t *testing.T) {
  createRandomAccount(t);
}

func TestGetOneAccounts(t *testing.T) {
  createAccounts := createRandomAccount(t);
  _, err := testQueries.GetAccount(context.Background(), createAccounts.ID)
  if (err != nil) {
    log.Fatal(err)
  }
}

func TestListAccounts(t *testing.T) {
  arg := ListAccountsParams {
    Limit: 1,
    Offset: 2,
  }

  allAccounts, err := testQueries.ListAccounts(context.Background(), arg) 
  require.NoError(t, err)
  require.Len(t, allAccounts, len(allAccounts))
}

func TestUpdateAccount(t *testing.T) {
  account := createRandomAccount(t)
  arg := UpdateAccountParams { 
    ID: account.ID,
    Balance: util.RandomInt(0, 11110),
  }
  updateAccount, err := testQueries.UpdateAccount(context.Background(), arg)
  require.NoError(t, err)
  require.NotEmpty(t, account)
  require.Equal(t, account.ID, updateAccount.ID)
  require.Equal(t, account.Owner, updateAccount.Owner)
  require.Equal(t, arg.Balance, updateAccount.Balance)
  require.Equal(t, account.Currency, updateAccount.Currency)
  require.WithinDuration(t, account.CreatedAt, updateAccount.CreatedAt, time.Second)
}

func TestDeleteAccounts(t *testing.T) {
  createAccounts := createRandomAccount(t)

  err := testQueries.DeleteAccount(context.Background(), createAccounts.ID)
  if err != nil { log.Fatal(err) }

  // for checking the account that we create
  checkAccounts, err := testQueries.GetAccount(context.Background(), createAccounts.ID)
  require.Error(t, err)
  require.Empty(t, checkAccounts)
}
