// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a DescribeLimits operation. Has no content.
type DescribeLimitsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeLimitsInput) String() string {
	return awsutil.Prettify(s)
}

// Represents the output of a DescribeLimits operation.
type DescribeLimitsOutput struct {
	_ struct{} `type:"structure"`

	// The maximum total read capacity units that your account allows you to provision
	// across all of your tables in this Region.
	AccountMaxReadCapacityUnits *int64 `min:"1" type:"long"`

	// The maximum total write capacity units that your account allows you to provision
	// across all of your tables in this Region.
	AccountMaxWriteCapacityUnits *int64 `min:"1" type:"long"`

	// The maximum read capacity units that your account allows you to provision
	// for a new table that you are creating in this Region, including the read
	// capacity units provisioned for its global secondary indexes (GSIs).
	TableMaxReadCapacityUnits *int64 `min:"1" type:"long"`

	// The maximum write capacity units that your account allows you to provision
	// for a new table that you are creating in this Region, including the write
	// capacity units provisioned for its global secondary indexes (GSIs).
	TableMaxWriteCapacityUnits *int64 `min:"1" type:"long"`
}

// String returns the string representation
func (s DescribeLimitsOutput) String() string {
	return awsutil.Prettify(s)
}
