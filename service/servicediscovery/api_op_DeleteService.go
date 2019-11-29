// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/servicediscovery/types"
)

const opDeleteService = "DeleteService"

// DeleteServiceRequest returns a request value for making API operation for
// AWS Cloud Map.
//
// Deletes a specified service. If the service still contains one or more registered
// instances, the request fails.
//
//    // Example sending a request using DeleteServiceRequest.
//    req := client.DeleteServiceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/servicediscovery-2017-03-14/DeleteService
func (c *Client) DeleteServiceRequest(input *types.DeleteServiceInput) DeleteServiceRequest {
	op := &aws.Operation{
		Name:       opDeleteService,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteServiceInput{}
	}

	req := c.newRequest(op, input, &types.DeleteServiceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteServiceMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteServiceRequest{Request: req, Input: input, Copy: c.DeleteServiceRequest}
}

// DeleteServiceRequest is the request type for the
// DeleteService API operation.
type DeleteServiceRequest struct {
	*aws.Request
	Input *types.DeleteServiceInput
	Copy  func(*types.DeleteServiceInput) DeleteServiceRequest
}

// Send marshals and sends the DeleteService API request.
func (r DeleteServiceRequest) Send(ctx context.Context) (*DeleteServiceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteServiceResponse{
		DeleteServiceOutput: r.Request.Data.(*types.DeleteServiceOutput),
		response:            &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteServiceResponse is the response type for the
// DeleteService API operation.
type DeleteServiceResponse struct {
	*types.DeleteServiceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteService request.
func (r *DeleteServiceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
