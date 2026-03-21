# VlanOuiModeQueryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OuiProfileId** | **string** | Oui profile ID. | 
**OuiProfileName** | Pointer to **string** | Oui profile name. (Used in Get Api.) | [optional] 
**Priority** | **int32** | Selected priority should be within the range of 0 or 7. | 
**VlanId** | **int32** | Selected vlan should be within the range of 2 or 4090 | 

## Methods

### NewVlanOuiModeQueryOpenApiVO

`func NewVlanOuiModeQueryOpenApiVO(ouiProfileId string, priority int32, vlanId int32, ) *VlanOuiModeQueryOpenApiVO`

NewVlanOuiModeQueryOpenApiVO instantiates a new VlanOuiModeQueryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanOuiModeQueryOpenApiVOWithDefaults

`func NewVlanOuiModeQueryOpenApiVOWithDefaults() *VlanOuiModeQueryOpenApiVO`

NewVlanOuiModeQueryOpenApiVOWithDefaults instantiates a new VlanOuiModeQueryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOuiProfileId

`func (o *VlanOuiModeQueryOpenApiVO) GetOuiProfileId() string`

GetOuiProfileId returns the OuiProfileId field if non-nil, zero value otherwise.

### GetOuiProfileIdOk

`func (o *VlanOuiModeQueryOpenApiVO) GetOuiProfileIdOk() (*string, bool)`

GetOuiProfileIdOk returns a tuple with the OuiProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuiProfileId

`func (o *VlanOuiModeQueryOpenApiVO) SetOuiProfileId(v string)`

SetOuiProfileId sets OuiProfileId field to given value.


### GetOuiProfileName

`func (o *VlanOuiModeQueryOpenApiVO) GetOuiProfileName() string`

GetOuiProfileName returns the OuiProfileName field if non-nil, zero value otherwise.

### GetOuiProfileNameOk

`func (o *VlanOuiModeQueryOpenApiVO) GetOuiProfileNameOk() (*string, bool)`

GetOuiProfileNameOk returns a tuple with the OuiProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuiProfileName

`func (o *VlanOuiModeQueryOpenApiVO) SetOuiProfileName(v string)`

SetOuiProfileName sets OuiProfileName field to given value.

### HasOuiProfileName

`func (o *VlanOuiModeQueryOpenApiVO) HasOuiProfileName() bool`

HasOuiProfileName returns a boolean if a field has been set.

### GetPriority

`func (o *VlanOuiModeQueryOpenApiVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *VlanOuiModeQueryOpenApiVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *VlanOuiModeQueryOpenApiVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetVlanId

`func (o *VlanOuiModeQueryOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *VlanOuiModeQueryOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *VlanOuiModeQueryOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


