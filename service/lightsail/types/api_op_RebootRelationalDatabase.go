// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RebootRelationalDatabaseInput struct {
	_ struct{} `type:"structure"`

	// The name of your database to reboot.
	//
	// RelationalDatabaseName is a required field
	RelationalDatabaseName *string `locationName:"relationalDatabaseName" type:"string" required:"true"`
}

// String returns the string representation
func (s RebootRelationalDatabaseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RebootRelationalDatabaseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RebootRelationalDatabaseInput"}

	if s.RelationalDatabaseName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RelationalDatabaseName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RebootRelationalDatabaseOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the result of your reboot relational database request.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s RebootRelationalDatabaseOutput) String() string {
	return awsutil.Prettify(s)
}
