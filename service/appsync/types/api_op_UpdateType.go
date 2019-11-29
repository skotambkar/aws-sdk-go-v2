// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/appsync/enums"
)

type UpdateTypeInput struct {
	_ struct{} `type:"structure"`

	// The API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The new definition.
	Definition *string `locationName:"definition" type:"string"`

	// The new type format: SDL or JSON.
	//
	// Format is a required field
	Format enums.TypeDefinitionFormat `locationName:"format" type:"string" required:"true" enum:"true"`

	// The new type name.
	//
	// TypeName is a required field
	TypeName *string `location:"uri" locationName:"typeName" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateTypeInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateTypeInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateTypeInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}
	if len(s.Format) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Format"))
	}

	if s.TypeName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TypeName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateTypeOutput struct {
	_ struct{} `type:"structure"`

	// The updated Type object.
	Type *Type `locationName:"type" type:"structure"`
}

// String returns the string representation
func (s UpdateTypeOutput) String() string {
	return awsutil.Prettify(s)
}