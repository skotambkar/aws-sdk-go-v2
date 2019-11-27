// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to delete a hosted zone.
type DeleteHostedZoneInput struct {
	_ struct{} `type:"structure"`

	// The ID of the hosted zone you want to delete.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteHostedZoneInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteHostedZoneInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteHostedZoneInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A complex type that contains the response to a DeleteHostedZone request.
type DeleteHostedZoneOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains the ID, the status, and the date and time of
	// a request to delete a hosted zone.
	//
	// ChangeInfo is a required field
	ChangeInfo *ChangeInfo `type:"structure" required:"true"`
}

// String returns the string representation
func (s DeleteHostedZoneOutput) String() string {
	return awsutil.Prettify(s)
}
