package service

import (
	"database/sql"

	"github.com/sinardyas/banking/domain"
	"github.com/sinardyas/banking/dto"
	"github.com/sinardyas/banking/helper"
)

type MemberService interface {
	GetAllMember(helper.PaginationRequest) ([]dto.MemberDto, *helper.AppError)
	GetById(string) (*dto.MemberDto, *helper.AppError)
	Create(dto.MemberDto) (*dto.MemberDto, *helper.AppError)
}

type DefaultMemberService struct {
	repo domain.IMember
}

func (d DefaultMemberService) GetAllMember(pagination helper.PaginationRequest) ([]dto.MemberDto, *helper.AppError) {
	members, err := d.repo.FindAll(pagination)
	if err != nil {
		return nil, err
	}
	response := make([]dto.MemberDto, 0)
	for _, c := range members {
		response = append(response, c.Convert())
	}
	return response, nil
}

func (d DefaultMemberService) GetById(id string) (*dto.MemberDto, *helper.AppError) {
	member, err := d.repo.ById(id)
	if err != nil {
		return nil, err
	}
	response := member.Convert()
	return &response, nil
}

func (d DefaultMemberService) Create(dto dto.MemberDto) (*dto.MemberDto, *helper.AppError) {
	_, appError, err := d.repo.ByColumn("nik_number", dto.NIKNumber)
	if appError != nil && err != sql.ErrNoRows {
		return nil, helper.UnexpectedError("Unexpected error when validating")
	}

	if appError == nil && err == nil {
		return nil, helper.ValidationError("NIK already exist!")
	}

	request := domain.Member{
		KTANumber:      dto.KTANumber,
		NIKNumber:      dto.NIKNumber,
		MemberName:     dto.MemberName,
		DateOfBirth:    dto.DateOfBirth,
		PlaceOfBirth:   dto.PlaceOfBirth,
		Address:        dto.Address,
		BusinessSector: dto.BusinessSector,
		PhoneNumber:    dto.PhoneNumber,
		Status:         dto.Status,
	}
	member, appError := d.repo.Create(request)
	if appError != nil {
		return nil, appError
	}
	response := member.Convert()
	return &response, nil
}

func NewMemberService(repository domain.IMember) DefaultMemberService {
	return DefaultMemberService{repo: repository}
}
