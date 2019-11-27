// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetCommentInput struct {
	_ struct{} `type:"structure"`

	// The unique, system-generated ID of the comment. To get this ID, use GetCommentsForComparedCommit
	// or GetCommentsForPullRequest.
	//
	// CommentId is a required field
	CommentId *string `locationName:"commentId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetCommentInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetCommentInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetCommentInput"}

	if s.CommentId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CommentId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetCommentOutput struct {
	_ struct{} `type:"structure"`

	// The contents of the comment.
	Comment *Comment `locationName:"comment" type:"structure"`
}

// String returns the string representation
func (s GetCommentOutput) String() string {
	return awsutil.Prettify(s)
}
