// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package types

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/service/directconnect/enums"
)

type ConfirmPrivateVirtualInterfaceInput struct {
	_ struct{} `type:"structure"`

	// The ID of the Direct Connect gateway.
	DirectConnectGatewayId *string `locationName:"directConnectGatewayId" type:"string"`

	// The ID of the virtual private gateway.
	VirtualGatewayId *string `locationName:"virtualGatewayId" type:"string"`

	// The ID of the virtual interface.
	//
	// VirtualInterfaceId is a required field
	VirtualInterfaceId *string `locationName:"virtualInterfaceId" type:"string" required:"true"`
}

// String returns the string representation
func (s ConfirmPrivateVirtualInterfaceInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ConfirmPrivateVirtualInterfaceInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "ConfirmPrivateVirtualInterfaceInput"}

	if s.VirtualInterfaceId == nil {
		invalidParams.Add(aws.NewErrParamRequired("VirtualInterfaceId"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type ConfirmPrivateVirtualInterfaceOutput struct {
	_ struct{} `type:"structure"`

	// The state of the virtual interface. The following are the possible values:
	//
	//    * confirming: The creation of the virtual interface is pending confirmation
	//    from the virtual interface owner. If the owner of the virtual interface
	//    is different from the owner of the connection on which it is provisioned,
	//    then the virtual interface will remain in this state until it is confirmed
	//    by the virtual interface owner.
	//
	//    * verifying: This state only applies to public virtual interfaces. Each
	//    public virtual interface needs validation before the virtual interface
	//    can be created.
	//
	//    * pending: A virtual interface is in this state from the time that it
	//    is created until the virtual interface is ready to forward traffic.
	//
	//    * available: A virtual interface that is able to forward traffic.
	//
	//    * down: A virtual interface that is BGP down.
	//
	//    * deleting: A virtual interface is in this state immediately after calling
	//    DeleteVirtualInterface until it can no longer forward traffic.
	//
	//    * deleted: A virtual interface that cannot forward traffic.
	//
	//    * rejected: The virtual interface owner has declined creation of the virtual
	//    interface. If a virtual interface in the Confirming state is deleted by
	//    the virtual interface owner, the virtual interface enters the Rejected
	//    state.
	//
	//    * unknown: The state of the virtual interface is not available.
	VirtualInterfaceState enums.VirtualInterfaceState `locationName:"virtualInterfaceState" type:"string" enum:"true"`
}

// String returns the string representation
func (s ConfirmPrivateVirtualInterfaceOutput) String() string {
	return awsutil.Prettify(s)
}
