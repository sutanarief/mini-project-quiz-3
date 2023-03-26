package helper

import (
	"errors"
	"regexp"
)

func ThicknessHandler(totalPage int) string {
	switch {
	case totalPage <= 100:
		return "tipis"
	case totalPage > 100 && totalPage <= 200:
		return "sedang"
	case totalPage >= 201:
		return "tebal"
	default:
		return "-"
	}
}

func UrlValidator(imageUrl string) bool {
	regex, err := regexp.Compile(`^(http|https)://`)

	if err != nil {
		return false
	}

	isValidUrl := regex.Match([]byte(imageUrl))

	return isValidUrl
}

func YearValidator(releasYear int64) bool {
	if releasYear < 1980 || releasYear > 2021 {
		return false
	} else {
		return true
	}
}

func ErrorHandler(isValidYear bool, isValidURL bool) error {
	if !isValidURL && !isValidYear {
		return errors.New("Image URL & Release Year tidak valid")
	} else if !isValidURL {
		return errors.New("Image URL tidak valid")
	} else if !isValidYear {
		return errors.New("Release Year tidak valid")
	}
	return nil
}
