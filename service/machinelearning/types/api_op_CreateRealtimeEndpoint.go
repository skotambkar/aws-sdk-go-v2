// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateRealtimeEndpointInput struct {
	_ struct{} `type:"structure"`

	// The ID assigned to the MLModel during creation.
	//
	// MLModelId is a required field
	MLModelId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateRealtimeEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateRealtimeEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateRealtimeEndpointInput"}

	if s.MLModelId == nil {
		invalidParams.Add(aws.NewErrParamRequired("MLModelId"))
	}
	if s.MLModelId != nil && len(*s.MLModelId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("MLModelId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Represents the output of an CreateRealtimeEndpoint operation.
//
// The result contains the MLModelId and the endpoint information for the MLModel.
//
// The endpoint information includes the URI of the MLModel; that is, the location
// to send online prediction requests for the specified MLModel.
type CreateRealtimeEndpointOutput struct {
	_ struct{} `type:"structure"`

	// A user-supplied ID that uniquely identifies the MLModel. This value should
	// be identical to the value of the MLModelId in the request.
	MLModelId *string `min:"1" type:"string"`

	// The endpoint information of the MLModel
	RealtimeEndpointInfo *RealtimeEndpointInfo `type:"structure"`
}

// String returns the string representation
func (s CreateRealtimeEndpointOutput) String() string {
	return awsutil.Prettify(s)
}
