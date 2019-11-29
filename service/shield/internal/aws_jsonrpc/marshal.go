// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package aws_jsonrpc service, marshal.go contains the generated operation marshalers and methods on it.

package aws_jsonrpc

import (
	"bytes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/service/shield/types"
)

var _ bytes.Buffer
var _ awserr.Error

// AssociateDRTLogBucketMarshaler defines marshaler for AssociateDRTLogBucket operation
type AssociateDRTLogBucketMarshaler struct {
	Input *types.AssociateDRTLogBucketInput
}

// AssociateDRTRoleMarshaler defines marshaler for AssociateDRTRole operation
type AssociateDRTRoleMarshaler struct {
	Input *types.AssociateDRTRoleInput
}

// CreateProtectionMarshaler defines marshaler for CreateProtection operation
type CreateProtectionMarshaler struct {
	Input *types.CreateProtectionInput
}

// CreateSubscriptionMarshaler defines marshaler for CreateSubscription operation
type CreateSubscriptionMarshaler struct {
	Input *types.CreateSubscriptionInput
}

// DeleteProtectionMarshaler defines marshaler for DeleteProtection operation
type DeleteProtectionMarshaler struct {
	Input *types.DeleteProtectionInput
}

// DeleteSubscriptionMarshaler defines marshaler for DeleteSubscription operation
type DeleteSubscriptionMarshaler struct {
	Input *types.DeleteSubscriptionInput
}

// DescribeAttackMarshaler defines marshaler for DescribeAttack operation
type DescribeAttackMarshaler struct {
	Input *types.DescribeAttackInput
}

// DescribeDRTAccessMarshaler defines marshaler for DescribeDRTAccess operation
type DescribeDRTAccessMarshaler struct {
	Input *types.DescribeDRTAccessInput
}

// DescribeEmergencyContactSettingsMarshaler defines marshaler for DescribeEmergencyContactSettings operation
type DescribeEmergencyContactSettingsMarshaler struct {
	Input *types.DescribeEmergencyContactSettingsInput
}

// DescribeProtectionMarshaler defines marshaler for DescribeProtection operation
type DescribeProtectionMarshaler struct {
	Input *types.DescribeProtectionInput
}

// DescribeSubscriptionMarshaler defines marshaler for DescribeSubscription operation
type DescribeSubscriptionMarshaler struct {
	Input *types.DescribeSubscriptionInput
}

// DisassociateDRTLogBucketMarshaler defines marshaler for DisassociateDRTLogBucket operation
type DisassociateDRTLogBucketMarshaler struct {
	Input *types.DisassociateDRTLogBucketInput
}

// DisassociateDRTRoleMarshaler defines marshaler for DisassociateDRTRole operation
type DisassociateDRTRoleMarshaler struct {
	Input *types.DisassociateDRTRoleInput
}

// GetSubscriptionStateMarshaler defines marshaler for GetSubscriptionState operation
type GetSubscriptionStateMarshaler struct {
	Input *types.GetSubscriptionStateInput
}

// ListAttacksMarshaler defines marshaler for ListAttacks operation
type ListAttacksMarshaler struct {
	Input *types.ListAttacksInput
}

// ListProtectionsMarshaler defines marshaler for ListProtections operation
type ListProtectionsMarshaler struct {
	Input *types.ListProtectionsInput
}

// UpdateEmergencyContactSettingsMarshaler defines marshaler for UpdateEmergencyContactSettings operation
type UpdateEmergencyContactSettingsMarshaler struct {
	Input *types.UpdateEmergencyContactSettingsInput
}

// UpdateSubscriptionMarshaler defines marshaler for UpdateSubscription operation
type UpdateSubscriptionMarshaler struct {
	Input *types.UpdateSubscriptionInput
}

func MarshalAssociateDRTLogBucketInputShapeAWSJSON(v *types.AssociateDRTLogBucketInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalAssociateDRTRoleInputShapeAWSJSON(v *types.AssociateDRTRoleInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalCreateProtectionInputShapeAWSJSON(v *types.CreateProtectionInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalCreateSubscriptionInputShapeAWSJSON(v *types.CreateSubscriptionInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDeleteProtectionInputShapeAWSJSON(v *types.DeleteProtectionInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDeleteSubscriptionInputShapeAWSJSON(v *types.DeleteSubscriptionInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeAttackInputShapeAWSJSON(v *types.DescribeAttackInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeDRTAccessInputShapeAWSJSON(v *types.DescribeDRTAccessInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeEmergencyContactSettingsInputShapeAWSJSON(v *types.DescribeEmergencyContactSettingsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeProtectionInputShapeAWSJSON(v *types.DescribeProtectionInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDescribeSubscriptionInputShapeAWSJSON(v *types.DescribeSubscriptionInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDisassociateDRTLogBucketInputShapeAWSJSON(v *types.DisassociateDRTLogBucketInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalDisassociateDRTRoleInputShapeAWSJSON(v *types.DisassociateDRTRoleInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalGetSubscriptionStateInputShapeAWSJSON(v *types.GetSubscriptionStateInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalListAttacksInputShapeAWSJSON(v *types.ListAttacksInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalListProtectionsInputShapeAWSJSON(v *types.ListProtectionsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalUpdateEmergencyContactSettingsInputShapeAWSJSON(v *types.UpdateEmergencyContactSettingsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}

func MarshalUpdateSubscriptionInputShapeAWSJSON(v *types.UpdateSubscriptionInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	jsonrpc.Build(r)
	return nil
}
func (m AssociateDRTLogBucketMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalAssociateDRTLogBucketInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m AssociateDRTLogBucketMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "AssociateDRTLogBucket.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m AssociateDRTRoleMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalAssociateDRTRoleInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m AssociateDRTRoleMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "AssociateDRTRole.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreateProtectionMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateProtectionInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateProtectionMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateProtection.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreateSubscriptionMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateSubscriptionInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateSubscriptionMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateSubscription.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteProtectionMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteProtectionInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteProtectionMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteProtection.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteSubscriptionMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteSubscriptionInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteSubscriptionMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteSubscription.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeAttackMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeAttackInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeAttackMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeAttack.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeDRTAccessMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeDRTAccessInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeDRTAccessMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeDRTAccess.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeEmergencyContactSettingsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeEmergencyContactSettingsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeEmergencyContactSettingsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeEmergencyContactSettings.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeProtectionMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeProtectionInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeProtectionMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeProtection.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeSubscriptionMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeSubscriptionInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeSubscriptionMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeSubscription.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DisassociateDRTLogBucketMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDisassociateDRTLogBucketInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DisassociateDRTLogBucketMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DisassociateDRTLogBucket.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DisassociateDRTRoleMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDisassociateDRTRoleInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DisassociateDRTRoleMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DisassociateDRTRole.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetSubscriptionStateMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetSubscriptionStateInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetSubscriptionStateMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetSubscriptionState.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListAttacksMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListAttacksInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListAttacksMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListAttacks.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListProtectionsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListProtectionsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListProtectionsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListProtections.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateEmergencyContactSettingsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateEmergencyContactSettingsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateEmergencyContactSettingsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateEmergencyContactSettings.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateSubscriptionMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateSubscriptionInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateSubscriptionMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateSubscription.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}
