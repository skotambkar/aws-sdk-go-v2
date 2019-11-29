// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costandusagereportservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/costandusagereportservice/types"
)

const opDeleteReportDefinition = "DeleteReportDefinition"

// DeleteReportDefinitionRequest returns a request value for making API operation for
// AWS Cost and Usage Report Service.
//
// Deletes the specified report.
//
//    // Example sending a request using DeleteReportDefinitionRequest.
//    req := client.DeleteReportDefinitionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cur-2017-01-06/DeleteReportDefinition
func (c *Client) DeleteReportDefinitionRequest(input *types.DeleteReportDefinitionInput) DeleteReportDefinitionRequest {
	op := &aws.Operation{
		Name:       opDeleteReportDefinition,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteReportDefinitionInput{}
	}

	req := c.newRequest(op, input, &types.DeleteReportDefinitionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteReportDefinitionMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteReportDefinitionRequest{Request: req, Input: input, Copy: c.DeleteReportDefinitionRequest}
}

// DeleteReportDefinitionRequest is the request type for the
// DeleteReportDefinition API operation.
type DeleteReportDefinitionRequest struct {
	*aws.Request
	Input *types.DeleteReportDefinitionInput
	Copy  func(*types.DeleteReportDefinitionInput) DeleteReportDefinitionRequest
}

// Send marshals and sends the DeleteReportDefinition API request.
func (r DeleteReportDefinitionRequest) Send(ctx context.Context) (*DeleteReportDefinitionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteReportDefinitionResponse{
		DeleteReportDefinitionOutput: r.Request.Data.(*types.DeleteReportDefinitionOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteReportDefinitionResponse is the response type for the
// DeleteReportDefinition API operation.
type DeleteReportDefinitionResponse struct {
	*types.DeleteReportDefinitionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteReportDefinition request.
func (r *DeleteReportDefinitionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
