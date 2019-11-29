// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package aws_restjson service, marshal.go contains the generated operation marshalers and methods on it.

package aws_restjson

import (
	"bytes"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/private/protocol/jsonrpc"
	"github.com/aws/aws-sdk-go-v2/private/protocol/rest"
	"github.com/aws/aws-sdk-go-v2/service/resourcegroups/types"
)

var _ bytes.Buffer
var _ awserr.Error

// CreateGroupMarshaler defines marshaler for CreateGroup operation
type CreateGroupMarshaler struct {
	Input *types.CreateGroupInput
}

// DeleteGroupMarshaler defines marshaler for DeleteGroup operation
type DeleteGroupMarshaler struct {
	Input *types.DeleteGroupInput
}

// GetGroupMarshaler defines marshaler for GetGroup operation
type GetGroupMarshaler struct {
	Input *types.GetGroupInput
}

// GetGroupQueryMarshaler defines marshaler for GetGroupQuery operation
type GetGroupQueryMarshaler struct {
	Input *types.GetGroupQueryInput
}

// GetTagsMarshaler defines marshaler for GetTags operation
type GetTagsMarshaler struct {
	Input *types.GetTagsInput
}

// ListGroupResourcesMarshaler defines marshaler for ListGroupResources operation
type ListGroupResourcesMarshaler struct {
	Input *types.ListGroupResourcesInput
}

// ListGroupsMarshaler defines marshaler for ListGroups operation
type ListGroupsMarshaler struct {
	Input *types.ListGroupsInput
}

// SearchResourcesMarshaler defines marshaler for SearchResources operation
type SearchResourcesMarshaler struct {
	Input *types.SearchResourcesInput
}

// TagMarshaler defines marshaler for Tag operation
type TagMarshaler struct {
	Input *types.TagInput
}

// UntagMarshaler defines marshaler for Untag operation
type UntagMarshaler struct {
	Input *types.UntagInput
}

// UpdateGroupMarshaler defines marshaler for UpdateGroup operation
type UpdateGroupMarshaler struct {
	Input *types.UpdateGroupInput
}

// UpdateGroupQueryMarshaler defines marshaler for UpdateGroupQuery operation
type UpdateGroupQueryMarshaler struct {
	Input *types.UpdateGroupQueryInput
}

func MarshalCreateGroupInputShapeAWSREST(v *types.CreateGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalCreateGroupInputShapeAWSJSON(v *types.CreateGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalDeleteGroupInputShapeAWSREST(v *types.DeleteGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalDeleteGroupInputShapeAWSJSON(v *types.DeleteGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalGetGroupInputShapeAWSREST(v *types.GetGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalGetGroupInputShapeAWSJSON(v *types.GetGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalGetGroupQueryInputShapeAWSREST(v *types.GetGroupQueryInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalGetGroupQueryInputShapeAWSJSON(v *types.GetGroupQueryInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalGetTagsInputShapeAWSREST(v *types.GetTagsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalGetTagsInputShapeAWSJSON(v *types.GetTagsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListGroupResourcesInputShapeAWSREST(v *types.ListGroupResourcesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListGroupResourcesInputShapeAWSJSON(v *types.ListGroupResourcesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalListGroupsInputShapeAWSREST(v *types.ListGroupsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalListGroupsInputShapeAWSJSON(v *types.ListGroupsInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalSearchResourcesInputShapeAWSREST(v *types.SearchResourcesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalSearchResourcesInputShapeAWSJSON(v *types.SearchResourcesInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalTagInputShapeAWSREST(v *types.TagInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalTagInputShapeAWSJSON(v *types.TagInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalUntagInputShapeAWSREST(v *types.UntagInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUntagInputShapeAWSJSON(v *types.UntagInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalUpdateGroupInputShapeAWSREST(v *types.UpdateGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUpdateGroupInputShapeAWSJSON(v *types.UpdateGroupInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}

func MarshalUpdateGroupQueryInputShapeAWSREST(v *types.UpdateGroupQueryInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	rest.Build(r)
	return nil
}
func MarshalUpdateGroupQueryInputShapeAWSJSON(v *types.UpdateGroupQueryInput, r *aws.Request) error {
	// delegate to reflection based marshaling
	if t := rest.PayloadType(r.Params); t == "structure" || t == "" {
		jsonrpc.Build(r)
	}
	return nil
}
func (m CreateGroupMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalCreateGroupInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalCreateGroupInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m CreateGroupMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "CreateGroup.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m DeleteGroupMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalDeleteGroupInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalDeleteGroupInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m DeleteGroupMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "DeleteGroup.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetGroupMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetGroupInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalGetGroupInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetGroupMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetGroup.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetGroupQueryMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetGroupQueryInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalGetGroupQueryInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetGroupQueryMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetGroupQuery.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m GetTagsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalGetTagsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalGetTagsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m GetTagsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "GetTags.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListGroupResourcesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListGroupResourcesInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListGroupResourcesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListGroupResourcesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListGroupResources.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m ListGroupsMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalListGroupsInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalListGroupsInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m ListGroupsMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "ListGroups.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m SearchResourcesMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalSearchResourcesInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalSearchResourcesInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m SearchResourcesMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "SearchResources.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m TagMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalTagInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalTagInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m TagMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "Tag.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UntagMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUntagInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUntagInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UntagMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "Untag.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateGroupMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateGroupInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUpdateGroupInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateGroupMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateGroup.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}

func (m UpdateGroupQueryMarshaler) MarshalOperation(r *aws.Request) {
	var err error

	err = MarshalUpdateGroupQueryInputShapeAWSREST(m.Input, r)
	if err != nil {
		r.Error = err
	}

	err = MarshalUpdateGroupQueryInputShapeAWSJSON(m.Input, r)
	if err != nil {
		r.Error = err
	}

}

// GetNamedBuildHandler returns a Named Build Handler for an operation marshal function
func (m UpdateGroupQueryMarshaler) GetNamedBuildHandler() aws.NamedHandler {
	const BuildHandler = "UpdateGroupQuery.BuildHandler"
	return aws.NamedHandler{
		Name: BuildHandler,
		Fn:   m.MarshalOperation,
	}
}
