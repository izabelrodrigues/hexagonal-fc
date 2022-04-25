package application_test

import (
	"testing"

	"github.com/izabelrodrigues/hexagonal-fc/application"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {

	product := application.Product{}
	product.Name = "Product 01"
	product.Status = application.DISABLED
	product.Price = 10

	err := product.Enable()
	require.Nil(t, err)

	product.Price = 0
	err = product.Enable()
	require.Equal(t, "The price must be greater than zero to enabled the product", err.Error())

}

func TestProduct_Disable(t *testing.T) {

	product := application.Product{}
	product.Name = "Product 01"
	product.Status = application.ENABLED
	product.Price = 0

	err := product.Disable()
	require.Nil(t, err)

	product.Price = 10
	err = product.Disable()
	require.Equal(t, "The price must be zero to disabled the product", err.Error())

}

func TestProduct_isValid(t *testing.T) {

	product := application.Product{}
	product.ID = uuid.NewV4().String()
	product.Name = "Product 01"
	product.Status = application.DISABLED
	product.Price = 10

	_, err := product.IsValid()
	require.Nil(t, err)

	product.Status = "INVALID"
	_, err = product.IsValid()
	require.Equal(t, "The status must be enabled or disabled", err.Error())

	product.Status = application.ENABLED
	_, err = product.IsValid()
	require.Nil(t, err)

	product.Price = -10
	_, err = product.IsValid()
	require.Equal(t, "The price must be greater or equal zero", err.Error())

}
