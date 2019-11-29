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

const opDeleteApplication = "DeleteApplication"

// DeleteApplicationRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Deletes the specified application along with all associated versions and
// configurations. The application versions will not be deleted from your Amazon
// S3 bucket.
//
// You cannot delete an application that has a running environment.
//
//    // Example sending a request using DeleteApplicationRequest.
//    req := client.DeleteApplicationRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/DeleteApplication
func (c *Client) DeleteApplicationRequest(input *types.DeleteApplicationInput) DeleteApplicationRequest {
	op := &aws.Operation{
		Name:       opDeleteApplication,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DeleteApplicationInput{}
	}

	req := c.newRequest(op, input, &types.DeleteApplicationOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.DeleteApplicationMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return DeleteApplicationRequest{Request: req, Input: input, Copy: c.DeleteApplicationRequest}
}

// DeleteApplicationRequest is the request type for the
// DeleteApplication API operation.
type DeleteApplicationRequest struct {
	*aws.Request
	Input *types.DeleteApplicationInput
	Copy  func(*types.DeleteApplicationInput) DeleteApplicationRequest
}

// Send marshals and sends the DeleteApplication API request.
func (r DeleteApplicationRequest) Send(ctx context.Context) (*DeleteApplicationResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteApplicationResponse{
		DeleteApplicationOutput: r.Request.Data.(*types.DeleteApplicationOutput),
		response:                &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteApplicationResponse is the response type for the
// DeleteApplication API operation.
type DeleteApplicationResponse struct {
	*types.DeleteApplicationOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteApplication request.
func (r *DeleteApplicationResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
