// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Request structure used to request projects list in AWS Mobile Hub.
type ListProjectsInput struct {
	_ struct{} `type:"structure"`

	// Maximum number of records to list in a single response.
	MaxResults *int64 `location:"querystring" locationName:"maxResults" type:"integer"`

	// Pagination token. Set to null to start listing projects from start. If non-null
	// pagination token is returned in a result, then pass its value in here in
	// another request to list more projects.
	NextToken *string `location:"querystring" locationName:"nextToken" type:"string"`
}

// String returns the string representation
func (s ListProjectsInput) String() string {
	return awsutil.Prettify(s)
}

// Result structure used for requests to list projects in AWS Mobile Hub.
type ListProjectsOutput struct {
	_ struct{} `type:"structure"`

	// Pagination token. Set to null to start listing records from start. If non-null
	// pagination token is returned in a result, then pass its value in here in
	// another request to list more entries.
	NextToken *string `locationName:"nextToken" type:"string"`

	// List of projects.
	Projects []ProjectSummary `locationName:"projects" type:"list"`
}

// String returns the string representation
func (s ListProjectsOutput) String() string {
	return awsutil.Prettify(s)
}