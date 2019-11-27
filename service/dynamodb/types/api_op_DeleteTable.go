// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DeleteTable operation.
type DeleteTableInput struct {
	_ struct{} `type:"structure"`

	// The name of the table to delete.
	//
	// TableName is a required field
	TableName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteTableInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteTableInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteTableInput"}

	if s.TableName == nil {
		invalidParams.Add(aws.NewErrParamRequired("TableName"))
	}
	if s.TableName != nil && len(*s.TableName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("TableName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of a DeleteTable operation.
type DeleteTableOutput struct {
	_ struct{} `type:"structure"`

	// Represents the properties of a table.
	TableDescription *TableDescription `type:"structure"`
}

// String returns the string representation
func (s DeleteTableOutput) String() string {
	return awsutil.Prettify(s)
}
