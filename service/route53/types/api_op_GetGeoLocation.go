// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request for information about whether a specified geographic location is
// supported for Amazon Route 53 geolocation resource record sets.
type GetGeoLocationInput struct {
	_ struct{} `type:"structure"`

	// Amazon Route 53 supports the following continent codes:
	//
	//    * AF: Africa
	//
	//    * AN: Antarctica
	//
	//    * AS: Asia
	//
	//    * EU: Europe
	//
	//    * OC: Oceania
	//
	//    * NA: North America
	//
	//    * SA: South America
	ContinentCode *string `location:"querystring" locationName:"continentcode" min:"2" type:"string"`

	// Amazon Route 53 uses the two-letter country codes that are specified in ISO
	// standard 3166-1 alpha-2 (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	CountryCode *string `location:"querystring" locationName:"countrycode" min:"1" type:"string"`

	// Amazon Route 53 uses the one- to three-letter subdivision codes that are
	// specified in ISO standard 3166-1 alpha-2 (https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2).
	// Route 53 doesn't support subdivision codes for all countries. If you specify
	// subdivisioncode, you must also specify countrycode.
	SubdivisionCode *string `location:"querystring" locationName:"subdivisioncode" min:"1" type:"string"`
}

// String returns the string representation
func (s GetGeoLocationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetGeoLocationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetGeoLocationInput"}
	if s.ContinentCode != nil && len(*s.ContinentCode) < 2 {
		invalidParams.Add(aws.NewErrParamMinLen("ContinentCode", 2))
	}
	if s.CountryCode != nil && len(*s.CountryCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CountryCode", 1))
	}
	if s.SubdivisionCode != nil && len(*s.SubdivisionCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("SubdivisionCode", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A complex type that contains the response information for the specified geolocation
// code.
type GetGeoLocationOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains the codes and full continent, country, and subdivision
	// names for the specified geolocation code.
	//
	// GeoLocationDetails is a required field
	GeoLocationDetails *GeoLocationDetails `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetGeoLocationOutput) String() string {
	return awsutil.Prettify(s)
}