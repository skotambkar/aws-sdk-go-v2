// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package quicksight

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
)

const opCreateGroupMembership = "CreateGroupMembership"

// CreateGroupMembershipRequest returns a request value for making API operation for
// Amazon QuickSight.
//
// Adds an Amazon QuickSight user to an Amazon QuickSight group.
//
// The permissions resource is arn:aws:quicksight:us-east-1:<aws-account-id>:group/default/<group-name> .
//
// The condition resource is the user name.
//
// The condition key is quicksight:UserName.
//
// The response is the group member object.
//
// CLI Sample:
//
// aws quicksight create-group-membership --aws-account-id=111122223333 --namespace=default
// --group-name=Sales --member-name=Pat
//
//    // Example sending a request using CreateGroupMembershipRequest.
//    req := client.CreateGroupMembershipRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/quicksight-2018-04-01/CreateGroupMembership
func (c *Client) CreateGroupMembershipRequest(input *types.CreateGroupMembershipInput) CreateGroupMembershipRequest {
	op := &aws.Operation{
		Name:       opCreateGroupMembership,
		HTTPMethod: "PUT",
		HTTPPath:   "/accounts/{AwsAccountId}/namespaces/{Namespace}/groups/{GroupName}/members/{MemberName}",
	}

	if input == nil {
		input = &types.CreateGroupMembershipInput{}
	}

	req := c.newRequest(op, input, &types.CreateGroupMembershipOutput{})
	return CreateGroupMembershipRequest{Request: req, Input: input, Copy: c.CreateGroupMembershipRequest}
}

// CreateGroupMembershipRequest is the request type for the
// CreateGroupMembership API operation.
type CreateGroupMembershipRequest struct {
	*aws.Request
	Input *types.CreateGroupMembershipInput
	Copy  func(*types.CreateGroupMembershipInput) CreateGroupMembershipRequest
}

// Send marshals and sends the CreateGroupMembership API request.
func (r CreateGroupMembershipRequest) Send(ctx context.Context) (*CreateGroupMembershipResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateGroupMembershipResponse{
		CreateGroupMembershipOutput: r.Request.Data.(*types.CreateGroupMembershipOutput),
		response:                    &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateGroupMembershipResponse is the response type for the
// CreateGroupMembership API operation.
type CreateGroupMembershipResponse struct {
	*types.CreateGroupMembershipOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateGroupMembership request.
func (r *CreateGroupMembershipResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
