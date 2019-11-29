// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package aws_restjson service, marshal.go contains the generated operation marshalers and methods on it.

package aws_restjson

import (
	"bytes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/private/protocol/rest"
	"github.com/aws/aws-sdk-go-v2/service/ram/types"
)

var _ bytes.Buffer
var _ awserr.Error

// AcceptResourceShareInvitationMarshaler defines marshaler for AcceptResourceShareInvitation operation
type AcceptResourceShareInvitationMarshaler struct {
	Input *types.AcceptResourceShareInvitationInput
}

// AssociateResourceShareMarshaler defines marshaler for AssociateResourceShare operation
type AssociateResourceShareMarshaler struct {
	Input *types.AssociateResourceShareInput
}

// CreateResourceShareMarshaler defines marshaler for CreateResourceShare operation
type CreateResourceShareMarshaler struct {
	Input *types.CreateResourceShareInput
}

// DeleteResourceShareMarshaler defines marshaler for DeleteResourceShare operation
type DeleteResourceShareMarshaler struct {
	Input *types.DeleteResourceShareInput
}

// DisassociateResourceShareMarshaler defines marshaler for DisassociateResourceShare operation
type DisassociateResourceShareMarshaler struct {
	Input *types.DisassociateResourceShareInput
}

// EnableSharingWithAwsOrganizationMarshaler defines marshaler for EnableSharingWithAwsOrganization operation
type EnableSharingWithAwsOrganizationMarshaler struct {
	Input *types.EnableSharingWithAwsOrganizationInput
}

// GetResourcePoliciesMarshaler defines marshaler for GetResourcePolicies operation
type GetResourcePoliciesMarshaler struct {
	Input *types.GetResourcePoliciesInput
}

// GetResourceShareAssociationsMarshaler defines marshaler for GetResourceShareAssociations operation
type GetResourceShareAssociationsMarshaler struct {
	Input *types.GetResourceShareAssociationsInput
}

// GetResourceShareInvitationsMarshaler defines marshaler for GetResourceShareInvitations operation
type GetResourceShareInvitationsMarshaler struct {
	Input *types.GetResourceShareInvitationsInput
}

// GetResourceSharesMarshaler defines marshaler for GetResourceShares operation
type GetResourceSharesMarshaler struct {
	Input *types.GetResourceSharesInput
}

// ListPendingInvitationResourcesMarshaler defines marshaler for ListPendingInvitationResources operation
type ListPendingInvitationResourcesMarshaler struct {
	Input *types.ListPendingInvitationResourcesInput
}

// ListPrincipalsMarshaler defines marshaler for ListPrincipals operation
type ListPrincipalsMarshaler struct {
	Input *types.ListPrincipalsInput
}

// ListResourcesMarshaler defines marshaler for ListResources operation
type ListResourcesMarshaler struct {
	Input *types.ListResourcesInput
}

// RejectResourceShareInvitationMarshaler defines marshaler for RejectResourceShareInvitation operation
type RejectResourceShareInvitationMarshaler struct {
	Input *types.RejectResourceShareInvitationInput
}

// TagResourceMarshaler defines marshaler for TagResource operation
type TagResourceMarshaler struct {
	Input *types.TagResourceInput
}

// UntagResourceMarshaler defines marshaler for UntagResource operation
type UntagResourceMarshaler struct {
	Input *types.UntagResourceInput
}

// UpdateResourceShareMarshaler defines marshaler for UpdateResourceShare operation
type UpdateResourceShareMarshaler struct {
	Input *types.UpdateResourceShareInput
}

func MarshalAcceptResourceShareInvitationInputShapeAWSREST(v *types.AcceptResourceShareInvitationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalAcceptResourceShareInvitationInputShapeAWSJSON(v *types.AcceptResourceShareInvitationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalAssociateResourceShareInputShapeAWSREST(v *types.AssociateResourceShareInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalAssociateResourceShareInputShapeAWSJSON(v *types.AssociateResourceShareInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalCreateResourceShareInputShapeAWSREST(v *types.CreateResourceShareInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalCreateResourceShareInputShapeAWSJSON(v *types.CreateResourceShareInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalDeleteResourceShareInputShapeAWSREST(v *types.DeleteResourceShareInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalDeleteResourceShareInputShapeAWSJSON(v *types.DeleteResourceShareInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalDisassociateResourceShareInputShapeAWSREST(v *types.DisassociateResourceShareInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalDisassociateResourceShareInputShapeAWSJSON(v *types.DisassociateResourceShareInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalEnableSharingWithAwsOrganizationInputShapeAWSREST(v *types.EnableSharingWithAwsOrganizationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalEnableSharingWithAwsOrganizationInputShapeAWSJSON(v *types.EnableSharingWithAwsOrganizationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalGetResourcePoliciesInputShapeAWSREST(v *types.GetResourcePoliciesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalGetResourcePoliciesInputShapeAWSJSON(v *types.GetResourcePoliciesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalGetResourceShareAssociationsInputShapeAWSREST(v *types.GetResourceShareAssociationsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalGetResourceShareAssociationsInputShapeAWSJSON(v *types.GetResourceShareAssociationsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalGetResourceShareInvitationsInputShapeAWSREST(v *types.GetResourceShareInvitationsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalGetResourceShareInvitationsInputShapeAWSJSON(v *types.GetResourceShareInvitationsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalGetResourceSharesInputShapeAWSREST(v *types.GetResourceSharesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalGetResourceSharesInputShapeAWSJSON(v *types.GetResourceSharesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListPendingInvitationResourcesInputShapeAWSREST(v *types.ListPendingInvitationResourcesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListPendingInvitationResourcesInputShapeAWSJSON(v *types.ListPendingInvitationResourcesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListPrincipalsInputShapeAWSREST(v *types.ListPrincipalsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListPrincipalsInputShapeAWSJSON(v *types.ListPrincipalsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListResourcesInputShapeAWSREST(v *types.ListResourcesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListResourcesInputShapeAWSJSON(v *types.ListResourcesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalRejectResourceShareInvitationInputShapeAWSREST(v *types.RejectResourceShareInvitationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalRejectResourceShareInvitationInputShapeAWSJSON(v *types.RejectResourceShareInvitationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalTagResourceInputShapeAWSREST(v *types.TagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalTagResourceInputShapeAWSJSON(v *types.TagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalUntagResourceInputShapeAWSREST(v *types.UntagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUntagResourceInputShapeAWSJSON(v *types.UntagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalUpdateResourceShareInputShapeAWSREST(v *types.UpdateResourceShareInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUpdateResourceShareInputShapeAWSJSON(v *types.UpdateResourceShareInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}
func (m AcceptResourceShareInvitationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalAcceptResourceShareInvitationInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalAcceptResourceShareInvitationInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m AcceptResourceShareInvitationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "AcceptResourceShareInvitation.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m AssociateResourceShareMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalAssociateResourceShareInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalAssociateResourceShareInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m AssociateResourceShareMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "AssociateResourceShare.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreateResourceShareMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateResourceShareInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalCreateResourceShareInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateResourceShareMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateResourceShare.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteResourceShareMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteResourceShareInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalDeleteResourceShareInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteResourceShareMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteResourceShare.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DisassociateResourceShareMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDisassociateResourceShareInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalDisassociateResourceShareInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DisassociateResourceShareMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DisassociateResourceShare.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m EnableSharingWithAwsOrganizationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalEnableSharingWithAwsOrganizationInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalEnableSharingWithAwsOrganizationInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m EnableSharingWithAwsOrganizationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "EnableSharingWithAwsOrganization.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetResourcePoliciesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetResourcePoliciesInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalGetResourcePoliciesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetResourcePoliciesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetResourcePolicies.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetResourceShareAssociationsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetResourceShareAssociationsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalGetResourceShareAssociationsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetResourceShareAssociationsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetResourceShareAssociations.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetResourceShareInvitationsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetResourceShareInvitationsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalGetResourceShareInvitationsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetResourceShareInvitationsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetResourceShareInvitations.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetResourceSharesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetResourceSharesInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalGetResourceSharesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetResourceSharesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetResourceShares.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListPendingInvitationResourcesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListPendingInvitationResourcesInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListPendingInvitationResourcesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListPendingInvitationResourcesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListPendingInvitationResources.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListPrincipalsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListPrincipalsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListPrincipalsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListPrincipalsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListPrincipals.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListResourcesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListResourcesInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListResourcesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListResourcesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListResources.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m RejectResourceShareInvitationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalRejectResourceShareInvitationInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalRejectResourceShareInvitationInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m RejectResourceShareInvitationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "RejectResourceShareInvitation.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m TagResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalTagResourceInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalTagResourceInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m TagResourceMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "TagResource.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UntagResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUntagResourceInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUntagResourceInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UntagResourceMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UntagResource.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateResourceShareMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateResourceShareInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUpdateResourceShareInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateResourceShareMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateResourceShare.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}
