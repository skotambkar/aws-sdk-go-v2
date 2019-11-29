// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to get information about a specified health check.
type GetHealthCheckInput struct {
	_ struct{} `type:"structure"`

	// The identifier that Amazon Route 53 assigned to the health check when you
	// created it. When you add or update a resource record set, you use this value
	// to specify which health check to use. The value can be up to 64 characters
	// long.
	//
	// HealthCheckId is a required field
	HealthCheckId *string `location:"uri" locationName:"HealthCheckId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetHealthCheckInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetHealthCheckInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetHealthCheckInput"}

	if s.HealthCheckId == nil {
		invalidParams.Add(aws.NewErrParamRequired("HealthCheckId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A complex type that contains the response to a GetHealthCheck request.
type GetHealthCheckOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains information about one health check that is associated
	// with the current AWS account.
	//
	// HealthCheck is a required field
	HealthCheck *HealthCheck `type:"structure" required:"true"`
}

// String returns the string representation
func (s GetHealthCheckOutput) String() string {
	return awsutil.Prettify(s)
}
