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
