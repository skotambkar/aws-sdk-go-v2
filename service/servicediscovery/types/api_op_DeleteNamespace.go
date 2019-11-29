// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteNamespaceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the namespace that you want to delete.
	//
	// Id is a required field
	Id *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteNamespaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteNamespaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteNamespaceInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteNamespaceOutput struct {
	_ struct{} `type:"structure"`

	// A value that you can use to determine whether the request completed successfully.
	// To get the status of the operation, see GetOperation.
	OperationId *string `type:"string"`
}

// String returns the string representation
func (s DeleteNamespaceOutput) String() string {
	return awsutil.Prettify(s)
}