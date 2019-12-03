package service

import (
	"encoding/json"
	"github.com/hurtuh/indriver/domain"
	"log"
	"net/http"
	"time"
)

type Logic struct {
	repo domain.Repository
}

func NewLogic(repository domain.Repository) *Logic {
	return &Logic{repo:repository}
}

func (l *Logic) NewCandidate(r *http.Request) (code int64, msg error, result interface{}) {
	candidate := new(domain.Candidate)
	var err error

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&candidate); err != nil {
		msg = domain.ErrorWithUnmarshal
		code = domain.CodeBadRequest
		return
	}

	if err = l.repo.NewCandidate(candidate); err != nil {
		log.Printf("Error with save new candidate, %v", candidate)
		msg = domain.ErrorWithDatabase
		code = domain.CodeDatabaseError
		return
	}

	code = domain.CodeSuccess
	result = struct {
		result string
	}{result:"success"}

	return
}

func (l *Logic) EditCandidate(r *http.Request) (code int64, msg error, result interface{}) {
	editCandReq := new(domain.EditCandidateReq)
	var err error

	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(&editCandReq); err != nil {
		msg = domain.ErrorWithUnmarshal
		code = domain.CodeBadRequest
		return
	}

	if editCandReq.Description != "" {
		err = l.repo.EditDescription(editCandReq.CandidateID, editCandReq.Description)
		if err != nil {
			log.Printf("Error with edit description candidate, %v", editCandReq)
			msg = domain.ErrorWithDatabase
			code = domain.CodeDatabaseError
			return
		}
	}

	zeroTime := time.Time{}

	if editCandReq.NewTime != zeroTime {
		err = l.repo.EditInterview(editCandReq.CandidateID, editCandReq.NewTime)
		if err != nil {
			log.Printf("Error with edit interview candidate, %v", editCandReq)
			msg = domain.ErrorWithDatabase
			code = domain.CodeDatabaseError
			return
		}
	}

	code = domain.CodeSuccess
	result = struct {
		result string
	}{result:"success"}

	return
}

func (l *Logic) DeleteCandidate(r *http.Request) (code int64, msg error, result interface{}) {
	delCandReq := new(domain.DeleteCandidateReq)
	var err error

	defer r.Body.Close()

	if err = json.NewDecoder(r.Body).Decode(&delCandReq); err != nil {
		msg = domain.ErrorWithUnmarshal
		code = domain.CodeBadRequest
		return
	}

	if err = l.repo.DeleteCandidate(delCandReq.CandidateID); err != nil {
		log.Printf("Error with delete candidate, %v", delCandReq)
		msg = domain.ErrorWithDatabase
		code = domain.CodeDatabaseError
		return
	}

	code = domain.CodeSuccess
	result = struct {
		result string
	}{result:"success"}

	return
}

func (l *Logic) GetCandidate(r *http.Request) (code int64, msg error, result interface{}) {
	getCandidateReq := new(domain.GetCandidateReq)
	var err error

	defer r.Body.Close()

	if err = json.NewDecoder(r.Body).Decode(&getCandidateReq); err != nil {
		msg = domain.ErrorWithUnmarshal
		code = domain.CodeBadRequest
		return
	}

	if getCandidateReq.CandidateID != 0 {
		result, err = l.repo.SelectByID(getCandidateReq.CandidateID)
		if err != nil {
			log.Printf("Error with get candidate, %v", getCandidateReq)
			msg = domain.ErrorWithDatabase
			code = domain.CodeDatabaseError
		}
		return
	}

	result, err = l.repo.SelectAll()
	if err != nil {
		log.Printf("Error with get candidate, %v", getCandidateReq)
		msg = domain.ErrorWithDatabase
		code = domain.CodeDatabaseError
	}

	return
}
