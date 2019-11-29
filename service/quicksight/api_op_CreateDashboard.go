// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opCreateDashboard = "CreateDashboard"

// CreateDashboardRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Creates a dashboard from a template. To first create a template, see the
// CreateTemplate API.
//
// A dashboard is an entity in QuickSight which identifies Quicksight reports,
// created from analyses. QuickSight dashboards are sharable. With the right
// permissions, you can create scheduled email reports from them. The CreateDashboard,
// DescribeDashboard and ListDashboardsByUser APIs act on the dashboard entity.
// If you have the correct permissions, you can create a dashboard from a template
// that exists in a different AWS account.
//
// CLI syntax:
//
// aws quicksight create-dashboard --cli-input-json file://create-dashboard.json
//
//    // Example sending a request using CreateDashboardRequest.
//    req := client.CreateDashboardRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/CreateDashboard
func (c *Client) CreateDashboardRequest(input *types.CreateDashboardInput) CreateDashboardRequest {
	op := &aws.Operation{
		Name:       opCreateDashboard,
		HTTPMethod: "POST",
		HTTPPath:   "/accounts/{AwsAccountId}/dashboards/{DashboardId}",
	}

	if input == nil {
		input = &types.CreateDashboardInput{}
	}

	req := c.newRequest(op, input, &types.CreateDashboardOutput{})
	return CreateDashboardRequest{Request: req, Input: input, Copy: c.CreateDashboardRequest}
}

// CreateDashboardRequest is the request type for the
// CreateDashboard API operation.
type CreateDashboardRequest struct {
	*aws.Request
	Input *types.CreateDashboardInput
	Copy  func(*types.CreateDashboardInput) CreateDashboardRequest
}

// Send marshals and sends the CreateDashboard API request.
func (r CreateDashboardRequest) Send(ctx context.Context) (*CreateDashboardResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateDashboardResponse{
		CreateDashboardOutput: r.Request.Data.(*types.CreateDashboardOutput),
		response:              &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateDashboardResponse is the response type for the
// CreateDashboard API operation.
type CreateDashboardResponse struct {
	*types.CreateDashboardOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateDashboard request.
func (r *CreateDashboardResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
