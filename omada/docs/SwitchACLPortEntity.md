# SwitchACLPortEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomLagIds** | **[]int32** | Custom lag IDs | 
**CustomPortIds** | **[]int32** | Custom port IDs | 
**Mac** | **string** | MAC | 
**StackId** | Pointer to **string** | Stack ID | [optional] 
**StandardCustomPortIds** | Pointer to **[]string** | Custom standard port IDs | [optional] 
**Vrf** | **string** | VRF should be 1 to 15 characters consisting of numbers (0 to 9), uppercase and lowercase letters (A to Z, a to z), and symbols -_@.+ but cannot be . or .. only. | 
**VrfId** | Pointer to **string** | VRF ID | [optional] 

## Methods

### NewSwitchACLPortEntity

`func NewSwitchACLPortEntity(customLagIds []int32, customPortIds []int32, mac string, vrf string, ) *SwitchACLPortEntity`

NewSwitchACLPortEntity instantiates a new SwitchACLPortEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchACLPortEntityWithDefaults

`func NewSwitchACLPortEntityWithDefaults() *SwitchACLPortEntity`

NewSwitchACLPortEntityWithDefaults instantiates a new SwitchACLPortEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomLagIds

`func (o *SwitchACLPortEntity) GetCustomLagIds() []int32`

GetCustomLagIds returns the CustomLagIds field if non-nil, zero value otherwise.

### GetCustomLagIdsOk

`func (o *SwitchACLPortEntity) GetCustomLagIdsOk() (*[]int32, bool)`

GetCustomLagIdsOk returns a tuple with the CustomLagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLagIds

`func (o *SwitchACLPortEntity) SetCustomLagIds(v []int32)`

SetCustomLagIds sets CustomLagIds field to given value.


### GetCustomPortIds

`func (o *SwitchACLPortEntity) GetCustomPortIds() []int32`

GetCustomPortIds returns the CustomPortIds field if non-nil, zero value otherwise.

### GetCustomPortIdsOk

`func (o *SwitchACLPortEntity) GetCustomPortIdsOk() (*[]int32, bool)`

GetCustomPortIdsOk returns a tuple with the CustomPortIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPortIds

`func (o *SwitchACLPortEntity) SetCustomPortIds(v []int32)`

SetCustomPortIds sets CustomPortIds field to given value.


### GetMac

`func (o *SwitchACLPortEntity) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SwitchACLPortEntity) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SwitchACLPortEntity) SetMac(v string)`

SetMac sets Mac field to given value.


### GetStackId

`func (o *SwitchACLPortEntity) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *SwitchACLPortEntity) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *SwitchACLPortEntity) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *SwitchACLPortEntity) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStandardCustomPortIds

`func (o *SwitchACLPortEntity) GetStandardCustomPortIds() []string`

GetStandardCustomPortIds returns the StandardCustomPortIds field if non-nil, zero value otherwise.

### GetStandardCustomPortIdsOk

`func (o *SwitchACLPortEntity) GetStandardCustomPortIdsOk() (*[]string, bool)`

GetStandardCustomPortIdsOk returns a tuple with the StandardCustomPortIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandardCustomPortIds

`func (o *SwitchACLPortEntity) SetStandardCustomPortIds(v []string)`

SetStandardCustomPortIds sets StandardCustomPortIds field to given value.

### HasStandardCustomPortIds

`func (o *SwitchACLPortEntity) HasStandardCustomPortIds() bool`

HasStandardCustomPortIds returns a boolean if a field has been set.

### GetVrf

`func (o *SwitchACLPortEntity) GetVrf() string`

GetVrf returns the Vrf field if non-nil, zero value otherwise.

### GetVrfOk

`func (o *SwitchACLPortEntity) GetVrfOk() (*string, bool)`

GetVrfOk returns a tuple with the Vrf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrf

`func (o *SwitchACLPortEntity) SetVrf(v string)`

SetVrf sets Vrf field to given value.


### GetVrfId

`func (o *SwitchACLPortEntity) GetVrfId() string`

GetVrfId returns the VrfId field if non-nil, zero value otherwise.

### GetVrfIdOk

`func (o *SwitchACLPortEntity) GetVrfIdOk() (*string, bool)`

GetVrfIdOk returns a tuple with the VrfId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVrfId

`func (o *SwitchACLPortEntity) SetVrfId(v string)`

SetVrfId sets VrfId field to given value.

### HasVrfId

`func (o *SwitchACLPortEntity) HasVrfId() bool`

HasVrfId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


