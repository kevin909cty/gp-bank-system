package db

import (
	"context"
	"github.com/stretchr/testify/require"
	"simplebank/db/util"
	"testing"
	"time"
)

func createRandomRecord(t *testing.T, account Account) Record {
	arg := CreateRecordParams{
		AccountID: account.ID,
		Amount:    util.RandomBalance(),
	}

	record, err := testQueries.CreateRecord(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, record)

	require.Equal(t, arg.AccountID, record.AccountID)
	require.Equal(t, arg.Amount, record.Amount)

	require.NotZero(t, record.AccountID)
	require.NotZero(t, record.CreatedAt)

	return record
}

func TestGetRecord(t *testing.T) {
	account := createRandomAccount(t)
	record1 := createRandomRecord(t, account)
	record2, err := testQueries.GetRecord(context.Background(), record1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, record2)

	require.Equal(t, record1.ID, record2.ID)
	require.Equal(t, record1.AccountID, record2.AccountID)
	require.Equal(t, record1.Amount, record2.Amount)
	require.WithinDuration(t, record1.CreatedAt, record2.CreatedAt, time.Second)
}

func TestCreateRecord(t *testing.T) {
	createRandomRecord(t, createRandomAccount(t))
}

func TestUpdateRecord(t *testing.T) {
	// TODO - BK2534: develop unit test method for update records operation
}

func TestDeleteRecord(t *testing.T) {
	// TODO - BK2535: develop unit test method for delete records operation
}

func TestListRecords(t *testing.T) {
	// TODO - BK2536: develop unit test method for get records operation
}
