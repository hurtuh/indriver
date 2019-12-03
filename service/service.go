package service

import (
	"database/sql"
	"github.com/hurtuh/indriver/domain"
	"log"
	"time"
)

type Logic struct {
	repo domain.Repository
}

func NewLogic(repository domain.Repository) *Logic {
	return &Logic{repo: repository}
}

func (l *Logic) NewCandidate(candidate *domain.Candidate) (code int64, msg error, result interface{}) {
	var err error

	if err = l.repo.NewCandidate(candidate); err != nil {

		if err == sql.ErrNoRows {
			msg = domain.ErrorNoFound
			code = domain.CodeNotFound
			return
		}

		log.Printf("Error with save new candidate, %v", candidate)
		msg = domain.ErrorWithDatabase
		code = domain.CodeDatabaseError
		return
	}

	code = domain.CodeSuccess
	result = struct {
		Result string `json:"result"`
	}{Result: "success"}

	return
}

func (l *Logic) EditCandidate(req *domain.EditCandidateReq) (code int64, msg error, result interface{}) {
	var err error

	if req.Description != "" {
		err = l.repo.EditDescription(req.CandidateID, req.Description)
		if err != nil {

			if err == sql.ErrNoRows {
				msg = domain.ErrorNoFound
				code = domain.CodeNotFound
				return
			}

			log.Printf("Error with edit description candidate, %v, err: %s", req, err)
			msg = domain.ErrorWithDatabase
			code = domain.CodeDatabaseError
			return
		}
	}

	zeroTime := time.Time{}

	if req.NewTime != zeroTime {
		err = l.repo.EditInterview(req.CandidateID, req.NewTime)
		if err != nil {

			if err == sql.ErrNoRows {
				msg = domain.ErrorNoFound
				code = domain.CodeNotFound
				return
			}

			log.Printf("Error with edit interview candidate, %v, err: %s", req, err)
			msg = domain.ErrorWithDatabase
			code = domain.CodeDatabaseError
			return
		}
	}

	code = domain.CodeSuccess
	result = struct {
		Result string `json:"result"`
	}{Result: "success"}

	return
}

func (l *Logic) DeleteCandidate(req *domain.DeleteCandidateReq) (code int64, msg error, result interface{}) {
	var err error
	if err = l.repo.DeleteCandidate(req.CandidateID); err != nil {

		if err == sql.ErrNoRows {
			msg = domain.ErrorNoFound
			code = domain.CodeNotFound
			return
		}

		log.Printf("Error with delete candidate, %v, err: %s", req, err)
		msg = domain.ErrorWithDatabase
		code = domain.CodeDatabaseError
		return
	}

	code = domain.CodeSuccess
	result = struct {
		Result string `json:"result"`
	}{Result: "success"}

	return
}

func (l *Logic) GetCandidate(req *domain.GetCandidateReq) (code int64, msg error, result interface{}) {
	var err error

	if req.CandidateID != 0 {
		result, err = l.repo.SelectByID(req.CandidateID)
		if err != nil {

			if err == sql.ErrNoRows {
				msg = domain.ErrorNoFound
				code = domain.CodeNotFound
				return
			}

			log.Printf("Error with get candidate, %v, err: %s", req, err)
			msg = domain.ErrorWithDatabase
			code = domain.CodeDatabaseError
		}
		return
	}

	result, err = l.repo.SelectAll()
	if err != nil {

		if err == sql.ErrNoRows {
			msg = domain.ErrorNoFound
			code = domain.CodeNotFound
			return
		}

		log.Printf("Error with get candidate, %v, err: %s", req, err)
		msg = domain.ErrorWithDatabase
		code = domain.CodeDatabaseError
	}

	return
}
