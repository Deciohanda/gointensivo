package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_If_I_Get_An_Error_If_ID_Is_Blank(t *testing.T) {
	order := Order{}
	// err := order.Validate()
	// if err == nil {
	// 	t.Error("Expected error, got nil")
	// }
	assert.Error(t, order.Validate(), "invalid id")
}

func Test_If_I_Get_An_Error_If_Price_Is_Blank(t *testing.T) {
	order := Order{ID: "123"}
	assert.Error(t, order.Validate(), "invalid price")
}

func Test_If_I_Get_An_Error_If_Tax_Is_Blank(t *testing.T) {
	order := Order{ID: "123", Price: 10.0}
	assert.Error(t, order.Validate(), "invalid tax")
}

func Test_With_All_Valid_Params(t *testing.T) {
	order := Order{ID: "123", Price: 10.0, Tax: 1.0}
	assert.NoError(t, order.Validate())
	assert.Equal(t, 10.0, order.Price)
	assert.Equal(t, 1.0, order.Tax)
	order.CalculateFinalPrice()
	assert.Equal(t, 11.0, order.FinalPrice)
}
