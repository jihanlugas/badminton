package request

type CreateCompany struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:""`
	Balance     int64  `json:"balance" validate:""`
}

type UpdateCompany struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:""`
	Balance     int64  `json:"balance" validate:""`
}

type PageCompany struct {
	Paging
	Name        string `json:"name" form:"name" query:"name"`
	Description string `json:"description" form:"description" query:"description"`
}
