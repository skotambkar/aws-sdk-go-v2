// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetRelationalDatabaseLogStreamsInput struct {
	_ struct{} `type:"structure"`

	// The name of your database for which to get log streams.
	//
	// RelationalDatabaseName is a required field
	RelationalDatabaseName *string `locationName:"relationalDatabaseName" type:"string" required:"true"`
}

// String returns the string representation
func (s GetRelationalDatabaseLogStreamsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetRelationalDatabaseLogStreamsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetRelationalDatabaseLogStreamsInput"}

	if s.RelationalDatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetRelationalDatabaseLogStreamsOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the result of your get relational database log streams
	// request.
	LogStreams []string `locationName:"logStreams" type:"list"`
}

// String returns the string representation
func (s GetRelationalDatabaseLogStreamsOutput) String() string {
	return awsutil.Prettify(s)
}