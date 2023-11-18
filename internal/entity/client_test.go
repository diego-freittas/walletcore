package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("Diego Freitas", "tpddiego@gmail.com")

	if err != nil {
		assert.Nil(t, err)
		assert.NotNil(t, client)
		assert.Equal(t, "Diego Freitas", client.Name)
		assert.Equal(t, "tpddiego@gmail.com", client.Email)
	}
}

func TestCreateNewClientWhenArgumentsInvalid(t *testing.T) {
	client, err := NewClient("", "")

	assert.NotNil(t, err)
	assert.Nil(t, client)

}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("Diego Freitas", "tpddiego@gmail.com")
	err := client.Update("Diego", "tpddiego@gmail.com.br")
	assert.Nil(t, err)
	assert.Equal(t, "Diego", client.Name)
	assert.Equal(t, "tpddiego@gmail.com.br", client.Email)
}

func TestUpdateClientIWithInvalidArguments(t *testing.T) {
	client, _ := NewClient("Diego Freitas", "tpddiego@gmail.com")
	err := client.Update("", "tpddiego@gmail.com.br")
	assert.Error(t, err, "Name is required")

}
