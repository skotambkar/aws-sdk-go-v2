// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListBucketsInput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ListBucketsInput) String() string {
	return awsutil.Prettify(s)
}

type ListBucketsOutput struct {
	_ struct{} `type:"structure"`

	// The list of buckets owned by the requestor.
	Buckets []Bucket `locationNameList:"Bucket" type:"list"`

	// The owner of the buckets listed.
	Owner *Owner `type:"structure"`
}

// String returns the string representation
func (s ListBucketsOutput) String() string {
	return awsutil.Prettify(s)
}
