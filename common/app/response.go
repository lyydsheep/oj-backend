package app

import (
	"back/common/errcode"
	log "back/common/logger"
	"errors"
	"github.com/gin-gonic/gin"
)

// 统一响应
type response struct {
	c          *gin.Context
	Code       int         `json:"code"`
	Msg        string      `json:"message"`
	Data       any         `json:"data,omitempty"`
	RequestId  string      `json:"request_id"`
	Pagination *Pagination `json:"Pagination,omitempty"`
}

func NewResponse(c *gin.Context) *response {
	return &response{
		c: c,
	}
}

func (r *response) SetPagination(pagination *Pagination) *response {
	r.Pagination = pagination
	return r
}

func (r *response) Success(data any) {
	r.Code = errcode.Success.GetCode()
	r.Msg = errcode.Success.GetMsg()
	if _, ok := r.c.Get("traceId"); ok {
		r.RequestId = r.c.GetString("traceId")
	}
	r.Data = data
	r.c.JSON(errcode.Success.HttpStatusCode(), r)
}

func (r *response) SuccessOk() {
	r.Success("")
}

func (r *response) Error(err error) {
	newErr := errcode.ErrServer.Clone()
	if !errors.As(err, &newErr) {
		newErr = newErr.WithCause(err)
	}
	r.Code = newErr.GetCode()
	r.Msg = newErr.GetMsg()
	if _, ok := r.c.Get("traceId"); ok {
		r.RequestId = r.c.GetString("traceId")
	}
	log.New(r.c).Error("api_response_err", "err", err)
	r.c.JSON(newErr.HttpStatusCode(), r)
}
