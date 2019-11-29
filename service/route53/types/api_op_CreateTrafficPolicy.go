// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A complex type that contains information about the traffic policy that you
// want to create.
type CreateTrafficPolicyInput struct {
	_ struct{} `locationName:"CreateTrafficPolicyRequest" type:"structure" xmlURI:"https://route53.amazonaws.com/doc/2013-04-01/"`

	// (Optional) Any comments that you want to include about the traffic policy.
	Comment *string `type:"string"`

	// The definition of this traffic policy in JSON format. For more information,
	// see Traffic Policy Document Format (https://docs.aws.amazon.com/Route53/latest/APIReference/api-policies-traffic-policy-document-format.html).
	//
	// Document is a required field
	Document *string `type:"string" required:"true"`

	// The name of the traffic policy.
	//
	// Name is a required field
	Name *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateTrafficPolicyInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateTrafficPolicyInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateTrafficPolicyInput"}

	if s.Document == nil {
		invalidParams.Add(aws.NewErrParamRequired("Document"))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A complex type that contains the response information for the CreateTrafficPolicy
// request.
type CreateTrafficPolicyOutput struct {
	_ struct{} `type:"structure"`

	// A unique URL that represents a new traffic policy.
	//
	// Location is a required field
	Location *string `location:"header" locationName:"Location" type:"string" required:"true"`

	// A complex type that contains settings for the new traffic policy.
	//
	// TrafficPolicy is a required field
	TrafficPolicy *TrafficPolicy `type:"structure" required:"true"`
}

// String returns the string representation
func (s CreateTrafficPolicyOutput) String() string {
	return awsutil.Prettify(s)
}