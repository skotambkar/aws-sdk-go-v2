// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package aws_jsonrpc service, marshal.go contains the generated operation marshalers and methods on it.

package aws_jsonrpc

import (
	"bytes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
)

var _ bytes.Buffer
var _ awserr.Error

// AssociateIpGroupsMarshaler defines marshaler for AssociateIpGroups operation
type AssociateIpGroupsMarshaler struct {
	Input *types.AssociateIpGroupsInput
}

// AuthorizeIpRulesMarshaler defines marshaler for AuthorizeIpRules operation
type AuthorizeIpRulesMarshaler struct {
	Input *types.AuthorizeIpRulesInput
}

// CopyWorkspaceImageMarshaler defines marshaler for CopyWorkspaceImage operation
type CopyWorkspaceImageMarshaler struct {
	Input *types.CopyWorkspaceImageInput
}

// CreateIpGroupMarshaler defines marshaler for CreateIpGroup operation
type CreateIpGroupMarshaler struct {
	Input *types.CreateIpGroupInput
}

// CreateTagsMarshaler defines marshaler for CreateTags operation
type CreateTagsMarshaler struct {
	Input *types.CreateTagsInput
}

// CreateWorkspacesMarshaler defines marshaler for CreateWorkspaces operation
type CreateWorkspacesMarshaler struct {
	Input *types.CreateWorkspacesInput
}

// DeleteIpGroupMarshaler defines marshaler for DeleteIpGroup operation
type DeleteIpGroupMarshaler struct {
	Input *types.DeleteIpGroupInput
}

// DeleteTagsMarshaler defines marshaler for DeleteTags operation
type DeleteTagsMarshaler struct {
	Input *types.DeleteTagsInput
}

// DeleteWorkspaceImageMarshaler defines marshaler for DeleteWorkspaceImage operation
type DeleteWorkspaceImageMarshaler struct {
	Input *types.DeleteWorkspaceImageInput
}

// DeregisterWorkspaceDirectoryMarshaler defines marshaler for DeregisterWorkspaceDirectory operation
type DeregisterWorkspaceDirectoryMarshaler struct {
	Input *types.DeregisterWorkspaceDirectoryInput
}

// DescribeAccountMarshaler defines marshaler for DescribeAccount operation
type DescribeAccountMarshaler struct {
	Input *types.DescribeAccountInput
}

// DescribeAccountModificationsMarshaler defines marshaler for DescribeAccountModifications operation
type DescribeAccountModificationsMarshaler struct {
	Input *types.DescribeAccountModificationsInput
}

// DescribeClientPropertiesMarshaler defines marshaler for DescribeClientProperties operation
type DescribeClientPropertiesMarshaler struct {
	Input *types.DescribeClientPropertiesInput
}

// DescribeIpGroupsMarshaler defines marshaler for DescribeIpGroups operation
type DescribeIpGroupsMarshaler struct {
	Input *types.DescribeIpGroupsInput
}

// DescribeTagsMarshaler defines marshaler for DescribeTags operation
type DescribeTagsMarshaler struct {
	Input *types.DescribeTagsInput
}

// DescribeWorkspaceBundlesMarshaler defines marshaler for DescribeWorkspaceBundles operation
type DescribeWorkspaceBundlesMarshaler struct {
	Input *types.DescribeWorkspaceBundlesInput
}

// DescribeWorkspaceDirectoriesMarshaler defines marshaler for DescribeWorkspaceDirectories operation
type DescribeWorkspaceDirectoriesMarshaler struct {
	Input *types.DescribeWorkspaceDirectoriesInput
}

// DescribeWorkspaceImagesMarshaler defines marshaler for DescribeWorkspaceImages operation
type DescribeWorkspaceImagesMarshaler struct {
	Input *types.DescribeWorkspaceImagesInput
}

// DescribeWorkspaceSnapshotsMarshaler defines marshaler for DescribeWorkspaceSnapshots operation
type DescribeWorkspaceSnapshotsMarshaler struct {
	Input *types.DescribeWorkspaceSnapshotsInput
}

// DescribeWorkspacesMarshaler defines marshaler for DescribeWorkspaces operation
type DescribeWorkspacesMarshaler struct {
	Input *types.DescribeWorkspacesInput
}

// DescribeWorkspacesConnectionStatusMarshaler defines marshaler for DescribeWorkspacesConnectionStatus operation
type DescribeWorkspacesConnectionStatusMarshaler struct {
	Input *types.DescribeWorkspacesConnectionStatusInput
}

// DisassociateIpGroupsMarshaler defines marshaler for DisassociateIpGroups operation
type DisassociateIpGroupsMarshaler struct {
	Input *types.DisassociateIpGroupsInput
}

// ImportWorkspaceImageMarshaler defines marshaler for ImportWorkspaceImage operation
type ImportWorkspaceImageMarshaler struct {
	Input *types.ImportWorkspaceImageInput
}

// ListAvailableManagementCidrRangesMarshaler defines marshaler for ListAvailableManagementCidrRanges operation
type ListAvailableManagementCidrRangesMarshaler struct {
	Input *types.ListAvailableManagementCidrRangesInput
}

// ModifyAccountMarshaler defines marshaler for ModifyAccount operation
type ModifyAccountMarshaler struct {
	Input *types.ModifyAccountInput
}

// ModifyClientPropertiesMarshaler defines marshaler for ModifyClientProperties operation
type ModifyClientPropertiesMarshaler struct {
	Input *types.ModifyClientPropertiesInput
}

// ModifySelfservicePermissionsMarshaler defines marshaler for ModifySelfservicePermissions operation
type ModifySelfservicePermissionsMarshaler struct {
	Input *types.ModifySelfservicePermissionsInput
}

// ModifyWorkspaceAccessPropertiesMarshaler defines marshaler for ModifyWorkspaceAccessProperties operation
type ModifyWorkspaceAccessPropertiesMarshaler struct {
	Input *types.ModifyWorkspaceAccessPropertiesInput
}

// ModifyWorkspaceCreationPropertiesMarshaler defines marshaler for ModifyWorkspaceCreationProperties operation
type ModifyWorkspaceCreationPropertiesMarshaler struct {
	Input *types.ModifyWorkspaceCreationPropertiesInput
}

// ModifyWorkspacePropertiesMarshaler defines marshaler for ModifyWorkspaceProperties operation
type ModifyWorkspacePropertiesMarshaler struct {
	Input *types.ModifyWorkspacePropertiesInput
}

// ModifyWorkspaceStateMarshaler defines marshaler for ModifyWorkspaceState operation
type ModifyWorkspaceStateMarshaler struct {
	Input *types.ModifyWorkspaceStateInput
}

// RebootWorkspacesMarshaler defines marshaler for RebootWorkspaces operation
type RebootWorkspacesMarshaler struct {
	Input *types.RebootWorkspacesInput
}

// RebuildWorkspacesMarshaler defines marshaler for RebuildWorkspaces operation
type RebuildWorkspacesMarshaler struct {
	Input *types.RebuildWorkspacesInput
}

// RegisterWorkspaceDirectoryMarshaler defines marshaler for RegisterWorkspaceDirectory operation
type RegisterWorkspaceDirectoryMarshaler struct {
	Input *types.RegisterWorkspaceDirectoryInput
}

// RestoreWorkspaceMarshaler defines marshaler for RestoreWorkspace operation
type RestoreWorkspaceMarshaler struct {
	Input *types.RestoreWorkspaceInput
}

// RevokeIpRulesMarshaler defines marshaler for RevokeIpRules operation
type RevokeIpRulesMarshaler struct {
	Input *types.RevokeIpRulesInput
}

// StartWorkspacesMarshaler defines marshaler for StartWorkspaces operation
type StartWorkspacesMarshaler struct {
	Input *types.StartWorkspacesInput
}

// StopWorkspacesMarshaler defines marshaler for StopWorkspaces operation
type StopWorkspacesMarshaler struct {
	Input *types.StopWorkspacesInput
}

// TerminateWorkspacesMarshaler defines marshaler for TerminateWorkspaces operation
type TerminateWorkspacesMarshaler struct {
	Input *types.TerminateWorkspacesInput
}

// UpdateRulesOfIpGroupMarshaler defines marshaler for UpdateRulesOfIpGroup operation
type UpdateRulesOfIpGroupMarshaler struct {
	Input *types.UpdateRulesOfIpGroupInput
}

func MarshalAssociateIpGroupsInputShapeAWSJSON(v *types.AssociateIpGroupsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalAuthorizeIpRulesInputShapeAWSJSON(v *types.AuthorizeIpRulesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalCopyWorkspaceImageInputShapeAWSJSON(v *types.CopyWorkspaceImageInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalCreateIpGroupInputShapeAWSJSON(v *types.CreateIpGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalCreateTagsInputShapeAWSJSON(v *types.CreateTagsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalCreateWorkspacesInputShapeAWSJSON(v *types.CreateWorkspacesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDeleteIpGroupInputShapeAWSJSON(v *types.DeleteIpGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDeleteTagsInputShapeAWSJSON(v *types.DeleteTagsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDeleteWorkspaceImageInputShapeAWSJSON(v *types.DeleteWorkspaceImageInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDeregisterWorkspaceDirectoryInputShapeAWSJSON(v *types.DeregisterWorkspaceDirectoryInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeAccountInputShapeAWSJSON(v *types.DescribeAccountInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeAccountModificationsInputShapeAWSJSON(v *types.DescribeAccountModificationsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeClientPropertiesInputShapeAWSJSON(v *types.DescribeClientPropertiesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeIpGroupsInputShapeAWSJSON(v *types.DescribeIpGroupsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeTagsInputShapeAWSJSON(v *types.DescribeTagsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeWorkspaceBundlesInputShapeAWSJSON(v *types.DescribeWorkspaceBundlesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeWorkspaceDirectoriesInputShapeAWSJSON(v *types.DescribeWorkspaceDirectoriesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeWorkspaceImagesInputShapeAWSJSON(v *types.DescribeWorkspaceImagesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeWorkspaceSnapshotsInputShapeAWSJSON(v *types.DescribeWorkspaceSnapshotsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeWorkspacesInputShapeAWSJSON(v *types.DescribeWorkspacesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeWorkspacesConnectionStatusInputShapeAWSJSON(v *types.DescribeWorkspacesConnectionStatusInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDisassociateIpGroupsInputShapeAWSJSON(v *types.DisassociateIpGroupsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalImportWorkspaceImageInputShapeAWSJSON(v *types.ImportWorkspaceImageInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalListAvailableManagementCidrRangesInputShapeAWSJSON(v *types.ListAvailableManagementCidrRangesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalModifyAccountInputShapeAWSJSON(v *types.ModifyAccountInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalModifyClientPropertiesInputShapeAWSJSON(v *types.ModifyClientPropertiesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalModifySelfservicePermissionsInputShapeAWSJSON(v *types.ModifySelfservicePermissionsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalModifyWorkspaceAccessPropertiesInputShapeAWSJSON(v *types.ModifyWorkspaceAccessPropertiesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalModifyWorkspaceCreationPropertiesInputShapeAWSJSON(v *types.ModifyWorkspaceCreationPropertiesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalModifyWorkspacePropertiesInputShapeAWSJSON(v *types.ModifyWorkspacePropertiesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalModifyWorkspaceStateInputShapeAWSJSON(v *types.ModifyWorkspaceStateInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalRebootWorkspacesInputShapeAWSJSON(v *types.RebootWorkspacesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalRebuildWorkspacesInputShapeAWSJSON(v *types.RebuildWorkspacesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalRegisterWorkspaceDirectoryInputShapeAWSJSON(v *types.RegisterWorkspaceDirectoryInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalRestoreWorkspaceInputShapeAWSJSON(v *types.RestoreWorkspaceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalRevokeIpRulesInputShapeAWSJSON(v *types.RevokeIpRulesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalStartWorkspacesInputShapeAWSJSON(v *types.StartWorkspacesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalStopWorkspacesInputShapeAWSJSON(v *types.StopWorkspacesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalTerminateWorkspacesInputShapeAWSJSON(v *types.TerminateWorkspacesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalUpdateRulesOfIpGroupInputShapeAWSJSON(v *types.UpdateRulesOfIpGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}
func (m AssociateIpGroupsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalAssociateIpGroupsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m AssociateIpGroupsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "AssociateIpGroups.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m AuthorizeIpRulesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalAuthorizeIpRulesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m AuthorizeIpRulesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "AuthorizeIpRules.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CopyWorkspaceImageMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCopyWorkspaceImageInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CopyWorkspaceImageMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CopyWorkspaceImage.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreateIpGroupMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateIpGroupInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateIpGroupMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateIpGroup.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreateTagsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateTagsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateTagsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateTags.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreateWorkspacesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateWorkspacesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateWorkspacesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateWorkspaces.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteIpGroupMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteIpGroupInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteIpGroupMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteIpGroup.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteTagsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteTagsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteTagsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteTags.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteWorkspaceImageMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteWorkspaceImageInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteWorkspaceImageMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteWorkspaceImage.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeregisterWorkspaceDirectoryMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeregisterWorkspaceDirectoryInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeregisterWorkspaceDirectoryMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeregisterWorkspaceDirectory.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeAccountMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeAccountInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeAccountMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeAccount.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeAccountModificationsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeAccountModificationsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeAccountModificationsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeAccountModifications.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeClientPropertiesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeClientPropertiesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeClientPropertiesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeClientProperties.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeIpGroupsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeIpGroupsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeIpGroupsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeIpGroups.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeTagsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeTagsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeTagsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeTags.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeWorkspaceBundlesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeWorkspaceBundlesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeWorkspaceBundlesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeWorkspaceBundles.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeWorkspaceDirectoriesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeWorkspaceDirectoriesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeWorkspaceDirectoriesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeWorkspaceDirectories.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeWorkspaceImagesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeWorkspaceImagesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeWorkspaceImagesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeWorkspaceImages.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeWorkspaceSnapshotsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeWorkspaceSnapshotsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeWorkspaceSnapshotsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeWorkspaceSnapshots.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeWorkspacesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeWorkspacesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeWorkspacesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeWorkspaces.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeWorkspacesConnectionStatusMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeWorkspacesConnectionStatusInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeWorkspacesConnectionStatusMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeWorkspacesConnectionStatus.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DisassociateIpGroupsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDisassociateIpGroupsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DisassociateIpGroupsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DisassociateIpGroups.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ImportWorkspaceImageMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalImportWorkspaceImageInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ImportWorkspaceImageMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ImportWorkspaceImage.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListAvailableManagementCidrRangesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListAvailableManagementCidrRangesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListAvailableManagementCidrRangesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListAvailableManagementCidrRanges.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ModifyAccountMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalModifyAccountInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ModifyAccountMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ModifyAccount.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ModifyClientPropertiesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalModifyClientPropertiesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ModifyClientPropertiesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ModifyClientProperties.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ModifySelfservicePermissionsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalModifySelfservicePermissionsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ModifySelfservicePermissionsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ModifySelfservicePermissions.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ModifyWorkspaceAccessPropertiesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalModifyWorkspaceAccessPropertiesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ModifyWorkspaceAccessPropertiesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ModifyWorkspaceAccessProperties.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ModifyWorkspaceCreationPropertiesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalModifyWorkspaceCreationPropertiesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ModifyWorkspaceCreationPropertiesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ModifyWorkspaceCreationProperties.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ModifyWorkspacePropertiesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalModifyWorkspacePropertiesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ModifyWorkspacePropertiesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ModifyWorkspaceProperties.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ModifyWorkspaceStateMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalModifyWorkspaceStateInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ModifyWorkspaceStateMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ModifyWorkspaceState.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m RebootWorkspacesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalRebootWorkspacesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m RebootWorkspacesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "RebootWorkspaces.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m RebuildWorkspacesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalRebuildWorkspacesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m RebuildWorkspacesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "RebuildWorkspaces.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m RegisterWorkspaceDirectoryMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalRegisterWorkspaceDirectoryInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m RegisterWorkspaceDirectoryMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "RegisterWorkspaceDirectory.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m RestoreWorkspaceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalRestoreWorkspaceInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m RestoreWorkspaceMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "RestoreWorkspace.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m RevokeIpRulesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalRevokeIpRulesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m RevokeIpRulesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "RevokeIpRules.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m StartWorkspacesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalStartWorkspacesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m StartWorkspacesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "StartWorkspaces.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m StopWorkspacesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalStopWorkspacesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m StopWorkspacesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "StopWorkspaces.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m TerminateWorkspacesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalTerminateWorkspacesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m TerminateWorkspacesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "TerminateWorkspaces.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateRulesOfIpGroupMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateRulesOfIpGroupInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateRulesOfIpGroupMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateRulesOfIpGroup.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}
