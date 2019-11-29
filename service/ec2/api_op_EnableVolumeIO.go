// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opEnableVolumeIO = "EnableVolumeIO"

// EnableVolumeIORequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Enables I/O operations for a volume that had I/O operations disabled because
// the data on the volume was potentially inconsistent.
//
//    // Example sending a request using EnableVolumeIORequest.
//    req := client.EnableVolumeIORequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/EnableVolumeIO
func (c *Client) EnableVolumeIORequest(input *types.EnableVolumeIOInput) EnableVolumeIORequest {
	op := &aws.Operation{
		Name:       opEnableVolumeIO,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.EnableVolumeIOInput{}
	}

	req := c.newRequest(op, input, &types.EnableVolumeIOOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.EnableVolumeIOMarshaler{Input: input}.GetNamedBuildHandler())

	req.Handlers.Unmarshal.Remove(ec2query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return EnableVolumeIORequest{Request: req, Input: input, Copy: c.EnableVolumeIORequest}
}

// EnableVolumeIORequest is the request type for the
// EnableVolumeIO API operation.
type EnableVolumeIORequest struct {
	*aws.Request
	Input *types.EnableVolumeIOInput
	Copy  func(*types.EnableVolumeIOInput) EnableVolumeIORequest
}

// Send marshals and sends the EnableVolumeIO API request.
func (r EnableVolumeIORequest) Send(ctx context.Context) (*EnableVolumeIOResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &EnableVolumeIOResponse{
		EnableVolumeIOOutput: r.Request.Data.(*types.EnableVolumeIOOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// EnableVolumeIOResponse is the response type for the
// EnableVolumeIO API operation.
type EnableVolumeIOResponse struct {
	*types.EnableVolumeIOOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// EnableVolumeIO request.
func (r *EnableVolumeIOResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
