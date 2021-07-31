package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sinardyas/banking/dto"
	"github.com/sinardyas/banking/helper"
	"github.com/sinardyas/banking/service"
)

type MemberHandler struct {
	Service service.MemberService
}

var response helper.Response

func (mh *MemberHandler) GetAllMember(w http.ResponseWriter, r *http.Request) {
	var requstBody helper.PaginationRequest

	err := json.NewDecoder(r.Body).Decode(&requstBody)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, helper.Response{
			Status: "ERROR",
			Error:  err.Error(),
		})
		return
	}

	data, appError := mh.Service.GetAllMember(requstBody)
	if appError != nil {
		response.Send(w, appError.Code, helper.Response{
			Status: "ERROR",
			Error:  appError.AsMessage(),
		})
		return
	}

	response.Send(w, http.StatusOK, helper.Response{
		Status: "SUCCESS",
		Data: helper.PaginationResponse{
			Page: requstBody.Page,
			Size: requstBody.Size,
			Data: data,
		},
	})
}

func (mh *MemberHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	member_id := vars["member_id"]

	data, appError := mh.Service.GetById(member_id)
	if appError != nil {
		response.Send(w, appError.Code, helper.Response{
			Status: "ERROR",
			Error:  appError.AsMessage(),
		})
		return
	}

	response.Send(w, http.StatusOK, helper.Response{
		Status: "SUCCESS",
		Data:   data,
	})
}

func (mh *MemberHandler) Create(w http.ResponseWriter, r *http.Request) {
	var requstBody dto.MemberDto

	err := json.NewDecoder(r.Body).Decode(&requstBody)
	if err != nil {
		response.Send(w, http.StatusInternalServerError, helper.Response{
			Status: "ERROR",
			Error:  err.Error(),
		})
		return
	}

	data, appError := mh.Service.Create(requstBody)
	if appError != nil {
		response.Send(w, appError.Code, helper.Response{
			Status: "ERROR",
			Error:  appError.AsMessage(),
		})
		return
	}

	response.Send(w, http.StatusOK, helper.Response{
		Status: "SUCCESS",
		Data:   data,
	})
}
