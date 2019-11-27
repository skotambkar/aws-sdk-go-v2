// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ListBackupPlanTemplatesInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of items to be returned.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListBackupPlanTemplatesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListBackupPlanTemplatesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListBackupPlanTemplatesInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListBackupPlanTemplatesOutput struct {
	_ struct{} `type:"structure"`

	// An array of template list items containing metadata about your saved templates.
	BackupPlanTemplatesList []BackupPlanTemplatesListMember `type:"list"`

	// The next item following a partial list of returned items. For example, if
	// a request is made to return maxResults number of items, NextToken allows
	// you to return more items in your list starting at the location pointed to
	// by the next token.
	NextToken *string `type:"string"`
}

// String returns the string representation
func (s ListBackupPlanTemplatesOutput) String() string {
	return awsutil.Prettify(s)
}
