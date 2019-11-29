// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to delete an existing Method resource.
type DeleteMethodInput struct {
	_ struct{} `type:"structure"`

	// [Required] The HTTP verb of the Method resource.
	//
	// HttpMethod is a required field
	HttpMethod *string `location:"uri" locationName:"http_method" type:"string" required:"true"`

	// [Required] The Resource identifier for the Method resource.
	//
	// ResourceId is a required field
	ResourceId *string `location:"uri" locationName:"resource_id" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMethodInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMethodInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteMethodInput"}

	if s.HttpMethod == nil {
		invalidParams.Add(aws.NewErrParamRequired("HttpMethod"))
	}

	if s.ResourceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ResourceId"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteMethodOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMethodOutput) String() string {
	return awsutil.Prettify(s)
}