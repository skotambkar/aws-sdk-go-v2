// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package aws_query service, marshal.go contains the generated operation marshalers and methods on it.

package aws_query

import (
	"bytes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/private/protocol/query"
	"github.com/aws/aws-sdk-go-v2/service/sns/types"
)

var _ bytes.Buffer
var _ awserr.Error

// AddPermissionMarshaler defines marshaler for AddPermission operation
type AddPermissionMarshaler struct {
	Input *types.AddPermissionInput
}

// CheckIfPhoneNumberIsOptedOutMarshaler defines marshaler for CheckIfPhoneNumberIsOptedOut operation
type CheckIfPhoneNumberIsOptedOutMarshaler struct {
	Input *types.CheckIfPhoneNumberIsOptedOutInput
}

// ConfirmSubscriptionMarshaler defines marshaler for ConfirmSubscription operation
type ConfirmSubscriptionMarshaler struct {
	Input *types.ConfirmSubscriptionInput
}

// CreatePlatformApplicationMarshaler defines marshaler for CreatePlatformApplication operation
type CreatePlatformApplicationMarshaler struct {
	Input *types.CreatePlatformApplicationInput
}

// CreatePlatformEndpointMarshaler defines marshaler for CreatePlatformEndpoint operation
type CreatePlatformEndpointMarshaler struct {
	Input *types.CreatePlatformEndpointInput
}

// CreateTopicMarshaler defines marshaler for CreateTopic operation
type CreateTopicMarshaler struct {
	Input *types.CreateTopicInput
}

// DeleteEndpointMarshaler defines marshaler for DeleteEndpoint operation
type DeleteEndpointMarshaler struct {
	Input *types.DeleteEndpointInput
}

// DeletePlatformApplicationMarshaler defines marshaler for DeletePlatformApplication operation
type DeletePlatformApplicationMarshaler struct {
	Input *types.DeletePlatformApplicationInput
}

// DeleteTopicMarshaler defines marshaler for DeleteTopic operation
type DeleteTopicMarshaler struct {
	Input *types.DeleteTopicInput
}

// GetEndpointAttributesMarshaler defines marshaler for GetEndpointAttributes operation
type GetEndpointAttributesMarshaler struct {
	Input *types.GetEndpointAttributesInput
}

// GetPlatformApplicationAttributesMarshaler defines marshaler for GetPlatformApplicationAttributes operation
type GetPlatformApplicationAttributesMarshaler struct {
	Input *types.GetPlatformApplicationAttributesInput
}

// GetSMSAttributesMarshaler defines marshaler for GetSMSAttributes operation
type GetSMSAttributesMarshaler struct {
	Input *types.GetSMSAttributesInput
}

// GetSubscriptionAttributesMarshaler defines marshaler for GetSubscriptionAttributes operation
type GetSubscriptionAttributesMarshaler struct {
	Input *types.GetSubscriptionAttributesInput
}

// GetTopicAttributesMarshaler defines marshaler for GetTopicAttributes operation
type GetTopicAttributesMarshaler struct {
	Input *types.GetTopicAttributesInput
}

// ListEndpointsByPlatformApplicationMarshaler defines marshaler for ListEndpointsByPlatformApplication operation
type ListEndpointsByPlatformApplicationMarshaler struct {
	Input *types.ListEndpointsByPlatformApplicationInput
}

// ListPhoneNumbersOptedOutMarshaler defines marshaler for ListPhoneNumbersOptedOut operation
type ListPhoneNumbersOptedOutMarshaler struct {
	Input *types.ListPhoneNumbersOptedOutInput
}

// ListPlatformApplicationsMarshaler defines marshaler for ListPlatformApplications operation
type ListPlatformApplicationsMarshaler struct {
	Input *types.ListPlatformApplicationsInput
}

// ListSubscriptionsMarshaler defines marshaler for ListSubscriptions operation
type ListSubscriptionsMarshaler struct {
	Input *types.ListSubscriptionsInput
}

// ListSubscriptionsByTopicMarshaler defines marshaler for ListSubscriptionsByTopic operation
type ListSubscriptionsByTopicMarshaler struct {
	Input *types.ListSubscriptionsByTopicInput
}

// ListTagsForResourceMarshaler defines marshaler for ListTagsForResource operation
type ListTagsForResourceMarshaler struct {
	Input *types.ListTagsForResourceInput
}

// ListTopicsMarshaler defines marshaler for ListTopics operation
type ListTopicsMarshaler struct {
	Input *types.ListTopicsInput
}

// OptInPhoneNumberMarshaler defines marshaler for OptInPhoneNumber operation
type OptInPhoneNumberMarshaler struct {
	Input *types.OptInPhoneNumberInput
}

// PublishMarshaler defines marshaler for Publish operation
type PublishMarshaler struct {
	Input *types.PublishInput
}

// RemovePermissionMarshaler defines marshaler for RemovePermission operation
type RemovePermissionMarshaler struct {
	Input *types.RemovePermissionInput
}

// SetEndpointAttributesMarshaler defines marshaler for SetEndpointAttributes operation
type SetEndpointAttributesMarshaler struct {
	Input *types.SetEndpointAttributesInput
}

// SetPlatformApplicationAttributesMarshaler defines marshaler for SetPlatformApplicationAttributes operation
type SetPlatformApplicationAttributesMarshaler struct {
	Input *types.SetPlatformApplicationAttributesInput
}

// SetSMSAttributesMarshaler defines marshaler for SetSMSAttributes operation
type SetSMSAttributesMarshaler struct {
	Input *types.SetSMSAttributesInput
}

// SetSubscriptionAttributesMarshaler defines marshaler for SetSubscriptionAttributes operation
type SetSubscriptionAttributesMarshaler struct {
	Input *types.SetSubscriptionAttributesInput
}

// SetTopicAttributesMarshaler defines marshaler for SetTopicAttributes operation
type SetTopicAttributesMarshaler struct {
	Input *types.SetTopicAttributesInput
}

// SubscribeMarshaler defines marshaler for Subscribe operation
type SubscribeMarshaler struct {
	Input *types.SubscribeInput
}

// TagResourceMarshaler defines marshaler for TagResource operation
type TagResourceMarshaler struct {
	Input *types.TagResourceInput
}

// UnsubscribeMarshaler defines marshaler for Unsubscribe operation
type UnsubscribeMarshaler struct {
	Input *types.UnsubscribeInput
}

// UntagResourceMarshaler defines marshaler for UntagResource operation
type UntagResourceMarshaler struct {
	Input *types.UntagResourceInput
}

func MarshalAddPermissionInputShapeAWSQuery(v *types.AddPermissionInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalCheckIfPhoneNumberIsOptedOutInputShapeAWSQuery(v *types.CheckIfPhoneNumberIsOptedOutInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalConfirmSubscriptionInputShapeAWSQuery(v *types.ConfirmSubscriptionInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalCreatePlatformApplicationInputShapeAWSQuery(v *types.CreatePlatformApplicationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalCreatePlatformEndpointInputShapeAWSQuery(v *types.CreatePlatformEndpointInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalCreateTopicInputShapeAWSQuery(v *types.CreateTopicInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDeleteEndpointInputShapeAWSQuery(v *types.DeleteEndpointInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDeletePlatformApplicationInputShapeAWSQuery(v *types.DeletePlatformApplicationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalDeleteTopicInputShapeAWSQuery(v *types.DeleteTopicInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalGetEndpointAttributesInputShapeAWSQuery(v *types.GetEndpointAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalGetPlatformApplicationAttributesInputShapeAWSQuery(v *types.GetPlatformApplicationAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalGetSMSAttributesInputShapeAWSQuery(v *types.GetSMSAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalGetSubscriptionAttributesInputShapeAWSQuery(v *types.GetSubscriptionAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalGetTopicAttributesInputShapeAWSQuery(v *types.GetTopicAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalListEndpointsByPlatformApplicationInputShapeAWSQuery(v *types.ListEndpointsByPlatformApplicationInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalListPhoneNumbersOptedOutInputShapeAWSQuery(v *types.ListPhoneNumbersOptedOutInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalListPlatformApplicationsInputShapeAWSQuery(v *types.ListPlatformApplicationsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalListSubscriptionsInputShapeAWSQuery(v *types.ListSubscriptionsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalListSubscriptionsByTopicInputShapeAWSQuery(v *types.ListSubscriptionsByTopicInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalListTagsForResourceInputShapeAWSQuery(v *types.ListTagsForResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalListTopicsInputShapeAWSQuery(v *types.ListTopicsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalOptInPhoneNumberInputShapeAWSQuery(v *types.OptInPhoneNumberInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalPublishInputShapeAWSQuery(v *types.PublishInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalRemovePermissionInputShapeAWSQuery(v *types.RemovePermissionInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalSetEndpointAttributesInputShapeAWSQuery(v *types.SetEndpointAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalSetPlatformApplicationAttributesInputShapeAWSQuery(v *types.SetPlatformApplicationAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalSetSMSAttributesInputShapeAWSQuery(v *types.SetSMSAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalSetSubscriptionAttributesInputShapeAWSQuery(v *types.SetSubscriptionAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalSetTopicAttributesInputShapeAWSQuery(v *types.SetTopicAttributesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalSubscribeInputShapeAWSQuery(v *types.SubscribeInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalTagResourceInputShapeAWSQuery(v *types.TagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalUnsubscribeInputShapeAWSQuery(v *types.UnsubscribeInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}

func MarshalUntagResourceInputShapeAWSQuery(v *types.UntagResourceInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	query.Build(r)
	return nil
}
func (m AddPermissionMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalAddPermissionInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m AddPermissionMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "AddPermission.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CheckIfPhoneNumberIsOptedOutMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCheckIfPhoneNumberIsOptedOutInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CheckIfPhoneNumberIsOptedOutMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CheckIfPhoneNumberIsOptedOut.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ConfirmSubscriptionMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalConfirmSubscriptionInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ConfirmSubscriptionMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ConfirmSubscription.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreatePlatformApplicationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreatePlatformApplicationInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreatePlatformApplicationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreatePlatformApplication.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreatePlatformEndpointMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreatePlatformEndpointInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreatePlatformEndpointMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreatePlatformEndpoint.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m CreateTopicMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateTopicInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateTopicMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateTopic.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteEndpointMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteEndpointInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteEndpointMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteEndpoint.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeletePlatformApplicationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeletePlatformApplicationInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeletePlatformApplicationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeletePlatformApplication.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteTopicMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteTopicInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteTopicMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteTopic.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetEndpointAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetEndpointAttributesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetEndpointAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetEndpointAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetPlatformApplicationAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetPlatformApplicationAttributesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetPlatformApplicationAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetPlatformApplicationAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetSMSAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetSMSAttributesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetSMSAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetSMSAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetSubscriptionAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetSubscriptionAttributesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetSubscriptionAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetSubscriptionAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetTopicAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetTopicAttributesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetTopicAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetTopicAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListEndpointsByPlatformApplicationMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListEndpointsByPlatformApplicationInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListEndpointsByPlatformApplicationMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListEndpointsByPlatformApplication.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListPhoneNumbersOptedOutMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListPhoneNumbersOptedOutInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListPhoneNumbersOptedOutMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListPhoneNumbersOptedOut.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListPlatformApplicationsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListPlatformApplicationsInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListPlatformApplicationsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListPlatformApplications.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListSubscriptionsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListSubscriptionsInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListSubscriptionsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListSubscriptions.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListSubscriptionsByTopicMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListSubscriptionsByTopicInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListSubscriptionsByTopicMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListSubscriptionsByTopic.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListTagsForResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListTagsForResourceInputShapeAWSQuery(m.Input, r)
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

func (m ListTopicsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListTopicsInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListTopicsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListTopics.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m OptInPhoneNumberMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalOptInPhoneNumberInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m OptInPhoneNumberMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "OptInPhoneNumber.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m PublishMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalPublishInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m PublishMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "Publish.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m RemovePermissionMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalRemovePermissionInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m RemovePermissionMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "RemovePermission.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m SetEndpointAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalSetEndpointAttributesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m SetEndpointAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "SetEndpointAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m SetPlatformApplicationAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalSetPlatformApplicationAttributesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m SetPlatformApplicationAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "SetPlatformApplicationAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m SetSMSAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalSetSMSAttributesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m SetSMSAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "SetSMSAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m SetSubscriptionAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalSetSubscriptionAttributesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m SetSubscriptionAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "SetSubscriptionAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m SetTopicAttributesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalSetTopicAttributesInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m SetTopicAttributesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "SetTopicAttributes.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m SubscribeMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalSubscribeInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m SubscribeMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "Subscribe.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m TagResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalTagResourceInputShapeAWSQuery(m.Input, r)
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

func (m UnsubscribeMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUnsubscribeInputShapeAWSQuery(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UnsubscribeMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "Unsubscribe.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UntagResourceMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUntagResourceInputShapeAWSQuery(m.Input, r)
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
