// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the parameters for SetLoadBalancerListenerSSLCertificate.
type SetLoadBalancerListenerSSLCertificateInput struct {
	_ struct{} `type:"structure"`

	// The name of the load balancer.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `type:"string" required:"true"`

	// The port that uses the specified SSL certificate.
	//
	// LoadBalancerPort is a required field
	LoadBalancerPort *int64 `type:"integer" required:"true"`

	// The Amazon Resource Name (ARN) of the SSL certificate.
	//
	// SSLCertificateId is a required field
	SSLCertificateId *string `type:"string" required:"true"`
}

// String returns the string representation
func (s SetLoadBalancerListenerSSLCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SetLoadBalancerListenerSSLCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SetLoadBalancerListenerSSLCertificateInput"}

	if s.LoadBalancerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerName"))
	}

	if s.LoadBalancerPort == nil {
		invalidParams.Add(aws.NewErrParamRequired("LoadBalancerPort"))
	}

	if s.SSLCertificateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("SSLCertificateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the output of SetLoadBalancerListenerSSLCertificate.
type SetLoadBalancerListenerSSLCertificateOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s SetLoadBalancerListenerSSLCertificateOutput) String() string {
	return awsutil.Prettify(s)
}