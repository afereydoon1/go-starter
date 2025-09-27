package utils

import (
	"testing"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestInitValidator(t *testing.T) {
	InitValidator()

	v, ok := binding.Validator.Engine().(*validator.Validate)
	assert.True(t, ok, "Validator should be initialized")
	//assert.Equal(t, "binding", v.GetTagName(), "Validator tag should be set to 'binding'")

	// Test validation with a sample struct
	type TestStruct struct {
		Name string `binding:"required"`
	}
	data := TestStruct{Name: ""}
	err := v.Struct(data)
	assert.Error(t, err, "Validator should return error for invalid struct")

	data = TestStruct{Name: "test"}
	err = v.Struct(data)
	assert.NoError(t, err, "Validator should not return error for valid struct")
}