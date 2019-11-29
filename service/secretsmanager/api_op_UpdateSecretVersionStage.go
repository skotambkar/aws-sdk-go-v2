// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package secretsmanager

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/internal/aws_jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager/types"
)

const opUpdateSecretVersionStage = "UpdateSecretVersionStage"

// UpdateSecretVersionStageRequest returns a request value for making API operation for
// AWS Secrets Manager.
//
// Modifies the staging labels attached to a version of a secret. Staging labels
// are used to track a version as it progresses through the secret rotation
// process. You can attach a staging label to only one version of a secret at
// a time. If a staging label to be added is already attached to another version,
// then it is moved--removed from the other version first and then attached
// to this one. For more information about staging labels, see Staging Labels
// (https://docs.aws.amazon.com/secretsmanager/latest/userguide/terms-concepts.html#term_staging-label)
// in the AWS Secrets Manager User Guide.
//
// The staging labels that you specify in the VersionStage parameter are added
// to the existing list of staging labels--they don't replace it.
//
// You can move the AWSCURRENT staging label to this version by including it
// in this call.
//
// Whenever you move AWSCURRENT, Secrets Manager automatically moves the label
// AWSPREVIOUS to the version that AWSCURRENT was removed from.
//
// If this action results in the last label being removed from a version, then
// the version is considered to be 'deprecated' and can be deleted by Secrets
// Manager.
//
// Minimum permissions
//
// To run this command, you must have the following permissions:
//
//    * secretsmanager:UpdateSecretVersionStage
//
// Related operations
//
//    * To get the list of staging labels that are currently associated with
//    a version of a secret, use DescribeSecret and examine the SecretVersionsToStages
//    response value.
//
//    // Example sending a request using UpdateSecretVersionStageRequest.
//    req := client.UpdateSecretVersionStageRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/secretsmanager-2017-10-17/UpdateSecretVersionStage
func (c *Client) UpdateSecretVersionStageRequest(input *types.UpdateSecretVersionStageInput) UpdateSecretVersionStageRequest {
	op := &aws.Operation{
		Name:       opUpdateSecretVersionStage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.UpdateSecretVersionStageInput{}
	}

	req := c.newRequest(op, input, &types.UpdateSecretVersionStageOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(jsonrpc.BuildHandler.Name, aws_jsonrpc.UpdateSecretVersionStageMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateSecretVersionStageRequest{Request: req, Input: input, Copy: c.UpdateSecretVersionStageRequest}
}

// UpdateSecretVersionStageRequest is the request type for the
// UpdateSecretVersionStage API operation.
type UpdateSecretVersionStageRequest struct {
	*aws.Request
	Input *types.UpdateSecretVersionStageInput
	Copy  func(*types.UpdateSecretVersionStageInput) UpdateSecretVersionStageRequest
}

// Send marshals and sends the UpdateSecretVersionStage API request.
func (r UpdateSecretVersionStageRequest) Send(ctx context.Context) (*UpdateSecretVersionStageResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateSecretVersionStageResponse{
		UpdateSecretVersionStageOutput: r.Request.Data.(*types.UpdateSecretVersionStageOutput),
		response:                       &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateSecretVersionStageResponse is the response type for the
// UpdateSecretVersionStage API operation.
type UpdateSecretVersionStageResponse struct {
	*types.UpdateSecretVersionStageOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateSecretVersionStage request.
func (r *UpdateSecretVersionStageResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
