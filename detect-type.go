package main

import (
	"strings"
)

func detectType(what string) string {
	what = strings.ToLower(what)

	food_keyword := []string{"bbq", "food", "rau", "nước", "củ", "trái", "bánh", "kem", "mì", "bún", "phở", "cơm", "cháo", "chè", "cá", "thịt"}
	for _, keyword := range food_keyword {
		if strings.Contains(what, keyword) {
			return "Food"
		}
	}

	medicine_keyword := [5]string{"thuốc", "viên", "hộp", "gói", "chai"}
	for _, keyword := range medicine_keyword {
		if strings.Contains(what, keyword) {
			return "Thuốc"
		}
	}

	return "Gia dụng"
}
