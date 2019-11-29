// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package guardduty

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/guardduty/types"
)

const opGetFindings = "GetFindings"

// GetFindingsRequest returns a request value for making API operation for
// Amazon GuardDuty.
//
// Describes Amazon GuardDuty findings specified by finding IDs.
//
//    // Example sending a request using GetFindingsRequest.
//    req := client.GetFindingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/guardduty-2017-11-28/GetFindings
func (c *Client) GetFindingsRequest(input *types.GetFindingsInput) GetFindingsRequest {
	op := &aws.Operation{
		Name:       opGetFindings,
		HTTPMethod: "POST",
		HTTPPath:   "/detector/{detectorId}/findings/get",
	}

	if input == nil {
		input = &types.GetFindingsInput{}
	}

	req := c.newRequest(op, input, &types.GetFindingsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.GetFindingsMarshaler{Input: input}.GetNamedBuildHandler())

	return GetFindingsRequest{Request: req, Input: input, Copy: c.GetFindingsRequest}
}

// GetFindingsRequest is the request type for the
// GetFindings API operation.
type GetFindingsRequest struct {
	*aws.Request
	Input *types.GetFindingsInput
	Copy  func(*types.GetFindingsInput) GetFindingsRequest
}

// Send marshals and sends the GetFindings API request.
func (r GetFindingsRequest) Send(ctx context.Context) (*GetFindingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &GetFindingsResponse{
		GetFindingsOutput: r.Request.Data.(*types.GetFindingsOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// GetFindingsResponse is the response type for the
// GetFindings API operation.
type GetFindingsResponse struct {
	*types.GetFindingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// GetFindings request.
func (r *GetFindingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
