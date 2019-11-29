// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/iot/enums"
)

type ListThingRegistrationTaskReportsInput struct {
	_ struct{} `type:"structure"`

	// The maximum number of results to return per request.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" min:"1" type:"integer"`

	// The token to retrieve the next set of results.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`

	// The type of task report.
	//
	// ReportType is a required field
	ReportType enums.ReportType `location:"querystring" locationName:"reportType" type:"string" required:"true" enum:"true"`

	// The id of the task.
	//
	// TaskId is a required field
	TaskId *string `location:"uri" locationName:"taskId" type:"string" required:"true"`
}

// String returns the string representation
func (s ListThingRegistrationTaskReportsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ListThingRegistrationTaskReportsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ListThingRegistrationTaskReportsInput"}
	if s.MaxResults != nil && *s.MaxResults < 1 {
		invalidParams.Add(aws.NewErrParamMinValue("MaxResults", 1))
	}
	if len(s.ReportType) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("ReportType"))
	}

	if s.TaskId == nil {
		invalidParams.Add(aws.NewErrParamRequired("TaskId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ListThingRegistrationTaskReportsOutput struct {
	_ struct{} `type:"structure"`

	// The token used to get the next set of results, or null if there are no additional
	// results.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The type of task report.
	ReportType enums.ReportType `locationName:"reportType" type:"string" enum:"true"`

	// Links to the task resources.
	ResourceLinks []string `locationName:"resourceLinks" type:"list"`
}

// String returns the string representation
func (s ListThingRegistrationTaskReportsOutput) String() string {
	return awsutil.Prettify(s)
}
