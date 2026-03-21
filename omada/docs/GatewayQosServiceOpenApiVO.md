# GatewayQosServiceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | The code cannot be null and should be within the range of 0-255 when protocol is 3(ICMP). | [optional] 
**Description** | Pointer to **string** | The description of Gateway QoS Service should contain 0 to 128 characters. | [optional] 
**DestEndPort** | Pointer to **int32** | The end port of Destination Port Range. It must be more than the start port. It should be within the range of 0-65535 when protocol is 0(TCP), 1(UDP) or 2(TCP/UDP). | [optional] 
**DestStartPort** | Pointer to **int32** | The start port of Destination Port Range. It must be less than the end port. It should be within the range of 0-65535 when protocol is 0(TCP), 1(UDP) or 2(TCP/UDP). | [optional] 
**Name** | **string** | The name of Gateway QoS Service should contain 1 to 64 characters. | 
**ProtoNum** | Pointer to **int32** | The protoNum cannot be null and should be within the range of 1-255 when protocol is 4(Other). | [optional] 
**Protocol** | **int32** | The protocol of Gateway QoS Service should be a value as follows: 0: TCP, 1: UDP, 2: TCP/UDP, 3: ICMP, 4: Other. | 
**SourceEndPort** | Pointer to **int32** | The end port of Source Port Range. It must be more than the start port. It should be within the range of 0-65535 when protocol is 0(TCP), 1(UDP) or 2(TCP/UDP). | [optional] 
**SourceStartPort** | Pointer to **int32** | The start port of Source Port Range. It must be less than the end port. It should be within the range of 0-65535 when protocol is 0(TCP), 1(UDP) or 2(TCP/UDP). | [optional] 
**Type** | Pointer to **int32** | The type cannot be null and should be within the range of 0-255 when protocol is 3(ICMP). | [optional] 

## Methods

### NewGatewayQosServiceOpenApiVO

`func NewGatewayQosServiceOpenApiVO(name string, protocol int32, ) *GatewayQosServiceOpenApiVO`

NewGatewayQosServiceOpenApiVO instantiates a new GatewayQosServiceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayQosServiceOpenApiVOWithDefaults

`func NewGatewayQosServiceOpenApiVOWithDefaults() *GatewayQosServiceOpenApiVO`

NewGatewayQosServiceOpenApiVOWithDefaults instantiates a new GatewayQosServiceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GatewayQosServiceOpenApiVO) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GatewayQosServiceOpenApiVO) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GatewayQosServiceOpenApiVO) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GatewayQosServiceOpenApiVO) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *GatewayQosServiceOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayQosServiceOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayQosServiceOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GatewayQosServiceOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestEndPort

`func (o *GatewayQosServiceOpenApiVO) GetDestEndPort() int32`

GetDestEndPort returns the DestEndPort field if non-nil, zero value otherwise.

### GetDestEndPortOk

`func (o *GatewayQosServiceOpenApiVO) GetDestEndPortOk() (*int32, bool)`

GetDestEndPortOk returns a tuple with the DestEndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestEndPort

`func (o *GatewayQosServiceOpenApiVO) SetDestEndPort(v int32)`

SetDestEndPort sets DestEndPort field to given value.

### HasDestEndPort

`func (o *GatewayQosServiceOpenApiVO) HasDestEndPort() bool`

HasDestEndPort returns a boolean if a field has been set.

### GetDestStartPort

`func (o *GatewayQosServiceOpenApiVO) GetDestStartPort() int32`

GetDestStartPort returns the DestStartPort field if non-nil, zero value otherwise.

### GetDestStartPortOk

`func (o *GatewayQosServiceOpenApiVO) GetDestStartPortOk() (*int32, bool)`

GetDestStartPortOk returns a tuple with the DestStartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestStartPort

`func (o *GatewayQosServiceOpenApiVO) SetDestStartPort(v int32)`

SetDestStartPort sets DestStartPort field to given value.

### HasDestStartPort

`func (o *GatewayQosServiceOpenApiVO) HasDestStartPort() bool`

HasDestStartPort returns a boolean if a field has been set.

### GetName

`func (o *GatewayQosServiceOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayQosServiceOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayQosServiceOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetProtoNum

`func (o *GatewayQosServiceOpenApiVO) GetProtoNum() int32`

GetProtoNum returns the ProtoNum field if non-nil, zero value otherwise.

### GetProtoNumOk

`func (o *GatewayQosServiceOpenApiVO) GetProtoNumOk() (*int32, bool)`

GetProtoNumOk returns a tuple with the ProtoNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoNum

`func (o *GatewayQosServiceOpenApiVO) SetProtoNum(v int32)`

SetProtoNum sets ProtoNum field to given value.

### HasProtoNum

`func (o *GatewayQosServiceOpenApiVO) HasProtoNum() bool`

HasProtoNum returns a boolean if a field has been set.

### GetProtocol

`func (o *GatewayQosServiceOpenApiVO) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GatewayQosServiceOpenApiVO) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GatewayQosServiceOpenApiVO) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.


### GetSourceEndPort

`func (o *GatewayQosServiceOpenApiVO) GetSourceEndPort() int32`

GetSourceEndPort returns the SourceEndPort field if non-nil, zero value otherwise.

### GetSourceEndPortOk

`func (o *GatewayQosServiceOpenApiVO) GetSourceEndPortOk() (*int32, bool)`

GetSourceEndPortOk returns a tuple with the SourceEndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEndPort

`func (o *GatewayQosServiceOpenApiVO) SetSourceEndPort(v int32)`

SetSourceEndPort sets SourceEndPort field to given value.

### HasSourceEndPort

`func (o *GatewayQosServiceOpenApiVO) HasSourceEndPort() bool`

HasSourceEndPort returns a boolean if a field has been set.

### GetSourceStartPort

`func (o *GatewayQosServiceOpenApiVO) GetSourceStartPort() int32`

GetSourceStartPort returns the SourceStartPort field if non-nil, zero value otherwise.

### GetSourceStartPortOk

`func (o *GatewayQosServiceOpenApiVO) GetSourceStartPortOk() (*int32, bool)`

GetSourceStartPortOk returns a tuple with the SourceStartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStartPort

`func (o *GatewayQosServiceOpenApiVO) SetSourceStartPort(v int32)`

SetSourceStartPort sets SourceStartPort field to given value.

### HasSourceStartPort

`func (o *GatewayQosServiceOpenApiVO) HasSourceStartPort() bool`

HasSourceStartPort returns a boolean if a field has been set.

### GetType

`func (o *GatewayQosServiceOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayQosServiceOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayQosServiceOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayQosServiceOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


