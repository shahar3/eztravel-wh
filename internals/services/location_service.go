package services

import (
	"context"
	"eztravel-wh/internals/models"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	GetLocation(string) (*models.Location, error)
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
	findOptions := options.FindOptions{}
	opt := *findOptions.SetLimit(100)
	query := bson.M{}

	cursor, err := ls.locationsCollection.Find(ls.ctx, query, &opt)
	if err != nil {
		fmt.Println(err)
	}

	defer cursor.Close(ls.ctx)

	ls.locations = make(map[string]*models.Location)

	for cursor.Next(ls.ctx) {
		locationObject, locationName := ls.extractLocation(cursor)

		ls.locations[locationName] = locationObject
	}

	if err = cursor.Err(); err != nil {
		fmt.Println(err)
	}

	fmt.Println("loaded locations")
}

func (ls *LocationServiceImplementation) extractLocation(cursor *mongo.Cursor) (*models.Location, string) {
	var locationMongoObject map[string]interface{}
	if err := cursor.Decode(&locationMongoObject); err != nil {
		fmt.Println(err)
	}

	mongoId := locationMongoObject["_id"].(primitive.ObjectID)
	var locationObject *models.Location
	var locationName string
	switch locationMongoObject[locationType] {
	case city:
		return extractCityNameAndLocation(locationMongoObject, locationObject, locationName, mongoId)
	case country:
		return extractCountryNameAndLocation(locationMongoObject, locationObject, locationName, mongoId)
	case state:
		return extractStateNameAndLocation(locationMongoObject, locationObject, locationName, mongoId)
	}
	return locationObject, locationName
}

func (ls LocationServiceImplementation) GetLocation(mongoIdHex string) (*models.Location, error) {
	objectID, _ := primitive.ObjectIDFromHex(mongoIdHex)
	filter := bson.M{"_id": objectID}
	result := ls.locationsCollection.FindOne(ls.ctx, filter)
	if result.Err() != nil {
		return nil, result.Err()
	}

	var locationMongoObject map[string]interface{}
	if err := result.Decode(&locationMongoObject); err != nil {
		return nil, err
	}

	//convert to models.location
	return convertToLocation(locationMongoObject), nil
}

func (ls LocationServiceImplementation) GetLocationsAutoComplete(term string) []*models.AutocompleteOption {
	term = strings.ToLower(term)

	var matchedOptions []*models.AutocompleteOption
	for locationName, location := range ls.locations {
		locationName = strings.ToLower(locationName)
		if strings.HasPrefix(locationName, term) {
			autocompleteOption := &models.AutocompleteOption{
				Id:           extractMongoId(location),
				Name:         locationName,
				LocationType: location.LocationType,
				Country:      getCountryName(location.Data, location.LocationType),
			}
			matchedOptions = append(matchedOptions, autocompleteOption)
		}
	}

	return matchedOptions
}
