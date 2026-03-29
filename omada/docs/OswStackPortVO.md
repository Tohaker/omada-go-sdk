# OswStackPortVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Device Mac | [optional] 
**Ports** | Pointer to [**[]OswStackMemberPortVO**](OswStackMemberPortVO.md) | Port list | [optional] 
**Unit** | Pointer to **int32** | Unit | [optional] 

## Methods

### NewOswStackPortVO

`func NewOswStackPortVO() *OswStackPortVO`

NewOswStackPortVO instantiates a new OswStackPortVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackPortVOWithDefaults

`func NewOswStackPortVOWithDefaults() *OswStackPortVO`

NewOswStackPortVOWithDefaults instantiates a new OswStackPortVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *OswStackPortVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswStackPortVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswStackPortVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswStackPortVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetPorts

`func (o *OswStackPortVO) GetPorts() []OswStackMemberPortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *OswStackPortVO) GetPortsOk() (*[]OswStackMemberPortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *OswStackPortVO) SetPorts(v []OswStackMemberPortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *OswStackPortVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetUnit

`func (o *OswStackPortVO) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *OswStackPortVO) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *OswStackPortVO) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *OswStackPortVO) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


