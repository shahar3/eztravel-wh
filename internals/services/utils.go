package services

import (
	"eztravel-wh/internals/models"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func extractStateNameAndLocation(locationMongoObject map[string]interface{}, locationObject *models.Location, locationName string, mongoId primitive.ObjectID) (*models.Location, string) {
	var stateObj models.State
	if err := mapstructure.Decode(locationMongoObject, &stateObj); err != nil {
		fmt.Println(err)
	}
	stateObj.MongoID = mongoId
	locationObject = &models.Location{
		LocationType: state,
		Data:         stateObj,
	}
	locationName = stateObj.Name
	return locationObject, locationName
}

func extractCountryNameAndLocation(locationMongoObject map[string]interface{}, locationObject *models.Location, locationName string, mongoId primitive.ObjectID) (*models.Location, string) {
	var countryObj models.Country
	if err := mapstructure.Decode(locationMongoObject, &countryObj); err != nil {
		fmt.Println(err)
	}
	countryObj.MongoID = mongoId
	locationObject = &models.Location{
		LocationType: country,
		Data:         countryObj,
	}
	locationName = countryObj.Name
	return locationObject, locationName
}

func extractCityNameAndLocation(locationMongoObject map[string]interface{}, locationObject *models.Location, locationName string, mongoId primitive.ObjectID) (*models.Location, string) {
	var cityObj models.City
	if err := mapstructure.Decode(locationMongoObject, &cityObj); err != nil {
		fmt.Println(err)
	}
	cityObj.MongoID = mongoId
	locationObject = &models.Location{
		LocationType: city,
		Data:         cityObj,
	}
	locationName = cityObj.Name
	return locationObject, locationName
}

func extractMongoId(location *models.Location) string {
	switch location.LocationType {
	case city:
		var cityObj models.City
		_ = mapstructure.Decode(location.Data, &cityObj)
		return cityObj.MongoID.Hex()
	case country:
		var countryObj models.Country
		_ = mapstructure.Decode(location.Data, &countryObj)
		return countryObj.MongoID.Hex()
	case state:
		var stateObj models.State
		_ = mapstructure.Decode(location.Data, &stateObj)
		return stateObj.MongoID.Hex()
	}
	return ""
}

func convertToLocation(locationMongoObject map[string]interface{}) *models.Location {
	var location *models.Location
	switch locationMongoObject[locationType] {
	case city:
		var cityObj models.City
		mapstructure.Decode(locationMongoObject, &cityObj)
		location = &models.Location{
			LocationType: city,
			Data:         cityObj,
		}
	case country:
		var countryObj models.Country
		mapstructure.Decode(locationMongoObject, &countryObj)
		location = &models.Location{
			LocationType: country,
			Data:         countryObj,
		}
	case state:
		var stateObj models.State
		mapstructure.Decode(locationMongoObject, &stateObj)
		location = &models.Location{
			LocationType: state,
			Data:         stateObj,
		}
	}

	return location
}
