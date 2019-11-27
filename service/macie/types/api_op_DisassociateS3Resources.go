// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DisassociateS3ResourcesInput struct {
	_ struct{} `type:"structure"`

	// The S3 resources (buckets or prefixes) that you want to remove from being
	// monitored and classified by Amazon Macie.
	//
	// AssociatedS3Resources is a required field
	AssociatedS3Resources []S3Resource `locationName:"associatedS3Resources" type:"list" required:"true"`

	// The ID of the Amazon Macie member account whose resources you want to remove
	// from being monitored by Amazon Macie.
	MemberAccountId *string `locationName:"memberAccountId" type:"string"`
}

// String returns the string representation
func (s DisassociateS3ResourcesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DisassociateS3ResourcesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DisassociateS3ResourcesInput"}

	if s.AssociatedS3Resources == nil {
		invalidParams.Add(aws.NewErrParamRequired("AssociatedS3Resources"))
	}
	if s.AssociatedS3Resources != nil {
		for i, v := range s.AssociatedS3Resources {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "AssociatedS3Resources", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type DisassociateS3ResourcesOutput struct {
	_ struct{} `type:"structure"`

	// S3 resources that couldn't be removed from being monitored and classified
	// by Amazon Macie. An error code and an error message are provided for each
	// failed item.
	FailedS3Resources []FailedS3Resource `locationName:"failedS3Resources" type:"list"`
}

// String returns the string representation
func (s DisassociateS3ResourcesOutput) String() string {
	return awsutil.Prettify(s)
}
