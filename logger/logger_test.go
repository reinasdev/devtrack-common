package logger_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"github.com/reinasdev/devtrack-common/logger"
)

func TestInitLogger(t *testing.T) {
	t.Run("deve inicializar o logger sem erros", func(t *testing.T) {
		cfg := zap.NewDevelopmentConfig()
		logger.InitLogger(cfg)

		assert.NotNil(t, logger.Logger)
	})
}

func TestWithFields(t *testing.T) {
	t.Run("deve converter mapa em campos zap", func(t *testing.T) {
		fields := map[string]interface{}{
			"user_id": 123,
			"active":  true,
		}

		zapFields := logger.WithFields(fields)

		assert.Len(t, zapFields, 2)
	})
}

func TestLogLevels(t *testing.T) {
	t.Run("deve registrar logs de todos os níveis sem erro", func(t *testing.T) {
		mockLogger := zaptest.NewLogger(t)
		logger.Logger = mockLogger // substitui o logger global por um mock

		logger.LogDebug("debug test")
		logger.LogInfo("info test")
		logger.LogWarn("warn test")
		logger.LogError("error test")

		// Não testamos LogFatal porque ele chama os.Exit(1)
	})
}

func TestSync(t *testing.T) {
	t.Run("deve chamar Sync sem erro", func(t *testing.T) {
		mockLogger := zaptest.NewLogger(t)
		logger.Logger = mockLogger

		// Apenas garantindo que não panica
		logger.Sync()
	})
}
