// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package applicationinsights

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/applicationinsights/types"
)

const opUpdateComponentConfiguration = "UpdateComponentConfiguration"

// UpdateComponentConfigurationRequest returns a request value for making API operation for
// Amazon CloudWatch Application Insights.
//
// Updates the monitoring configurations for the component. The configuration
// input parameter is an escaped JSON of the configuration and should match
// the schema of what is returned by DescribeComponentConfigurationRecommendation.
//
//    // Example sending a request using UpdateComponentConfigurationRequest.
//    req := client.UpdateComponentConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/application-insights-2018-11-25/UpdateComponentConfiguration
func (c *Client) UpdateComponentConfigurationRequest(input *types.UpdateComponentConfigurationInput) UpdateComponentConfigurationRequest {
	op := &aws.Operation{
		Name:       opUpdateComponentConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateComponentConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.UpdateComponentConfigurationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateComponentConfigurationMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateComponentConfigurationRequest{Request: req, Input: input, Copy: c.UpdateComponentConfigurationRequest}
}

// UpdateComponentConfigurationRequest is the request type for the
// UpdateComponentConfiguration API operation.
type UpdateComponentConfigurationRequest struct {
	*aws.Request
	Input *types.UpdateComponentConfigurationInput
	Copy  func(*types.UpdateComponentConfigurationInput) UpdateComponentConfigurationRequest
}

// Send marshals and sends the UpdateComponentConfiguration API request.
func (r UpdateComponentConfigurationRequest) Send(ctx context.Context) (*UpdateComponentConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateComponentConfigurationResponse{
		UpdateComponentConfigurationOutput: r.Request.Data.(*types.UpdateComponentConfigurationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateComponentConfigurationResponse is the response type for the
// UpdateComponentConfiguration API operation.
type UpdateComponentConfigurationResponse struct {
	*types.UpdateComponentConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateComponentConfiguration request.
func (r *UpdateComponentConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
