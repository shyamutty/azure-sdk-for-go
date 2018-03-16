package network

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
	"bytes"
	"context"
	"encoding/json"
	"github.com/Azure/azure-pipeline-go/pipeline"
	"io/ioutil"
	"net/http"
)

// PublicIPAddressesClient is the network Client
type PublicIPAddressesClient struct {
	ManagementClient
}

// NewPublicIPAddressesClient creates an instance of the PublicIPAddressesClient client.
func NewPublicIPAddressesClient(p pipeline.Pipeline) PublicIPAddressesClient {
	return PublicIPAddressesClient{NewManagementClient(p)}
}

// CreateOrUpdate creates or updates a static or dynamic public IP address. This method may poll for completion.
// Polling can be canceled by passing the cancel channel argument. The channel will be used to cancel polling and any
// outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. publicIPAddressName is the name of the public IP address.
// parameters is parameters supplied to the create or update public IP address operation.
func (client PublicIPAddressesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress) (*PublicIPAddress, error) {
	if err := validate([]validation{
		{targetValue: parameters,
			constraints: []constraint{{target: "parameters.PublicIPAddressPropertiesFormat", name: null, rule: false,
				chain: []constraint{{target: "parameters.PublicIPAddressPropertiesFormat.IPConfiguration", name: null, rule: false,
					chain: []constraint{{target: "parameters.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat", name: null, rule: false,
						chain: []constraint{{target: "parameters.PublicIPAddressPropertiesFormat.IPConfiguration.IPConfigurationPropertiesFormat.PublicIPAddress", name: null, rule: false, chain: nil}}},
					}},
				}}}}}); err != nil {
		return nil, err
	}
	req, err := client.createOrUpdatePreparer(resourceGroupName, publicIPAddressName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.createOrUpdateResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*PublicIPAddress), err
}

// createOrUpdatePreparer prepares the CreateOrUpdate request.
func (client PublicIPAddressesClient) createOrUpdatePreparer(resourceGroupName string, publicIPAddressName string, parameters PublicIPAddress) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	req, err := pipeline.NewRequest("PUT", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2018-01-01")
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(parameters)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// createOrUpdateResponder handles the response to the CreateOrUpdate request.
func (client PublicIPAddressesClient) createOrUpdateResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusCreated, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &PublicIPAddress{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// Delete deletes the specified public IP address. This method may poll for completion. Polling can be canceled by
// passing the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. publicIPAddressName is the name of the subnet.
func (client PublicIPAddressesClient) Delete(ctx context.Context, resourceGroupName string, publicIPAddressName string) (*http.Response, error) {
	req, err := client.deletePreparer(resourceGroupName, publicIPAddressName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.deleteResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.Response(), err
}

// deletePreparer prepares the Delete request.
func (client PublicIPAddressesClient) deletePreparer(resourceGroupName string, publicIPAddressName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	req, err := pipeline.NewRequest("DELETE", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2018-01-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// deleteResponder handles the response to the Delete request.
func (client PublicIPAddressesClient) deleteResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusNoContent, http.StatusAccepted, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	resp.Response().Body.Close()
	return resp, err
}

// Get gets the specified public IP address in a specified resource group.
//
// resourceGroupName is the name of the resource group. publicIPAddressName is the name of the subnet. expand is
// expands referenced resources.
func (client PublicIPAddressesClient) Get(ctx context.Context, resourceGroupName string, publicIPAddressName string, expand *string) (*PublicIPAddress, error) {
	req, err := client.getPreparer(resourceGroupName, publicIPAddressName, expand)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*PublicIPAddress), err
}

// getPreparer prepares the Get request.
func (client PublicIPAddressesClient) getPreparer(resourceGroupName string, publicIPAddressName string, expand *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2018-01-01")
	if expand != nil {
		params.Set("$expand", *expand)
	}
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getResponder handles the response to the Get request.
func (client PublicIPAddressesClient) getResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &PublicIPAddress{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// GetVirtualMachineScaleSetPublicIPAddress get the specified public IP address in a virtual machine scale set.
//
// resourceGroupName is the name of the resource group. virtualMachineScaleSetName is the name of the virtual machine
// scale set. virtualmachineIndex is the virtual machine index. networkInterfaceName is the name of the network
// interface. IPConfigurationName is the name of the IP configuration. publicIPAddressName is the name of the public IP
// Address. expand is expands referenced resources.
func (client PublicIPAddressesClient) GetVirtualMachineScaleSetPublicIPAddress(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string, publicIPAddressName string, expand *string) (*PublicIPAddress, error) {
	req, err := client.getVirtualMachineScaleSetPublicIPAddressPreparer(resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName, IPConfigurationName, publicIPAddressName, expand)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.getVirtualMachineScaleSetPublicIPAddressResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*PublicIPAddress), err
}

// getVirtualMachineScaleSetPublicIPAddressPreparer prepares the GetVirtualMachineScaleSetPublicIPAddress request.
func (client PublicIPAddressesClient) getVirtualMachineScaleSetPublicIPAddressPreparer(resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string, publicIPAddressName string, expand *string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses/{publicIpAddressName}"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-03-30")
	if expand != nil {
		params.Set("$expand", *expand)
	}
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// getVirtualMachineScaleSetPublicIPAddressResponder handles the response to the GetVirtualMachineScaleSetPublicIPAddress request.
func (client PublicIPAddressesClient) getVirtualMachineScaleSetPublicIPAddressResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &PublicIPAddress{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// List gets all public IP addresses in a resource group.
//
// resourceGroupName is the name of the resource group.
func (client PublicIPAddressesClient) List(ctx context.Context, resourceGroupName string) (*PublicIPAddressListResult, error) {
	req, err := client.listPreparer(resourceGroupName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*PublicIPAddressListResult), err
}

// listPreparer prepares the List request.
func (client PublicIPAddressesClient) listPreparer(resourceGroupName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2018-01-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listResponder handles the response to the List request.
func (client PublicIPAddressesClient) listResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &PublicIPAddressListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// ListAll gets all the public IP addresses in a subscription.
func (client PublicIPAddressesClient) ListAll(ctx context.Context) (*PublicIPAddressListResult, error) {
	req, err := client.listAllPreparer()
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listAllResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*PublicIPAddressListResult), err
}

// listAllPreparer prepares the ListAll request.
func (client PublicIPAddressesClient) listAllPreparer() (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/providers/Microsoft.Network/publicIPAddresses"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2018-01-01")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listAllResponder handles the response to the ListAll request.
func (client PublicIPAddressesClient) listAllResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &PublicIPAddressListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// ListVirtualMachineScaleSetPublicIPAddresses gets information about all public IP addresses on a virtual machine
// scale set level.
//
// resourceGroupName is the name of the resource group. virtualMachineScaleSetName is the name of the virtual machine
// scale set.
func (client PublicIPAddressesClient) ListVirtualMachineScaleSetPublicIPAddresses(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string) (*PublicIPAddressListResult, error) {
	req, err := client.listVirtualMachineScaleSetPublicIPAddressesPreparer(resourceGroupName, virtualMachineScaleSetName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listVirtualMachineScaleSetPublicIPAddressesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*PublicIPAddressListResult), err
}

// listVirtualMachineScaleSetPublicIPAddressesPreparer prepares the ListVirtualMachineScaleSetPublicIPAddresses request.
func (client PublicIPAddressesClient) listVirtualMachineScaleSetPublicIPAddressesPreparer(resourceGroupName string, virtualMachineScaleSetName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/publicipaddresses"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-03-30")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listVirtualMachineScaleSetPublicIPAddressesResponder handles the response to the ListVirtualMachineScaleSetPublicIPAddresses request.
func (client PublicIPAddressesClient) listVirtualMachineScaleSetPublicIPAddressesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &PublicIPAddressListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// ListVirtualMachineScaleSetVMPublicIPAddresses gets information about all public IP addresses in a virtual machine IP
// configuration in a virtual machine scale set.
//
// resourceGroupName is the name of the resource group. virtualMachineScaleSetName is the name of the virtual machine
// scale set. virtualmachineIndex is the virtual machine index. networkInterfaceName is the network interface name.
// IPConfigurationName is the IP configuration name.
func (client PublicIPAddressesClient) ListVirtualMachineScaleSetVMPublicIPAddresses(ctx context.Context, resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string) (*PublicIPAddressListResult, error) {
	req, err := client.listVirtualMachineScaleSetVMPublicIPAddressesPreparer(resourceGroupName, virtualMachineScaleSetName, virtualmachineIndex, networkInterfaceName, IPConfigurationName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.listVirtualMachineScaleSetVMPublicIPAddressesResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*PublicIPAddressListResult), err
}

// listVirtualMachineScaleSetVMPublicIPAddressesPreparer prepares the ListVirtualMachineScaleSetVMPublicIPAddresses request.
func (client PublicIPAddressesClient) listVirtualMachineScaleSetVMPublicIPAddressesPreparer(resourceGroupName string, virtualMachineScaleSetName string, virtualmachineIndex string, networkInterfaceName string, IPConfigurationName string) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{virtualMachineScaleSetName}/virtualMachines/{virtualmachineIndex}/networkInterfaces/{networkInterfaceName}/ipconfigurations/{ipConfigurationName}/publicipaddresses"
	req, err := pipeline.NewRequest("GET", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2017-03-30")
	req.URL.RawQuery = params.Encode()
	return req, nil
}

// listVirtualMachineScaleSetVMPublicIPAddressesResponder handles the response to the ListVirtualMachineScaleSetVMPublicIPAddresses request.
func (client PublicIPAddressesClient) listVirtualMachineScaleSetVMPublicIPAddressesResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &PublicIPAddressListResult{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}

// UpdateTags updates public IP address tags. This method may poll for completion. Polling can be canceled by passing
// the cancel channel argument. The channel will be used to cancel polling and any outstanding HTTP requests.
//
// resourceGroupName is the name of the resource group. publicIPAddressName is the name of the public IP address.
// parameters is parameters supplied to update public IP address tags.
func (client PublicIPAddressesClient) UpdateTags(ctx context.Context, resourceGroupName string, publicIPAddressName string, parameters TagsObject) (*PublicIPAddress, error) {
	req, err := client.updateTagsPreparer(resourceGroupName, publicIPAddressName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(ctx, responderPolicyFactory{responder: client.updateTagsResponder}, req)
	if err != nil {
		return nil, err
	}
	return resp.(*PublicIPAddress), err
}

// updateTagsPreparer prepares the UpdateTags request.
func (client PublicIPAddressesClient) updateTagsPreparer(resourceGroupName string, publicIPAddressName string, parameters TagsObject) (pipeline.Request, error) {
	u := client.url
	u.Path = "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/publicIPAddresses/{publicIpAddressName}"
	req, err := pipeline.NewRequest("PATCH", u, nil)
	if err != nil {
		return req, pipeline.NewError(err, "failed to create request")
	}
	params := req.URL.Query()
	params.Set("api-version", "2018-01-01")
	req.URL.RawQuery = params.Encode()
	b, err := json.Marshal(parameters)
	if err != nil {
		return req, pipeline.NewError(err, "failed to marshal request body")
	}
	req.Header.Set("Content-Type", "application/json")
	err = req.SetBody(bytes.NewReader(b))
	if err != nil {
		return req, pipeline.NewError(err, "failed to set request body")
	}
	return req, nil
}

// updateTagsResponder handles the response to the UpdateTags request.
func (client PublicIPAddressesClient) updateTagsResponder(resp pipeline.Response) (pipeline.Response, error) {
	err := validateResponse(resp, http.StatusOK)
	if resp == nil {
		return nil, err
	}
	result := &PublicIPAddress{rawResponse: resp.Response()}
	if err != nil {
		return result, err
	}
	defer resp.Response().Body.Close()
	b, err := ioutil.ReadAll(resp.Response().Body)
	if err != nil {
		return result, NewResponseError(err, resp.Response(), "failed to read response body")
	}
	if len(b) > 0 {
		err = json.Unmarshal(b, result)
		if err != nil {
			return result, NewResponseError(err, resp.Response(), "failed to unmarshal response body")
		}
	}
	return result, nil
}
