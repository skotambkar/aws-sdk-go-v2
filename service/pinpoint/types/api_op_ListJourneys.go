// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListJourneysInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`

	PageSize *string `location:"querystring" locationName:"page-size" type:"string"`

	Token *string `location:"querystring" locationName:"token" type:"string"`
}

// String returns the string representation
func (s ListJourneysInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListJourneysInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListJourneysInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListJourneysOutput struct {
	_ struct{} `type:"structure" payload:"JourneysResponse"`

	// Provides information about the status, configuration, and other settings
	// for all the journeys that are associated with an application.
	//
	// JourneysResponse is a required field
	JourneysResponse *JourneysResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s ListJourneysOutput) String() string {
	return awsutil.Prettify(s)
}
