// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Updates a conditional forwarder.
type UpdateConditionalForwarderInput struct {
	_ struct{} `type:"structure"`

	// The directory ID of the AWS directory for which to update the conditional
	// forwarder.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`

	// The updated IP addresses of the remote DNS server associated with the conditional
	// forwarder.
	//
	// DnsIpAddrs is a required field
	DnsIpAddrs []string `type:"list" required:"true"`

	// The fully qualified domain name (FQDN) of the remote domain with which you
	// will set up a trust relationship.
	//
	// RemoteDomainName is a required field
	RemoteDomainName *string `type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateConditionalForwarderInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateConditionalForwarderInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateConditionalForwarderInput"}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}

	if s.DnsIpAddrs == nil {
		invalidParams.Add(aws.NewErrParamRequired("DnsIpAddrs"))
	}

	if s.RemoteDomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("RemoteDomainName"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of an UpdateConditionalForwarder request.
type UpdateConditionalForwarderOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateConditionalForwarderOutput) String() string {
	return awsutil.Prettify(s)
}