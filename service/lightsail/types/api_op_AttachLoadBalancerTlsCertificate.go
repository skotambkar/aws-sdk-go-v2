// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AttachLoadBalancerTlsCertificateInput struct {
	_ struct{} `type:"structure"`

	// The name of your SSL/TLS certificate.
	//
	// CertificateName is a required field
	CertificateName *string `locationName:"certificateName" type:"string" required:"true"`

	// The name of the load balancer to which you want to associate the SSL/TLS
	// certificate.
	//
	// LoadBalancerName is a required field
	LoadBalancerName *string `locationName:"loadBalancerName" type:"string" required:"true"`
}

// String returns the string representation
func (s AttachLoadBalancerTlsCertificateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AttachLoadBalancerTlsCertificateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AttachLoadBalancerTlsCertificateInput"}

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

type AttachLoadBalancerTlsCertificateOutput struct {
	_ struct{} `type:"structure"`

	// An object representing the API operations.
	//
	// These SSL/TLS certificates are only usable by Lightsail load balancers. You
	// can't get the certificate and use it for another purpose.
	Operations []Operation `locationName:"operations" type:"list"`
}

// String returns the string representation
func (s AttachLoadBalancerTlsCertificateOutput) String() string {
	return awsutil.Prettify(s)
}