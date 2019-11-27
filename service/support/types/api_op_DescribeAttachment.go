// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribeAttachmentInput struct {
	_ struct{} `type:"structure"`

	// The ID of the attachment to return. Attachment IDs are returned by the DescribeCommunications
	// operation.
	//
	// AttachmentId is a required field
	AttachmentId *string `locationName:"attachmentId" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeAttachmentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeAttachmentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeAttachmentInput"}

	if s.AttachmentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AttachmentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The content and file name of the attachment returned by the DescribeAttachment
// operation.
type DescribeAttachmentOutput struct {
	_ struct{} `type:"structure"`

	// The attachment content and file name.
	Attachment *Attachment `locationName:"attachment" type:"structure"`
}

// String returns the string representation
func (s DescribeAttachmentOutput) String() string {
	return awsutil.Prettify(s)
}
