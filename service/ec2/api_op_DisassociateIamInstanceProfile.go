// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/internal/aws_ec2query"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opDisassociateIamInstanceProfile = "DisassociateIamInstanceProfile"

// DisassociateIamInstanceProfileRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Disassociates an IAM instance profile from a running or stopped instance.
//
// Use DescribeIamInstanceProfileAssociations to get the association ID.
//
//    // Example sending a request using DisassociateIamInstanceProfileRequest.
//    req := client.DisassociateIamInstanceProfileRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/DisassociateIamInstanceProfile
func (c *Client) DisassociateIamInstanceProfileRequest(input *types.DisassociateIamInstanceProfileInput) DisassociateIamInstanceProfileRequest {
	op := &aws.Operation{
		Name:       opDisassociateIamInstanceProfile,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.DisassociateIamInstanceProfileInput{}
	}

	req := c.newRequest(op, input, &types.DisassociateIamInstanceProfileOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(ec2query.BuildHandler.Name, aws_ec2query.DisassociateIamInstanceProfileMarshaler{Input: input}.GetNamedBuildHandler())

	return DisassociateIamInstanceProfileRequest{Request: req, Input: input, Copy: c.DisassociateIamInstanceProfileRequest}
}

// DisassociateIamInstanceProfileRequest is the request type for the
// DisassociateIamInstanceProfile API operation.
type DisassociateIamInstanceProfileRequest struct {
	*aws.Request
	Input *types.DisassociateIamInstanceProfileInput
	Copy  func(*types.DisassociateIamInstanceProfileInput) DisassociateIamInstanceProfileRequest
}

// Send marshals and sends the DisassociateIamInstanceProfile API request.
func (r DisassociateIamInstanceProfileRequest) Send(ctx context.Context) (*DisassociateIamInstanceProfileResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DisassociateIamInstanceProfileResponse{
		DisassociateIamInstanceProfileOutput: r.Request.Data.(*types.DisassociateIamInstanceProfileOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DisassociateIamInstanceProfileResponse is the response type for the
// DisassociateIamInstanceProfile API operation.
type DisassociateIamInstanceProfileResponse struct {
	*types.DisassociateIamInstanceProfileOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DisassociateIamInstanceProfile request.
func (r *DisassociateIamInstanceProfileResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
