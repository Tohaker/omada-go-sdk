# UpnpSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | Whether to enable UPnP | 
**NetworkIds** | Pointer to **[]string** | This field represents LAN Network ID. LAN Network ID can be obtained from \&quot;Get all \&quot;single\&quot;/\&quot;multi\&quot; interface lan network\&quot; interface. | [optional] 
**WanPortIds** | Pointer to **[]string** | This field represents WAN Port ID, WAN Port ID can be obtained from \&quot;Get internet basic info\&quot; interface. | [optional] 

## Methods

### NewUpnpSettingOpenApiVO

`func NewUpnpSettingOpenApiVO(enable bool, ) *UpnpSettingOpenApiVO`

NewUpnpSettingOpenApiVO instantiates a new UpnpSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpnpSettingOpenApiVOWithDefaults

`func NewUpnpSettingOpenApiVOWithDefaults() *UpnpSettingOpenApiVO`

NewUpnpSettingOpenApiVOWithDefaults instantiates a new UpnpSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *UpnpSettingOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *UpnpSettingOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *UpnpSettingOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetNetworkIds

`func (o *UpnpSettingOpenApiVO) GetNetworkIds() []string`

GetNetworkIds returns the NetworkIds field if non-nil, zero value otherwise.

### GetNetworkIdsOk

`func (o *UpnpSettingOpenApiVO) GetNetworkIdsOk() (*[]string, bool)`

GetNetworkIdsOk returns a tuple with the NetworkIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIds

`func (o *UpnpSettingOpenApiVO) SetNetworkIds(v []string)`

SetNetworkIds sets NetworkIds field to given value.

### HasNetworkIds

`func (o *UpnpSettingOpenApiVO) HasNetworkIds() bool`

HasNetworkIds returns a boolean if a field has been set.

### GetWanPortIds

`func (o *UpnpSettingOpenApiVO) GetWanPortIds() []string`

GetWanPortIds returns the WanPortIds field if non-nil, zero value otherwise.

### GetWanPortIdsOk

`func (o *UpnpSettingOpenApiVO) GetWanPortIdsOk() (*[]string, bool)`

GetWanPortIdsOk returns a tuple with the WanPortIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortIds

`func (o *UpnpSettingOpenApiVO) SetWanPortIds(v []string)`

SetWanPortIds sets WanPortIds field to given value.

### HasWanPortIds

`func (o *UpnpSettingOpenApiVO) HasWanPortIds() bool`

HasWanPortIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


