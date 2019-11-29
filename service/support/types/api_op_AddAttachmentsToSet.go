// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type AddAttachmentsToSetInput struct {
	_ struct{} `type:"structure"`

	// The ID of the attachment set. If an attachmentSetId is not specified, a new
	// attachment set is created, and the ID of the set is returned in the response.
	// If an attachmentSetId is specified, the attachments are added to the specified
	// set, if it exists.
	AttachmentSetId *string `locationName:"attachmentSetId" type:"string"`

	// One or more attachments to add to the set. The limit is 3 attachments per
	// set, and the size limit is 5 MB per attachment.
	//
	// Attachments is a required field
	Attachments []Attachment `locationName:"attachments" type:"list" required:"true"`
}

// String returns the string representation
func (s AddAttachmentsToSetInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *AddAttachmentsToSetInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "AddAttachmentsToSetInput"}

	if s.Attachments == nil {
		invalidParams.Add(aws.NewErrParamRequired("Attachments"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// The ID and expiry time of the attachment set returned by the AddAttachmentsToSet
// operation.
type AddAttachmentsToSetOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the attachment set. If an attachmentSetId was not specified, a
	// new attachment set is created, and the ID of the set is returned in the response.
	// If an attachmentSetId was specified, the attachments are added to the specified
	// set, if it exists.
	AttachmentSetId *string `locationName:"attachmentSetId" type:"string"`

	// The time and date when the attachment set expires.
	ExpiryTime *string `locationName:"expiryTime" type:"string"`
}

// String returns the string representation
func (s AddAttachmentsToSetOutput) String() string {
	return awsutil.Prettify(s)
}