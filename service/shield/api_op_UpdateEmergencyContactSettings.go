// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package shield

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/shield/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
)

const opUpdateEmergencyContactSettings = "UpdateEmergencyContactSettings"

// UpdateEmergencyContactSettingsRequest returns a request value for making API operation for
// AWS Shield.
//
// Updates the details of the list of email addresses that the DRT can use to
// contact you during a suspected attack.
//
//    // Example sending a request using UpdateEmergencyContactSettingsRequest.
//    req := client.UpdateEmergencyContactSettingsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/shield-2016-06-02/UpdateEmergencyContactSettings
func (c *Client) UpdateEmergencyContactSettingsRequest(input *types.UpdateEmergencyContactSettingsInput) UpdateEmergencyContactSettingsRequest {
	op := &aws.Operation{
		Name:       opUpdateEmergencyContactSettings,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateEmergencyContactSettingsInput{}
	}

	req := c.newRequest(op, input, &types.UpdateEmergencyContactSettingsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateEmergencyContactSettingsMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateEmergencyContactSettingsRequest{Request: req, Input: input, Copy: c.UpdateEmergencyContactSettingsRequest}
}

// UpdateEmergencyContactSettingsRequest is the request type for the
// UpdateEmergencyContactSettings API operation.
type UpdateEmergencyContactSettingsRequest struct {
	*aws.Request
	Input *types.UpdateEmergencyContactSettingsInput
	Copy  func(*types.UpdateEmergencyContactSettingsInput) UpdateEmergencyContactSettingsRequest
}

// Send marshals and sends the UpdateEmergencyContactSettings API request.
func (r UpdateEmergencyContactSettingsRequest) Send(ctx context.Context) (*UpdateEmergencyContactSettingsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateEmergencyContactSettingsResponse{
		UpdateEmergencyContactSettingsOutput: r.Request.Data.(*types.UpdateEmergencyContactSettingsOutput),
		response:                             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateEmergencyContactSettingsResponse is the response type for the
// UpdateEmergencyContactSettings API operation.
type UpdateEmergencyContactSettingsResponse struct {
	*types.UpdateEmergencyContactSettingsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateEmergencyContactSettings request.
func (r *UpdateEmergencyContactSettingsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
