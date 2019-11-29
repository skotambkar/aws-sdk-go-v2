// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// A request to update the comment for a hosted zone.
type UpdateHostedZoneCommentInput struct {
	_ struct{} `locationName:"UpdateHostedZoneCommentRequest" type:"structure" xmlURI:"https://route53.amazonaws.com/doc/2013-04-01/"`

	// The new comment for the hosted zone. If you don't specify a value for Comment,
	// Amazon Route 53 deletes the existing value of the Comment element, if any.
	Comment *string `type:"string"`

	// The ID for the hosted zone that you want to update the comment for.
	//
	// Id is a required field
	Id *string `location:"uri" locationName:"Id" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateHostedZoneCommentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateHostedZoneCommentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateHostedZoneCommentInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// A complex type that contains the response to the UpdateHostedZoneComment
// request.
type UpdateHostedZoneCommentOutput struct {
	_ struct{} `type:"structure"`

	// A complex type that contains the response to the UpdateHostedZoneComment
	// request.
	//
	// HostedZone is a required field
	HostedZone *HostedZone `type:"structure" required:"true"`
}

// String returns the string representation
func (s UpdateHostedZoneCommentOutput) String() string {
	return awsutil.Prettify(s)
}
