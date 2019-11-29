// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateModelInput struct {
	_ struct{} `type:"structure"`

	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// A string with a length between [1-256].
	ContentType *string `locationName:"contentType" type:"string"`

	// A string with a length between [0-1024].
	Description *string `locationName:"description" type:"string"`

	// A string with a length between [1-128].
	//
	// Name is a required field
	Name *string `locationName:"name" type:"string" required:"true"`

	// A string with a length between [0-32768].
	//
	// Schema is a required field
	Schema *string `locationName:"schema" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateModelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateModelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateModelInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if s.Schema == nil {
		invalidParams.Add(aws.NewErrParamRequired("Schema"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateModelOutput struct {
	_ struct{} `type:"structure"`

	// A string with a length between [1-256].
	ContentType *string `locationName:"contentType" type:"string"`

	// A string with a length between [0-1024].
	Description *string `locationName:"description" type:"string"`

	// The identifier.
	ModelId *string `locationName:"modelId" type:"string"`

	// A string with a length between [1-128].
	Name *string `locationName:"name" type:"string"`

	// A string with a length between [0-32768].
	Schema *string `locationName:"schema" type:"string"`
}

// String returns the string representation
func (s CreateModelOutput) String() string {
	return awsutil.Prettify(s)
}
