// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A BatchMeterUsageRequest contains UsageRecords, which indicate quantities
// of usage within your application.
type BatchMeterUsageInput struct {
	_ struct{} `type:"structure"`

	// Product code is used to uniquely identify a product in AWS Marketplace. The
	// product code should be the same as the one used during the publishing of
	// a new product.
	//
	// ProductCode is a required field
	ProductCode *string `min:"1" type:"string" required:"true"`

	// The set of UsageRecords to submit. BatchMeterUsage accepts up to 25 UsageRecords
	// at a time.
	//
	// UsageRecords is a required field
	UsageRecords []UsageRecord `type:"list" required:"true"`
}

// String returns the string representation
func (s BatchMeterUsageInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *BatchMeterUsageInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "BatchMeterUsageInput"}

	if s.ProductCode == nil {
		invalidParams.Add(aws.NewErrParamRequired("ProductCode"))
	}
	if s.ProductCode != nil && len(*s.ProductCode) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ProductCode", 1))
	}

	if s.UsageRecords == nil {
		invalidParams.Add(aws.NewErrParamRequired("UsageRecords"))
	}
	if s.UsageRecords != nil {
		for i, v := range s.UsageRecords {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "UsageRecords", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the UsageRecords processed by BatchMeterUsage and any records that
// have failed due to transient error.
type BatchMeterUsageOutput struct {
	_ struct{} `type:"structure"`

	// Contains all UsageRecords processed by BatchMeterUsage. These records were
	// either honored by AWS Marketplace Metering Service or were invalid.
	Results []UsageRecordResult `type:"list"`

	// Contains all UsageRecords that were not processed by BatchMeterUsage. This
	// is a list of UsageRecords. You can retry the failed request by making another
	// BatchMeterUsage call with this list as input in the BatchMeterUsageRequest.
	UnprocessedRecords []UsageRecord `type:"list"`
}

// String returns the string representation
func (s BatchMeterUsageOutput) String() string {
	return awsutil.Prettify(s)
}