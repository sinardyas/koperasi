package domain

import (
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/sinardyas/banking/helper"
)

type IMemberDbAdapter struct {
	client *sqlx.DB
}

func (m IMemberDbAdapter) FindAll(pagination helper.PaginationRequest) ([]Member, *helper.AppError) {
	var err error
	var page = (pagination.Page - 1) * pagination.Size
	var size = pagination.Size
	members := make([]Member, 0)

	findAllSql := `SELECT BIN_TO_UUID(id) id, member_name, kta_number, nik_number, date_of_birth, place_of_birth, address, phone_number, status, created_at, updated_at, deleted_at FROM api25o82ny8gixl4.member`
	whereSql := " WHERE " + pagination.Filter.Field + " LIKE " + "'%" + pagination.Filter.Value + "%' " + "AND deleted_at IS NULL"
	orderSql := " ORDER BY " + pagination.Order.Field + " " + pagination.Order.Direction + " "
	limitSql := " LIMIT " + strconv.Itoa(page) + ", " + strconv.Itoa(size)

	if pagination.Filter.Value != "" && pagination.Filter.Field != "" {
		findAllSql = findAllSql + whereSql
	}

	if pagination.Order.Direction != "" && pagination.Order.Field != "" {
		findAllSql = findAllSql + orderSql
	}

	findAllSql = findAllSql + limitSql
	fmt.Printf("SQL Query :: ", findAllSql)

	err = m.client.Select(&members, findAllSql)
	if err != nil {
		fmt.Println("Error when find all members :: " + err.Error())
		return nil, helper.UnexpectedError("Unexpected error occured from database")
	}

	return members, nil
}

func (m IMemberDbAdapter) ById(id string) (*Member, *helper.AppError) {
	var member Member
	byIdSql := `SELECT BIN_TO_UUID(id) id, member_name, kta_number, nik_number, date_of_birth, place_of_birth, address,
			phone_number, status, created_at, updated_at, deleted_at FROM api25o82ny8gixl4.member ` + "WHERE id = UUID_TO_BIN(?)" + "AND deleted_at IS NULL"

	err := m.client.Get(&member, byIdSql, id)
	if err != nil {
		fmt.Println("Error when get by id :: " + err.Error())
		return nil, helper.UnexpectedError("Unexpected error occured from database")
	}

	return &member, nil
}

func (m IMemberDbAdapter) ByColumn(columnName string, value interface{}) (*Member, *helper.AppError, error) {
	var additionalSql string

	if columnName == "id" {
		additionalSql = "UUID_TO_BIN(?)"
	} else {
		additionalSql = "?"
	}

	var member Member
	byIdSql := `SELECT BIN_TO_UUID(id) id, member_name, kta_number, nik_number, date_of_birth, place_of_birth, address,
			phone_number, status, created_at, updated_at, deleted_at FROM api25o82ny8gixl4.member WHERE ` + columnName + " = " + additionalSql

	err := m.client.Get(&member, byIdSql, value)
	if err != nil {
		fmt.Println("Error when get by column :: " + err.Error())
		return nil, helper.UnexpectedError("Unexpected error occured from database"), err
	}

	return &member, nil, nil
}

func (m IMemberDbAdapter) Create(d Member) (*Member, *helper.AppError) {
	sqlInsert := `INSERT INTO api25o82ny8gixl4.member (member_name, kta_number, nik_number, date_of_birth, place_of_birth, address, business_sector, phone_number, status)  
		values (?, ?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := m.client.Exec(sqlInsert, d.MemberName, d.KTANumber, d.NIKNumber, d.DateOfBirth, d.PlaceOfBirth, d.Address, d.BusinessSector, d.PhoneNumber, d.Status)
	if err != nil {
		fmt.Println("Error when creating new record :: " + err.Error())
		return nil, helper.UnexpectedError("Unexpected error occured from database")
	}

	id, err := result.LastInsertId()
	if err != nil {
		fmt.Println("Error when creating new record :: " + err.Error())
		return nil, helper.UnexpectedError("Unexpected error occured from database")
	}

	fmt.Println(id)
	d.MemberId = strconv.FormatInt(id, 10)
	return &d, nil
}

func (m IMemberDbAdapter) Update(d Member) *helper.AppError {
	fmt.Println("DATA :: ", d)
	sqlUpdate := "UPDATE api25o82ny8gixl4.member SET "
	sqlWhere := "WHERE id = UUID_TO_BIN(?)"

	if d.KTANumber != "" {
		sqlUpdate = sqlUpdate + "kta_number = '" + d.KTANumber + "', "
	}

	if d.NIKNumber != "" {
		sqlUpdate = sqlUpdate + "nik_number = '" + d.NIKNumber + "', "
	}

	if d.MemberName != "" {
		sqlUpdate = sqlUpdate + "member_name = '" + d.MemberName + "', "
	}

	if d.DateOfBirth != "" {
		sqlUpdate = sqlUpdate + "date_of_birth = '" + d.DateOfBirth + "', "
	}

	if d.PlaceOfBirth != "" {
		sqlUpdate = sqlUpdate + "place_of_birth = '" + d.PlaceOfBirth + "', "
	}

	if d.Address != "" {
		sqlUpdate = sqlUpdate + "address = '" + d.Address + "', "
	}

	if d.BusinessSector != "" {
		sqlUpdate = sqlUpdate + "business_sector = '" + d.BusinessSector + "', "
	}

	if d.PhoneNumber != "" {
		sqlUpdate = sqlUpdate + "phone_number = '" + d.PhoneNumber + "', "
	}

	if d.Status != "" {
		sqlUpdate = sqlUpdate + "status = '" + d.Status + "', "
	}

	sqlUpdate = sqlUpdate + "updated_at = '" + d.UpdatedAt.String + "' "
	sqlUpdate = sqlUpdate + sqlWhere
	fmt.Println("SQL Query :: ", sqlUpdate)

	_, err := m.client.Exec(sqlUpdate, d.MemberId)
	if err != nil {
		fmt.Println("Error when updating a record :: " + err.Error())
		return helper.UnexpectedError("Unexpected error occured from database")
	}

	return nil
}

func (m IMemberDbAdapter) SoftDelete(d Member) *helper.AppError {
	fmt.Println("DATA :: ", d)
	sqlUpdate := "UPDATE api25o82ny8gixl4.member SET "
	sqlWhere := "WHERE id = UUID_TO_BIN(?)"

	sqlUpdate = sqlUpdate + "updated_at = '" + d.UpdatedAt.String + "', "
	sqlUpdate = sqlUpdate + "deleted_at = '" + d.DeletedAt.String + "' "
	sqlUpdate = sqlUpdate + sqlWhere
	fmt.Println("SQL Query :: ", sqlUpdate)

	_, err := m.client.Exec(sqlUpdate, d.MemberId)
	if err != nil {
		fmt.Println("Error when deleting a record :: " + err.Error())
		return helper.UnexpectedError("Unexpected error occured from database")
	}

	return nil
}

func NewMemberRepository(db *sqlx.DB) IMemberDbAdapter {
	return IMemberDbAdapter{db}
}
