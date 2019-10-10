// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPatchClusterRoleBindingParams creates a new PatchClusterRoleBindingParams object
// with the default values initialized.
func NewPatchClusterRoleBindingParams() *PatchClusterRoleBindingParams {
	var ()
	return &PatchClusterRoleBindingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchClusterRoleBindingParamsWithTimeout creates a new PatchClusterRoleBindingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchClusterRoleBindingParamsWithTimeout(timeout time.Duration) *PatchClusterRoleBindingParams {
	var ()
	return &PatchClusterRoleBindingParams{

		timeout: timeout,
	}
}

// NewPatchClusterRoleBindingParamsWithContext creates a new PatchClusterRoleBindingParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchClusterRoleBindingParamsWithContext(ctx context.Context) *PatchClusterRoleBindingParams {
	var ()
	return &PatchClusterRoleBindingParams{

		Context: ctx,
	}
}

// NewPatchClusterRoleBindingParamsWithHTTPClient creates a new PatchClusterRoleBindingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchClusterRoleBindingParamsWithHTTPClient(client *http.Client) *PatchClusterRoleBindingParams {
	var ()
	return &PatchClusterRoleBindingParams{
		HTTPClient: client,
	}
}

/*PatchClusterRoleBindingParams contains all the parameters to send to the API endpoint
for the patch cluster role binding operation typically these are written to a http.Request
*/
type PatchClusterRoleBindingParams struct {

	/*Patch*/
	Patch []uint8
	/*BindingID*/
	BindingID string
	/*ClusterID*/
	ClusterID string
	/*Dc*/
	Dc string
	/*ProjectID*/
	ProjectID string
	/*RoleID*/
	RoleID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) WithTimeout(timeout time.Duration) *PatchClusterRoleBindingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) WithContext(ctx context.Context) *PatchClusterRoleBindingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) WithHTTPClient(client *http.Client) *PatchClusterRoleBindingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPatch adds the patch to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) WithPatch(patch []uint8) *PatchClusterRoleBindingParams {
	o.SetPatch(patch)
	return o
}

// SetPatch adds the patch to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) SetPatch(patch []uint8) {
	o.Patch = patch
}

// WithBindingID adds the bindingID to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) WithBindingID(bindingID string) *PatchClusterRoleBindingParams {
	o.SetBindingID(bindingID)
	return o
}

// SetBindingID adds the bindingId to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) SetBindingID(bindingID string) {
	o.BindingID = bindingID
}

// WithClusterID adds the clusterID to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) WithClusterID(clusterID string) *PatchClusterRoleBindingParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithDc adds the dc to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) WithDc(dc string) *PatchClusterRoleBindingParams {
	o.SetDc(dc)
	return o
}

// SetDc adds the dc to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) SetDc(dc string) {
	o.Dc = dc
}

// WithProjectID adds the projectID to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) WithProjectID(projectID string) *PatchClusterRoleBindingParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WithRoleID adds the roleID to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) WithRoleID(roleID string) *PatchClusterRoleBindingParams {
	o.SetRoleID(roleID)
	return o
}

// SetRoleID adds the roleId to the patch cluster role binding params
func (o *PatchClusterRoleBindingParams) SetRoleID(roleID string) {
	o.RoleID = roleID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchClusterRoleBindingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Patch != nil {
		if err := r.SetBodyParam(o.Patch); err != nil {
			return err
		}
	}

	// path param binding_id
	if err := r.SetPathParam("binding_id", o.BindingID); err != nil {
		return err
	}

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param dc
	if err := r.SetPathParam("dc", o.Dc); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	// path param role_id
	if err := r.SetPathParam("role_id", o.RoleID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}