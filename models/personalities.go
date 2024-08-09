package models

type Personality struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Bio  string `json:"bio"`
}

// GORM associates the Personality struct with the personalities table automatically
// If the table name is different from the struct name, you can use the following:
// func (Personality) TableName() string {
// 	return "table_name"
// }
// This is useful when you have a struct with a name that is not the plural of the table name
// It's a struct method, so it should be defined in the same file as the struct
