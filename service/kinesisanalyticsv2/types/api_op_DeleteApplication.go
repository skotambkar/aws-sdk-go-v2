// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteApplicationInput struct {
	_ struct{} `type:"structure"`

	// The name of the application to delete.
	//
	// ApplicationName is a required field
	ApplicationName *string `min:"1" type:"string" required:"true"`

	// Use the DescribeApplication operation to get this value.
	//
	// CreateTimestamp is a required field
	CreateTimestamp *time.Time `type:"timestamp" required:"true"`
}

// String returns the string representation
func (s DeleteApplicationInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteApplicationInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteApplicationInput"}

	if s.ApplicationName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationName"))
	}
	if s.ApplicationName != nil && len(*s.ApplicationName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ApplicationName", 1))
	}

	if s.CreateTimestamp == nil {
		invalidParams.Add(aws.NewErrParamRequired("CreateTimestamp"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteApplicationOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteApplicationOutput) String() string {
	return awsutil.Prettify(s)
}
