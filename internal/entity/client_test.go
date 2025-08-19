package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "j@j.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("John Doe Update", "j@j.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe Update", client.Name)
	assert.Equal(t, "j@j.com", client.Email)
}

func TestUpdateClientWithInvalidArgs(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	err := client.Update("", "")
	assert.NotNil(t, err, "name is required")
}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "j@j.com")
	account := NewAccount(client)
	assert.NotNil(t, account)

	err := client.AddAccount(account)
	assert.Nil(t, err)
	assert.Contains(t, client.Accounts, account)
}

func TestAddAccountToClientWithMismatchedClient(t *testing.T) {
	client1, _ := NewClient("John Doe", "j@j.com")
	client2, _ := NewClient("Jane Doe", "jane@j.com")
	account := NewAccount(client1)
	assert.NotNil(t, account)

	err := client2.AddAccount(account)
	assert.NotNil(t, err)
	assert.Equal(t, "account client does not match", err.Error())
}
