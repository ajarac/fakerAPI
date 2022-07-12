package dto

type BasicSchemaDTO struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func NewBasicSchemaDTO(id string, name string) *BasicSchemaDTO {
	return &BasicSchemaDTO{Id: id, Name: name}
}
