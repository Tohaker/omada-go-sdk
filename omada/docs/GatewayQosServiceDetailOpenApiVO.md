# GatewayQosServiceDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** | The code cannot be null and should be within the range of 0-255 when protocol is 3(ICMP). | [optional] 
**DefaultProfile** | Pointer to **bool** | Indicating that the profile is the default which cannot be modified. | [optional] 
**Description** | Pointer to **string** | The description of Gateway QoS Service should contain 0 to 128 characters. | [optional] 
**DestEndPort** | Pointer to **int32** | The end port of Destination Port Range. It must be more than the start port. It should be within the range of 0-65535 when protocol is 0(TCP), 1(UDP) or 2(TCP/UDP). | [optional] 
**DestStartPort** | Pointer to **int32** | The start port of Destination Port Range. It must be less than the end port. It should be within the range of 0-65535 when protocol is 0(TCP), 1(UDP) or 2(TCP/UDP). | [optional] 
**Id** | Pointer to **string** | The ID of Gateway QoS Service. | [optional] 
**Name** | Pointer to **string** | The name of Gateway QoS Service should contain 1 to 64 characters. | [optional] 
**ProtoNum** | Pointer to **int32** | The protoNum cannot be null and should be within the range of 1-255 when protocol is 4(Other). | [optional] 
**Protocol** | Pointer to **int32** | The protocol of Gateway QoS Service should be a value as follows: 0: TCP, 1: UDP, 2: TCP/UDP, 3: ICMP, 4: Other. | [optional] 
**SourceEndPort** | Pointer to **int32** | The end port of Source Port Range. It must be more than the start port. It should be within the range of 0-65535 when protocol is 0(TCP), 1(UDP) or 2(TCP/UDP). | [optional] 
**SourceStartPort** | Pointer to **int32** | The start port of Source Port Range. It must be less than the end port. It should be within the range of 0-65535 when protocol is 0(TCP), 1(UDP) or 2(TCP/UDP). | [optional] 
**Type** | Pointer to **int32** | The type cannot be null and should be within the range of 0-255 when protocol is 3(ICMP). | [optional] 

## Methods

### NewGatewayQosServiceDetailOpenApiVO

`func NewGatewayQosServiceDetailOpenApiVO() *GatewayQosServiceDetailOpenApiVO`

NewGatewayQosServiceDetailOpenApiVO instantiates a new GatewayQosServiceDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayQosServiceDetailOpenApiVOWithDefaults

`func NewGatewayQosServiceDetailOpenApiVOWithDefaults() *GatewayQosServiceDetailOpenApiVO`

NewGatewayQosServiceDetailOpenApiVOWithDefaults instantiates a new GatewayQosServiceDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *GatewayQosServiceDetailOpenApiVO) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GatewayQosServiceDetailOpenApiVO) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GatewayQosServiceDetailOpenApiVO) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GatewayQosServiceDetailOpenApiVO) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDefaultProfile

`func (o *GatewayQosServiceDetailOpenApiVO) GetDefaultProfile() bool`

GetDefaultProfile returns the DefaultProfile field if non-nil, zero value otherwise.

### GetDefaultProfileOk

`func (o *GatewayQosServiceDetailOpenApiVO) GetDefaultProfileOk() (*bool, bool)`

GetDefaultProfileOk returns a tuple with the DefaultProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProfile

`func (o *GatewayQosServiceDetailOpenApiVO) SetDefaultProfile(v bool)`

SetDefaultProfile sets DefaultProfile field to given value.

### HasDefaultProfile

`func (o *GatewayQosServiceDetailOpenApiVO) HasDefaultProfile() bool`

HasDefaultProfile returns a boolean if a field has been set.

### GetDescription

`func (o *GatewayQosServiceDetailOpenApiVO) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayQosServiceDetailOpenApiVO) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayQosServiceDetailOpenApiVO) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GatewayQosServiceDetailOpenApiVO) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDestEndPort

`func (o *GatewayQosServiceDetailOpenApiVO) GetDestEndPort() int32`

GetDestEndPort returns the DestEndPort field if non-nil, zero value otherwise.

### GetDestEndPortOk

`func (o *GatewayQosServiceDetailOpenApiVO) GetDestEndPortOk() (*int32, bool)`

GetDestEndPortOk returns a tuple with the DestEndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestEndPort

`func (o *GatewayQosServiceDetailOpenApiVO) SetDestEndPort(v int32)`

SetDestEndPort sets DestEndPort field to given value.

### HasDestEndPort

`func (o *GatewayQosServiceDetailOpenApiVO) HasDestEndPort() bool`

HasDestEndPort returns a boolean if a field has been set.

### GetDestStartPort

`func (o *GatewayQosServiceDetailOpenApiVO) GetDestStartPort() int32`

GetDestStartPort returns the DestStartPort field if non-nil, zero value otherwise.

### GetDestStartPortOk

`func (o *GatewayQosServiceDetailOpenApiVO) GetDestStartPortOk() (*int32, bool)`

GetDestStartPortOk returns a tuple with the DestStartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestStartPort

`func (o *GatewayQosServiceDetailOpenApiVO) SetDestStartPort(v int32)`

SetDestStartPort sets DestStartPort field to given value.

### HasDestStartPort

`func (o *GatewayQosServiceDetailOpenApiVO) HasDestStartPort() bool`

HasDestStartPort returns a boolean if a field has been set.

### GetId

`func (o *GatewayQosServiceDetailOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GatewayQosServiceDetailOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GatewayQosServiceDetailOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GatewayQosServiceDetailOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GatewayQosServiceDetailOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayQosServiceDetailOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayQosServiceDetailOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GatewayQosServiceDetailOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProtoNum

`func (o *GatewayQosServiceDetailOpenApiVO) GetProtoNum() int32`

GetProtoNum returns the ProtoNum field if non-nil, zero value otherwise.

### GetProtoNumOk

`func (o *GatewayQosServiceDetailOpenApiVO) GetProtoNumOk() (*int32, bool)`

GetProtoNumOk returns a tuple with the ProtoNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtoNum

`func (o *GatewayQosServiceDetailOpenApiVO) SetProtoNum(v int32)`

SetProtoNum sets ProtoNum field to given value.

### HasProtoNum

`func (o *GatewayQosServiceDetailOpenApiVO) HasProtoNum() bool`

HasProtoNum returns a boolean if a field has been set.

### GetProtocol

`func (o *GatewayQosServiceDetailOpenApiVO) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *GatewayQosServiceDetailOpenApiVO) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *GatewayQosServiceDetailOpenApiVO) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *GatewayQosServiceDetailOpenApiVO) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSourceEndPort

`func (o *GatewayQosServiceDetailOpenApiVO) GetSourceEndPort() int32`

GetSourceEndPort returns the SourceEndPort field if non-nil, zero value otherwise.

### GetSourceEndPortOk

`func (o *GatewayQosServiceDetailOpenApiVO) GetSourceEndPortOk() (*int32, bool)`

GetSourceEndPortOk returns a tuple with the SourceEndPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceEndPort

`func (o *GatewayQosServiceDetailOpenApiVO) SetSourceEndPort(v int32)`

SetSourceEndPort sets SourceEndPort field to given value.

### HasSourceEndPort

`func (o *GatewayQosServiceDetailOpenApiVO) HasSourceEndPort() bool`

HasSourceEndPort returns a boolean if a field has been set.

### GetSourceStartPort

`func (o *GatewayQosServiceDetailOpenApiVO) GetSourceStartPort() int32`

GetSourceStartPort returns the SourceStartPort field if non-nil, zero value otherwise.

### GetSourceStartPortOk

`func (o *GatewayQosServiceDetailOpenApiVO) GetSourceStartPortOk() (*int32, bool)`

GetSourceStartPortOk returns a tuple with the SourceStartPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceStartPort

`func (o *GatewayQosServiceDetailOpenApiVO) SetSourceStartPort(v int32)`

SetSourceStartPort sets SourceStartPort field to given value.

### HasSourceStartPort

`func (o *GatewayQosServiceDetailOpenApiVO) HasSourceStartPort() bool`

HasSourceStartPort returns a boolean if a field has been set.

### GetType

`func (o *GatewayQosServiceDetailOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GatewayQosServiceDetailOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GatewayQosServiceDetailOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *GatewayQosServiceDetailOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


