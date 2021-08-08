package model

// instance entity
type Instance struct {
	ID     uint64 `json:"id,omitempty"`
	Name   string `json:"name,omitempty" gorm:"type:varchar(255)"`
	Status int    `json:"age,omitempty"`
	BaseModel
}

// user DAO
type InstanceDAO interface {
	BaseDAO
}
