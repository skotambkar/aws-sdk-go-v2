// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type ConnectCustomKeyStoreInput struct {
	_ struct{} `type:"structure"`

	// Enter the key store ID of the custom key store that you want to connect.
	// To find the ID of a custom key store, use the DescribeCustomKeyStores operation.
	//
	// CustomKeyStoreId is a required field
	CustomKeyStoreId *string `min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s ConnectCustomKeyStoreInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConnectCustomKeyStoreInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConnectCustomKeyStoreInput"}

	if s.CustomKeyStoreId == nil {
		invalidParams.Add(aws.NewErrParamRequired("CustomKeyStoreId"))
	}
	if s.CustomKeyStoreId != nil && len(*s.CustomKeyStoreId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("CustomKeyStoreId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ConnectCustomKeyStoreOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ConnectCustomKeyStoreOutput) String() string {
	return awsutil.Prettify(s)
}