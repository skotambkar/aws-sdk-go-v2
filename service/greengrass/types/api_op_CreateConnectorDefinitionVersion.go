// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateConnectorDefinitionVersionInput struct {
	_ struct{} `type:"structure"`

	AmznClientToken *string `location:"header" locationName:"X-Amzn-Client-Token" type:"string"`

	// ConnectorDefinitionId is a required field
	ConnectorDefinitionId *string `location:"uri" locationName:"ConnectorDefinitionId" type:"string" required:"true"`

	Connectors []Connector `type:"list"`
}

// String returns the string representation
func (s CreateConnectorDefinitionVersionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateConnectorDefinitionVersionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateConnectorDefinitionVersionInput"}

	if s.ConnectorDefinitionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectorDefinitionId"))
	}
	if s.Connectors != nil {
		for i, v := range s.Connectors {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Connectors", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateConnectorDefinitionVersionOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `type:"string"`

	CreationTimestamp *string `type:"string"`

	Id *string `type:"string"`

	Version *string `type:"string"`
}

// String returns the string representation
func (s CreateConnectorDefinitionVersionOutput) String() string {
	return awsutil.Prettify(s)
}
