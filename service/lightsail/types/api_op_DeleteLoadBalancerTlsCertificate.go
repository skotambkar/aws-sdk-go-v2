// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DeleteLoadBalancerTlsCertificateInput struct {
	_ struct{} `type:"structure"`

	// The SSL/TLS certificate name.
	//
	// CertificateName is a required field
	CertificateName *string `locationName:"certificateName" type:"string" required:"true"`

	// When true, forces the deletion of an SSL/TLS certificate.
	//
	// There can be two certificates associated with a Lightsail load balancer:
	// the primary and the backup. The force parameter is required when the primary
	// SSL/TLS certificate is in use by an instance attached to the load balancer.
	Force *bool `locationName:"force" type:"boolean"`

	// The load balancer name.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `locationName:"loadBalancerName" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteLoadBalancerTlsCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteLoadBalancerTlsCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DeleteLoadBalancerTlsCertificateInput"}

	if s.CertificateName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CertificateName"))
	}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DeleteLoadBalancerTlsCertificateOutput struct {
	_ struct{} `type:"structure"`

	// An object describing the API operations.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s DeleteLoadBalancerTlsCertificateOutput) String() string {
	return awsutil.Prettify(s)
}
