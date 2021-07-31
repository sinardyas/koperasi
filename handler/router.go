package handler

import (
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

	router.HandleFunc("/member/{member_id}", mh.GetById).Methods("GET")
	router.HandleFunc("/member/list", mh.GetAllMember).Methods("POST")
	router.HandleFunc("/member", mh.Create).Methods("POST")
}
