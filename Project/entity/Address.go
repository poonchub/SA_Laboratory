package entity

import "gorm.io/gorm"

type Address struct {
	gorm.Model
	Province		string
	District		string
	Subdistrict		string
	ZipCode			string
	AddressDetail	string

	Order	[]Order	`gorm:"foreignKey:AddressID"`

	CustomerID		uint
}