// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
)

const opResetEbsDefaultKmsKeyId = "ResetEbsDefaultKmsKeyId"

// ResetEbsDefaultKmsKeyIdRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Resets the default customer master key (CMK) for EBS encryption for your
// account in this Region to the AWS managed CMK for EBS.
//
// After resetting the default CMK to the AWS managed CMK, you can continue
// to encrypt by a customer managed CMK by specifying it when you create the
// volume. For more information, see Amazon EBS Encryption (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/EBSEncryption.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
//    // Example sending a request using ResetEbsDefaultKmsKeyIdRequest.
//    req := client.ResetEbsDefaultKmsKeyIdRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/ResetEbsDefaultKmsKeyId
func (c *Client) ResetEbsDefaultKmsKeyIdRequest(input *types.ResetEbsDefaultKmsKeyIdInput) ResetEbsDefaultKmsKeyIdRequest {
	op := &aws.Operation{
		Name:       opResetEbsDefaultKmsKeyId,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.ResetEbsDefaultKmsKeyIdInput{}
	}

	req := c.newRequest(op, input, &types.ResetEbsDefaultKmsKeyIdOutput{})
	return ResetEbsDefaultKmsKeyIdRequest{Request: req, Input: input, Copy: c.ResetEbsDefaultKmsKeyIdRequest}
}

// ResetEbsDefaultKmsKeyIdRequest is the request type for the
// ResetEbsDefaultKmsKeyId API operation.
type ResetEbsDefaultKmsKeyIdRequest struct {
	*aws.Request
	Input *types.ResetEbsDefaultKmsKeyIdInput
	Copy  func(*types.ResetEbsDefaultKmsKeyIdInput) ResetEbsDefaultKmsKeyIdRequest
}

// Send marshals and sends the ResetEbsDefaultKmsKeyId API request.
func (r ResetEbsDefaultKmsKeyIdRequest) Send(ctx context.Context) (*ResetEbsDefaultKmsKeyIdResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &ResetEbsDefaultKmsKeyIdResponse{
		ResetEbsDefaultKmsKeyIdOutput: r.Request.Data.(*types.ResetEbsDefaultKmsKeyIdOutput),
		response:                      &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// ResetEbsDefaultKmsKeyIdResponse is the response type for the
// ResetEbsDefaultKmsKeyId API operation.
type ResetEbsDefaultKmsKeyIdResponse struct {
	*types.ResetEbsDefaultKmsKeyIdOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// ResetEbsDefaultKmsKeyId request.
func (r *ResetEbsDefaultKmsKeyIdResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
