// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request for the number of health checks that are associated with the current
// AWS account.
type GetHealthCheckCountInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetHealthCheckCountInput) String() string {
	return awsutil.Prettify(s)
}

// A complex type that contains the response to a GetHealthCheckCount request.
type GetHealthCheckCountOutput struct {
	_ struct{} `type:"structure"`

	// The number of health checks associated with the current AWS account.
	//
	// HealthCheckCount is a required field
	HealthCheckCount *int64 `type:"long" required:"true"`
}

// String returns the string representation
func (s GetHealthCheckCountOutput) String() string {
	return awsutil.Prettify(s)
}