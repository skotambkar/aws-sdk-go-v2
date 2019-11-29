// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package amplify

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/amplify/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/amplify/types"
)

const opUpdateApp = "UpdateApp"

// UpdateAppRequest returns a request value for making API operation for
// AWS Amplify.
//
// Updates an existing Amplify App.
//
//    // Example sending a request using UpdateAppRequest.
//    req := client.UpdateAppRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/amplify-2017-07-25/UpdateApp
func (c *Client) UpdateAppRequest(input *types.UpdateAppInput) UpdateAppRequest {
	op := &aws.Operation{
		Name:       opUpdateApp,
		HTTPMethod: "POST",
		HTTPPath:   "/apps/{appId}",
	}

	if input == nil {
		input = &types.UpdateAppInput{}
	}

	req := c.newRequest(op, input, &types.UpdateAppOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateAppMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateAppRequest{Request: req, Input: input, Copy: c.UpdateAppRequest}
}

// UpdateAppRequest is the request type for the
// UpdateApp API operation.
type UpdateAppRequest struct {
	*aws.Request
	Input *types.UpdateAppInput
	Copy  func(*types.UpdateAppInput) UpdateAppRequest
}

// Send marshals and sends the UpdateApp API request.
func (r UpdateAppRequest) Send(ctx context.Context) (*UpdateAppResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateAppResponse{
		UpdateAppOutput: r.Request.Data.(*types.UpdateAppOutput),
		response:        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateAppResponse is the response type for the
// UpdateApp API operation.
type UpdateAppResponse struct {
	*types.UpdateAppOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateApp request.
func (r *UpdateAppResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
