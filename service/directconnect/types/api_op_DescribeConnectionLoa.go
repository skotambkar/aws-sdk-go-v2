// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/enums"
)

type DescribeConnectionLoaInput struct {
	_ struct{} `type:"structure"`

	// The ID of the connection.
	//
	// ConnectionId is a required field
	ConnectionId *string `locationName:"connectionId" type:"string" required:"true"`

	// The standard media type for the LOA-CFA document. The only supported value
	// is application/pdf.
	LoaContentType enums.LoaContentType `locationName:"loaContentType" type:"string" enum:"true"`

	// The name of the APN partner or service provider who establishes connectivity
	// on your behalf. If you specify this parameter, the LOA-CFA lists the provider
	// name alongside your company name as the requester of the cross connect.
	ProviderName *string `locationName:"providerName" type:"string"`
}

// String returns the string representation
func (s DescribeConnectionLoaInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeConnectionLoaInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeConnectionLoaInput"}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DescribeConnectionLoaOutput struct {
	_ struct{} `type:"structure"`

	// The Letter of Authorization - Connecting Facility Assignment (LOA-CFA).
	Loa *Loa `locationName:"loa" type:"structure"`
}

// String returns the string representation
func (s DescribeConnectionLoaOutput) String() string {
	return awsutil.Prettify(s)
}