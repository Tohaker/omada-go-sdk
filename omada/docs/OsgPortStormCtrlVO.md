# OsgPortStormCtrlVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | Pointer to **int32** |  | [optional] 
**Broadcast** | Pointer to **int32** |  | [optional] 
**BroadcastEnable** | **bool** |  | 
**Multicast** | Pointer to **int32** |  | [optional] 
**MulticastEnable** | **bool** |  | 
**RecoverTime** | Pointer to **int32** |  | [optional] 
**UnknownUnicast** | Pointer to **int32** |  | [optional] 
**UnknownUnicastEnable** | **bool** |  | 

## Methods

### NewOsgPortStormCtrlVO

`func NewOsgPortStormCtrlVO(broadcastEnable bool, multicastEnable bool, unknownUnicastEnable bool, ) *OsgPortStormCtrlVO`

NewOsgPortStormCtrlVO instantiates a new OsgPortStormCtrlVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgPortStormCtrlVOWithDefaults

`func NewOsgPortStormCtrlVOWithDefaults() *OsgPortStormCtrlVO`

NewOsgPortStormCtrlVOWithDefaults instantiates a new OsgPortStormCtrlVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *OsgPortStormCtrlVO) GetAction() int32`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *OsgPortStormCtrlVO) GetActionOk() (*int32, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *OsgPortStormCtrlVO) SetAction(v int32)`

SetAction sets Action field to given value.

### HasAction

`func (o *OsgPortStormCtrlVO) HasAction() bool`

HasAction returns a boolean if a field has been set.

### GetBroadcast

`func (o *OsgPortStormCtrlVO) GetBroadcast() int32`

GetBroadcast returns the Broadcast field if non-nil, zero value otherwise.

### GetBroadcastOk

`func (o *OsgPortStormCtrlVO) GetBroadcastOk() (*int32, bool)`

GetBroadcastOk returns a tuple with the Broadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcast

`func (o *OsgPortStormCtrlVO) SetBroadcast(v int32)`

SetBroadcast sets Broadcast field to given value.

### HasBroadcast

`func (o *OsgPortStormCtrlVO) HasBroadcast() bool`

HasBroadcast returns a boolean if a field has been set.

### GetBroadcastEnable

`func (o *OsgPortStormCtrlVO) GetBroadcastEnable() bool`

GetBroadcastEnable returns the BroadcastEnable field if non-nil, zero value otherwise.

### GetBroadcastEnableOk

`func (o *OsgPortStormCtrlVO) GetBroadcastEnableOk() (*bool, bool)`

GetBroadcastEnableOk returns a tuple with the BroadcastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroadcastEnable

`func (o *OsgPortStormCtrlVO) SetBroadcastEnable(v bool)`

SetBroadcastEnable sets BroadcastEnable field to given value.


### GetMulticast

`func (o *OsgPortStormCtrlVO) GetMulticast() int32`

GetMulticast returns the Multicast field if non-nil, zero value otherwise.

### GetMulticastOk

`func (o *OsgPortStormCtrlVO) GetMulticastOk() (*int32, bool)`

GetMulticastOk returns a tuple with the Multicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticast

`func (o *OsgPortStormCtrlVO) SetMulticast(v int32)`

SetMulticast sets Multicast field to given value.

### HasMulticast

`func (o *OsgPortStormCtrlVO) HasMulticast() bool`

HasMulticast returns a boolean if a field has been set.

### GetMulticastEnable

`func (o *OsgPortStormCtrlVO) GetMulticastEnable() bool`

GetMulticastEnable returns the MulticastEnable field if non-nil, zero value otherwise.

### GetMulticastEnableOk

`func (o *OsgPortStormCtrlVO) GetMulticastEnableOk() (*bool, bool)`

GetMulticastEnableOk returns a tuple with the MulticastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMulticastEnable

`func (o *OsgPortStormCtrlVO) SetMulticastEnable(v bool)`

SetMulticastEnable sets MulticastEnable field to given value.


### GetRecoverTime

`func (o *OsgPortStormCtrlVO) GetRecoverTime() int32`

GetRecoverTime returns the RecoverTime field if non-nil, zero value otherwise.

### GetRecoverTimeOk

`func (o *OsgPortStormCtrlVO) GetRecoverTimeOk() (*int32, bool)`

GetRecoverTimeOk returns a tuple with the RecoverTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecoverTime

`func (o *OsgPortStormCtrlVO) SetRecoverTime(v int32)`

SetRecoverTime sets RecoverTime field to given value.

### HasRecoverTime

`func (o *OsgPortStormCtrlVO) HasRecoverTime() bool`

HasRecoverTime returns a boolean if a field has been set.

### GetUnknownUnicast

`func (o *OsgPortStormCtrlVO) GetUnknownUnicast() int32`

GetUnknownUnicast returns the UnknownUnicast field if non-nil, zero value otherwise.

### GetUnknownUnicastOk

`func (o *OsgPortStormCtrlVO) GetUnknownUnicastOk() (*int32, bool)`

GetUnknownUnicastOk returns a tuple with the UnknownUnicast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownUnicast

`func (o *OsgPortStormCtrlVO) SetUnknownUnicast(v int32)`

SetUnknownUnicast sets UnknownUnicast field to given value.

### HasUnknownUnicast

`func (o *OsgPortStormCtrlVO) HasUnknownUnicast() bool`

HasUnknownUnicast returns a boolean if a field has been set.

### GetUnknownUnicastEnable

`func (o *OsgPortStormCtrlVO) GetUnknownUnicastEnable() bool`

GetUnknownUnicastEnable returns the UnknownUnicastEnable field if non-nil, zero value otherwise.

### GetUnknownUnicastEnableOk

`func (o *OsgPortStormCtrlVO) GetUnknownUnicastEnableOk() (*bool, bool)`

GetUnknownUnicastEnableOk returns a tuple with the UnknownUnicastEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnknownUnicastEnable

`func (o *OsgPortStormCtrlVO) SetUnknownUnicastEnable(v bool)`

SetUnknownUnicastEnable sets UnknownUnicastEnable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


