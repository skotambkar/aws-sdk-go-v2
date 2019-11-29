// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/enums"
)

type CreateDashboardInput struct {
	_ struct{} `type:"structure"`

	// AWS account ID where you want to create the dashboard.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the dashboard, also added to IAM policy.
	//
	// DashboardId is a required field
	DashboardId *string `location:"uri" locationName:"DashboardId" min:"1" type:"string" required:"true"`

	// Publishing options when creating dashboard.
	//
	//    * AvailabilityStatus for AdHocFilteringOption - This can be either ENABLED
	//    or DISABLED. When This is set to set to DISABLED, QuickSight disables
	//    the left filter pane on the published dashboard, which can be used for
	//    AdHoc filtering. Enabled by default.
	//
	//    * AvailabilityStatus for ExportToCSVOption - This can be either ENABLED
	//    or DISABLED. The visual option to export data to CSV is disabled when
	//    this is set to DISABLED. Enabled by default.
	//
	//    * VisibilityState for SheetControlsOption - This can be either COLLAPSED
	//    or EXPANDED. The sheet controls pane is collapsed by default when set
	//    to true. Collapsed by default.
	//
	// Shorthand Syntax:
	//
	// AdHocFilteringDisabled=boolean,ExportToCSVDisabled=boolean,SheetControlsCollapsed=boolean
	DashboardPublishOptions *DashboardPublishOptions `type:"structure"`

	// The display name of the dashboard.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// A structure that contains the parameters of the dashboard. These are parameter
	// overrides for a dashboard. A dashboard can have any type of parameters and
	// some parameters might accept multiple values. You could use the following
	// structure to override two string parameters that accept multiple values:
	Parameters *Parameters `type:"structure"`

	// A structure that contains the permissions of the dashboard. You can use this
	// for granting permissions with principal and action information.
	Permissions []ResourcePermission `min:"1" type:"list"`

	// Source entity from which the dashboard is created. The souce entity accepts
	// the ARN of the source template or analysis and also references the replacement
	// datasets for the placeholders set when creating the template. The replacement
	// datasets need to follow the same schema as the datasets for which placeholders
	// were created when creating the template.
	//
	// If you are creating a dashboard from a source entity in a different AWS account,
	// use the ARN of the source template.
	//
	// SourceEntity is a required field
	SourceEntity *DashboardSourceEntity `type:"structure" required:"true"`

	// Contains a map of the key-value pairs for the resource tag or tags assigned
	// to the dashboard.
	Tags []Tag `min:"1" type:"list"`

	// A description for the first version of the dashboard being created.
	VersionDescription *string `min:"1" type:"string"`
}

// String returns the string representation
func (s CreateDashboardInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateDashboardInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateDashboardInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DashboardId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DashboardId"))
	}
	if s.DashboardId != nil && len(*s.DashboardId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DashboardId", 1))
	}

	if s.Name == nil {
		invalidParams.Add(aws.NewErrParamRequired("Name"))
	}
	if s.Name != nil && len(*s.Name) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Name", 1))
	}
	if s.Permissions != nil && len(s.Permissions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Permissions", 1))
	}

	if s.SourceEntity == nil {
		invalidParams.Add(aws.NewErrParamRequired("SourceEntity"))
	}
	if s.Tags != nil && len(s.Tags) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Tags", 1))
	}
	if s.VersionDescription != nil && len(*s.VersionDescription) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("VersionDescription", 1))
	}
	if s.Parameters != nil {
		if err := s.Parameters.Validate(); err != nil {
			invalidParams.AddNested("Parameters", err.(aws.ErrInvalidParams))
		}
	}
	if s.Permissions != nil {
		for i, v := range s.Permissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Permissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.SourceEntity != nil {
		if err := s.SourceEntity.Validate(); err != nil {
			invalidParams.AddNested("SourceEntity", err.(aws.ErrInvalidParams))
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

type CreateDashboardOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the dashboard.
	Arn *string `type:"string"`

	// The creation status of the dashboard create request.
	CreationStatus enums.ResourceStatus `type:"string" enum:"true"`

	// The ID for the dashboard.
	DashboardId *string `min:"1" type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// The ARN of the dashboard, including the version number of the first version
	// that is created.
	VersionArn *string `type:"string"`
}

// String returns the string representation
func (s CreateDashboardOutput) String() string {
	return awsutil.Prettify(s)
}