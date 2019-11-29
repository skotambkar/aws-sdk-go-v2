// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/internal/aws_restjson"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opDeleteGroupMembership = "DeleteGroupMembership"

// DeleteGroupMembershipRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Removes a user from a group so that the user is no longer a member of the
// group.
//
// The permissions resource is arn:aws:quicksight:us-east-1:<aws-account-id>:group/default/<group-name> .
//
// The condition resource is the user name.
//
// The condition key is quicksight:UserName.
//
// CLI Sample:
//
// aws quicksight delete-group-membership --aws-account-id=111122223333 --namespace=default
// --group-name=Sales-Management --member-name=Charlie
//
//    // Example sending a request using DeleteGroupMembershipRequest.
//    req := client.DeleteGroupMembershipRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/DeleteGroupMembership
func (c *Client) DeleteGroupMembershipRequest(input *types.DeleteGroupMembershipInput) DeleteGroupMembershipRequest {
	op := &aws.Operation{
		Name:       opDeleteGroupMembership,
		HTTPMethod: "DELETE",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/groups/{GroupName}/members/{MemberName}",
	}

	if input == nil {
		input = &types.DeleteGroupMembershipInput{}
	}

	req := c.newRequest(op, input, &types.DeleteGroupMembershipOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(restjson.BuildHandler.Name, aws_restjson.DeleteGroupMembershipMarshaler{Input: input}.GetNamedBuildHandler())

	return DeleteGroupMembershipRequest{Request: req, Input: input, Copy: c.DeleteGroupMembershipRequest}
}

// DeleteGroupMembershipRequest is the request type for the
// DeleteGroupMembership API operation.
type DeleteGroupMembershipRequest struct {
	*aws.Request
	Input *types.DeleteGroupMembershipInput
	Copy  func(*types.DeleteGroupMembershipInput) DeleteGroupMembershipRequest
}

// Send marshals and sends the DeleteGroupMembership API request.
func (r DeleteGroupMembershipRequest) Send(ctx context.Context) (*DeleteGroupMembershipResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DeleteGroupMembershipResponse{
		DeleteGroupMembershipOutput: r.Request.Data.(*types.DeleteGroupMembershipOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DeleteGroupMembershipResponse is the response type for the
// DeleteGroupMembership API operation.
type DeleteGroupMembershipResponse struct {
	*types.DeleteGroupMembershipOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DeleteGroupMembership request.
func (r *DeleteGroupMembershipResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
