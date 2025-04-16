package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/reinasdev/devtrack-common/config"
)

func TestLoadEnv(t *testing.T) {
	t.Run("deve carregar variáveis de ambiente a partir do arquivo .env", func(t *testing.T) {
		// cria um arquivo temporário .env para o teste
		envContent := []byte("TESTE_VAR=valor_teste\n")
		tmpFile := ".env.test"
		err := os.WriteFile(tmpFile, envContent, 0644)
		assert.NoError(t, err)
		defer os.Remove(tmpFile)

		err = config.LoadEnv(tmpFile)
		assert.NoError(t, err)
		assert.Equal(t, "valor_teste", os.Getenv("TESTE_VAR"))
	})

	t.Run("deve retornar erro se o arquivo não existir", func(t *testing.T) {
		err := config.LoadEnv("arquivo_inexistente.env")
		assert.Error(t, err)
	})
}

func TestGet(t *testing.T) {
	t.Run("deve retornar valor da variável se estiver definida", func(t *testing.T) {
		os.Setenv("VAR_TESTE", "valor_definido")
		defer os.Unsetenv("VAR_TESTE")

		valor := config.Get("VAR_TESTE", "valor_padrao")
		assert.Equal(t, "valor_definido", valor)
	})

	t.Run("deve retornar valor padrão se variável não estiver definida", func(t *testing.T) {
		os.Unsetenv("VAR_INEXISTENTE")

		valor := config.Get("VAR_INEXISTENTE", "valor_padrao")
		assert.Equal(t, "valor_padrao", valor)
	})
}
