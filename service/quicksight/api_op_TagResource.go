// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opTagResource = "TagResource"

// TagResourceRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Assigns a tag or tags to a resource.
//
// Assigns one or more tags (key-value pairs) to the specified QuickSight resource.
// Tags can help you organize and categorize your resources. You can also use
// them to scope user permissions, by granting a user permission to access or
// change only resources with certain tag values. You can use the TagResource
// action with a resource that already has tags. If you specify a new tag key
// for the resource, this tag is appended to the list of tags associated with
// the resource. If you specify a tag key that is already associated with the
// resource, the new tag value that you specify replaces the previous value
// for that tag.
//
// You can associate as many as 50 tags with a resource. QuickSight supports
// tagging on data-set, data-source, dashboard, template.
//
// Tagging for QuickSight works in a similar was to tagging for other AWS services,
// except for the following:
//
//    * You can't use tags to track AWS costs for QuickSight, because QuickSight
//    costs are based on users and SPICE capacity, which aren't taggable resources.
//
//    * QuickSight doesn't currently support the Tag Editor for AWS Resource
//    Groups.
//
// CLI syntax to tag a resource:
//
//    * aws quicksight tag-resource --resource-arn arn:aws:quicksight:us-east-1:111111111111:dataset/dataset1
//    --tags Key=K1,Value=V1 Key=K2,Value=V2 --region us-east-1
//
//    // Example sending a request using TagResourceRequest.
//    req := client.TagResourceRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/TagResource
func (c *Client) TagResourceRequest(input *types.TagResourceInput) TagResourceRequest {
	op := &aws.Operation{
		Name:       opTagResource,
		HTTPMethod: "POST",
		HTTPPath:   "/resources/{ResourceArn}/tags",
	}

	if input == nil {
		input = &types.TagResourceInput{}
	}

	req := c.newRequest(op, input, &types.TagResourceOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.TagResourceMarshaler{Input: input}.GetNamedBuildHandler())

	return TagResourceRequest{Request: req, Input: input, Copy: c.TagResourceRequest}
}

// TagResourceRequest is the request type for the
// TagResource API operation.
type TagResourceRequest struct {
	*aws.Request
	Input *types.TagResourceInput
	Copy  func(*types.TagResourceInput) TagResourceRequest
}

// Send marshals and sends the TagResource API request.
func (r TagResourceRequest) Send(ctx context.Context) (*TagResourceResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &TagResourceResponse{
		TagResourceOutput: r.Request.Data.(*types.TagResourceOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// TagResourceResponse is the response type for the
// TagResource API operation.
type TagResourceResponse struct {
	*types.TagResourceOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// TagResource request.
func (r *TagResourceResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
