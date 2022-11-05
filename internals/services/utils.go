package services

import (
	"eztravel-wh/internals/models"
)

func getCountryName(locationObject interface{}, locationType string) string {
	switch locationType {
	case city:
		return locationObject.(models.City).CountryName
	case country:
		return locationObject.(models.Country).Name
	case state:
		return locationObject.(models.State).CountryName
	}

	return ""
}
