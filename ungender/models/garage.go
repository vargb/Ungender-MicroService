package models

type Garage struct {
	CarID            string `json:"carid" gorm:"primaryKey"`
	DateofManufac    string `json:"dom"`
	LastServicedDate string `json:"lsd"`
	LastUsedDate     string `json:"lud"`
	Available        bool   `json:"available"`
}
