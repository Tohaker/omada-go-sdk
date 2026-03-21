# IptvPortConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortId** | **string** | Port ID can be obtained from &#39;Get internet basic info&#39; interface. | 
**Type** | **int32** | Type should be a value as follows: 1: Internet; 2:IPTV; 3:IP-Phone, 3 is valid only IPTV mode is custom. | 

## Methods

### NewIptvPortConfigOpenApiVO

`func NewIptvPortConfigOpenApiVO(portId string, type_ int32, ) *IptvPortConfigOpenApiVO`

NewIptvPortConfigOpenApiVO instantiates a new IptvPortConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIptvPortConfigOpenApiVOWithDefaults

`func NewIptvPortConfigOpenApiVOWithDefaults() *IptvPortConfigOpenApiVO`

NewIptvPortConfigOpenApiVOWithDefaults instantiates a new IptvPortConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortId

`func (o *IptvPortConfigOpenApiVO) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *IptvPortConfigOpenApiVO) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *IptvPortConfigOpenApiVO) SetPortId(v string)`

SetPortId sets PortId field to given value.


### GetType

`func (o *IptvPortConfigOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IptvPortConfigOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IptvPortConfigOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


