package models

type Property struct {
	PropertyId int `json:"propertyid"`
	Name string `json:"name"`
	Category string `json:"category"`
	SubCategory string `json:"subcategory"`
	Data string `json:"data"`
	Url string `json:"uri"`
	Divisible bool `json:"divisible"`
}
