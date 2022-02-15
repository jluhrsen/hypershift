package authorization

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// DenyAssignmentsClient is the client for the DenyAssignments methods of the Authorization service.
type DenyAssignmentsClient struct {
	BaseClient
}

// NewDenyAssignmentsClient creates an instance of the DenyAssignmentsClient client.
func NewDenyAssignmentsClient(subscriptionID string) DenyAssignmentsClient {
	return NewDenyAssignmentsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewDenyAssignmentsClientWithBaseURI creates an instance of the DenyAssignmentsClient client using a custom endpoint.
// Use this when interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewDenyAssignmentsClientWithBaseURI(baseURI string, subscriptionID string) DenyAssignmentsClient {
	return DenyAssignmentsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Get get the specified deny assignment.
// Parameters:
// scope - the scope of the deny assignment.
// denyAssignmentID - the ID of the deny assignment to get.
func (client DenyAssignmentsClient) Get(ctx context.Context, scope string, denyAssignmentID string) (result DenyAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DenyAssignmentsClient.Get")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetPreparer(ctx, scope, denyAssignmentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "Get", resp, "Failure responding to request")
		return
	}

	return
}

// GetPreparer prepares the Get request.
func (client DenyAssignmentsClient) GetPreparer(ctx context.Context, scope string, denyAssignmentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"denyAssignmentId": autorest.Encode("path", denyAssignmentID),
		"scope":            scope,
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/denyAssignments/{denyAssignmentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client DenyAssignmentsClient) GetSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client DenyAssignmentsClient) GetResponder(resp *http.Response) (result DenyAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// GetByID gets a deny assignment by ID.
// Parameters:
// denyAssignmentID - the fully qualified deny assignment ID. For example, use the format,
// /subscriptions/{guid}/providers/Microsoft.Authorization/denyAssignments/{denyAssignmentId} for subscription
// level deny assignments, or /providers/Microsoft.Authorization/denyAssignments/{denyAssignmentId} for tenant
// level deny assignments.
func (client DenyAssignmentsClient) GetByID(ctx context.Context, denyAssignmentID string) (result DenyAssignment, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DenyAssignmentsClient.GetByID")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	req, err := client.GetByIDPreparer(ctx, denyAssignmentID)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "GetByID", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetByIDSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "GetByID", resp, "Failure sending request")
		return
	}

	result, err = client.GetByIDResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "GetByID", resp, "Failure responding to request")
		return
	}

	return
}

// GetByIDPreparer prepares the GetByID request.
func (client DenyAssignmentsClient) GetByIDPreparer(ctx context.Context, denyAssignmentID string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"denyAssignmentId": denyAssignmentID,
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{denyAssignmentId}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetByIDSender sends the GetByID request. The method will close the
// http.Response Body if it receives an error.
func (client DenyAssignmentsClient) GetByIDSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// GetByIDResponder handles the response to the GetByID request. The method always
// closes the http.Response Body.
func (client DenyAssignmentsClient) GetByIDResponder(resp *http.Response) (result DenyAssignment, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets all deny assignments for the subscription.
// Parameters:
// filter - the filter to apply on the operation. Use $filter=atScope() to return all deny assignments at or
// above the scope. Use $filter=denyAssignmentName eq '{name}' to search deny assignments by name at specified
// scope. Use $filter=principalId eq '{id}' to return all deny assignments at, above and below the scope for
// the specified principal. Use $filter=gdprExportPrincipalId eq '{id}' to return all deny assignments at,
// above and below the scope for the specified principal. This filter is different from the principalId filter
// as it returns not only those deny assignments that contain the specified principal is the Principals list
// but also those deny assignments that contain the specified principal is the ExcludePrincipals list.
// Additionally, when gdprExportPrincipalId filter is used, only the deny assignment name and description
// properties are returned.
func (client DenyAssignmentsClient) List(ctx context.Context, filter string) (result DenyAssignmentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DenyAssignmentsClient.List")
		defer func() {
			sc := -1
			if result.dalr.Response.Response != nil {
				sc = result.dalr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}}}); err != nil {
		return result, validation.NewError("authorization.DenyAssignmentsClient", "List", err.Error())
	}

	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.dalr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "List", resp, "Failure sending request")
		return
	}

	result.dalr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "List", resp, "Failure responding to request")
		return
	}
	if result.dalr.hasNextLink() && result.dalr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListPreparer prepares the List request.
func (client DenyAssignmentsClient) ListPreparer(ctx context.Context, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.Authorization/denyAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client DenyAssignmentsClient) ListSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client DenyAssignmentsClient) ListResponder(resp *http.Response) (result DenyAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client DenyAssignmentsClient) listNextResults(ctx context.Context, lastResults DenyAssignmentListResult) (result DenyAssignmentListResult, err error) {
	req, err := lastResults.denyAssignmentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client DenyAssignmentsClient) ListComplete(ctx context.Context, filter string) (result DenyAssignmentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DenyAssignmentsClient.List")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.List(ctx, filter)
	return
}

// ListForResource gets deny assignments for a resource.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// resourceProviderNamespace - the namespace of the resource provider.
// parentResourcePath - the parent resource identity.
// resourceType - the resource type of the resource.
// resourceName - the name of the resource to get deny assignments for.
// filter - the filter to apply on the operation. Use $filter=atScope() to return all deny assignments at or
// above the scope. Use $filter=denyAssignmentName eq '{name}' to search deny assignments by name at specified
// scope. Use $filter=principalId eq '{id}' to return all deny assignments at, above and below the scope for
// the specified principal. Use $filter=gdprExportPrincipalId eq '{id}' to return all deny assignments at,
// above and below the scope for the specified principal. This filter is different from the principalId filter
// as it returns not only those deny assignments that contain the specified principal is the Principals list
// but also those deny assignments that contain the specified principal is the ExcludePrincipals list.
// Additionally, when gdprExportPrincipalId filter is used, only the deny assignment name and description
// properties are returned.
func (client DenyAssignmentsClient) ListForResource(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result DenyAssignmentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DenyAssignmentsClient.ListForResource")
		defer func() {
			sc := -1
			if result.dalr.Response.Response != nil {
				sc = result.dalr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("authorization.DenyAssignmentsClient", "ListForResource", err.Error())
	}

	result.fn = client.listForResourceNextResults
	req, err := client.ListForResourcePreparer(ctx, resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "ListForResource", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForResourceSender(req)
	if err != nil {
		result.dalr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "ListForResource", resp, "Failure sending request")
		return
	}

	result.dalr, err = client.ListForResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "ListForResource", resp, "Failure responding to request")
		return
	}
	if result.dalr.hasNextLink() && result.dalr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListForResourcePreparer prepares the ListForResource request.
func (client DenyAssignmentsClient) ListForResourcePreparer(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"parentResourcePath":        parentResourcePath,
		"resourceGroupName":         autorest.Encode("path", resourceGroupName),
		"resourceName":              autorest.Encode("path", resourceName),
		"resourceProviderNamespace": autorest.Encode("path", resourceProviderNamespace),
		"resourceType":              resourceType,
		"subscriptionId":            autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{parentResourcePath}/{resourceType}/{resourceName}/providers/Microsoft.Authorization/denyAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListForResourceSender sends the ListForResource request. The method will close the
// http.Response Body if it receives an error.
func (client DenyAssignmentsClient) ListForResourceSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListForResourceResponder handles the response to the ListForResource request. The method always
// closes the http.Response Body.
func (client DenyAssignmentsClient) ListForResourceResponder(resp *http.Response) (result DenyAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listForResourceNextResults retrieves the next set of results, if any.
func (client DenyAssignmentsClient) listForResourceNextResults(ctx context.Context, lastResults DenyAssignmentListResult) (result DenyAssignmentListResult, err error) {
	req, err := lastResults.denyAssignmentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "listForResourceNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListForResourceSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "listForResourceNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListForResourceResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "listForResourceNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListForResourceComplete enumerates all values, automatically crossing page boundaries as required.
func (client DenyAssignmentsClient) ListForResourceComplete(ctx context.Context, resourceGroupName string, resourceProviderNamespace string, parentResourcePath string, resourceType string, resourceName string, filter string) (result DenyAssignmentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DenyAssignmentsClient.ListForResource")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListForResource(ctx, resourceGroupName, resourceProviderNamespace, parentResourcePath, resourceType, resourceName, filter)
	return
}

// ListForResourceGroup gets deny assignments for a resource group.
// Parameters:
// resourceGroupName - the name of the resource group. The name is case insensitive.
// filter - the filter to apply on the operation. Use $filter=atScope() to return all deny assignments at or
// above the scope. Use $filter=denyAssignmentName eq '{name}' to search deny assignments by name at specified
// scope. Use $filter=principalId eq '{id}' to return all deny assignments at, above and below the scope for
// the specified principal. Use $filter=gdprExportPrincipalId eq '{id}' to return all deny assignments at,
// above and below the scope for the specified principal. This filter is different from the principalId filter
// as it returns not only those deny assignments that contain the specified principal is the Principals list
// but also those deny assignments that contain the specified principal is the ExcludePrincipals list.
// Additionally, when gdprExportPrincipalId filter is used, only the deny assignment name and description
// properties are returned.
func (client DenyAssignmentsClient) ListForResourceGroup(ctx context.Context, resourceGroupName string, filter string) (result DenyAssignmentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DenyAssignmentsClient.ListForResourceGroup")
		defer func() {
			sc := -1
			if result.dalr.Response.Response != nil {
				sc = result.dalr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: client.SubscriptionID,
			Constraints: []validation.Constraint{{Target: "client.SubscriptionID", Name: validation.MinLength, Rule: 1, Chain: nil}}},
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 90, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `^[-\w\._\(\)]+$`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("authorization.DenyAssignmentsClient", "ListForResourceGroup", err.Error())
	}

	result.fn = client.listForResourceGroupNextResults
	req, err := client.ListForResourceGroupPreparer(ctx, resourceGroupName, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "ListForResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForResourceGroupSender(req)
	if err != nil {
		result.dalr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "ListForResourceGroup", resp, "Failure sending request")
		return
	}

	result.dalr, err = client.ListForResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "ListForResourceGroup", resp, "Failure responding to request")
		return
	}
	if result.dalr.hasNextLink() && result.dalr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListForResourceGroupPreparer prepares the ListForResourceGroup request.
func (client DenyAssignmentsClient) ListForResourceGroupPreparer(ctx context.Context, resourceGroupName string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Authorization/denyAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListForResourceGroupSender sends the ListForResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client DenyAssignmentsClient) ListForResourceGroupSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// ListForResourceGroupResponder handles the response to the ListForResourceGroup request. The method always
// closes the http.Response Body.
func (client DenyAssignmentsClient) ListForResourceGroupResponder(resp *http.Response) (result DenyAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listForResourceGroupNextResults retrieves the next set of results, if any.
func (client DenyAssignmentsClient) listForResourceGroupNextResults(ctx context.Context, lastResults DenyAssignmentListResult) (result DenyAssignmentListResult, err error) {
	req, err := lastResults.denyAssignmentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "listForResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListForResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "listForResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListForResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "listForResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListForResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client DenyAssignmentsClient) ListForResourceGroupComplete(ctx context.Context, resourceGroupName string, filter string) (result DenyAssignmentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DenyAssignmentsClient.ListForResourceGroup")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListForResourceGroup(ctx, resourceGroupName, filter)
	return
}

// ListForScope gets deny assignments for a scope.
// Parameters:
// scope - the scope of the deny assignments.
// filter - the filter to apply on the operation. Use $filter=atScope() to return all deny assignments at or
// above the scope. Use $filter=denyAssignmentName eq '{name}' to search deny assignments by name at specified
// scope. Use $filter=principalId eq '{id}' to return all deny assignments at, above and below the scope for
// the specified principal. Use $filter=gdprExportPrincipalId eq '{id}' to return all deny assignments at,
// above and below the scope for the specified principal. This filter is different from the principalId filter
// as it returns not only those deny assignments that contain the specified principal is the Principals list
// but also those deny assignments that contain the specified principal is the ExcludePrincipals list.
// Additionally, when gdprExportPrincipalId filter is used, only the deny assignment name and description
// properties are returned.
func (client DenyAssignmentsClient) ListForScope(ctx context.Context, scope string, filter string) (result DenyAssignmentListResultPage, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DenyAssignmentsClient.ListForScope")
		defer func() {
			sc := -1
			if result.dalr.Response.Response != nil {
				sc = result.dalr.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.fn = client.listForScopeNextResults
	req, err := client.ListForScopePreparer(ctx, scope, filter)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "ListForScope", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListForScopeSender(req)
	if err != nil {
		result.dalr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "ListForScope", resp, "Failure sending request")
		return
	}

	result.dalr, err = client.ListForScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "ListForScope", resp, "Failure responding to request")
		return
	}
	if result.dalr.hasNextLink() && result.dalr.IsEmpty() {
		err = result.NextWithContext(ctx)
		return
	}

	return
}

// ListForScopePreparer prepares the ListForScope request.
func (client DenyAssignmentsClient) ListForScopePreparer(ctx context.Context, scope string, filter string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"scope": scope,
	}

	const APIVersion = "2018-07-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(filter) > 0 {
		queryParameters["$filter"] = autorest.Encode("query", filter)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/{scope}/providers/Microsoft.Authorization/denyAssignments", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListForScopeSender sends the ListForScope request. The method will close the
// http.Response Body if it receives an error.
func (client DenyAssignmentsClient) ListForScopeSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, autorest.DoRetryForStatusCodes(client.RetryAttempts, client.RetryDuration, autorest.StatusCodesForRetry...))
}

// ListForScopeResponder handles the response to the ListForScope request. The method always
// closes the http.Response Body.
func (client DenyAssignmentsClient) ListForScopeResponder(resp *http.Response) (result DenyAssignmentListResult, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listForScopeNextResults retrieves the next set of results, if any.
func (client DenyAssignmentsClient) listForScopeNextResults(ctx context.Context, lastResults DenyAssignmentListResult) (result DenyAssignmentListResult, err error) {
	req, err := lastResults.denyAssignmentListResultPreparer(ctx)
	if err != nil {
		return result, autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "listForScopeNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListForScopeSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "listForScopeNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListForScopeResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "authorization.DenyAssignmentsClient", "listForScopeNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListForScopeComplete enumerates all values, automatically crossing page boundaries as required.
func (client DenyAssignmentsClient) ListForScopeComplete(ctx context.Context, scope string, filter string) (result DenyAssignmentListResultIterator, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/DenyAssignmentsClient.ListForScope")
		defer func() {
			sc := -1
			if result.Response().Response.Response != nil {
				sc = result.page.Response().Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	result.page, err = client.ListForScope(ctx, scope, filter)
	return
}
