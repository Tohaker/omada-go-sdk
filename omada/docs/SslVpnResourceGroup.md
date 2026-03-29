# SslVpnResourceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the SSL VPN resource group | [optional] 
**Name** | **string** | Name of the SSL VPN resource group | 
**ResourcesList** | Pointer to [**[]SslVpnResourceBriefInfo**](SslVpnResourceBriefInfo.md) | Resources list of the SSL VPN resource group | [optional] 

## Methods

### NewSslVpnResourceGroup

`func NewSslVpnResourceGroup(name string, ) *SslVpnResourceGroup`

NewSslVpnResourceGroup instantiates a new SslVpnResourceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSslVpnResourceGroupWithDefaults

`func NewSslVpnResourceGroupWithDefaults() *SslVpnResourceGroup`

NewSslVpnResourceGroupWithDefaults instantiates a new SslVpnResourceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SslVpnResourceGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SslVpnResourceGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SslVpnResourceGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SslVpnResourceGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SslVpnResourceGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SslVpnResourceGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SslVpnResourceGroup) SetName(v string)`

SetName sets Name field to given value.


### GetResourcesList

`func (o *SslVpnResourceGroup) GetResourcesList() []SslVpnResourceBriefInfo`

GetResourcesList returns the ResourcesList field if non-nil, zero value otherwise.

### GetResourcesListOk

`func (o *SslVpnResourceGroup) GetResourcesListOk() (*[]SslVpnResourceBriefInfo, bool)`

GetResourcesListOk returns a tuple with the ResourcesList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourcesList

`func (o *SslVpnResourceGroup) SetResourcesList(v []SslVpnResourceBriefInfo)`

SetResourcesList sets ResourcesList field to given value.

### HasResourcesList

`func (o *SslVpnResourceGroup) HasResourcesList() bool`

HasResourcesList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


