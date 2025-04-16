package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

// Logger é a instância global do logger
var Logger *zap.Logger

// InitLogger inicializa o logger com a configuração fornecida ou padrão
func InitLogger(config zap.Config) {
	logger, err := config.Build()
	if err != nil {
		fmt.Printf("Falha ao inicializar o logger: %v\n", err)
		os.Exit(1)
	}

	Logger = logger
}

// Sync faz o flush dos logs antes do encerramento do programa
func Sync() {
	if err := Logger.Sync(); err != nil {
		fmt.Printf("Falha ao sincronizar o logger: %v\n", err)
	}
}

// LogInfo registra uma mensagem de log com nível INFO
func LogInfo(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

// LogWarn registra uma mensagem de log com nível WARN
func LogWarn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

// LogError registra uma mensagem de log com nível ERROR
func LogError(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

// LogDebug registra uma mensagem de log com nível DEBUG
func LogDebug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

// LogFatal registra uma mensagem de log com nível FATAL e termina a execução
func LogFatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

// WithFields permite adicionar campos estruturados a um log
func WithFields(fields map[string]interface{}) []zap.Field {
	var zapFields []zap.Field
	for key, value := range fields {
		zapFields = append(zapFields, zap.Any(key, value))
	}
	return zapFields
}
