# DhcpL2RelayVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **bool** | Enable | [optional] 
**Format** | Pointer to **int32** | Format should be a value as follows: 0: normal, 1: private | [optional] 

## Methods

### NewDhcpL2RelayVO

`func NewDhcpL2RelayVO() *DhcpL2RelayVO`

NewDhcpL2RelayVO instantiates a new DhcpL2RelayVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpL2RelayVOWithDefaults

`func NewDhcpL2RelayVOWithDefaults() *DhcpL2RelayVO`

NewDhcpL2RelayVOWithDefaults instantiates a new DhcpL2RelayVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *DhcpL2RelayVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *DhcpL2RelayVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *DhcpL2RelayVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *DhcpL2RelayVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetFormat

`func (o *DhcpL2RelayVO) GetFormat() int32`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *DhcpL2RelayVO) GetFormatOk() (*int32, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *DhcpL2RelayVO) SetFormat(v int32)`

SetFormat sets Format field to given value.

### HasFormat

`func (o *DhcpL2RelayVO) HasFormat() bool`

HasFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


