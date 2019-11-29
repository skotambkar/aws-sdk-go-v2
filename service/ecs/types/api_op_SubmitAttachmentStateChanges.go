// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type SubmitAttachmentStateChangesInput struct {
	_ struct{} `type:"structure"`

	// Any attachments associated with the state change request.
	//
	// Attachments is a required field
	Attachments []AttachmentStateChange `locationName:"attachments" type:"list" required:"true"`

	// The short name or full ARN of the cluster that hosts the container instance
	// the attachment belongs to.
	Cluster *string `locationName:"cluster" type:"string"`
}

// String returns the string representation
func (s SubmitAttachmentStateChangesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *SubmitAttachmentStateChangesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "SubmitAttachmentStateChangesInput"}

	if s.Attachments == nil {
		invalidParams.Add(aws.NewErrParamRequired("Attachments"))
	}
	if s.Attachments != nil {
		for i, v := range s.Attachments {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Attachments", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type SubmitAttachmentStateChangesOutput struct {
	_ struct{} `type:"structure"`

	// Acknowledgement of the state change.
	Acknowledgment *string `locationName:"acknowledgment" type:"string"`
}

// String returns the string representation
func (s SubmitAttachmentStateChangesOutput) String() string {
	return awsutil.Prettify(s)
}