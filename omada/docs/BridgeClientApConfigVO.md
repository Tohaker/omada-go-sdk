# BridgeClientApConfigVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | **string** | Bridge TDMA Client AP mac. | 
**Priority** | Pointer to **int32** | Bridge TDMA Client AP priority config. 0: high, 1:base, 2:low. | [optional] 

## Methods

### NewBridgeClientApConfigVO

`func NewBridgeClientApConfigVO(mac string, ) *BridgeClientApConfigVO`

NewBridgeClientApConfigVO instantiates a new BridgeClientApConfigVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBridgeClientApConfigVOWithDefaults

`func NewBridgeClientApConfigVOWithDefaults() *BridgeClientApConfigVO`

NewBridgeClientApConfigVOWithDefaults instantiates a new BridgeClientApConfigVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *BridgeClientApConfigVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *BridgeClientApConfigVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *BridgeClientApConfigVO) SetMac(v string)`

SetMac sets Mac field to given value.


### GetPriority

`func (o *BridgeClientApConfigVO) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *BridgeClientApConfigVO) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *BridgeClientApConfigVO) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *BridgeClientApConfigVO) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


