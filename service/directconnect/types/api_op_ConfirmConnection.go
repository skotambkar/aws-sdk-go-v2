// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/enums"
)

type ConfirmConnectionInput struct {
	_ struct{} `type:"structure"`

	// The ID of the hosted connection.
	//
	// ConnectionId is a required field
	ConnectionId *string `locationName:"connectionId" type:"string" required:"true"`
}

// String returns the string representation
func (s ConfirmConnectionInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConfirmConnectionInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConfirmConnectionInput"}

	if s.ConnectionId == nil {
		invalidParams.Add(aws.NewErrParamRequired("ConnectionId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ConfirmConnectionOutput struct {
	_ struct{} `type:"structure"`

	// The state of the connection. The following are the possible values:
	//
	//    * ordering: The initial state of a hosted connection provisioned on an
	//    interconnect. The connection stays in the ordering state until the owner
	//    of the hosted connection confirms or declines the connection order.
	//
	//    * requested: The initial state of a standard connection. The connection
	//    stays in the requested state until the Letter of Authorization (LOA) is
	//    sent to the customer.
	//
	//    * pending: The connection has been approved and is being initialized.
	//
	//    * available: The network link is up and the connection is ready for use.
	//
	//    * down: The network link is down.
	//
	//    * deleting: The connection is being deleted.
	//
	//    * deleted: The connection has been deleted.
	//
	//    * rejected: A hosted connection in the ordering state enters the rejected
	//    state if it is deleted by the customer.
	//
	//    * unknown: The state of the connection is not available.
	ConnectionState enums.ConnectionState `locationName:"connectionState" type:"string" enum:"true"`
}

// String returns the string representation
func (s ConfirmConnectionOutput) String() string {
	return awsutil.Prettify(s)
}
