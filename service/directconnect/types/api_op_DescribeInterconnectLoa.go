// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/enums"
)

type DescribeInterconnectLoaInput struct {
	_ struct{} `type:"structure"`

	// The ID of the interconnect.
	//
	// InterconnectId is a required field
	InterconnectId *string `locationName:"interconnectId" type:"string" required:"true"`

	// The standard media type for the LOA-CFA document. The only supported value
	// is application/pdf.
	LoaContentType enums.LoaContentType `locationName:"loaContentType" type:"string" enum:"true"`

	// The name of the service provider who establishes connectivity on your behalf.
	// If you supply this parameter, the LOA-CFA lists the provider name alongside
	// your company name as the requester of the cross connect.
	ProviderName *string `locationName:"providerName" type:"string"`
}

// String returns the string representation
func (s DescribeInterconnectLoaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeInterconnectLoaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeInterconnectLoaInput"}

	if s.InterconnectId == nil {
		invalidParams.Add(aws.NewErrParamRequired("InterconnectId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeInterconnectLoaOutput struct {
	_ struct{} `type:"structure"`

	// The Letter of Authorization - Connecting Facility Assignment (LOA-CFA).
	Loa *Loa `locationName:"loa" type:"structure"`
}

// String returns the string representation
func (s DescribeInterconnectLoaOutput) String() string {
	return awsutil.Prettify(s)
}
