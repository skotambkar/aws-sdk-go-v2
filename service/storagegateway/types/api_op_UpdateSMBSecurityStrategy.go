// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/storagegateway/enums"
)

type UpdateSMBSecurityStrategyInput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	//
	// GatewayARN is a required field
	GatewayARN *string `min:"50" type:"string" required:"true"`

	// Specifies the type of security strategy.
	//
	// ClientSpecified: if you use this option, requests are established based on
	// what is negotiated by the client. This option is recommended when you want
	// to maximize compatibility across different clients in your environment.
	//
	// MandatorySigning: if you use this option, file gateway only allows connections
	// from SMBv2 or SMBv3 clients that have signing enabled. This option works
	// with SMB clients on Microsoft Windows Vista, Windows Server 2008 or newer.
	//
	// MandatoryEncryption: if you use this option, file gateway only allows connections
	// from SMBv3 clients that have encryption enabled. This option is highly recommended
	// for environments that handle sensitive data. This option works with SMB clients
	// on Microsoft Windows 8, Windows Server 2012 or newer.
	//
	// SMBSecurityStrategy is a required field
	SMBSecurityStrategy enums.SMBSecurityStrategy `type:"string" required:"true" enum:"true"`
}

// String returns the string representation
func (s UpdateSMBSecurityStrategyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateSMBSecurityStrategyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateSMBSecurityStrategyInput"}

	if s.GatewayARN == nil {
		invalidParams.Add(aws.NewErrParamRequired("GatewayARN"))
	}
	if s.GatewayARN != nil && len(*s.GatewayARN) < 50 {
		invalidParams.Add(aws.NewErrParamMinLen("GatewayARN", 50))
	}
	if len(s.SMBSecurityStrategy) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("SMBSecurityStrategy"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateSMBSecurityStrategyOutput struct {
	_ struct{} `type:"structure"`

	// The Amazon Resource Name (ARN) of the gateway. Use the ListGateways operation
	// to return a list of gateways for your account and AWS Region.
	GatewayARN *string `min:"50" type:"string"`
}

// String returns the string representation
func (s UpdateSMBSecurityStrategyOutput) String() string {
	return awsutil.Prettify(s)
}