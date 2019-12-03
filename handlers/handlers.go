package handlers

import (
	"encoding/json"
	"github.com/hurtuh/indriver/domain"
	"net/http"
)

type Handlers struct {
	service domain.Service
}

func NewHandlers(service domain.Service) *Handlers {
	return &Handlers{service: service}
}

type ResponseMessage struct {
	Code    int64       `json:"code"`
	Message string      `json:"message,omitempty"`
	Result  interface{} `json:"data,omitempty"`
}

func ApiResponse(w http.ResponseWriter, code int64, msg error, response interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := ResponseMessage{Code: code, Result: response}
	if nil != msg {
		resp.Message = msg.Error()
	}
	json.NewEncoder(w).Encode(resp)
}

func (h *Handlers) AddRecord(w http.ResponseWriter, r *http.Request) {
	code, msg, result := h.service.NewCandidate(r)
	ApiResponse(w, code, msg, result)
}

func (h *Handlers) DelRecord(w http.ResponseWriter, r *http.Request) {
	code, msg, result := h.service.DeleteCandidate(r)
	ApiResponse(w, code, msg, result)
}

func (h *Handlers) GetRecord(w http.ResponseWriter, r *http.Request) {
	code, msg, result := h.service.GetCandidate(r)
	ApiResponse(w, code, msg, result)
}

func (h *Handlers) EditRecord(w http.ResponseWriter, r *http.Request) {
	code, msg, result := h.service.EditCandidate(r)
	ApiResponse(w, code, msg, result)
}
