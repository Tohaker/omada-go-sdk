# Dot1xPortInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dot1xEnable** | Pointer to **bool** | 802.1x enable status | [optional] 
**MabEnable** | Pointer to **bool** | MAB enable status | [optional] 
**Operation** | Pointer to **string** | switching or mirroring or aggregating | [optional] 
**Port** | Pointer to **int32** | Port number | [optional] 

## Methods

### NewDot1xPortInfoOpenApiVO

`func NewDot1xPortInfoOpenApiVO() *Dot1xPortInfoOpenApiVO`

NewDot1xPortInfoOpenApiVO instantiates a new Dot1xPortInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDot1xPortInfoOpenApiVOWithDefaults

`func NewDot1xPortInfoOpenApiVOWithDefaults() *Dot1xPortInfoOpenApiVO`

NewDot1xPortInfoOpenApiVOWithDefaults instantiates a new Dot1xPortInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDot1xEnable

`func (o *Dot1xPortInfoOpenApiVO) GetDot1xEnable() bool`

GetDot1xEnable returns the Dot1xEnable field if non-nil, zero value otherwise.

### GetDot1xEnableOk

`func (o *Dot1xPortInfoOpenApiVO) GetDot1xEnableOk() (*bool, bool)`

GetDot1xEnableOk returns a tuple with the Dot1xEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDot1xEnable

`func (o *Dot1xPortInfoOpenApiVO) SetDot1xEnable(v bool)`

SetDot1xEnable sets Dot1xEnable field to given value.

### HasDot1xEnable

`func (o *Dot1xPortInfoOpenApiVO) HasDot1xEnable() bool`

HasDot1xEnable returns a boolean if a field has been set.

### GetMabEnable

`func (o *Dot1xPortInfoOpenApiVO) GetMabEnable() bool`

GetMabEnable returns the MabEnable field if non-nil, zero value otherwise.

### GetMabEnableOk

`func (o *Dot1xPortInfoOpenApiVO) GetMabEnableOk() (*bool, bool)`

GetMabEnableOk returns a tuple with the MabEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMabEnable

`func (o *Dot1xPortInfoOpenApiVO) SetMabEnable(v bool)`

SetMabEnable sets MabEnable field to given value.

### HasMabEnable

`func (o *Dot1xPortInfoOpenApiVO) HasMabEnable() bool`

HasMabEnable returns a boolean if a field has been set.

### GetOperation

`func (o *Dot1xPortInfoOpenApiVO) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *Dot1xPortInfoOpenApiVO) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *Dot1xPortInfoOpenApiVO) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *Dot1xPortInfoOpenApiVO) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetPort

`func (o *Dot1xPortInfoOpenApiVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *Dot1xPortInfoOpenApiVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *Dot1xPortInfoOpenApiVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *Dot1xPortInfoOpenApiVO) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


