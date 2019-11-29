// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package aws_jsonrpc service, marshal.go contains the generated operation marshalers and methods on it.

package aws_jsonrpc

import (
	"bytes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/marketplaceentitlementservice/types"
)

var _ bytes.Buffer
var _ awserr.Error

// GetEntitlementsMarshaler defines marshaler for GetEntitlements operation
type GetEntitlementsMarshaler struct {
	Input *types.GetEntitlementsInput
}

func MarshalGetEntitlementsInputShapeAWSJSON(v *types.GetEntitlementsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}
func (m GetEntitlementsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetEntitlementsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetEntitlementsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetEntitlements.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}
