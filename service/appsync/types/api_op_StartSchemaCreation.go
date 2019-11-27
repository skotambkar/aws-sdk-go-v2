// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/appsync/enums"
)

type StartSchemaCreationInput struct {
	_ struct{} `type:"structure"`

	// The API ID.
	//
	// ApiId is a required field
	ApiId *string `location:"uri" locationName:"apiId" type:"string" required:"true"`

	// The schema definition, in GraphQL schema language format.
	//
	// Definition is automatically base64 encoded/decoded by the SDK.
	//
	// Definition is a required field
	Definition []byte `locationName:"definition" type:"blob" required:"true"`
}

// String returns the string representation
func (s StartSchemaCreationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartSchemaCreationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartSchemaCreationInput"}

	if s.ApiId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApiId"))
	}

	if s.Definition == nil {
		invalidParams.Add(aws.NewErrParamRequired("Definition"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartSchemaCreationOutput struct {
	_ struct{} `type:"structure"`

	// The current state of the schema (PROCESSING, FAILED, SUCCESS, or NOT_APPLICABLE).
	// When the schema is in the ACTIVE state, you can add data.
	Status enums.SchemaStatus `locationName:"status" type:"string" enum:"true"`
}

// String returns the string representation
func (s StartSchemaCreationOutput) String() string {
	return awsutil.Prettify(s)
}
