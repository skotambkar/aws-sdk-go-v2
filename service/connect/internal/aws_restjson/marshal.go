// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package aws_restjson service, marshal.go contains the generated operation marshalers and methods on it.

package aws_restjson

import (
	"bytes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/private/protocol/rest"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
)

var _ bytes.Buffer
var _ awserr.Error

// CreateUserMarshaler defines marshaler for CreateUser operation
type CreateUserMarshaler struct {
	Input *types.CreateUserInput
}

// DeleteUserMarshaler defines marshaler for DeleteUser operation
type DeleteUserMarshaler struct {
	Input *types.DeleteUserInput
}

// DescribeUserMarshaler defines marshaler for DescribeUser operation
type DescribeUserMarshaler struct {
	Input *types.DescribeUserInput
}

// DescribeUserHierarchyGroupMarshaler defines marshaler for DescribeUserHierarchyGroup operation
type DescribeUserHierarchyGroupMarshaler struct {
	Input *types.DescribeUserHierarchyGroupInput
}

// DescribeUserHierarchyStructureMarshaler defines marshaler for DescribeUserHierarchyStructure operation
type DescribeUserHierarchyStructureMarshaler struct {
	Input *types.DescribeUserHierarchyStructureInput
}

// GetContactAttributesMarshaler defines marshaler for GetContactAttributes operation
type GetContactAttributesMarshaler struct {
	Input *types.GetContactAttributesInput
}

// GetCurrentMetricDataMarshaler defines marshaler for GetCurrentMetricData operation
type GetCurrentMetricDataMarshaler struct {
	Input *types.GetCurrentMetricDataInput
}

// GetFederationTokenMarshaler defines marshaler for GetFederationToken operation
type GetFederationTokenMarshaler struct {
	Input *types.GetFederationTokenInput
}

// GetMetricDataMarshaler defines marshaler for GetMetricData operation
type GetMetricDataMarshaler struct {
	Input *types.GetMetricDataInput
}

// ListContactFlowsMarshaler defines marshaler for ListContactFlows operation
type ListContactFlowsMarshaler struct {
	Input *types.ListContactFlowsInput
}

// ListHoursOfOperationsMarshaler defines marshaler for ListHoursOfOperations operation
type ListHoursOfOperationsMarshaler struct {
	Input *types.ListHoursOfOperationsInput
}

// ListPhoneNumbersMarshaler defines marshaler for ListPhoneNumbers operation
type ListPhoneNumbersMarshaler struct {
	Input *types.ListPhoneNumbersInput
}

// ListQueuesMarshaler defines marshaler for ListQueues operation
type ListQueuesMarshaler struct {
	Input *types.ListQueuesInput
}

// ListRoutingProfilesMarshaler defines marshaler for ListRoutingProfiles operation
type ListRoutingProfilesMarshaler struct {
	Input *types.ListRoutingProfilesInput
}

// ListSecurityProfilesMarshaler defines marshaler for ListSecurityProfiles operation
type ListSecurityProfilesMarshaler struct {
	Input *types.ListSecurityProfilesInput
}

// ListTagsForResourceMarshaler defines marshaler for ListTagsForResource operation
type ListTagsForResourceMarshaler struct {
	Input *types.ListTagsForResourceInput
}

// ListUserHierarchyGroupsMarshaler defines marshaler for ListUserHierarchyGroups operation
type ListUserHierarchyGroupsMarshaler struct {
	Input *types.ListUserHierarchyGroupsInput
}

// ListUsersMarshaler defines marshaler for ListUsers operation
type ListUsersMarshaler struct {
	Input *types.ListUsersInput
}

// StartOutboundVoiceContactMarshaler defines marshaler for StartOutboundVoiceContact operation
type StartOutboundVoiceContactMarshaler struct {
	Input *types.StartOutboundVoiceContactInput
}

// StopContactMarshaler defines marshaler for StopContact operation
type StopContactMarshaler struct {
	Input *types.StopContactInput
}

// TagResourceMarshaler defines marshaler for TagResource operation
type TagResourceMarshaler struct {
	Input *types.TagResourceInput
}

// UntagResourceMarshaler defines marshaler for UntagResource operation
type UntagResourceMarshaler struct {
	Input *types.UntagResourceInput
}

// UpdateContactAttributesMarshaler defines marshaler for UpdateContactAttributes operation
type UpdateContactAttributesMarshaler struct {
	Input *types.UpdateContactAttributesInput
}

// UpdateUserHierarchyMarshaler defines marshaler for UpdateUserHierarchy operation
type UpdateUserHierarchyMarshaler struct {
	Input *types.UpdateUserHierarchyInput
}

// UpdateUserIdentityInfoMarshaler defines marshaler for UpdateUserIdentityInfo operation
type UpdateUserIdentityInfoMarshaler struct {
	Input *types.UpdateUserIdentityInfoInput
}

// UpdateUserPhoneConfigMarshaler defines marshaler for UpdateUserPhoneConfig operation
type UpdateUserPhoneConfigMarshaler struct {
	Input *types.UpdateUserPhoneConfigInput
}

// UpdateUserRoutingProfileMarshaler defines marshaler for UpdateUserRoutingProfile operation
type UpdateUserRoutingProfileMarshaler struct {
	Input *types.UpdateUserRoutingProfileInput
}

// UpdateUserSecurityProfilesMarshaler defines marshaler for UpdateUserSecurityProfiles operation
type UpdateUserSecurityProfilesMarshaler struct {
	Input *types.UpdateUserSecurityProfilesInput
}

func MarshalCreateUserInputShapeAWSREST(v *types.CreateUserInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalCreateUserInputShapeAWSJSON(v *types.CreateUserInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalDeleteUserInputShapeAWSREST(v *types.DeleteUserInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalDeleteUserInputShapeAWSJSON(v *types.DeleteUserInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalDescribeUserInputShapeAWSREST(v *types.DescribeUserInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalDescribeUserInputShapeAWSJSON(v *types.DescribeUserInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalDescribeUserHierarchyGroupInputShapeAWSREST(v *types.DescribeUserHierarchyGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalDescribeUserHierarchyGroupInputShapeAWSJSON(v *types.DescribeUserHierarchyGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalDescribeUserHierarchyStructureInputShapeAWSREST(v *types.DescribeUserHierarchyStructureInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalDescribeUserHierarchyStructureInputShapeAWSJSON(v *types.DescribeUserHierarchyStructureInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalGetContactAttributesInputShapeAWSREST(v *types.GetContactAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalGetContactAttributesInputShapeAWSJSON(v *types.GetContactAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalGetCurrentMetricDataInputShapeAWSREST(v *types.GetCurrentMetricDataInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalGetCurrentMetricDataInputShapeAWSJSON(v *types.GetCurrentMetricDataInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalGetFederationTokenInputShapeAWSREST(v *types.GetFederationTokenInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalGetFederationTokenInputShapeAWSJSON(v *types.GetFederationTokenInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalGetMetricDataInputShapeAWSREST(v *types.GetMetricDataInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalGetMetricDataInputShapeAWSJSON(v *types.GetMetricDataInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListContactFlowsInputShapeAWSREST(v *types.ListContactFlowsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListContactFlowsInputShapeAWSJSON(v *types.ListContactFlowsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListHoursOfOperationsInputShapeAWSREST(v *types.ListHoursOfOperationsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListHoursOfOperationsInputShapeAWSJSON(v *types.ListHoursOfOperationsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListPhoneNumbersInputShapeAWSREST(v *types.ListPhoneNumbersInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListPhoneNumbersInputShapeAWSJSON(v *types.ListPhoneNumbersInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListQueuesInputShapeAWSREST(v *types.ListQueuesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListQueuesInputShapeAWSJSON(v *types.ListQueuesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListRoutingProfilesInputShapeAWSREST(v *types.ListRoutingProfilesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListRoutingProfilesInputShapeAWSJSON(v *types.ListRoutingProfilesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListSecurityProfilesInputShapeAWSREST(v *types.ListSecurityProfilesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListSecurityProfilesInputShapeAWSJSON(v *types.ListSecurityProfilesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListTagsForResourceInputShapeAWSREST(v *types.ListTagsForResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListTagsForResourceInputShapeAWSJSON(v *types.ListTagsForResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListUserHierarchyGroupsInputShapeAWSREST(v *types.ListUserHierarchyGroupsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListUserHierarchyGroupsInputShapeAWSJSON(v *types.ListUserHierarchyGroupsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListUsersInputShapeAWSREST(v *types.ListUsersInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListUsersInputShapeAWSJSON(v *types.ListUsersInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalStartOutboundVoiceContactInputShapeAWSREST(v *types.StartOutboundVoiceContactInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalStartOutboundVoiceContactInputShapeAWSJSON(v *types.StartOutboundVoiceContactInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalStopContactInputShapeAWSREST(v *types.StopContactInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalStopContactInputShapeAWSJSON(v *types.StopContactInput, r *aws.Request) error {
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

func MarshalUpdateContactAttributesInputShapeAWSREST(v *types.UpdateContactAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUpdateContactAttributesInputShapeAWSJSON(v *types.UpdateContactAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalUpdateUserHierarchyInputShapeAWSREST(v *types.UpdateUserHierarchyInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUpdateUserHierarchyInputShapeAWSJSON(v *types.UpdateUserHierarchyInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalUpdateUserIdentityInfoInputShapeAWSREST(v *types.UpdateUserIdentityInfoInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUpdateUserIdentityInfoInputShapeAWSJSON(v *types.UpdateUserIdentityInfoInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalUpdateUserPhoneConfigInputShapeAWSREST(v *types.UpdateUserPhoneConfigInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUpdateUserPhoneConfigInputShapeAWSJSON(v *types.UpdateUserPhoneConfigInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalUpdateUserRoutingProfileInputShapeAWSREST(v *types.UpdateUserRoutingProfileInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUpdateUserRoutingProfileInputShapeAWSJSON(v *types.UpdateUserRoutingProfileInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalUpdateUserSecurityProfilesInputShapeAWSREST(v *types.UpdateUserSecurityProfilesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUpdateUserSecurityProfilesInputShapeAWSJSON(v *types.UpdateUserSecurityProfilesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}
func (m CreateUserMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateUserInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalCreateUserInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateUserMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateUser.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteUserMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteUserInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalDeleteUserInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteUserMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteUser.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeUserMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeUserInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalDescribeUserInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeUserMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeUser.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeUserHierarchyGroupMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeUserHierarchyGroupInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalDescribeUserHierarchyGroupInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeUserHierarchyGroupMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeUserHierarchyGroup.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DescribeUserHierarchyStructureMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDescribeUserHierarchyStructureInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalDescribeUserHierarchyStructureInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DescribeUserHierarchyStructureMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DescribeUserHierarchyStructure.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetContactAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetContactAttributesInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalGetContactAttributesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetContactAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetContactAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetCurrentMetricDataMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetCurrentMetricDataInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalGetCurrentMetricDataInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetCurrentMetricDataMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetCurrentMetricData.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetFederationTokenMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetFederationTokenInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalGetFederationTokenInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetFederationTokenMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetFederationToken.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetMetricDataMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetMetricDataInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalGetMetricDataInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetMetricDataMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetMetricData.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListContactFlowsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListContactFlowsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListContactFlowsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListContactFlowsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListContactFlows.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListHoursOfOperationsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListHoursOfOperationsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListHoursOfOperationsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListHoursOfOperationsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListHoursOfOperations.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListPhoneNumbersMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListPhoneNumbersInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListPhoneNumbersInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListPhoneNumbersMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListPhoneNumbers.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListQueuesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListQueuesInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListQueuesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListQueuesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListQueues.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListRoutingProfilesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListRoutingProfilesInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListRoutingProfilesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListRoutingProfilesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListRoutingProfiles.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListSecurityProfilesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListSecurityProfilesInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListSecurityProfilesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListSecurityProfilesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListSecurityProfiles.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListTagsForResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListTagsForResourceInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

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

func (m ListUserHierarchyGroupsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListUserHierarchyGroupsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListUserHierarchyGroupsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListUserHierarchyGroupsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListUserHierarchyGroups.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListUsersMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListUsersInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListUsersInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListUsersMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListUsers.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m StartOutboundVoiceContactMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalStartOutboundVoiceContactInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalStartOutboundVoiceContactInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m StartOutboundVoiceContactMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "StartOutboundVoiceContact.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m StopContactMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalStopContactInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalStopContactInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m StopContactMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "StopContact.BuildHandler"
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

func (m UpdateContactAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateContactAttributesInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUpdateContactAttributesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateContactAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateContactAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateUserHierarchyMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateUserHierarchyInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUpdateUserHierarchyInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateUserHierarchyMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateUserHierarchy.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateUserIdentityInfoMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateUserIdentityInfoInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUpdateUserIdentityInfoInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateUserIdentityInfoMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateUserIdentityInfo.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateUserPhoneConfigMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateUserPhoneConfigInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUpdateUserPhoneConfigInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateUserPhoneConfigMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateUserPhoneConfig.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateUserRoutingProfileMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateUserRoutingProfileInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUpdateUserRoutingProfileInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateUserRoutingProfileMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateUserRoutingProfile.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateUserSecurityProfilesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateUserSecurityProfilesInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUpdateUserSecurityProfilesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateUserSecurityProfilesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateUserSecurityProfiles.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}
