# VlanOuiModeOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OuiProfileId** | **string** | Oui profile ID. | 
**Priority** | **int32** | Selected priority, valid range is 0 to 7. | 
**VlanId** | **int32** | Selected vlan, valid range is 1 to 4090. | 

## Methods

### NewVlanOuiModeOpenApiVO

`func NewVlanOuiModeOpenApiVO(ouiProfileId string, priority int32, vlanId int32, ) *VlanOuiModeOpenApiVO`

NewVlanOuiModeOpenApiVO instantiates a new VlanOuiModeOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanOuiModeOpenApiVOWithDefaults

`func NewVlanOuiModeOpenApiVOWithDefaults() *VlanOuiModeOpenApiVO`

NewVlanOuiModeOpenApiVOWithDefaults instantiates a new VlanOuiModeOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOuiProfileId

`func (o *VlanOuiModeOpenApiVO) GetOuiProfileId() string`

GetOuiProfileId returns the OuiProfileId field if non-nil, zero value otherwise.

### GetOuiProfileIdOk

`func (o *VlanOuiModeOpenApiVO) GetOuiProfileIdOk() (*string, bool)`

GetOuiProfileIdOk returns a tuple with the OuiProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOuiProfileId

`func (o *VlanOuiModeOpenApiVO) SetOuiProfileId(v string)`

SetOuiProfileId sets OuiProfileId field to given value.


### GetPriority

`func (o *VlanOuiModeOpenApiVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *VlanOuiModeOpenApiVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *VlanOuiModeOpenApiVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.


### GetVlanId

`func (o *VlanOuiModeOpenApiVO) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *VlanOuiModeOpenApiVO) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *VlanOuiModeOpenApiVO) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


