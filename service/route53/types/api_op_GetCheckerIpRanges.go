// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Empty request.
type GetCheckerIpRangesInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s GetCheckerIpRangesInput) String() string {
	return awsutil.Prettify(s)
}

// A complex type that contains the CheckerIpRanges element.
type GetCheckerIpRangesOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains sorted list of IP ranges in CIDR format for
	// Amazon Route 53 health checkers.
	//
	// CheckerIpRanges is a required field
	CheckerIpRanges []string `type:"list" required:"true"`
}

// String returns the string representation
func (s GetCheckerIpRangesOutput) String() string {
	return awsutil.Prettify(s)
}