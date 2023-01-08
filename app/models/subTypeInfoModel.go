package models

type SubTypeInfo struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	TypeId      uint   `json:"type_id"`
}
