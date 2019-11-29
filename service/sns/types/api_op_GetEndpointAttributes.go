// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Input for GetEndpointAttributes action.
type GetEndpointAttributesInput struct {
	_ struct{} `type:"structure"`

	// EndpointArn for GetEndpointAttributes input.
	//
	// EndpointArn is a required field
	EndpointArn *string `type:"string" required:"true"`
}

// String returns the string representation
func (s GetEndpointAttributesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetEndpointAttributesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetEndpointAttributesInput"}

	if s.EndpointArn == nil {
		invalidParams.Add(aws.NewErrParamRequired("EndpointArn"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Response from GetEndpointAttributes of the EndpointArn.
type GetEndpointAttributesOutput struct {
	_ struct{} `type:"structure"`

	// Attributes include the following:
	//
	//    * CustomUserData – arbitrary user data to associate with the endpoint.
	//    Amazon SNS does not use this data. The data must be in UTF-8 format and
	//    less than 2KB.
	//
	//    * Enabled – flag that enables/disables delivery to the endpoint. Amazon
	//    SNS will set this to false when a notification service indicates to Amazon
	//    SNS that the endpoint is invalid. Users can set it back to true, typically
	//    after updating Token.
	//
	//    * Token – device token, also referred to as a registration id, for an
	//    app and mobile device. This is returned from the notification service
	//    when an app and mobile device are registered with the notification service.
	Attributes map[string]string `type:"map"`
}

// String returns the string representation
func (s GetEndpointAttributesOutput) String() string {
	return awsutil.Prettify(s)
}