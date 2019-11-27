// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type DescribePendingMaintenanceActionsInput struct {
	_ struct{} `type:"structure"`

	// A filter that specifies one or more resources to return pending maintenance
	// actions for.
	//
	// Supported filters:
	//
	//    * db-cluster-id - Accepts DB cluster identifiers and DB cluster Amazon
	//    Resource Names (ARNs). The results list will only include pending maintenance
	//    actions for the DB clusters identified by these ARNs.
	//
	//    * db-instance-id - Accepts DB instance identifiers and DB instance ARNs.
	//    The results list will only include pending maintenance actions for the
	//    DB instances identified by these ARNs.
	Filters []Filter `locationNameList:"Filter" type:"list"`

	// An optional pagination token provided by a previous DescribePendingMaintenanceActions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to a number of records specified by MaxRecords.
	Marker *string `type:"string"`

	// The maximum number of records to include in the response. If more records
	// exist than the specified MaxRecords value, a pagination token called a marker
	// is included in the response so that you can retrieve the remaining results.
	//
	// Default: 100
	//
	// Constraints: Minimum 20, maximum 100.
	MaxRecords *int64 `type:"integer"`

	// The ARN of a resource to return pending maintenance actions for.
	ResourceIdentifier *string `type:"string"`
}

// String returns the string representation
func (s DescribePendingMaintenanceActionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribePendingMaintenanceActionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribePendingMaintenanceActionsInput"}
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

// Data returned from the DescribePendingMaintenanceActions action.
type DescribePendingMaintenanceActionsOutput struct {
	_ struct{} `type:"structure"`

	// An optional pagination token provided by a previous DescribePendingMaintenanceActions
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to a number of records specified by MaxRecords.
	Marker *string `type:"string"`

	// A list of the pending maintenance actions for the resource.
	PendingMaintenanceActions []ResourcePendingMaintenanceActions `locationNameList:"ResourcePendingMaintenanceActions" type:"list"`
}

// String returns the string representation
func (s DescribePendingMaintenanceActionsOutput) String() string {
	return awsutil.Prettify(s)
}
