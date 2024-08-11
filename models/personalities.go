package models

type Personality struct {
	Id int `json:"id" validate:"number"`

	// validate package is used to add validation rules to struct fields

	Name string `json:"name" validate:"required"`
	Bio  string `json:"bio" validate:"required"`
}

type CreatePersonality struct {
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

func (p CreatePersonality) TableName() string {
	return "personalities"
}

// GORM associates the Personality struct with the personalities table automatically
// If the table name is different from the struct name, you can use the following:
// func (Personality) TableName() string {
// 	return "table_name"
// }
// This is useful when you have a struct with a name that is not the plural of the table name
// It's a struct method, so it should be defined in the same file as the struct
