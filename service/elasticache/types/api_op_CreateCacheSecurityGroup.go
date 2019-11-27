// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Represents the input of a CreateCacheSecurityGroup operation.
type CreateCacheSecurityGroupInput struct {
	_ struct{} `type:"structure"`

	// A name for the cache security group. This value is stored as a lowercase
	// string.
	//
	// Constraints: Must contain no more than 255 alphanumeric characters. Cannot
	// be the word "Default".
	//
	// Example: mysecuritygroup
	//
	// CacheSecurityGroupName is a required field
	CacheSecurityGroupName *string `type:"string" required:"true"`

	// A description for the cache security group.
	//
	// Description is a required field
	Description *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateCacheSecurityGroupInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateCacheSecurityGroupInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateCacheSecurityGroupInput"}

	if s.CacheSecurityGroupName == nil {
		invalidParams.Add(aws.NewErrParamRequired("CacheSecurityGroupName"))
	}

	if s.Description == nil {
		invalidParams.Add(aws.NewErrParamRequired("Description"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type CreateCacheSecurityGroupOutput struct {
	_ struct{} `type:"structure"`

	// Represents the output of one of the following operations:
	//
	//    * AuthorizeCacheSecurityGroupIngress
	//
	//    * CreateCacheSecurityGroup
	//
	//    * RevokeCacheSecurityGroupIngress
	CacheSecurityGroup *CacheSecurityGroup `type:"structure"`
}

// String returns the string representation
func (s CreateCacheSecurityGroupOutput) String() string {
	return awsutil.Prettify(s)
}
