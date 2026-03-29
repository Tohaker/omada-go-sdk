# OswPortDhcpL2RelayVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CircuitId** | Pointer to **string** | Circuit ID | [optional] 
**Enable** | Pointer to **bool** | Enable | [optional] 
**Format** | Pointer to **int32** | Format should be a value as follows: 0: Normal; 1: Private | [optional] 
**RemoteId** | Pointer to **string** | Remote ID | [optional] 

## Methods

### NewOswPortDhcpL2RelayVO

`func NewOswPortDhcpL2RelayVO() *OswPortDhcpL2RelayVO`

NewOswPortDhcpL2RelayVO instantiates a new OswPortDhcpL2RelayVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswPortDhcpL2RelayVOWithDefaults

`func NewOswPortDhcpL2RelayVOWithDefaults() *OswPortDhcpL2RelayVO`

NewOswPortDhcpL2RelayVOWithDefaults instantiates a new OswPortDhcpL2RelayVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCircuitId

`func (o *OswPortDhcpL2RelayVO) GetCircuitId() string`

GetCircuitId returns the CircuitId field if non-nil, zero value otherwise.

### GetCircuitIdOk

`func (o *OswPortDhcpL2RelayVO) GetCircuitIdOk() (*string, bool)`

GetCircuitIdOk returns a tuple with the CircuitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCircuitId

`func (o *OswPortDhcpL2RelayVO) SetCircuitId(v string)`

SetCircuitId sets CircuitId field to given value.

### HasCircuitId

`func (o *OswPortDhcpL2RelayVO) HasCircuitId() bool`

HasCircuitId returns a boolean if a field has been set.

### GetEnable

`func (o *OswPortDhcpL2RelayVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *OswPortDhcpL2RelayVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *OswPortDhcpL2RelayVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *OswPortDhcpL2RelayVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetFormat

`func (o *OswPortDhcpL2RelayVO) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *OswPortDhcpL2RelayVO) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *OswPortDhcpL2RelayVO) SetFormat(v int32)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *OswPortDhcpL2RelayVO) HasFormat() bool`

HasFormat returns a boolean if a field has been set.

### GetRemoteId

`func (o *OswPortDhcpL2RelayVO) GetRemoteId() string`

GetRemoteId returns the RemoteId field if non-nil, zero value otherwise.

### GetRemoteIdOk

`func (o *OswPortDhcpL2RelayVO) GetRemoteIdOk() (*string, bool)`

GetRemoteIdOk returns a tuple with the RemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteId

`func (o *OswPortDhcpL2RelayVO) SetRemoteId(v string)`

SetRemoteId sets RemoteId field to given value.

### HasRemoteId

`func (o *OswPortDhcpL2RelayVO) HasRemoteId() bool`

HasRemoteId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


