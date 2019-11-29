// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package elasticbeanstalk

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/elasticbeanstalk/types"
)

const opUpdateTagsForResource = "UpdateTagsForResource"

// UpdateTagsForResourceRequest returns a request value for making API operation for
// AWS Elastic Beanstalk.
//
// Update the list of tags applied to an AWS Elastic Beanstalk resource. Two
// lists can be passed: TagsToAdd for tags to add or update, and TagsToRemove.
//
// Currently, Elastic Beanstalk only supports tagging of Elastic Beanstalk environments.
// For details about environment tagging, see Tagging Resources in Your Elastic
// Beanstalk Environment (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/using-features.tagging.html).
//
// If you create a custom IAM user policy to control permission to this operation,
// specify one of the following two virtual actions (or both) instead of the
// API operation name:
//
// elasticbeanstalk:AddTags
//
// Controls permission to call UpdateTagsForResource and pass a list of tags
// to add in the TagsToAdd parameter.
//
// elasticbeanstalk:RemoveTags
//
// Controls permission to call UpdateTagsForResource and pass a list of tag
// keys to remove in the TagsToRemove parameter.
//
// For details about creating a custom user policy, see Creating a Custom User
// Policy (https://docs.aws.amazon.com/elasticbeanstalk/latest/dg/AWSHowTo.iam.managed-policies.html#AWSHowTo.iam.policies).
//
//    // Example sending a request using UpdateTagsForResourceRequest.
//    req := client.UpdateTagsForResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/elasticbeanstalk-2010-12-01/UpdateTagsForResource
func (c *Client) UpdateTagsForResourceRequest(input *types.UpdateTagsForResourceInput) UpdateTagsForResourceRequest {
	op := &aws.Operation{
		Name:       opUpdateTagsForResource,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateTagsForResourceInput{}
	}

	req := c.newRequest(op, input, &types.UpdateTagsForResourceOutput{})
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return UpdateTagsForResourceRequest{Request: req, Input: input, Copy: c.UpdateTagsForResourceRequest}
}

// UpdateTagsForResourceRequest is the request type for the
// UpdateTagsForResource API operation.
type UpdateTagsForResourceRequest struct {
	*aws.Request
	Input *types.UpdateTagsForResourceInput
	Copy  func(*types.UpdateTagsForResourceInput) UpdateTagsForResourceRequest
}

// Send marshals and sends the UpdateTagsForResource API request.
func (r UpdateTagsForResourceRequest) Send(ctx context.Context) (*UpdateTagsForResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTagsForResourceResponse{
		UpdateTagsForResourceOutput: r.Request.Data.(*types.UpdateTagsForResourceOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTagsForResourceResponse is the response type for the
// UpdateTagsForResource API operation.
type UpdateTagsForResourceResponse struct {
	*types.UpdateTagsForResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTagsForResource request.
func (r *UpdateTagsForResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
