// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package aws_jsonrpc service, marshal.go contains the generated operation marshalers and methods on it.

package aws_jsonrpc

import (
	"bytes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
)

var _ bytes.Buffer
var _ awserr.Error

// CreateLicenseConfigurationMarshaler defines marshaler for CreateLicenseConfiguration operation
type CreateLicenseConfigurationMarshaler struct {
	Input *types.CreateLicenseConfigurationInput
}

// DeleteLicenseConfigurationMarshaler defines marshaler for DeleteLicenseConfiguration operation
type DeleteLicenseConfigurationMarshaler struct {
	Input *types.DeleteLicenseConfigurationInput
}

// GetLicenseConfigurationMarshaler defines marshaler for GetLicenseConfiguration operation
type GetLicenseConfigurationMarshaler struct {
	Input *types.GetLicenseConfigurationInput
}

// GetServiceSettingsMarshaler defines marshaler for GetServiceSettings operation
type GetServiceSettingsMarshaler struct {
	Input *types.GetServiceSettingsInput
}

// ListAssociationsForLicenseConfigurationMarshaler defines marshaler for ListAssociationsForLicenseConfiguration operation
type ListAssociationsForLicenseConfigurationMarshaler struct {
	Input *types.ListAssociationsForLicenseConfigurationInput
}

// ListLicenseConfigurationsMarshaler defines marshaler for ListLicenseConfigurations operation
type ListLicenseConfigurationsMarshaler struct {
	Input *types.ListLicenseConfigurationsInput
}

// ListLicenseSpecificationsForResourceMarshaler defines marshaler for ListLicenseSpecificationsForResource operation
type ListLicenseSpecificationsForResourceMarshaler struct {
	Input *types.ListLicenseSpecificationsForResourceInput
}

// ListResourceInventoryMarshaler defines marshaler for ListResourceInventory operation
type ListResourceInventoryMarshaler struct {
	Input *types.ListResourceInventoryInput
}

// ListTagsForResourceMarshaler defines marshaler for ListTagsForResource operation
type ListTagsForResourceMarshaler struct {
	Input *types.ListTagsForResourceInput
}

// ListUsageForLicenseConfigurationMarshaler defines marshaler for ListUsageForLicenseConfiguration operation
type ListUsageForLicenseConfigurationMarshaler struct {
	Input *types.ListUsageForLicenseConfigurationInput
}

// TagResourceMarshaler defines marshaler for TagResource operation
type TagResourceMarshaler struct {
	Input *types.TagResourceInput
}

// UntagResourceMarshaler defines marshaler for UntagResource operation
type UntagResourceMarshaler struct {
	Input *types.UntagResourceInput
}

// UpdateLicenseConfigurationMarshaler defines marshaler for UpdateLicenseConfiguration operation
type UpdateLicenseConfigurationMarshaler struct {
	Input *types.UpdateLicenseConfigurationInput
}

// UpdateLicenseSpecificationsForResourceMarshaler defines marshaler for UpdateLicenseSpecificationsForResource operation
type UpdateLicenseSpecificationsForResourceMarshaler struct {
	Input *types.UpdateLicenseSpecificationsForResourceInput
}

// UpdateServiceSettingsMarshaler defines marshaler for UpdateServiceSettings operation
type UpdateServiceSettingsMarshaler struct {
	Input *types.UpdateServiceSettingsInput
}

func MarshalCreateLicenseConfigurationInputShapeAWSJSON(v *types.CreateLicenseConfigurationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDeleteLicenseConfigurationInputShapeAWSJSON(v *types.DeleteLicenseConfigurationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalGetLicenseConfigurationInputShapeAWSJSON(v *types.GetLicenseConfigurationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalGetServiceSettingsInputShapeAWSJSON(v *types.GetServiceSettingsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalListAssociationsForLicenseConfigurationInputShapeAWSJSON(v *types.ListAssociationsForLicenseConfigurationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalListLicenseConfigurationsInputShapeAWSJSON(v *types.ListLicenseConfigurationsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalListLicenseSpecificationsForResourceInputShapeAWSJSON(v *types.ListLicenseSpecificationsForResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalListResourceInventoryInputShapeAWSJSON(v *types.ListResourceInventoryInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalListTagsForResourceInputShapeAWSJSON(v *types.ListTagsForResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalListUsageForLicenseConfigurationInputShapeAWSJSON(v *types.ListUsageForLicenseConfigurationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalTagResourceInputShapeAWSJSON(v *types.TagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalUntagResourceInputShapeAWSJSON(v *types.UntagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalUpdateLicenseConfigurationInputShapeAWSJSON(v *types.UpdateLicenseConfigurationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalUpdateLicenseSpecificationsForResourceInputShapeAWSJSON(v *types.UpdateLicenseSpecificationsForResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalUpdateServiceSettingsInputShapeAWSJSON(v *types.UpdateServiceSettingsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}
func (m CreateLicenseConfigurationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateLicenseConfigurationInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateLicenseConfigurationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateLicenseConfiguration.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteLicenseConfigurationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteLicenseConfigurationInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteLicenseConfigurationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteLicenseConfiguration.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetLicenseConfigurationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetLicenseConfigurationInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetLicenseConfigurationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetLicenseConfiguration.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetServiceSettingsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetServiceSettingsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetServiceSettingsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetServiceSettings.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListAssociationsForLicenseConfigurationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListAssociationsForLicenseConfigurationInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListAssociationsForLicenseConfigurationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListAssociationsForLicenseConfiguration.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListLicenseConfigurationsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListLicenseConfigurationsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListLicenseConfigurationsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListLicenseConfigurations.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListLicenseSpecificationsForResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListLicenseSpecificationsForResourceInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListLicenseSpecificationsForResourceMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListLicenseSpecificationsForResource.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListResourceInventoryMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListResourceInventoryInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListResourceInventoryMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListResourceInventory.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListTagsForResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListTagsForResourceInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListTagsForResourceMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListTagsForResource.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListUsageForLicenseConfigurationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListUsageForLicenseConfigurationInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListUsageForLicenseConfigurationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListUsageForLicenseConfiguration.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m TagResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

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

func (m UpdateLicenseConfigurationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateLicenseConfigurationInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateLicenseConfigurationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateLicenseConfiguration.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateLicenseSpecificationsForResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateLicenseSpecificationsForResourceInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateLicenseSpecificationsForResourceMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateLicenseSpecificationsForResource.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateServiceSettingsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateServiceSettingsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateServiceSettingsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateServiceSettings.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}
