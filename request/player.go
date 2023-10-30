package request

type CreatePlayer struct {
	CompanyID string `json:"companyId" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:""`
	NoHp      string `json:"noHp" validate:""`
	Address   string `json:"address" validate:""`
	IsActive  bool   `json:"isActive" validate:""`
}

type UpdatePlayer struct {
	CompanyID string `json:"companyId" validate:"required"`
	Name      string `json:"name" validate:"required"`
	Email     string `json:"email" validate:""`
	NoHp      string `json:"noHp" validate:""`
	Address   string `json:"address" validate:""`
	IsActive  bool   `json:"isActive" validate:""`
}

type PagePlayer struct {
	Paging
	CompanyID string `json:"companyId" form:"companyId" query:"companyId" `
	Name      string `json:"name" form:"name" query:"name" `
	Email     string `json:"email" form:"email" query:"email" `
	NoHp      string `json:"noHp" form:"noHp" query:"noHp" `
	Address   string `json:"address" form:"address" query:"address" `
}
