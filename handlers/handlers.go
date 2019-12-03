package handlers

import (
	"encoding/json"
	"github.com/hurtuh/indriver/domain"
	"net/http"
	"strconv"
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
	candidate := new(domain.Candidate)
	var err, msg error
	var code int64
	var result interface{}

	defer r.Body.Close()

	if err = json.NewDecoder(r.Body).Decode(&candidate); err != nil {
		msg = domain.ErrorWithDecoding
		code = domain.CodeBadRequest
		ApiResponse(w, code, msg, result)
		return
	}

	code, msg, result = h.service.NewCandidate(candidate)
	ApiResponse(w, code, msg, result)
}

func (h *Handlers) DelRecord(w http.ResponseWriter, r *http.Request) {
	req := new(domain.DeleteCandidateReq)
	var err, msg error
	var code int64
	var result interface{}

	defer r.Body.Close()

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		msg = domain.ErrorWithDecoding
		code = domain.CodeBadRequest
		ApiResponse(w, code, msg, result)
		return
	}

	code, msg, result = h.service.DeleteCandidate(req)
	ApiResponse(w, code, msg, result)
}

func (h *Handlers) GetRecord(w http.ResponseWriter, r *http.Request) {
	req := new(domain.GetCandidateReq)
	var err, msg error
	var code int64
	var result interface{}

	id := r.FormValue("id")
	if id != "" {
		req.CandidateID, err = strconv.ParseUint(id, 10, 64)
		if err != nil {
			msg = domain.ErrorWithDecoding
			code = domain.CodeBadRequest
			ApiResponse(w, code, msg, result)
			return
		}
	}

	code, msg, result = h.service.GetCandidate(req)
	ApiResponse(w, code, msg, result)
}

func (h *Handlers) EditRecord(w http.ResponseWriter, r *http.Request) {
	req := new(domain.EditCandidateReq)
	var err, msg error
	var code int64
	var result interface{}

	defer r.Body.Close()

	if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
		msg = domain.ErrorWithDecoding
		code = domain.CodeBadRequest
		return
	}

	code, msg, result = h.service.EditCandidate(req)
	ApiResponse(w, code, msg, result)
}
