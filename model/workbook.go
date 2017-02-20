package model

type Workbook struct {
	Name        string `csv:"0" prn:"0"  json:"name"`
	Address     string `csv:"1" prn:"1"  json:"address"`
	Postcode    string `csv:"2" prn:"2"  json:"postcode"`
	Phone       string `csv:"3" prn:"3"  json:"phone"`
	CreditLimit string `csv:"4" prn:"4"  json:"creditlimit"`
	Birthday    string `csv:"5" prn:"5"  json:"birthday"`
}
