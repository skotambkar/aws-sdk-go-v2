// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetAccessKeyLastUsedInput struct {
	_ struct{} `type:"structure"`

	// The identifier of an access key.
	//
	// This parameter allows (through its regex pattern (http://wikipedia.org/wiki/regex))
	// a string of characters that can consist of any upper or lowercased letter
	// or digit.
	//
	// AccessKeyId is a required field
	AccessKeyId *string `min:"16" type:"string" required:"true"`
}

// String returns the string representation
func (s GetAccessKeyLastUsedInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetAccessKeyLastUsedInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetAccessKeyLastUsedInput"}

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

// Contains the response to a successful GetAccessKeyLastUsed request. It is
// also returned as a member of the AccessKeyMetaData structure returned by
// the ListAccessKeys action.
type GetAccessKeyLastUsedOutput struct {
	_ struct{} `type:"structure"`

	// Contains information about the last time the access key was used.
	AccessKeyLastUsed *AccessKeyLastUsed `type:"structure"`

	// The name of the AWS IAM user that owns this access key.
	UserName *string `min:"1" type:"string"`
}

// String returns the string representation
func (s GetAccessKeyLastUsedOutput) String() string {
	return awsutil.Prettify(s)
}
