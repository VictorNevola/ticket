package company

type (
	CreateData struct {
		Name  string `json:"name" validate:"required,min=3,max=255"`
		TaxID string `json:"tax_id" validate:"required,min=3,max=255,isCNPJORCPF"`
	}
)
