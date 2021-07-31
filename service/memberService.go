package service

import (
	"database/sql"
	"time"

	"github.com/sinardyas/banking/domain"
	"github.com/sinardyas/banking/dto"
	"github.com/sinardyas/banking/helper"
)

type MemberService interface {
	GetAllMember(helper.PaginationRequest) ([]dto.MemberDto, *helper.AppError)
	GetById(string) (*dto.MemberDto, *helper.AppError)
	Create(dto.MemberDto) (*dto.MemberDto, *helper.AppError)
	Update(dto.MemberDto) (bool, *helper.AppError)
	Delete(dto.MemberDto) (bool, *helper.AppError)
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
	_, validationErr, err := d.repo.ByColumn("nik_number", dto.NIKNumber)
	if validationErr != nil && err != sql.ErrNoRows {
		return nil, helper.UnexpectedError("Unexpected error when validating")
	}

	if validationErr == nil && err == nil {
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

func (d DefaultMemberService) Update(dto dto.MemberDto) (bool, *helper.AppError) {
	_, validationErr, err := d.repo.ByColumn("id", dto.MemberId)
	if err != nil && err != sql.ErrNoRows {
		return false, helper.ValidationError("Unexpected error when validating")
	}

	if validationErr != nil && err != nil {
		return false, helper.ValidationError("Member ID not found")
	}

	currentTime := time.Now()

	request := domain.Member{
		MemberId:       dto.MemberId,
		KTANumber:      dto.KTANumber,
		NIKNumber:      dto.NIKNumber,
		MemberName:     dto.MemberName,
		DateOfBirth:    dto.DateOfBirth,
		PlaceOfBirth:   dto.PlaceOfBirth,
		Address:        dto.Address,
		BusinessSector: dto.BusinessSector,
		PhoneNumber:    dto.PhoneNumber,
		Status:         dto.Status,
		UpdatedAt:      sql.NullString{String: currentTime.Format("2006-01-02 15:04:05"), Valid: true},
	}

	appError := d.repo.Update(request)
	if appError != nil {
		return false, appError
	}
	return true, nil
}

func (d DefaultMemberService) Delete(dto dto.MemberDto) (bool, *helper.AppError) {
	_, validationErr, err := d.repo.ByColumn("id", dto.MemberId)
	if err != nil && err != sql.ErrNoRows {
		return false, helper.ValidationError("Unexpected error when validating")
	}

	if validationErr != nil && err != nil {
		return false, helper.ValidationError("Member ID not found")
	}

	currentTime := time.Now()

	request := domain.Member{
		UpdatedAt: sql.NullString{String: currentTime.Format("2006-01-02 15:04:05"), Valid: true},
		DeletedAt: sql.NullString{String: currentTime.Format("2006-01-02 15:04:05"), Valid: true},
	}

	appError := d.repo.SoftDelete(request)
	if appError != nil {
		return false, appError
	}
	return true, nil
}

func NewMemberService(repository domain.IMember) DefaultMemberService {
	return DefaultMemberService{repo: repository}
}
