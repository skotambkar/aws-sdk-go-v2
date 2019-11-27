// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type GetBackupPlanFromTemplateInput struct {
	_ struct{} `type:"structure"`

	// Uniquely identifies a stored backup plan template.
	//
	// BackupPlanTemplateId is a required field
	BackupPlanTemplateId *string `location:"uri" locationName:"templateId" type:"string" required:"true"`
}

// String returns the string representation
func (s GetBackupPlanFromTemplateInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetBackupPlanFromTemplateInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "GetBackupPlanFromTemplateInput"}

	if s.BackupPlanTemplateId == nil {
		invalidParams.Add(aws.NewErrParamRequired("BackupPlanTemplateId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type GetBackupPlanFromTemplateOutput struct {
	_ struct{} `type:"structure"`

	// Returns the body of a backup plan based on the target template, including
	// the name, rules, and backup vault of the plan.
	BackupPlanDocument *BackupPlan `type:"structure"`
}

// String returns the string representation
func (s GetBackupPlanFromTemplateOutput) String() string {
	return awsutil.Prettify(s)
}
