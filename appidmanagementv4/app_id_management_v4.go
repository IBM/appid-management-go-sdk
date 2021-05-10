/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 99-SNAPSHOT-a675267a-20210510-115948
 */

// Package appidmanagementv4 : Operations and models for the AppIdManagementV4 service
package appidmanagementv4

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	common "github.com/IBM/appid-go-sdk/common"
	"github.com/IBM/go-sdk-core/v5/core"
)

// AppIdManagementV4 : You can use the following APIs to configure your instances of IBM Cloud App ID. To define fine
// grain access policies, you must have an instance of App ID that was created after March 15, 2018.</br> New to the
// APIs? Try them out by using the <a href="https://github.com/ibm-cloud-security/appid-postman">App ID Postman
// collection</a>!</br> </br> <b>Important:</b> You must have an <a
// href="https://cloud.ibm.com/docs/account?topic=account-iamoverview">IBM Cloud Identity and Access Management</a>
// token to access the APIs. For help obtaining a token, check out <a
// href="https://cloud.ibm.com/docs/account?topic=account-iamtoken_from_apikey">Getting an IAM token with an API
// key</a>.
//
// Version: 4
type AppIdManagementV4 struct {
	Service *core.BaseService

	// The service tenantId. The tenantId can be found in the service credentials.
	TenantID *string
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://app-id-management.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "app_id_management"

// AppIdManagementV4Options : Service options
type AppIdManagementV4Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator

	// The service tenantId. The tenantId can be found in the service credentials.
	TenantID *string `validate:"required"`
}

// NewAppIdManagementV4UsingExternalConfig : constructs an instance of AppIdManagementV4 with passed in options and external configuration.
func NewAppIdManagementV4UsingExternalConfig(options *AppIdManagementV4Options) (appIdManagement *AppIdManagementV4, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	appIdManagement, err = NewAppIdManagementV4(options)
	if err != nil {
		return
	}

	err = appIdManagement.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = appIdManagement.Service.SetServiceURL(options.URL)
	}
	return
}

// NewAppIdManagementV4 : constructs an instance of AppIdManagementV4 with passed in options.
func NewAppIdManagementV4(options *AppIdManagementV4Options) (service *AppIdManagementV4, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	err = core.ValidateStruct(options, "options")
	if err != nil {
		return
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &AppIdManagementV4{
		Service: baseService,
		TenantID: options.TenantID,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "appIdManagement" suitable for processing requests.
func (appIdManagement *AppIdManagementV4) Clone() *AppIdManagementV4 {
	if core.IsNil(appIdManagement) {
		return nil
	}
	clone := *appIdManagement
	clone.Service = appIdManagement.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (appIdManagement *AppIdManagementV4) SetServiceURL(url string) error {
	return appIdManagement.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (appIdManagement *AppIdManagementV4) GetServiceURL() string {
	return appIdManagement.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (appIdManagement *AppIdManagementV4) SetDefaultHeaders(headers http.Header) {
	appIdManagement.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (appIdManagement *AppIdManagementV4) SetEnableGzipCompression(enableGzip bool) {
	appIdManagement.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (appIdManagement *AppIdManagementV4) GetEnableGzipCompression() bool {
	return appIdManagement.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (appIdManagement *AppIdManagementV4) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	appIdManagement.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (appIdManagement *AppIdManagementV4) DisableRetries() {
	appIdManagement.Service.DisableRetries()
}

// ListApplications : List applications
// Returns a list of all applications registered with the App ID Instance.
func (appIdManagement *AppIdManagementV4) ListApplications(listApplicationsOptions *ListApplicationsOptions) (result *ApplicationsList, response *core.DetailedResponse, err error) {
	return appIdManagement.ListApplicationsWithContext(context.Background(), listApplicationsOptions)
}

// ListApplicationsWithContext is an alternate form of the ListApplications method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) ListApplicationsWithContext(ctx context.Context, listApplicationsOptions *ListApplicationsOptions) (result *ApplicationsList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listApplicationsOptions, "listApplicationsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/applications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listApplicationsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "ListApplications")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApplicationsList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// RegisterApplication : Create application
// Register a new application with the App ID instance.
func (appIdManagement *AppIdManagementV4) RegisterApplication(registerApplicationOptions *RegisterApplicationOptions) (result *Application, response *core.DetailedResponse, err error) {
	return appIdManagement.RegisterApplicationWithContext(context.Background(), registerApplicationOptions)
}

// RegisterApplicationWithContext is an alternate form of the RegisterApplication method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) RegisterApplicationWithContext(ctx context.Context, registerApplicationOptions *RegisterApplicationOptions) (result *Application, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(registerApplicationOptions, "registerApplicationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(registerApplicationOptions, "registerApplicationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/applications`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range registerApplicationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "RegisterApplication")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if registerApplicationOptions.Name != nil {
		body["name"] = registerApplicationOptions.Name
	}
	if registerApplicationOptions.Type != nil {
		body["type"] = registerApplicationOptions.Type
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApplication)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetApplication : Get application
// Returns a specific application registered with the App ID Instance.
func (appIdManagement *AppIdManagementV4) GetApplication(getApplicationOptions *GetApplicationOptions) (result *Application, response *core.DetailedResponse, err error) {
	return appIdManagement.GetApplicationWithContext(context.Background(), getApplicationOptions)
}

// GetApplicationWithContext is an alternate form of the GetApplication method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetApplicationWithContext(ctx context.Context, getApplicationOptions *GetApplicationOptions) (result *Application, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getApplicationOptions, "getApplicationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getApplicationOptions, "getApplicationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"clientId": *getApplicationOptions.ClientID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/applications/{clientId}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getApplicationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetApplication")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApplication)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateApplication : Update application
// Update an application registered with the App ID instance.
func (appIdManagement *AppIdManagementV4) UpdateApplication(updateApplicationOptions *UpdateApplicationOptions) (result *Application, response *core.DetailedResponse, err error) {
	return appIdManagement.UpdateApplicationWithContext(context.Background(), updateApplicationOptions)
}

// UpdateApplicationWithContext is an alternate form of the UpdateApplication method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UpdateApplicationWithContext(ctx context.Context, updateApplicationOptions *UpdateApplicationOptions) (result *Application, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateApplicationOptions, "updateApplicationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateApplicationOptions, "updateApplicationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"clientId": *updateApplicationOptions.ClientID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/applications/{clientId}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateApplicationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UpdateApplication")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateApplicationOptions.Name != nil {
		body["name"] = updateApplicationOptions.Name
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApplication)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteApplication : Delete application
// Delete an application registered with the App ID instance. Note: This action cannot be undone.
func (appIdManagement *AppIdManagementV4) DeleteApplication(deleteApplicationOptions *DeleteApplicationOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.DeleteApplicationWithContext(context.Background(), deleteApplicationOptions)
}

// DeleteApplicationWithContext is an alternate form of the DeleteApplication method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) DeleteApplicationWithContext(ctx context.Context, deleteApplicationOptions *DeleteApplicationOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteApplicationOptions, "deleteApplicationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteApplicationOptions, "deleteApplicationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"clientId": *deleteApplicationOptions.ClientID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/applications/{clientId}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteApplicationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "DeleteApplication")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// GetApplicationScopes : Get application scopes
// View the defined scopes for an application that is registered with an App ID instance.
func (appIdManagement *AppIdManagementV4) GetApplicationScopes(getApplicationScopesOptions *GetApplicationScopesOptions) (result *GetScopesForApplication, response *core.DetailedResponse, err error) {
	return appIdManagement.GetApplicationScopesWithContext(context.Background(), getApplicationScopesOptions)
}

// GetApplicationScopesWithContext is an alternate form of the GetApplicationScopes method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetApplicationScopesWithContext(ctx context.Context, getApplicationScopesOptions *GetApplicationScopesOptions) (result *GetScopesForApplication, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getApplicationScopesOptions, "getApplicationScopesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getApplicationScopesOptions, "getApplicationScopesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"clientId": *getApplicationScopesOptions.ClientID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/applications/{clientId}/scopes`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getApplicationScopesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetApplicationScopes")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetScopesForApplication)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PutApplicationsScopes : Add application scope
// Update the scopes for a registered application.</br> <b>Important</b>: Removing a scope from an array deletes it from
// any roles that it is associated with and the action cannot be undone.
func (appIdManagement *AppIdManagementV4) PutApplicationsScopes(putApplicationsScopesOptions *PutApplicationsScopesOptions) (result *GetScopesForApplication, response *core.DetailedResponse, err error) {
	return appIdManagement.PutApplicationsScopesWithContext(context.Background(), putApplicationsScopesOptions)
}

// PutApplicationsScopesWithContext is an alternate form of the PutApplicationsScopes method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) PutApplicationsScopesWithContext(ctx context.Context, putApplicationsScopesOptions *PutApplicationsScopesOptions) (result *GetScopesForApplication, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(putApplicationsScopesOptions, "putApplicationsScopesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(putApplicationsScopesOptions, "putApplicationsScopesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"clientId": *putApplicationsScopesOptions.ClientID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/applications/{clientId}/scopes`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range putApplicationsScopesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "PutApplicationsScopes")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if putApplicationsScopesOptions.Scopes != nil {
		body["scopes"] = putApplicationsScopesOptions.Scopes
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetScopesForApplication)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetApplicationRoles : Get application roles
// View the defined roles for an application that is registered with an App ID instance.
func (appIdManagement *AppIdManagementV4) GetApplicationRoles(getApplicationRolesOptions *GetApplicationRolesOptions) (result *GetUserRolesResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.GetApplicationRolesWithContext(context.Background(), getApplicationRolesOptions)
}

// GetApplicationRolesWithContext is an alternate form of the GetApplicationRoles method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetApplicationRolesWithContext(ctx context.Context, getApplicationRolesOptions *GetApplicationRolesOptions) (result *GetUserRolesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getApplicationRolesOptions, "getApplicationRolesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getApplicationRolesOptions, "getApplicationRolesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"clientId": *getApplicationRolesOptions.ClientID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/applications/{clientId}/roles`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getApplicationRolesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetApplicationRoles")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetUserRolesResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PutApplicationsRoles : Add application role
// Update the roles for a registered application.</br>.
func (appIdManagement *AppIdManagementV4) PutApplicationsRoles(putApplicationsRolesOptions *PutApplicationsRolesOptions) (result *AssignRoleToUser, response *core.DetailedResponse, err error) {
	return appIdManagement.PutApplicationsRolesWithContext(context.Background(), putApplicationsRolesOptions)
}

// PutApplicationsRolesWithContext is an alternate form of the PutApplicationsRoles method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) PutApplicationsRolesWithContext(ctx context.Context, putApplicationsRolesOptions *PutApplicationsRolesOptions) (result *AssignRoleToUser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(putApplicationsRolesOptions, "putApplicationsRolesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(putApplicationsRolesOptions, "putApplicationsRolesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"clientId": *putApplicationsRolesOptions.ClientID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/applications/{clientId}/roles`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range putApplicationsRolesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "PutApplicationsRoles")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if putApplicationsRolesOptions.Roles != nil {
		body["roles"] = putApplicationsRolesOptions.Roles
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAssignRoleToUser)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListCloudDirectoryUsers : List Cloud Directory users
// Get the list of Cloud Directory users. <a href="https://cloud.ibm.com/docs/appid?topic=appid-cloud-directory"
// target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) ListCloudDirectoryUsers(listCloudDirectoryUsersOptions *ListCloudDirectoryUsersOptions) (result *UsersList, response *core.DetailedResponse, err error) {
	return appIdManagement.ListCloudDirectoryUsersWithContext(context.Background(), listCloudDirectoryUsersOptions)
}

// ListCloudDirectoryUsersWithContext is an alternate form of the ListCloudDirectoryUsers method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) ListCloudDirectoryUsersWithContext(ctx context.Context, listCloudDirectoryUsersOptions *ListCloudDirectoryUsersOptions) (result *UsersList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listCloudDirectoryUsersOptions, "listCloudDirectoryUsersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/Users`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCloudDirectoryUsersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "ListCloudDirectoryUsers")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listCloudDirectoryUsersOptions.StartIndex != nil {
		builder.AddQuery("startIndex", fmt.Sprint(*listCloudDirectoryUsersOptions.StartIndex))
	}
	if listCloudDirectoryUsersOptions.Count != nil {
		builder.AddQuery("count", fmt.Sprint(*listCloudDirectoryUsersOptions.Count))
	}
	if listCloudDirectoryUsersOptions.Query != nil {
		builder.AddQuery("query", fmt.Sprint(*listCloudDirectoryUsersOptions.Query))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUsersList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateCloudDirectoryUser : Create a Cloud Directory user
// Create a new record for Cloud Directory (no verification email is sent, and no profile is created).</br> To create a
// new Cloud Directory user use the  <a href="/swagger-ui/#/Management API - Cloud Directory Workflows/mgmt.startSignUp"
// target="_blank">sign_up</a> API. <a href="https://cloud.ibm.com/docs/appid?topic=appid-cloud-directory"
// target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) CreateCloudDirectoryUser(createCloudDirectoryUserOptions *CreateCloudDirectoryUserOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.CreateCloudDirectoryUserWithContext(context.Background(), createCloudDirectoryUserOptions)
}

// CreateCloudDirectoryUserWithContext is an alternate form of the CreateCloudDirectoryUser method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) CreateCloudDirectoryUserWithContext(ctx context.Context, createCloudDirectoryUserOptions *CreateCloudDirectoryUserOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCloudDirectoryUserOptions, "createCloudDirectoryUserOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCloudDirectoryUserOptions, "createCloudDirectoryUserOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/Users`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCloudDirectoryUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "CreateCloudDirectoryUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createCloudDirectoryUserOptions.Emails != nil {
		body["emails"] = createCloudDirectoryUserOptions.Emails
	}
	if createCloudDirectoryUserOptions.Password != nil {
		body["password"] = createCloudDirectoryUserOptions.Password
	}
	if createCloudDirectoryUserOptions.Active != nil {
		body["active"] = createCloudDirectoryUserOptions.Active
	}
	if createCloudDirectoryUserOptions.UserName != nil {
		body["userName"] = createCloudDirectoryUserOptions.UserName
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// GetCloudDirectoryUser : Get a Cloud Directory user
// Returns the requested Cloud Directory user object. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cloud-directory" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) GetCloudDirectoryUser(getCloudDirectoryUserOptions *GetCloudDirectoryUserOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.GetCloudDirectoryUserWithContext(context.Background(), getCloudDirectoryUserOptions)
}

// GetCloudDirectoryUserWithContext is an alternate form of the GetCloudDirectoryUser method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetCloudDirectoryUserWithContext(ctx context.Context, getCloudDirectoryUserOptions *GetCloudDirectoryUserOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCloudDirectoryUserOptions, "getCloudDirectoryUserOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCloudDirectoryUserOptions, "getCloudDirectoryUserOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"userId": *getCloudDirectoryUserOptions.UserID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/Users/{userId}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCloudDirectoryUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetCloudDirectoryUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// UpdateCloudDirectoryUser : Update a Cloud Directory user
// Updates an existing Cloud Directory user. <a href="https://cloud.ibm.com/docs/appid?topic=appid-cd-users"
// target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) UpdateCloudDirectoryUser(updateCloudDirectoryUserOptions *UpdateCloudDirectoryUserOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.UpdateCloudDirectoryUserWithContext(context.Background(), updateCloudDirectoryUserOptions)
}

// UpdateCloudDirectoryUserWithContext is an alternate form of the UpdateCloudDirectoryUser method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UpdateCloudDirectoryUserWithContext(ctx context.Context, updateCloudDirectoryUserOptions *UpdateCloudDirectoryUserOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCloudDirectoryUserOptions, "updateCloudDirectoryUserOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCloudDirectoryUserOptions, "updateCloudDirectoryUserOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"userId": *updateCloudDirectoryUserOptions.UserID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/Users/{userId}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCloudDirectoryUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UpdateCloudDirectoryUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateCloudDirectoryUserOptions.Emails != nil {
		body["emails"] = updateCloudDirectoryUserOptions.Emails
	}
	if updateCloudDirectoryUserOptions.Active != nil {
		body["active"] = updateCloudDirectoryUserOptions.Active
	}
	if updateCloudDirectoryUserOptions.UserName != nil {
		body["userName"] = updateCloudDirectoryUserOptions.UserName
	}
	if updateCloudDirectoryUserOptions.Password != nil {
		body["password"] = updateCloudDirectoryUserOptions.Password
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// DeleteCloudDirectoryUser : Delete a Cloud Directory user
// Deletes an existing Cloud Directory recored (without removing the associated profile). <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cd-users" target="_blank">Learn more</a>.</br> To remove a Cloud
// Directory user use the <a href="/swagger-ui/#/Management API - Cloud Directory Workflows/mgmt.cloud_directory_remove"
// target="_blank">remove</a> API. </br> <b>Note: This action cannot be undone</b>.
func (appIdManagement *AppIdManagementV4) DeleteCloudDirectoryUser(deleteCloudDirectoryUserOptions *DeleteCloudDirectoryUserOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.DeleteCloudDirectoryUserWithContext(context.Background(), deleteCloudDirectoryUserOptions)
}

// DeleteCloudDirectoryUserWithContext is an alternate form of the DeleteCloudDirectoryUser method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) DeleteCloudDirectoryUserWithContext(ctx context.Context, deleteCloudDirectoryUserOptions *DeleteCloudDirectoryUserOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCloudDirectoryUserOptions, "deleteCloudDirectoryUserOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCloudDirectoryUserOptions, "deleteCloudDirectoryUserOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"userId": *deleteCloudDirectoryUserOptions.UserID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/Users/{userId}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCloudDirectoryUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "DeleteCloudDirectoryUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// SsoLogoutFromAllApps : Invalidate all SSO sessions
// Invalidate all the user's SSO sessions. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cd-sso#ending-all-sessions-for-a-user" target="_blank">Learn
// more</a>.
func (appIdManagement *AppIdManagementV4) SsoLogoutFromAllApps(ssoLogoutFromAllAppsOptions *SsoLogoutFromAllAppsOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.SsoLogoutFromAllAppsWithContext(context.Background(), ssoLogoutFromAllAppsOptions)
}

// SsoLogoutFromAllAppsWithContext is an alternate form of the SsoLogoutFromAllApps method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) SsoLogoutFromAllAppsWithContext(ctx context.Context, ssoLogoutFromAllAppsOptions *SsoLogoutFromAllAppsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(ssoLogoutFromAllAppsOptions, "ssoLogoutFromAllAppsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(ssoLogoutFromAllAppsOptions, "ssoLogoutFromAllAppsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"userId": *ssoLogoutFromAllAppsOptions.UserID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/Users/{userId}/sso/logout`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range ssoLogoutFromAllAppsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "SsoLogoutFromAllApps")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// CloudDirectoryExport : Export Cloud Directory users
// Exports Cloud Directory users with their profile attributes and hashed passwords. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cd-users" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) CloudDirectoryExport(cloudDirectoryExportOptions *CloudDirectoryExportOptions) (result *ExportUser, response *core.DetailedResponse, err error) {
	return appIdManagement.CloudDirectoryExportWithContext(context.Background(), cloudDirectoryExportOptions)
}

// CloudDirectoryExportWithContext is an alternate form of the CloudDirectoryExport method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) CloudDirectoryExportWithContext(ctx context.Context, cloudDirectoryExportOptions *CloudDirectoryExportOptions) (result *ExportUser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(cloudDirectoryExportOptions, "cloudDirectoryExportOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(cloudDirectoryExportOptions, "cloudDirectoryExportOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/export`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range cloudDirectoryExportOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "CloudDirectoryExport")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("encryption_secret", fmt.Sprint(*cloudDirectoryExportOptions.EncryptionSecret))
	if cloudDirectoryExportOptions.StartIndex != nil {
		builder.AddQuery("startIndex", fmt.Sprint(*cloudDirectoryExportOptions.StartIndex))
	}
	if cloudDirectoryExportOptions.Count != nil {
		builder.AddQuery("count", fmt.Sprint(*cloudDirectoryExportOptions.Count))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExportUser)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CloudDirectoryImport : Import Cloud Directory users
// Imports Cloud Directory users list that was exported using the /export API. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cd-users" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) CloudDirectoryImport(cloudDirectoryImportOptions *CloudDirectoryImportOptions) (result *ImportResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.CloudDirectoryImportWithContext(context.Background(), cloudDirectoryImportOptions)
}

// CloudDirectoryImportWithContext is an alternate form of the CloudDirectoryImport method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) CloudDirectoryImportWithContext(ctx context.Context, cloudDirectoryImportOptions *CloudDirectoryImportOptions) (result *ImportResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(cloudDirectoryImportOptions, "cloudDirectoryImportOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(cloudDirectoryImportOptions, "cloudDirectoryImportOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/import`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range cloudDirectoryImportOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "CloudDirectoryImport")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("encryption_secret", fmt.Sprint(*cloudDirectoryImportOptions.EncryptionSecret))

	body := make(map[string]interface{})
	if cloudDirectoryImportOptions.Users != nil {
		body["users"] = cloudDirectoryImportOptions.Users
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalImportResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CloudDirectoryGetUserinfo : Get Cloud Directory SCIM and Attributes
// Returns the Cloud Directory user SCIM and the Profile related to it. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cd-users" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) CloudDirectoryGetUserinfo(cloudDirectoryGetUserinfoOptions *CloudDirectoryGetUserinfoOptions) (result *GetUserAndProfile, response *core.DetailedResponse, err error) {
	return appIdManagement.CloudDirectoryGetUserinfoWithContext(context.Background(), cloudDirectoryGetUserinfoOptions)
}

// CloudDirectoryGetUserinfoWithContext is an alternate form of the CloudDirectoryGetUserinfo method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) CloudDirectoryGetUserinfoWithContext(ctx context.Context, cloudDirectoryGetUserinfoOptions *CloudDirectoryGetUserinfoOptions) (result *GetUserAndProfile, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(cloudDirectoryGetUserinfoOptions, "cloudDirectoryGetUserinfoOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(cloudDirectoryGetUserinfoOptions, "cloudDirectoryGetUserinfoOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"userId": *cloudDirectoryGetUserinfoOptions.UserID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/{userId}/userinfo`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range cloudDirectoryGetUserinfoOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "CloudDirectoryGetUserinfo")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetUserAndProfile)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// StartSignUp : Sign up
// Start the sign up process <a href="https://cloud.ibm.com/docs/appid?topic=appid-branded" target="_blank">Learn
// more</a>.
func (appIdManagement *AppIdManagementV4) StartSignUp(startSignUpOptions *StartSignUpOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.StartSignUpWithContext(context.Background(), startSignUpOptions)
}

// StartSignUpWithContext is an alternate form of the StartSignUp method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) StartSignUpWithContext(ctx context.Context, startSignUpOptions *StartSignUpOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(startSignUpOptions, "startSignUpOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(startSignUpOptions, "startSignUpOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/sign_up`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range startSignUpOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "StartSignUp")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	builder.AddQuery("shouldCreateProfile", fmt.Sprint(*startSignUpOptions.ShouldCreateProfile))
	if startSignUpOptions.Language != nil {
		builder.AddQuery("language", fmt.Sprint(*startSignUpOptions.Language))
	}

	body := make(map[string]interface{})
	if startSignUpOptions.Emails != nil {
		body["emails"] = startSignUpOptions.Emails
	}
	if startSignUpOptions.Password != nil {
		body["password"] = startSignUpOptions.Password
	}
	if startSignUpOptions.Active != nil {
		body["active"] = startSignUpOptions.Active
	}
	if startSignUpOptions.UserName != nil {
		body["userName"] = startSignUpOptions.UserName
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// UserVerificationResult : Get signup confirmation result
// Returns the sign up confirmation result. <a href="https://cloud.ibm.com/docs/appid?topic=appid-branded"
// target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) UserVerificationResult(userVerificationResultOptions *UserVerificationResultOptions) (result *ConfirmationResultOK, response *core.DetailedResponse, err error) {
	return appIdManagement.UserVerificationResultWithContext(context.Background(), userVerificationResultOptions)
}

// UserVerificationResultWithContext is an alternate form of the UserVerificationResult method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UserVerificationResultWithContext(ctx context.Context, userVerificationResultOptions *UserVerificationResultOptions) (result *ConfirmationResultOK, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(userVerificationResultOptions, "userVerificationResultOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(userVerificationResultOptions, "userVerificationResultOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/sign_up/confirmation_result`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range userVerificationResultOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UserVerificationResult")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if userVerificationResultOptions.Context != nil {
		body["context"] = userVerificationResultOptions.Context
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConfirmationResultOK)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// StartForgotPassword : Forgot password
// Starts the forgot password process. <a href="https://cloud.ibm.com/docs/appid?topic=appid-branded"
// target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) StartForgotPassword(startForgotPasswordOptions *StartForgotPasswordOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.StartForgotPasswordWithContext(context.Background(), startForgotPasswordOptions)
}

// StartForgotPasswordWithContext is an alternate form of the StartForgotPassword method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) StartForgotPasswordWithContext(ctx context.Context, startForgotPasswordOptions *StartForgotPasswordOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(startForgotPasswordOptions, "startForgotPasswordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(startForgotPasswordOptions, "startForgotPasswordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/forgot_password`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range startForgotPasswordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "StartForgotPassword")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if startForgotPasswordOptions.Language != nil {
		builder.AddQuery("language", fmt.Sprint(*startForgotPasswordOptions.Language))
	}

	body := make(map[string]interface{})
	if startForgotPasswordOptions.User != nil {
		body["user"] = startForgotPasswordOptions.User
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// ForgotPasswordResult : Forgot password confirmation result
// Returns the forgot password flow confirmation result. <a href="https://cloud.ibm.com/docs/appid?topic=appid-branded"
// target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) ForgotPasswordResult(forgotPasswordResultOptions *ForgotPasswordResultOptions) (result *ConfirmationResultOK, response *core.DetailedResponse, err error) {
	return appIdManagement.ForgotPasswordResultWithContext(context.Background(), forgotPasswordResultOptions)
}

// ForgotPasswordResultWithContext is an alternate form of the ForgotPasswordResult method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) ForgotPasswordResultWithContext(ctx context.Context, forgotPasswordResultOptions *ForgotPasswordResultOptions) (result *ConfirmationResultOK, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(forgotPasswordResultOptions, "forgotPasswordResultOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(forgotPasswordResultOptions, "forgotPasswordResultOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/forgot_password/confirmation_result`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range forgotPasswordResultOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "ForgotPasswordResult")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if forgotPasswordResultOptions.Context != nil {
		body["context"] = forgotPasswordResultOptions.Context
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalConfirmationResultOK)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ChangePassword : Change password
// Changes the Cloud Directory user password. <a href="https://cloud.ibm.com/docs/appid?topic=appid-branded"
// target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) ChangePassword(changePasswordOptions *ChangePasswordOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.ChangePasswordWithContext(context.Background(), changePasswordOptions)
}

// ChangePasswordWithContext is an alternate form of the ChangePassword method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) ChangePasswordWithContext(ctx context.Context, changePasswordOptions *ChangePasswordOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(changePasswordOptions, "changePasswordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(changePasswordOptions, "changePasswordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/change_password`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range changePasswordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "ChangePassword")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if changePasswordOptions.Language != nil {
		builder.AddQuery("language", fmt.Sprint(*changePasswordOptions.Language))
	}

	body := make(map[string]interface{})
	if changePasswordOptions.NewPassword != nil {
		body["newPassword"] = changePasswordOptions.NewPassword
	}
	if changePasswordOptions.UUID != nil {
		body["uuid"] = changePasswordOptions.UUID
	}
	if changePasswordOptions.ChangedIpAddress != nil {
		body["changedIpAddress"] = changePasswordOptions.ChangedIpAddress
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// ResendNotification : Resend user notifications
// Resend user email notifications (e.g. resend user verification email). <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-branded" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) ResendNotification(resendNotificationOptions *ResendNotificationOptions) (result *ResendNotificationResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.ResendNotificationWithContext(context.Background(), resendNotificationOptions)
}

// ResendNotificationWithContext is an alternate form of the ResendNotification method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) ResendNotificationWithContext(ctx context.Context, resendNotificationOptions *ResendNotificationOptions) (result *ResendNotificationResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(resendNotificationOptions, "resendNotificationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(resendNotificationOptions, "resendNotificationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"templateName": *resendNotificationOptions.TemplateName,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/resend/{templateName}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range resendNotificationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "ResendNotification")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	if resendNotificationOptions.Language != nil {
		builder.AddQuery("language", fmt.Sprint(*resendNotificationOptions.Language))
	}

	body := make(map[string]interface{})
	if resendNotificationOptions.UUID != nil {
		body["uuid"] = resendNotificationOptions.UUID
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalResendNotificationResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CloudDirectoryRemove : Delete Cloud Directory User and Profile
// Deletes an existing Cloud Directory user and the Profile related to it. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cd-users" target="_blank">Learn more</a>. <b>Note: This action
// cannot be undone</b>.
func (appIdManagement *AppIdManagementV4) CloudDirectoryRemove(cloudDirectoryRemoveOptions *CloudDirectoryRemoveOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.CloudDirectoryRemoveWithContext(context.Background(), cloudDirectoryRemoveOptions)
}

// CloudDirectoryRemoveWithContext is an alternate form of the CloudDirectoryRemove method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) CloudDirectoryRemoveWithContext(ctx context.Context, cloudDirectoryRemoveOptions *CloudDirectoryRemoveOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(cloudDirectoryRemoveOptions, "cloudDirectoryRemoveOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(cloudDirectoryRemoveOptions, "cloudDirectoryRemoveOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"userId": *cloudDirectoryRemoveOptions.UserID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/cloud_directory/remove/{userId}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range cloudDirectoryRemoveOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "CloudDirectoryRemove")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// GetTokensConfig : Get tokens configuration
// Returns the token configuration. <a href="https://cloud.ibm.com/docs/appid?topic=appid-key-concepts"
// target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) GetTokensConfig(getTokensConfigOptions *GetTokensConfigOptions) (result *TokensConfigResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.GetTokensConfigWithContext(context.Background(), getTokensConfigOptions)
}

// GetTokensConfigWithContext is an alternate form of the GetTokensConfig method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetTokensConfigWithContext(ctx context.Context, getTokensConfigOptions *GetTokensConfigOptions) (result *TokensConfigResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getTokensConfigOptions, "getTokensConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/tokens`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTokensConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetTokensConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTokensConfigResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PutTokensConfig : Update tokens configuration
// Update the tokens' configuration to fine-tune the expiration times of access, id and refresh tokens, to
// enable/disable refresh and anonymous tokens, and to configure custom claims. When a token config object is not
// included in the set, its value will be reset back to default. <br> For more information, check out the <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-key-concepts" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) PutTokensConfig(putTokensConfigOptions *PutTokensConfigOptions) (result *TokensConfigResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.PutTokensConfigWithContext(context.Background(), putTokensConfigOptions)
}

// PutTokensConfigWithContext is an alternate form of the PutTokensConfig method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) PutTokensConfigWithContext(ctx context.Context, putTokensConfigOptions *PutTokensConfigOptions) (result *TokensConfigResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(putTokensConfigOptions, "putTokensConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(putTokensConfigOptions, "putTokensConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/tokens`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range putTokensConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "PutTokensConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if putTokensConfigOptions.IdTokenClaims != nil {
		body["idTokenClaims"] = putTokensConfigOptions.IdTokenClaims
	}
	if putTokensConfigOptions.AccessTokenClaims != nil {
		body["accessTokenClaims"] = putTokensConfigOptions.AccessTokenClaims
	}
	if putTokensConfigOptions.Access != nil {
		body["access"] = putTokensConfigOptions.Access
	}
	if putTokensConfigOptions.Refresh != nil {
		body["refresh"] = putTokensConfigOptions.Refresh
	}
	if putTokensConfigOptions.AnonymousAccess != nil {
		body["anonymousAccess"] = putTokensConfigOptions.AnonymousAccess
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTokensConfigResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetRedirectUris : Get redirect URIs
// Returns the list of the redirect URIs that can be used as callbacks of App ID authentication flow. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-managing-idp#add-redirect-uri" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) GetRedirectUris(getRedirectUrisOptions *GetRedirectUrisOptions) (result *RedirectUriResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.GetRedirectUrisWithContext(context.Background(), getRedirectUrisOptions)
}

// GetRedirectUrisWithContext is an alternate form of the GetRedirectUris method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetRedirectUrisWithContext(ctx context.Context, getRedirectUrisOptions *GetRedirectUrisOptions) (result *RedirectUriResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getRedirectUrisOptions, "getRedirectUrisOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/redirect_uris`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRedirectUrisOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetRedirectUris")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRedirectUriResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateRedirectUris : Update redirect URIs
// Update the list of the redirect URIs that can be used as callbacks of App ID authentication flow. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-managing-idp#add-redirect-uri" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) UpdateRedirectUris(updateRedirectUrisOptions *UpdateRedirectUrisOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.UpdateRedirectUrisWithContext(context.Background(), updateRedirectUrisOptions)
}

// UpdateRedirectUrisWithContext is an alternate form of the UpdateRedirectUris method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UpdateRedirectUrisWithContext(ctx context.Context, updateRedirectUrisOptions *UpdateRedirectUrisOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateRedirectUrisOptions, "updateRedirectUrisOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateRedirectUrisOptions, "updateRedirectUrisOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/redirect_uris`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateRedirectUrisOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UpdateRedirectUris")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(updateRedirectUrisOptions.RedirectUrisArray)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// GetUserProfilesConfig : Get user profiles configuration
// A user profile is an entity that is stored and maintained by App ID. The profile holds a user's attributes and
// identity. It can be anonymous or linked to an identity that is managed by an identity provider. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-profiles" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) GetUserProfilesConfig(getUserProfilesConfigOptions *GetUserProfilesConfigOptions) (result *GetUserProfilesConfigResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.GetUserProfilesConfigWithContext(context.Background(), getUserProfilesConfigOptions)
}

// GetUserProfilesConfigWithContext is an alternate form of the GetUserProfilesConfig method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetUserProfilesConfigWithContext(ctx context.Context, getUserProfilesConfigOptions *GetUserProfilesConfigOptions) (result *GetUserProfilesConfigResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getUserProfilesConfigOptions, "getUserProfilesConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/users_profile`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getUserProfilesConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetUserProfilesConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetUserProfilesConfigResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateUserProfilesConfig : Update user profiles configuration
// A user profile is an entity that is stored and maintained by App ID. The profile holds a user's attributes and
// identity. It can be anonymous or linked to an identity that is managed by an identity provider. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-profiles" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) UpdateUserProfilesConfig(updateUserProfilesConfigOptions *UpdateUserProfilesConfigOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.UpdateUserProfilesConfigWithContext(context.Background(), updateUserProfilesConfigOptions)
}

// UpdateUserProfilesConfigWithContext is an alternate form of the UpdateUserProfilesConfig method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UpdateUserProfilesConfigWithContext(ctx context.Context, updateUserProfilesConfigOptions *UpdateUserProfilesConfigOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateUserProfilesConfigOptions, "updateUserProfilesConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateUserProfilesConfigOptions, "updateUserProfilesConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/users_profile`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateUserProfilesConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UpdateUserProfilesConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateUserProfilesConfigOptions.IsActive != nil {
		body["isActive"] = updateUserProfilesConfigOptions.IsActive
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// GetThemeText : Get widget texts
// Get the theme texts of the App ID login widget. <a href="https://cloud.ibm.com/docs/appid?topic=appid-login-widget"
// target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) GetThemeText(getThemeTextOptions *GetThemeTextOptions) (result *GetThemeTextResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.GetThemeTextWithContext(context.Background(), getThemeTextOptions)
}

// GetThemeTextWithContext is an alternate form of the GetThemeText method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetThemeTextWithContext(ctx context.Context, getThemeTextOptions *GetThemeTextOptions) (result *GetThemeTextResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getThemeTextOptions, "getThemeTextOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/ui/theme_text`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getThemeTextOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetThemeText")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetThemeTextResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PostThemeText : Update widget texts
// Update the texts of the App ID login widget. <a href="https://cloud.ibm.com/docs/appid?topic=appid-login-widget"
// target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) PostThemeText(postThemeTextOptions *PostThemeTextOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.PostThemeTextWithContext(context.Background(), postThemeTextOptions)
}

// PostThemeTextWithContext is an alternate form of the PostThemeText method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) PostThemeTextWithContext(ctx context.Context, postThemeTextOptions *PostThemeTextOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postThemeTextOptions, "postThemeTextOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(postThemeTextOptions, "postThemeTextOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/ui/theme_text`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range postThemeTextOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "PostThemeText")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if postThemeTextOptions.TabTitle != nil {
		body["tabTitle"] = postThemeTextOptions.TabTitle
	}
	if postThemeTextOptions.Footnote != nil {
		body["footnote"] = postThemeTextOptions.Footnote
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// GetThemeColor : Get widget colors
// Get the colors of the App ID login widget. <a href="https://cloud.ibm.com/docs/appid?topic=appid-login-widget"
// target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) GetThemeColor(getThemeColorOptions *GetThemeColorOptions) (result *GetThemeColorResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.GetThemeColorWithContext(context.Background(), getThemeColorOptions)
}

// GetThemeColorWithContext is an alternate form of the GetThemeColor method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetThemeColorWithContext(ctx context.Context, getThemeColorOptions *GetThemeColorOptions) (result *GetThemeColorResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getThemeColorOptions, "getThemeColorOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/ui/theme_color`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getThemeColorOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetThemeColor")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetThemeColorResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PostThemeColor : Update widget colors
// Update the colors of the App ID login widget. <a href="https://cloud.ibm.com/docs/appid?topic=appid-login-widget"
// target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) PostThemeColor(postThemeColorOptions *PostThemeColorOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.PostThemeColorWithContext(context.Background(), postThemeColorOptions)
}

// PostThemeColorWithContext is an alternate form of the PostThemeColor method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) PostThemeColorWithContext(ctx context.Context, postThemeColorOptions *PostThemeColorOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postThemeColorOptions, "postThemeColorOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(postThemeColorOptions, "postThemeColorOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/ui/theme_color`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range postThemeColorOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "PostThemeColor")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if postThemeColorOptions.HeaderColor != nil {
		body["headerColor"] = postThemeColorOptions.HeaderColor
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// GetMedia : Get widget logo
// Returns the link to the custom logo image of the login widget. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-login-widget" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) GetMedia(getMediaOptions *GetMediaOptions) (result *GetMediaResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.GetMediaWithContext(context.Background(), getMediaOptions)
}

// GetMediaWithContext is an alternate form of the GetMedia method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetMediaWithContext(ctx context.Context, getMediaOptions *GetMediaOptions) (result *GetMediaResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMediaOptions, "getMediaOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/ui/media`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getMediaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetMedia")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetMediaResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PostMedia : Update widget logo
// You can update the image file shown in the login widget. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-login-widget" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) PostMedia(postMediaOptions *PostMediaOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.PostMediaWithContext(context.Background(), postMediaOptions)
}

// PostMediaWithContext is an alternate form of the PostMedia method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) PostMediaWithContext(ctx context.Context, postMediaOptions *PostMediaOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postMediaOptions, "postMediaOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(postMediaOptions, "postMediaOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/ui/media`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range postMediaOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "PostMedia")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("mediaType", fmt.Sprint(*postMediaOptions.MediaType))

	builder.AddFormData("file", "filename",
		core.StringNilMapper(postMediaOptions.FileContentType), postMediaOptions.File)

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// GetSAMLMetadata : Get the SAML metadata
// Returns the SAML metadata required in order to integrate App ID with a SAML identity provider. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-enterprise" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) GetSAMLMetadata(getSAMLMetadataOptions *GetSAMLMetadataOptions) (result *string, response *core.DetailedResponse, err error) {
	return appIdManagement.GetSAMLMetadataWithContext(context.Background(), getSAMLMetadataOptions)
}

// GetSAMLMetadataWithContext is an alternate form of the GetSAMLMetadata method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetSAMLMetadataWithContext(ctx context.Context, getSAMLMetadataOptions *GetSAMLMetadataOptions) (result *string, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSAMLMetadataOptions, "getSAMLMetadataOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/saml_metadata`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSAMLMetadataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetSAMLMetadata")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/xml")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, &result)

	return
}

// GetTemplate : Get an email template
// Returns the content of a custom email template or the default template in case it wasn't customized. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cd-types" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) GetTemplate(getTemplateOptions *GetTemplateOptions) (result *GetTemplate, response *core.DetailedResponse, err error) {
	return appIdManagement.GetTemplateWithContext(context.Background(), getTemplateOptions)
}

// GetTemplateWithContext is an alternate form of the GetTemplate method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetTemplateWithContext(ctx context.Context, getTemplateOptions *GetTemplateOptions) (result *GetTemplate, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getTemplateOptions, "getTemplateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getTemplateOptions, "getTemplateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"templateName": *getTemplateOptions.TemplateName,
		"language": *getTemplateOptions.Language,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/templates/{templateName}/{language}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetTemplate)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateTemplate : Update an email template
// Updates the Cloud Directory email template. <a href="https://cloud.ibm.com/docs/appid?topic=appid-cd-types"
// target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) UpdateTemplate(updateTemplateOptions *UpdateTemplateOptions) (result *GetTemplate, response *core.DetailedResponse, err error) {
	return appIdManagement.UpdateTemplateWithContext(context.Background(), updateTemplateOptions)
}

// UpdateTemplateWithContext is an alternate form of the UpdateTemplate method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UpdateTemplateWithContext(ctx context.Context, updateTemplateOptions *UpdateTemplateOptions) (result *GetTemplate, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateTemplateOptions, "updateTemplateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateTemplateOptions, "updateTemplateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"templateName": *updateTemplateOptions.TemplateName,
		"language": *updateTemplateOptions.Language,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/templates/{templateName}/{language}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UpdateTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateTemplateOptions.Subject != nil {
		body["subject"] = updateTemplateOptions.Subject
	}
	if updateTemplateOptions.HTMLBody != nil {
		body["html_body"] = updateTemplateOptions.HTMLBody
	}
	if updateTemplateOptions.Base64EncodedHTMLBody != nil {
		body["base64_encoded_html_body"] = updateTemplateOptions.Base64EncodedHTMLBody
	}
	if updateTemplateOptions.PlainTextBody != nil {
		body["plain_text_body"] = updateTemplateOptions.PlainTextBody
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetTemplate)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteTemplate : Delete an email template
// Delete the customized email template and reverts to App ID default template. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cd-users" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) DeleteTemplate(deleteTemplateOptions *DeleteTemplateOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.DeleteTemplateWithContext(context.Background(), deleteTemplateOptions)
}

// DeleteTemplateWithContext is an alternate form of the DeleteTemplate method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) DeleteTemplateWithContext(ctx context.Context, deleteTemplateOptions *DeleteTemplateOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteTemplateOptions, "deleteTemplateOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteTemplateOptions, "deleteTemplateOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"templateName": *deleteTemplateOptions.TemplateName,
		"language": *deleteTemplateOptions.Language,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/templates/{templateName}/{language}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteTemplateOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "DeleteTemplate")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// GetLocalization : Get languages
// Returns the list of languages that can be used to customize email templates for Cloud Directory.
func (appIdManagement *AppIdManagementV4) GetLocalization(getLocalizationOptions *GetLocalizationOptions) (result *GetLanguages, response *core.DetailedResponse, err error) {
	return appIdManagement.GetLocalizationWithContext(context.Background(), getLocalizationOptions)
}

// GetLocalizationWithContext is an alternate form of the GetLocalization method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetLocalizationWithContext(ctx context.Context, getLocalizationOptions *GetLocalizationOptions) (result *GetLanguages, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getLocalizationOptions, "getLocalizationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/ui/languages`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getLocalizationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetLocalization")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetLanguages)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateLocalization : Update languages
// Update the list of languages that can be used to customize email templates for Cloud Directory.
func (appIdManagement *AppIdManagementV4) UpdateLocalization(updateLocalizationOptions *UpdateLocalizationOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.UpdateLocalizationWithContext(context.Background(), updateLocalizationOptions)
}

// UpdateLocalizationWithContext is an alternate form of the UpdateLocalization method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UpdateLocalizationWithContext(ctx context.Context, updateLocalizationOptions *UpdateLocalizationOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(updateLocalizationOptions, "updateLocalizationOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/ui/languages`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateLocalizationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UpdateLocalization")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateLocalizationOptions.Languages != nil {
		body["languages"] = updateLocalizationOptions.Languages
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// GetCloudDirectorySenderDetails : Get sender details
// Returns the sender details configuration that is used by Cloud Directory when sending emails. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cd-types" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) GetCloudDirectorySenderDetails(getCloudDirectorySenderDetailsOptions *GetCloudDirectorySenderDetailsOptions) (result *CloudDirectorySenderDetails, response *core.DetailedResponse, err error) {
	return appIdManagement.GetCloudDirectorySenderDetailsWithContext(context.Background(), getCloudDirectorySenderDetailsOptions)
}

// GetCloudDirectorySenderDetailsWithContext is an alternate form of the GetCloudDirectorySenderDetails method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetCloudDirectorySenderDetailsWithContext(ctx context.Context, getCloudDirectorySenderDetailsOptions *GetCloudDirectorySenderDetailsOptions) (result *CloudDirectorySenderDetails, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCloudDirectorySenderDetailsOptions, "getCloudDirectorySenderDetailsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/sender_details`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCloudDirectorySenderDetailsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetCloudDirectorySenderDetails")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCloudDirectorySenderDetails)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetCloudDirectorySenderDetails : Update the sender details
// Updates the sender details configuration that is used by Cloud Directory when sending emails. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cd-types" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) SetCloudDirectorySenderDetails(setCloudDirectorySenderDetailsOptions *SetCloudDirectorySenderDetailsOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.SetCloudDirectorySenderDetailsWithContext(context.Background(), setCloudDirectorySenderDetailsOptions)
}

// SetCloudDirectorySenderDetailsWithContext is an alternate form of the SetCloudDirectorySenderDetails method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) SetCloudDirectorySenderDetailsWithContext(ctx context.Context, setCloudDirectorySenderDetailsOptions *SetCloudDirectorySenderDetailsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setCloudDirectorySenderDetailsOptions, "setCloudDirectorySenderDetailsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setCloudDirectorySenderDetailsOptions, "setCloudDirectorySenderDetailsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/sender_details`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setCloudDirectorySenderDetailsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "SetCloudDirectorySenderDetails")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setCloudDirectorySenderDetailsOptions.SenderDetails != nil {
		body["senderDetails"] = setCloudDirectorySenderDetailsOptions.SenderDetails
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// GetCloudDirectoryActionURL : Get action url
// Get the custom url to redirect to when <b>action</b> is executed. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cloud-directory" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) GetCloudDirectoryActionURL(getCloudDirectoryActionUrlOptions *GetCloudDirectoryActionUrlOptions) (result *ActionUrlResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.GetCloudDirectoryActionURLWithContext(context.Background(), getCloudDirectoryActionUrlOptions)
}

// GetCloudDirectoryActionURLWithContext is an alternate form of the GetCloudDirectoryActionURL method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetCloudDirectoryActionURLWithContext(ctx context.Context, getCloudDirectoryActionUrlOptions *GetCloudDirectoryActionUrlOptions) (result *ActionUrlResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCloudDirectoryActionUrlOptions, "getCloudDirectoryActionUrlOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCloudDirectoryActionUrlOptions, "getCloudDirectoryActionUrlOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"action": *getCloudDirectoryActionUrlOptions.Action,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/action_url/{action}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCloudDirectoryActionUrlOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetCloudDirectoryActionURL")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalActionUrlResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetCloudDirectoryAction : Update action url
// Updates the custom url to redirect to when <b>action</b> is executed. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cloud-directory" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) SetCloudDirectoryAction(setCloudDirectoryActionOptions *SetCloudDirectoryActionOptions) (result *ActionUrlResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.SetCloudDirectoryActionWithContext(context.Background(), setCloudDirectoryActionOptions)
}

// SetCloudDirectoryActionWithContext is an alternate form of the SetCloudDirectoryAction method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) SetCloudDirectoryActionWithContext(ctx context.Context, setCloudDirectoryActionOptions *SetCloudDirectoryActionOptions) (result *ActionUrlResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setCloudDirectoryActionOptions, "setCloudDirectoryActionOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setCloudDirectoryActionOptions, "setCloudDirectoryActionOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"action": *setCloudDirectoryActionOptions.Action,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/action_url/{action}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setCloudDirectoryActionOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "SetCloudDirectoryAction")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setCloudDirectoryActionOptions.ActionURL != nil {
		body["actionUrl"] = setCloudDirectoryActionOptions.ActionURL
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalActionUrlResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteActionURL : Delete action url
// Delete the custom url to redirect to when <b>action</b> is executed. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cloud-directory" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) DeleteActionURL(deleteActionUrlOptions *DeleteActionUrlOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.DeleteActionURLWithContext(context.Background(), deleteActionUrlOptions)
}

// DeleteActionURLWithContext is an alternate form of the DeleteActionURL method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) DeleteActionURLWithContext(ctx context.Context, deleteActionUrlOptions *DeleteActionUrlOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteActionUrlOptions, "deleteActionUrlOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteActionUrlOptions, "deleteActionUrlOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"action": *deleteActionUrlOptions.Action,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/action_url/{action}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteActionUrlOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "DeleteActionURL")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// GetCloudDirectoryPasswordRegex : Get password regex
// Returns the regular expression used by App ID for password strength validation. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cd-strength" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) GetCloudDirectoryPasswordRegex(getCloudDirectoryPasswordRegexOptions *GetCloudDirectoryPasswordRegexOptions) (result *PasswordRegexConfigParamsGet, response *core.DetailedResponse, err error) {
	return appIdManagement.GetCloudDirectoryPasswordRegexWithContext(context.Background(), getCloudDirectoryPasswordRegexOptions)
}

// GetCloudDirectoryPasswordRegexWithContext is an alternate form of the GetCloudDirectoryPasswordRegex method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetCloudDirectoryPasswordRegexWithContext(ctx context.Context, getCloudDirectoryPasswordRegexOptions *GetCloudDirectoryPasswordRegexOptions) (result *PasswordRegexConfigParamsGet, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCloudDirectoryPasswordRegexOptions, "getCloudDirectoryPasswordRegexOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/password_regex`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCloudDirectoryPasswordRegexOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetCloudDirectoryPasswordRegex")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPasswordRegexConfigParamsGet)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetCloudDirectoryPasswordRegex : Update password regex
// Updates the regular expression used by App ID for password strength validation.<br />For example, the regular
// expression: <code>"^[A-Za-z\d]*$"</code> should be passed as:<br /><code>{<br />&nbsp;&nbsp;"base64_encoded_regex":
// "XltBLVphLXpcZF0qJA==", <br />&nbsp;&nbsp;"error_message": "Must only contain letters and digits"<br />}</code> <br
// /><br /> <a href="https://cloud.ibm.com/docs/appid?topic=appid-cd-strength" target="_blank" rel="noopener">Learn
// more</a>.
func (appIdManagement *AppIdManagementV4) SetCloudDirectoryPasswordRegex(setCloudDirectoryPasswordRegexOptions *SetCloudDirectoryPasswordRegexOptions) (result *PasswordRegexConfigParamsGet, response *core.DetailedResponse, err error) {
	return appIdManagement.SetCloudDirectoryPasswordRegexWithContext(context.Background(), setCloudDirectoryPasswordRegexOptions)
}

// SetCloudDirectoryPasswordRegexWithContext is an alternate form of the SetCloudDirectoryPasswordRegex method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) SetCloudDirectoryPasswordRegexWithContext(ctx context.Context, setCloudDirectoryPasswordRegexOptions *SetCloudDirectoryPasswordRegexOptions) (result *PasswordRegexConfigParamsGet, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setCloudDirectoryPasswordRegexOptions, "setCloudDirectoryPasswordRegexOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setCloudDirectoryPasswordRegexOptions, "setCloudDirectoryPasswordRegexOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/password_regex`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setCloudDirectoryPasswordRegexOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "SetCloudDirectoryPasswordRegex")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setCloudDirectoryPasswordRegexOptions.Regex != nil {
		body["regex"] = setCloudDirectoryPasswordRegexOptions.Regex
	}
	if setCloudDirectoryPasswordRegexOptions.Base64EncodedRegex != nil {
		body["base64_encoded_regex"] = setCloudDirectoryPasswordRegexOptions.Base64EncodedRegex
	}
	if setCloudDirectoryPasswordRegexOptions.ErrorMessage != nil {
		body["error_message"] = setCloudDirectoryPasswordRegexOptions.ErrorMessage
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPasswordRegexConfigParamsGet)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCloudDirectoryEmailDispatcher : Get email dispatcher configuration
// Get the configuration of email dispatcher that is used by Cloud Directory when sending emails.
func (appIdManagement *AppIdManagementV4) GetCloudDirectoryEmailDispatcher(getCloudDirectoryEmailDispatcherOptions *GetCloudDirectoryEmailDispatcherOptions) (result *EmailDispatcherParams, response *core.DetailedResponse, err error) {
	return appIdManagement.GetCloudDirectoryEmailDispatcherWithContext(context.Background(), getCloudDirectoryEmailDispatcherOptions)
}

// GetCloudDirectoryEmailDispatcherWithContext is an alternate form of the GetCloudDirectoryEmailDispatcher method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetCloudDirectoryEmailDispatcherWithContext(ctx context.Context, getCloudDirectoryEmailDispatcherOptions *GetCloudDirectoryEmailDispatcherOptions) (result *EmailDispatcherParams, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCloudDirectoryEmailDispatcherOptions, "getCloudDirectoryEmailDispatcherOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/email_dispatcher`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCloudDirectoryEmailDispatcherOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetCloudDirectoryEmailDispatcher")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEmailDispatcherParams)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetCloudDirectoryEmailDispatcher : Update email dispatcher configuration
// App ID allows you to use your own email provider. You can use your own Sendgrid account by providing your Sendgrind
// API key. Alternatively, you can define a custom email dispatcher by providing App ID with URL. The URL is called for
// sending emails. Optionally, you can determine a specific authorization method – either basic, such as a username and
// password, or a custom value. By default, App ID's email provider will be used.
func (appIdManagement *AppIdManagementV4) SetCloudDirectoryEmailDispatcher(setCloudDirectoryEmailDispatcherOptions *SetCloudDirectoryEmailDispatcherOptions) (result *EmailDispatcherParams, response *core.DetailedResponse, err error) {
	return appIdManagement.SetCloudDirectoryEmailDispatcherWithContext(context.Background(), setCloudDirectoryEmailDispatcherOptions)
}

// SetCloudDirectoryEmailDispatcherWithContext is an alternate form of the SetCloudDirectoryEmailDispatcher method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) SetCloudDirectoryEmailDispatcherWithContext(ctx context.Context, setCloudDirectoryEmailDispatcherOptions *SetCloudDirectoryEmailDispatcherOptions) (result *EmailDispatcherParams, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setCloudDirectoryEmailDispatcherOptions, "setCloudDirectoryEmailDispatcherOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setCloudDirectoryEmailDispatcherOptions, "setCloudDirectoryEmailDispatcherOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/email_dispatcher`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setCloudDirectoryEmailDispatcherOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "SetCloudDirectoryEmailDispatcher")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setCloudDirectoryEmailDispatcherOptions.Provider != nil {
		body["provider"] = setCloudDirectoryEmailDispatcherOptions.Provider
	}
	if setCloudDirectoryEmailDispatcherOptions.Sendgrid != nil {
		body["sendgrid"] = setCloudDirectoryEmailDispatcherOptions.Sendgrid
	}
	if setCloudDirectoryEmailDispatcherOptions.Custom != nil {
		body["custom"] = setCloudDirectoryEmailDispatcherOptions.Custom
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalEmailDispatcherParams)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// EmailSettingTest : Test the email provider configuration
// You can send a message to a specific email to test your settings.
func (appIdManagement *AppIdManagementV4) EmailSettingTest(emailSettingTestOptions *EmailSettingTestOptions) (result *RespEmailSettingsTest, response *core.DetailedResponse, err error) {
	return appIdManagement.EmailSettingTestWithContext(context.Background(), emailSettingTestOptions)
}

// EmailSettingTestWithContext is an alternate form of the EmailSettingTest method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) EmailSettingTestWithContext(ctx context.Context, emailSettingTestOptions *EmailSettingTestOptions) (result *RespEmailSettingsTest, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(emailSettingTestOptions, "emailSettingTestOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(emailSettingTestOptions, "emailSettingTestOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/email_settings/test`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range emailSettingTestOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "EmailSettingTest")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if emailSettingTestOptions.EmailTo != nil {
		body["emailTo"] = emailSettingTestOptions.EmailTo
	}
	if emailSettingTestOptions.EmailSettings != nil {
		body["emailSettings"] = emailSettingTestOptions.EmailSettings
	}
	if emailSettingTestOptions.SenderDetails != nil {
		body["senderDetails"] = emailSettingTestOptions.SenderDetails
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRespEmailSettingsTest)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PostEmailDispatcherTest : Test the email dispatcher configuration
// You can send a message to a specific email to test your configuration.
func (appIdManagement *AppIdManagementV4) PostEmailDispatcherTest(postEmailDispatcherTestOptions *PostEmailDispatcherTestOptions) (result *RespCustomEmailDisParams, response *core.DetailedResponse, err error) {
	return appIdManagement.PostEmailDispatcherTestWithContext(context.Background(), postEmailDispatcherTestOptions)
}

// PostEmailDispatcherTestWithContext is an alternate form of the PostEmailDispatcherTest method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) PostEmailDispatcherTestWithContext(ctx context.Context, postEmailDispatcherTestOptions *PostEmailDispatcherTestOptions) (result *RespCustomEmailDisParams, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postEmailDispatcherTestOptions, "postEmailDispatcherTestOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(postEmailDispatcherTestOptions, "postEmailDispatcherTestOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/email_dispatcher/test`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range postEmailDispatcherTestOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "PostEmailDispatcherTest")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if postEmailDispatcherTestOptions.Email != nil {
		body["email"] = postEmailDispatcherTestOptions.Email
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRespCustomEmailDisParams)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PostSmsDispatcherTest : Test the MFA SMS dispatcher configuration
// You can send a message to a specific phone number to test your MFA SMS configuration.
func (appIdManagement *AppIdManagementV4) PostSmsDispatcherTest(postSmsDispatcherTestOptions *PostSmsDispatcherTestOptions) (result *RespSMSDisParams, response *core.DetailedResponse, err error) {
	return appIdManagement.PostSmsDispatcherTestWithContext(context.Background(), postSmsDispatcherTestOptions)
}

// PostSmsDispatcherTestWithContext is an alternate form of the PostSmsDispatcherTest method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) PostSmsDispatcherTestWithContext(ctx context.Context, postSmsDispatcherTestOptions *PostSmsDispatcherTestOptions) (result *RespSMSDisParams, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postSmsDispatcherTestOptions, "postSmsDispatcherTestOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(postSmsDispatcherTestOptions, "postSmsDispatcherTestOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/sms_dispatcher/test`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range postSmsDispatcherTestOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "PostSmsDispatcherTest")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if postSmsDispatcherTestOptions.PhoneNumber != nil {
		body["phone_number"] = postSmsDispatcherTestOptions.PhoneNumber
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRespSMSDisParams)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCloudDirectoryAdvancedPasswordManagement : Get APM configuration
// Get the configuration of the advanced password management.
func (appIdManagement *AppIdManagementV4) GetCloudDirectoryAdvancedPasswordManagement(getCloudDirectoryAdvancedPasswordManagementOptions *GetCloudDirectoryAdvancedPasswordManagementOptions) (result *ApmSchema, response *core.DetailedResponse, err error) {
	return appIdManagement.GetCloudDirectoryAdvancedPasswordManagementWithContext(context.Background(), getCloudDirectoryAdvancedPasswordManagementOptions)
}

// GetCloudDirectoryAdvancedPasswordManagementWithContext is an alternate form of the GetCloudDirectoryAdvancedPasswordManagement method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetCloudDirectoryAdvancedPasswordManagementWithContext(ctx context.Context, getCloudDirectoryAdvancedPasswordManagementOptions *GetCloudDirectoryAdvancedPasswordManagementOptions) (result *ApmSchema, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCloudDirectoryAdvancedPasswordManagementOptions, "getCloudDirectoryAdvancedPasswordManagementOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/advanced_password_management`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCloudDirectoryAdvancedPasswordManagementOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetCloudDirectoryAdvancedPasswordManagement")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApmSchema)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetCloudDirectoryAdvancedPasswordManagement : Update APM configuration
// Updates the advanced password management configuration for the provided tenantId. By turning this on, any
// authentication event is also charged as advanced security event.
func (appIdManagement *AppIdManagementV4) SetCloudDirectoryAdvancedPasswordManagement(setCloudDirectoryAdvancedPasswordManagementOptions *SetCloudDirectoryAdvancedPasswordManagementOptions) (result *ApmSchema, response *core.DetailedResponse, err error) {
	return appIdManagement.SetCloudDirectoryAdvancedPasswordManagementWithContext(context.Background(), setCloudDirectoryAdvancedPasswordManagementOptions)
}

// SetCloudDirectoryAdvancedPasswordManagementWithContext is an alternate form of the SetCloudDirectoryAdvancedPasswordManagement method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) SetCloudDirectoryAdvancedPasswordManagementWithContext(ctx context.Context, setCloudDirectoryAdvancedPasswordManagementOptions *SetCloudDirectoryAdvancedPasswordManagementOptions) (result *ApmSchema, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setCloudDirectoryAdvancedPasswordManagementOptions, "setCloudDirectoryAdvancedPasswordManagementOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setCloudDirectoryAdvancedPasswordManagementOptions, "setCloudDirectoryAdvancedPasswordManagementOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/advanced_password_management`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setCloudDirectoryAdvancedPasswordManagementOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "SetCloudDirectoryAdvancedPasswordManagement")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setCloudDirectoryAdvancedPasswordManagementOptions.AdvancedPasswordManagement != nil {
		body["advancedPasswordManagement"] = setCloudDirectoryAdvancedPasswordManagementOptions.AdvancedPasswordManagement
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalApmSchema)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetAuditStatus : Get tenant audit status
// Returns a JSON object containing the auditing status of the tenant.
func (appIdManagement *AppIdManagementV4) GetAuditStatus(getAuditStatusOptions *GetAuditStatusOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.GetAuditStatusWithContext(context.Background(), getAuditStatusOptions)
}

// GetAuditStatusWithContext is an alternate form of the GetAuditStatus method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetAuditStatusWithContext(ctx context.Context, getAuditStatusOptions *GetAuditStatusOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getAuditStatusOptions, "getAuditStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/capture_runtime_activity`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getAuditStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetAuditStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// SetAuditStatus : Update tenant audit status
// Capture app user sign-in, sign-up and other runtime events in Activity Tracker for you to search, analyse and report.
// By turning this On, any authentication event is also charged as advanced security event. Activity Tracker with LogDNA
// is available in select regions. <a href="https://cloud.ibm.com/docs/appid?topic=appid-at-events">Learn more</a>.
func (appIdManagement *AppIdManagementV4) SetAuditStatus(setAuditStatusOptions *SetAuditStatusOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.SetAuditStatusWithContext(context.Background(), setAuditStatusOptions)
}

// SetAuditStatusWithContext is an alternate form of the SetAuditStatus method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) SetAuditStatusWithContext(ctx context.Context, setAuditStatusOptions *SetAuditStatusOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setAuditStatusOptions, "setAuditStatusOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setAuditStatusOptions, "setAuditStatusOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/capture_runtime_activity`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setAuditStatusOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "SetAuditStatus")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setAuditStatusOptions.IsActive != nil {
		body["isActive"] = setAuditStatusOptions.IsActive
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// ListChannels : List channels
// Returns all MFA channels registered with the App ID Instance.
func (appIdManagement *AppIdManagementV4) ListChannels(listChannelsOptions *ListChannelsOptions) (result *MfaChannelsList, response *core.DetailedResponse, err error) {
	return appIdManagement.ListChannelsWithContext(context.Background(), listChannelsOptions)
}

// ListChannelsWithContext is an alternate form of the ListChannels method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) ListChannelsWithContext(ctx context.Context, listChannelsOptions *ListChannelsOptions) (result *MfaChannelsList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listChannelsOptions, "listChannelsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/mfa/channels`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listChannelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "ListChannels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalMfaChannelsList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetChannel : Get channel
// Returns a specific MFA channel registered with the App ID Instance.
func (appIdManagement *AppIdManagementV4) GetChannel(getChannelOptions *GetChannelOptions) (result *GetSMSChannel, response *core.DetailedResponse, err error) {
	return appIdManagement.GetChannelWithContext(context.Background(), getChannelOptions)
}

// GetChannelWithContext is an alternate form of the GetChannel method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetChannelWithContext(ctx context.Context, getChannelOptions *GetChannelOptions) (result *GetSMSChannel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getChannelOptions, "getChannelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getChannelOptions, "getChannelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"channel": *getChannelOptions.Channel,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/mfa/channels/{channel}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getChannelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetChannel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetSMSChannel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateChannel : Update channel
// Enable or disable a registered MFA channel on the App ID instance.
func (appIdManagement *AppIdManagementV4) UpdateChannel(updateChannelOptions *UpdateChannelOptions) (result *GetSMSChannel, response *core.DetailedResponse, err error) {
	return appIdManagement.UpdateChannelWithContext(context.Background(), updateChannelOptions)
}

// UpdateChannelWithContext is an alternate form of the UpdateChannel method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UpdateChannelWithContext(ctx context.Context, updateChannelOptions *UpdateChannelOptions) (result *GetSMSChannel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateChannelOptions, "updateChannelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateChannelOptions, "updateChannelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"channel": *updateChannelOptions.Channel,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/mfa/channels/{channel}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateChannelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UpdateChannel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateChannelOptions.IsActive != nil {
		body["isActive"] = updateChannelOptions.IsActive
	}
	if updateChannelOptions.Config != nil {
		body["config"] = updateChannelOptions.Config
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetSMSChannel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetExtensionConfig : Get an extension configuration
// View a registered extension's configuration for an instance of App ID. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cd-mfa#cd-mfa-extensions" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) GetExtensionConfig(getExtensionConfigOptions *GetExtensionConfigOptions) (result *UpdateExtensionConfig, response *core.DetailedResponse, err error) {
	return appIdManagement.GetExtensionConfigWithContext(context.Background(), getExtensionConfigOptions)
}

// GetExtensionConfigWithContext is an alternate form of the GetExtensionConfig method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetExtensionConfigWithContext(ctx context.Context, getExtensionConfigOptions *GetExtensionConfigOptions) (result *UpdateExtensionConfig, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getExtensionConfigOptions, "getExtensionConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getExtensionConfigOptions, "getExtensionConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"name": *getExtensionConfigOptions.Name,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/mfa/extensions/{name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getExtensionConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetExtensionConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUpdateExtensionConfig)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateExtensionConfig : Update an extension configuration
// Set or update a registered extension's configuration for an instance of App ID. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cd-mfa#cd-mfa-extensions" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) UpdateExtensionConfig(updateExtensionConfigOptions *UpdateExtensionConfigOptions) (result *UpdateExtensionConfig, response *core.DetailedResponse, err error) {
	return appIdManagement.UpdateExtensionConfigWithContext(context.Background(), updateExtensionConfigOptions)
}

// UpdateExtensionConfigWithContext is an alternate form of the UpdateExtensionConfig method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UpdateExtensionConfigWithContext(ctx context.Context, updateExtensionConfigOptions *UpdateExtensionConfigOptions) (result *UpdateExtensionConfig, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateExtensionConfigOptions, "updateExtensionConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateExtensionConfigOptions, "updateExtensionConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"name": *updateExtensionConfigOptions.Name,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/mfa/extensions/{name}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateExtensionConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UpdateExtensionConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateExtensionConfigOptions.IsActive != nil {
		body["isActive"] = updateExtensionConfigOptions.IsActive
	}
	if updateExtensionConfigOptions.Config != nil {
		body["config"] = updateExtensionConfigOptions.Config
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUpdateExtensionConfig)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateExtensionActive : Enable or disable an extension
// Update the status of a registered extension for an instance of App ID to enabled or disabled. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cd-mfa#cd-mfa-extensions" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) UpdateExtensionActive(updateExtensionActiveOptions *UpdateExtensionActiveOptions) (result *ExtensionActive, response *core.DetailedResponse, err error) {
	return appIdManagement.UpdateExtensionActiveWithContext(context.Background(), updateExtensionActiveOptions)
}

// UpdateExtensionActiveWithContext is an alternate form of the UpdateExtensionActive method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UpdateExtensionActiveWithContext(ctx context.Context, updateExtensionActiveOptions *UpdateExtensionActiveOptions) (result *ExtensionActive, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateExtensionActiveOptions, "updateExtensionActiveOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateExtensionActiveOptions, "updateExtensionActiveOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"name": *updateExtensionActiveOptions.Name,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/mfa/extensions/{name}/active`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateExtensionActiveOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UpdateExtensionActive")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateExtensionActiveOptions.IsActive != nil {
		body["isActive"] = updateExtensionActiveOptions.IsActive
	}
	if updateExtensionActiveOptions.Config != nil {
		body["config"] = updateExtensionActiveOptions.Config
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExtensionActive)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// PostExtensionsTest : Test the extension configuration
// Test an extension configuration. <a href="https://cloud.ibm.com/docs/appid?topic=appid-cd-mfa#cd-mfa-extensions"
// target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) PostExtensionsTest(postExtensionsTestOptions *PostExtensionsTestOptions) (result *ExtensionTest, response *core.DetailedResponse, err error) {
	return appIdManagement.PostExtensionsTestWithContext(context.Background(), postExtensionsTestOptions)
}

// PostExtensionsTestWithContext is an alternate form of the PostExtensionsTest method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) PostExtensionsTestWithContext(ctx context.Context, postExtensionsTestOptions *PostExtensionsTestOptions) (result *ExtensionTest, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(postExtensionsTestOptions, "postExtensionsTestOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(postExtensionsTestOptions, "postExtensionsTestOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"name": *postExtensionsTestOptions.Name,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/mfa/extensions/{name}/test`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range postExtensionsTestOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "PostExtensionsTest")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExtensionTest)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetMFAConfig : Get MFA configuration
// Returns MFA configuration registered with the App ID Instance.
func (appIdManagement *AppIdManagementV4) GetMFAConfig(getMFAConfigOptions *GetMFAConfigOptions) (result *GetMFAConfiguration, response *core.DetailedResponse, err error) {
	return appIdManagement.GetMFAConfigWithContext(context.Background(), getMFAConfigOptions)
}

// GetMFAConfigWithContext is an alternate form of the GetMFAConfig method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetMFAConfigWithContext(ctx context.Context, getMFAConfigOptions *GetMFAConfigOptions) (result *GetMFAConfiguration, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getMFAConfigOptions, "getMFAConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/mfa`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getMFAConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetMFAConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetMFAConfiguration)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateMFAConfig : Update MFA configuration
// Update MFA configuration on the App ID instance.
func (appIdManagement *AppIdManagementV4) UpdateMFAConfig(updateMFAConfigOptions *UpdateMFAConfigOptions) (result *GetMFAConfiguration, response *core.DetailedResponse, err error) {
	return appIdManagement.UpdateMFAConfigWithContext(context.Background(), updateMFAConfigOptions)
}

// UpdateMFAConfigWithContext is an alternate form of the UpdateMFAConfig method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UpdateMFAConfigWithContext(ctx context.Context, updateMFAConfigOptions *UpdateMFAConfigOptions) (result *GetMFAConfiguration, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateMFAConfigOptions, "updateMFAConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateMFAConfigOptions, "updateMFAConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/mfa`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateMFAConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UpdateMFAConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateMFAConfigOptions.IsActive != nil {
		body["isActive"] = updateMFAConfigOptions.IsActive
	}
	if updateMFAConfigOptions.Config != nil {
		body["config"] = updateMFAConfigOptions.Config
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetMFAConfiguration)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetSSOConfig : Get SSO configuration
// Returns SSO configuration registered with the App ID Instance.
func (appIdManagement *AppIdManagementV4) GetSSOConfig(getSSOConfigOptions *GetSSOConfigOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.GetSSOConfigWithContext(context.Background(), getSSOConfigOptions)
}

// GetSSOConfigWithContext is an alternate form of the GetSSOConfig method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetSSOConfigWithContext(ctx context.Context, getSSOConfigOptions *GetSSOConfigOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSSOConfigOptions, "getSSOConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/sso`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSSOConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetSSOConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// UpdateSSOConfig : Update SSO configuration
// Update SSO configuration on the App ID instance.
func (appIdManagement *AppIdManagementV4) UpdateSSOConfig(updateSSOConfigOptions *UpdateSSOConfigOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.UpdateSSOConfigWithContext(context.Background(), updateSSOConfigOptions)
}

// UpdateSSOConfigWithContext is an alternate form of the UpdateSSOConfig method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UpdateSSOConfigWithContext(ctx context.Context, updateSSOConfigOptions *UpdateSSOConfigOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateSSOConfigOptions, "updateSSOConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateSSOConfigOptions, "updateSSOConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/sso`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateSSOConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UpdateSSOConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateSSOConfigOptions.IsActive != nil {
		body["isActive"] = updateSSOConfigOptions.IsActive
	}
	if updateSSOConfigOptions.InactivityTimeoutSeconds != nil {
		body["inactivityTimeoutSeconds"] = updateSSOConfigOptions.InactivityTimeoutSeconds
	}
	if updateSSOConfigOptions.LogoutRedirectUris != nil {
		body["logoutRedirectUris"] = updateSSOConfigOptions.LogoutRedirectUris
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// GetRateLimitConfig : Get the rate limit configuration
// Returns the rate limit configuration registered with the App ID Instance.
func (appIdManagement *AppIdManagementV4) GetRateLimitConfig(getRateLimitConfigOptions *GetRateLimitConfigOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.GetRateLimitConfigWithContext(context.Background(), getRateLimitConfigOptions)
}

// GetRateLimitConfigWithContext is an alternate form of the GetRateLimitConfig method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetRateLimitConfigWithContext(ctx context.Context, getRateLimitConfigOptions *GetRateLimitConfigOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getRateLimitConfigOptions, "getRateLimitConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/rate_limit`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRateLimitConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetRateLimitConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// UpdateRateLimitConfig : Update the rate limit configuration
// Update the rate limit configuration on the App ID instance.
func (appIdManagement *AppIdManagementV4) UpdateRateLimitConfig(updateRateLimitConfigOptions *UpdateRateLimitConfigOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.UpdateRateLimitConfigWithContext(context.Background(), updateRateLimitConfigOptions)
}

// UpdateRateLimitConfigWithContext is an alternate form of the UpdateRateLimitConfig method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UpdateRateLimitConfigWithContext(ctx context.Context, updateRateLimitConfigOptions *UpdateRateLimitConfigOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateRateLimitConfigOptions, "updateRateLimitConfigOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateRateLimitConfigOptions, "updateRateLimitConfigOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/cloud_directory/rate_limit`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateRateLimitConfigOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UpdateRateLimitConfig")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateRateLimitConfigOptions.SignUpLimitPerMinute != nil {
		body["signUpLimitPerMinute"] = updateRateLimitConfigOptions.SignUpLimitPerMinute
	}
	if updateRateLimitConfigOptions.SignInLimitPerMinute != nil {
		body["signInLimitPerMinute"] = updateRateLimitConfigOptions.SignInLimitPerMinute
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// GetFacebookIDP : Get Facebook IDP configuration
// Returns the Facebook identity provider configuration.
func (appIdManagement *AppIdManagementV4) GetFacebookIDP(getFacebookIDPOptions *GetFacebookIDPOptions) (result *FacebookConfigParams, response *core.DetailedResponse, err error) {
	return appIdManagement.GetFacebookIDPWithContext(context.Background(), getFacebookIDPOptions)
}

// GetFacebookIDPWithContext is an alternate form of the GetFacebookIDP method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetFacebookIDPWithContext(ctx context.Context, getFacebookIDPOptions *GetFacebookIDPOptions) (result *FacebookConfigParams, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getFacebookIDPOptions, "getFacebookIDPOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/idps/facebook`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getFacebookIDPOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetFacebookIDP")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalFacebookConfigParams)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetFacebookIDP : Update Facebook IDP configuration
// Configure Facebook to set up a single sign-on experience for your users. By using Facebook, users are able to sign in
// with credentials with which they are already familiar. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-social#facebook" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) SetFacebookIDP(setFacebookIDPOptions *SetFacebookIDPOptions) (result *FacebookConfigParamsPUT, response *core.DetailedResponse, err error) {
	return appIdManagement.SetFacebookIDPWithContext(context.Background(), setFacebookIDPOptions)
}

// SetFacebookIDPWithContext is an alternate form of the SetFacebookIDP method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) SetFacebookIDPWithContext(ctx context.Context, setFacebookIDPOptions *SetFacebookIDPOptions) (result *FacebookConfigParamsPUT, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setFacebookIDPOptions, "setFacebookIDPOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setFacebookIDPOptions, "setFacebookIDPOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/idps/facebook`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setFacebookIDPOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "SetFacebookIDP")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(setFacebookIDPOptions.IDP)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalFacebookConfigParamsPUT)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetGoogleIDP : Get Google IDP configuration
// Returns the Google identity provider configuration.
func (appIdManagement *AppIdManagementV4) GetGoogleIDP(getGoogleIDPOptions *GetGoogleIDPOptions) (result *GoogleConfigParams, response *core.DetailedResponse, err error) {
	return appIdManagement.GetGoogleIDPWithContext(context.Background(), getGoogleIDPOptions)
}

// GetGoogleIDPWithContext is an alternate form of the GetGoogleIDP method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetGoogleIDPWithContext(ctx context.Context, getGoogleIDPOptions *GetGoogleIDPOptions) (result *GoogleConfigParams, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getGoogleIDPOptions, "getGoogleIDPOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/idps/google`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getGoogleIDPOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetGoogleIDP")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGoogleConfigParams)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetGoogleIDP : Update Google IDP configuration
// Configure Google to set up a single sign-on experience for your users. By using Google, users are able to sign in
// with credentials with which they are already familiar. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-social#google" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) SetGoogleIDP(setGoogleIDPOptions *SetGoogleIDPOptions) (result *GoogleConfigParamsPUT, response *core.DetailedResponse, err error) {
	return appIdManagement.SetGoogleIDPWithContext(context.Background(), setGoogleIDPOptions)
}

// SetGoogleIDPWithContext is an alternate form of the SetGoogleIDP method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) SetGoogleIDPWithContext(ctx context.Context, setGoogleIDPOptions *SetGoogleIDPOptions) (result *GoogleConfigParamsPUT, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setGoogleIDPOptions, "setGoogleIDPOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setGoogleIDPOptions, "setGoogleIDPOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/idps/google`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setGoogleIDPOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "SetGoogleIDP")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	_, err = builder.SetBodyContentJSON(setGoogleIDPOptions.IDP)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGoogleConfigParamsPUT)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCustomIDP : Returns the Custom identity configuration
func (appIdManagement *AppIdManagementV4) GetCustomIDP(getCustomIDPOptions *GetCustomIDPOptions) (result *CustomIdPConfigParams, response *core.DetailedResponse, err error) {
	return appIdManagement.GetCustomIDPWithContext(context.Background(), getCustomIDPOptions)
}

// GetCustomIDPWithContext is an alternate form of the GetCustomIDP method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetCustomIDPWithContext(ctx context.Context, getCustomIDPOptions *GetCustomIDPOptions) (result *CustomIdPConfigParams, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCustomIDPOptions, "getCustomIDPOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/idps/custom`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCustomIDPOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetCustomIDP")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomIdPConfigParams)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetCustomIDP : Update or change the configuration of the Custom identity
// Configure App ID Custom identity to allow users to sign-in using your own identity provider.
func (appIdManagement *AppIdManagementV4) SetCustomIDP(setCustomIDPOptions *SetCustomIDPOptions) (result *CustomIdPConfigParams, response *core.DetailedResponse, err error) {
	return appIdManagement.SetCustomIDPWithContext(context.Background(), setCustomIDPOptions)
}

// SetCustomIDPWithContext is an alternate form of the SetCustomIDP method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) SetCustomIDPWithContext(ctx context.Context, setCustomIDPOptions *SetCustomIDPOptions) (result *CustomIdPConfigParams, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setCustomIDPOptions, "setCustomIDPOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setCustomIDPOptions, "setCustomIDPOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/idps/custom`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setCustomIDPOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "SetCustomIDP")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setCustomIDPOptions.IsActive != nil {
		body["isActive"] = setCustomIDPOptions.IsActive
	}
	if setCustomIDPOptions.Config != nil {
		body["config"] = setCustomIDPOptions.Config
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomIdPConfigParams)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCloudDirectoryIDP : Get Cloud Directory IDP configuration
// Returns the Cloud Directory identity provider configuration. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cloud-directory" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) GetCloudDirectoryIDP(getCloudDirectoryIDPOptions *GetCloudDirectoryIDPOptions) (result *CloudDirectoryResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.GetCloudDirectoryIDPWithContext(context.Background(), getCloudDirectoryIDPOptions)
}

// GetCloudDirectoryIDPWithContext is an alternate form of the GetCloudDirectoryIDP method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetCloudDirectoryIDPWithContext(ctx context.Context, getCloudDirectoryIDPOptions *GetCloudDirectoryIDPOptions) (result *CloudDirectoryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getCloudDirectoryIDPOptions, "getCloudDirectoryIDPOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/idps/cloud_directory`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCloudDirectoryIDPOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetCloudDirectoryIDP")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCloudDirectoryResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetCloudDirectoryIDP : Update Cloud Directory IDP configuration
// Configure Cloud Directory to set up a single sign-on experience for your users. With Cloud Directory users can use
// their email and a password of their choice to log in to your applications. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-cloud-directory" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) SetCloudDirectoryIDP(setCloudDirectoryIDPOptions *SetCloudDirectoryIDPOptions) (result *CloudDirectoryResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.SetCloudDirectoryIDPWithContext(context.Background(), setCloudDirectoryIDPOptions)
}

// SetCloudDirectoryIDPWithContext is an alternate form of the SetCloudDirectoryIDP method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) SetCloudDirectoryIDPWithContext(ctx context.Context, setCloudDirectoryIDPOptions *SetCloudDirectoryIDPOptions) (result *CloudDirectoryResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setCloudDirectoryIDPOptions, "setCloudDirectoryIDPOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setCloudDirectoryIDPOptions, "setCloudDirectoryIDPOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/idps/cloud_directory`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setCloudDirectoryIDPOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "SetCloudDirectoryIDP")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setCloudDirectoryIDPOptions.IsActive != nil {
		body["isActive"] = setCloudDirectoryIDPOptions.IsActive
	}
	if setCloudDirectoryIDPOptions.Config != nil {
		body["config"] = setCloudDirectoryIDPOptions.Config
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCloudDirectoryResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetSAMLIDP : Get SAML IDP configuration
// Returns the SAML identity provider configuration, including status and credentials. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-enterprise" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) GetSAMLIDP(getSAMLIDPOptions *GetSAMLIDPOptions) (result *SAMLResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.GetSAMLIDPWithContext(context.Background(), getSAMLIDPOptions)
}

// GetSAMLIDPWithContext is an alternate form of the GetSAMLIDP method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetSAMLIDPWithContext(ctx context.Context, getSAMLIDPOptions *GetSAMLIDPOptions) (result *SAMLResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(getSAMLIDPOptions, "getSAMLIDPOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/idps/saml`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSAMLIDPOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetSAMLIDP")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSAMLResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// SetSAMLIDP : Update SAML IDP configuration
// Configure SAML to set up a single sign-on experience for your users. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-enterprise" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) SetSAMLIDP(setSAMLIDPOptions *SetSAMLIDPOptions) (result *SAMLResponseWithValidationData, response *core.DetailedResponse, err error) {
	return appIdManagement.SetSAMLIDPWithContext(context.Background(), setSAMLIDPOptions)
}

// SetSAMLIDPWithContext is an alternate form of the SetSAMLIDP method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) SetSAMLIDPWithContext(ctx context.Context, setSAMLIDPOptions *SetSAMLIDPOptions) (result *SAMLResponseWithValidationData, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(setSAMLIDPOptions, "setSAMLIDPOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(setSAMLIDPOptions, "setSAMLIDPOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/config/idps/saml`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range setSAMLIDPOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "SetSAMLIDP")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if setSAMLIDPOptions.IsActive != nil {
		body["isActive"] = setSAMLIDPOptions.IsActive
	}
	if setSAMLIDPOptions.Config != nil {
		body["config"] = setSAMLIDPOptions.Config
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSAMLResponseWithValidationData)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListRoles : List all roles
// Obtain a list of the roles that are associated with your registered application.
func (appIdManagement *AppIdManagementV4) ListRoles(listRolesOptions *ListRolesOptions) (result *RolesList, response *core.DetailedResponse, err error) {
	return appIdManagement.ListRolesWithContext(context.Background(), listRolesOptions)
}

// ListRolesWithContext is an alternate form of the ListRoles method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) ListRolesWithContext(ctx context.Context, listRolesOptions *ListRolesOptions) (result *RolesList, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listRolesOptions, "listRolesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/roles`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listRolesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "ListRoles")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalRolesList)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateRole : Create a role
// Create a new role for a registered application.
func (appIdManagement *AppIdManagementV4) CreateRole(createRoleOptions *CreateRoleOptions) (result *CreateRolesResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.CreateRoleWithContext(context.Background(), createRoleOptions)
}

// CreateRoleWithContext is an alternate form of the CreateRole method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) CreateRoleWithContext(ctx context.Context, createRoleOptions *CreateRoleOptions) (result *CreateRolesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createRoleOptions, "createRoleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createRoleOptions, "createRoleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/roles`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createRoleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "CreateRole")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createRoleOptions.Name != nil {
		body["name"] = createRoleOptions.Name
	}
	if createRoleOptions.Access != nil {
		body["access"] = createRoleOptions.Access
	}
	if createRoleOptions.Description != nil {
		body["description"] = createRoleOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCreateRolesResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetRole : View a specific role
// By using the role ID, obtain the information for a specific role that is associated with a registered application.
func (appIdManagement *AppIdManagementV4) GetRole(getRoleOptions *GetRoleOptions) (result *GetRoleResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.GetRoleWithContext(context.Background(), getRoleOptions)
}

// GetRoleWithContext is an alternate form of the GetRole method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetRoleWithContext(ctx context.Context, getRoleOptions *GetRoleOptions) (result *GetRoleResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getRoleOptions, "getRoleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getRoleOptions, "getRoleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"roleId": *getRoleOptions.RoleID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/roles/{roleId}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getRoleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetRole")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetRoleResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateRole : Update a role
// Update an existing role.
func (appIdManagement *AppIdManagementV4) UpdateRole(updateRoleOptions *UpdateRoleOptions) (result *UpdateRolesResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.UpdateRoleWithContext(context.Background(), updateRoleOptions)
}

// UpdateRoleWithContext is an alternate form of the UpdateRole method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UpdateRoleWithContext(ctx context.Context, updateRoleOptions *UpdateRoleOptions) (result *UpdateRolesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateRoleOptions, "updateRoleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateRoleOptions, "updateRoleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"roleId": *updateRoleOptions.RoleID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/roles/{roleId}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateRoleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UpdateRole")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateRoleOptions.Name != nil {
		body["name"] = updateRoleOptions.Name
	}
	if updateRoleOptions.Access != nil {
		body["access"] = updateRoleOptions.Access
	}
	if updateRoleOptions.Description != nil {
		body["description"] = updateRoleOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUpdateRolesResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteRole : Delete a role
// Delete an existing role.
func (appIdManagement *AppIdManagementV4) DeleteRole(deleteRoleOptions *DeleteRoleOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.DeleteRoleWithContext(context.Background(), deleteRoleOptions)
}

// DeleteRoleWithContext is an alternate form of the DeleteRole method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) DeleteRoleWithContext(ctx context.Context, deleteRoleOptions *DeleteRoleOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteRoleOptions, "deleteRoleOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteRoleOptions, "deleteRoleOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"roleId": *deleteRoleOptions.RoleID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/roles/{roleId}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteRoleOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "DeleteRole")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// UsersSearchUserProfile : Search users
// Returns list of users, if given email/id returns only users which match the email/id - not including anonymous
// profiles. <a href="https://cloud.ibm.com/docs/appid?topic=appid-profiles" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) UsersSearchUserProfile(usersSearchUserProfileOptions *UsersSearchUserProfileOptions) (result *UserSearchResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.UsersSearchUserProfileWithContext(context.Background(), usersSearchUserProfileOptions)
}

// UsersSearchUserProfileWithContext is an alternate form of the UsersSearchUserProfile method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UsersSearchUserProfileWithContext(ctx context.Context, usersSearchUserProfileOptions *UsersSearchUserProfileOptions) (result *UserSearchResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(usersSearchUserProfileOptions, "usersSearchUserProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(usersSearchUserProfileOptions, "usersSearchUserProfileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/users`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range usersSearchUserProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UsersSearchUserProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("dataScope", fmt.Sprint(*usersSearchUserProfileOptions.DataScope))
	if usersSearchUserProfileOptions.Email != nil {
		builder.AddQuery("email", fmt.Sprint(*usersSearchUserProfileOptions.Email))
	}
	if usersSearchUserProfileOptions.ID != nil {
		builder.AddQuery("id", fmt.Sprint(*usersSearchUserProfileOptions.ID))
	}
	if usersSearchUserProfileOptions.StartIndex != nil {
		builder.AddQuery("startIndex", fmt.Sprint(*usersSearchUserProfileOptions.StartIndex))
	}
	if usersSearchUserProfileOptions.Count != nil {
		builder.AddQuery("count", fmt.Sprint(*usersSearchUserProfileOptions.Count))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUserSearchResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UsersNominateUser : Pre-register a user profile
// Create a profile for a user that you know needs access to your app before they sign in to your app for the first
// time. <a href="https://cloud.ibm.com/docs/appid?topic=appid-preregister" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) UsersNominateUser(usersNominateUserOptions *UsersNominateUserOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.UsersNominateUserWithContext(context.Background(), usersNominateUserOptions)
}

// UsersNominateUserWithContext is an alternate form of the UsersNominateUser method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UsersNominateUserWithContext(ctx context.Context, usersNominateUserOptions *UsersNominateUserOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(usersNominateUserOptions, "usersNominateUserOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(usersNominateUserOptions, "usersNominateUserOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/users`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range usersNominateUserOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UsersNominateUser")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if usersNominateUserOptions.IDP != nil {
		body["idp"] = usersNominateUserOptions.IDP
	}
	if usersNominateUserOptions.IDPIdentity != nil {
		body["idp-identity"] = usersNominateUserOptions.IDPIdentity
	}
	if usersNominateUserOptions.Profile != nil {
		body["profile"] = usersNominateUserOptions.Profile
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// UserProfilesExport : Export user profiles
// Exports App ID user profiles, not including Cloud Directory and anonymous users.
func (appIdManagement *AppIdManagementV4) UserProfilesExport(userProfilesExportOptions *UserProfilesExportOptions) (result *ExportUserProfile, response *core.DetailedResponse, err error) {
	return appIdManagement.UserProfilesExportWithContext(context.Background(), userProfilesExportOptions)
}

// UserProfilesExportWithContext is an alternate form of the UserProfilesExport method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UserProfilesExportWithContext(ctx context.Context, userProfilesExportOptions *UserProfilesExportOptions) (result *ExportUserProfile, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(userProfilesExportOptions, "userProfilesExportOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/users/export`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range userProfilesExportOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UserProfilesExport")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if userProfilesExportOptions.StartIndex != nil {
		builder.AddQuery("startIndex", fmt.Sprint(*userProfilesExportOptions.StartIndex))
	}
	if userProfilesExportOptions.Count != nil {
		builder.AddQuery("count", fmt.Sprint(*userProfilesExportOptions.Count))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalExportUserProfile)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UserProfilesImport : Import user profiles
// Imports App ID user profiles, not including Cloud Directory and anonymous users.
func (appIdManagement *AppIdManagementV4) UserProfilesImport(userProfilesImportOptions *UserProfilesImportOptions) (result *ImportProfilesResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.UserProfilesImportWithContext(context.Background(), userProfilesImportOptions)
}

// UserProfilesImportWithContext is an alternate form of the UserProfilesImport method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UserProfilesImportWithContext(ctx context.Context, userProfilesImportOptions *UserProfilesImportOptions) (result *ImportProfilesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(userProfilesImportOptions, "userProfilesImportOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(userProfilesImportOptions, "userProfilesImportOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/users/import`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range userProfilesImportOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UserProfilesImport")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if userProfilesImportOptions.Users != nil {
		body["users"] = userProfilesImportOptions.Users
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalImportProfilesResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UsersDeleteUserProfile : Delete user
// Deletes a user by id. <a href="https://cloud.ibm.com/docs/appid?topic=appid-profiles" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) UsersDeleteUserProfile(usersDeleteUserProfileOptions *UsersDeleteUserProfileOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.UsersDeleteUserProfileWithContext(context.Background(), usersDeleteUserProfileOptions)
}

// UsersDeleteUserProfileWithContext is an alternate form of the UsersDeleteUserProfile method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UsersDeleteUserProfileWithContext(ctx context.Context, usersDeleteUserProfileOptions *UsersDeleteUserProfileOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(usersDeleteUserProfileOptions, "usersDeleteUserProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(usersDeleteUserProfileOptions, "usersDeleteUserProfileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"id": *usersDeleteUserProfileOptions.ID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/users/{id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range usersDeleteUserProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UsersDeleteUserProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// UsersRevokeRefreshToken : Revoke refresh token
// Revokes all the refresh tokens issued for the given user. <a
// href="https://cloud.ibm.com/docs/appid?topic=appid-profiles" target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) UsersRevokeRefreshToken(usersRevokeRefreshTokenOptions *UsersRevokeRefreshTokenOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.UsersRevokeRefreshTokenWithContext(context.Background(), usersRevokeRefreshTokenOptions)
}

// UsersRevokeRefreshTokenWithContext is an alternate form of the UsersRevokeRefreshToken method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UsersRevokeRefreshTokenWithContext(ctx context.Context, usersRevokeRefreshTokenOptions *UsersRevokeRefreshTokenOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(usersRevokeRefreshTokenOptions, "usersRevokeRefreshTokenOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(usersRevokeRefreshTokenOptions, "usersRevokeRefreshTokenOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"id": *usersRevokeRefreshTokenOptions.ID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/users/{id}/revoke_refresh_token`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range usersRevokeRefreshTokenOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UsersRevokeRefreshToken")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// UsersGetUserProfile : Get user profile
// Returns the profile of a given user. <a href="https://cloud.ibm.com/docs/appid?topic=appid-profiles"
// target="_blank">Learn more</a>.
func (appIdManagement *AppIdManagementV4) UsersGetUserProfile(usersGetUserProfileOptions *UsersGetUserProfileOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.UsersGetUserProfileWithContext(context.Background(), usersGetUserProfileOptions)
}

// UsersGetUserProfileWithContext is an alternate form of the UsersGetUserProfile method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UsersGetUserProfileWithContext(ctx context.Context, usersGetUserProfileOptions *UsersGetUserProfileOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(usersGetUserProfileOptions, "usersGetUserProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(usersGetUserProfileOptions, "usersGetUserProfileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"id": *usersGetUserProfileOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/users/{id}/profile`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range usersGetUserProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UsersGetUserProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// UsersSetUserProfile : Update user profile
// Updates a user profile. <a href="https://cloud.ibm.com/docs/appid?topic=appid-profiles" target="_blank">Learn
// more</a>.
func (appIdManagement *AppIdManagementV4) UsersSetUserProfile(usersSetUserProfileOptions *UsersSetUserProfileOptions) (response *core.DetailedResponse, err error) {
	return appIdManagement.UsersSetUserProfileWithContext(context.Background(), usersSetUserProfileOptions)
}

// UsersSetUserProfileWithContext is an alternate form of the UsersSetUserProfile method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UsersSetUserProfileWithContext(ctx context.Context, usersSetUserProfileOptions *UsersSetUserProfileOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(usersSetUserProfileOptions, "usersSetUserProfileOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(usersSetUserProfileOptions, "usersSetUserProfileOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"id": *usersSetUserProfileOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/users/{id}/profile`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range usersSetUserProfileOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UsersSetUserProfile")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if usersSetUserProfileOptions.Attributes != nil {
		body["attributes"] = usersSetUserProfileOptions.Attributes
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = appIdManagement.Service.Request(request, nil)

	return
}

// GetUserRoles : Get a user's roles
// View a list of roles that are associated with a specific user.
func (appIdManagement *AppIdManagementV4) GetUserRoles(getUserRolesOptions *GetUserRolesOptions) (result *GetUserRolesResponse, response *core.DetailedResponse, err error) {
	return appIdManagement.GetUserRolesWithContext(context.Background(), getUserRolesOptions)
}

// GetUserRolesWithContext is an alternate form of the GetUserRoles method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) GetUserRolesWithContext(ctx context.Context, getUserRolesOptions *GetUserRolesOptions) (result *GetUserRolesResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getUserRolesOptions, "getUserRolesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getUserRolesOptions, "getUserRolesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"id": *getUserRolesOptions.ID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/users/{id}/roles`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getUserRolesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "GetUserRoles")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalGetUserRolesResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateUserRoles : Update a user's roles
// Update which roles are associated with a specific user or assign a role to a user for the first time.
func (appIdManagement *AppIdManagementV4) UpdateUserRoles(updateUserRolesOptions *UpdateUserRolesOptions) (result *AssignRoleToUser, response *core.DetailedResponse, err error) {
	return appIdManagement.UpdateUserRolesWithContext(context.Background(), updateUserRolesOptions)
}

// UpdateUserRolesWithContext is an alternate form of the UpdateUserRoles method which supports a Context parameter
func (appIdManagement *AppIdManagementV4) UpdateUserRolesWithContext(ctx context.Context, updateUserRolesOptions *UpdateUserRolesOptions) (result *AssignRoleToUser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateUserRolesOptions, "updateUserRolesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateUserRolesOptions, "updateUserRolesOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"tenantId": *appIdManagement.TenantID,
		"id": *updateUserRolesOptions.ID,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = appIdManagement.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(appIdManagement.Service.Options.URL, `/management/v4/{tenantId}/users/{id}/roles`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateUserRolesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("app_id_management", "V4", "UpdateUserRoles")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateUserRolesOptions.Roles != nil {
		body["roles"] = updateUserRolesOptions.Roles
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = appIdManagement.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalAssignRoleToUser)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ApmSchemaAdvancedPasswordManagement : ApmSchemaAdvancedPasswordManagement struct
type ApmSchemaAdvancedPasswordManagement struct {
	Enabled *bool `json:"enabled" validate:"required"`

	PasswordReuse *ApmSchemaAdvancedPasswordManagementPasswordReuse `json:"passwordReuse" validate:"required"`

	PreventPasswordWithUsername *ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername `json:"preventPasswordWithUsername" validate:"required"`

	PasswordExpiration *ApmSchemaAdvancedPasswordManagementPasswordExpiration `json:"passwordExpiration" validate:"required"`

	LockOutPolicy *ApmSchemaAdvancedPasswordManagementLockOutPolicy `json:"lockOutPolicy" validate:"required"`

	MinPasswordChangeInterval *ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval `json:"minPasswordChangeInterval,omitempty"`
}

// NewApmSchemaAdvancedPasswordManagement : Instantiate ApmSchemaAdvancedPasswordManagement (Generic Model Constructor)
func (*AppIdManagementV4) NewApmSchemaAdvancedPasswordManagement(enabled bool, passwordReuse *ApmSchemaAdvancedPasswordManagementPasswordReuse, preventPasswordWithUsername *ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername, passwordExpiration *ApmSchemaAdvancedPasswordManagementPasswordExpiration, lockOutPolicy *ApmSchemaAdvancedPasswordManagementLockOutPolicy) (model *ApmSchemaAdvancedPasswordManagement, err error) {
	model = &ApmSchemaAdvancedPasswordManagement{
		Enabled: core.BoolPtr(enabled),
		PasswordReuse: passwordReuse,
		PreventPasswordWithUsername: preventPasswordWithUsername,
		PasswordExpiration: passwordExpiration,
		LockOutPolicy: lockOutPolicy,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalApmSchemaAdvancedPasswordManagement unmarshals an instance of ApmSchemaAdvancedPasswordManagement from the specified map of raw messages.
func UnmarshalApmSchemaAdvancedPasswordManagement(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApmSchemaAdvancedPasswordManagement)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "passwordReuse", &obj.PasswordReuse, UnmarshalApmSchemaAdvancedPasswordManagementPasswordReuse)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "preventPasswordWithUsername", &obj.PreventPasswordWithUsername, UnmarshalApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "passwordExpiration", &obj.PasswordExpiration, UnmarshalApmSchemaAdvancedPasswordManagementPasswordExpiration)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "lockOutPolicy", &obj.LockOutPolicy, UnmarshalApmSchemaAdvancedPasswordManagementLockOutPolicy)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "minPasswordChangeInterval", &obj.MinPasswordChangeInterval, UnmarshalApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApmSchemaAdvancedPasswordManagementLockOutPolicy : ApmSchemaAdvancedPasswordManagementLockOutPolicy struct
type ApmSchemaAdvancedPasswordManagementLockOutPolicy struct {
	Enabled *bool `json:"enabled" validate:"required"`

	Config *ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig `json:"config,omitempty"`
}

// NewApmSchemaAdvancedPasswordManagementLockOutPolicy : Instantiate ApmSchemaAdvancedPasswordManagementLockOutPolicy (Generic Model Constructor)
func (*AppIdManagementV4) NewApmSchemaAdvancedPasswordManagementLockOutPolicy(enabled bool) (model *ApmSchemaAdvancedPasswordManagementLockOutPolicy, err error) {
	model = &ApmSchemaAdvancedPasswordManagementLockOutPolicy{
		Enabled: core.BoolPtr(enabled),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalApmSchemaAdvancedPasswordManagementLockOutPolicy unmarshals an instance of ApmSchemaAdvancedPasswordManagementLockOutPolicy from the specified map of raw messages.
func UnmarshalApmSchemaAdvancedPasswordManagementLockOutPolicy(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApmSchemaAdvancedPasswordManagementLockOutPolicy)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalApmSchemaAdvancedPasswordManagementLockOutPolicyConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig : ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig struct
type ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig struct {
	LockOutTimeSec *float64 `json:"lockOutTimeSec" validate:"required"`

	NumOfAttempts *float64 `json:"numOfAttempts" validate:"required"`
}

// NewApmSchemaAdvancedPasswordManagementLockOutPolicyConfig : Instantiate ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig (Generic Model Constructor)
func (*AppIdManagementV4) NewApmSchemaAdvancedPasswordManagementLockOutPolicyConfig(lockOutTimeSec float64, numOfAttempts float64) (model *ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig, err error) {
	model = &ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig{
		LockOutTimeSec: core.Float64Ptr(lockOutTimeSec),
		NumOfAttempts: core.Float64Ptr(numOfAttempts),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalApmSchemaAdvancedPasswordManagementLockOutPolicyConfig unmarshals an instance of ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig from the specified map of raw messages.
func UnmarshalApmSchemaAdvancedPasswordManagementLockOutPolicyConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig)
	err = core.UnmarshalPrimitive(m, "lockOutTimeSec", &obj.LockOutTimeSec)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "numOfAttempts", &obj.NumOfAttempts)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval : ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval struct
type ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval struct {
	Enabled *bool `json:"enabled" validate:"required"`

	Config *ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig `json:"config,omitempty"`
}

// NewApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval : Instantiate ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval (Generic Model Constructor)
func (*AppIdManagementV4) NewApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval(enabled bool) (model *ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval, err error) {
	model = &ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval{
		Enabled: core.BoolPtr(enabled),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval unmarshals an instance of ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval from the specified map of raw messages.
func UnmarshalApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig : ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig struct
type ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig struct {
	MinHoursToChangePassword *float64 `json:"minHoursToChangePassword" validate:"required"`
}

// NewApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig : Instantiate ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig (Generic Model Constructor)
func (*AppIdManagementV4) NewApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig(minHoursToChangePassword float64) (model *ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig, err error) {
	model = &ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig{
		MinHoursToChangePassword: core.Float64Ptr(minHoursToChangePassword),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig unmarshals an instance of ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig from the specified map of raw messages.
func UnmarshalApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig)
	err = core.UnmarshalPrimitive(m, "minHoursToChangePassword", &obj.MinHoursToChangePassword)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApmSchemaAdvancedPasswordManagementPasswordExpiration : ApmSchemaAdvancedPasswordManagementPasswordExpiration struct
type ApmSchemaAdvancedPasswordManagementPasswordExpiration struct {
	Enabled *bool `json:"enabled" validate:"required"`

	Config *ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig `json:"config,omitempty"`
}

// NewApmSchemaAdvancedPasswordManagementPasswordExpiration : Instantiate ApmSchemaAdvancedPasswordManagementPasswordExpiration (Generic Model Constructor)
func (*AppIdManagementV4) NewApmSchemaAdvancedPasswordManagementPasswordExpiration(enabled bool) (model *ApmSchemaAdvancedPasswordManagementPasswordExpiration, err error) {
	model = &ApmSchemaAdvancedPasswordManagementPasswordExpiration{
		Enabled: core.BoolPtr(enabled),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalApmSchemaAdvancedPasswordManagementPasswordExpiration unmarshals an instance of ApmSchemaAdvancedPasswordManagementPasswordExpiration from the specified map of raw messages.
func UnmarshalApmSchemaAdvancedPasswordManagementPasswordExpiration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApmSchemaAdvancedPasswordManagementPasswordExpiration)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalApmSchemaAdvancedPasswordManagementPasswordExpirationConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig : ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig struct
type ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig struct {
	DaysToExpire *float64 `json:"daysToExpire" validate:"required"`
}

// NewApmSchemaAdvancedPasswordManagementPasswordExpirationConfig : Instantiate ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig (Generic Model Constructor)
func (*AppIdManagementV4) NewApmSchemaAdvancedPasswordManagementPasswordExpirationConfig(daysToExpire float64) (model *ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig, err error) {
	model = &ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig{
		DaysToExpire: core.Float64Ptr(daysToExpire),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalApmSchemaAdvancedPasswordManagementPasswordExpirationConfig unmarshals an instance of ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig from the specified map of raw messages.
func UnmarshalApmSchemaAdvancedPasswordManagementPasswordExpirationConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig)
	err = core.UnmarshalPrimitive(m, "daysToExpire", &obj.DaysToExpire)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApmSchemaAdvancedPasswordManagementPasswordReuse : ApmSchemaAdvancedPasswordManagementPasswordReuse struct
type ApmSchemaAdvancedPasswordManagementPasswordReuse struct {
	Enabled *bool `json:"enabled" validate:"required"`

	Config *ApmSchemaAdvancedPasswordManagementPasswordReuseConfig `json:"config,omitempty"`
}

// NewApmSchemaAdvancedPasswordManagementPasswordReuse : Instantiate ApmSchemaAdvancedPasswordManagementPasswordReuse (Generic Model Constructor)
func (*AppIdManagementV4) NewApmSchemaAdvancedPasswordManagementPasswordReuse(enabled bool) (model *ApmSchemaAdvancedPasswordManagementPasswordReuse, err error) {
	model = &ApmSchemaAdvancedPasswordManagementPasswordReuse{
		Enabled: core.BoolPtr(enabled),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalApmSchemaAdvancedPasswordManagementPasswordReuse unmarshals an instance of ApmSchemaAdvancedPasswordManagementPasswordReuse from the specified map of raw messages.
func UnmarshalApmSchemaAdvancedPasswordManagementPasswordReuse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApmSchemaAdvancedPasswordManagementPasswordReuse)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalApmSchemaAdvancedPasswordManagementPasswordReuseConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApmSchemaAdvancedPasswordManagementPasswordReuseConfig : ApmSchemaAdvancedPasswordManagementPasswordReuseConfig struct
type ApmSchemaAdvancedPasswordManagementPasswordReuseConfig struct {
	MaxPasswordReuse *float64 `json:"maxPasswordReuse" validate:"required"`
}

// NewApmSchemaAdvancedPasswordManagementPasswordReuseConfig : Instantiate ApmSchemaAdvancedPasswordManagementPasswordReuseConfig (Generic Model Constructor)
func (*AppIdManagementV4) NewApmSchemaAdvancedPasswordManagementPasswordReuseConfig(maxPasswordReuse float64) (model *ApmSchemaAdvancedPasswordManagementPasswordReuseConfig, err error) {
	model = &ApmSchemaAdvancedPasswordManagementPasswordReuseConfig{
		MaxPasswordReuse: core.Float64Ptr(maxPasswordReuse),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalApmSchemaAdvancedPasswordManagementPasswordReuseConfig unmarshals an instance of ApmSchemaAdvancedPasswordManagementPasswordReuseConfig from the specified map of raw messages.
func UnmarshalApmSchemaAdvancedPasswordManagementPasswordReuseConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApmSchemaAdvancedPasswordManagementPasswordReuseConfig)
	err = core.UnmarshalPrimitive(m, "maxPasswordReuse", &obj.MaxPasswordReuse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername : ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername struct
type ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername struct {
	Enabled *bool `json:"enabled" validate:"required"`
}

// NewApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername : Instantiate ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername (Generic Model Constructor)
func (*AppIdManagementV4) NewApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername(enabled bool) (model *ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername, err error) {
	model = &ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername{
		Enabled: core.BoolPtr(enabled),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername unmarshals an instance of ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername from the specified map of raw messages.
func UnmarshalApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername)
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Application : Application struct
type Application struct {
	// The application clientId.
	ClientID *string `json:"clientId,omitempty"`

	// The service tenantId.
	TenantID *string `json:"tenantId,omitempty"`

	Secret *string `json:"secret,omitempty"`

	// The application name.
	Name *string `json:"name,omitempty"`

	OAuthServerURL *string `json:"oAuthServerUrl,omitempty"`

	// The type of application. Allowed types are regularwebapp and singlepageapp.
	Type *string `json:"type,omitempty"`
}

// UnmarshalApplication unmarshals an instance of Application from the specified map of raw messages.
func UnmarshalApplication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Application)
	err = core.UnmarshalPrimitive(m, "clientId", &obj.ClientID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tenantId", &obj.TenantID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "secret", &obj.Secret)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "oAuthServerUrl", &obj.OAuthServerURL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApplicationsList : ApplicationsList struct
type ApplicationsList struct {
	Applications []Application `json:"applications" validate:"required"`
}

// UnmarshalApplicationsList unmarshals an instance of ApplicationsList from the specified map of raw messages.
func UnmarshalApplicationsList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApplicationsList)
	err = core.UnmarshalModel(m, "applications", &obj.Applications, UnmarshalApplication)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssignRoleToUserRolesItem : AssignRoleToUserRolesItem struct
type AssignRoleToUserRolesItem struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

// UnmarshalAssignRoleToUserRolesItem unmarshals an instance of AssignRoleToUserRolesItem from the specified map of raw messages.
func UnmarshalAssignRoleToUserRolesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssignRoleToUserRolesItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ChangePasswordOptions : The ChangePassword options.
type ChangePasswordOptions struct {
	// The new password.
	NewPassword *string `validate:"required"`

	// The Cloud Directory unique user Id.
	UUID *string `validate:"required"`

	// The ip address the password changed from.
	ChangedIpAddress *string

	// Preferred language for resource. Format as described at RFC5646.
	Language *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewChangePasswordOptions : Instantiate ChangePasswordOptions
func (*AppIdManagementV4) NewChangePasswordOptions(newPassword string, uuid string) *ChangePasswordOptions {
	return &ChangePasswordOptions{
		NewPassword: core.StringPtr(newPassword),
		UUID: core.StringPtr(uuid),
	}
}

// SetNewPassword : Allow user to set NewPassword
func (options *ChangePasswordOptions) SetNewPassword(newPassword string) *ChangePasswordOptions {
	options.NewPassword = core.StringPtr(newPassword)
	return options
}

// SetUUID : Allow user to set UUID
func (options *ChangePasswordOptions) SetUUID(uuid string) *ChangePasswordOptions {
	options.UUID = core.StringPtr(uuid)
	return options
}

// SetChangedIpAddress : Allow user to set ChangedIpAddress
func (options *ChangePasswordOptions) SetChangedIpAddress(changedIpAddress string) *ChangePasswordOptions {
	options.ChangedIpAddress = core.StringPtr(changedIpAddress)
	return options
}

// SetLanguage : Allow user to set Language
func (options *ChangePasswordOptions) SetLanguage(language string) *ChangePasswordOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ChangePasswordOptions) SetHeaders(param map[string]string) *ChangePasswordOptions {
	options.Headers = param
	return options
}

// CloudDirectoryConfigParamsInteractions : CloudDirectoryConfigParamsInteractions struct
type CloudDirectoryConfigParamsInteractions struct {
	IdentityConfirmation *CloudDirectoryConfigParamsInteractionsIdentityConfirmation `json:"identityConfirmation" validate:"required"`

	WelcomeEnabled *bool `json:"welcomeEnabled" validate:"required"`

	ResetPasswordEnabled *bool `json:"resetPasswordEnabled" validate:"required"`

	ResetPasswordNotificationEnable *bool `json:"resetPasswordNotificationEnable" validate:"required"`
}

// NewCloudDirectoryConfigParamsInteractions : Instantiate CloudDirectoryConfigParamsInteractions (Generic Model Constructor)
func (*AppIdManagementV4) NewCloudDirectoryConfigParamsInteractions(identityConfirmation *CloudDirectoryConfigParamsInteractionsIdentityConfirmation, welcomeEnabled bool, resetPasswordEnabled bool, resetPasswordNotificationEnable bool) (model *CloudDirectoryConfigParamsInteractions, err error) {
	model = &CloudDirectoryConfigParamsInteractions{
		IdentityConfirmation: identityConfirmation,
		WelcomeEnabled: core.BoolPtr(welcomeEnabled),
		ResetPasswordEnabled: core.BoolPtr(resetPasswordEnabled),
		ResetPasswordNotificationEnable: core.BoolPtr(resetPasswordNotificationEnable),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCloudDirectoryConfigParamsInteractions unmarshals an instance of CloudDirectoryConfigParamsInteractions from the specified map of raw messages.
func UnmarshalCloudDirectoryConfigParamsInteractions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CloudDirectoryConfigParamsInteractions)
	err = core.UnmarshalModel(m, "identityConfirmation", &obj.IdentityConfirmation, UnmarshalCloudDirectoryConfigParamsInteractionsIdentityConfirmation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "welcomeEnabled", &obj.WelcomeEnabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resetPasswordEnabled", &obj.ResetPasswordEnabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resetPasswordNotificationEnable", &obj.ResetPasswordNotificationEnable)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CloudDirectoryConfigParamsInteractionsIdentityConfirmation : CloudDirectoryConfigParamsInteractionsIdentityConfirmation struct
type CloudDirectoryConfigParamsInteractionsIdentityConfirmation struct {
	AccessMode *string `json:"accessMode" validate:"required"`

	Methods []string `json:"methods,omitempty"`
}

// Constants associated with the CloudDirectoryConfigParamsInteractionsIdentityConfirmation.AccessMode property.
const (
	CloudDirectoryConfigParamsInteractionsIdentityConfirmation_AccessMode_False = "false"
	CloudDirectoryConfigParamsInteractionsIdentityConfirmation_AccessMode_Full = "FULL"
	CloudDirectoryConfigParamsInteractionsIdentityConfirmation_AccessMode_Restrictive = "RESTRICTIVE"
)

// Constants associated with the CloudDirectoryConfigParamsInteractionsIdentityConfirmation.Methods property.
const (
	CloudDirectoryConfigParamsInteractionsIdentityConfirmation_Methods_Email = "email"
)

// NewCloudDirectoryConfigParamsInteractionsIdentityConfirmation : Instantiate CloudDirectoryConfigParamsInteractionsIdentityConfirmation (Generic Model Constructor)
func (*AppIdManagementV4) NewCloudDirectoryConfigParamsInteractionsIdentityConfirmation(accessMode string) (model *CloudDirectoryConfigParamsInteractionsIdentityConfirmation, err error) {
	model = &CloudDirectoryConfigParamsInteractionsIdentityConfirmation{
		AccessMode: core.StringPtr(accessMode),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCloudDirectoryConfigParamsInteractionsIdentityConfirmation unmarshals an instance of CloudDirectoryConfigParamsInteractionsIdentityConfirmation from the specified map of raw messages.
func UnmarshalCloudDirectoryConfigParamsInteractionsIdentityConfirmation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CloudDirectoryConfigParamsInteractionsIdentityConfirmation)
	err = core.UnmarshalPrimitive(m, "accessMode", &obj.AccessMode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "methods", &obj.Methods)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CloudDirectoryExportOptions : The CloudDirectoryExport options.
type CloudDirectoryExportOptions struct {
	// A custom string that will be use to encrypt and decrypt the users hashed password.
	EncryptionSecret *string `validate:"required"`

	// The first result in a set list of results.
	StartIndex *int64

	// The maximum number of results per page. Limit to 50 users per request.
	Count *int64

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCloudDirectoryExportOptions : Instantiate CloudDirectoryExportOptions
func (*AppIdManagementV4) NewCloudDirectoryExportOptions(encryptionSecret string) *CloudDirectoryExportOptions {
	return &CloudDirectoryExportOptions{
		EncryptionSecret: core.StringPtr(encryptionSecret),
	}
}

// SetEncryptionSecret : Allow user to set EncryptionSecret
func (options *CloudDirectoryExportOptions) SetEncryptionSecret(encryptionSecret string) *CloudDirectoryExportOptions {
	options.EncryptionSecret = core.StringPtr(encryptionSecret)
	return options
}

// SetStartIndex : Allow user to set StartIndex
func (options *CloudDirectoryExportOptions) SetStartIndex(startIndex int64) *CloudDirectoryExportOptions {
	options.StartIndex = core.Int64Ptr(startIndex)
	return options
}

// SetCount : Allow user to set Count
func (options *CloudDirectoryExportOptions) SetCount(count int64) *CloudDirectoryExportOptions {
	options.Count = core.Int64Ptr(count)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CloudDirectoryExportOptions) SetHeaders(param map[string]string) *CloudDirectoryExportOptions {
	options.Headers = param
	return options
}

// CloudDirectoryGetUserinfoOptions : The CloudDirectoryGetUserinfo options.
type CloudDirectoryGetUserinfoOptions struct {
	// The ID assigned to a user when they sign in by using Cloud Directory.
	UserID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCloudDirectoryGetUserinfoOptions : Instantiate CloudDirectoryGetUserinfoOptions
func (*AppIdManagementV4) NewCloudDirectoryGetUserinfoOptions(userID string) *CloudDirectoryGetUserinfoOptions {
	return &CloudDirectoryGetUserinfoOptions{
		UserID: core.StringPtr(userID),
	}
}

// SetUserID : Allow user to set UserID
func (options *CloudDirectoryGetUserinfoOptions) SetUserID(userID string) *CloudDirectoryGetUserinfoOptions {
	options.UserID = core.StringPtr(userID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CloudDirectoryGetUserinfoOptions) SetHeaders(param map[string]string) *CloudDirectoryGetUserinfoOptions {
	options.Headers = param
	return options
}

// CloudDirectoryImportOptions : The CloudDirectoryImport options.
type CloudDirectoryImportOptions struct {
	// A custom string that will be use to encrypt and decrypt the users hashed password.
	EncryptionSecret *string `validate:"required"`

	Users []ExportUserUsersItem `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCloudDirectoryImportOptions : Instantiate CloudDirectoryImportOptions
func (*AppIdManagementV4) NewCloudDirectoryImportOptions(encryptionSecret string, users []ExportUserUsersItem) *CloudDirectoryImportOptions {
	return &CloudDirectoryImportOptions{
		EncryptionSecret: core.StringPtr(encryptionSecret),
		Users: users,
	}
}

// SetEncryptionSecret : Allow user to set EncryptionSecret
func (options *CloudDirectoryImportOptions) SetEncryptionSecret(encryptionSecret string) *CloudDirectoryImportOptions {
	options.EncryptionSecret = core.StringPtr(encryptionSecret)
	return options
}

// SetUsers : Allow user to set Users
func (options *CloudDirectoryImportOptions) SetUsers(users []ExportUserUsersItem) *CloudDirectoryImportOptions {
	options.Users = users
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CloudDirectoryImportOptions) SetHeaders(param map[string]string) *CloudDirectoryImportOptions {
	options.Headers = param
	return options
}

// CloudDirectoryRemoveOptions : The CloudDirectoryRemove options.
type CloudDirectoryRemoveOptions struct {
	// The ID assigned to a user when they sign in by using Cloud Directory.
	UserID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCloudDirectoryRemoveOptions : Instantiate CloudDirectoryRemoveOptions
func (*AppIdManagementV4) NewCloudDirectoryRemoveOptions(userID string) *CloudDirectoryRemoveOptions {
	return &CloudDirectoryRemoveOptions{
		UserID: core.StringPtr(userID),
	}
}

// SetUserID : Allow user to set UserID
func (options *CloudDirectoryRemoveOptions) SetUserID(userID string) *CloudDirectoryRemoveOptions {
	options.UserID = core.StringPtr(userID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CloudDirectoryRemoveOptions) SetHeaders(param map[string]string) *CloudDirectoryRemoveOptions {
	options.Headers = param
	return options
}

// CloudDirectorySenderDetailsSenderDetails : CloudDirectorySenderDetailsSenderDetails struct
type CloudDirectorySenderDetailsSenderDetails struct {
	From *CloudDirectorySenderDetailsSenderDetailsFrom `json:"from" validate:"required"`

	ReplyTo *CloudDirectorySenderDetailsSenderDetailsReplyTo `json:"reply_to,omitempty"`

	LinkExpirationSec *float64 `json:"linkExpirationSec,omitempty"`
}

// NewCloudDirectorySenderDetailsSenderDetails : Instantiate CloudDirectorySenderDetailsSenderDetails (Generic Model Constructor)
func (*AppIdManagementV4) NewCloudDirectorySenderDetailsSenderDetails(from *CloudDirectorySenderDetailsSenderDetailsFrom) (model *CloudDirectorySenderDetailsSenderDetails, err error) {
	model = &CloudDirectorySenderDetailsSenderDetails{
		From: from,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCloudDirectorySenderDetailsSenderDetails unmarshals an instance of CloudDirectorySenderDetailsSenderDetails from the specified map of raw messages.
func UnmarshalCloudDirectorySenderDetailsSenderDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CloudDirectorySenderDetailsSenderDetails)
	err = core.UnmarshalModel(m, "from", &obj.From, UnmarshalCloudDirectorySenderDetailsSenderDetailsFrom)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "reply_to", &obj.ReplyTo, UnmarshalCloudDirectorySenderDetailsSenderDetailsReplyTo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "linkExpirationSec", &obj.LinkExpirationSec)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CloudDirectorySenderDetailsSenderDetailsFrom : CloudDirectorySenderDetailsSenderDetailsFrom struct
type CloudDirectorySenderDetailsSenderDetailsFrom struct {
	Name *string `json:"name,omitempty"`

	Email *string `json:"email" validate:"required"`
}

// NewCloudDirectorySenderDetailsSenderDetailsFrom : Instantiate CloudDirectorySenderDetailsSenderDetailsFrom (Generic Model Constructor)
func (*AppIdManagementV4) NewCloudDirectorySenderDetailsSenderDetailsFrom(email string) (model *CloudDirectorySenderDetailsSenderDetailsFrom, err error) {
	model = &CloudDirectorySenderDetailsSenderDetailsFrom{
		Email: core.StringPtr(email),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCloudDirectorySenderDetailsSenderDetailsFrom unmarshals an instance of CloudDirectorySenderDetailsSenderDetailsFrom from the specified map of raw messages.
func UnmarshalCloudDirectorySenderDetailsSenderDetailsFrom(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CloudDirectorySenderDetailsSenderDetailsFrom)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CloudDirectorySenderDetailsSenderDetailsReplyTo : CloudDirectorySenderDetailsSenderDetailsReplyTo struct
type CloudDirectorySenderDetailsSenderDetailsReplyTo struct {
	Name *string `json:"name,omitempty"`

	Email *string `json:"email,omitempty"`
}

// UnmarshalCloudDirectorySenderDetailsSenderDetailsReplyTo unmarshals an instance of CloudDirectorySenderDetailsSenderDetailsReplyTo from the specified map of raw messages.
func UnmarshalCloudDirectorySenderDetailsSenderDetailsReplyTo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CloudDirectorySenderDetailsSenderDetailsReplyTo)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateCloudDirectoryUserOptions : The CreateCloudDirectoryUser options.
type CreateCloudDirectoryUserOptions struct {
	Emails []CreateNewUserEmailsItem `validate:"required"`

	Password *string `validate:"required"`

	Active *bool

	UserName *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateCloudDirectoryUserOptions : Instantiate CreateCloudDirectoryUserOptions
func (*AppIdManagementV4) NewCreateCloudDirectoryUserOptions(emails []CreateNewUserEmailsItem, password string) *CreateCloudDirectoryUserOptions {
	return &CreateCloudDirectoryUserOptions{
		Emails: emails,
		Password: core.StringPtr(password),
	}
}

// SetEmails : Allow user to set Emails
func (options *CreateCloudDirectoryUserOptions) SetEmails(emails []CreateNewUserEmailsItem) *CreateCloudDirectoryUserOptions {
	options.Emails = emails
	return options
}

// SetPassword : Allow user to set Password
func (options *CreateCloudDirectoryUserOptions) SetPassword(password string) *CreateCloudDirectoryUserOptions {
	options.Password = core.StringPtr(password)
	return options
}

// SetActive : Allow user to set Active
func (options *CreateCloudDirectoryUserOptions) SetActive(active bool) *CreateCloudDirectoryUserOptions {
	options.Active = core.BoolPtr(active)
	return options
}

// SetUserName : Allow user to set UserName
func (options *CreateCloudDirectoryUserOptions) SetUserName(userName string) *CreateCloudDirectoryUserOptions {
	options.UserName = core.StringPtr(userName)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCloudDirectoryUserOptions) SetHeaders(param map[string]string) *CreateCloudDirectoryUserOptions {
	options.Headers = param
	return options
}

// CreateNewUserEmailsItem : CreateNewUserEmailsItem struct
type CreateNewUserEmailsItem struct {
	Value *string `json:"value" validate:"required"`

	Primary *bool `json:"primary,omitempty"`
}

// NewCreateNewUserEmailsItem : Instantiate CreateNewUserEmailsItem (Generic Model Constructor)
func (*AppIdManagementV4) NewCreateNewUserEmailsItem(value string) (model *CreateNewUserEmailsItem, err error) {
	model = &CreateNewUserEmailsItem{
		Value: core.StringPtr(value),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCreateNewUserEmailsItem unmarshals an instance of CreateNewUserEmailsItem from the specified map of raw messages.
func UnmarshalCreateNewUserEmailsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateNewUserEmailsItem)
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "primary", &obj.Primary)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateRoleOptions : The CreateRole options.
type CreateRoleOptions struct {
	Name *string `validate:"required"`

	Access []CreateRoleParamsAccessItem `validate:"required"`

	Description *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateRoleOptions : Instantiate CreateRoleOptions
func (*AppIdManagementV4) NewCreateRoleOptions(name string, access []CreateRoleParamsAccessItem) *CreateRoleOptions {
	return &CreateRoleOptions{
		Name: core.StringPtr(name),
		Access: access,
	}
}

// SetName : Allow user to set Name
func (options *CreateRoleOptions) SetName(name string) *CreateRoleOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetAccess : Allow user to set Access
func (options *CreateRoleOptions) SetAccess(access []CreateRoleParamsAccessItem) *CreateRoleOptions {
	options.Access = access
	return options
}

// SetDescription : Allow user to set Description
func (options *CreateRoleOptions) SetDescription(description string) *CreateRoleOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *CreateRoleOptions) SetHeaders(param map[string]string) *CreateRoleOptions {
	options.Headers = param
	return options
}

// CreateRoleParamsAccessItem : CreateRoleParamsAccessItem struct
type CreateRoleParamsAccessItem struct {
	ApplicationID *string `json:"application_id" validate:"required"`

	Scopes []string `json:"scopes" validate:"required"`
}

// NewCreateRoleParamsAccessItem : Instantiate CreateRoleParamsAccessItem (Generic Model Constructor)
func (*AppIdManagementV4) NewCreateRoleParamsAccessItem(applicationID string, scopes []string) (model *CreateRoleParamsAccessItem, err error) {
	model = &CreateRoleParamsAccessItem{
		ApplicationID: core.StringPtr(applicationID),
		Scopes: scopes,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCreateRoleParamsAccessItem unmarshals an instance of CreateRoleParamsAccessItem from the specified map of raw messages.
func UnmarshalCreateRoleParamsAccessItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateRoleParamsAccessItem)
	err = core.UnmarshalPrimitive(m, "application_id", &obj.ApplicationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scopes", &obj.Scopes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateRolesResponseAccessItem : CreateRolesResponseAccessItem struct
type CreateRolesResponseAccessItem struct {
	ApplicationID *string `json:"application_id" validate:"required"`

	Scopes []string `json:"scopes" validate:"required"`
}

// UnmarshalCreateRolesResponseAccessItem unmarshals an instance of CreateRolesResponseAccessItem from the specified map of raw messages.
func UnmarshalCreateRolesResponseAccessItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateRolesResponseAccessItem)
	err = core.UnmarshalPrimitive(m, "application_id", &obj.ApplicationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scopes", &obj.Scopes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CustomIdPConfigParamsConfig : CustomIdPConfigParamsConfig struct
type CustomIdPConfigParamsConfig struct {
	PublicKey *string `json:"publicKey,omitempty"`
}

// UnmarshalCustomIdPConfigParamsConfig unmarshals an instance of CustomIdPConfigParamsConfig from the specified map of raw messages.
func UnmarshalCustomIdPConfigParamsConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CustomIdPConfigParamsConfig)
	err = core.UnmarshalPrimitive(m, "publicKey", &obj.PublicKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteActionUrlOptions : The DeleteActionURL options.
type DeleteActionUrlOptions struct {
	// The type of the action. on_user_verified - the URL of your custom user verified page, on_reset_password - the URL of
	// your custom reset password page.
	Action *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the DeleteActionUrlOptions.Action property.
// The type of the action. on_user_verified - the URL of your custom user verified page, on_reset_password - the URL of
// your custom reset password page.
const (
	DeleteActionUrlOptions_Action_OnResetPassword = "on_reset_password"
	DeleteActionUrlOptions_Action_OnUserVerified = "on_user_verified"
)

// NewDeleteActionUrlOptions : Instantiate DeleteActionUrlOptions
func (*AppIdManagementV4) NewDeleteActionUrlOptions(action string) *DeleteActionUrlOptions {
	return &DeleteActionUrlOptions{
		Action: core.StringPtr(action),
	}
}

// SetAction : Allow user to set Action
func (options *DeleteActionUrlOptions) SetAction(action string) *DeleteActionUrlOptions {
	options.Action = core.StringPtr(action)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteActionUrlOptions) SetHeaders(param map[string]string) *DeleteActionUrlOptions {
	options.Headers = param
	return options
}

// DeleteApplicationOptions : The DeleteApplication options.
type DeleteApplicationOptions struct {
	// The application clientId.
	ClientID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteApplicationOptions : Instantiate DeleteApplicationOptions
func (*AppIdManagementV4) NewDeleteApplicationOptions(clientID string) *DeleteApplicationOptions {
	return &DeleteApplicationOptions{
		ClientID: core.StringPtr(clientID),
	}
}

// SetClientID : Allow user to set ClientID
func (options *DeleteApplicationOptions) SetClientID(clientID string) *DeleteApplicationOptions {
	options.ClientID = core.StringPtr(clientID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteApplicationOptions) SetHeaders(param map[string]string) *DeleteApplicationOptions {
	options.Headers = param
	return options
}

// DeleteCloudDirectoryUserOptions : The DeleteCloudDirectoryUser options.
type DeleteCloudDirectoryUserOptions struct {
	// The ID assigned to a user when they sign in by using Cloud Directory.
	UserID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCloudDirectoryUserOptions : Instantiate DeleteCloudDirectoryUserOptions
func (*AppIdManagementV4) NewDeleteCloudDirectoryUserOptions(userID string) *DeleteCloudDirectoryUserOptions {
	return &DeleteCloudDirectoryUserOptions{
		UserID: core.StringPtr(userID),
	}
}

// SetUserID : Allow user to set UserID
func (options *DeleteCloudDirectoryUserOptions) SetUserID(userID string) *DeleteCloudDirectoryUserOptions {
	options.UserID = core.StringPtr(userID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCloudDirectoryUserOptions) SetHeaders(param map[string]string) *DeleteCloudDirectoryUserOptions {
	options.Headers = param
	return options
}

// DeleteRoleOptions : The DeleteRole options.
type DeleteRoleOptions struct {
	// The role identifier.
	RoleID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteRoleOptions : Instantiate DeleteRoleOptions
func (*AppIdManagementV4) NewDeleteRoleOptions(roleID string) *DeleteRoleOptions {
	return &DeleteRoleOptions{
		RoleID: core.StringPtr(roleID),
	}
}

// SetRoleID : Allow user to set RoleID
func (options *DeleteRoleOptions) SetRoleID(roleID string) *DeleteRoleOptions {
	options.RoleID = core.StringPtr(roleID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteRoleOptions) SetHeaders(param map[string]string) *DeleteRoleOptions {
	options.Headers = param
	return options
}

// DeleteTemplateOptions : The DeleteTemplate options.
type DeleteTemplateOptions struct {
	// The type of email template. This can be "USER_VERIFICATION", "WELCOME", "PASSWORD_CHANGED", "RESET_PASSWORD" or
	// "MFA_VERIFICATION".
	TemplateName *string `validate:"required,ne="`

	// Preferred language for resource. Format as described at RFC5646. According to the configured languages codes
	// returned from the `GET /management/v4/{tenantId}/config/ui/languages` API.
	Language *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the DeleteTemplateOptions.TemplateName property.
// The type of email template. This can be "USER_VERIFICATION", "WELCOME", "PASSWORD_CHANGED", "RESET_PASSWORD" or
// "MFA_VERIFICATION".
const (
	DeleteTemplateOptions_TemplateName_MfaVerification = "MFA_VERIFICATION"
	DeleteTemplateOptions_TemplateName_PasswordChanged = "PASSWORD_CHANGED"
	DeleteTemplateOptions_TemplateName_ResetPassword = "RESET_PASSWORD"
	DeleteTemplateOptions_TemplateName_UserVerification = "USER_VERIFICATION"
	DeleteTemplateOptions_TemplateName_Welcome = "WELCOME"
)

// NewDeleteTemplateOptions : Instantiate DeleteTemplateOptions
func (*AppIdManagementV4) NewDeleteTemplateOptions(templateName string, language string) *DeleteTemplateOptions {
	return &DeleteTemplateOptions{
		TemplateName: core.StringPtr(templateName),
		Language: core.StringPtr(language),
	}
}

// SetTemplateName : Allow user to set TemplateName
func (options *DeleteTemplateOptions) SetTemplateName(templateName string) *DeleteTemplateOptions {
	options.TemplateName = core.StringPtr(templateName)
	return options
}

// SetLanguage : Allow user to set Language
func (options *DeleteTemplateOptions) SetLanguage(language string) *DeleteTemplateOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteTemplateOptions) SetHeaders(param map[string]string) *DeleteTemplateOptions {
	options.Headers = param
	return options
}

// EmailDispatcherParamsCustom : EmailDispatcherParamsCustom struct
type EmailDispatcherParamsCustom struct {
	URL *string `json:"url" validate:"required"`

	Authorization *EmailDispatcherParamsCustomAuthorization `json:"authorization" validate:"required"`
}

// NewEmailDispatcherParamsCustom : Instantiate EmailDispatcherParamsCustom (Generic Model Constructor)
func (*AppIdManagementV4) NewEmailDispatcherParamsCustom(url string, authorization *EmailDispatcherParamsCustomAuthorization) (model *EmailDispatcherParamsCustom, err error) {
	model = &EmailDispatcherParamsCustom{
		URL: core.StringPtr(url),
		Authorization: authorization,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalEmailDispatcherParamsCustom unmarshals an instance of EmailDispatcherParamsCustom from the specified map of raw messages.
func UnmarshalEmailDispatcherParamsCustom(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmailDispatcherParamsCustom)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authorization", &obj.Authorization, UnmarshalEmailDispatcherParamsCustomAuthorization)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EmailDispatcherParamsCustomAuthorization : EmailDispatcherParamsCustomAuthorization struct
type EmailDispatcherParamsCustomAuthorization struct {
	Type *string `json:"type" validate:"required"`

	Value *string `json:"value,omitempty"`

	Username *string `json:"username,omitempty"`

	Password *string `json:"password,omitempty"`
}

// Constants associated with the EmailDispatcherParamsCustomAuthorization.Type property.
const (
	EmailDispatcherParamsCustomAuthorization_Type_Basic = "basic"
	EmailDispatcherParamsCustomAuthorization_Type_None = "none"
	EmailDispatcherParamsCustomAuthorization_Type_Value = "value"
)

// NewEmailDispatcherParamsCustomAuthorization : Instantiate EmailDispatcherParamsCustomAuthorization (Generic Model Constructor)
func (*AppIdManagementV4) NewEmailDispatcherParamsCustomAuthorization(typeVar string) (model *EmailDispatcherParamsCustomAuthorization, err error) {
	model = &EmailDispatcherParamsCustomAuthorization{
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalEmailDispatcherParamsCustomAuthorization unmarshals an instance of EmailDispatcherParamsCustomAuthorization from the specified map of raw messages.
func UnmarshalEmailDispatcherParamsCustomAuthorization(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmailDispatcherParamsCustomAuthorization)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EmailDispatcherParamsSendgrid : EmailDispatcherParamsSendgrid struct
type EmailDispatcherParamsSendgrid struct {
	ApiKey *string `json:"apiKey" validate:"required"`
}

// NewEmailDispatcherParamsSendgrid : Instantiate EmailDispatcherParamsSendgrid (Generic Model Constructor)
func (*AppIdManagementV4) NewEmailDispatcherParamsSendgrid(apiKey string) (model *EmailDispatcherParamsSendgrid, err error) {
	model = &EmailDispatcherParamsSendgrid{
		ApiKey: core.StringPtr(apiKey),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalEmailDispatcherParamsSendgrid unmarshals an instance of EmailDispatcherParamsSendgrid from the specified map of raw messages.
func UnmarshalEmailDispatcherParamsSendgrid(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmailDispatcherParamsSendgrid)
	err = core.UnmarshalPrimitive(m, "apiKey", &obj.ApiKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EmailSettingTestOptions : The EmailSettingTest options.
type EmailSettingTestOptions struct {
	EmailTo *string `validate:"required"`

	EmailSettings *EmailSettingsTestParamsEmailSettings `validate:"required"`

	SenderDetails *EmailSettingsTestParamsSenderDetails `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewEmailSettingTestOptions : Instantiate EmailSettingTestOptions
func (*AppIdManagementV4) NewEmailSettingTestOptions(emailTo string, emailSettings *EmailSettingsTestParamsEmailSettings, senderDetails *EmailSettingsTestParamsSenderDetails) *EmailSettingTestOptions {
	return &EmailSettingTestOptions{
		EmailTo: core.StringPtr(emailTo),
		EmailSettings: emailSettings,
		SenderDetails: senderDetails,
	}
}

// SetEmailTo : Allow user to set EmailTo
func (options *EmailSettingTestOptions) SetEmailTo(emailTo string) *EmailSettingTestOptions {
	options.EmailTo = core.StringPtr(emailTo)
	return options
}

// SetEmailSettings : Allow user to set EmailSettings
func (options *EmailSettingTestOptions) SetEmailSettings(emailSettings *EmailSettingsTestParamsEmailSettings) *EmailSettingTestOptions {
	options.EmailSettings = emailSettings
	return options
}

// SetSenderDetails : Allow user to set SenderDetails
func (options *EmailSettingTestOptions) SetSenderDetails(senderDetails *EmailSettingsTestParamsSenderDetails) *EmailSettingTestOptions {
	options.SenderDetails = senderDetails
	return options
}

// SetHeaders : Allow user to set Headers
func (options *EmailSettingTestOptions) SetHeaders(param map[string]string) *EmailSettingTestOptions {
	options.Headers = param
	return options
}

// EmailSettingsTestParamsEmailSettings : EmailSettingsTestParamsEmailSettings struct
type EmailSettingsTestParamsEmailSettings struct {
	Provider *string `json:"provider" validate:"required"`

	Sendgrid *EmailSettingsTestParamsEmailSettingsSendgrid `json:"sendgrid,omitempty"`

	Custom *EmailSettingsTestParamsEmailSettingsCustom `json:"custom,omitempty"`
}

// Constants associated with the EmailSettingsTestParamsEmailSettings.Provider property.
const (
	EmailSettingsTestParamsEmailSettings_Provider_Custom = "custom"
	EmailSettingsTestParamsEmailSettings_Provider_Sendgrid = "sendgrid"
)

// NewEmailSettingsTestParamsEmailSettings : Instantiate EmailSettingsTestParamsEmailSettings (Generic Model Constructor)
func (*AppIdManagementV4) NewEmailSettingsTestParamsEmailSettings(provider string) (model *EmailSettingsTestParamsEmailSettings, err error) {
	model = &EmailSettingsTestParamsEmailSettings{
		Provider: core.StringPtr(provider),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalEmailSettingsTestParamsEmailSettings unmarshals an instance of EmailSettingsTestParamsEmailSettings from the specified map of raw messages.
func UnmarshalEmailSettingsTestParamsEmailSettings(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmailSettingsTestParamsEmailSettings)
	err = core.UnmarshalPrimitive(m, "provider", &obj.Provider)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sendgrid", &obj.Sendgrid, UnmarshalEmailSettingsTestParamsEmailSettingsSendgrid)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "custom", &obj.Custom, UnmarshalEmailSettingsTestParamsEmailSettingsCustom)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EmailSettingsTestParamsEmailSettingsCustom : EmailSettingsTestParamsEmailSettingsCustom struct
type EmailSettingsTestParamsEmailSettingsCustom struct {
	URL *string `json:"url" validate:"required"`

	Authorization *EmailSettingsTestParamsEmailSettingsCustomAuthorization `json:"authorization" validate:"required"`
}

// NewEmailSettingsTestParamsEmailSettingsCustom : Instantiate EmailSettingsTestParamsEmailSettingsCustom (Generic Model Constructor)
func (*AppIdManagementV4) NewEmailSettingsTestParamsEmailSettingsCustom(url string, authorization *EmailSettingsTestParamsEmailSettingsCustomAuthorization) (model *EmailSettingsTestParamsEmailSettingsCustom, err error) {
	model = &EmailSettingsTestParamsEmailSettingsCustom{
		URL: core.StringPtr(url),
		Authorization: authorization,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalEmailSettingsTestParamsEmailSettingsCustom unmarshals an instance of EmailSettingsTestParamsEmailSettingsCustom from the specified map of raw messages.
func UnmarshalEmailSettingsTestParamsEmailSettingsCustom(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmailSettingsTestParamsEmailSettingsCustom)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "authorization", &obj.Authorization, UnmarshalEmailSettingsTestParamsEmailSettingsCustomAuthorization)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EmailSettingsTestParamsEmailSettingsCustomAuthorization : EmailSettingsTestParamsEmailSettingsCustomAuthorization struct
type EmailSettingsTestParamsEmailSettingsCustomAuthorization struct {
	Type *string `json:"type" validate:"required"`

	Value *string `json:"value,omitempty"`

	Username *string `json:"username,omitempty"`

	Password *string `json:"password,omitempty"`
}

// Constants associated with the EmailSettingsTestParamsEmailSettingsCustomAuthorization.Type property.
const (
	EmailSettingsTestParamsEmailSettingsCustomAuthorization_Type_Basic = "basic"
	EmailSettingsTestParamsEmailSettingsCustomAuthorization_Type_None = "none"
	EmailSettingsTestParamsEmailSettingsCustomAuthorization_Type_Value = "value"
)

// NewEmailSettingsTestParamsEmailSettingsCustomAuthorization : Instantiate EmailSettingsTestParamsEmailSettingsCustomAuthorization (Generic Model Constructor)
func (*AppIdManagementV4) NewEmailSettingsTestParamsEmailSettingsCustomAuthorization(typeVar string) (model *EmailSettingsTestParamsEmailSettingsCustomAuthorization, err error) {
	model = &EmailSettingsTestParamsEmailSettingsCustomAuthorization{
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalEmailSettingsTestParamsEmailSettingsCustomAuthorization unmarshals an instance of EmailSettingsTestParamsEmailSettingsCustomAuthorization from the specified map of raw messages.
func UnmarshalEmailSettingsTestParamsEmailSettingsCustomAuthorization(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmailSettingsTestParamsEmailSettingsCustomAuthorization)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "username", &obj.Username)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EmailSettingsTestParamsEmailSettingsSendgrid : EmailSettingsTestParamsEmailSettingsSendgrid struct
type EmailSettingsTestParamsEmailSettingsSendgrid struct {
	ApiKey *string `json:"apiKey" validate:"required"`
}

// NewEmailSettingsTestParamsEmailSettingsSendgrid : Instantiate EmailSettingsTestParamsEmailSettingsSendgrid (Generic Model Constructor)
func (*AppIdManagementV4) NewEmailSettingsTestParamsEmailSettingsSendgrid(apiKey string) (model *EmailSettingsTestParamsEmailSettingsSendgrid, err error) {
	model = &EmailSettingsTestParamsEmailSettingsSendgrid{
		ApiKey: core.StringPtr(apiKey),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalEmailSettingsTestParamsEmailSettingsSendgrid unmarshals an instance of EmailSettingsTestParamsEmailSettingsSendgrid from the specified map of raw messages.
func UnmarshalEmailSettingsTestParamsEmailSettingsSendgrid(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmailSettingsTestParamsEmailSettingsSendgrid)
	err = core.UnmarshalPrimitive(m, "apiKey", &obj.ApiKey)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EmailSettingsTestParamsSenderDetails : EmailSettingsTestParamsSenderDetails struct
type EmailSettingsTestParamsSenderDetails struct {
	From *EmailSettingsTestParamsSenderDetailsFrom `json:"from" validate:"required"`

	ReplyTo *EmailSettingsTestParamsSenderDetailsReplyTo `json:"reply_to,omitempty"`
}

// NewEmailSettingsTestParamsSenderDetails : Instantiate EmailSettingsTestParamsSenderDetails (Generic Model Constructor)
func (*AppIdManagementV4) NewEmailSettingsTestParamsSenderDetails(from *EmailSettingsTestParamsSenderDetailsFrom) (model *EmailSettingsTestParamsSenderDetails, err error) {
	model = &EmailSettingsTestParamsSenderDetails{
		From: from,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalEmailSettingsTestParamsSenderDetails unmarshals an instance of EmailSettingsTestParamsSenderDetails from the specified map of raw messages.
func UnmarshalEmailSettingsTestParamsSenderDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmailSettingsTestParamsSenderDetails)
	err = core.UnmarshalModel(m, "from", &obj.From, UnmarshalEmailSettingsTestParamsSenderDetailsFrom)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "reply_to", &obj.ReplyTo, UnmarshalEmailSettingsTestParamsSenderDetailsReplyTo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EmailSettingsTestParamsSenderDetailsFrom : EmailSettingsTestParamsSenderDetailsFrom struct
type EmailSettingsTestParamsSenderDetailsFrom struct {
	Email *string `json:"email" validate:"required"`

	Name *string `json:"name,omitempty"`
}

// NewEmailSettingsTestParamsSenderDetailsFrom : Instantiate EmailSettingsTestParamsSenderDetailsFrom (Generic Model Constructor)
func (*AppIdManagementV4) NewEmailSettingsTestParamsSenderDetailsFrom(email string) (model *EmailSettingsTestParamsSenderDetailsFrom, err error) {
	model = &EmailSettingsTestParamsSenderDetailsFrom{
		Email: core.StringPtr(email),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalEmailSettingsTestParamsSenderDetailsFrom unmarshals an instance of EmailSettingsTestParamsSenderDetailsFrom from the specified map of raw messages.
func UnmarshalEmailSettingsTestParamsSenderDetailsFrom(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmailSettingsTestParamsSenderDetailsFrom)
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EmailSettingsTestParamsSenderDetailsReplyTo : EmailSettingsTestParamsSenderDetailsReplyTo struct
type EmailSettingsTestParamsSenderDetailsReplyTo struct {
	Email *string `json:"email" validate:"required"`

	Name *string `json:"name,omitempty"`
}

// NewEmailSettingsTestParamsSenderDetailsReplyTo : Instantiate EmailSettingsTestParamsSenderDetailsReplyTo (Generic Model Constructor)
func (*AppIdManagementV4) NewEmailSettingsTestParamsSenderDetailsReplyTo(email string) (model *EmailSettingsTestParamsSenderDetailsReplyTo, err error) {
	model = &EmailSettingsTestParamsSenderDetailsReplyTo{
		Email: core.StringPtr(email),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalEmailSettingsTestParamsSenderDetailsReplyTo unmarshals an instance of EmailSettingsTestParamsSenderDetailsReplyTo from the specified map of raw messages.
func UnmarshalEmailSettingsTestParamsSenderDetailsReplyTo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmailSettingsTestParamsSenderDetailsReplyTo)
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExportUserProfileUsersItem : ExportUserProfileUsersItem struct
type ExportUserProfileUsersItem struct {
	ID *string `json:"id" validate:"required"`

	Identities []ExportUserProfileUsersItemIdentitiesItem `json:"identities" validate:"required"`

	Attributes interface{} `json:"attributes" validate:"required"`

	Name *string `json:"name,omitempty"`

	Email *string `json:"email,omitempty"`

	Picture *string `json:"picture,omitempty"`

	Gender *string `json:"gender,omitempty"`

	Locale *string `json:"locale,omitempty"`

	PreferredUsername *string `json:"preferred_username,omitempty"`

	IDP *string `json:"idp,omitempty"`

	HashedIDPID *string `json:"hashedIdpId,omitempty"`

	HashedEmail *string `json:"hashedEmail,omitempty"`

	Roles []string `json:"roles,omitempty"`
}

// NewExportUserProfileUsersItem : Instantiate ExportUserProfileUsersItem (Generic Model Constructor)
func (*AppIdManagementV4) NewExportUserProfileUsersItem(id string, identities []ExportUserProfileUsersItemIdentitiesItem, attributes interface{}) (model *ExportUserProfileUsersItem, err error) {
	model = &ExportUserProfileUsersItem{
		ID: core.StringPtr(id),
		Identities: identities,
		Attributes: attributes,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalExportUserProfileUsersItem unmarshals an instance of ExportUserProfileUsersItem from the specified map of raw messages.
func UnmarshalExportUserProfileUsersItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExportUserProfileUsersItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "identities", &obj.Identities, UnmarshalExportUserProfileUsersItemIdentitiesItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "attributes", &obj.Attributes)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "picture", &obj.Picture)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "gender", &obj.Gender)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "locale", &obj.Locale)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "preferred_username", &obj.PreferredUsername)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "idp", &obj.IDP)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hashedIdpId", &obj.HashedIDPID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "hashedEmail", &obj.HashedEmail)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "roles", &obj.Roles)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExportUserProfileUsersItemIdentitiesItem : ExportUserProfileUsersItemIdentitiesItem struct
type ExportUserProfileUsersItemIdentitiesItem struct {
	Provider *string `json:"provider,omitempty"`

	ID *string `json:"id,omitempty"`

	IDPUserInfo interface{} `json:"idpUserInfo,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// SetProperty allows the user to set an arbitrary property on an instance of ExportUserProfileUsersItemIdentitiesItem
func (o *ExportUserProfileUsersItemIdentitiesItem) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of ExportUserProfileUsersItemIdentitiesItem
func (o *ExportUserProfileUsersItemIdentitiesItem) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of ExportUserProfileUsersItemIdentitiesItem
func (o *ExportUserProfileUsersItemIdentitiesItem) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of ExportUserProfileUsersItemIdentitiesItem
func (o *ExportUserProfileUsersItemIdentitiesItem) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.Provider != nil {
		m["provider"] = o.Provider
	}
	if o.ID != nil {
		m["id"] = o.ID
	}
	if o.IDPUserInfo != nil {
		m["idpUserInfo"] = o.IDPUserInfo
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalExportUserProfileUsersItemIdentitiesItem unmarshals an instance of ExportUserProfileUsersItemIdentitiesItem from the specified map of raw messages.
func UnmarshalExportUserProfileUsersItemIdentitiesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExportUserProfileUsersItemIdentitiesItem)
	err = core.UnmarshalPrimitive(m, "provider", &obj.Provider)
	if err != nil {
		return
	}
	delete(m, "provider")
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	delete(m, "id")
	err = core.UnmarshalPrimitive(m, "idpUserInfo", &obj.IDPUserInfo)
	if err != nil {
		return
	}
	delete(m, "idpUserInfo")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExportUserUsersItem : ExportUserUsersItem struct
type ExportUserUsersItem struct {
	ScimUser interface{} `json:"scimUser" validate:"required"`

	PasswordHash *string `json:"passwordHash" validate:"required"`

	PasswordHashAlg *string `json:"passwordHashAlg" validate:"required"`

	Profile *ExportUserUsersItemProfile `json:"profile" validate:"required"`

	Roles []string `json:"roles" validate:"required"`
}

// NewExportUserUsersItem : Instantiate ExportUserUsersItem (Generic Model Constructor)
func (*AppIdManagementV4) NewExportUserUsersItem(scimUser interface{}, passwordHash string, passwordHashAlg string, profile *ExportUserUsersItemProfile, roles []string) (model *ExportUserUsersItem, err error) {
	model = &ExportUserUsersItem{
		ScimUser: scimUser,
		PasswordHash: core.StringPtr(passwordHash),
		PasswordHashAlg: core.StringPtr(passwordHashAlg),
		Profile: profile,
		Roles: roles,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalExportUserUsersItem unmarshals an instance of ExportUserUsersItem from the specified map of raw messages.
func UnmarshalExportUserUsersItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExportUserUsersItem)
	err = core.UnmarshalPrimitive(m, "scimUser", &obj.ScimUser)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "passwordHash", &obj.PasswordHash)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "passwordHashAlg", &obj.PasswordHashAlg)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "profile", &obj.Profile, UnmarshalExportUserUsersItemProfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "roles", &obj.Roles)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExportUserUsersItemProfile : ExportUserUsersItemProfile struct
type ExportUserUsersItemProfile struct {
	Attributes interface{} `json:"attributes" validate:"required"`
}

// NewExportUserUsersItemProfile : Instantiate ExportUserUsersItemProfile (Generic Model Constructor)
func (*AppIdManagementV4) NewExportUserUsersItemProfile(attributes interface{}) (model *ExportUserUsersItemProfile, err error) {
	model = &ExportUserUsersItemProfile{
		Attributes: attributes,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalExportUserUsersItemProfile unmarshals an instance of ExportUserUsersItemProfile from the specified map of raw messages.
func UnmarshalExportUserUsersItemProfile(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExportUserUsersItemProfile)
	err = core.UnmarshalPrimitive(m, "attributes", &obj.Attributes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FacebookConfigParamsConfig : FacebookConfigParamsConfig struct
type FacebookConfigParamsConfig struct {
	IDPID *string `json:"idpId" validate:"required"`

	Secret *string `json:"secret" validate:"required"`
}

// UnmarshalFacebookConfigParamsConfig unmarshals an instance of FacebookConfigParamsConfig from the specified map of raw messages.
func UnmarshalFacebookConfigParamsConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FacebookConfigParamsConfig)
	err = core.UnmarshalPrimitive(m, "idpId", &obj.IDPID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "secret", &obj.Secret)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FacebookConfigParamsPUTConfig : FacebookConfigParamsPUTConfig struct
type FacebookConfigParamsPUTConfig struct {
	IDPID *string `json:"idpId" validate:"required"`

	Secret *string `json:"secret" validate:"required"`
}

// UnmarshalFacebookConfigParamsPUTConfig unmarshals an instance of FacebookConfigParamsPUTConfig from the specified map of raw messages.
func UnmarshalFacebookConfigParamsPUTConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FacebookConfigParamsPUTConfig)
	err = core.UnmarshalPrimitive(m, "idpId", &obj.IDPID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "secret", &obj.Secret)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FacebookGoogleConfigParamsConfig : FacebookGoogleConfigParamsConfig struct
type FacebookGoogleConfigParamsConfig struct {
	IDPID *string `json:"idpId" validate:"required"`

	Secret *string `json:"secret" validate:"required"`
}

// NewFacebookGoogleConfigParamsConfig : Instantiate FacebookGoogleConfigParamsConfig (Generic Model Constructor)
func (*AppIdManagementV4) NewFacebookGoogleConfigParamsConfig(iDPID string, secret string) (model *FacebookGoogleConfigParamsConfig, err error) {
	model = &FacebookGoogleConfigParamsConfig{
		IDPID: core.StringPtr(iDPID),
		Secret: core.StringPtr(secret),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalFacebookGoogleConfigParamsConfig unmarshals an instance of FacebookGoogleConfigParamsConfig from the specified map of raw messages.
func UnmarshalFacebookGoogleConfigParamsConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FacebookGoogleConfigParamsConfig)
	err = core.UnmarshalPrimitive(m, "idpId", &obj.IDPID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "secret", &obj.Secret)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ForgotPasswordResultOptions : The ForgotPasswordResult options.
type ForgotPasswordResultOptions struct {
	// The context that will be use to get the forgot password confirmation result.
	Context *string `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewForgotPasswordResultOptions : Instantiate ForgotPasswordResultOptions
func (*AppIdManagementV4) NewForgotPasswordResultOptions(context string) *ForgotPasswordResultOptions {
	return &ForgotPasswordResultOptions{
		Context: core.StringPtr(context),
	}
}

// SetContext : Allow user to set Context
func (options *ForgotPasswordResultOptions) SetContext(context string) *ForgotPasswordResultOptions {
	options.Context = core.StringPtr(context)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ForgotPasswordResultOptions) SetHeaders(param map[string]string) *ForgotPasswordResultOptions {
	options.Headers = param
	return options
}

// GetApplicationOptions : The GetApplication options.
type GetApplicationOptions struct {
	// The application clientId.
	ClientID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetApplicationOptions : Instantiate GetApplicationOptions
func (*AppIdManagementV4) NewGetApplicationOptions(clientID string) *GetApplicationOptions {
	return &GetApplicationOptions{
		ClientID: core.StringPtr(clientID),
	}
}

// SetClientID : Allow user to set ClientID
func (options *GetApplicationOptions) SetClientID(clientID string) *GetApplicationOptions {
	options.ClientID = core.StringPtr(clientID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetApplicationOptions) SetHeaders(param map[string]string) *GetApplicationOptions {
	options.Headers = param
	return options
}

// GetApplicationRolesOptions : The GetApplicationRoles options.
type GetApplicationRolesOptions struct {
	// The application clientId.
	ClientID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetApplicationRolesOptions : Instantiate GetApplicationRolesOptions
func (*AppIdManagementV4) NewGetApplicationRolesOptions(clientID string) *GetApplicationRolesOptions {
	return &GetApplicationRolesOptions{
		ClientID: core.StringPtr(clientID),
	}
}

// SetClientID : Allow user to set ClientID
func (options *GetApplicationRolesOptions) SetClientID(clientID string) *GetApplicationRolesOptions {
	options.ClientID = core.StringPtr(clientID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetApplicationRolesOptions) SetHeaders(param map[string]string) *GetApplicationRolesOptions {
	options.Headers = param
	return options
}

// GetApplicationScopesOptions : The GetApplicationScopes options.
type GetApplicationScopesOptions struct {
	// The application clientId.
	ClientID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetApplicationScopesOptions : Instantiate GetApplicationScopesOptions
func (*AppIdManagementV4) NewGetApplicationScopesOptions(clientID string) *GetApplicationScopesOptions {
	return &GetApplicationScopesOptions{
		ClientID: core.StringPtr(clientID),
	}
}

// SetClientID : Allow user to set ClientID
func (options *GetApplicationScopesOptions) SetClientID(clientID string) *GetApplicationScopesOptions {
	options.ClientID = core.StringPtr(clientID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetApplicationScopesOptions) SetHeaders(param map[string]string) *GetApplicationScopesOptions {
	options.Headers = param
	return options
}

// GetAuditStatusOptions : The GetAuditStatus options.
type GetAuditStatusOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetAuditStatusOptions : Instantiate GetAuditStatusOptions
func (*AppIdManagementV4) NewGetAuditStatusOptions() *GetAuditStatusOptions {
	return &GetAuditStatusOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetAuditStatusOptions) SetHeaders(param map[string]string) *GetAuditStatusOptions {
	options.Headers = param
	return options
}

// GetChannelOptions : The GetChannel options.
type GetChannelOptions struct {
	// The MFA channel.
	Channel *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetChannelOptions.Channel property.
// The MFA channel.
const (
	GetChannelOptions_Channel_Email = "email"
	GetChannelOptions_Channel_Nexmo = "nexmo"
)

// NewGetChannelOptions : Instantiate GetChannelOptions
func (*AppIdManagementV4) NewGetChannelOptions(channel string) *GetChannelOptions {
	return &GetChannelOptions{
		Channel: core.StringPtr(channel),
	}
}

// SetChannel : Allow user to set Channel
func (options *GetChannelOptions) SetChannel(channel string) *GetChannelOptions {
	options.Channel = core.StringPtr(channel)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetChannelOptions) SetHeaders(param map[string]string) *GetChannelOptions {
	options.Headers = param
	return options
}

// GetCloudDirectoryActionUrlOptions : The GetCloudDirectoryActionURL options.
type GetCloudDirectoryActionUrlOptions struct {
	// The type of the action. on_user_verified - the URL of your custom user verified page, on_reset_password - the URL of
	// your custom reset password page.
	Action *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetCloudDirectoryActionUrlOptions.Action property.
// The type of the action. on_user_verified - the URL of your custom user verified page, on_reset_password - the URL of
// your custom reset password page.
const (
	GetCloudDirectoryActionUrlOptions_Action_OnResetPassword = "on_reset_password"
	GetCloudDirectoryActionUrlOptions_Action_OnUserVerified = "on_user_verified"
)

// NewGetCloudDirectoryActionUrlOptions : Instantiate GetCloudDirectoryActionUrlOptions
func (*AppIdManagementV4) NewGetCloudDirectoryActionUrlOptions(action string) *GetCloudDirectoryActionUrlOptions {
	return &GetCloudDirectoryActionUrlOptions{
		Action: core.StringPtr(action),
	}
}

// SetAction : Allow user to set Action
func (options *GetCloudDirectoryActionUrlOptions) SetAction(action string) *GetCloudDirectoryActionUrlOptions {
	options.Action = core.StringPtr(action)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCloudDirectoryActionUrlOptions) SetHeaders(param map[string]string) *GetCloudDirectoryActionUrlOptions {
	options.Headers = param
	return options
}

// GetCloudDirectoryAdvancedPasswordManagementOptions : The GetCloudDirectoryAdvancedPasswordManagement options.
type GetCloudDirectoryAdvancedPasswordManagementOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCloudDirectoryAdvancedPasswordManagementOptions : Instantiate GetCloudDirectoryAdvancedPasswordManagementOptions
func (*AppIdManagementV4) NewGetCloudDirectoryAdvancedPasswordManagementOptions() *GetCloudDirectoryAdvancedPasswordManagementOptions {
	return &GetCloudDirectoryAdvancedPasswordManagementOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetCloudDirectoryAdvancedPasswordManagementOptions) SetHeaders(param map[string]string) *GetCloudDirectoryAdvancedPasswordManagementOptions {
	options.Headers = param
	return options
}

// GetCloudDirectoryEmailDispatcherOptions : The GetCloudDirectoryEmailDispatcher options.
type GetCloudDirectoryEmailDispatcherOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCloudDirectoryEmailDispatcherOptions : Instantiate GetCloudDirectoryEmailDispatcherOptions
func (*AppIdManagementV4) NewGetCloudDirectoryEmailDispatcherOptions() *GetCloudDirectoryEmailDispatcherOptions {
	return &GetCloudDirectoryEmailDispatcherOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetCloudDirectoryEmailDispatcherOptions) SetHeaders(param map[string]string) *GetCloudDirectoryEmailDispatcherOptions {
	options.Headers = param
	return options
}

// GetCloudDirectoryIDPOptions : The GetCloudDirectoryIDP options.
type GetCloudDirectoryIDPOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCloudDirectoryIDPOptions : Instantiate GetCloudDirectoryIDPOptions
func (*AppIdManagementV4) NewGetCloudDirectoryIDPOptions() *GetCloudDirectoryIDPOptions {
	return &GetCloudDirectoryIDPOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetCloudDirectoryIDPOptions) SetHeaders(param map[string]string) *GetCloudDirectoryIDPOptions {
	options.Headers = param
	return options
}

// GetCloudDirectoryPasswordRegexOptions : The GetCloudDirectoryPasswordRegex options.
type GetCloudDirectoryPasswordRegexOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCloudDirectoryPasswordRegexOptions : Instantiate GetCloudDirectoryPasswordRegexOptions
func (*AppIdManagementV4) NewGetCloudDirectoryPasswordRegexOptions() *GetCloudDirectoryPasswordRegexOptions {
	return &GetCloudDirectoryPasswordRegexOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetCloudDirectoryPasswordRegexOptions) SetHeaders(param map[string]string) *GetCloudDirectoryPasswordRegexOptions {
	options.Headers = param
	return options
}

// GetCloudDirectorySenderDetailsOptions : The GetCloudDirectorySenderDetails options.
type GetCloudDirectorySenderDetailsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCloudDirectorySenderDetailsOptions : Instantiate GetCloudDirectorySenderDetailsOptions
func (*AppIdManagementV4) NewGetCloudDirectorySenderDetailsOptions() *GetCloudDirectorySenderDetailsOptions {
	return &GetCloudDirectorySenderDetailsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetCloudDirectorySenderDetailsOptions) SetHeaders(param map[string]string) *GetCloudDirectorySenderDetailsOptions {
	options.Headers = param
	return options
}

// GetCloudDirectoryUserOptions : The GetCloudDirectoryUser options.
type GetCloudDirectoryUserOptions struct {
	// The ID assigned to a user when they sign in by using Cloud Directory.
	UserID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCloudDirectoryUserOptions : Instantiate GetCloudDirectoryUserOptions
func (*AppIdManagementV4) NewGetCloudDirectoryUserOptions(userID string) *GetCloudDirectoryUserOptions {
	return &GetCloudDirectoryUserOptions{
		UserID: core.StringPtr(userID),
	}
}

// SetUserID : Allow user to set UserID
func (options *GetCloudDirectoryUserOptions) SetUserID(userID string) *GetCloudDirectoryUserOptions {
	options.UserID = core.StringPtr(userID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetCloudDirectoryUserOptions) SetHeaders(param map[string]string) *GetCloudDirectoryUserOptions {
	options.Headers = param
	return options
}

// GetCustomIDPOptions : The GetCustomIDP options.
type GetCustomIDPOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCustomIDPOptions : Instantiate GetCustomIDPOptions
func (*AppIdManagementV4) NewGetCustomIDPOptions() *GetCustomIDPOptions {
	return &GetCustomIDPOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetCustomIDPOptions) SetHeaders(param map[string]string) *GetCustomIDPOptions {
	options.Headers = param
	return options
}

// GetExtensionConfigOptions : The GetExtensionConfig options.
type GetExtensionConfigOptions struct {
	// The name of the extension.
	Name *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetExtensionConfigOptions.Name property.
// The name of the extension.
const (
	GetExtensionConfigOptions_Name_Postmfa = "postmfa"
	GetExtensionConfigOptions_Name_Premfa = "premfa"
)

// NewGetExtensionConfigOptions : Instantiate GetExtensionConfigOptions
func (*AppIdManagementV4) NewGetExtensionConfigOptions(name string) *GetExtensionConfigOptions {
	return &GetExtensionConfigOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (options *GetExtensionConfigOptions) SetName(name string) *GetExtensionConfigOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetExtensionConfigOptions) SetHeaders(param map[string]string) *GetExtensionConfigOptions {
	options.Headers = param
	return options
}

// GetFacebookIDPOptions : The GetFacebookIDP options.
type GetFacebookIDPOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetFacebookIDPOptions : Instantiate GetFacebookIDPOptions
func (*AppIdManagementV4) NewGetFacebookIDPOptions() *GetFacebookIDPOptions {
	return &GetFacebookIDPOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetFacebookIDPOptions) SetHeaders(param map[string]string) *GetFacebookIDPOptions {
	options.Headers = param
	return options
}

// GetGoogleIDPOptions : The GetGoogleIDP options.
type GetGoogleIDPOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetGoogleIDPOptions : Instantiate GetGoogleIDPOptions
func (*AppIdManagementV4) NewGetGoogleIDPOptions() *GetGoogleIDPOptions {
	return &GetGoogleIDPOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetGoogleIDPOptions) SetHeaders(param map[string]string) *GetGoogleIDPOptions {
	options.Headers = param
	return options
}

// GetLocalizationOptions : The GetLocalization options.
type GetLocalizationOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetLocalizationOptions : Instantiate GetLocalizationOptions
func (*AppIdManagementV4) NewGetLocalizationOptions() *GetLocalizationOptions {
	return &GetLocalizationOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetLocalizationOptions) SetHeaders(param map[string]string) *GetLocalizationOptions {
	options.Headers = param
	return options
}

// GetMFAConfigOptions : The GetMFAConfig options.
type GetMFAConfigOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetMFAConfigOptions : Instantiate GetMFAConfigOptions
func (*AppIdManagementV4) NewGetMFAConfigOptions() *GetMFAConfigOptions {
	return &GetMFAConfigOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetMFAConfigOptions) SetHeaders(param map[string]string) *GetMFAConfigOptions {
	options.Headers = param
	return options
}

// GetMediaOptions : The GetMedia options.
type GetMediaOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetMediaOptions : Instantiate GetMediaOptions
func (*AppIdManagementV4) NewGetMediaOptions() *GetMediaOptions {
	return &GetMediaOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetMediaOptions) SetHeaders(param map[string]string) *GetMediaOptions {
	options.Headers = param
	return options
}

// GetMediaResponse : GetMediaResponse struct
type GetMediaResponse struct {
	Image *string `json:"image" validate:"required"`
}

// UnmarshalGetMediaResponse unmarshals an instance of GetMediaResponse from the specified map of raw messages.
func UnmarshalGetMediaResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetMediaResponse)
	err = core.UnmarshalPrimitive(m, "image", &obj.Image)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetRateLimitConfigOptions : The GetRateLimitConfig options.
type GetRateLimitConfigOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRateLimitConfigOptions : Instantiate GetRateLimitConfigOptions
func (*AppIdManagementV4) NewGetRateLimitConfigOptions() *GetRateLimitConfigOptions {
	return &GetRateLimitConfigOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetRateLimitConfigOptions) SetHeaders(param map[string]string) *GetRateLimitConfigOptions {
	options.Headers = param
	return options
}

// GetRedirectUrisOptions : The GetRedirectUris options.
type GetRedirectUrisOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRedirectUrisOptions : Instantiate GetRedirectUrisOptions
func (*AppIdManagementV4) NewGetRedirectUrisOptions() *GetRedirectUrisOptions {
	return &GetRedirectUrisOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetRedirectUrisOptions) SetHeaders(param map[string]string) *GetRedirectUrisOptions {
	options.Headers = param
	return options
}

// GetRoleOptions : The GetRole options.
type GetRoleOptions struct {
	// The role identifier.
	RoleID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetRoleOptions : Instantiate GetRoleOptions
func (*AppIdManagementV4) NewGetRoleOptions(roleID string) *GetRoleOptions {
	return &GetRoleOptions{
		RoleID: core.StringPtr(roleID),
	}
}

// SetRoleID : Allow user to set RoleID
func (options *GetRoleOptions) SetRoleID(roleID string) *GetRoleOptions {
	options.RoleID = core.StringPtr(roleID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetRoleOptions) SetHeaders(param map[string]string) *GetRoleOptions {
	options.Headers = param
	return options
}

// GetRoleResponseAccessItem : GetRoleResponseAccessItem struct
type GetRoleResponseAccessItem struct {
	ApplicationID *string `json:"application_id,omitempty"`

	Scopes []string `json:"scopes,omitempty"`
}

// UnmarshalGetRoleResponseAccessItem unmarshals an instance of GetRoleResponseAccessItem from the specified map of raw messages.
func UnmarshalGetRoleResponseAccessItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetRoleResponseAccessItem)
	err = core.UnmarshalPrimitive(m, "application_id", &obj.ApplicationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scopes", &obj.Scopes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetSMSChannelConfig : GetSMSChannelConfig struct
type GetSMSChannelConfig struct {
	Key *string `json:"key,omitempty"`

	Secret *string `json:"secret,omitempty"`

	From *string `json:"from,omitempty"`

	Provider *string `json:"provider,omitempty"`
}

// UnmarshalGetSMSChannelConfig unmarshals an instance of GetSMSChannelConfig from the specified map of raw messages.
func UnmarshalGetSMSChannelConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetSMSChannelConfig)
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "secret", &obj.Secret)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "from", &obj.From)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provider", &obj.Provider)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetSSOConfigOptions : The GetSSOConfig options.
type GetSSOConfigOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSSOConfigOptions : Instantiate GetSSOConfigOptions
func (*AppIdManagementV4) NewGetSSOConfigOptions() *GetSSOConfigOptions {
	return &GetSSOConfigOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetSSOConfigOptions) SetHeaders(param map[string]string) *GetSSOConfigOptions {
	options.Headers = param
	return options
}

// GetSAMLIDPOptions : The GetSAMLIDP options.
type GetSAMLIDPOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSAMLIDPOptions : Instantiate GetSAMLIDPOptions
func (*AppIdManagementV4) NewGetSAMLIDPOptions() *GetSAMLIDPOptions {
	return &GetSAMLIDPOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetSAMLIDPOptions) SetHeaders(param map[string]string) *GetSAMLIDPOptions {
	options.Headers = param
	return options
}

// GetSAMLMetadataOptions : The GetSAMLMetadata options.
type GetSAMLMetadataOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSAMLMetadataOptions : Instantiate GetSAMLMetadataOptions
func (*AppIdManagementV4) NewGetSAMLMetadataOptions() *GetSAMLMetadataOptions {
	return &GetSAMLMetadataOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetSAMLMetadataOptions) SetHeaders(param map[string]string) *GetSAMLMetadataOptions {
	options.Headers = param
	return options
}

// GetTemplateOptions : The GetTemplate options.
type GetTemplateOptions struct {
	// The type of email template. This can be "USER_VERIFICATION", "WELCOME", "PASSWORD_CHANGED", "RESET_PASSWORD" or
	// "MFA_VERIFICATION".
	TemplateName *string `validate:"required,ne="`

	// Preferred language for resource. Format as described at RFC5646. According to the configured languages codes
	// returned from the `GET /management/v4/{tenantId}/config/ui/languages` API.
	Language *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetTemplateOptions.TemplateName property.
// The type of email template. This can be "USER_VERIFICATION", "WELCOME", "PASSWORD_CHANGED", "RESET_PASSWORD" or
// "MFA_VERIFICATION".
const (
	GetTemplateOptions_TemplateName_MfaVerification = "MFA_VERIFICATION"
	GetTemplateOptions_TemplateName_PasswordChanged = "PASSWORD_CHANGED"
	GetTemplateOptions_TemplateName_ResetPassword = "RESET_PASSWORD"
	GetTemplateOptions_TemplateName_UserVerification = "USER_VERIFICATION"
	GetTemplateOptions_TemplateName_Welcome = "WELCOME"
)

// NewGetTemplateOptions : Instantiate GetTemplateOptions
func (*AppIdManagementV4) NewGetTemplateOptions(templateName string, language string) *GetTemplateOptions {
	return &GetTemplateOptions{
		TemplateName: core.StringPtr(templateName),
		Language: core.StringPtr(language),
	}
}

// SetTemplateName : Allow user to set TemplateName
func (options *GetTemplateOptions) SetTemplateName(templateName string) *GetTemplateOptions {
	options.TemplateName = core.StringPtr(templateName)
	return options
}

// SetLanguage : Allow user to set Language
func (options *GetTemplateOptions) SetLanguage(language string) *GetTemplateOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetTemplateOptions) SetHeaders(param map[string]string) *GetTemplateOptions {
	options.Headers = param
	return options
}

// GetThemeColorOptions : The GetThemeColor options.
type GetThemeColorOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetThemeColorOptions : Instantiate GetThemeColorOptions
func (*AppIdManagementV4) NewGetThemeColorOptions() *GetThemeColorOptions {
	return &GetThemeColorOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetThemeColorOptions) SetHeaders(param map[string]string) *GetThemeColorOptions {
	options.Headers = param
	return options
}

// GetThemeColorResponse : GetThemeColorResponse struct
type GetThemeColorResponse struct {
	HeaderColor *string `json:"headerColor" validate:"required"`
}

// UnmarshalGetThemeColorResponse unmarshals an instance of GetThemeColorResponse from the specified map of raw messages.
func UnmarshalGetThemeColorResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetThemeColorResponse)
	err = core.UnmarshalPrimitive(m, "headerColor", &obj.HeaderColor)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetThemeTextOptions : The GetThemeText options.
type GetThemeTextOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetThemeTextOptions : Instantiate GetThemeTextOptions
func (*AppIdManagementV4) NewGetThemeTextOptions() *GetThemeTextOptions {
	return &GetThemeTextOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetThemeTextOptions) SetHeaders(param map[string]string) *GetThemeTextOptions {
	options.Headers = param
	return options
}

// GetThemeTextResponse : GetThemeTextResponse struct
type GetThemeTextResponse struct {
	Footnote *string `json:"footnote" validate:"required"`

	TabTitle *string `json:"tabTitle" validate:"required"`
}

// UnmarshalGetThemeTextResponse unmarshals an instance of GetThemeTextResponse from the specified map of raw messages.
func UnmarshalGetThemeTextResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetThemeTextResponse)
	err = core.UnmarshalPrimitive(m, "footnote", &obj.Footnote)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "tabTitle", &obj.TabTitle)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetTokensConfigOptions : The GetTokensConfig options.
type GetTokensConfigOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetTokensConfigOptions : Instantiate GetTokensConfigOptions
func (*AppIdManagementV4) NewGetTokensConfigOptions() *GetTokensConfigOptions {
	return &GetTokensConfigOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetTokensConfigOptions) SetHeaders(param map[string]string) *GetTokensConfigOptions {
	options.Headers = param
	return options
}

// GetUserAndProfileIdentitiesItem : GetUserAndProfileIdentitiesItem struct
type GetUserAndProfileIdentitiesItem struct {
	Provider *string `json:"provider,omitempty"`

	ID *string `json:"id,omitempty"`

	IDPUserInfo interface{} `json:"idpUserInfo,omitempty"`
}

// UnmarshalGetUserAndProfileIdentitiesItem unmarshals an instance of GetUserAndProfileIdentitiesItem from the specified map of raw messages.
func UnmarshalGetUserAndProfileIdentitiesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetUserAndProfileIdentitiesItem)
	err = core.UnmarshalPrimitive(m, "provider", &obj.Provider)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "idpUserInfo", &obj.IDPUserInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetUserProfilesConfigOptions : The GetUserProfilesConfig options.
type GetUserProfilesConfigOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetUserProfilesConfigOptions : Instantiate GetUserProfilesConfigOptions
func (*AppIdManagementV4) NewGetUserProfilesConfigOptions() *GetUserProfilesConfigOptions {
	return &GetUserProfilesConfigOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *GetUserProfilesConfigOptions) SetHeaders(param map[string]string) *GetUserProfilesConfigOptions {
	options.Headers = param
	return options
}

// GetUserProfilesConfigResponse : GetUserProfilesConfigResponse struct
type GetUserProfilesConfigResponse struct {
	IsActive *bool `json:"isActive" validate:"required"`
}

// UnmarshalGetUserProfilesConfigResponse unmarshals an instance of GetUserProfilesConfigResponse from the specified map of raw messages.
func UnmarshalGetUserProfilesConfigResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetUserProfilesConfigResponse)
	err = core.UnmarshalPrimitive(m, "isActive", &obj.IsActive)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetUserRolesOptions : The GetUserRoles options.
type GetUserRolesOptions struct {
	// The user's identifier ('subject' in identity token) You can search user in <a
	// href="#!/Users/users_search_user_profile" target="_blank">/users</a>.
	ID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetUserRolesOptions : Instantiate GetUserRolesOptions
func (*AppIdManagementV4) NewGetUserRolesOptions(id string) *GetUserRolesOptions {
	return &GetUserRolesOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *GetUserRolesOptions) SetID(id string) *GetUserRolesOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *GetUserRolesOptions) SetHeaders(param map[string]string) *GetUserRolesOptions {
	options.Headers = param
	return options
}

// GetUserRolesResponseRolesItem : GetUserRolesResponseRolesItem struct
type GetUserRolesResponseRolesItem struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`
}

// UnmarshalGetUserRolesResponseRolesItem unmarshals an instance of GetUserRolesResponseRolesItem from the specified map of raw messages.
func UnmarshalGetUserRolesResponseRolesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetUserRolesResponseRolesItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GoogleConfigParamsConfig : GoogleConfigParamsConfig struct
type GoogleConfigParamsConfig struct {
	IDPID *string `json:"idpId" validate:"required"`

	Secret *string `json:"secret" validate:"required"`
}

// UnmarshalGoogleConfigParamsConfig unmarshals an instance of GoogleConfigParamsConfig from the specified map of raw messages.
func UnmarshalGoogleConfigParamsConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GoogleConfigParamsConfig)
	err = core.UnmarshalPrimitive(m, "idpId", &obj.IDPID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "secret", &obj.Secret)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GoogleConfigParamsPUTConfig : GoogleConfigParamsPUTConfig struct
type GoogleConfigParamsPUTConfig struct {
	IDPID *string `json:"idpId" validate:"required"`

	Secret *string `json:"secret" validate:"required"`
}

// UnmarshalGoogleConfigParamsPUTConfig unmarshals an instance of GoogleConfigParamsPUTConfig from the specified map of raw messages.
func UnmarshalGoogleConfigParamsPUTConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GoogleConfigParamsPUTConfig)
	err = core.UnmarshalPrimitive(m, "idpId", &obj.IDPID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "secret", &obj.Secret)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImportProfilesResponseFailReasonsItem : ImportProfilesResponseFailReasonsItem struct
type ImportProfilesResponseFailReasonsItem struct {
	OriginalID *string `json:"originalId,omitempty"`

	IDP *string `json:"idp,omitempty"`

	Error interface{} `json:"error,omitempty"`
}

// UnmarshalImportProfilesResponseFailReasonsItem unmarshals an instance of ImportProfilesResponseFailReasonsItem from the specified map of raw messages.
func UnmarshalImportProfilesResponseFailReasonsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImportProfilesResponseFailReasonsItem)
	err = core.UnmarshalPrimitive(m, "originalId", &obj.OriginalID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "idp", &obj.IDP)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "error", &obj.Error)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImportResponseFailReasonsItem : ImportResponseFailReasonsItem struct
type ImportResponseFailReasonsItem struct {
	OriginalID *string `json:"originalId,omitempty"`

	ID *string `json:"id,omitempty"`

	Email *string `json:"email,omitempty"`

	UserName *string `json:"userName,omitempty"`

	Error interface{} `json:"error,omitempty"`
}

// UnmarshalImportResponseFailReasonsItem unmarshals an instance of ImportResponseFailReasonsItem from the specified map of raw messages.
func UnmarshalImportResponseFailReasonsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImportResponseFailReasonsItem)
	err = core.UnmarshalPrimitive(m, "originalId", &obj.OriginalID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "userName", &obj.UserName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "error", &obj.Error)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListApplicationsOptions : The ListApplications options.
type ListApplicationsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListApplicationsOptions : Instantiate ListApplicationsOptions
func (*AppIdManagementV4) NewListApplicationsOptions() *ListApplicationsOptions {
	return &ListApplicationsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListApplicationsOptions) SetHeaders(param map[string]string) *ListApplicationsOptions {
	options.Headers = param
	return options
}

// ListChannelsOptions : The ListChannels options.
type ListChannelsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListChannelsOptions : Instantiate ListChannelsOptions
func (*AppIdManagementV4) NewListChannelsOptions() *ListChannelsOptions {
	return &ListChannelsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListChannelsOptions) SetHeaders(param map[string]string) *ListChannelsOptions {
	options.Headers = param
	return options
}

// ListCloudDirectoryUsersOptions : The ListCloudDirectoryUsers options.
type ListCloudDirectoryUsersOptions struct {
	// The first result in a set list of results.
	StartIndex *int64

	// The maximum number of results per page.
	Count *int64

	// Filter users by identity field.
	Query *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListCloudDirectoryUsersOptions : Instantiate ListCloudDirectoryUsersOptions
func (*AppIdManagementV4) NewListCloudDirectoryUsersOptions() *ListCloudDirectoryUsersOptions {
	return &ListCloudDirectoryUsersOptions{}
}

// SetStartIndex : Allow user to set StartIndex
func (options *ListCloudDirectoryUsersOptions) SetStartIndex(startIndex int64) *ListCloudDirectoryUsersOptions {
	options.StartIndex = core.Int64Ptr(startIndex)
	return options
}

// SetCount : Allow user to set Count
func (options *ListCloudDirectoryUsersOptions) SetCount(count int64) *ListCloudDirectoryUsersOptions {
	options.Count = core.Int64Ptr(count)
	return options
}

// SetQuery : Allow user to set Query
func (options *ListCloudDirectoryUsersOptions) SetQuery(query string) *ListCloudDirectoryUsersOptions {
	options.Query = core.StringPtr(query)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ListCloudDirectoryUsersOptions) SetHeaders(param map[string]string) *ListCloudDirectoryUsersOptions {
	options.Headers = param
	return options
}

// ListRolesOptions : The ListRoles options.
type ListRolesOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListRolesOptions : Instantiate ListRolesOptions
func (*AppIdManagementV4) NewListRolesOptions() *ListRolesOptions {
	return &ListRolesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListRolesOptions) SetHeaders(param map[string]string) *ListRolesOptions {
	options.Headers = param
	return options
}

// MfaChannelsList : MfaChannelsList struct
type MfaChannelsList struct {
	Channels []MfaChannelsListChannelsItem `json:"channels" validate:"required"`
}

// UnmarshalMfaChannelsList unmarshals an instance of MfaChannelsList from the specified map of raw messages.
func UnmarshalMfaChannelsList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MfaChannelsList)
	err = core.UnmarshalModel(m, "channels", &obj.Channels, UnmarshalMfaChannelsListChannelsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MfaChannelsListChannelsItem : MfaChannelsListChannelsItem struct
type MfaChannelsListChannelsItem struct {
	Type *string `json:"type" validate:"required"`

	IsActive *bool `json:"isActive" validate:"required"`

	Config *MfaChannelsListChannelsItemConfig `json:"config,omitempty"`
}

// UnmarshalMfaChannelsListChannelsItem unmarshals an instance of MfaChannelsListChannelsItem from the specified map of raw messages.
func UnmarshalMfaChannelsListChannelsItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MfaChannelsListChannelsItem)
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "isActive", &obj.IsActive)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalMfaChannelsListChannelsItemConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// MfaChannelsListChannelsItemConfig : MfaChannelsListChannelsItemConfig struct
type MfaChannelsListChannelsItemConfig struct {
	Key *string `json:"key,omitempty"`

	Secret *string `json:"secret,omitempty"`

	From *string `json:"from,omitempty"`

	Provider *string `json:"provider,omitempty"`
}

// UnmarshalMfaChannelsListChannelsItemConfig unmarshals an instance of MfaChannelsListChannelsItemConfig from the specified map of raw messages.
func UnmarshalMfaChannelsListChannelsItemConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(MfaChannelsListChannelsItemConfig)
	err = core.UnmarshalPrimitive(m, "key", &obj.Key)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "secret", &obj.Secret)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "from", &obj.From)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "provider", &obj.Provider)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PostEmailDispatcherTestOptions : The PostEmailDispatcherTest options.
type PostEmailDispatcherTestOptions struct {
	// The email address where you want to send your test message.
	Email *string `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPostEmailDispatcherTestOptions : Instantiate PostEmailDispatcherTestOptions
func (*AppIdManagementV4) NewPostEmailDispatcherTestOptions(email string) *PostEmailDispatcherTestOptions {
	return &PostEmailDispatcherTestOptions{
		Email: core.StringPtr(email),
	}
}

// SetEmail : Allow user to set Email
func (options *PostEmailDispatcherTestOptions) SetEmail(email string) *PostEmailDispatcherTestOptions {
	options.Email = core.StringPtr(email)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PostEmailDispatcherTestOptions) SetHeaders(param map[string]string) *PostEmailDispatcherTestOptions {
	options.Headers = param
	return options
}

// PostExtensionsTestOptions : The PostExtensionsTest options.
type PostExtensionsTestOptions struct {
	// The name of the extension.
	Name *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the PostExtensionsTestOptions.Name property.
// The name of the extension.
const (
	PostExtensionsTestOptions_Name_Postmfa = "postmfa"
	PostExtensionsTestOptions_Name_Premfa = "premfa"
)

// NewPostExtensionsTestOptions : Instantiate PostExtensionsTestOptions
func (*AppIdManagementV4) NewPostExtensionsTestOptions(name string) *PostExtensionsTestOptions {
	return &PostExtensionsTestOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (options *PostExtensionsTestOptions) SetName(name string) *PostExtensionsTestOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PostExtensionsTestOptions) SetHeaders(param map[string]string) *PostExtensionsTestOptions {
	options.Headers = param
	return options
}

// PostMediaOptions : The PostMedia options.
type PostMediaOptions struct {
	// The type of media. You can upload JPG or PNG files.
	MediaType *string `validate:"required"`

	// The image file. The recommended size is 320x320 px. The maxmimum files size is 100kb.
	File io.ReadCloser `validate:"required"`

	// The content type of file.
	FileContentType *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the PostMediaOptions.MediaType property.
// The type of media. You can upload JPG or PNG files.
const (
	PostMediaOptions_MediaType_Logo = "logo"
)

// NewPostMediaOptions : Instantiate PostMediaOptions
func (*AppIdManagementV4) NewPostMediaOptions(mediaType string, file io.ReadCloser) *PostMediaOptions {
	return &PostMediaOptions{
		MediaType: core.StringPtr(mediaType),
		File: file,
	}
}

// SetMediaType : Allow user to set MediaType
func (options *PostMediaOptions) SetMediaType(mediaType string) *PostMediaOptions {
	options.MediaType = core.StringPtr(mediaType)
	return options
}

// SetFile : Allow user to set File
func (options *PostMediaOptions) SetFile(file io.ReadCloser) *PostMediaOptions {
	options.File = file
	return options
}

// SetFileContentType : Allow user to set FileContentType
func (options *PostMediaOptions) SetFileContentType(fileContentType string) *PostMediaOptions {
	options.FileContentType = core.StringPtr(fileContentType)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PostMediaOptions) SetHeaders(param map[string]string) *PostMediaOptions {
	options.Headers = param
	return options
}

// PostSmsDispatcherTestOptions : The PostSmsDispatcherTest options.
type PostSmsDispatcherTestOptions struct {
	PhoneNumber *string `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPostSmsDispatcherTestOptions : Instantiate PostSmsDispatcherTestOptions
func (*AppIdManagementV4) NewPostSmsDispatcherTestOptions(phoneNumber string) *PostSmsDispatcherTestOptions {
	return &PostSmsDispatcherTestOptions{
		PhoneNumber: core.StringPtr(phoneNumber),
	}
}

// SetPhoneNumber : Allow user to set PhoneNumber
func (options *PostSmsDispatcherTestOptions) SetPhoneNumber(phoneNumber string) *PostSmsDispatcherTestOptions {
	options.PhoneNumber = core.StringPtr(phoneNumber)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PostSmsDispatcherTestOptions) SetHeaders(param map[string]string) *PostSmsDispatcherTestOptions {
	options.Headers = param
	return options
}

// PostThemeColorOptions : The PostThemeColor options.
type PostThemeColorOptions struct {
	HeaderColor *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPostThemeColorOptions : Instantiate PostThemeColorOptions
func (*AppIdManagementV4) NewPostThemeColorOptions() *PostThemeColorOptions {
	return &PostThemeColorOptions{}
}

// SetHeaderColor : Allow user to set HeaderColor
func (options *PostThemeColorOptions) SetHeaderColor(headerColor string) *PostThemeColorOptions {
	options.HeaderColor = core.StringPtr(headerColor)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PostThemeColorOptions) SetHeaders(param map[string]string) *PostThemeColorOptions {
	options.Headers = param
	return options
}

// PostThemeTextOptions : The PostThemeText options.
type PostThemeTextOptions struct {
	TabTitle *string

	Footnote *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPostThemeTextOptions : Instantiate PostThemeTextOptions
func (*AppIdManagementV4) NewPostThemeTextOptions() *PostThemeTextOptions {
	return &PostThemeTextOptions{}
}

// SetTabTitle : Allow user to set TabTitle
func (options *PostThemeTextOptions) SetTabTitle(tabTitle string) *PostThemeTextOptions {
	options.TabTitle = core.StringPtr(tabTitle)
	return options
}

// SetFootnote : Allow user to set Footnote
func (options *PostThemeTextOptions) SetFootnote(footnote string) *PostThemeTextOptions {
	options.Footnote = core.StringPtr(footnote)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PostThemeTextOptions) SetHeaders(param map[string]string) *PostThemeTextOptions {
	options.Headers = param
	return options
}

// PutApplicationsRolesOptions : The PutApplicationsRoles options.
type PutApplicationsRolesOptions struct {
	// The application clientId.
	ClientID *string `validate:"required,ne="`

	Roles *UpdateUserRolesParamsRoles `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPutApplicationsRolesOptions : Instantiate PutApplicationsRolesOptions
func (*AppIdManagementV4) NewPutApplicationsRolesOptions(clientID string, roles *UpdateUserRolesParamsRoles) *PutApplicationsRolesOptions {
	return &PutApplicationsRolesOptions{
		ClientID: core.StringPtr(clientID),
		Roles: roles,
	}
}

// SetClientID : Allow user to set ClientID
func (options *PutApplicationsRolesOptions) SetClientID(clientID string) *PutApplicationsRolesOptions {
	options.ClientID = core.StringPtr(clientID)
	return options
}

// SetRoles : Allow user to set Roles
func (options *PutApplicationsRolesOptions) SetRoles(roles *UpdateUserRolesParamsRoles) *PutApplicationsRolesOptions {
	options.Roles = roles
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PutApplicationsRolesOptions) SetHeaders(param map[string]string) *PutApplicationsRolesOptions {
	options.Headers = param
	return options
}

// PutApplicationsScopesOptions : The PutApplicationsScopes options.
type PutApplicationsScopesOptions struct {
	// The application clientId.
	ClientID *string `validate:"required,ne="`

	Scopes []string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPutApplicationsScopesOptions : Instantiate PutApplicationsScopesOptions
func (*AppIdManagementV4) NewPutApplicationsScopesOptions(clientID string) *PutApplicationsScopesOptions {
	return &PutApplicationsScopesOptions{
		ClientID: core.StringPtr(clientID),
	}
}

// SetClientID : Allow user to set ClientID
func (options *PutApplicationsScopesOptions) SetClientID(clientID string) *PutApplicationsScopesOptions {
	options.ClientID = core.StringPtr(clientID)
	return options
}

// SetScopes : Allow user to set Scopes
func (options *PutApplicationsScopesOptions) SetScopes(scopes []string) *PutApplicationsScopesOptions {
	options.Scopes = scopes
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PutApplicationsScopesOptions) SetHeaders(param map[string]string) *PutApplicationsScopesOptions {
	options.Headers = param
	return options
}

// PutTokensConfigOptions : The PutTokensConfig options.
type PutTokensConfigOptions struct {
	IdTokenClaims []TokenClaimMapping

	AccessTokenClaims []TokenClaimMapping

	Access []TokenConfigParams

	Refresh []RefreshTokenConfigParams

	AnonymousAccess []TokenConfigParams

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewPutTokensConfigOptions : Instantiate PutTokensConfigOptions
func (*AppIdManagementV4) NewPutTokensConfigOptions() *PutTokensConfigOptions {
	return &PutTokensConfigOptions{}
}

// SetIdTokenClaims : Allow user to set IdTokenClaims
func (options *PutTokensConfigOptions) SetIdTokenClaims(idTokenClaims []TokenClaimMapping) *PutTokensConfigOptions {
	options.IdTokenClaims = idTokenClaims
	return options
}

// SetAccessTokenClaims : Allow user to set AccessTokenClaims
func (options *PutTokensConfigOptions) SetAccessTokenClaims(accessTokenClaims []TokenClaimMapping) *PutTokensConfigOptions {
	options.AccessTokenClaims = accessTokenClaims
	return options
}

// SetAccess : Allow user to set Access
func (options *PutTokensConfigOptions) SetAccess(access []TokenConfigParams) *PutTokensConfigOptions {
	options.Access = access
	return options
}

// SetRefresh : Allow user to set Refresh
func (options *PutTokensConfigOptions) SetRefresh(refresh []RefreshTokenConfigParams) *PutTokensConfigOptions {
	options.Refresh = refresh
	return options
}

// SetAnonymousAccess : Allow user to set AnonymousAccess
func (options *PutTokensConfigOptions) SetAnonymousAccess(anonymousAccess []TokenConfigParams) *PutTokensConfigOptions {
	options.AnonymousAccess = anonymousAccess
	return options
}

// SetHeaders : Allow user to set Headers
func (options *PutTokensConfigOptions) SetHeaders(param map[string]string) *PutTokensConfigOptions {
	options.Headers = param
	return options
}

// RegisterApplicationOptions : The RegisterApplication options.
type RegisterApplicationOptions struct {
	// The application name to be registered. Application name cannot exceed 50 characters.
	Name *string `validate:"required"`

	// The type of application to be registered. Allowed types are regularwebapp and singlepageapp.
	Type *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewRegisterApplicationOptions : Instantiate RegisterApplicationOptions
func (*AppIdManagementV4) NewRegisterApplicationOptions(name string) *RegisterApplicationOptions {
	return &RegisterApplicationOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (options *RegisterApplicationOptions) SetName(name string) *RegisterApplicationOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetType : Allow user to set Type
func (options *RegisterApplicationOptions) SetType(typeVar string) *RegisterApplicationOptions {
	options.Type = core.StringPtr(typeVar)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *RegisterApplicationOptions) SetHeaders(param map[string]string) *RegisterApplicationOptions {
	options.Headers = param
	return options
}

// ResendNotificationOptions : The ResendNotification options.
type ResendNotificationOptions struct {
	// The type of email template. This can be "USER_VERIFICATION", "WELCOME", "PASSWORD_CHANGED" or "RESET_PASSWORD".
	TemplateName *string `validate:"required,ne="`

	// The Cloud Directory unique user Id.
	UUID *string `validate:"required"`

	// Preferred language for resource. Format as described at RFC5646.
	Language *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ResendNotificationOptions.TemplateName property.
// The type of email template. This can be "USER_VERIFICATION", "WELCOME", "PASSWORD_CHANGED" or "RESET_PASSWORD".
const (
	ResendNotificationOptions_TemplateName_PasswordChanged = "PASSWORD_CHANGED"
	ResendNotificationOptions_TemplateName_ResetPassword = "RESET_PASSWORD"
	ResendNotificationOptions_TemplateName_UserVerification = "USER_VERIFICATION"
	ResendNotificationOptions_TemplateName_Welcome = "WELCOME"
)

// NewResendNotificationOptions : Instantiate ResendNotificationOptions
func (*AppIdManagementV4) NewResendNotificationOptions(templateName string, uuid string) *ResendNotificationOptions {
	return &ResendNotificationOptions{
		TemplateName: core.StringPtr(templateName),
		UUID: core.StringPtr(uuid),
	}
}

// SetTemplateName : Allow user to set TemplateName
func (options *ResendNotificationOptions) SetTemplateName(templateName string) *ResendNotificationOptions {
	options.TemplateName = core.StringPtr(templateName)
	return options
}

// SetUUID : Allow user to set UUID
func (options *ResendNotificationOptions) SetUUID(uuid string) *ResendNotificationOptions {
	options.UUID = core.StringPtr(uuid)
	return options
}

// SetLanguage : Allow user to set Language
func (options *ResendNotificationOptions) SetLanguage(language string) *ResendNotificationOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *ResendNotificationOptions) SetHeaders(param map[string]string) *ResendNotificationOptions {
	options.Headers = param
	return options
}

// ResendNotificationResponse : ResendNotificationResponse struct
type ResendNotificationResponse struct {
	Message *string `json:"message,omitempty"`
}

// UnmarshalResendNotificationResponse unmarshals an instance of ResendNotificationResponse from the specified map of raw messages.
func UnmarshalResendNotificationResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ResendNotificationResponse)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RolesList : RolesList struct
type RolesList struct {
	Roles []RolesListRolesItem `json:"roles,omitempty"`
}

// UnmarshalRolesList unmarshals an instance of RolesList from the specified map of raw messages.
func UnmarshalRolesList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RolesList)
	err = core.UnmarshalModel(m, "roles", &obj.Roles, UnmarshalRolesListRolesItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RolesListRolesItem : RolesListRolesItem struct
type RolesListRolesItem struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Access []RolesListRolesItemAccessItem `json:"access,omitempty"`
}

// UnmarshalRolesListRolesItem unmarshals an instance of RolesListRolesItem from the specified map of raw messages.
func UnmarshalRolesListRolesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RolesListRolesItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "access", &obj.Access, UnmarshalRolesListRolesItemAccessItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RolesListRolesItemAccessItem : RolesListRolesItemAccessItem struct
type RolesListRolesItemAccessItem struct {
	ApplicationID *string `json:"application_id,omitempty"`

	Scopes []string `json:"scopes,omitempty"`
}

// UnmarshalRolesListRolesItemAccessItem unmarshals an instance of RolesListRolesItemAccessItem from the specified map of raw messages.
func UnmarshalRolesListRolesItemAccessItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RolesListRolesItemAccessItem)
	err = core.UnmarshalPrimitive(m, "application_id", &obj.ApplicationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scopes", &obj.Scopes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SAMLConfigParamsAuthnContext : SAMLConfigParamsAuthnContext struct
type SAMLConfigParamsAuthnContext struct {
	Class []string `json:"class,omitempty"`

	Comparison *string `json:"comparison,omitempty"`
}

// Constants associated with the SAMLConfigParamsAuthnContext.Class property.
const (
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesAuthenticatedtelephony = "urn:oasis:names:tc:SAML:2.0:ac:classes:AuthenticatedTelephony"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesInternetprotocol = "urn:oasis:names:tc:SAML:2.0:ac:classes:InternetProtocol"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesInternetprotocolpassword = "urn:oasis:names:tc:SAML:2.0:ac:classes:InternetProtocolPassword"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesKerberos = "urn:oasis:names:tc:SAML:2.0:ac:classes:Kerberos"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesMobileonefactorcontract = "urn:oasis:names:tc:SAML:2.0:ac:classes:MobileOneFactorContract"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesMobileonefactorunregistered = "urn:oasis:names:tc:SAML:2.0:ac:classes:MobileOneFactorUnregistered"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesMobiletwofactorcontract = "urn:oasis:names:tc:SAML:2.0:ac:classes:MobileTwoFactorContract"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesMobiletwofactorunregistered = "urn:oasis:names:tc:SAML:2.0:ac:classes:MobileTwoFactorUnregistered"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesNomadtelephony = "urn:oasis:names:tc:SAML:2.0:ac:classes:NomadTelephony"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesPassword = "urn:oasis:names:tc:SAML:2.0:ac:classes:Password"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesPasswordprotectedtransport = "urn:oasis:names:tc:SAML:2.0:ac:classes:PasswordProtectedTransport"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesPersonaltelephony = "urn:oasis:names:tc:SAML:2.0:ac:classes:PersonalTelephony"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesPgp = "urn:oasis:names:tc:SAML:2.0:ac:classes:PGP"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesPrevioussession = "urn:oasis:names:tc:SAML:2.0:ac:classes:PreviousSession"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesSecureremotepassword = "urn:oasis:names:tc:SAML:2.0:ac:classes:SecureRemotePassword"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesSmartcard = "urn:oasis:names:tc:SAML:2.0:ac:classes:Smartcard"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesSmartcardpki = "urn:oasis:names:tc:SAML:2.0:ac:classes:SmartcardPKI"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesSoftwarepki = "urn:oasis:names:tc:SAML:2.0:ac:classes:SoftwarePKI"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesSpki = "urn:oasis:names:tc:SAML:2.0:ac:classes:SPKI"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesTelephony = "urn:oasis:names:tc:SAML:2.0:ac:classes:Telephony"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesTimesynctoken = "urn:oasis:names:tc:SAML:2.0:ac:classes:TimeSyncToken"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesTlsclient = "urn:oasis:names:tc:SAML:2.0:ac:classes:TLSClient"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesUnspecified = "urn:oasis:names:tc:SAML:2.0:ac:classes:unspecified"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesX509 = "urn:oasis:names:tc:SAML:2.0:ac:classes:X509"
	SAMLConfigParamsAuthnContext_Class_UrnOasisNamesTcSAML20AcClassesXmldsig = "urn:oasis:names:tc:SAML:2.0:ac:classes:XMLDSig"
)

// Constants associated with the SAMLConfigParamsAuthnContext.Comparison property.
const (
	SAMLConfigParamsAuthnContext_Comparison_Better = "better"
	SAMLConfigParamsAuthnContext_Comparison_Exact = "exact"
	SAMLConfigParamsAuthnContext_Comparison_Maximum = "maximum"
	SAMLConfigParamsAuthnContext_Comparison_Minimum = "minimum"
)

// UnmarshalSAMLConfigParamsAuthnContext unmarshals an instance of SAMLConfigParamsAuthnContext from the specified map of raw messages.
func UnmarshalSAMLConfigParamsAuthnContext(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SAMLConfigParamsAuthnContext)
	err = core.UnmarshalPrimitive(m, "class", &obj.Class)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "comparison", &obj.Comparison)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SAMLResponseWithValidationDataValidationData : SAMLResponseWithValidationDataValidationData struct
type SAMLResponseWithValidationDataValidationData struct {
	Certificates []SAMLResponseWithValidationDataValidationDataCertificatesItem `json:"certificates" validate:"required"`
}

// UnmarshalSAMLResponseWithValidationDataValidationData unmarshals an instance of SAMLResponseWithValidationDataValidationData from the specified map of raw messages.
func UnmarshalSAMLResponseWithValidationDataValidationData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SAMLResponseWithValidationDataValidationData)
	err = core.UnmarshalModel(m, "certificates", &obj.Certificates, UnmarshalSAMLResponseWithValidationDataValidationDataCertificatesItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SAMLResponseWithValidationDataValidationDataCertificatesItem : SAMLResponseWithValidationDataValidationDataCertificatesItem struct
type SAMLResponseWithValidationDataValidationDataCertificatesItem struct {
	CertificateIndex *float64 `json:"certificate_index" validate:"required"`

	ExpirationTimestamp *float64 `json:"expiration_timestamp" validate:"required"`

	Warning *string `json:"warning,omitempty"`
}

// UnmarshalSAMLResponseWithValidationDataValidationDataCertificatesItem unmarshals an instance of SAMLResponseWithValidationDataValidationDataCertificatesItem from the specified map of raw messages.
func UnmarshalSAMLResponseWithValidationDataValidationDataCertificatesItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SAMLResponseWithValidationDataValidationDataCertificatesItem)
	err = core.UnmarshalPrimitive(m, "certificate_index", &obj.CertificateIndex)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "expiration_timestamp", &obj.ExpirationTimestamp)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "warning", &obj.Warning)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SetAuditStatusOptions : The SetAuditStatus options.
type SetAuditStatusOptions struct {
	IsActive *bool `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetAuditStatusOptions : Instantiate SetAuditStatusOptions
func (*AppIdManagementV4) NewSetAuditStatusOptions(isActive bool) *SetAuditStatusOptions {
	return &SetAuditStatusOptions{
		IsActive: core.BoolPtr(isActive),
	}
}

// SetIsActive : Allow user to set IsActive
func (options *SetAuditStatusOptions) SetIsActive(isActive bool) *SetAuditStatusOptions {
	options.IsActive = core.BoolPtr(isActive)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetAuditStatusOptions) SetHeaders(param map[string]string) *SetAuditStatusOptions {
	options.Headers = param
	return options
}

// SetCloudDirectoryActionOptions : The SetCloudDirectoryAction options.
type SetCloudDirectoryActionOptions struct {
	// The type of the action. on_user_verified - the URL of your custom user verified page, on_reset_password - the URL of
	// your custom reset password page.
	Action *string `validate:"required,ne="`

	// The action URL.
	ActionURL *string `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the SetCloudDirectoryActionOptions.Action property.
// The type of the action. on_user_verified - the URL of your custom user verified page, on_reset_password - the URL of
// your custom reset password page.
const (
	SetCloudDirectoryActionOptions_Action_OnResetPassword = "on_reset_password"
	SetCloudDirectoryActionOptions_Action_OnUserVerified = "on_user_verified"
)

// NewSetCloudDirectoryActionOptions : Instantiate SetCloudDirectoryActionOptions
func (*AppIdManagementV4) NewSetCloudDirectoryActionOptions(action string, actionURL string) *SetCloudDirectoryActionOptions {
	return &SetCloudDirectoryActionOptions{
		Action: core.StringPtr(action),
		ActionURL: core.StringPtr(actionURL),
	}
}

// SetAction : Allow user to set Action
func (options *SetCloudDirectoryActionOptions) SetAction(action string) *SetCloudDirectoryActionOptions {
	options.Action = core.StringPtr(action)
	return options
}

// SetActionURL : Allow user to set ActionURL
func (options *SetCloudDirectoryActionOptions) SetActionURL(actionURL string) *SetCloudDirectoryActionOptions {
	options.ActionURL = core.StringPtr(actionURL)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetCloudDirectoryActionOptions) SetHeaders(param map[string]string) *SetCloudDirectoryActionOptions {
	options.Headers = param
	return options
}

// SetCloudDirectoryAdvancedPasswordManagementOptions : The SetCloudDirectoryAdvancedPasswordManagement options.
type SetCloudDirectoryAdvancedPasswordManagementOptions struct {
	AdvancedPasswordManagement *ApmSchemaAdvancedPasswordManagement `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetCloudDirectoryAdvancedPasswordManagementOptions : Instantiate SetCloudDirectoryAdvancedPasswordManagementOptions
func (*AppIdManagementV4) NewSetCloudDirectoryAdvancedPasswordManagementOptions(advancedPasswordManagement *ApmSchemaAdvancedPasswordManagement) *SetCloudDirectoryAdvancedPasswordManagementOptions {
	return &SetCloudDirectoryAdvancedPasswordManagementOptions{
		AdvancedPasswordManagement: advancedPasswordManagement,
	}
}

// SetAdvancedPasswordManagement : Allow user to set AdvancedPasswordManagement
func (options *SetCloudDirectoryAdvancedPasswordManagementOptions) SetAdvancedPasswordManagement(advancedPasswordManagement *ApmSchemaAdvancedPasswordManagement) *SetCloudDirectoryAdvancedPasswordManagementOptions {
	options.AdvancedPasswordManagement = advancedPasswordManagement
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetCloudDirectoryAdvancedPasswordManagementOptions) SetHeaders(param map[string]string) *SetCloudDirectoryAdvancedPasswordManagementOptions {
	options.Headers = param
	return options
}

// SetCloudDirectoryEmailDispatcherOptions : The SetCloudDirectoryEmailDispatcher options.
type SetCloudDirectoryEmailDispatcherOptions struct {
	Provider *string `validate:"required"`

	Sendgrid *EmailDispatcherParamsSendgrid

	Custom *EmailDispatcherParamsCustom

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the SetCloudDirectoryEmailDispatcherOptions.Provider property.
const (
	SetCloudDirectoryEmailDispatcherOptions_Provider_Appid = "appid"
	SetCloudDirectoryEmailDispatcherOptions_Provider_Custom = "custom"
	SetCloudDirectoryEmailDispatcherOptions_Provider_Sendgrid = "sendgrid"
)

// NewSetCloudDirectoryEmailDispatcherOptions : Instantiate SetCloudDirectoryEmailDispatcherOptions
func (*AppIdManagementV4) NewSetCloudDirectoryEmailDispatcherOptions(provider string) *SetCloudDirectoryEmailDispatcherOptions {
	return &SetCloudDirectoryEmailDispatcherOptions{
		Provider: core.StringPtr(provider),
	}
}

// SetProvider : Allow user to set Provider
func (options *SetCloudDirectoryEmailDispatcherOptions) SetProvider(provider string) *SetCloudDirectoryEmailDispatcherOptions {
	options.Provider = core.StringPtr(provider)
	return options
}

// SetSendgrid : Allow user to set Sendgrid
func (options *SetCloudDirectoryEmailDispatcherOptions) SetSendgrid(sendgrid *EmailDispatcherParamsSendgrid) *SetCloudDirectoryEmailDispatcherOptions {
	options.Sendgrid = sendgrid
	return options
}

// SetCustom : Allow user to set Custom
func (options *SetCloudDirectoryEmailDispatcherOptions) SetCustom(custom *EmailDispatcherParamsCustom) *SetCloudDirectoryEmailDispatcherOptions {
	options.Custom = custom
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetCloudDirectoryEmailDispatcherOptions) SetHeaders(param map[string]string) *SetCloudDirectoryEmailDispatcherOptions {
	options.Headers = param
	return options
}

// SetCloudDirectoryIDPOptions : The SetCloudDirectoryIDP options.
type SetCloudDirectoryIDPOptions struct {
	IsActive *bool `validate:"required"`

	Config *CloudDirectoryConfigParams `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetCloudDirectoryIDPOptions : Instantiate SetCloudDirectoryIDPOptions
func (*AppIdManagementV4) NewSetCloudDirectoryIDPOptions(isActive bool, config *CloudDirectoryConfigParams) *SetCloudDirectoryIDPOptions {
	return &SetCloudDirectoryIDPOptions{
		IsActive: core.BoolPtr(isActive),
		Config: config,
	}
}

// SetIsActive : Allow user to set IsActive
func (options *SetCloudDirectoryIDPOptions) SetIsActive(isActive bool) *SetCloudDirectoryIDPOptions {
	options.IsActive = core.BoolPtr(isActive)
	return options
}

// SetConfig : Allow user to set Config
func (options *SetCloudDirectoryIDPOptions) SetConfig(config *CloudDirectoryConfigParams) *SetCloudDirectoryIDPOptions {
	options.Config = config
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetCloudDirectoryIDPOptions) SetHeaders(param map[string]string) *SetCloudDirectoryIDPOptions {
	options.Headers = param
	return options
}

// SetCloudDirectoryPasswordRegexOptions : The SetCloudDirectoryPasswordRegex options.
type SetCloudDirectoryPasswordRegexOptions struct {
	Regex *string

	Base64EncodedRegex *string

	ErrorMessage *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetCloudDirectoryPasswordRegexOptions : Instantiate SetCloudDirectoryPasswordRegexOptions
func (*AppIdManagementV4) NewSetCloudDirectoryPasswordRegexOptions() *SetCloudDirectoryPasswordRegexOptions {
	return &SetCloudDirectoryPasswordRegexOptions{}
}

// SetRegex : Allow user to set Regex
func (options *SetCloudDirectoryPasswordRegexOptions) SetRegex(regex string) *SetCloudDirectoryPasswordRegexOptions {
	options.Regex = core.StringPtr(regex)
	return options
}

// SetBase64EncodedRegex : Allow user to set Base64EncodedRegex
func (options *SetCloudDirectoryPasswordRegexOptions) SetBase64EncodedRegex(base64EncodedRegex string) *SetCloudDirectoryPasswordRegexOptions {
	options.Base64EncodedRegex = core.StringPtr(base64EncodedRegex)
	return options
}

// SetErrorMessage : Allow user to set ErrorMessage
func (options *SetCloudDirectoryPasswordRegexOptions) SetErrorMessage(errorMessage string) *SetCloudDirectoryPasswordRegexOptions {
	options.ErrorMessage = core.StringPtr(errorMessage)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetCloudDirectoryPasswordRegexOptions) SetHeaders(param map[string]string) *SetCloudDirectoryPasswordRegexOptions {
	options.Headers = param
	return options
}

// SetCloudDirectorySenderDetailsOptions : The SetCloudDirectorySenderDetails options.
type SetCloudDirectorySenderDetailsOptions struct {
	SenderDetails *CloudDirectorySenderDetailsSenderDetails `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetCloudDirectorySenderDetailsOptions : Instantiate SetCloudDirectorySenderDetailsOptions
func (*AppIdManagementV4) NewSetCloudDirectorySenderDetailsOptions(senderDetails *CloudDirectorySenderDetailsSenderDetails) *SetCloudDirectorySenderDetailsOptions {
	return &SetCloudDirectorySenderDetailsOptions{
		SenderDetails: senderDetails,
	}
}

// SetSenderDetails : Allow user to set SenderDetails
func (options *SetCloudDirectorySenderDetailsOptions) SetSenderDetails(senderDetails *CloudDirectorySenderDetailsSenderDetails) *SetCloudDirectorySenderDetailsOptions {
	options.SenderDetails = senderDetails
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetCloudDirectorySenderDetailsOptions) SetHeaders(param map[string]string) *SetCloudDirectorySenderDetailsOptions {
	options.Headers = param
	return options
}

// SetCustomIDPOptions : The SetCustomIDP options.
type SetCustomIDPOptions struct {
	IsActive *bool `validate:"required"`

	Config *CustomIdPConfigParamsConfig

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetCustomIDPOptions : Instantiate SetCustomIDPOptions
func (*AppIdManagementV4) NewSetCustomIDPOptions(isActive bool) *SetCustomIDPOptions {
	return &SetCustomIDPOptions{
		IsActive: core.BoolPtr(isActive),
	}
}

// SetIsActive : Allow user to set IsActive
func (options *SetCustomIDPOptions) SetIsActive(isActive bool) *SetCustomIDPOptions {
	options.IsActive = core.BoolPtr(isActive)
	return options
}

// SetConfig : Allow user to set Config
func (options *SetCustomIDPOptions) SetConfig(config *CustomIdPConfigParamsConfig) *SetCustomIDPOptions {
	options.Config = config
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetCustomIDPOptions) SetHeaders(param map[string]string) *SetCustomIDPOptions {
	options.Headers = param
	return options
}

// SetFacebookIDPOptions : The SetFacebookIDP options.
type SetFacebookIDPOptions struct {
	// The identity provider configuration as a JSON object. If the configuration is not set, IBM default credentials are
	// used.
	IDP *FacebookGoogleConfigParams `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetFacebookIDPOptions : Instantiate SetFacebookIDPOptions
func (*AppIdManagementV4) NewSetFacebookIDPOptions(iDP *FacebookGoogleConfigParams) *SetFacebookIDPOptions {
	return &SetFacebookIDPOptions{
		IDP: iDP,
	}
}

// SetIDP : Allow user to set IDP
func (options *SetFacebookIDPOptions) SetIDP(iDP *FacebookGoogleConfigParams) *SetFacebookIDPOptions {
	options.IDP = iDP
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetFacebookIDPOptions) SetHeaders(param map[string]string) *SetFacebookIDPOptions {
	options.Headers = param
	return options
}

// SetGoogleIDPOptions : The SetGoogleIDP options.
type SetGoogleIDPOptions struct {
	// The identity provider configuration as a JSON object. If the configuration is not set, IBM default credentials are
	// used.
	IDP *FacebookGoogleConfigParams `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetGoogleIDPOptions : Instantiate SetGoogleIDPOptions
func (*AppIdManagementV4) NewSetGoogleIDPOptions(iDP *FacebookGoogleConfigParams) *SetGoogleIDPOptions {
	return &SetGoogleIDPOptions{
		IDP: iDP,
	}
}

// SetIDP : Allow user to set IDP
func (options *SetGoogleIDPOptions) SetIDP(iDP *FacebookGoogleConfigParams) *SetGoogleIDPOptions {
	options.IDP = iDP
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetGoogleIDPOptions) SetHeaders(param map[string]string) *SetGoogleIDPOptions {
	options.Headers = param
	return options
}

// SetSAMLIDPOptions : The SetSAMLIDP options.
type SetSAMLIDPOptions struct {
	IsActive *bool `validate:"required"`

	Config *SAMLConfigParams

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSetSAMLIDPOptions : Instantiate SetSAMLIDPOptions
func (*AppIdManagementV4) NewSetSAMLIDPOptions(isActive bool) *SetSAMLIDPOptions {
	return &SetSAMLIDPOptions{
		IsActive: core.BoolPtr(isActive),
	}
}

// SetIsActive : Allow user to set IsActive
func (options *SetSAMLIDPOptions) SetIsActive(isActive bool) *SetSAMLIDPOptions {
	options.IsActive = core.BoolPtr(isActive)
	return options
}

// SetConfig : Allow user to set Config
func (options *SetSAMLIDPOptions) SetConfig(config *SAMLConfigParams) *SetSAMLIDPOptions {
	options.Config = config
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SetSAMLIDPOptions) SetHeaders(param map[string]string) *SetSAMLIDPOptions {
	options.Headers = param
	return options
}

// SsoLogoutFromAllAppsOptions : The SsoLogoutFromAllApps options.
type SsoLogoutFromAllAppsOptions struct {
	// The ID assigned to a user when they sign in by using Cloud Directory.
	UserID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewSsoLogoutFromAllAppsOptions : Instantiate SsoLogoutFromAllAppsOptions
func (*AppIdManagementV4) NewSsoLogoutFromAllAppsOptions(userID string) *SsoLogoutFromAllAppsOptions {
	return &SsoLogoutFromAllAppsOptions{
		UserID: core.StringPtr(userID),
	}
}

// SetUserID : Allow user to set UserID
func (options *SsoLogoutFromAllAppsOptions) SetUserID(userID string) *SsoLogoutFromAllAppsOptions {
	options.UserID = core.StringPtr(userID)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *SsoLogoutFromAllAppsOptions) SetHeaders(param map[string]string) *SsoLogoutFromAllAppsOptions {
	options.Headers = param
	return options
}

// StartForgotPasswordOptions : The StartForgotPassword options.
type StartForgotPasswordOptions struct {
	// The user identitier - email or userName based on the identityField property in <a
	// href="#!/Identity_Providers/set_cloud_directory_idp" target="_blank"> cloud directory configuration.</a>.
	User *string `validate:"required"`

	// Preferred language for resource. Format as described at RFC5646.
	Language *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewStartForgotPasswordOptions : Instantiate StartForgotPasswordOptions
func (*AppIdManagementV4) NewStartForgotPasswordOptions(user string) *StartForgotPasswordOptions {
	return &StartForgotPasswordOptions{
		User: core.StringPtr(user),
	}
}

// SetUser : Allow user to set User
func (options *StartForgotPasswordOptions) SetUser(user string) *StartForgotPasswordOptions {
	options.User = core.StringPtr(user)
	return options
}

// SetLanguage : Allow user to set Language
func (options *StartForgotPasswordOptions) SetLanguage(language string) *StartForgotPasswordOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *StartForgotPasswordOptions) SetHeaders(param map[string]string) *StartForgotPasswordOptions {
	options.Headers = param
	return options
}

// StartSignUpOptions : The StartSignUp options.
type StartSignUpOptions struct {
	// A boolean indication if a profile should be created for the Cloud Directory user.
	ShouldCreateProfile *bool `validate:"required"`

	Emails []CreateNewUserEmailsItem `validate:"required"`

	Password *string `validate:"required"`

	Active *bool

	UserName *string

	// Preferred language for resource. Format as described at RFC5646.
	Language *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewStartSignUpOptions : Instantiate StartSignUpOptions
func (*AppIdManagementV4) NewStartSignUpOptions(shouldCreateProfile bool, emails []CreateNewUserEmailsItem, password string) *StartSignUpOptions {
	return &StartSignUpOptions{
		ShouldCreateProfile: core.BoolPtr(shouldCreateProfile),
		Emails: emails,
		Password: core.StringPtr(password),
	}
}

// SetShouldCreateProfile : Allow user to set ShouldCreateProfile
func (options *StartSignUpOptions) SetShouldCreateProfile(shouldCreateProfile bool) *StartSignUpOptions {
	options.ShouldCreateProfile = core.BoolPtr(shouldCreateProfile)
	return options
}

// SetEmails : Allow user to set Emails
func (options *StartSignUpOptions) SetEmails(emails []CreateNewUserEmailsItem) *StartSignUpOptions {
	options.Emails = emails
	return options
}

// SetPassword : Allow user to set Password
func (options *StartSignUpOptions) SetPassword(password string) *StartSignUpOptions {
	options.Password = core.StringPtr(password)
	return options
}

// SetActive : Allow user to set Active
func (options *StartSignUpOptions) SetActive(active bool) *StartSignUpOptions {
	options.Active = core.BoolPtr(active)
	return options
}

// SetUserName : Allow user to set UserName
func (options *StartSignUpOptions) SetUserName(userName string) *StartSignUpOptions {
	options.UserName = core.StringPtr(userName)
	return options
}

// SetLanguage : Allow user to set Language
func (options *StartSignUpOptions) SetLanguage(language string) *StartSignUpOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *StartSignUpOptions) SetHeaders(param map[string]string) *StartSignUpOptions {
	options.Headers = param
	return options
}

// UpdateApplicationOptions : The UpdateApplication options.
type UpdateApplicationOptions struct {
	// The application clientId.
	ClientID *string `validate:"required,ne="`

	// The application name to be updated. Application name cannot exceed 50 characters.
	Name *string `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateApplicationOptions : Instantiate UpdateApplicationOptions
func (*AppIdManagementV4) NewUpdateApplicationOptions(clientID string, name string) *UpdateApplicationOptions {
	return &UpdateApplicationOptions{
		ClientID: core.StringPtr(clientID),
		Name: core.StringPtr(name),
	}
}

// SetClientID : Allow user to set ClientID
func (options *UpdateApplicationOptions) SetClientID(clientID string) *UpdateApplicationOptions {
	options.ClientID = core.StringPtr(clientID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateApplicationOptions) SetName(name string) *UpdateApplicationOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateApplicationOptions) SetHeaders(param map[string]string) *UpdateApplicationOptions {
	options.Headers = param
	return options
}

// UpdateChannelOptions : The UpdateChannel options.
type UpdateChannelOptions struct {
	// The MFA channel.
	Channel *string `validate:"required,ne="`

	IsActive *bool `validate:"required"`

	Config interface{}

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateChannelOptions.Channel property.
// The MFA channel.
const (
	UpdateChannelOptions_Channel_Email = "email"
	UpdateChannelOptions_Channel_Nexmo = "nexmo"
)

// NewUpdateChannelOptions : Instantiate UpdateChannelOptions
func (*AppIdManagementV4) NewUpdateChannelOptions(channel string, isActive bool) *UpdateChannelOptions {
	return &UpdateChannelOptions{
		Channel: core.StringPtr(channel),
		IsActive: core.BoolPtr(isActive),
	}
}

// SetChannel : Allow user to set Channel
func (options *UpdateChannelOptions) SetChannel(channel string) *UpdateChannelOptions {
	options.Channel = core.StringPtr(channel)
	return options
}

// SetIsActive : Allow user to set IsActive
func (options *UpdateChannelOptions) SetIsActive(isActive bool) *UpdateChannelOptions {
	options.IsActive = core.BoolPtr(isActive)
	return options
}

// SetConfig : Allow user to set Config
func (options *UpdateChannelOptions) SetConfig(config interface{}) *UpdateChannelOptions {
	options.Config = config
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateChannelOptions) SetHeaders(param map[string]string) *UpdateChannelOptions {
	options.Headers = param
	return options
}

// UpdateCloudDirectoryUserOptions : The UpdateCloudDirectoryUser options.
type UpdateCloudDirectoryUserOptions struct {
	// The ID assigned to a user when they sign in by using Cloud Directory.
	UserID *string `validate:"required,ne="`

	Emails []CreateNewUserEmailsItem `validate:"required"`

	Active *bool

	UserName *string

	Password *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCloudDirectoryUserOptions : Instantiate UpdateCloudDirectoryUserOptions
func (*AppIdManagementV4) NewUpdateCloudDirectoryUserOptions(userID string, emails []CreateNewUserEmailsItem) *UpdateCloudDirectoryUserOptions {
	return &UpdateCloudDirectoryUserOptions{
		UserID: core.StringPtr(userID),
		Emails: emails,
	}
}

// SetUserID : Allow user to set UserID
func (options *UpdateCloudDirectoryUserOptions) SetUserID(userID string) *UpdateCloudDirectoryUserOptions {
	options.UserID = core.StringPtr(userID)
	return options
}

// SetEmails : Allow user to set Emails
func (options *UpdateCloudDirectoryUserOptions) SetEmails(emails []CreateNewUserEmailsItem) *UpdateCloudDirectoryUserOptions {
	options.Emails = emails
	return options
}

// SetActive : Allow user to set Active
func (options *UpdateCloudDirectoryUserOptions) SetActive(active bool) *UpdateCloudDirectoryUserOptions {
	options.Active = core.BoolPtr(active)
	return options
}

// SetUserName : Allow user to set UserName
func (options *UpdateCloudDirectoryUserOptions) SetUserName(userName string) *UpdateCloudDirectoryUserOptions {
	options.UserName = core.StringPtr(userName)
	return options
}

// SetPassword : Allow user to set Password
func (options *UpdateCloudDirectoryUserOptions) SetPassword(password string) *UpdateCloudDirectoryUserOptions {
	options.Password = core.StringPtr(password)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCloudDirectoryUserOptions) SetHeaders(param map[string]string) *UpdateCloudDirectoryUserOptions {
	options.Headers = param
	return options
}

// UpdateExtensionActiveOptions : The UpdateExtensionActive options.
type UpdateExtensionActiveOptions struct {
	// The name of the extension.
	Name *string `validate:"required,ne="`

	IsActive *bool `validate:"required"`

	Config interface{}

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateExtensionActiveOptions.Name property.
// The name of the extension.
const (
	UpdateExtensionActiveOptions_Name_Postmfa = "postmfa"
	UpdateExtensionActiveOptions_Name_Premfa = "premfa"
)

// NewUpdateExtensionActiveOptions : Instantiate UpdateExtensionActiveOptions
func (*AppIdManagementV4) NewUpdateExtensionActiveOptions(name string, isActive bool) *UpdateExtensionActiveOptions {
	return &UpdateExtensionActiveOptions{
		Name: core.StringPtr(name),
		IsActive: core.BoolPtr(isActive),
	}
}

// SetName : Allow user to set Name
func (options *UpdateExtensionActiveOptions) SetName(name string) *UpdateExtensionActiveOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetIsActive : Allow user to set IsActive
func (options *UpdateExtensionActiveOptions) SetIsActive(isActive bool) *UpdateExtensionActiveOptions {
	options.IsActive = core.BoolPtr(isActive)
	return options
}

// SetConfig : Allow user to set Config
func (options *UpdateExtensionActiveOptions) SetConfig(config interface{}) *UpdateExtensionActiveOptions {
	options.Config = config
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateExtensionActiveOptions) SetHeaders(param map[string]string) *UpdateExtensionActiveOptions {
	options.Headers = param
	return options
}

// UpdateExtensionConfigConfig : UpdateExtensionConfigConfig struct
type UpdateExtensionConfigConfig struct {
	URL *string `json:"url,omitempty"`

	HeadersVar interface{} `json:"headers,omitempty"`
}

// UnmarshalUpdateExtensionConfigConfig unmarshals an instance of UpdateExtensionConfigConfig from the specified map of raw messages.
func UnmarshalUpdateExtensionConfigConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateExtensionConfigConfig)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "headers", &obj.HeadersVar)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateExtensionConfigOptions : The UpdateExtensionConfig options.
type UpdateExtensionConfigOptions struct {
	// The name of the extension.
	Name *string `validate:"required,ne="`

	IsActive *bool `validate:"required"`

	Config *UpdateExtensionConfigConfig

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateExtensionConfigOptions.Name property.
// The name of the extension.
const (
	UpdateExtensionConfigOptions_Name_Postmfa = "postmfa"
	UpdateExtensionConfigOptions_Name_Premfa = "premfa"
)

// NewUpdateExtensionConfigOptions : Instantiate UpdateExtensionConfigOptions
func (*AppIdManagementV4) NewUpdateExtensionConfigOptions(name string, isActive bool) *UpdateExtensionConfigOptions {
	return &UpdateExtensionConfigOptions{
		Name: core.StringPtr(name),
		IsActive: core.BoolPtr(isActive),
	}
}

// SetName : Allow user to set Name
func (options *UpdateExtensionConfigOptions) SetName(name string) *UpdateExtensionConfigOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetIsActive : Allow user to set IsActive
func (options *UpdateExtensionConfigOptions) SetIsActive(isActive bool) *UpdateExtensionConfigOptions {
	options.IsActive = core.BoolPtr(isActive)
	return options
}

// SetConfig : Allow user to set Config
func (options *UpdateExtensionConfigOptions) SetConfig(config *UpdateExtensionConfigConfig) *UpdateExtensionConfigOptions {
	options.Config = config
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateExtensionConfigOptions) SetHeaders(param map[string]string) *UpdateExtensionConfigOptions {
	options.Headers = param
	return options
}

// UpdateLocalizationOptions : The UpdateLocalization options.
type UpdateLocalizationOptions struct {
	Languages []string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateLocalizationOptions : Instantiate UpdateLocalizationOptions
func (*AppIdManagementV4) NewUpdateLocalizationOptions() *UpdateLocalizationOptions {
	return &UpdateLocalizationOptions{}
}

// SetLanguages : Allow user to set Languages
func (options *UpdateLocalizationOptions) SetLanguages(languages []string) *UpdateLocalizationOptions {
	options.Languages = languages
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateLocalizationOptions) SetHeaders(param map[string]string) *UpdateLocalizationOptions {
	options.Headers = param
	return options
}

// UpdateMFAConfigOptions : The UpdateMFAConfig options.
type UpdateMFAConfigOptions struct {
	IsActive *bool `validate:"required"`

	Config interface{}

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateMFAConfigOptions : Instantiate UpdateMFAConfigOptions
func (*AppIdManagementV4) NewUpdateMFAConfigOptions(isActive bool) *UpdateMFAConfigOptions {
	return &UpdateMFAConfigOptions{
		IsActive: core.BoolPtr(isActive),
	}
}

// SetIsActive : Allow user to set IsActive
func (options *UpdateMFAConfigOptions) SetIsActive(isActive bool) *UpdateMFAConfigOptions {
	options.IsActive = core.BoolPtr(isActive)
	return options
}

// SetConfig : Allow user to set Config
func (options *UpdateMFAConfigOptions) SetConfig(config interface{}) *UpdateMFAConfigOptions {
	options.Config = config
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateMFAConfigOptions) SetHeaders(param map[string]string) *UpdateMFAConfigOptions {
	options.Headers = param
	return options
}

// UpdateRateLimitConfigOptions : The UpdateRateLimitConfig options.
type UpdateRateLimitConfigOptions struct {
	SignUpLimitPerMinute *int64 `validate:"required"`

	SignInLimitPerMinute *int64 `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateRateLimitConfigOptions : Instantiate UpdateRateLimitConfigOptions
func (*AppIdManagementV4) NewUpdateRateLimitConfigOptions(signUpLimitPerMinute int64, signInLimitPerMinute int64) *UpdateRateLimitConfigOptions {
	return &UpdateRateLimitConfigOptions{
		SignUpLimitPerMinute: core.Int64Ptr(signUpLimitPerMinute),
		SignInLimitPerMinute: core.Int64Ptr(signInLimitPerMinute),
	}
}

// SetSignUpLimitPerMinute : Allow user to set SignUpLimitPerMinute
func (options *UpdateRateLimitConfigOptions) SetSignUpLimitPerMinute(signUpLimitPerMinute int64) *UpdateRateLimitConfigOptions {
	options.SignUpLimitPerMinute = core.Int64Ptr(signUpLimitPerMinute)
	return options
}

// SetSignInLimitPerMinute : Allow user to set SignInLimitPerMinute
func (options *UpdateRateLimitConfigOptions) SetSignInLimitPerMinute(signInLimitPerMinute int64) *UpdateRateLimitConfigOptions {
	options.SignInLimitPerMinute = core.Int64Ptr(signInLimitPerMinute)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateRateLimitConfigOptions) SetHeaders(param map[string]string) *UpdateRateLimitConfigOptions {
	options.Headers = param
	return options
}

// UpdateRedirectUrisOptions : The UpdateRedirectUris options.
type UpdateRedirectUrisOptions struct {
	// The redirect URIs JSON object. If IBM default credentials are used, the redirect URIs are ignored.
	RedirectUrisArray *RedirectUriConfig `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateRedirectUrisOptions : Instantiate UpdateRedirectUrisOptions
func (*AppIdManagementV4) NewUpdateRedirectUrisOptions(redirectUrisArray *RedirectUriConfig) *UpdateRedirectUrisOptions {
	return &UpdateRedirectUrisOptions{
		RedirectUrisArray: redirectUrisArray,
	}
}

// SetRedirectUrisArray : Allow user to set RedirectUrisArray
func (options *UpdateRedirectUrisOptions) SetRedirectUrisArray(redirectUrisArray *RedirectUriConfig) *UpdateRedirectUrisOptions {
	options.RedirectUrisArray = redirectUrisArray
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateRedirectUrisOptions) SetHeaders(param map[string]string) *UpdateRedirectUrisOptions {
	options.Headers = param
	return options
}

// UpdateRoleOptions : The UpdateRole options.
type UpdateRoleOptions struct {
	// The role identifier.
	RoleID *string `validate:"required,ne="`

	Name *string `validate:"required"`

	Access []UpdateRoleParamsAccessItem `validate:"required"`

	Description *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateRoleOptions : Instantiate UpdateRoleOptions
func (*AppIdManagementV4) NewUpdateRoleOptions(roleID string, name string, access []UpdateRoleParamsAccessItem) *UpdateRoleOptions {
	return &UpdateRoleOptions{
		RoleID: core.StringPtr(roleID),
		Name: core.StringPtr(name),
		Access: access,
	}
}

// SetRoleID : Allow user to set RoleID
func (options *UpdateRoleOptions) SetRoleID(roleID string) *UpdateRoleOptions {
	options.RoleID = core.StringPtr(roleID)
	return options
}

// SetName : Allow user to set Name
func (options *UpdateRoleOptions) SetName(name string) *UpdateRoleOptions {
	options.Name = core.StringPtr(name)
	return options
}

// SetAccess : Allow user to set Access
func (options *UpdateRoleOptions) SetAccess(access []UpdateRoleParamsAccessItem) *UpdateRoleOptions {
	options.Access = access
	return options
}

// SetDescription : Allow user to set Description
func (options *UpdateRoleOptions) SetDescription(description string) *UpdateRoleOptions {
	options.Description = core.StringPtr(description)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateRoleOptions) SetHeaders(param map[string]string) *UpdateRoleOptions {
	options.Headers = param
	return options
}

// UpdateRoleParamsAccessItem : UpdateRoleParamsAccessItem struct
type UpdateRoleParamsAccessItem struct {
	ApplicationID *string `json:"application_id" validate:"required"`

	Scopes []string `json:"scopes" validate:"required"`
}

// NewUpdateRoleParamsAccessItem : Instantiate UpdateRoleParamsAccessItem (Generic Model Constructor)
func (*AppIdManagementV4) NewUpdateRoleParamsAccessItem(applicationID string, scopes []string) (model *UpdateRoleParamsAccessItem, err error) {
	model = &UpdateRoleParamsAccessItem{
		ApplicationID: core.StringPtr(applicationID),
		Scopes: scopes,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalUpdateRoleParamsAccessItem unmarshals an instance of UpdateRoleParamsAccessItem from the specified map of raw messages.
func UnmarshalUpdateRoleParamsAccessItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateRoleParamsAccessItem)
	err = core.UnmarshalPrimitive(m, "application_id", &obj.ApplicationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scopes", &obj.Scopes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateRolesResponseAccessItem : UpdateRolesResponseAccessItem struct
type UpdateRolesResponseAccessItem struct {
	ApplicationID *string `json:"application_id" validate:"required"`

	Scopes []string `json:"scopes" validate:"required"`
}

// UnmarshalUpdateRolesResponseAccessItem unmarshals an instance of UpdateRolesResponseAccessItem from the specified map of raw messages.
func UnmarshalUpdateRolesResponseAccessItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateRolesResponseAccessItem)
	err = core.UnmarshalPrimitive(m, "application_id", &obj.ApplicationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "scopes", &obj.Scopes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateSSOConfigOptions : The UpdateSSOConfig options.
type UpdateSSOConfigOptions struct {
	IsActive *bool `validate:"required"`

	InactivityTimeoutSeconds *float64 `validate:"required"`

	LogoutRedirectUris []string `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateSSOConfigOptions : Instantiate UpdateSSOConfigOptions
func (*AppIdManagementV4) NewUpdateSSOConfigOptions(isActive bool, inactivityTimeoutSeconds float64, logoutRedirectUris []string) *UpdateSSOConfigOptions {
	return &UpdateSSOConfigOptions{
		IsActive: core.BoolPtr(isActive),
		InactivityTimeoutSeconds: core.Float64Ptr(inactivityTimeoutSeconds),
		LogoutRedirectUris: logoutRedirectUris,
	}
}

// SetIsActive : Allow user to set IsActive
func (options *UpdateSSOConfigOptions) SetIsActive(isActive bool) *UpdateSSOConfigOptions {
	options.IsActive = core.BoolPtr(isActive)
	return options
}

// SetInactivityTimeoutSeconds : Allow user to set InactivityTimeoutSeconds
func (options *UpdateSSOConfigOptions) SetInactivityTimeoutSeconds(inactivityTimeoutSeconds float64) *UpdateSSOConfigOptions {
	options.InactivityTimeoutSeconds = core.Float64Ptr(inactivityTimeoutSeconds)
	return options
}

// SetLogoutRedirectUris : Allow user to set LogoutRedirectUris
func (options *UpdateSSOConfigOptions) SetLogoutRedirectUris(logoutRedirectUris []string) *UpdateSSOConfigOptions {
	options.LogoutRedirectUris = logoutRedirectUris
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateSSOConfigOptions) SetHeaders(param map[string]string) *UpdateSSOConfigOptions {
	options.Headers = param
	return options
}

// UpdateTemplateOptions : The UpdateTemplate options.
type UpdateTemplateOptions struct {
	// The type of email template. This can be "USER_VERIFICATION", "WELCOME", "PASSWORD_CHANGED", "RESET_PASSWORD" or
	// "MFA_VERIFICATION".
	TemplateName *string `validate:"required,ne="`

	// Preferred language for resource. Format as described at RFC5646. According to the configured languages codes
	// returned from the `GET /management/v4/{tenantId}/config/ui/languages` API.
	Language *string `validate:"required,ne="`

	Subject *string `validate:"required"`

	HTMLBody *string

	Base64EncodedHTMLBody *string

	PlainTextBody *string

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UpdateTemplateOptions.TemplateName property.
// The type of email template. This can be "USER_VERIFICATION", "WELCOME", "PASSWORD_CHANGED", "RESET_PASSWORD" or
// "MFA_VERIFICATION".
const (
	UpdateTemplateOptions_TemplateName_MfaVerification = "MFA_VERIFICATION"
	UpdateTemplateOptions_TemplateName_PasswordChanged = "PASSWORD_CHANGED"
	UpdateTemplateOptions_TemplateName_ResetPassword = "RESET_PASSWORD"
	UpdateTemplateOptions_TemplateName_UserVerification = "USER_VERIFICATION"
	UpdateTemplateOptions_TemplateName_Welcome = "WELCOME"
)

// NewUpdateTemplateOptions : Instantiate UpdateTemplateOptions
func (*AppIdManagementV4) NewUpdateTemplateOptions(templateName string, language string, subject string) *UpdateTemplateOptions {
	return &UpdateTemplateOptions{
		TemplateName: core.StringPtr(templateName),
		Language: core.StringPtr(language),
		Subject: core.StringPtr(subject),
	}
}

// SetTemplateName : Allow user to set TemplateName
func (options *UpdateTemplateOptions) SetTemplateName(templateName string) *UpdateTemplateOptions {
	options.TemplateName = core.StringPtr(templateName)
	return options
}

// SetLanguage : Allow user to set Language
func (options *UpdateTemplateOptions) SetLanguage(language string) *UpdateTemplateOptions {
	options.Language = core.StringPtr(language)
	return options
}

// SetSubject : Allow user to set Subject
func (options *UpdateTemplateOptions) SetSubject(subject string) *UpdateTemplateOptions {
	options.Subject = core.StringPtr(subject)
	return options
}

// SetHTMLBody : Allow user to set HTMLBody
func (options *UpdateTemplateOptions) SetHTMLBody(htmlBody string) *UpdateTemplateOptions {
	options.HTMLBody = core.StringPtr(htmlBody)
	return options
}

// SetBase64EncodedHTMLBody : Allow user to set Base64EncodedHTMLBody
func (options *UpdateTemplateOptions) SetBase64EncodedHTMLBody(base64EncodedHTMLBody string) *UpdateTemplateOptions {
	options.Base64EncodedHTMLBody = core.StringPtr(base64EncodedHTMLBody)
	return options
}

// SetPlainTextBody : Allow user to set PlainTextBody
func (options *UpdateTemplateOptions) SetPlainTextBody(plainTextBody string) *UpdateTemplateOptions {
	options.PlainTextBody = core.StringPtr(plainTextBody)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateTemplateOptions) SetHeaders(param map[string]string) *UpdateTemplateOptions {
	options.Headers = param
	return options
}

// UpdateUserProfilesConfigOptions : The UpdateUserProfilesConfig options.
type UpdateUserProfilesConfigOptions struct {
	IsActive *bool `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateUserProfilesConfigOptions : Instantiate UpdateUserProfilesConfigOptions
func (*AppIdManagementV4) NewUpdateUserProfilesConfigOptions(isActive bool) *UpdateUserProfilesConfigOptions {
	return &UpdateUserProfilesConfigOptions{
		IsActive: core.BoolPtr(isActive),
	}
}

// SetIsActive : Allow user to set IsActive
func (options *UpdateUserProfilesConfigOptions) SetIsActive(isActive bool) *UpdateUserProfilesConfigOptions {
	options.IsActive = core.BoolPtr(isActive)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateUserProfilesConfigOptions) SetHeaders(param map[string]string) *UpdateUserProfilesConfigOptions {
	options.Headers = param
	return options
}

// UpdateUserRolesOptions : The UpdateUserRoles options.
type UpdateUserRolesOptions struct {
	// The user's identifier ('subject' in identity token) You can search user in <a
	// href="#!/Users/users_search_user_profile" target="_blank">/users</a>.
	ID *string `validate:"required,ne="`

	Roles *UpdateUserRolesParamsRoles `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateUserRolesOptions : Instantiate UpdateUserRolesOptions
func (*AppIdManagementV4) NewUpdateUserRolesOptions(id string, roles *UpdateUserRolesParamsRoles) *UpdateUserRolesOptions {
	return &UpdateUserRolesOptions{
		ID: core.StringPtr(id),
		Roles: roles,
	}
}

// SetID : Allow user to set ID
func (options *UpdateUserRolesOptions) SetID(id string) *UpdateUserRolesOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetRoles : Allow user to set Roles
func (options *UpdateUserRolesOptions) SetRoles(roles *UpdateUserRolesParamsRoles) *UpdateUserRolesOptions {
	options.Roles = roles
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateUserRolesOptions) SetHeaders(param map[string]string) *UpdateUserRolesOptions {
	options.Headers = param
	return options
}

// UpdateUserRolesParamsRoles : UpdateUserRolesParamsRoles struct
type UpdateUserRolesParamsRoles struct {
	Ids []string `json:"ids,omitempty"`
}

// UnmarshalUpdateUserRolesParamsRoles unmarshals an instance of UpdateUserRolesParamsRoles from the specified map of raw messages.
func UnmarshalUpdateUserRolesParamsRoles(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateUserRolesParamsRoles)
	err = core.UnmarshalPrimitive(m, "ids", &obj.Ids)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UserProfilesExportOptions : The UserProfilesExport options.
type UserProfilesExportOptions struct {
	// The first result in a set list of results.
	StartIndex *int64

	// The maximum number of results per page.
	Count *int64

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUserProfilesExportOptions : Instantiate UserProfilesExportOptions
func (*AppIdManagementV4) NewUserProfilesExportOptions() *UserProfilesExportOptions {
	return &UserProfilesExportOptions{}
}

// SetStartIndex : Allow user to set StartIndex
func (options *UserProfilesExportOptions) SetStartIndex(startIndex int64) *UserProfilesExportOptions {
	options.StartIndex = core.Int64Ptr(startIndex)
	return options
}

// SetCount : Allow user to set Count
func (options *UserProfilesExportOptions) SetCount(count int64) *UserProfilesExportOptions {
	options.Count = core.Int64Ptr(count)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UserProfilesExportOptions) SetHeaders(param map[string]string) *UserProfilesExportOptions {
	options.Headers = param
	return options
}

// UserProfilesImportOptions : The UserProfilesImport options.
type UserProfilesImportOptions struct {
	Users []ExportUserProfileUsersItem `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUserProfilesImportOptions : Instantiate UserProfilesImportOptions
func (*AppIdManagementV4) NewUserProfilesImportOptions(users []ExportUserProfileUsersItem) *UserProfilesImportOptions {
	return &UserProfilesImportOptions{
		Users: users,
	}
}

// SetUsers : Allow user to set Users
func (options *UserProfilesImportOptions) SetUsers(users []ExportUserProfileUsersItem) *UserProfilesImportOptions {
	options.Users = users
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UserProfilesImportOptions) SetHeaders(param map[string]string) *UserProfilesImportOptions {
	options.Headers = param
	return options
}

// UserSearchResponseRequestOptions : UserSearchResponseRequestOptions struct
type UserSearchResponseRequestOptions struct {
	StartIndex *float64 `json:"startIndex,omitempty"`

	Count *float64 `json:"count,omitempty"`
}

// UnmarshalUserSearchResponseRequestOptions unmarshals an instance of UserSearchResponseRequestOptions from the specified map of raw messages.
func UnmarshalUserSearchResponseRequestOptions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UserSearchResponseRequestOptions)
	err = core.UnmarshalPrimitive(m, "startIndex", &obj.StartIndex)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UserSearchResponseUsersItem : UserSearchResponseUsersItem struct
type UserSearchResponseUsersItem struct {
	ID *string `json:"id,omitempty"`

	IDP *string `json:"idp,omitempty"`

	Email *string `json:"email,omitempty"`
}

// UnmarshalUserSearchResponseUsersItem unmarshals an instance of UserSearchResponseUsersItem from the specified map of raw messages.
func UnmarshalUserSearchResponseUsersItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UserSearchResponseUsersItem)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "idp", &obj.IDP)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "email", &obj.Email)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UserVerificationResultOptions : The UserVerificationResult options.
type UserVerificationResultOptions struct {
	// The context that will be use to get the verification result.
	Context *string `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUserVerificationResultOptions : Instantiate UserVerificationResultOptions
func (*AppIdManagementV4) NewUserVerificationResultOptions(context string) *UserVerificationResultOptions {
	return &UserVerificationResultOptions{
		Context: core.StringPtr(context),
	}
}

// SetContext : Allow user to set Context
func (options *UserVerificationResultOptions) SetContext(context string) *UserVerificationResultOptions {
	options.Context = core.StringPtr(context)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UserVerificationResultOptions) SetHeaders(param map[string]string) *UserVerificationResultOptions {
	options.Headers = param
	return options
}

// UsersDeleteUserProfileOptions : The UsersDeleteUserProfile options.
type UsersDeleteUserProfileOptions struct {
	// The user's identifier ('subject' in identity token) You can search user in <a
	// href="#!/Users/users_search_user_profile" target="_blank">/users</a>.
	ID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUsersDeleteUserProfileOptions : Instantiate UsersDeleteUserProfileOptions
func (*AppIdManagementV4) NewUsersDeleteUserProfileOptions(id string) *UsersDeleteUserProfileOptions {
	return &UsersDeleteUserProfileOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *UsersDeleteUserProfileOptions) SetID(id string) *UsersDeleteUserProfileOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UsersDeleteUserProfileOptions) SetHeaders(param map[string]string) *UsersDeleteUserProfileOptions {
	options.Headers = param
	return options
}

// UsersGetUserProfileOptions : The UsersGetUserProfile options.
type UsersGetUserProfileOptions struct {
	// The user's identifier ('subject' in identity token) You can search user in <a
	// href="#!/Users/users_search_user_profile" target="_blank">/users</a>.
	ID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUsersGetUserProfileOptions : Instantiate UsersGetUserProfileOptions
func (*AppIdManagementV4) NewUsersGetUserProfileOptions(id string) *UsersGetUserProfileOptions {
	return &UsersGetUserProfileOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *UsersGetUserProfileOptions) SetID(id string) *UsersGetUserProfileOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UsersGetUserProfileOptions) SetHeaders(param map[string]string) *UsersGetUserProfileOptions {
	options.Headers = param
	return options
}

// UsersList : UsersList struct
type UsersList struct {
	TotalResults *float64 `json:"totalResults,omitempty"`

	ItemsPerPage *float64 `json:"itemsPerPage,omitempty"`

	Resources []interface{} `json:"Resources" validate:"required"`
}

// UnmarshalUsersList unmarshals an instance of UsersList from the specified map of raw messages.
func UnmarshalUsersList(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UsersList)
	err = core.UnmarshalPrimitive(m, "totalResults", &obj.TotalResults)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "itemsPerPage", &obj.ItemsPerPage)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "Resources", &obj.Resources)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UsersNominateUserOptions : The UsersNominateUser options.
type UsersNominateUserOptions struct {
	IDP *string `validate:"required"`

	IDPIdentity *string `validate:"required"`

	Profile *UsersNominateUserParamsProfile

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UsersNominateUserOptions.IDP property.
const (
	UsersNominateUserOptions_IDP_AppidCustom = "appid_custom"
	UsersNominateUserOptions_IDP_CloudDirectory = "cloud_directory"
	UsersNominateUserOptions_IDP_Facebook = "facebook"
	UsersNominateUserOptions_IDP_Google = "google"
	UsersNominateUserOptions_IDP_Ibmid = "ibmid"
	UsersNominateUserOptions_IDP_SAML = "saml"
)

// NewUsersNominateUserOptions : Instantiate UsersNominateUserOptions
func (*AppIdManagementV4) NewUsersNominateUserOptions(iDP string, iDPIdentity string) *UsersNominateUserOptions {
	return &UsersNominateUserOptions{
		IDP: core.StringPtr(iDP),
		IDPIdentity: core.StringPtr(iDPIdentity),
	}
}

// SetIDP : Allow user to set IDP
func (options *UsersNominateUserOptions) SetIDP(iDP string) *UsersNominateUserOptions {
	options.IDP = core.StringPtr(iDP)
	return options
}

// SetIDPIdentity : Allow user to set IDPIdentity
func (options *UsersNominateUserOptions) SetIDPIdentity(iDPIdentity string) *UsersNominateUserOptions {
	options.IDPIdentity = core.StringPtr(iDPIdentity)
	return options
}

// SetProfile : Allow user to set Profile
func (options *UsersNominateUserOptions) SetProfile(profile *UsersNominateUserParamsProfile) *UsersNominateUserOptions {
	options.Profile = profile
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UsersNominateUserOptions) SetHeaders(param map[string]string) *UsersNominateUserOptions {
	options.Headers = param
	return options
}

// UsersNominateUserParamsProfile : UsersNominateUserParamsProfile struct
type UsersNominateUserParamsProfile struct {
	Attributes map[string]interface{} `json:"attributes,omitempty"`
}

// UnmarshalUsersNominateUserParamsProfile unmarshals an instance of UsersNominateUserParamsProfile from the specified map of raw messages.
func UnmarshalUsersNominateUserParamsProfile(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UsersNominateUserParamsProfile)
	err = core.UnmarshalPrimitive(m, "attributes", &obj.Attributes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UsersRevokeRefreshTokenOptions : The UsersRevokeRefreshToken options.
type UsersRevokeRefreshTokenOptions struct {
	// The user's identifier ('subject' in identity token) You can search user in <a
	// href="#!/Users/users_search_user_profile" target="_blank">/users</a>.
	ID *string `validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUsersRevokeRefreshTokenOptions : Instantiate UsersRevokeRefreshTokenOptions
func (*AppIdManagementV4) NewUsersRevokeRefreshTokenOptions(id string) *UsersRevokeRefreshTokenOptions {
	return &UsersRevokeRefreshTokenOptions{
		ID: core.StringPtr(id),
	}
}

// SetID : Allow user to set ID
func (options *UsersRevokeRefreshTokenOptions) SetID(id string) *UsersRevokeRefreshTokenOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UsersRevokeRefreshTokenOptions) SetHeaders(param map[string]string) *UsersRevokeRefreshTokenOptions {
	options.Headers = param
	return options
}

// UsersSearchUserProfileOptions : The UsersSearchUserProfile options.
type UsersSearchUserProfileOptions struct {
	// display user data.
	DataScope *string `validate:"required"`

	// Email (as retrieved from the Identity Provider).
	Email *string

	// The IDP specific user identifier.
	ID *string

	// The first result in a set list of results.
	StartIndex *int64

	// The maximum number of results per page.
	Count *int64

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the UsersSearchUserProfileOptions.DataScope property.
// display user data.
const (
	UsersSearchUserProfileOptions_DataScope_Full = "full"
	UsersSearchUserProfileOptions_DataScope_Index = "index"
)

// NewUsersSearchUserProfileOptions : Instantiate UsersSearchUserProfileOptions
func (*AppIdManagementV4) NewUsersSearchUserProfileOptions(dataScope string) *UsersSearchUserProfileOptions {
	return &UsersSearchUserProfileOptions{
		DataScope: core.StringPtr(dataScope),
	}
}

// SetDataScope : Allow user to set DataScope
func (options *UsersSearchUserProfileOptions) SetDataScope(dataScope string) *UsersSearchUserProfileOptions {
	options.DataScope = core.StringPtr(dataScope)
	return options
}

// SetEmail : Allow user to set Email
func (options *UsersSearchUserProfileOptions) SetEmail(email string) *UsersSearchUserProfileOptions {
	options.Email = core.StringPtr(email)
	return options
}

// SetID : Allow user to set ID
func (options *UsersSearchUserProfileOptions) SetID(id string) *UsersSearchUserProfileOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetStartIndex : Allow user to set StartIndex
func (options *UsersSearchUserProfileOptions) SetStartIndex(startIndex int64) *UsersSearchUserProfileOptions {
	options.StartIndex = core.Int64Ptr(startIndex)
	return options
}

// SetCount : Allow user to set Count
func (options *UsersSearchUserProfileOptions) SetCount(count int64) *UsersSearchUserProfileOptions {
	options.Count = core.Int64Ptr(count)
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UsersSearchUserProfileOptions) SetHeaders(param map[string]string) *UsersSearchUserProfileOptions {
	options.Headers = param
	return options
}

// UsersSetUserProfileOptions : The UsersSetUserProfile options.
type UsersSetUserProfileOptions struct {
	// The user's identifier ('subject' in identity token) You can search user in <a
	// href="#!/Users/users_search_user_profile" target="_blank">/users</a>.
	ID *string `validate:"required,ne="`

	Attributes map[string]interface{} `validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUsersSetUserProfileOptions : Instantiate UsersSetUserProfileOptions
func (*AppIdManagementV4) NewUsersSetUserProfileOptions(id string, attributes map[string]interface{}) *UsersSetUserProfileOptions {
	return &UsersSetUserProfileOptions{
		ID: core.StringPtr(id),
		Attributes: attributes,
	}
}

// SetID : Allow user to set ID
func (options *UsersSetUserProfileOptions) SetID(id string) *UsersSetUserProfileOptions {
	options.ID = core.StringPtr(id)
	return options
}

// SetAttributes : Allow user to set Attributes
func (options *UsersSetUserProfileOptions) SetAttributes(attributes map[string]interface{}) *UsersSetUserProfileOptions {
	options.Attributes = attributes
	return options
}

// SetHeaders : Allow user to set Headers
func (options *UsersSetUserProfileOptions) SetHeaders(param map[string]string) *UsersSetUserProfileOptions {
	options.Headers = param
	return options
}

// ActionUrlResponse : ActionUrlResponse struct
type ActionUrlResponse struct {
	ActionURL *string `json:"actionUrl" validate:"required"`
}

// UnmarshalActionUrlResponse unmarshals an instance of ActionUrlResponse from the specified map of raw messages.
func UnmarshalActionUrlResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ActionUrlResponse)
	err = core.UnmarshalPrimitive(m, "actionUrl", &obj.ActionURL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ApmSchema : ApmSchema struct
type ApmSchema struct {
	AdvancedPasswordManagement *ApmSchemaAdvancedPasswordManagement `json:"advancedPasswordManagement" validate:"required"`
}

// NewApmSchema : Instantiate ApmSchema (Generic Model Constructor)
func (*AppIdManagementV4) NewApmSchema(advancedPasswordManagement *ApmSchemaAdvancedPasswordManagement) (model *ApmSchema, err error) {
	model = &ApmSchema{
		AdvancedPasswordManagement: advancedPasswordManagement,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalApmSchema unmarshals an instance of ApmSchema from the specified map of raw messages.
func UnmarshalApmSchema(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ApmSchema)
	err = core.UnmarshalModel(m, "advancedPasswordManagement", &obj.AdvancedPasswordManagement, UnmarshalApmSchemaAdvancedPasswordManagement)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// AssignRoleToUser : AssignRoleToUser struct
type AssignRoleToUser struct {
	Roles []AssignRoleToUserRolesItem `json:"roles,omitempty"`
}

// UnmarshalAssignRoleToUser unmarshals an instance of AssignRoleToUser from the specified map of raw messages.
func UnmarshalAssignRoleToUser(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(AssignRoleToUser)
	err = core.UnmarshalModel(m, "roles", &obj.Roles, UnmarshalAssignRoleToUserRolesItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CloudDirectoryConfigParams : CloudDirectoryConfigParams struct
type CloudDirectoryConfigParams struct {
	SelfServiceEnabled *bool `json:"selfServiceEnabled" validate:"required"`

	SignupEnabled *bool `json:"signupEnabled,omitempty"`

	Interactions *CloudDirectoryConfigParamsInteractions `json:"interactions" validate:"required"`

	IdentityField *string `json:"identityField,omitempty"`
}

// Constants associated with the CloudDirectoryConfigParams.IdentityField property.
const (
	CloudDirectoryConfigParams_IdentityField_Email = "email"
	CloudDirectoryConfigParams_IdentityField_Username = "userName"
)

// NewCloudDirectoryConfigParams : Instantiate CloudDirectoryConfigParams (Generic Model Constructor)
func (*AppIdManagementV4) NewCloudDirectoryConfigParams(selfServiceEnabled bool, interactions *CloudDirectoryConfigParamsInteractions) (model *CloudDirectoryConfigParams, err error) {
	model = &CloudDirectoryConfigParams{
		SelfServiceEnabled: core.BoolPtr(selfServiceEnabled),
		Interactions: interactions,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCloudDirectoryConfigParams unmarshals an instance of CloudDirectoryConfigParams from the specified map of raw messages.
func UnmarshalCloudDirectoryConfigParams(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CloudDirectoryConfigParams)
	err = core.UnmarshalPrimitive(m, "selfServiceEnabled", &obj.SelfServiceEnabled)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "signupEnabled", &obj.SignupEnabled)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "interactions", &obj.Interactions, UnmarshalCloudDirectoryConfigParamsInteractions)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "identityField", &obj.IdentityField)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CloudDirectoryResponse : CloudDirectoryResponse struct
type CloudDirectoryResponse struct {
	IsActive *bool `json:"isActive" validate:"required"`

	Config *CloudDirectoryConfigParams `json:"config,omitempty"`
}

// UnmarshalCloudDirectoryResponse unmarshals an instance of CloudDirectoryResponse from the specified map of raw messages.
func UnmarshalCloudDirectoryResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CloudDirectoryResponse)
	err = core.UnmarshalPrimitive(m, "isActive", &obj.IsActive)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalCloudDirectoryConfigParams)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CloudDirectorySenderDetails : CloudDirectorySenderDetails struct
type CloudDirectorySenderDetails struct {
	SenderDetails *CloudDirectorySenderDetailsSenderDetails `json:"senderDetails" validate:"required"`
}

// NewCloudDirectorySenderDetails : Instantiate CloudDirectorySenderDetails (Generic Model Constructor)
func (*AppIdManagementV4) NewCloudDirectorySenderDetails(senderDetails *CloudDirectorySenderDetailsSenderDetails) (model *CloudDirectorySenderDetails, err error) {
	model = &CloudDirectorySenderDetails{
		SenderDetails: senderDetails,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCloudDirectorySenderDetails unmarshals an instance of CloudDirectorySenderDetails from the specified map of raw messages.
func UnmarshalCloudDirectorySenderDetails(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CloudDirectorySenderDetails)
	err = core.UnmarshalModel(m, "senderDetails", &obj.SenderDetails, UnmarshalCloudDirectorySenderDetailsSenderDetails)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ConfirmationResultOK : ConfirmationResultOK struct
type ConfirmationResultOK struct {
	Success *bool `json:"success" validate:"required"`

	UUID *string `json:"uuid" validate:"required"`
}

// UnmarshalConfirmationResultOK unmarshals an instance of ConfirmationResultOK from the specified map of raw messages.
func UnmarshalConfirmationResultOK(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ConfirmationResultOK)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uuid", &obj.UUID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateRolesResponse : CreateRolesResponse struct
type CreateRolesResponse struct {
	ID *string `json:"id" validate:"required"`

	Name *string `json:"name" validate:"required"`

	Description *string `json:"description,omitempty"`

	Access []CreateRolesResponseAccessItem `json:"access" validate:"required"`
}

// UnmarshalCreateRolesResponse unmarshals an instance of CreateRolesResponse from the specified map of raw messages.
func UnmarshalCreateRolesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CreateRolesResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "access", &obj.Access, UnmarshalCreateRolesResponseAccessItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CustomIdPConfigParams : CustomIdPConfigParams struct
type CustomIdPConfigParams struct {
	IsActive *bool `json:"isActive" validate:"required"`

	Config *CustomIdPConfigParamsConfig `json:"config,omitempty"`
}

// NewCustomIdPConfigParams : Instantiate CustomIdPConfigParams (Generic Model Constructor)
func (*AppIdManagementV4) NewCustomIdPConfigParams(isActive bool) (model *CustomIdPConfigParams, err error) {
	model = &CustomIdPConfigParams{
		IsActive: core.BoolPtr(isActive),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalCustomIdPConfigParams unmarshals an instance of CustomIdPConfigParams from the specified map of raw messages.
func UnmarshalCustomIdPConfigParams(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CustomIdPConfigParams)
	err = core.UnmarshalPrimitive(m, "isActive", &obj.IsActive)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalCustomIdPConfigParamsConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// EmailDispatcherParams : EmailDispatcherParams struct
type EmailDispatcherParams struct {
	Provider *string `json:"provider" validate:"required"`

	Sendgrid *EmailDispatcherParamsSendgrid `json:"sendgrid,omitempty"`

	Custom *EmailDispatcherParamsCustom `json:"custom,omitempty"`
}

// Constants associated with the EmailDispatcherParams.Provider property.
const (
	EmailDispatcherParams_Provider_Appid = "appid"
	EmailDispatcherParams_Provider_Custom = "custom"
	EmailDispatcherParams_Provider_Sendgrid = "sendgrid"
)

// NewEmailDispatcherParams : Instantiate EmailDispatcherParams (Generic Model Constructor)
func (*AppIdManagementV4) NewEmailDispatcherParams(provider string) (model *EmailDispatcherParams, err error) {
	model = &EmailDispatcherParams{
		Provider: core.StringPtr(provider),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalEmailDispatcherParams unmarshals an instance of EmailDispatcherParams from the specified map of raw messages.
func UnmarshalEmailDispatcherParams(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(EmailDispatcherParams)
	err = core.UnmarshalPrimitive(m, "provider", &obj.Provider)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "sendgrid", &obj.Sendgrid, UnmarshalEmailDispatcherParamsSendgrid)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "custom", &obj.Custom, UnmarshalEmailDispatcherParamsCustom)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExportUser : ExportUser struct
type ExportUser struct {
	Users []ExportUserUsersItem `json:"users" validate:"required"`
}

// NewExportUser : Instantiate ExportUser (Generic Model Constructor)
func (*AppIdManagementV4) NewExportUser(users []ExportUserUsersItem) (model *ExportUser, err error) {
	model = &ExportUser{
		Users: users,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalExportUser unmarshals an instance of ExportUser from the specified map of raw messages.
func UnmarshalExportUser(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExportUser)
	err = core.UnmarshalModel(m, "users", &obj.Users, UnmarshalExportUserUsersItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExportUserProfile : ExportUserProfile struct
type ExportUserProfile struct {
	Users []ExportUserProfileUsersItem `json:"users" validate:"required"`
}

// NewExportUserProfile : Instantiate ExportUserProfile (Generic Model Constructor)
func (*AppIdManagementV4) NewExportUserProfile(users []ExportUserProfileUsersItem) (model *ExportUserProfile, err error) {
	model = &ExportUserProfile{
		Users: users,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalExportUserProfile unmarshals an instance of ExportUserProfile from the specified map of raw messages.
func UnmarshalExportUserProfile(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExportUserProfile)
	err = core.UnmarshalModel(m, "users", &obj.Users, UnmarshalExportUserProfileUsersItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExtensionActive : ExtensionActive struct
type ExtensionActive struct {
	IsActive *bool `json:"isActive" validate:"required"`

	Config interface{} `json:"config,omitempty"`
}

// NewExtensionActive : Instantiate ExtensionActive (Generic Model Constructor)
func (*AppIdManagementV4) NewExtensionActive(isActive bool) (model *ExtensionActive, err error) {
	model = &ExtensionActive{
		IsActive: core.BoolPtr(isActive),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalExtensionActive unmarshals an instance of ExtensionActive from the specified map of raw messages.
func UnmarshalExtensionActive(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExtensionActive)
	err = core.UnmarshalPrimitive(m, "isActive", &obj.IsActive)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "config", &obj.Config)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ExtensionTest : ExtensionTest struct
type ExtensionTest struct {
	StatusCode *float64 `json:"statusCode,omitempty"`

	HeadersVar interface{} `json:"headers,omitempty"`
}

// UnmarshalExtensionTest unmarshals an instance of ExtensionTest from the specified map of raw messages.
func UnmarshalExtensionTest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ExtensionTest)
	err = core.UnmarshalPrimitive(m, "statusCode", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "headers", &obj.HeadersVar)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FacebookConfigParams : FacebookConfigParams struct
type FacebookConfigParams struct {
	IsActive *bool `json:"isActive" validate:"required"`

	Config *FacebookConfigParamsConfig `json:"config,omitempty"`
}

// UnmarshalFacebookConfigParams unmarshals an instance of FacebookConfigParams from the specified map of raw messages.
func UnmarshalFacebookConfigParams(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FacebookConfigParams)
	err = core.UnmarshalPrimitive(m, "isActive", &obj.IsActive)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalFacebookConfigParamsConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FacebookConfigParamsPUT : FacebookConfigParamsPUT struct
type FacebookConfigParamsPUT struct {
	IsActive *bool `json:"isActive" validate:"required"`

	Config *FacebookConfigParamsPUTConfig `json:"config,omitempty"`
}

// UnmarshalFacebookConfigParamsPUT unmarshals an instance of FacebookConfigParamsPUT from the specified map of raw messages.
func UnmarshalFacebookConfigParamsPUT(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FacebookConfigParamsPUT)
	err = core.UnmarshalPrimitive(m, "isActive", &obj.IsActive)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalFacebookConfigParamsPUTConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FacebookGoogleConfigParams : FacebookGoogleConfigParams struct
type FacebookGoogleConfigParams struct {
	IsActive *bool `json:"isActive" validate:"required"`

	Config *FacebookGoogleConfigParamsConfig `json:"config,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// NewFacebookGoogleConfigParams : Instantiate FacebookGoogleConfigParams (Generic Model Constructor)
func (*AppIdManagementV4) NewFacebookGoogleConfigParams(isActive bool) (model *FacebookGoogleConfigParams, err error) {
	model = &FacebookGoogleConfigParams{
		IsActive: core.BoolPtr(isActive),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// SetProperty allows the user to set an arbitrary property on an instance of FacebookGoogleConfigParams
func (o *FacebookGoogleConfigParams) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of FacebookGoogleConfigParams
func (o *FacebookGoogleConfigParams) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of FacebookGoogleConfigParams
func (o *FacebookGoogleConfigParams) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of FacebookGoogleConfigParams
func (o *FacebookGoogleConfigParams) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.IsActive != nil {
		m["isActive"] = o.IsActive
	}
	if o.Config != nil {
		m["config"] = o.Config
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalFacebookGoogleConfigParams unmarshals an instance of FacebookGoogleConfigParams from the specified map of raw messages.
func UnmarshalFacebookGoogleConfigParams(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FacebookGoogleConfigParams)
	err = core.UnmarshalPrimitive(m, "isActive", &obj.IsActive)
	if err != nil {
		return
	}
	delete(m, "isActive")
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalFacebookGoogleConfigParamsConfig)
	if err != nil {
		return
	}
	delete(m, "config")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetLanguages : GetLanguages struct
type GetLanguages struct {
	Languages []string `json:"languages" validate:"required"`
}

// NewGetLanguages : Instantiate GetLanguages (Generic Model Constructor)
func (*AppIdManagementV4) NewGetLanguages(languages []string) (model *GetLanguages, err error) {
	model = &GetLanguages{
		Languages: languages,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalGetLanguages unmarshals an instance of GetLanguages from the specified map of raw messages.
func UnmarshalGetLanguages(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetLanguages)
	err = core.UnmarshalPrimitive(m, "languages", &obj.Languages)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetMFAConfiguration : GetMFAConfiguration struct
type GetMFAConfiguration struct {
	IsActive *bool `json:"isActive" validate:"required"`
}

// UnmarshalGetMFAConfiguration unmarshals an instance of GetMFAConfiguration from the specified map of raw messages.
func UnmarshalGetMFAConfiguration(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetMFAConfiguration)
	err = core.UnmarshalPrimitive(m, "isActive", &obj.IsActive)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetRoleResponse : GetRoleResponse struct
type GetRoleResponse struct {
	ID *string `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Access []GetRoleResponseAccessItem `json:"access,omitempty"`
}

// UnmarshalGetRoleResponse unmarshals an instance of GetRoleResponse from the specified map of raw messages.
func UnmarshalGetRoleResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetRoleResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "access", &obj.Access, UnmarshalGetRoleResponseAccessItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetSMSChannel : GetSMSChannel struct
type GetSMSChannel struct {
	IsActive *bool `json:"isActive" validate:"required"`

	Type *string `json:"type" validate:"required"`

	Config *GetSMSChannelConfig `json:"config,omitempty"`
}

// UnmarshalGetSMSChannel unmarshals an instance of GetSMSChannel from the specified map of raw messages.
func UnmarshalGetSMSChannel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetSMSChannel)
	err = core.UnmarshalPrimitive(m, "isActive", &obj.IsActive)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalGetSMSChannelConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetScopesForApplication : GetScopesForApplication struct
type GetScopesForApplication struct {
	Scopes []string `json:"scopes,omitempty"`
}

// UnmarshalGetScopesForApplication unmarshals an instance of GetScopesForApplication from the specified map of raw messages.
func UnmarshalGetScopesForApplication(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetScopesForApplication)
	err = core.UnmarshalPrimitive(m, "scopes", &obj.Scopes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetTemplate : GetTemplate struct
type GetTemplate struct {
	Subject *string `json:"subject" validate:"required"`

	HTMLBody *string `json:"html_body,omitempty"`

	Base64EncodedHTMLBody *string `json:"base64_encoded_html_body,omitempty"`

	PlainTextBody *string `json:"plain_text_body,omitempty"`
}

// UnmarshalGetTemplate unmarshals an instance of GetTemplate from the specified map of raw messages.
func UnmarshalGetTemplate(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetTemplate)
	err = core.UnmarshalPrimitive(m, "subject", &obj.Subject)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "html_body", &obj.HTMLBody)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "base64_encoded_html_body", &obj.Base64EncodedHTMLBody)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "plain_text_body", &obj.PlainTextBody)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetUserAndProfile : GetUserAndProfile struct
type GetUserAndProfile struct {
	Sub *string `json:"sub,omitempty"`

	Identities []GetUserAndProfileIdentitiesItem `json:"identities,omitempty"`

	Attributes interface{} `json:"attributes,omitempty"`
}

// UnmarshalGetUserAndProfile unmarshals an instance of GetUserAndProfile from the specified map of raw messages.
func UnmarshalGetUserAndProfile(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetUserAndProfile)
	err = core.UnmarshalPrimitive(m, "sub", &obj.Sub)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "identities", &obj.Identities, UnmarshalGetUserAndProfileIdentitiesItem)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "attributes", &obj.Attributes)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetUserRolesResponse : GetUserRolesResponse struct
type GetUserRolesResponse struct {
	Roles []GetUserRolesResponseRolesItem `json:"roles,omitempty"`
}

// UnmarshalGetUserRolesResponse unmarshals an instance of GetUserRolesResponse from the specified map of raw messages.
func UnmarshalGetUserRolesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GetUserRolesResponse)
	err = core.UnmarshalModel(m, "roles", &obj.Roles, UnmarshalGetUserRolesResponseRolesItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GoogleConfigParams : GoogleConfigParams struct
type GoogleConfigParams struct {
	IsActive *bool `json:"isActive" validate:"required"`

	Config *GoogleConfigParamsConfig `json:"config,omitempty"`
}

// UnmarshalGoogleConfigParams unmarshals an instance of GoogleConfigParams from the specified map of raw messages.
func UnmarshalGoogleConfigParams(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GoogleConfigParams)
	err = core.UnmarshalPrimitive(m, "isActive", &obj.IsActive)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalGoogleConfigParamsConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GoogleConfigParamsPUT : GoogleConfigParamsPUT struct
type GoogleConfigParamsPUT struct {
	IsActive *bool `json:"isActive" validate:"required"`

	Config *GoogleConfigParamsPUTConfig `json:"config,omitempty"`
}

// UnmarshalGoogleConfigParamsPUT unmarshals an instance of GoogleConfigParamsPUT from the specified map of raw messages.
func UnmarshalGoogleConfigParamsPUT(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(GoogleConfigParamsPUT)
	err = core.UnmarshalPrimitive(m, "isActive", &obj.IsActive)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalGoogleConfigParamsPUTConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImportProfilesResponse : ImportProfilesResponse struct
type ImportProfilesResponse struct {
	Added *float64 `json:"added,omitempty"`

	Failed *float64 `json:"failed,omitempty"`

	FailReasons []ImportProfilesResponseFailReasonsItem `json:"failReasons,omitempty"`
}

// UnmarshalImportProfilesResponse unmarshals an instance of ImportProfilesResponse from the specified map of raw messages.
func UnmarshalImportProfilesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImportProfilesResponse)
	err = core.UnmarshalPrimitive(m, "added", &obj.Added)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "failed", &obj.Failed)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "failReasons", &obj.FailReasons, UnmarshalImportProfilesResponseFailReasonsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ImportResponse : ImportResponse struct
type ImportResponse struct {
	Added *float64 `json:"added,omitempty"`

	Failed *float64 `json:"failed,omitempty"`

	FailReasons []ImportResponseFailReasonsItem `json:"failReasons,omitempty"`
}

// UnmarshalImportResponse unmarshals an instance of ImportResponse from the specified map of raw messages.
func UnmarshalImportResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ImportResponse)
	err = core.UnmarshalPrimitive(m, "added", &obj.Added)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "failed", &obj.Failed)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "failReasons", &obj.FailReasons, UnmarshalImportResponseFailReasonsItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PasswordRegexConfigParamsGet : PasswordRegexConfigParamsGet struct
type PasswordRegexConfigParamsGet struct {
	Regex *string `json:"regex,omitempty"`

	Base64EncodedRegex *string `json:"base64_encoded_regex,omitempty"`

	ErrorMessage *string `json:"error_message,omitempty"`
}

// UnmarshalPasswordRegexConfigParamsGet unmarshals an instance of PasswordRegexConfigParamsGet from the specified map of raw messages.
func UnmarshalPasswordRegexConfigParamsGet(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PasswordRegexConfigParamsGet)
	err = core.UnmarshalPrimitive(m, "regex", &obj.Regex)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "base64_encoded_regex", &obj.Base64EncodedRegex)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "error_message", &obj.ErrorMessage)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RedirectUriConfig : RedirectUriConfig struct
type RedirectUriConfig struct {
	RedirectUris []string `json:"redirectUris,omitempty"`

	TrustCloudIAMRedirectUris *bool `json:"trustCloudIAMRedirectUris,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// SetProperty allows the user to set an arbitrary property on an instance of RedirectUriConfig
func (o *RedirectUriConfig) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of RedirectUriConfig
func (o *RedirectUriConfig) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of RedirectUriConfig
func (o *RedirectUriConfig) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of RedirectUriConfig
func (o *RedirectUriConfig) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.RedirectUris != nil {
		m["redirectUris"] = o.RedirectUris
	}
	if o.TrustCloudIAMRedirectUris != nil {
		m["trustCloudIAMRedirectUris"] = o.TrustCloudIAMRedirectUris
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalRedirectUriConfig unmarshals an instance of RedirectUriConfig from the specified map of raw messages.
func UnmarshalRedirectUriConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RedirectUriConfig)
	err = core.UnmarshalPrimitive(m, "redirectUris", &obj.RedirectUris)
	if err != nil {
		return
	}
	delete(m, "redirectUris")
	err = core.UnmarshalPrimitive(m, "trustCloudIAMRedirectUris", &obj.TrustCloudIAMRedirectUris)
	if err != nil {
		return
	}
	delete(m, "trustCloudIAMRedirectUris")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RedirectUriResponse : RedirectUriResponse struct
type RedirectUriResponse struct {
	RedirectUris []string `json:"redirectUris,omitempty"`
}

// UnmarshalRedirectUriResponse unmarshals an instance of RedirectUriResponse from the specified map of raw messages.
func UnmarshalRedirectUriResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RedirectUriResponse)
	err = core.UnmarshalPrimitive(m, "redirectUris", &obj.RedirectUris)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RefreshTokenConfigParams : RefreshTokenConfigParams struct
type RefreshTokenConfigParams struct {
	ExpiresIn *float64 `json:"expires_in" validate:"required"`

	Enabled *bool `json:"enabled" validate:"required"`
}

// NewRefreshTokenConfigParams : Instantiate RefreshTokenConfigParams (Generic Model Constructor)
func (*AppIdManagementV4) NewRefreshTokenConfigParams(expiresIn float64, enabled bool) (model *RefreshTokenConfigParams, err error) {
	model = &RefreshTokenConfigParams{
		ExpiresIn: core.Float64Ptr(expiresIn),
		Enabled: core.BoolPtr(enabled),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalRefreshTokenConfigParams unmarshals an instance of RefreshTokenConfigParams from the specified map of raw messages.
func UnmarshalRefreshTokenConfigParams(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RefreshTokenConfigParams)
	err = core.UnmarshalPrimitive(m, "expires_in", &obj.ExpiresIn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "enabled", &obj.Enabled)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RespCustomEmailDisParams : RespCustomEmailDisParams struct
type RespCustomEmailDisParams struct {
	StatusCode *float64 `json:"statusCode,omitempty"`

	HeadersVar interface{} `json:"headers,omitempty"`
}

// UnmarshalRespCustomEmailDisParams unmarshals an instance of RespCustomEmailDisParams from the specified map of raw messages.
func UnmarshalRespCustomEmailDisParams(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RespCustomEmailDisParams)
	err = core.UnmarshalPrimitive(m, "statusCode", &obj.StatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "headers", &obj.HeadersVar)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RespEmailSettingsTest : RespEmailSettingsTest struct
type RespEmailSettingsTest struct {
	Success *bool `json:"success" validate:"required"`

	DispatcherStatusCode *float64 `json:"dispatcherStatusCode" validate:"required"`

	DispatcherResponse interface{} `json:"dispatcherResponse,omitempty"`
}

// UnmarshalRespEmailSettingsTest unmarshals an instance of RespEmailSettingsTest from the specified map of raw messages.
func UnmarshalRespEmailSettingsTest(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RespEmailSettingsTest)
	err = core.UnmarshalPrimitive(m, "success", &obj.Success)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dispatcherStatusCode", &obj.DispatcherStatusCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "dispatcherResponse", &obj.DispatcherResponse)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RespSMSDisParams : RespSMSDisParams struct
type RespSMSDisParams struct {
	ConfirmationCode *float64 `json:"confirmationCode,omitempty"`

	PhoneNumber *string `json:"phoneNumber,omitempty"`
}

// UnmarshalRespSMSDisParams unmarshals an instance of RespSMSDisParams from the specified map of raw messages.
func UnmarshalRespSMSDisParams(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RespSMSDisParams)
	err = core.UnmarshalPrimitive(m, "confirmationCode", &obj.ConfirmationCode)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "phoneNumber", &obj.PhoneNumber)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SAMLConfigParams : SAMLConfigParams struct
type SAMLConfigParams struct {
	EntityID *string `json:"entityID" validate:"required"`

	SignInURL *string `json:"signInUrl" validate:"required"`

	Certificates []string `json:"certificates" validate:"required"`

	DisplayName *string `json:"displayName,omitempty"`

	AuthnContext *SAMLConfigParamsAuthnContext `json:"authnContext,omitempty"`

	SignRequest *bool `json:"signRequest,omitempty"`

	EncryptResponse *bool `json:"encryptResponse,omitempty"`

	IncludeScoping *bool `json:"includeScoping,omitempty"`

	// Allows users to set arbitrary properties
	additionalProperties map[string]interface{}
}

// NewSAMLConfigParams : Instantiate SAMLConfigParams (Generic Model Constructor)
func (*AppIdManagementV4) NewSAMLConfigParams(entityID string, signInURL string, certificates []string) (model *SAMLConfigParams, err error) {
	model = &SAMLConfigParams{
		EntityID: core.StringPtr(entityID),
		SignInURL: core.StringPtr(signInURL),
		Certificates: certificates,
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// SetProperty allows the user to set an arbitrary property on an instance of SAMLConfigParams
func (o *SAMLConfigParams) SetProperty(key string, value interface{}) {
	if o.additionalProperties == nil {
		o.additionalProperties = make(map[string]interface{})
	}
	o.additionalProperties[key] = value
}

// GetProperty allows the user to retrieve an arbitrary property from an instance of SAMLConfigParams
func (o *SAMLConfigParams) GetProperty(key string) interface{} {
	return o.additionalProperties[key]
}

// GetProperties allows the user to retrieve the map of arbitrary properties from an instance of SAMLConfigParams
func (o *SAMLConfigParams) GetProperties() map[string]interface{} {
	return o.additionalProperties
}

// MarshalJSON performs custom serialization for instances of SAMLConfigParams
func (o *SAMLConfigParams) MarshalJSON() (buffer []byte, err error) {
	m := make(map[string]interface{})
	if len(o.additionalProperties) > 0 {
		for k, v := range o.additionalProperties {
			m[k] = v
		}
	}
	if o.EntityID != nil {
		m["entityID"] = o.EntityID
	}
	if o.SignInURL != nil {
		m["signInUrl"] = o.SignInURL
	}
	if o.Certificates != nil {
		m["certificates"] = o.Certificates
	}
	if o.DisplayName != nil {
		m["displayName"] = o.DisplayName
	}
	if o.AuthnContext != nil {
		m["authnContext"] = o.AuthnContext
	}
	if o.SignRequest != nil {
		m["signRequest"] = o.SignRequest
	}
	if o.EncryptResponse != nil {
		m["encryptResponse"] = o.EncryptResponse
	}
	if o.IncludeScoping != nil {
		m["includeScoping"] = o.IncludeScoping
	}
	buffer, err = json.Marshal(m)
	return
}

// UnmarshalSAMLConfigParams unmarshals an instance of SAMLConfigParams from the specified map of raw messages.
func UnmarshalSAMLConfigParams(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SAMLConfigParams)
	err = core.UnmarshalPrimitive(m, "entityID", &obj.EntityID)
	if err != nil {
		return
	}
	delete(m, "entityID")
	err = core.UnmarshalPrimitive(m, "signInUrl", &obj.SignInURL)
	if err != nil {
		return
	}
	delete(m, "signInUrl")
	err = core.UnmarshalPrimitive(m, "certificates", &obj.Certificates)
	if err != nil {
		return
	}
	delete(m, "certificates")
	err = core.UnmarshalPrimitive(m, "displayName", &obj.DisplayName)
	if err != nil {
		return
	}
	delete(m, "displayName")
	err = core.UnmarshalModel(m, "authnContext", &obj.AuthnContext, UnmarshalSAMLConfigParamsAuthnContext)
	if err != nil {
		return
	}
	delete(m, "authnContext")
	err = core.UnmarshalPrimitive(m, "signRequest", &obj.SignRequest)
	if err != nil {
		return
	}
	delete(m, "signRequest")
	err = core.UnmarshalPrimitive(m, "encryptResponse", &obj.EncryptResponse)
	if err != nil {
		return
	}
	delete(m, "encryptResponse")
	err = core.UnmarshalPrimitive(m, "includeScoping", &obj.IncludeScoping)
	if err != nil {
		return
	}
	delete(m, "includeScoping")
	for k := range m {
		var v interface{}
		e := core.UnmarshalPrimitive(m, k, &v)
		if e != nil {
			err = e
			return
		}
		obj.SetProperty(k, v)
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SAMLResponse : SAMLResponse struct
type SAMLResponse struct {
	IsActive *bool `json:"isActive" validate:"required"`

	Config *SAMLConfigParams `json:"config,omitempty"`
}

// UnmarshalSAMLResponse unmarshals an instance of SAMLResponse from the specified map of raw messages.
func UnmarshalSAMLResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SAMLResponse)
	err = core.UnmarshalPrimitive(m, "isActive", &obj.IsActive)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalSAMLConfigParams)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SAMLResponseWithValidationData : SAMLResponseWithValidationData struct
type SAMLResponseWithValidationData struct {
	IsActive *bool `json:"isActive" validate:"required"`

	Config *SAMLConfigParams `json:"config,omitempty"`

	ValidationData *SAMLResponseWithValidationDataValidationData `json:"validation_data,omitempty"`
}

// UnmarshalSAMLResponseWithValidationData unmarshals an instance of SAMLResponseWithValidationData from the specified map of raw messages.
func UnmarshalSAMLResponseWithValidationData(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SAMLResponseWithValidationData)
	err = core.UnmarshalPrimitive(m, "isActive", &obj.IsActive)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalSAMLConfigParams)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "validation_data", &obj.ValidationData, UnmarshalSAMLResponseWithValidationDataValidationData)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TokenClaimMapping : TokenClaimMapping struct
type TokenClaimMapping struct {
	Source *string `json:"source" validate:"required"`

	SourceClaim *string `json:"sourceClaim,omitempty"`

	DestinationClaim *string `json:"destinationClaim,omitempty"`
}

// Constants associated with the TokenClaimMapping.Source property.
const (
	TokenClaimMapping_Source_AppidCustom = "appid_custom"
	TokenClaimMapping_Source_Attributes = "attributes"
	TokenClaimMapping_Source_CloudDirectory = "cloud_directory"
	TokenClaimMapping_Source_Facebook = "facebook"
	TokenClaimMapping_Source_Google = "google"
	TokenClaimMapping_Source_Ibmid = "ibmid"
	TokenClaimMapping_Source_Roles = "roles"
	TokenClaimMapping_Source_SAML = "saml"
)

// NewTokenClaimMapping : Instantiate TokenClaimMapping (Generic Model Constructor)
func (*AppIdManagementV4) NewTokenClaimMapping(source string) (model *TokenClaimMapping, err error) {
	model = &TokenClaimMapping{
		Source: core.StringPtr(source),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalTokenClaimMapping unmarshals an instance of TokenClaimMapping from the specified map of raw messages.
func UnmarshalTokenClaimMapping(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TokenClaimMapping)
	err = core.UnmarshalPrimitive(m, "source", &obj.Source)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "sourceClaim", &obj.SourceClaim)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "destinationClaim", &obj.DestinationClaim)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TokenConfigParams : TokenConfigParams struct
type TokenConfigParams struct {
	ExpiresIn *float64 `json:"expires_in" validate:"required"`
}

// NewTokenConfigParams : Instantiate TokenConfigParams (Generic Model Constructor)
func (*AppIdManagementV4) NewTokenConfigParams(expiresIn float64) (model *TokenConfigParams, err error) {
	model = &TokenConfigParams{
		ExpiresIn: core.Float64Ptr(expiresIn),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalTokenConfigParams unmarshals an instance of TokenConfigParams from the specified map of raw messages.
func UnmarshalTokenConfigParams(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TokenConfigParams)
	err = core.UnmarshalPrimitive(m, "expires_in", &obj.ExpiresIn)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// TokensConfigResponse : TokensConfigResponse struct
type TokensConfigResponse struct {
	IdTokenClaims []TokenClaimMapping `json:"idTokenClaims,omitempty"`

	AccessTokenClaims []TokenClaimMapping `json:"accessTokenClaims,omitempty"`

	Access []TokenConfigParams `json:"access,omitempty"`

	Refresh []RefreshTokenConfigParams `json:"refresh,omitempty"`

	AnonymousAccess []TokenConfigParams `json:"anonymousAccess,omitempty"`
}

// UnmarshalTokensConfigResponse unmarshals an instance of TokensConfigResponse from the specified map of raw messages.
func UnmarshalTokensConfigResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(TokensConfigResponse)
	err = core.UnmarshalModel(m, "idTokenClaims", &obj.IdTokenClaims, UnmarshalTokenClaimMapping)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "accessTokenClaims", &obj.AccessTokenClaims, UnmarshalTokenClaimMapping)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "access", &obj.Access, UnmarshalTokenConfigParams)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "refresh", &obj.Refresh, UnmarshalRefreshTokenConfigParams)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "anonymousAccess", &obj.AnonymousAccess, UnmarshalTokenConfigParams)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateExtensionConfig : UpdateExtensionConfig struct
type UpdateExtensionConfig struct {
	IsActive *bool `json:"isActive" validate:"required"`

	Config *UpdateExtensionConfigConfig `json:"config,omitempty"`
}

// NewUpdateExtensionConfig : Instantiate UpdateExtensionConfig (Generic Model Constructor)
func (*AppIdManagementV4) NewUpdateExtensionConfig(isActive bool) (model *UpdateExtensionConfig, err error) {
	model = &UpdateExtensionConfig{
		IsActive: core.BoolPtr(isActive),
	}
	err = core.ValidateStruct(model, "required parameters")
	return
}

// UnmarshalUpdateExtensionConfig unmarshals an instance of UpdateExtensionConfig from the specified map of raw messages.
func UnmarshalUpdateExtensionConfig(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateExtensionConfig)
	err = core.UnmarshalPrimitive(m, "isActive", &obj.IsActive)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "config", &obj.Config, UnmarshalUpdateExtensionConfigConfig)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateRolesResponse : UpdateRolesResponse struct
type UpdateRolesResponse struct {
	ID *string `json:"id" validate:"required"`

	Name *string `json:"name" validate:"required"`

	Description *string `json:"description,omitempty"`

	Access []UpdateRolesResponseAccessItem `json:"access" validate:"required"`
}

// UnmarshalUpdateRolesResponse unmarshals an instance of UpdateRolesResponse from the specified map of raw messages.
func UnmarshalUpdateRolesResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateRolesResponse)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "access", &obj.Access, UnmarshalUpdateRolesResponseAccessItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UserSearchResponse : UserSearchResponse struct
type UserSearchResponse struct {
	TotalResults *float64 `json:"totalResults,omitempty"`

	ItemsPerPage *float64 `json:"itemsPerPage,omitempty"`

	RequestOptions *UserSearchResponseRequestOptions `json:"requestOptions,omitempty"`

	Users []UserSearchResponseUsersItem `json:"users,omitempty"`
}

// UnmarshalUserSearchResponse unmarshals an instance of UserSearchResponse from the specified map of raw messages.
func UnmarshalUserSearchResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UserSearchResponse)
	err = core.UnmarshalPrimitive(m, "totalResults", &obj.TotalResults)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "itemsPerPage", &obj.ItemsPerPage)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "requestOptions", &obj.RequestOptions, UnmarshalUserSearchResponseRequestOptions)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "users", &obj.Users, UnmarshalUserSearchResponseUsersItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
