# ApUplinkConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auto** | **bool** | uplink auto | 
**Port** | Pointer to **string** | port id | [optional] 

## Methods

### NewApUplinkConfigOpenApiVO

`func NewApUplinkConfigOpenApiVO(auto bool, ) *ApUplinkConfigOpenApiVO`

NewApUplinkConfigOpenApiVO instantiates a new ApUplinkConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApUplinkConfigOpenApiVOWithDefaults

`func NewApUplinkConfigOpenApiVOWithDefaults() *ApUplinkConfigOpenApiVO`

NewApUplinkConfigOpenApiVOWithDefaults instantiates a new ApUplinkConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuto

`func (o *ApUplinkConfigOpenApiVO) GetAuto() bool`

GetAuto returns the Auto field if non-nil, zero value otherwise.

### GetAutoOk

`func (o *ApUplinkConfigOpenApiVO) GetAutoOk() (*bool, bool)`

GetAutoOk returns a tuple with the Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuto

`func (o *ApUplinkConfigOpenApiVO) SetAuto(v bool)`

SetAuto sets Auto field to given value.


### GetPort

`func (o *ApUplinkConfigOpenApiVO) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ApUplinkConfigOpenApiVO) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ApUplinkConfigOpenApiVO) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *ApUplinkConfigOpenApiVO) HasPort() bool`

HasPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


