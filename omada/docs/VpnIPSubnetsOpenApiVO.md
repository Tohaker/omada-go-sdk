# VpnIPSubnetsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ip** | **string** | IP address, should be a valid IP format. | 
**Mask** | **int32** | IP mask, mask should be within the range of 0-32. | 

## Methods

### NewVpnIPSubnetsOpenApiVO

`func NewVpnIPSubnetsOpenApiVO(ip string, mask int32, ) *VpnIPSubnetsOpenApiVO`

NewVpnIPSubnetsOpenApiVO instantiates a new VpnIPSubnetsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnIPSubnetsOpenApiVOWithDefaults

`func NewVpnIPSubnetsOpenApiVOWithDefaults() *VpnIPSubnetsOpenApiVO`

NewVpnIPSubnetsOpenApiVOWithDefaults instantiates a new VpnIPSubnetsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIp

`func (o *VpnIPSubnetsOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *VpnIPSubnetsOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *VpnIPSubnetsOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.


### GetMask

`func (o *VpnIPSubnetsOpenApiVO) GetMask() int32`

GetMask returns the Mask field if non-nil, zero value otherwise.

### GetMaskOk

`func (o *VpnIPSubnetsOpenApiVO) GetMaskOk() (*int32, bool)`

GetMaskOk returns a tuple with the Mask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMask

`func (o *VpnIPSubnetsOpenApiVO) SetMask(v int32)`

SetMask sets Mask field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


