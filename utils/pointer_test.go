package utils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/reinasdev/devtrack-common/utils"
)

func TestGet(t *testing.T) {
	t.Run("deve retornar ponteiro para o valor", func(t *testing.T) {
		val := 42
		ptr := utils.Get(val)

		assert.NotNil(t, ptr)
		assert.Equal(t, &val, ptr)
		assert.Equal(t, val, *ptr)
	})
}

func TestCoalesce(t *testing.T) {
	t.Run("deve retornar o primeiro ponteiro não-nulo", func(t *testing.T) {
		a, b := 1, 2
		result := utils.Coalesce(&a, &b)
		assert.NotNil(t, result)
		assert.Equal(t, a, *result)
		assert.Equal(t, &a, result)
	})

	t.Run("deve retornar nil se todos os valores forem nil", func(t *testing.T) {
		assert.Nil(t, utils.Coalesce[int](nil, nil, nil))
	})

	t.Run("deve retornar o primeiro valor quando todos forem não-nulos", func(t *testing.T) {
		x, y, z := 10, 20, 30
		result := utils.Coalesce(&x, &y, &z)
		assert.NotNil(t, result)
		assert.Equal(t, x, *result)
		assert.Equal(t, &x, result)
	})

	t.Run("deve retornar nil quando não houver entrada", func(t *testing.T) {
		assert.Nil(t, utils.Coalesce[int]())
	})
}
