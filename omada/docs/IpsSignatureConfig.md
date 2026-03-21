# IpsSignatureConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to **int32** | Direction should be a value as follow: 0: both direction; 1: source direction; 2: destination direction | [optional] 
**Ip** | Pointer to **string** | IPS signature traffic Source. If parameter [trafficType] is 0, parameter [ip] is needed | [optional] 
**Subnet** | Pointer to **string** | IPS signature traffic Source. If parameter [trafficType] is 1, parameter [subnet] is needed | [optional] 
**TrafficType** | Pointer to **int32** | TrafficType should be a value as follow: 0: ip address; 1: subnet | [optional] 
**Type** | **int32** | Type should be a value as follow: 0: all traffic; 1: packet tracking | 

## Methods

### NewIpsSignatureConfig

`func NewIpsSignatureConfig(type_ int32, ) *IpsSignatureConfig`

NewIpsSignatureConfig instantiates a new IpsSignatureConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsSignatureConfigWithDefaults

`func NewIpsSignatureConfigWithDefaults() *IpsSignatureConfig`

NewIpsSignatureConfigWithDefaults instantiates a new IpsSignatureConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *IpsSignatureConfig) GetDirection() int32`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *IpsSignatureConfig) GetDirectionOk() (*int32, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *IpsSignatureConfig) SetDirection(v int32)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *IpsSignatureConfig) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetIp

`func (o *IpsSignatureConfig) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *IpsSignatureConfig) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *IpsSignatureConfig) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *IpsSignatureConfig) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetSubnet

`func (o *IpsSignatureConfig) GetSubnet() string`

GetSubnet returns the Subnet field if non-nil, zero value otherwise.

### GetSubnetOk

`func (o *IpsSignatureConfig) GetSubnetOk() (*string, bool)`

GetSubnetOk returns a tuple with the Subnet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubnet

`func (o *IpsSignatureConfig) SetSubnet(v string)`

SetSubnet sets Subnet field to given value.

### HasSubnet

`func (o *IpsSignatureConfig) HasSubnet() bool`

HasSubnet returns a boolean if a field has been set.

### GetTrafficType

`func (o *IpsSignatureConfig) GetTrafficType() int32`

GetTrafficType returns the TrafficType field if non-nil, zero value otherwise.

### GetTrafficTypeOk

`func (o *IpsSignatureConfig) GetTrafficTypeOk() (*int32, bool)`

GetTrafficTypeOk returns a tuple with the TrafficType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficType

`func (o *IpsSignatureConfig) SetTrafficType(v int32)`

SetTrafficType sets TrafficType field to given value.

### HasTrafficType

`func (o *IpsSignatureConfig) HasTrafficType() bool`

HasTrafficType returns a boolean if a field has been set.

### GetType

`func (o *IpsSignatureConfig) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IpsSignatureConfig) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IpsSignatureConfig) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


