// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package glacier

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/glacier/types"
)

const opSetVaultNotifications = "SetVaultNotifications"

// SetVaultNotificationsRequest returns a request value for making API operation for
// Amazon Glacier.
//
// This operation configures notifications that will be sent when specific events
// happen to a vault. By default, you don't get any notifications.
//
// To configure vault notifications, send a PUT request to the notification-configuration
// subresource of the vault. The request should include a JSON document that
// provides an Amazon SNS topic and specific events for which you want Amazon
// S3 Glacier to send notifications to the topic.
//
// Amazon SNS topics must grant permission to the vault to be allowed to publish
// notifications to the topic. You can configure a vault to publish a notification
// for the following vault events:
//
//    * ArchiveRetrievalCompleted This event occurs when a job that was initiated
//    for an archive retrieval is completed (InitiateJob). The status of the
//    completed job can be "Succeeded" or "Failed". The notification sent to
//    the SNS topic is the same output as returned from DescribeJob.
//
//    * InventoryRetrievalCompleted This event occurs when a job that was initiated
//    for an inventory retrieval is completed (InitiateJob). The status of the
//    completed job can be "Succeeded" or "Failed". The notification sent to
//    the SNS topic is the same output as returned from DescribeJob.
//
// An AWS account has full permission to perform all operations (actions). However,
// AWS Identity and Access Management (IAM) users don't have any permissions
// by default. You must grant them explicit permission to perform specific actions.
// For more information, see Access Control Using AWS Identity and Access Management
// (IAM) (https://docs.aws.amazon.com/amazonglacier/latest/dev/using-iam-with-amazon-glacier.html).
//
// For conceptual information and underlying REST API, see Configuring Vault
// Notifications in Amazon S3 Glacier (https://docs.aws.amazon.com/amazonglacier/latest/dev/configuring-notifications.html)
// and Set Vault Notification Configuration (https://docs.aws.amazon.com/amazonglacier/latest/dev/api-vault-notifications-put.html)
// in the Amazon Glacier Developer Guide.
//
//    // Example sending a request using SetVaultNotificationsRequest.
//    req := client.SetVaultNotificationsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
func (c *Client) SetVaultNotificationsRequest(input *types.SetVaultNotificationsInput) SetVaultNotificationsRequest {
	op := &aws.Operation{
		Name:       opSetVaultNotifications,
		HTTPMethod: "PUT",
		HTTPPath:   "/{accountId}/vaults/{vaultName}/notification-configuration",
	}

	if input == nil {
		input = &types.SetVaultNotificationsInput{}
	}

	req := c.newRequest(op, input, &types.SetVaultNotificationsOutput{})
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return SetVaultNotificationsRequest{Request: req, Input: input, Copy: c.SetVaultNotificationsRequest}
}

// SetVaultNotificationsRequest is the request type for the
// SetVaultNotifications API operation.
type SetVaultNotificationsRequest struct {
	*aws.Request
	Input *types.SetVaultNotificationsInput
	Copy  func(*types.SetVaultNotificationsInput) SetVaultNotificationsRequest
}

// Send marshals and sends the SetVaultNotifications API request.
func (r SetVaultNotificationsRequest) Send(ctx context.Context) (*SetVaultNotificationsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &SetVaultNotificationsResponse{
		SetVaultNotificationsOutput: r.Request.Data.(*types.SetVaultNotificationsOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// SetVaultNotificationsResponse is the response type for the
// SetVaultNotifications API operation.
type SetVaultNotificationsResponse struct {
	*types.SetVaultNotificationsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// SetVaultNotifications request.
func (r *SetVaultNotificationsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
