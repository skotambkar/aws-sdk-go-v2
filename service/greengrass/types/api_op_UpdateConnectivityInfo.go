// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Information required to update a Greengrass core's connectivity.
type UpdateConnectivityInfoInput struct {
	_ struct{} `type:"structure"`

	// A list of connectivity info.
	ConnectivityInfo []ConnectivityInfo `type:"list"`

	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"ThingName" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateConnectivityInfoInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConnectivityInfoInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConnectivityInfoInput"}

	if s.ThingName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ThingName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateConnectivityInfoOutput struct {
	_ struct{} `type:"structure"`

	// A message about the connectivity info update request.
	Message *string `locationName:"message" type:"string"`

	// The new version of the connectivity info.
	Version *string `type:"string"`
}

// String returns the string representation
func (s UpdateConnectivityInfoOutput) String() string {
	return awsutil.Prettify(s)
}