package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	client, _ := NewClient("Jhon Doe", "j@j.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client, account.Client)
	assert.Equal(t, 0.0, account.Balance)
	assert.NotEmpty(t, account.ID)
	assert.NotZero(t, account.CreatedAt)
	assert.NotZero(t, account.UpdatedAt)
}

func TestCreateAccountWithNilClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("Jhon Doe", "j@j.com")
	account := NewAccount(client)
	assert.NotNil(t, account)

	account.Credit(100)
	assert.Equal(t, 100.0, account.Balance)
	assert.NotZero(t, account.UpdatedAt)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("Jhon Doe", "j@j.com")
	account := NewAccount(client)
	assert.NotNil(t, account)

	account.Credit(100)
	assert.Equal(t, 100.0, account.Balance)

	err := account.Debit(50)
	assert.NoError(t, err)
	assert.Equal(t, 50.0, account.Balance)

	err = account.Debit(100)
	assert.Error(t, err)
	assert.Equal(t, "insufficient funds", err.Error())
}
