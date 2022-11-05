package services

import (
	"context"
	"eztravel-wh/internals/models"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"strings"
)

const (
	locationType = "locationType"
	city         = "city"
	state        = "State"
	country      = "country"
	name         = "name"
)

type LocationService interface {
	GetLocation() (*models.Location, error)
	GetLocationsAutoComplete(string) []*models.AutocompleteOption
	Init()
}

type LocationServiceImplementation struct {
	locationsCollection *mongo.Collection
	locations           map[string]*models.Location
	ctx                 context.Context
}

func NewLocationService(locationsCollection *mongo.Collection, ctx context.Context) LocationService {
	return &LocationServiceImplementation{
		locationsCollection: locationsCollection,
		ctx:                 ctx,
	}
}

func (ls *LocationServiceImplementation) Init() {
	query := bson.M{}

	cursor, err := ls.locationsCollection.Find(ls.ctx, query)
	if err != nil {
		fmt.Println(err)
	}

	defer cursor.Close(ls.ctx)

	ls.locations = make(map[string]*models.Location)

	for cursor.Next(ls.ctx) {
		var locationMongoObject map[string]interface{}
		if err = cursor.Decode(&locationMongoObject); err != nil {
			fmt.Println(err)
		}

		var locationObject *models.Location
		var locationName string
		switch locationMongoObject[locationType] {
		case city:
			var cityObj models.City
			if err = mapstructure.Decode(locationMongoObject, &cityObj); err != nil {
				fmt.Println(err)
			}
			locationObject = &models.Location{
				LocationType: city,
				Data:         cityObj,
			}
			locationName = cityObj.Name
		case country:
			var countryObj models.Country
			if err = mapstructure.Decode(locationMongoObject, &countryObj); err != nil {
				fmt.Println(err)
			}
			locationObject = &models.Location{
				LocationType: country,
				Data:         countryObj,
			}
			locationName = countryObj.Name
		case state:
			var stateObj models.State
			if err = mapstructure.Decode(locationMongoObject, &stateObj); err != nil {
				fmt.Println(err)
			}
			locationObject = &models.Location{
				LocationType: state,
				Data:         stateObj,
			}
			locationName = stateObj.Name
		}

		ls.locations[locationName] = locationObject
	}

	if err = cursor.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("loaded locations")
}

func (ls LocationServiceImplementation) GetLocation() (*models.Location, error) {
	return nil, nil
}

func (ls LocationServiceImplementation) GetLocationsAutoComplete(term string) []*models.AutocompleteOption {
	term = strings.ToLower(term)

	var matchedOptions []*models.AutocompleteOption
	for locationName, location := range ls.locations {
		locationName = strings.ToLower(locationName)
		if strings.HasPrefix(locationName, term) {
			autocompleteOption := &models.AutocompleteOption{
				Name:         locationName,
				LocationType: location.LocationType,
				Country:      getCountryName(location.Data, location.LocationType),
			}
			matchedOptions = append(matchedOptions, autocompleteOption)
		}
	}

	return matchedOptions
}
