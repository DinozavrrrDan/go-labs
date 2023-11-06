package app

import "github.com/DinozavrrrDan/go-module/pkg/contains"

func Search(file string, subString string) (bool, error) {
	return contains.Contains(file, subString)
}
