package models

type Garage struct {
	CarID            string `json:"carid"`
	DateofManufac    string `json:"dom"`
	LastServicedDate string `json:"lsd"`
	UniqueID         string `json:"uniqueid"`
	LastUsedDate     string `json:"lud"`
	Available        bool   `json:"available"`
}
