// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request to delete an existing model in an existing RestApi resource.
type DeleteModelInput struct {
	_ struct{} `type:"structure"`

	// [Required] The name of the model to delete.
	//
	// ModelName is a required field
	ModelName *string `location:"uri" locationName:"model_name" type:"string" required:"true"`

	// [Required] The string identifier of the associated RestApi.
	//
	// RestApiId is a required field
	RestApiId *string `location:"uri" locationName:"restapi_id" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteModelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteModelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteModelInput"}

	if s.ModelName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ModelName"))
	}

	if s.RestApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("RestApiId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteModelOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteModelOutput) String() string {
	return awsutil.Prettify(s)
}