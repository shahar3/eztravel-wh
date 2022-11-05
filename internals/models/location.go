package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type City struct {
	MongoID     primitive.ObjectID `bson:"_id,omitempty"`
	ID          int                `json:"id" bson:"id,omitempty"`
	Name        string             `json:"name"`
	StateID     int                `json:"state_id"`
	StateCode   string             `json:"state_code"`
	StateName   string             `json:"state_name"`
	CountryID   int                `json:"country_id"`
	CountryCode string             `json:"country_code"`
	CountryName string             `json:"country_name"`
	Latitude    string             `json:"latitude"`
	Longitude   string             `json:"longitude"`
	Wikidataid  string             `json:"wikiDataId"`
}

type Country struct {
	MongoID        primitive.ObjectID `bson:"_id,omitempty"`
	ID             int                `json:"id"`
	Name           string             `json:"name"`
	Iso3           string             `json:"iso3"`
	Iso2           string             `json:"iso2"`
	NumericCode    string             `json:"numeric_code"`
	PhoneCode      string             `json:"phone_code"`
	Capital        string             `json:"capital"`
	Currency       string             `json:"currency"`
	CurrencyName   string             `json:"currency_name"`
	CurrencySymbol string             `json:"currency_symbol"`
	Tld            string             `json:"tld"`
	Native         string             `json:"native"`
	Region         string             `json:"region"`
	Subregion      string             `json:"subregion"`
	Timezones      []struct {
		Zonename      string `json:"zoneName"`
		Gmtoffset     int    `json:"gmtOffset"`
		Gmtoffsetname string `json:"gmtOffsetName"`
		Abbreviation  string `json:"abbreviation"`
		Tzname        string `json:"tzName"`
	} `json:"timezones"`
	Translations struct {
		Kr   string `json:"kr"`
		PtBr string `json:"pt-BR"`
		Pt   string `json:"pt"`
		Nl   string `json:"nl"`
		Hr   string `json:"hr"`
		Fa   string `json:"fa"`
		De   string `json:"de"`
		Es   string `json:"es"`
		Fr   string `json:"fr"`
		Ja   string `json:"ja"`
		It   string `json:"it"`
		Cn   string `json:"cn"`
		Tr   string `json:"tr"`
	} `json:"translations"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Emoji     string `json:"emoji"`
	Emojiu    string `json:"emojiU"`
}

type State struct {
	MongoID     primitive.ObjectID `bson:"_id,omitempty"`
	ID          int                `json:"id"`
	Name        string             `json:"name"`
	CountryID   int                `json:"country_id"`
	CountryCode string             `json:"country_code"`
	CountryName string             `json:"country_name"`
	StateCode   string             `json:"state_code"`
	Type        interface{}        `json:"type"`
	Latitude    string             `json:"latitude"`
	Longitude   string             `json:"longitude"`
}

type Location struct {
	LocationType string      `json:"location_type"`
	Data         interface{} `json:"data"`
}

type AutocompleteOption struct {
	Name         string `json:"name"`
	LocationType string `json:"locationType"`
	Country      string `json:"country"`
}
