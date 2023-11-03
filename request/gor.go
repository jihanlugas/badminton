package request

type CreateGor struct {
	CompanyID       string `json:"companyId" validate:"required,existsdata=company_id"`
	Name            string `json:"name" validate:"required"`
	Description     string `json:"description" validate:""`
	Address         string `json:"address" validate:""`
	NormalGamePrice int64  `json:"normalGamePrice" validate:""`
	RubberGamePrice int64  `json:"rubberGamePrice" validate:""`
	BallPrice       int64  `json:"ballPrice" validate:""`
}

type UpdateGor struct {
	CompanyID       string `json:"companyId" validate:"required,existsdata=company_id"`
	Name            string `json:"name" validate:"required"`
	Description     string `json:"description" validate:""`
	Address         string `json:"address" validate:""`
	NormalGamePrice int64  `json:"normalGamePrice" validate:""`
	RubberGamePrice int64  `json:"rubberGamePrice" validate:""`
	BallPrice       int64  `json:"ballPrice" validate:""`
}

type PageGor struct {
	Paging
	CompanyID   string `json:"companyId" form:"companyId" query:"companyId"`
	Name        string `json:"name" form:"name" query:"name"`
	Description string `json:"description" form:"description" query:"description"`
	Address     string `json:"address" form:"address" query:"address"`
}
