// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PostToConnectionInput struct {
	_ struct{} `type:"structure" payload:"Data"`

	// ConnectionId is a required field
	ConnectionId *string `location:"uri" locationName:"connectionId" type:"string" required:"true"`

	// The data to be sent to the client specified by its connection id.
	//
	// Data is a required field
	Data []byte `type:"blob" required:"true"`
}

// String returns the string representation
func (s PostToConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PostToConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PostToConnectionInput"}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

	if s.Data == nil {
		invalidParams.Add(aws.NewErrParamRequired("Data"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PostToConnectionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PostToConnectionOutput) String() string {
	return awsutil.Prettify(s)
}
