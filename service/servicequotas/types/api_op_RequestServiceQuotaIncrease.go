// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type RequestServiceQuotaIncreaseInput struct {
	_ struct{} `type:"structure"`

	// Specifies the value submitted in the service quota increase request.
	//
	// DesiredValue is a required field
	DesiredValue *float64 `type:"double" required:"true"`

	// Specifies the service quota that you want to use.
	//
	// QuotaCode is a required field
	QuotaCode *string `min:"1" type:"string" required:"true"`

	// Specifies the service that you want to use.
	//
	// ServiceCode is a required field
	ServiceCode *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s RequestServiceQuotaIncreaseInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *RequestServiceQuotaIncreaseInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "RequestServiceQuotaIncreaseInput"}

	if s.DesiredValue == nil {
		invalidParams.Add(aws.NewErrParamRequired("DesiredValue"))
	}

	if s.QuotaCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("QuotaCode"))
	}
	if s.QuotaCode != nil && len(*s.QuotaCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("QuotaCode", 1))
	}

	if s.ServiceCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ServiceCode"))
	}
	if s.ServiceCode != nil && len(*s.ServiceCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ServiceCode", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type RequestServiceQuotaIncreaseOutput struct {
	_ struct{} `type:"structure"`

	// Returns a list of service quota requests.
	RequestedQuota *RequestedServiceQuotaChange `type:"structure"`
}

// String returns the string representation
func (s RequestServiceQuotaIncreaseOutput) String() string {
	return awsutil.Prettify(s)
}
