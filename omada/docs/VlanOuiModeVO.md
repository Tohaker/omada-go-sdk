# VlanOuiModeVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OuiProfileId** | **string** |  | 
**OuiProfileName** | Pointer to **string** |  | [optional] 
**Priority** | **int32** |  | 
**VlanId** | **int32** |  | 

## Methods

### NewVlanOuiModeVO

`func NewVlanOuiModeVO(ouiProfileId string, priority int32, vlanId int32, ) *VlanOuiModeVO`

NewVlanOuiModeVO instantiates a new VlanOuiModeVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanOuiModeVOWithDefaults

`func NewVlanOuiModeVOWithDefaults() *VlanOuiModeVO`

NewVlanOuiModeVOWithDefaults instantiates a new VlanOuiModeVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOuiProfileId

`func (o *VlanOuiModeVO) GetOuiProfileId() string`

GetOuiProfileId returns the OuiProfileId field if non-nil, zero value otherwise.

### GetOuiProfileIdOk

`func (o *VlanOuiModeVO) GetOuiProfileIdOk() (*string, bool)`

GetOuiProfileIdOk returns a tuple with the OuiProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuiProfileId

`func (o *VlanOuiModeVO) SetOuiProfileId(v string)`

SetOuiProfileId sets OuiProfileId field to given value.


### GetOuiProfileName

`func (o *VlanOuiModeVO) GetOuiProfileName() string`

GetOuiProfileName returns the OuiProfileName field if non-nil, zero value otherwise.

### GetOuiProfileNameOk

`func (o *VlanOuiModeVO) GetOuiProfileNameOk() (*string, bool)`

GetOuiProfileNameOk returns a tuple with the OuiProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuiProfileName

`func (o *VlanOuiModeVO) SetOuiProfileName(v string)`

SetOuiProfileName sets OuiProfileName field to given value.

### HasOuiProfileName

`func (o *VlanOuiModeVO) HasOuiProfileName() bool`

HasOuiProfileName returns a boolean if a field has been set.

### GetPriority

`func (o *VlanOuiModeVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *VlanOuiModeVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *VlanOuiModeVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetVlanId

`func (o *VlanOuiModeVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *VlanOuiModeVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *VlanOuiModeVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


