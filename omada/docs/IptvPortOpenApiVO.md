# IptvPortOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | This field represents Port&#39;s Name. | [optional] 
**PortId** | **string** | This field represents Port ID. | 
**Type** | **int32** | Type should be a value as follows: 1: Internet; 2:IPTV; 3:IP-Phone | 

## Methods

### NewIptvPortOpenApiVO

`func NewIptvPortOpenApiVO(portId string, type_ int32, ) *IptvPortOpenApiVO`

NewIptvPortOpenApiVO instantiates a new IptvPortOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIptvPortOpenApiVOWithDefaults

`func NewIptvPortOpenApiVOWithDefaults() *IptvPortOpenApiVO`

NewIptvPortOpenApiVOWithDefaults instantiates a new IptvPortOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *IptvPortOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IptvPortOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IptvPortOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IptvPortOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortId

`func (o *IptvPortOpenApiVO) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *IptvPortOpenApiVO) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *IptvPortOpenApiVO) SetPortId(v string)`

SetPortId sets PortId field to given value.


### GetType

`func (o *IptvPortOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IptvPortOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IptvPortOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


