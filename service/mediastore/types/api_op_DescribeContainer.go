// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeContainerInput struct {
	_ struct{} `type:"structure"`

	// The name of the container to query.
	ContainerName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s DescribeContainerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeContainerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeContainerInput"}
	if s.ContainerName != nil && len(*s.ContainerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ContainerName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeContainerOutput struct {
	_ struct{} `type:"structure"`

	// The name of the queried container.
	Container *Container `type:"structure"`
}

// String returns the string representation
func (s DescribeContainerOutput) String() string {
	return awsutil.Prettify(s)
}