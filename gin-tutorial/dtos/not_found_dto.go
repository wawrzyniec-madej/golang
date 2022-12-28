package dtos

type NotFoundDTO struct {
	Message string `json:"message"`
}

func NewNotFoundDTO() *NotFoundDTO {
	return &NotFoundDTO{
		Message: "Not found",
	}
}
