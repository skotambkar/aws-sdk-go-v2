// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opUpdateTemplatePermissions = "UpdateTemplatePermissions"

// UpdateTemplatePermissionsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Updates the permissions on a template.
//
// CLI syntax:
//
//    * aws quicksight describe-template-permissions —aws-account-id 111122223333
//    —template-id reports_test_template
//
//    * aws quicksight update-template-permissions —cli-input-json file://update-permission.json
//
//    * The structure of update-permissions.json to add permissions: { "AwsAccountId":
//    "111122223333", "DashboardId": "reports_test_template", "GrantPermissions":
//    [ { "Principal": "arn:aws:quicksight:us-east-1:196359894473:user/default/user3",
//    "Actions": [ "quicksight:DescribeTemplate", "quicksight:ListTemplateVersions"
//    ] } ] } The structure of update-permissions.json to add permissions: {
//    "AwsAccountId": "111122223333", "DashboardId": "reports_test_template",
//    "RevokePermissions": [ { "Principal": "arn:aws:quicksight:us-east-1:196359894473:user/default/user3",
//    "Actions": [ "quicksight:DescribeTemplate", "quicksight:ListTemplateVersions"
//    ] } ] } To obtain the principal name of a QuickSight group or user, use
//    user describe-group or describe-user. For example: aws quicksight describe-user
//    --aws-account-id 111122223333 --namespace default --user-name user2 --region
//    us-east-1 { "User": { "Arn": "arn:aws:quicksight:us-east-1:111122223333:user/default/user2",
//    "Active": true, "Email": "user2@example.com", "Role": "ADMIN", "UserName":
//    "user2", "PrincipalId": "federated/iam/abcd2abcdabcdeabc5ab5" }, "RequestId":
//    "8f74bb31-6291-448a-a71c-a765a44bae31", "Status": 200 }
//
//    // Example sending a request using UpdateTemplatePermissionsRequest.
//    req := client.UpdateTemplatePermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateTemplatePermissions
func (c *Client) UpdateTemplatePermissionsRequest(input *types.UpdateTemplatePermissionsInput) UpdateTemplatePermissionsRequest {
	op := &aws.Operation{
		Name:       opUpdateTemplatePermissions,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/templates/{TemplateId}/permissions",
	}

	if input == nil {
		input = &types.UpdateTemplatePermissionsInput{}
	}

	req := c.newRequest(op, input, &types.UpdateTemplatePermissionsOutput{})
	return UpdateTemplatePermissionsRequest{Request: req, Input: input, Copy: c.UpdateTemplatePermissionsRequest}
}

// UpdateTemplatePermissionsRequest is the request type for the
// UpdateTemplatePermissions API operation.
type UpdateTemplatePermissionsRequest struct {
	*aws.Request
	Input *types.UpdateTemplatePermissionsInput
	Copy  func(*types.UpdateTemplatePermissionsInput) UpdateTemplatePermissionsRequest
}

// Send marshals and sends the UpdateTemplatePermissions API request.
func (r UpdateTemplatePermissionsRequest) Send(ctx context.Context) (*UpdateTemplatePermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateTemplatePermissionsResponse{
		UpdateTemplatePermissionsOutput: r.Request.Data.(*types.UpdateTemplatePermissionsOutput),
		response:                        &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateTemplatePermissionsResponse is the response type for the
// UpdateTemplatePermissions API operation.
type UpdateTemplatePermissionsResponse struct {
	*types.UpdateTemplatePermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateTemplatePermissions request.
func (r *UpdateTemplatePermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
