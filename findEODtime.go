package main

import (
	"time"
)

type PCCL struct {
	phoneCode                   string
	regionCode                  string
	currencyCode                string
	langCode                    string
	currencySymbol              string
	currencyNumberGroupingDelim string
	currencyHasCents            bool
	currencyDecimalSymbol       string
	dateFormat                  string
	domain                      string
	timeLocation                *time.Location
}

const (
	EN     = "en"      // English
	MS     = "ms"      // MY Bahasa
	ID     = "id"      // ID Bahasa
	VN     = "vn"      // Vietnamese
	TH     = "th"      // Thai
	ZhHant = "zh-Hant" // Traditional Chinese
	ZhHans = "zh-Hans" // Simplified Chinese
	PtBr   = "pt-BR"   // Portuguese
	HI     = "hi"      // Hindi
	MM     = "mm"      // Burmese
	FA     = "fa"      // Farsi
	ES     = "es-MX"   // Mexican
	ESCO   = "es-CO"   // Colombia
	ESCL   = "es-CL"   // Chile
	EsAr   = "es-AR"   // Argentina
	PL     = "pl"      // Poland
	EsEs   = "es-ES"   // Spanish-Spain
	FR     = "fr"      // France
)

var (
	pcclList = []*PCCL{
		{phoneCode: "65", regionCode: "SG", currencyCode: "SGD", langCode: EN, currencySymbol: "$",
			currencyNumberGroupingDelim: ",", currencyHasCents: true, currencyDecimalSymbol: ".",
			dateFormat: "02-01-2006", timeLocation: loadLocation("Etc/GMT-8"), domain: "shopee.sg"},

		{phoneCode: "66", regionCode: "TH", currencyCode: "THB", langCode: TH, currencySymbol: "฿",
			currencyNumberGroupingDelim: ",", currencyHasCents: false, currencyDecimalSymbol: ".",
			dateFormat: "02-01-2006", timeLocation: loadLocation("Etc/GMT-7"), domain: "shopee.co.th"},

		{phoneCode: "60", regionCode: "MY", currencyCode: "MYR", langCode: MS, currencySymbol: "RM",
			currencyNumberGroupingDelim: ",", currencyHasCents: true, currencyDecimalSymbol: ".",
			dateFormat: "02-01-2006", timeLocation: loadLocation("Etc/GMT-8"), domain: "shopee.com.my"},

		{phoneCode: "62", regionCode: "ID", currencyCode: "IDR", langCode: ID, currencySymbol: "Rp",
			currencyNumberGroupingDelim: ".", currencyHasCents: false, currencyDecimalSymbol: ",",
			dateFormat: "02-01-2006", timeLocation: loadLocation("Etc/GMT-7"), domain: "shopee.co.id"},

		{phoneCode: "63", regionCode: "PH", currencyCode: "PHP", langCode: EN, currencySymbol: "₱",
			currencyNumberGroupingDelim: ",", currencyHasCents: false, currencyDecimalSymbol: ".",
			dateFormat: "02-01-2006", timeLocation: loadLocation("Etc/GMT-8"), domain: "shopee.ph"},

		{phoneCode: "84", regionCode: "VN", currencyCode: "VND", langCode: VN, currencySymbol: "₫",
			currencyNumberGroupingDelim: ".", currencyHasCents: false, currencyDecimalSymbol: ",",
			dateFormat: "02-01-2006", timeLocation: loadLocation("Etc/GMT-7"), domain: "shopee.vn"},

		{phoneCode: "88", regionCode: "TW", currencyCode: "TWD", langCode: ZhHant, currencySymbol: "NT$",
			currencyNumberGroupingDelim: ",", currencyHasCents: false, currencyDecimalSymbol: ".",
			dateFormat: "02-01-2006", timeLocation: loadLocation("Etc/GMT-8"), domain: "shopee.tw"},

		{phoneCode: "98", regionCode: "IR", currencyCode: "IRR", langCode: FA, currencySymbol: "﷼",
			currencyNumberGroupingDelim: ",", currencyHasCents: false, currencyDecimalSymbol: ".",
			dateFormat: "02-01-2006", timeLocation: loadLocation("Iran"), domain: "shopee.co.ir"},

		{phoneCode: "95", regionCode: "MM", currencyCode: "MMK", langCode: MM, currencySymbol: "$",
			currencyNumberGroupingDelim: ",", currencyHasCents: false, currencyDecimalSymbol: ".",
			dateFormat: "02-01-2006", timeLocation: loadLocation("Asia/Rangoon"), domain: "shopee.com.mm"},

		{phoneCode: "852", regionCode: "HK", currencyCode: "HKD", langCode: ZhHant, currencySymbol: "$",
			currencyNumberGroupingDelim: ",", currencyHasCents: true, currencyDecimalSymbol: ".",
			dateFormat: "02-01-2006", timeLocation: loadLocation("Etc/GMT-8"), domain: "shopee.hk"},

		{phoneCode: "55", regionCode: "BR", currencyCode: "BRL", langCode: PtBr, currencySymbol: "R$",
			currencyNumberGroupingDelim: ".", currencyHasCents: true, currencyDecimalSymbol: ",",
			dateFormat: "02-01-2006", timeLocation: loadLocation("America/Fortaleza"), domain: "shopee.com.br"},

		{phoneCode: "91", regionCode: "IN", currencyCode: "INR", langCode: EN, currencySymbol: "₹",
			currencyNumberGroupingDelim: ",", currencyHasCents: true, currencyDecimalSymbol: ".",
			dateFormat: "02-01-2006", timeLocation: loadLocation("Asia/Calcutta"), domain: "shopee.in"},

		{phoneCode: "52", regionCode: "MX", currencyCode: "MXN", langCode: ES, currencySymbol: "MX$",
			currencyNumberGroupingDelim: ",", currencyHasCents: true, currencyDecimalSymbol: ".",
			dateFormat: "02-01-2006", timeLocation: loadLocation("America/Mexico_City"), domain: "shopee.com.mx"},

		{phoneCode: "57", regionCode: "CO", currencyCode: "COP", langCode: ESCO, currencySymbol: "COP",
			currencyNumberGroupingDelim: ".", currencyHasCents: false, currencyDecimalSymbol: ",",
			dateFormat: "02-01-2006", timeLocation: loadLocation("America/Bogota"), domain: "shopee.com.co"},

		{phoneCode: "56", regionCode: "CL", currencyCode: "CLP", langCode: ESCL, currencySymbol: "CLP",
			currencyNumberGroupingDelim: ".", currencyHasCents: false, currencyDecimalSymbol: ",",
			dateFormat: "02-01-2006", timeLocation: loadLocation("America/Santiago"), domain: "shopee.cl"},

		{phoneCode: "54", regionCode: "AR", currencyCode: "ARS", langCode: EsAr, currencySymbol: "$",
			currencyNumberGroupingDelim: ".", currencyHasCents: true, currencyDecimalSymbol: ",",
			dateFormat: "02-01-2006", timeLocation: loadLocation("America/Buenos_Aires"), domain: "shopee.com.ar"},

		{phoneCode: "48", regionCode: "PL", currencyCode: "PLN", langCode: PL, currencySymbol: "zł",
			currencyNumberGroupingDelim: " ", currencyHasCents: true, currencyDecimalSymbol: ",",
			dateFormat: "02-01-2006", timeLocation: loadLocation("Europe/Warsaw"), domain: "shopee.pl"},

		{phoneCode: "33", regionCode: "FR", currencyCode: "EUFR", langCode: FR, currencySymbol: "€",
			currencyNumberGroupingDelim: " ", currencyHasCents: true, currencyDecimalSymbol: ",",
			dateFormat: "02-01-2006", timeLocation: loadLocation("Europe/Paris"), domain: "shopee.fr"},

		{phoneCode: "34", regionCode: "ES", currencyCode: "EUES", langCode: EsEs, currencySymbol: "€",
			currencyNumberGroupingDelim: ".", currencyHasCents: true, currencyDecimalSymbol: ",",
			dateFormat: "02-01-2006", timeLocation: loadLocation("Europe/Warsaw"), domain: "shopee.es"},

		{phoneCode: "00", regionCode: "XX", currencyCode: "XXX", langCode: EN, currencySymbol: "$",
			currencyNumberGroupingDelim: ".", currencyHasCents: false, currencyDecimalSymbol: ",",
			dateFormat: "02-01-2006", timeLocation: loadLocation("Etc/GMT-7"), domain: "shopee.systems"},
	}
)

func (p *PCCL) GetTimeLocation() *time.Location {
	return p.timeLocation
}

func loadLocation(timeLocation string) *time.Location {
	loc, err := time.LoadLocation(timeLocation)
	if err != nil {
		panic(err)
	}
	return loc
}

func FindPCCL(code string) *PCCL {
	for _, pccl := range pcclList {
		if pccl.phoneCode == code || pccl.regionCode == code ||
			pccl.currencyCode == code || pccl.langCode == code {
			return pccl
		}
	}
	return nil
}

func GetLocalEOD(currTime int64, region string) int64 {
	convertedTime := time.Unix(currTime, 0)
	check := convertedTime.In(FindPCCL(region).GetTimeLocation())

	year, month, day := check.Date()
	return time.Date(year, month, day, 23, 59, 59, 0, check.Location()).Unix()
}
