// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package cloudtrail

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
)

const opCreateTrail = "CreateTrail"

// CreateTrailRequest returns a request value for making API operation for
// AWS CloudTrail.
//
// Creates a trail that specifies the settings for delivery of log data to an
// Amazon S3 bucket.
//
//    // Example sending a request using CreateTrailRequest.
//    req := client.CreateTrailRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/cloudtrail-2013-11-01/CreateTrail
func (c *Client) CreateTrailRequest(input *types.CreateTrailInput) CreateTrailRequest {
	op := &aws.Operation{
		Name:       opCreateTrail,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.CreateTrailInput{}
	}

	req := c.newRequest(op, input, &types.CreateTrailOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.CreateTrailMarshaler{Input: input}.GetNamedBuildHandler())

	return CreateTrailRequest{Request: req, Input: input, Copy: c.CreateTrailRequest}
}

// CreateTrailRequest is the request type for the
// CreateTrail API operation.
type CreateTrailRequest struct {
	*aws.Request
	Input *types.CreateTrailInput
	Copy  func(*types.CreateTrailInput) CreateTrailRequest
}

// Send marshals and sends the CreateTrail API request.
func (r CreateTrailRequest) Send(ctx context.Context) (*CreateTrailResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateTrailResponse{
		CreateTrailOutput: r.Request.Data.(*types.CreateTrailOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateTrailResponse is the response type for the
// CreateTrail API operation.
type CreateTrailResponse struct {
	*types.CreateTrailOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateTrail request.
func (r *CreateTrailResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
