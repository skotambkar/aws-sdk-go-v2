// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/ecs/enums"
)

type ListTasksInput struct {
	_ struct{} `type:"structure"`

	// The short name or full Amazon Resource Name (ARN) of the cluster that hosts
	// the tasks to list. If you do not specify a cluster, the default cluster is
	// assumed.
	Cluster *string `locationName:"cluster" type:"string"`

	// The container instance ID or full ARN of the container instance with which
	// to filter the ListTasks results. Specifying a containerInstance limits the
	// results to tasks that belong to that container instance.
	ContainerInstance *string `locationName:"containerInstance" type:"string"`

	// The task desired status with which to filter the ListTasks results. Specifying
	// a desiredStatus of STOPPED limits the results to tasks that Amazon ECS has
	// set the desired status to STOPPED. This can be useful for debugging tasks
	// that are not starting properly or have died or finished. The default status
	// filter is RUNNING, which shows tasks that Amazon ECS has set the desired
	// status to RUNNING.
	//
	// Although you can filter results based on a desired status of PENDING, this
	// does not return any results. Amazon ECS never sets the desired status of
	// a task to that value (only a task's lastStatus may have a value of PENDING).
	DesiredStatus enums.DesiredStatus `locationName:"desiredStatus" type:"string" enum:"true"`

	// The name of the family with which to filter the ListTasks results. Specifying
	// a family limits the results to tasks that belong to that family.
	Family *string `locationName:"family" type:"string"`

	// The launch type for services to list.
	LaunchType enums.LaunchType `locationName:"launchType" type:"string" enum:"true"`

	// The maximum number of task results returned by ListTasks in paginated output.
	// When this parameter is used, ListTasks only returns maxResults results in
	// a single page along with a nextToken response element. The remaining results
	// of the initial request can be seen by sending another ListTasks request with
	// the returned nextToken value. This value can be between 1 and 100. If this
	// parameter is not used, then ListTasks returns up to 100 results and a nextToken
	// value if applicable.
	MaxResults *int64 `locationName:"maxResults" type:"integer"`

	// The nextToken value returned from a previous paginated ListTasks request
	// where maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value.
	//
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The name of the service with which to filter the ListTasks results. Specifying
	// a serviceName limits the results to tasks that belong to that service.
	ServiceName *string `locationName:"serviceName" type:"string"`

	// The startedBy value with which to filter the task results. Specifying a startedBy
	// value limits the results to tasks that were started with that value.
	StartedBy *string `locationName:"startedBy" type:"string"`
}

// String returns the string representation
func (s ListTasksInput) String() string {
	return awsutil.Prettify(s)
}

type ListTasksOutput struct {
	_ struct{} `type:"structure"`

	// The nextToken value to include in a future ListTasks request. When the results
	// of a ListTasks request exceed maxResults, this value can be used to retrieve
	// the next page of results. This value is null when there are no more results
	// to return.
	NextToken *string `locationName:"nextToken" type:"string"`

	// The list of task ARN entries for the ListTasks request.
	TaskArns []string `locationName:"taskArns" type:"list"`
}

// String returns the string representation
func (s ListTasksOutput) String() string {
	return awsutil.Prettify(s)
}
