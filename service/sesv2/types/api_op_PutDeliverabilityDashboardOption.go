// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Enable or disable the Deliverability dashboard. When you enable the Deliverability
// dashboard, you gain access to reputation, deliverability, and other metrics
// for the domains that you use to send email using Amazon SES API v2. You also
// gain the ability to perform predictive inbox placement tests.
//
// When you use the Deliverability dashboard, you pay a monthly subscription
// charge, in addition to any other fees that you accrue by using Amazon SES
// and other AWS services. For more information about the features and cost
// of a Deliverability dashboard subscription, see Amazon Pinpoint Pricing (http://aws.amazon.com/pinpoint/pricing/).
type PutDeliverabilityDashboardOptionInput struct {
	_ struct{} `type:"structure"`

	// Specifies whether to enable the Deliverability dashboard. To enable the dashboard,
	// set this value to true.
	//
	// DashboardEnabled is a required field
	DashboardEnabled *bool `type:"boolean" required:"true"`

	// An array of objects, one for each verified domain that you use to send email
	// and enabled the Deliverability dashboard for.
	SubscribedDomains []DomainDeliverabilityTrackingOption `type:"list"`
}

// String returns the string representation
func (s PutDeliverabilityDashboardOptionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutDeliverabilityDashboardOptionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutDeliverabilityDashboardOptionInput"}

	if s.DashboardEnabled == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashboardEnabled"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A response that indicates whether the Deliverability dashboard is enabled.
type PutDeliverabilityDashboardOptionOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutDeliverabilityDashboardOptionOutput) String() string {
	return awsutil.Prettify(s)
}
