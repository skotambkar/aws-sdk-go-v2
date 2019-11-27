// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/enums"
)

type UpdateUserInput struct {
	_ struct{} `type:"structure"`

	// The ID for the AWS account that the user is in. Currently, you use the ID
	// for the AWS account that contains your Amazon QuickSight account.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The email address of the user that you want to update.
	//
	// Email is a required field
	Email *string `type:"string" required:"true"`

	// The namespace. Currently, you should set this to default.
	//
	// Namespace is a required field
	Namespace *string `location:"uri" locationName:"Namespace" type:"string" required:"true"`

	// The Amazon QuickSight role of the user. The user role can be one of the following:
	//
	//    * READER: A user who has read-only access to dashboards.
	//
	//    * AUTHOR: A user who can create data sources, data sets, analyses, and
	//    dashboards.
	//
	//    * ADMIN: A user who is an author, who can also manage Amazon QuickSight
	//    settings.
	//
	// Role is a required field
	Role enums.UserRole `type:"string" required:"true" enum:"true"`

	// The Amazon QuickSight user name that you want to update.
	//
	// UserName is a required field
	UserName *string `location:"uri" locationName:"UserName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateUserInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateUserInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateUserInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.Email == nil {
		invalidParams.Add(aws.NewErrParamRequired("Email"))
	}

	if s.Namespace == nil {
		invalidParams.Add(aws.NewErrParamRequired("Namespace"))
	}
	if len(s.Role) == 0 {
		invalidParams.Add(aws.NewErrParamRequired("Role"))
	}

	if s.UserName == nil {
		invalidParams.Add(aws.NewErrParamRequired("UserName"))
	}
	if s.UserName != nil && len(*s.UserName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("UserName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateUserOutput struct {
	_ struct{} `type:"structure"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`

	// The Amazon QuickSight user.
	User *User `type:"structure"`
}

// String returns the string representation
func (s UpdateUserOutput) String() string {
	return awsutil.Prettify(s)
}
