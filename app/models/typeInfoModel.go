package models

type TypeInfo struct {
	ID          uint           `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	SubType     []*SubTypeInfo `json:"sub_type" gorm:"foreignKey:TypeId"`
}

type TypeInfoResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	SubType     []*struct {
		ID          uint   `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
	} `json:"sub_type"`
}
