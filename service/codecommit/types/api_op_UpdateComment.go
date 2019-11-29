// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateCommentInput struct {
	_ struct{} `type:"structure"`

	// The system-generated ID of the comment you want to update. To get this ID,
	// use GetCommentsForComparedCommit or GetCommentsForPullRequest.
	//
	// CommentId is a required field
	CommentId *string `locationName:"commentId" type:"string" required:"true"`

	// The updated content to replace the existing content of the comment.
	//
	// Content is a required field
	Content *string `locationName:"content" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateCommentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateCommentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateCommentInput"}

	if s.CommentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CommentId"))
	}

	if s.Content == nil {
		invalidParams.Add(aws.NewErrParamRequired("Content"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateCommentOutput struct {
	_ struct{} `type:"structure"`

	// Information about the updated comment.
	Comment *Comment `locationName:"comment" type:"structure"`
}

// String returns the string representation
func (s UpdateCommentOutput) String() string {
	return awsutil.Prettify(s)
}
