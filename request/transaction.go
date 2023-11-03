package request

type CreateTransaction struct {
	CompanyID string `json:"companyId" validate:"required,existsdata=company_id"`
	Name      string `json:"name" validate:"required"`
	IsDebit   bool   `json:"isDebit" validate:""`
	Price     int64  `json:"price" validate:"required"`
}

type PageTransaction struct {
	Paging
	CompanyID string `json:"companyId" form:"companyId" query:"companyId"`
	Name      string `json:"name" form:"name" query:"name"`
}
