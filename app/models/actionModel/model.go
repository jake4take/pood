package actionModel

type Action struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}

type ActionCreateRequest struct {
	Name string `json:"name"`
}
