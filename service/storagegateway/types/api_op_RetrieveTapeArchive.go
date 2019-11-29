// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// RetrieveTapeArchiveInput
type RetrieveTapeArchiveInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway you want to retrieve the virtual
	// tape to. Use the ListGateways operation to return a list of gateways for
	// your account and AWS Region.
	//
	// You retrieve archived virtual tapes to only one gateway and the gateway must
	// be a tape gateway.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of the virtual tape you want to retrieve from
	// the virtual tape shelf (VTS).
	//
	// TapeARN is a required field
	TapeARN *string `min:"50" type:"string" required:"true"`
}

// String returns the string representation
func (s RetrieveTapeArchiveInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RetrieveTapeArchiveInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RetrieveTapeArchiveInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}

	if s.TapeARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("TapeARN"))
	}
	if s.TapeARN != nil && len(*s.TapeARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("TapeARN", 50))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// RetrieveTapeArchiveOutput
type RetrieveTapeArchiveOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the retrieved virtual tape.
	TapeARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s RetrieveTapeArchiveOutput) String() string {
	return awsutil.Prettify(s)
}