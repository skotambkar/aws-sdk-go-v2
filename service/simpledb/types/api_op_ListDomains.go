// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListDomainsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of domain names you want returned. The range is 1 to 100.
	// The default setting is 100.
	MaxNumberOfDomains *int64 `type:"integer"`

	// A string informing Amazon SimpleDB where to start the next list of domain
	// names.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDomainsInput) String() string {
	return awsutil.Prettify(s)
}

type ListDomainsOutput struct {
	_ struct{} `type:"structure"`

	// A list of domain names that match the expression.
	DomainNames []string `locationNameList:"DomainName" type:"list" flattened:"true"`

	// An opaque token indicating that there are more domains than the specified
	//    MaxNumberOfDomains
	//  still available.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListDomainsOutput) String() string {
	return awsutil.Prettify(s)
}
