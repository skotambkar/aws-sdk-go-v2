// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sts

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/sts/internal/aws_query"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
)

const opAssumeRole = "AssumeRole"

// AssumeRoleRequest returns a request value for making API operation for
// AWS Security Token Service.
//
// Returns a set of temporary security credentials that you can use to access
// AWS resources that you might not normally have access to. These temporary
// credentials consist of an access key ID, a secret access key, and a security
// token. Typically, you use AssumeRole within your account or for cross-account
// access. For a comparison of AssumeRole with other API operations that produce
// temporary credentials, see Requesting Temporary Security Credentials (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html)
// and Comparing the AWS STS API operations (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_credentials_temp_request.html#stsapi_comparison)
// in the IAM User Guide.
//
// You cannot use AWS account root user credentials to call AssumeRole. You
// must use credentials for an IAM user or an IAM role to call AssumeRole.
//
// For cross-account access, imagine that you own multiple accounts and need
// to access resources in each account. You could create long-term credentials
// in each account to access those resources. However, managing all those credentials
// and remembering which one can access which account can be time consuming.
// Instead, you can create one set of long-term credentials in one account.
// Then use temporary security credentials to access all the other accounts
// by assuming roles in those accounts. For more information about roles, see
// IAM Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles.html)
// in the IAM User Guide.
//
// By default, the temporary security credentials created by AssumeRole last
// for one hour. However, you can use the optional DurationSeconds parameter
// to specify the duration of your session. You can provide a value from 900
// seconds (15 minutes) up to the maximum session duration setting for the role.
// This setting can have a value from 1 hour to 12 hours. To learn how to view
// the maximum value for your role, see View the Maximum Session Duration Setting
// for a Role (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html#id_roles_use_view-role-max-session)
// in the IAM User Guide. The maximum session duration limit applies when you
// use the AssumeRole* API operations or the assume-role* CLI commands. However
// the limit does not apply when you use those operations to create a console
// URL. For more information, see Using IAM Roles (https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_use.html)
// in the IAM User Guide.
//
// The temporary security credentials created by AssumeRole can be used to make
// API calls to any AWS service with the following exception: You cannot call
// the AWS STS GetFederationToken or GetSessionToken API operations.
//
// (Optional) You can pass inline or managed session policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
// to this operation. You can pass a single JSON policy document to use as an
// inline session policy. You can also specify up to 10 managed policies to
// use as managed session policies. The plain text that you use for both inline
// and managed session policies shouldn't exceed 2048 characters. Passing policies
// to this operation returns new temporary credentials. The resulting session's
// permissions are the intersection of the role's identity-based policy and
// the session policies. You can use the role's temporary credentials in subsequent
// AWS API calls to access resources in the account that owns the role. You
// cannot use session policies to grant more permissions than those allowed
// by the identity-based policy of the role that is being assumed. For more
// information, see Session Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#policies_session)
// in the IAM User Guide.
//
// To assume a role from a different account, your AWS account must be trusted
// by the role. The trust relationship is defined in the role's trust policy
// when the role is created. That trust policy states which accounts are allowed
// to delegate that access to users in the account.
//
// A user who wants to access a role in a different account must also have permissions
// that are delegated from the user account administrator. The administrator
// must attach a policy that allows the user to call AssumeRole for the ARN
// of the role in the other account. If the user is in the same account as the
// role, then you can do either of the following:
//
//    * Attach a policy to the user (identical to the previous user in a different
//    account).
//
//    * Add the user as a principal directly in the role's trust policy.
//
// In this case, the trust policy acts as an IAM resource-based policy. Users
// in the same account as the role do not need explicit permission to assume
// the role. For more information about trust policies and resource-based policies,
// see IAM Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html)
// in the IAM User Guide.
//
// Using MFA with AssumeRole
//
// (Optional) You can include multi-factor authentication (MFA) information
// when you call AssumeRole. This is useful for cross-account scenarios to ensure
// that the user that assumes the role has been authenticated with an AWS MFA
// device. In that scenario, the trust policy of the role being assumed includes
// a condition that tests for MFA authentication. If the caller does not include
// valid MFA information, the request to assume the role is denied. The condition
// in a trust policy that tests for MFA authentication might look like the following
// example.
//
// "Condition": {"Bool": {"aws:MultiFactorAuthPresent": true}}
//
// For more information, see Configuring MFA-Protected API Access (https://docs.aws.amazon.com/IAM/latest/UserGuide/MFAProtectedAPI.html)
// in the IAM User Guide guide.
//
// To use MFA with AssumeRole, you pass values for the SerialNumber and TokenCode
// parameters. The SerialNumber value identifies the user's hardware or virtual
// MFA device. The TokenCode is the time-based one-time password (TOTP) that
// the MFA device produces.
//
//    // Example sending a request using AssumeRoleRequest.
//    req := client.AssumeRoleRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/sts-2011-06-15/AssumeRole
func (c *Client) AssumeRoleRequest(input *types.AssumeRoleInput) AssumeRoleRequest {
	op := &aws.Operation{
		Name:       opAssumeRole,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &types.AssumeRoleInput{}
	}

	req := c.newRequest(op, input, &types.AssumeRoleOutput{})

	// swap existing build handler on svc, with a new named build handler
	req.Handlers.Build.Swap(query.BuildHandler.Name, aws_query.AssumeRoleMarshaler{Input: input}.GetNamedBuildHandler())

	return AssumeRoleRequest{Request: req, Input: input, Copy: c.AssumeRoleRequest}
}

// AssumeRoleRequest is the request type for the
// AssumeRole API operation.
type AssumeRoleRequest struct {
	*aws.Request
	Input *types.AssumeRoleInput
	Copy  func(*types.AssumeRoleInput) AssumeRoleRequest
}

// Send marshals and sends the AssumeRole API request.
func (r AssumeRoleRequest) Send(ctx context.Context) (*AssumeRoleResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &AssumeRoleResponse{
		AssumeRoleOutput: r.Request.Data.(*types.AssumeRoleOutput),
		response:         &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// AssumeRoleResponse is the response type for the
// AssumeRole API operation.
type AssumeRoleResponse struct {
	*types.AssumeRoleOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// AssumeRole request.
func (r *AssumeRoleResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
