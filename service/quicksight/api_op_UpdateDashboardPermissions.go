// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opUpdateDashboardPermissions = "UpdateDashboardPermissions"

// UpdateDashboardPermissionsRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Updates read and write permissions on a dashboard.
//
// CLI syntax:
//
// aws quicksight update-dashboard-permissions —cli-input-json file://update-permission.json
//
// A sample update-permissions.json for granting read only permissions:
//
// { "AwsAccountId": "111122223333", "DashboardId": "reports_test_report", "GrantPermissions":
// [ { "Principal": "arn:aws:quicksight:us-east-1:111122223333:user/default/user2",
// "Actions": [ "quicksight:DescribeDashboard", "quicksight:ListDashboardVersions",
// "quicksight:DescribeDashboardVersion", "quicksight:QueryDashboard" ] } ]
// }
//
// A sample update-permissions.json for granting read and write permissions:
//
// { "AwsAccountId": "111122223333", "DashboardId": "reports_test_report", "GrantPermissions":
// [ { "Principal": "arn:aws:quicksight:us-east-1:111122223333:user/default/user2",
// "Actions": [ "quicksight:DescribeDashboard", "quicksight:ListDashboardVersions",
// "quicksight:DescribeDashboardVersion", "quicksight:QueryDashboard", "quicksight:DescribeDashboardPermissions",
// "quicksight:UpdateDashboardPermissions", "quicksight:DeleteDashboardVersion",
// "quicksight:DeleteDashboard", "quicksight:UpdateDashboard", "quicksight:UpdateDashboardPublishedVersion",
// ] } ] }
//
// A sample update-permissions.json for revoking write permissions:
//
// { "AwsAccountId": "111122223333", "DashboardId": "reports_test_report", "RevokePermissions":
// [ { "Principal": "arn:aws:quicksight:us-east-1:111122223333:user/default/user2",
// "Actions": [ "quicksight:DescribeDashboardPermissions", "quicksight:UpdateDashboardPermissions",
// "quicksight:DeleteDashboardVersion", "quicksight:DeleteDashboard", "quicksight:UpdateDashboard",
// "quicksight:UpdateDashboardPublishedVersion", ] } ] }
//
// A sample update-permissions.json for revoking read and write permissions:
//
// { "AwsAccountId": "111122223333", "DashboardId": "reports_test_report", "RevokePermissions":
// [ { "Principal": "arn:aws:quicksight:us-east-1:111122223333:user/default/user2",
// "Actions": [ "quicksight:DescribeDashboard", "quicksight:ListDashboardVersions",
// "quicksight:DescribeDashboardVersion", "quicksight:QueryDashboard", "quicksight:DescribeDashboardPermissions",
// "quicksight:UpdateDashboardPermissions", "quicksight:DeleteDashboardVersion",
// "quicksight:DeleteDashboard", "quicksight:UpdateDashboard", "quicksight:UpdateDashboardPublishedVersion",
// ] } ] }
//
// To obtain the principal name of a QuickSight user or group, you can use describe-group
// or describe-user. For example:
//
// aws quicksight describe-user --aws-account-id 111122223333 --namespace default
// --user-name user2 --region us-east-1 { "User": { "Arn": "arn:aws:quicksight:us-east-1:111122223333:user/default/user2",
// "Active": true, "Email": "user2@example.com", "Role": "ADMIN", "UserName":
// "user2", "PrincipalId": "federated/iam/abcd2abcdabcdeabc5ab5" }, "RequestId":
// "8f74bb31-6291-448a-a71c-a765a44bae31", "Status": 200 }
//
//    // Example sending a request using UpdateDashboardPermissionsRequest.
//    req := client.UpdateDashboardPermissionsRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/UpdateDashboardPermissions
func (c *Client) UpdateDashboardPermissionsRequest(input *types.UpdateDashboardPermissionsInput) UpdateDashboardPermissionsRequest {
	op := &aws.Operation{
		Name:       opUpdateDashboardPermissions,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/dashboards/{DashboardId}/permissions",
	}

	if input == nil {
		input = &types.UpdateDashboardPermissionsInput{}
	}

	req := c.newRequest(op, input, &types.UpdateDashboardPermissionsOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.UpdateDashboardPermissionsMarshaler{Input: input}.GetNamedBuildHandler())

	return UpdateDashboardPermissionsRequest{Request: req, Input: input, Copy: c.UpdateDashboardPermissionsRequest}
}

// UpdateDashboardPermissionsRequest is the request type for the
// UpdateDashboardPermissions API operation.
type UpdateDashboardPermissionsRequest struct {
	*aws.Request
	Input *types.UpdateDashboardPermissionsInput
	Copy  func(*types.UpdateDashboardPermissionsInput) UpdateDashboardPermissionsRequest
}

// Send marshals and sends the UpdateDashboardPermissions API request.
func (r UpdateDashboardPermissionsRequest) Send(ctx context.Context) (*UpdateDashboardPermissionsResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &UpdateDashboardPermissionsResponse{
		UpdateDashboardPermissionsOutput: r.Request.Data.(*types.UpdateDashboardPermissionsOutput),
		response:                         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// UpdateDashboardPermissionsResponse is the response type for the
// UpdateDashboardPermissions API operation.
type UpdateDashboardPermissionsResponse struct {
	*types.UpdateDashboardPermissionsOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// UpdateDashboardPermissions request.
func (r *UpdateDashboardPermissionsResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
