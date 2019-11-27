// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package configservice

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
)

const opGetDiscoveredResourceCounts = "GetDiscoveredResourceCounts"

// GetDiscoveredResourceCountsRequest returns a request value for making API operation for
// AWS Config.
//
// Returns the resource types, the number of each resource type, and the total
// number of resources that AWS Config is recording in this region for your
// AWS account.
//
// Example
//
// AWS Config is recording three resource types in the US East (Ohio) Region
// for your account: 25 EC2 instances, 20 IAM users, and 15 S3 buckets.
//
// You make a call to the GetDiscoveredResourceCounts action and specify that
// you want all resource types.
//
// AWS Config returns the following:
//
//    * The resource types (EC2 instances, IAM users, and S3 buckets).
//
//    * The number of each resource type (25, 20, and 15).
//
//    * The total number of all resources (60).
//
// The response is paginated. By default, AWS Config lists 100 ResourceCount
// objects on each page. You can customize this number with the limit parameter.
// The response includes a nextToken string. To get the next page of results,
// run the request again and specify the string for the nextToken parameter.
//
// If you make a call to the GetDiscoveredResourceCounts action, you might not
// immediately receive resource counts in the following situations:
//
//    * You are a new AWS Config customer.
//
//    * You just enabled resource recording.
//
// It might take a few minutes for AWS Config to record and count your resources.
// Wait a few minutes and then retry the GetDiscoveredResourceCounts action.
//
//    // Example sending a request using GetDiscoveredResourceCountsRequest.
//    req := client.GetDiscoveredResourceCountsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/config-2014-11-12/GetDiscoveredResourceCounts
func (c *Client) GetDiscoveredResourceCountsRequest(input *types.GetDiscoveredResourceCountsInput) GetDiscoveredResourceCountsRequest {
	op := &aws.Operation{
		Name:       opGetDiscoveredResourceCounts,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.GetDiscoveredResourceCountsInput{}
	}

	req := c.newRequest(op, input, &types.GetDiscoveredResourceCountsOutput{})
	return GetDiscoveredResourceCountsRequest{Request: req, Input: input, Copy: c.GetDiscoveredResourceCountsRequest}
}

// GetDiscoveredResourceCountsRequest is the request type for the
// GetDiscoveredResourceCounts API operation.
type GetDiscoveredResourceCountsRequest struct {
	*aws.Request
	Input *types.GetDiscoveredResourceCountsInput
	Copy  func(*types.GetDiscoveredResourceCountsInput) GetDiscoveredResourceCountsRequest
}

// Send marshals and sends the GetDiscoveredResourceCounts API request.
func (r GetDiscoveredResourceCountsRequest) Send(ctx context.Context) (*GetDiscoveredResourceCountsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetDiscoveredResourceCountsResponse{
		GetDiscoveredResourceCountsOutput: r.Request.Data.(*types.GetDiscoveredResourceCountsOutput),
		response:                          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetDiscoveredResourceCountsResponse is the response type for the
// GetDiscoveredResourceCounts API operation.
type GetDiscoveredResourceCountsResponse struct {
	*types.GetDiscoveredResourceCountsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetDiscoveredResourceCounts request.
func (r *GetDiscoveredResourceCountsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
