# VoipOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** | The status of First Priority for VoIP SIP/RTP, valid value is true or false. | 
**SipUdpPort** | Pointer to **int32** | The SIP UDP Port should be within the range of 0-65535 when parameter [enable] is true. | [optional] 

## Methods

### NewVoipOpenApiVO

`func NewVoipOpenApiVO(enable bool, ) *VoipOpenApiVO`

NewVoipOpenApiVO instantiates a new VoipOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoipOpenApiVOWithDefaults

`func NewVoipOpenApiVOWithDefaults() *VoipOpenApiVO`

NewVoipOpenApiVOWithDefaults instantiates a new VoipOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *VoipOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *VoipOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *VoipOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetSipUdpPort

`func (o *VoipOpenApiVO) GetSipUdpPort() int32`

GetSipUdpPort returns the SipUdpPort field if non-nil, zero value otherwise.

### GetSipUdpPortOk

`func (o *VoipOpenApiVO) GetSipUdpPortOk() (*int32, bool)`

GetSipUdpPortOk returns a tuple with the SipUdpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipUdpPort

`func (o *VoipOpenApiVO) SetSipUdpPort(v int32)`

SetSipUdpPort sets SipUdpPort field to given value.

### HasSipUdpPort

`func (o *VoipOpenApiVO) HasSipUdpPort() bool`

HasSipUdpPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


