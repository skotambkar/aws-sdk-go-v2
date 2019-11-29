// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type UpdateDataSetPermissionsInput struct {
	_ struct{} `type:"structure"`

	// The AWS Account ID.
	//
	// AwsAccountId is a required field
	AwsAccountId *string `location:"uri" locationName:"AwsAccountId" min:"12" type:"string" required:"true"`

	// The ID for the dataset you want to create. This is unique per region per
	// AWS account.
	//
	// DataSetId is a required field
	DataSetId *string `location:"uri" locationName:"DataSetId" type:"string" required:"true"`

	// The resource permissions that you want to grant to the dataset.
	GrantPermissions []ResourcePermission `min:"1" type:"list"`

	// The resource permissions that you want to revoke from the dataset.
	RevokePermissions []ResourcePermission `min:"1" type:"list"`
}

// String returns the string representation
func (s UpdateDataSetPermissionsInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateDataSetPermissionsInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "UpdateDataSetPermissionsInput"}

	if s.AwsAccountId == nil {
		invalidParams.Add(aws.NewErrParamRequired("AwsAccountId"))
	}
	if s.AwsAccountId != nil && len(*s.AwsAccountId) < 12 {
		invalidParams.Add(aws.NewErrParamMinLen("AwsAccountId", 12))
	}

	if s.DataSetId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DataSetId"))
	}
	if s.GrantPermissions != nil && len(s.GrantPermissions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("GrantPermissions", 1))
	}
	if s.RevokePermissions != nil && len(s.RevokePermissions) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("RevokePermissions", 1))
	}
	if s.GrantPermissions != nil {
		for i, v := range s.GrantPermissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "GrantPermissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}
	if s.RevokePermissions != nil {
		for i, v := range s.RevokePermissions {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "RevokePermissions", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type UpdateDataSetPermissionsOutput struct {
	_ struct{} `type:"structure"`

	// The ARN of the dataset.
	DataSetArn *string `type:"string"`

	// The ID for the dataset you want to create. This is unique per region per
	// AWS account.
	DataSetId *string `type:"string"`

	// The AWS request ID for this operation.
	RequestId *string `type:"string"`

	// The http status of the request.
	Status *int64 `location:"statusCode" type:"integer"`
}

// String returns the string representation
func (s UpdateDataSetPermissionsOutput) String() string {
	return awsutil.Prettify(s)
}