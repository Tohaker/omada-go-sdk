# SslVpnResourceGroupConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Name of the SSL VPN resource group should contain 1 to 20 characters. | 
**ResourcesList** | Pointer to **[]string** | Resources ID list of the SSL VPN resource group | [optional] 

## Methods

### NewSslVpnResourceGroupConfigOpenApiVO

`func NewSslVpnResourceGroupConfigOpenApiVO(name string, ) *SslVpnResourceGroupConfigOpenApiVO`

NewSslVpnResourceGroupConfigOpenApiVO instantiates a new SslVpnResourceGroupConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnResourceGroupConfigOpenApiVOWithDefaults

`func NewSslVpnResourceGroupConfigOpenApiVOWithDefaults() *SslVpnResourceGroupConfigOpenApiVO`

NewSslVpnResourceGroupConfigOpenApiVOWithDefaults instantiates a new SslVpnResourceGroupConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SslVpnResourceGroupConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SslVpnResourceGroupConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SslVpnResourceGroupConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetResourcesList

`func (o *SslVpnResourceGroupConfigOpenApiVO) GetResourcesList() []string`

GetResourcesList returns the ResourcesList field if non-nil, zero value otherwise.

### GetResourcesListOk

`func (o *SslVpnResourceGroupConfigOpenApiVO) GetResourcesListOk() (*[]string, bool)`

GetResourcesListOk returns a tuple with the ResourcesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesList

`func (o *SslVpnResourceGroupConfigOpenApiVO) SetResourcesList(v []string)`

SetResourcesList sets ResourcesList field to given value.

### HasResourcesList

`func (o *SslVpnResourceGroupConfigOpenApiVO) HasResourcesList() bool`

HasResourcesList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


