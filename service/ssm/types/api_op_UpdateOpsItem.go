// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ssm/enums"
)

type UpdateOpsItemInput struct {
	_ struct{} `type:"structure"`

	// Update the information about the OpsItem. Provide enough information so that
	// users reading this OpsItem for the first time understand the issue.
	Description *string `min:"1" type:"string"`

	// The Amazon Resource Name (ARN) of an SNS topic where notifications are sent
	// when this OpsItem is edited or changed.
	Notifications []OpsItemNotification `type:"list"`

	// Add new keys or edit existing key-value pairs of the OperationalData map
	// in the OpsItem object.
	//
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

	// Keys that you want to remove from the OperationalData map.
	OperationalDataToDelete []string `type:"list"`

	// The ID of the OpsItem.
	//
	// OpsItemId is a required field
	OpsItemId *string `type:"string" required:"true"`

	// The importance of this OpsItem in relation to other OpsItems in the system.
	Priority *int64 `min:"1" type:"integer"`

	// One or more OpsItems that share something in common with the current OpsItems.
	// For example, related OpsItems can include OpsItems with similar error messages,
	// impacted resources, or statuses for the impacted resource.
	RelatedOpsItems []RelatedOpsItem `type:"list"`

	// The OpsItem status. Status can be Open, In Progress, or Resolved. For more
	// information, see Editing OpsItem Details (http://docs.aws.amazon.com/systems-manager/latest/userguide/OpsCenter-working-with-OpsItems-editing-details.html)
	// in the AWS Systems Manager User Guide.
	Status enums.OpsItemStatus `type:"string" enum:"true"`

	// A short heading that describes the nature of the OpsItem and the impacted
	// resource.
	Title *string `min:"1" type:"string"`
}

// String returns the string representation
func (s UpdateOpsItemInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateOpsItemInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateOpsItemInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.OpsItemId == nil {
		invalidParams.Add(aws.NewErrParamRequired("OpsItemId"))
	}
	if s.Priority != nil && *s.Priority < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("Priority", 1))
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

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateOpsItemOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s UpdateOpsItemOutput) String() string {
	return awsutil.Prettify(s)
}
