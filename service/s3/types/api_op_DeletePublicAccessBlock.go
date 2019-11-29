// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeletePublicAccessBlockInput struct {
	_ struct{} `type:"structure"`

	// The Amazon S3 bucket whose PublicAccessBlock configuration you want to delete.
	//
	// Bucket is a required field
	Bucket *string `location:"uri" locationName:"Bucket" type:"string" required:"true"`
}

// String returns the string representation
func (s DeletePublicAccessBlockInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeletePublicAccessBlockInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeletePublicAccessBlockInput"}

	if s.Bucket == nil {
		invalidParams.Add(aws.NewErrParamRequired("Bucket"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (s *DeletePublicAccessBlockInput) GetBucket() (v string) {
	if s.Bucket == nil {
		return v
	}
	return *s.Bucket
}

type DeletePublicAccessBlockOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeletePublicAccessBlockOutput) String() string {
	return awsutil.Prettify(s)
}
