// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAccessKeyInfoInput struct {
	_ struct{} `type:"structure"`

	// The identifier of an access key.
	//
	// This parameter allows (through its regex pattern) a string of characters
	// that can consist of any upper- or lowercased letter or digit.
	//
	// AccessKeyId is a required field
	AccessKeyId *string `min:"16" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAccessKeyInfoInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAccessKeyInfoInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAccessKeyInfoInput"}

	if s.AccessKeyId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AccessKeyId"))
	}
	if s.AccessKeyId != nil && len(*s.AccessKeyId) < 16 {
		invalidParams.Add(aws.NewErrParamMinLen("AccessKeyId", 16))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetAccessKeyInfoOutput struct {
	_ struct{} `type:"structure"`

	// The number used to identify the AWS account.
	Account *string `type:"string"`
}

// String returns the string representation
func (s GetAccessKeyInfoOutput) String() string {
	return awsutil.Prettify(s)
}