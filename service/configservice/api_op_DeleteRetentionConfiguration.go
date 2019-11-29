// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/configservice/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
)

const opDeleteRetentionConfiguration = "DeleteRetentionConfiguration"

// DeleteRetentionConfigurationRequest returns a request value for making API operation for
// AWS Config.
//
// Deletes the retention configuration.
//
//    // Example sending a request using DeleteRetentionConfigurationRequest.
//    req := client.DeleteRetentionConfigurationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/DeleteRetentionConfiguration
func (c *Client) DeleteRetentionConfigurationRequest(input *types.DeleteRetentionConfigurationInput) DeleteRetentionConfigurationRequest {
	op := &aws.Operation{
		Name:       opDeleteRetentionConfiguration,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteRetentionConfigurationInput{}
	}

	req := c.newRequest(op, input, &types.DeleteRetentionConfigurationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.DeleteRetentionConfigurationMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(jsonrpc.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteRetentionConfigurationRequest{Request: req, Input: input, Copy: c.DeleteRetentionConfigurationRequest}
}

// DeleteRetentionConfigurationRequest is the request type for the
// DeleteRetentionConfiguration API operation.
type DeleteRetentionConfigurationRequest struct {
	*aws.Request
	Input *types.DeleteRetentionConfigurationInput
	Copy  func(*types.DeleteRetentionConfigurationInput) DeleteRetentionConfigurationRequest
}

// Send marshals and sends the DeleteRetentionConfiguration API request.
func (r DeleteRetentionConfigurationRequest) Send(ctx context.Context) (*DeleteRetentionConfigurationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteRetentionConfigurationResponse{
		DeleteRetentionConfigurationOutput: r.Request.Data.(*types.DeleteRetentionConfigurationOutput),
		response:                           &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteRetentionConfigurationResponse is the response type for the
// DeleteRetentionConfiguration API operation.
type DeleteRetentionConfigurationResponse struct {
	*types.DeleteRetentionConfigurationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteRetentionConfiguration request.
func (r *DeleteRetentionConfigurationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
