package domain

import (
	"database/sql"

	"github.com/sinardyas/banking/dto"
	"github.com/sinardyas/banking/helper"
)

type Member struct {
	MemberId       string         `db:"id"`
	KTANumber      string         `db:"kta_number"`
	NIKNumber      string         `db:"nik_number"`
	MemberName     string         `db:"member_name"`
	DateOfBirth    string         `db:"date_of_birth"`
	PlaceOfBirth   string         `db:"place_of_birth"`
	Address        string         `db:"address"`
	BusinessSector string         `db:"business_sector"`
	PhoneNumber    string         `db:"phone_number"`
	Status         string         `db:"status"`
	CreatedAt      string         `db:"created_at"`
	UpdatedAt      sql.NullString `db:"updated_at"`
	DeletedAt      sql.NullString `db:"deleted_at"`
}

type IMember interface {
	FindAll(helper.PaginationRequest) ([]Member, *helper.AppError)
	ById(id string) (*Member, *helper.AppError)
	ByColumn(string, interface{}) (*Member, *helper.AppError, error)
	Create(Member) (*Member, *helper.AppError)
	Update(Member) *helper.AppError
	SoftDelete(Member) *helper.AppError
}

func (m Member) Convert() dto.MemberDto {
	return dto.MemberDto{
		MemberId:       m.MemberId,
		KTANumber:      m.KTANumber,
		NIKNumber:      m.NIKNumber,
		MemberName:     m.MemberName,
		DateOfBirth:    m.DateOfBirth,
		PlaceOfBirth:   m.PlaceOfBirth,
		Address:        m.Address,
		BusinessSector: m.BusinessSector,
		PhoneNumber:    m.PhoneNumber,
		Status:         m.Status,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt.String,
		DeletedAt:      m.DeletedAt.String,
	}
}
