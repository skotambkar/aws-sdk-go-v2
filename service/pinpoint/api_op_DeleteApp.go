// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package pinpoint

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/pinpoint/types"
)

const opDeleteApp = "DeleteApp"

// DeleteAppRequest returns a request value for making API operation for
// Amazon Pinpoint.
//
// Deletes an application.
//
//    // Example sending a request using DeleteAppRequest.
//    req := client.DeleteAppRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/pinpoint-2016-12-01/DeleteApp
func (c *Client) DeleteAppRequest(input *types.DeleteAppInput) DeleteAppRequest {
	op := &aws.Operation{
		Name:       opDeleteApp,
		HTTPMethod: "DELETE",
		HTTPPath:   "/v1/apps/{application-id}",
	}

	if input == nil {
		input = &types.DeleteAppInput{}
	}

	req := c.newRequest(op, input, &types.DeleteAppOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeleteAppMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteAppRequest{Request: req, Input: input, Copy: c.DeleteAppRequest}
}

// DeleteAppRequest is the request type for the
// DeleteApp API operation.
type DeleteAppRequest struct {
	*aws.Request
	Input *types.DeleteAppInput
	Copy  func(*types.DeleteAppInput) DeleteAppRequest
}

// Send marshals and sends the DeleteApp API request.
func (r DeleteAppRequest) Send(ctx context.Context) (*DeleteAppResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteAppResponse{
		DeleteAppOutput: r.Request.Data.(*types.DeleteAppOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteAppResponse is the response type for the
// DeleteApp API operation.
type DeleteAppResponse struct {
	*types.DeleteAppOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApp request.
func (r *DeleteAppResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
