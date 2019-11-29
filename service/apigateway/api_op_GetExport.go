// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opGetExport = "GetExport"

// GetExportRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Exports a deployed version of a RestApi in a specified format.
//
//    // Example sending a request using GetExportRequest.
//    req := client.GetExportRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetExportRequest(input *types.GetExportInput) GetExportRequest {
	op := &aws.Operation{
		Name:       opGetExport,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/stages/{stage_name}/exports/{export_type}",
	}

	if input == nil {
		input = &types.GetExportInput{}
	}

	req := c.newRequest(op, input, &types.GetExportOutput{})
	return GetExportRequest{Request: req, Input: input, Copy: c.GetExportRequest}
}

// GetExportRequest is the request type for the
// GetExport API operation.
type GetExportRequest struct {
	*aws.Request
	Input *types.GetExportInput
	Copy  func(*types.GetExportInput) GetExportRequest
}

// Send marshals and sends the GetExport API request.
func (r GetExportRequest) Send(ctx context.Context) (*GetExportResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetExportResponse{
		GetExportOutput: r.Request.Data.(*types.GetExportOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetExportResponse is the response type for the
// GetExport API operation.
type GetExportResponse struct {
	*types.GetExportOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetExport request.
func (r *GetExportResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
