package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {

	prod, err := NewProduct("Produto 1", 10.0)
	assert.Nil(t, err)
	assert.NotNil(t, prod)
	assert.NotEmpty(t, prod.ID)
	assert.Equal(t, "Produto 1", prod.Name)
	assert.Equal(t, 10.0, prod.Price)

}

func TestProductWhenNameIsRequired(t *testing.T) {

	prod, err := NewProduct("", 10.0)
	assert.Nil(t, prod)
	assert.Equal(t, ErrNameIsRequired, err)

}

func TestProductWhenPriceIsRequired(t *testing.T) {

	prod, err := NewProduct("Product", 0)
	assert.Nil(t, prod)
	assert.Equal(t, ErrPriceIsRequired, err)

}

func TestProductWhenPriceIsInvalid(t *testing.T) {

	prod, err := NewProduct("Product", -10)
	assert.Nil(t, prod)
	assert.Equal(t, ErrInvalidPrice, err)

}

func TestProductValidate(t *testing.T) {
	prod, err := NewProduct("Product 01", 10)
	assert.Nil(t, err)
	assert.NotNil(t, prod)
	assert.Nil(t, prod.Validate())
}
