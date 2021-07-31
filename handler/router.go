package handler

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/sinardyas/banking/controller"
	"github.com/sinardyas/banking/domain"
	"github.com/sinardyas/banking/service"
)

type Route struct{}

func (*Route) MemberRouterHandling(router *mux.Router, db *sqlx.DB) {
	client := openDbConnection()
	memberRepository := domain.NewMemberRepository(client)
	mh := controller.MemberHandler{Service: service.NewMemberService(memberRepository)}

	router.HandleFunc("/member", mh.Create).Methods("POST")
	router.HandleFunc("/member/list", mh.GetAllMember).Methods("POST")
	router.HandleFunc("/member/delete", mh.Delete).Methods("POST")
	router.HandleFunc("/member/{member_id}", mh.GetById).Methods("GET")
	router.HandleFunc("/member/{member_id}", mh.Update).Methods("PUT")

	router.HandleFunc("/ping", func(rw http.ResponseWriter, r *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}).Methods("GET")
}
