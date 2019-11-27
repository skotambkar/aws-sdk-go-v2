// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeLagsInput struct {
	_ struct{} `type:"structure"`

	// The ID of the LAG.
	LagId *string `locationName:"lagId" type:"string"`
}

// String returns the string representation
func (s DescribeLagsInput) String() string {
	return awsutil.Prettify(s)
}

type DescribeLagsOutput struct {
	_ struct{} `type:"structure"`

	// The LAGs.
	Lags []Lag `locationName:"lags" type:"list"`
}

// String returns the string representation
func (s DescribeLagsOutput) String() string {
	return awsutil.Prettify(s)
}
