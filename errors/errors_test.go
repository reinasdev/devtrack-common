package errors_test

import (
	"database/sql"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	common "github.com/reinasdev/devtrack-common/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewErr(t *testing.T) {
	t.Run("deve criar erro com código e status correspondente", func(t *testing.T) {
		err := common.NewErr(common.CodeBadRequest, "erro de validação")
		customErr := err.(*common.Error)

		assert.Equal(t, common.CodeBadRequest, customErr.Code)
		assert.Equal(t, http.StatusBadRequest, customErr.StatusCode)
		assert.Equal(t, "erro de validação", customErr.Msg)
	})
}

func TestErr(t *testing.T) {
	t.Run("deve retornar erro do tipo *Error se já for *Error", func(t *testing.T) {
		original := common.NewErr(common.CodeNotFound, "não achei")
		converted := common.Err(original)

		assert.IsType(t, &common.Error{}, converted)
		assert.Equal(t, "não achei", converted.(*common.Error).Msg)
	})

	t.Run("deve converter sql.ErrNoRows para erro not found", func(t *testing.T) {
		err := common.Err(sql.ErrNoRows)
		rawErr, ok := err.(*common.Error)

		assert.True(t, ok)
		assert.Equal(t, common.CodeNotFound, rawErr.Code)
		assert.Equal(t, http.StatusNotFound, rawErr.StatusCode)
		assert.Contains(t, rawErr.Trace, "sql: no rows in result set")
	})

	t.Run("deve retornar erro genérico para erro desconhecido", func(t *testing.T) {
		err := common.Err(errors.New("erro qualquer")).(*common.Error)

		assert.Equal(t, common.CodeInternalError, err.Code)
		assert.Equal(t, http.StatusInternalServerError, err.StatusCode)
		assert.Contains(t, err.Trace, "erro qualquer")
	})
}

func TestWrap(t *testing.T) {
	t.Run("deve adicionar stack trace ao erro", func(t *testing.T) {
		base := common.NewErr(common.CodeInternalError, "erro base")
		wrapped := common.Wrap(base, "contexto adicional")

		assert.Error(t, wrapped)
		assert.Contains(t, wrapped.Error(), "contexto adicional")
		assert.Contains(t, wrapped.Error(), "erro base")
	})
}

func TestHandling(t *testing.T) {
	t.Run("deve retornar JSON estruturado com erro e status", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		router := gin.New()

		router.GET("/erro", func(c *gin.Context) {
			err := common.NewErr(common.CodeBadRequest, "erro simulado")
			common.Handling(c, err)
		})

		req := httptest.NewRequest("GET", "/erro", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), `"code":"bad_request"`)
		assert.Contains(t, rec.Body.String(), `"message":"[bad_request] erro simulado"`)
	})

	t.Run("deve incluir RID na mensagem se presente no contexto", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		router := gin.New()

		router.GET("/erro", func(c *gin.Context) {
			c.Set("RID", "abc123456789")
			err := common.NewErr(common.CodeInternalError, "problema sério")
			common.Handling(c, err)
		})

		req := httptest.NewRequest("GET", "/erro", nil)
		rec := httptest.NewRecorder()
		router.ServeHTTP(rec, req)

		assert.Contains(t, rec.Body.String(), `"message":"[internal_error] problema sério [abc123]"`)
	})
}
