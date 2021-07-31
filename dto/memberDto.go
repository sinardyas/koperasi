package dto

type MemberDto struct {
	MemberId       string `json:"id"`
	KTANumber      string `json:"kta_number"`
	NIKNumber      string `json:"nik_number"`
	MemberName     string `json:"member_name"`
	DateOfBirth    string `json:"date_of_birth"`
	PlaceOfBirth   string `json:"place_of_birth"`
	Address        string `json:"address"`
	BusinessSector string `json:"business_sector"`
	PhoneNumber    string `json:"phone_number"`
	Status         string `json:"status"`
	CreatedAt      string `json:"created_at"`
	UpdatedAt      string `json:"updated_at"`
	DeletedAt      string `json:"deleted_at"`
}
