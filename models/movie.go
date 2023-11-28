package models

type Movie struct{
	ID int `json:"id"`
	Movie string `json:"movie"`
	Director string `json:"director"`
	Year int `json:"year"`
} 