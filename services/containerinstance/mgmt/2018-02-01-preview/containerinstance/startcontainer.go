package containerinstance

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"net/http"
)

// StartContainerClient is the client for the StartContainer methods of the Containerinstance service.
type StartContainerClient struct {
	BaseClient
}

// NewStartContainerClient creates an instance of the StartContainerClient client.
func NewStartContainerClient(subscriptionID string) StartContainerClient {
	return NewStartContainerClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewStartContainerClientWithBaseURI creates an instance of the StartContainerClient client.
func NewStartContainerClientWithBaseURI(baseURI string, subscriptionID string) StartContainerClient {
	return StartContainerClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Exec starts the exec command for a specified container instance in a specified resource group and container group.
//
// resourceGroupName is the name of the resource group. containerGroupName is the name of the container group.
// containerName is the name of the container instance. containerExecRequest is the request for the exec command.
func (client StartContainerClient) Exec(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, containerExecRequest ContainerExecRequest) (result ContainerExecResponse, err error) {
	req, err := client.ExecPreparer(ctx, resourceGroupName, containerGroupName, containerName, containerExecRequest)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.StartContainerClient", "Exec", nil, "Failure preparing request")
		return
	}

	resp, err := client.ExecSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerinstance.StartContainerClient", "Exec", resp, "Failure sending request")
		return
	}

	result, err = client.ExecResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerinstance.StartContainerClient", "Exec", resp, "Failure responding to request")
	}

	return
}

// ExecPreparer prepares the Exec request.
func (client StartContainerClient) ExecPreparer(ctx context.Context, resourceGroupName string, containerGroupName string, containerName string, containerExecRequest ContainerExecRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerGroupName": autorest.Encode("path", containerGroupName),
		"containerName":      autorest.Encode("path", containerName),
		"resourceGroupName":  autorest.Encode("path", resourceGroupName),
		"subscriptionId":     autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2018-02-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerInstance/containerGroups/{containerGroupName}/containers/{containerName}/exec", pathParameters),
		autorest.WithJSON(containerExecRequest),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ExecSender sends the Exec request. The method will close the
// http.Response Body if it receives an error.
func (client StartContainerClient) ExecSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ExecResponder handles the response to the Exec request. The method always
// closes the http.Response Body.
func (client StartContainerClient) ExecResponder(resp *http.Response) (result ContainerExecResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
