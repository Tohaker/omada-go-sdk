# SslVpnResourceGroupModifyOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ResourcesList** | Pointer to **[]string** | Resources ID list of the SSL VPN resource group. Resource can be created using &#39;Create SSL VPN resource&#39; interface, and resource ID can be obtained from &#39;Get resource list for SSL VPN server&#39; interface. | [optional] 

## Methods

### NewSslVpnResourceGroupModifyOpenApiVO

`func NewSslVpnResourceGroupModifyOpenApiVO() *SslVpnResourceGroupModifyOpenApiVO`

NewSslVpnResourceGroupModifyOpenApiVO instantiates a new SslVpnResourceGroupModifyOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnResourceGroupModifyOpenApiVOWithDefaults

`func NewSslVpnResourceGroupModifyOpenApiVOWithDefaults() *SslVpnResourceGroupModifyOpenApiVO`

NewSslVpnResourceGroupModifyOpenApiVOWithDefaults instantiates a new SslVpnResourceGroupModifyOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetResourcesList

`func (o *SslVpnResourceGroupModifyOpenApiVO) GetResourcesList() []string`

GetResourcesList returns the ResourcesList field if non-nil, zero value otherwise.

### GetResourcesListOk

`func (o *SslVpnResourceGroupModifyOpenApiVO) GetResourcesListOk() (*[]string, bool)`

GetResourcesListOk returns a tuple with the ResourcesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesList

`func (o *SslVpnResourceGroupModifyOpenApiVO) SetResourcesList(v []string)`

SetResourcesList sets ResourcesList field to given value.

### HasResourcesList

`func (o *SslVpnResourceGroupModifyOpenApiVO) HasResourcesList() bool`

HasResourcesList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


