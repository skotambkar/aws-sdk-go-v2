// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type CreateOpsItemInput struct {
	_ struct{} `type:"structure"`

	// Information about the OpsItem.
	//
	// Description is a required field
	Description *string `min:"1" type:"string" required:"true"`

	// The Amazon Resource Name (ARN) of an SNS topic where notifications are sent
	// when this OpsItem is edited or changed.
	Notifications []OpsItemNotification `type:"list"`

	// Operational data is custom data that provides useful reference details about
	// the OpsItem. For example, you can specify log files, error strings, license
	// keys, troubleshooting tips, or other relevant data. You enter operational
	// data as key-value pairs. The key has a maximum length of 128 characters.
	// The value has a maximum size of 20 KB.
	//
	// Operational data keys can't begin with the following: amazon, aws, amzn,
	// ssm, /amazon, /aws, /amzn, /ssm.
	//
	// You can choose to make the data searchable by other users in the account
	// or you can restrict search access. Searchable data means that all users with
	// access to the OpsItem Overview page (as provided by the DescribeOpsItems
	// API action) can view and search on the specified data. Operational data that
	// is not searchable is only viewable by users who have access to the OpsItem
	// (as provided by the GetOpsItem API action).
	//
	// Use the /aws/resources key in OperationalData to specify a related resource
	// in the request. Use the /aws/automations key in OperationalData to associate
	// an Automation runbook with the OpsItem. To view AWS CLI example commands
	// that use these keys, see Creating OpsItems Manually (http://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-creating-OpsItems.html#OpsCenter-manually-create-OpsItems)
	// in the AWS Systems Manager User Guide.
	OperationalData map[string]OpsItemDataValue `type:"map"`

	// The importance of this OpsItem in relation to other OpsItems in the system.
	Priority *int64 `min:"1" type:"integer"`

	// One or more OpsItems that share something in common with the current OpsItems.
	// For example, related OpsItems can include OpsItems with similar error messages,
	// impacted resources, or statuses for the impacted resource.
	RelatedOpsItems []RelatedOpsItem `type:"list"`

	// The origin of the OpsItem, such as Amazon EC2 or AWS Systems Manager.
	//
	// Source is a required field
	Source *string `min:"1" type:"string" required:"true"`

	// Optional metadata that you assign to a resource. You can restrict access
	// to OpsItems by using an inline IAM policy that specifies tags. For more information,
	// see Getting Started with OpsCenter (http://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-getting-started.html#OpsCenter-getting-started-user-permissions)
	// in the AWS Systems Manager User Guide.
	//
	// Tags use a key-value pair. For example:
	//
	// Key=Department,Value=Finance
	//
	// To add tags to an existing OpsItem, use the AddTagsToResource action.
	Tags []Tag `type:"list"`

	// A short heading that describes the nature of the OpsItem and the impacted
	// resource.
	//
	// Title is a required field
	Title *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateOpsItemInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateOpsItemInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateOpsItemInput"}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}
	if s.Priority != nil && *s.Priority < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Priority", 1))
	}

	if s.Source == nil {
		invalidParams.Add(aws.NewErrParamRequired("Source"))
	}
	if s.Source != nil && len(*s.Source) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Source", 1))
	}

	if s.Title == nil {
		invalidParams.Add(aws.NewErrParamRequired("Title"))
	}
	if s.Title != nil && len(*s.Title) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Title", 1))
	}
	if s.RelatedOpsItems != nil {
		for i, v := range s.RelatedOpsItems {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RelatedOpsItems", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.Tags != nil {
		for i, v := range s.Tags {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Tags", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateOpsItemOutput struct {
	_ struct{} `type:"structure"`

	// The ID of the OpsItem.
	OpsItemId *string `type:"string"`
}

// String returns the string representation
func (s CreateOpsItemOutput) String() string {
	return awsutil.Prettify(s)
}
