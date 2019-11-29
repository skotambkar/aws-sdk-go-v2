// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
)

const opDeleteApplicationVersion = "DeleteApplicationVersion"

// DeleteApplicationVersionRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Deletes the specified version from the specified application.
//
// You cannot delete an application version that is associated with a running
// environment.
//
//    // Example sending a request using DeleteApplicationVersionRequest.
//    req := client.DeleteApplicationVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DeleteApplicationVersion
func (c *Client) DeleteApplicationVersionRequest(input *types.DeleteApplicationVersionInput) DeleteApplicationVersionRequest {
	op := &aws.Operation{
		Name:       opDeleteApplicationVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteApplicationVersionInput{}
	}

	req := c.newRequest(op, input, &types.DeleteApplicationVersionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DeleteApplicationVersionMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteApplicationVersionRequest{Request: req, Input: input, Copy: c.DeleteApplicationVersionRequest}
}

// DeleteApplicationVersionRequest is the request type for the
// DeleteApplicationVersion API operation.
type DeleteApplicationVersionRequest struct {
	*aws.Request
	Input *types.DeleteApplicationVersionInput
	Copy  func(*types.DeleteApplicationVersionInput) DeleteApplicationVersionRequest
}

// Send marshals and sends the DeleteApplicationVersion API request.
func (r DeleteApplicationVersionRequest) Send(ctx context.Context) (*DeleteApplicationVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApplicationVersionResponse{
		DeleteApplicationVersionOutput: r.Request.Data.(*types.DeleteApplicationVersionOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApplicationVersionResponse is the response type for the
// DeleteApplicationVersion API operation.
type DeleteApplicationVersionResponse struct {
	*types.DeleteApplicationVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApplicationVersion request.
func (r *DeleteApplicationVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
