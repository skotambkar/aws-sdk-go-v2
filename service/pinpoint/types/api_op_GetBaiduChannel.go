// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetBaiduChannelInput struct {
	_ struct{} `type:"structure"`

	// ApplicationId is a required field
	ApplicationId *string `location:"uri" locationName:"application-id" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBaiduChannelInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBaiduChannelInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBaiduChannelInput"}

	if s.ApplicationId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ApplicationId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetBaiduChannelOutput struct {
	_ struct{} `type:"structure" payload:"BaiduChannelResponse"`

	// Provides information about the status and settings of the Baidu (Baidu Cloud
	// Push) channel for an application.
	//
	// BaiduChannelResponse is a required field
	BaiduChannelResponse *BaiduChannelResponse `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetBaiduChannelOutput) String() string {
	return awsutil.Prettify(s)
}