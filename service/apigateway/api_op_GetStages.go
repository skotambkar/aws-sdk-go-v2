// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigateway

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/apigateway/types"
)

const opGetStages = "GetStages"

// GetStagesRequest returns a request value for making API operation for
// Amazon API Gateway.
//
// Gets information about one or more Stage resources.
//
//    // Example sending a request using GetStagesRequest.
//    req := client.GetStagesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) GetStagesRequest(input *types.GetStagesInput) GetStagesRequest {
	op := &aws.Operation{
		Name:       opGetStages,
		HTTPMethod: "GET",
		HTTPPath:   "/restapis/{restapi_id}/stages",
	}

	if input == nil {
		input = &types.GetStagesInput{}
	}

	req := c.newRequest(op, input, &types.GetStagesOutput{})
	return GetStagesRequest{Request: req, Input: input, Copy: c.GetStagesRequest}
}

// GetStagesRequest is the request type for the
// GetStages API operation.
type GetStagesRequest struct {
	*aws.Request
	Input *types.GetStagesInput
	Copy  func(*types.GetStagesInput) GetStagesRequest
}

// Send marshals and sends the GetStages API request.
func (r GetStagesRequest) Send(ctx context.Context) (*GetStagesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetStagesResponse{
		GetStagesOutput: r.Request.Data.(*types.GetStagesOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetStagesResponse is the response type for the
// GetStages API operation.
type GetStagesResponse struct {
	*types.GetStagesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetStages request.
func (r *GetStagesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
