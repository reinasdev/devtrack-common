package errors

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

type Code string

const (
	// Códigos de erro
	CodeInternalError Code = "internal_error"
	CodeBadRequest    Code = "bad_request"
	CodeNotFound      Code = "not_found"
	CodeUnauthorized  Code = "unauthorized"
	CodeForbidden     Code = "forbidden"
)

type Error struct {
	Code       Code   `json:"code"`
	StatusCode int    `json:"status_code"`
	Msg        string `json:"message"`
	Trace      string `json:"trace"`
}

// Error implementa a interface error
func (e *Error) Error() string {
	return fmt.Sprintf("%s [%s]", e.Msg, e.Code)
}

// NewErr cria um novo erro customizado
func NewErr(code Code, msg string) error {
	return &Error{
		Code:       code,
		StatusCode: mapCodeToStatusCode(code),
		Msg:        msg,
	}
}

// Err garante que o erro seja do tipo *Error
func Err(err error) error {
	return fromErr(err)
}

// fromErr mapeia erros padrão de bibliotecas para nosso tipo Error
func fromErr(err error) *Error {
	if err == nil {
		return nil
	}

	switch err := err.(type) {
	case *Error:
		return &Error{
			Code:       err.Code,
			StatusCode: err.StatusCode,
			Msg:        err.Msg,
			Trace:      extractStackTrace(err),
		}
	case error:
		switch err {
		case sql.ErrNoRows:
			return &Error{
				Code:       CodeNotFound,
				StatusCode: http.StatusNotFound,
				Msg:        "Recurso não encontrado",
				Trace:      extractStackTrace(err),
			}
		}
	}

	return &Error{
		Code:       CodeInternalError,
		StatusCode: http.StatusInternalServerError,
		Msg:        "Erro desconhecido",
		Trace:      extractStackTrace(err),
	}
}

// Wrap adiciona uma mensagem ao erro com stack trace
func Wrap(err error, message string) error {
	return errors.Wrap(Err(err), message)
}

// Handling lida com o erro e envia resposta estruturada via Gin
func Handling(ctx *gin.Context, err error) {
	e := Err(err).(*Error)
	e.Msg = err.Error()
	e.Trace = extractStackTrace(err)

	if reqIDVal, hasRID := ctx.Get("RID"); hasRID {
		if reqID, ok := reqIDVal.(string); ok && len(reqID) > 0 {
			e.Msg = fmt.Sprintf("%s [%s]", e.Msg, reqID[:6])
		}
	}

	ctx.JSON(e.StatusCode, e)
	ctx.Set("error", err)
	ctx.Abort()
}

// mapCodeToStatusCode traduz código para status HTTP
func mapCodeToStatusCode(code Code) int {
	switch code {
	case CodeBadRequest:
		return http.StatusBadRequest
	case CodeNotFound:
		return http.StatusNotFound
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	default:
		return http.StatusInternalServerError
	}
}

// extractStackTrace retorna stack trace formatado
func extractStackTrace(err error) string {
	if err == nil {
		return ""
	}

	return fmt.Sprintf("%+v", errors.WithStack(err))
}
