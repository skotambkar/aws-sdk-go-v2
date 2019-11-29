// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/applicationdiscoveryservice/enums"
)

type StartExportTaskInput struct {
	_ struct{} `type:"structure"`

	// The end timestamp for exported data from the single Application Discovery
	// Agent selected in the filters. If no value is specified, exported data includes
	// the most recent data collected by the agent.
	EndTime *time.Time `locationName:"endTime" type:"timestamp"`

	// The file format for the returned export data. Default value is CSV. Note:
	// The GRAPHML option has been deprecated.
	ExportDataFormat []enums.ExportDataFormat `locationName:"exportDataFormat" type:"list"`

	// If a filter is present, it selects the single agentId of the Application
	// Discovery Agent for which data is exported. The agentId can be found in the
	// results of the DescribeAgents API or CLI. If no filter is present, startTime
	// and endTime are ignored and exported data includes both Agentless Discovery
	// Connector data and summary data from Application Discovery agents.
	Filters []ExportFilter `locationName:"filters" type:"list"`

	// The start timestamp for exported data from the single Application Discovery
	// Agent selected in the filters. If no value is specified, data is exported
	// starting from the first data collected by the agent.
	StartTime *time.Time `locationName:"startTime" type:"timestamp"`
}

// String returns the string representation
func (s StartExportTaskInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *StartExportTaskInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "StartExportTaskInput"}
	if s.Filters != nil {
		for i, v := range s.Filters {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "Filters", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type StartExportTaskOutput struct {
	_ struct{} `type:"structure"`

	// A unique identifier used to query the status of an export request.
	ExportId *string `locationName:"exportId" type:"string"`
}

// String returns the string representation
func (s StartExportTaskOutput) String() string {
	return awsutil.Prettify(s)
}
