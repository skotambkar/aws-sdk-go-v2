// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opDescribeDashboardPermissions = "DescribeDashboardPermissions"

// DescribeDashboardPermissionsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Describes read and write permissions on a dashboard.
//
// CLI syntax:
//
// aws quicksight describe-dashboard-permissions --aws-account-id 735340738645
// —dashboard-id reports_test_bob_report
//
//    // Example sending a request using DescribeDashboardPermissionsRequest.
//    req := client.DescribeDashboardPermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DescribeDashboardPermissions
func (c *Client) DescribeDashboardPermissionsRequest(input *types.DescribeDashboardPermissionsInput) DescribeDashboardPermissionsRequest {
	op := &aws.Operation{
		Name:       opDescribeDashboardPermissions,
		HTTPMethod: "GET",
		HTTPPath:   "/accounts/{AwsAccountId}/dashboards/{DashboardId}/permissions",
	}

	if input == nil {
		input = &types.DescribeDashboardPermissionsInput{}
	}

	req := c.newRequest(op, input, &types.DescribeDashboardPermissionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DescribeDashboardPermissionsMarshaler{Input: input}.GetNamedBuildHandler())

	return DescribeDashboardPermissionsRequest{Request: req, Input: input, Copy: c.DescribeDashboardPermissionsRequest}
}

// DescribeDashboardPermissionsRequest is the request type for the
// DescribeDashboardPermissions API operation.
type DescribeDashboardPermissionsRequest struct {
	*aws.Request
	Input *types.DescribeDashboardPermissionsInput
	Copy  func(*types.DescribeDashboardPermissionsInput) DescribeDashboardPermissionsRequest
}

// Send marshals and sends the DescribeDashboardPermissions API request.
func (r DescribeDashboardPermissionsRequest) Send(ctx context.Context) (*DescribeDashboardPermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeDashboardPermissionsResponse{
		DescribeDashboardPermissionsOutput: r.Request.Data.(*types.DescribeDashboardPermissionsOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeDashboardPermissionsResponse is the response type for the
// DescribeDashboardPermissions API operation.
type DescribeDashboardPermissionsResponse struct {
	*types.DescribeDashboardPermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeDashboardPermissions request.
func (r *DescribeDashboardPermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
