// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateGroupVersionInput struct {
	_ struct{} `type:"structure"`

	AmznClientToken *string `location:"header" locationName:"X-Amzn-Client-Token" type:"string"`

	ConnectorDefinitionVersionArn *string `type:"string"`

	CoreDefinitionVersionArn *string `type:"string"`

	DeviceDefinitionVersionArn *string `type:"string"`

	FunctionDefinitionVersionArn *string `type:"string"`

	// GroupId is a required field
	GroupId *string `location:"uri" locationName:"GroupId" type:"string" required:"true"`

	LoggerDefinitionVersionArn *string `type:"string"`

	ResourceDefinitionVersionArn *string `type:"string"`

	SubscriptionDefinitionVersionArn *string `type:"string"`
}

// String returns the string representation
func (s CreateGroupVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateGroupVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateGroupVersionInput"}

	if s.GroupId == nil {
		invalidParams.Add(aws.NewErrParamRequired("GroupId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateGroupVersionOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `type:"string"`

	CreationTimestamp *string `type:"string"`

	Id *string `type:"string"`

	Version *string `type:"string"`
}

// String returns the string representation
func (s CreateGroupVersionOutput) String() string {
	return awsutil.Prettify(s)
}