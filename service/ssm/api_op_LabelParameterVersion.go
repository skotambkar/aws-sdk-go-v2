// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ssm

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

const opLabelParameterVersion = "LabelParameterVersion"

// LabelParameterVersionRequest returns a request value for making API operation for
// Amazon Simple Systems Manager (SSM).
//
// A parameter label is a user-defined alias to help you manage different versions
// of a parameter. When you modify a parameter, Systems Manager automatically
// saves a new version and increments the version number by one. A label can
// help you remember the purpose of a parameter when there are multiple versions.
//
// Parameter labels have the following requirements and restrictions.
//
//    * A version of a parameter can have a maximum of 10 labels.
//
//    * You can't attach the same label to different versions of the same parameter.
//    For example, if version 1 has the label Production, then you can't attach
//    Production to version 2.
//
//    * You can move a label from one version of a parameter to another.
//
//    * You can't create a label when you create a new parameter. You must attach
//    a label to a specific version of a parameter.
//
//    * You can't delete a parameter label. If you no longer want to use a parameter
//    label, then you must move it to a different version of a parameter.
//
//    * A label can have a maximum of 100 characters.
//
//    * Labels can contain letters (case sensitive), numbers, periods (.), hyphens
//    (-), or underscores (_).
//
//    * Labels can't begin with a number, "aws," or "ssm" (not case sensitive).
//    If a label fails to meet these requirements, then the label is not associated
//    with a parameter and the system displays it in the list of InvalidLabels.
//
//    // Example sending a request using LabelParameterVersionRequest.
//    req := client.LabelParameterVersionRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ssm-2014-11-06/LabelParameterVersion
func (c *Client) LabelParameterVersionRequest(input *types.LabelParameterVersionInput) LabelParameterVersionRequest {
	op := &aws.Operation{
		Name:       opLabelParameterVersion,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.LabelParameterVersionInput{}
	}

	req := c.newRequest(op, input, &types.LabelParameterVersionOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.LabelParameterVersionMarshaler{Input: input}.GetNamedBuildHandler())

	return LabelParameterVersionRequest{Request: req, Input: input, Copy: c.LabelParameterVersionRequest}
}

// LabelParameterVersionRequest is the request type for the
// LabelParameterVersion API operation.
type LabelParameterVersionRequest struct {
	*aws.Request
	Input *types.LabelParameterVersionInput
	Copy  func(*types.LabelParameterVersionInput) LabelParameterVersionRequest
}

// Send marshals and sends the LabelParameterVersion API request.
func (r LabelParameterVersionRequest) Send(ctx context.Context) (*LabelParameterVersionResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &LabelParameterVersionResponse{
		LabelParameterVersionOutput: r.Request.Data.(*types.LabelParameterVersionOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// LabelParameterVersionResponse is the response type for the
// LabelParameterVersion API operation.
type LabelParameterVersionResponse struct {
	*types.LabelParameterVersionOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// LabelParameterVersion request.
func (r *LabelParameterVersionResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
