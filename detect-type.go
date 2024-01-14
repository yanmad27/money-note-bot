package main

import (
	"strings"
)

func detectType(what string) string {
	what = strings.ToLower(what)

	device_keyword := []string{
		"điện",
		"đèn",
		"tivi",
		"tủ lạnh",
		"máy",
		"may",
		"lò",
		"rac",
		"rác",
		"rửa",
		"bếp",
		"chén",
		"bát",
		"nồi",
		"xoong",
		"chảo",
		"bình",
		"nước",
	}
	for _, keyword := range device_keyword {
		if strings.Contains(what, keyword) {
			return "Gia dụng"
		}
	}

	medicine_keyword := [5]string{"thuốc",
		"viên",
		"hộp",
		"gói",
		"chai"}
	for _, keyword := range medicine_keyword {
		if strings.Contains(what, keyword) {
			return "Thuốc"
		}
	}

	return "Food"
}
