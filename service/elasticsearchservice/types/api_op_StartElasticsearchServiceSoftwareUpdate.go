// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Container for the parameters to the StartElasticsearchServiceSoftwareUpdate
// operation. Specifies the name of the Elasticsearch domain that you wish to
// schedule a service software update on.
type StartElasticsearchServiceSoftwareUpdateInput struct {
	_ struct{} `type:"structure"`

	// The name of the domain that you want to update to the latest service software.
	//
	// DomainName is a required field
	DomainName *string `min:"3" type:"string" required:"true"`
}

// String returns the string representation
func (s StartElasticsearchServiceSoftwareUpdateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartElasticsearchServiceSoftwareUpdateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartElasticsearchServiceSoftwareUpdateInput"}

	if s.DomainName == nil {
		invalidParams.Add(aws.NewErrParamRequired("DomainName"))
	}
	if s.DomainName != nil && len(*s.DomainName) < 3 {
		invalidParams.Add(aws.NewErrParamMinLen("DomainName", 3))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The result of a StartElasticsearchServiceSoftwareUpdate operation. Contains
// the status of the update.
type StartElasticsearchServiceSoftwareUpdateOutput struct {
	_ struct{} `type:"structure"`

	// The current status of the Elasticsearch service software update.
	ServiceSoftwareOptions *ServiceSoftwareOptions `type:"structure"`
}

// String returns the string representation
func (s StartElasticsearchServiceSoftwareUpdateOutput) String() string {
	return awsutil.Prettify(s)
}
