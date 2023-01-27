package db

import (
	"context"
	"testing"
	"log"

	"github.com/stretchr/testify/require"
)

func createRandomEntries(t *testing.T) Entries {
	account := createRandomAccount(t)
	arg := CreateEntriesParams{
		AccountID: account.ID,
		Amount:    account.Balance,
	}
	entries, err := testQueries.CreateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, account.ID, entries.AccountID)
	require.Equal(t, account.Balance, entries.Amount)
	return entries
}

func TestCreateEntries(t *testing.T) {
	createRandomEntries(t)
}

func TestUpdateEntries(t *testing.T) {
	entries := createRandomEntries(t)
	arg := UpdateEntriesParams{
		ID:     entries.ID,
		Amount: entries.Amount,
	}
	updatedEntries, err := testQueries.UpdateEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Equal(t, entries.ID, updatedEntries.ID)
	require.Equal(t, entries.Amount, updatedEntries.Amount)
}

func TestGetEntries(t *testing.T) {
	entries := createRandomEntries(t)
	getEntries, err := testQueries.GetEntries(context.Background(), entries.ID)
	require.NoError(t, err)
	require.Equal(t, entries.ID, getEntries.ID)
}

func TestListEntries(t *testing.T) {
	arg := ListEntriesParams{
		Limit:  1,
		Offset: 10,
	}
	allEntries, err := testQueries.ListEntries(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, allEntries, len(allEntries))
}

func TestDeleteEntries(t *testing.T) {
  createEntries := createRandomEntries(t)

  err := testQueries.DeleteEntries(context.Background(), createEntries.ID)
  if err != nil { log.Fatal(err) }

  // for checking the account that we create
  checkEntries, err := testQueries.GetEntries(context.Background(), createEntries.ID)
  require.Error(t, err)
  require.Empty(t, checkEntries)
}
