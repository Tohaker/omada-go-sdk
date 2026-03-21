# OsgWanStatusVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | Pointer to **int32** |  | [optional] 
**LanId** | Pointer to **string** |  | [optional] 
**PdEnable** | Pointer to **int32** |  | [optional] 
**PdSize** | Pointer to **int32** |  | [optional] 
**PortName** | Pointer to **string** |  | [optional] 
**PortUuid** | Pointer to **string** |  | [optional] 
**Prefix** | Pointer to **string** |  | [optional] 
**Proto** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** |  | [optional] 

## Methods

### NewOsgWanStatusVO

`func NewOsgWanStatusVO() *OsgWanStatusVO`

NewOsgWanStatusVO instantiates a new OsgWanStatusVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOsgWanStatusVOWithDefaults

`func NewOsgWanStatusVOWithDefaults() *OsgWanStatusVO`

NewOsgWanStatusVOWithDefaults instantiates a new OsgWanStatusVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *OsgWanStatusVO) GetEnable() int32`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *OsgWanStatusVO) GetEnableOk() (*int32, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *OsgWanStatusVO) SetEnable(v int32)`

SetEnable sets Enable field to given value.

### HasEnable

`func (o *OsgWanStatusVO) HasEnable() bool`

HasEnable returns a boolean if a field has been set.

### GetLanId

`func (o *OsgWanStatusVO) GetLanId() string`

GetLanId returns the LanId field if non-nil, zero value otherwise.

### GetLanIdOk

`func (o *OsgWanStatusVO) GetLanIdOk() (*string, bool)`

GetLanIdOk returns a tuple with the LanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanId

`func (o *OsgWanStatusVO) SetLanId(v string)`

SetLanId sets LanId field to given value.

### HasLanId

`func (o *OsgWanStatusVO) HasLanId() bool`

HasLanId returns a boolean if a field has been set.

### GetPdEnable

`func (o *OsgWanStatusVO) GetPdEnable() int32`

GetPdEnable returns the PdEnable field if non-nil, zero value otherwise.

### GetPdEnableOk

`func (o *OsgWanStatusVO) GetPdEnableOk() (*int32, bool)`

GetPdEnableOk returns a tuple with the PdEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdEnable

`func (o *OsgWanStatusVO) SetPdEnable(v int32)`

SetPdEnable sets PdEnable field to given value.

### HasPdEnable

`func (o *OsgWanStatusVO) HasPdEnable() bool`

HasPdEnable returns a boolean if a field has been set.

### GetPdSize

`func (o *OsgWanStatusVO) GetPdSize() int32`

GetPdSize returns the PdSize field if non-nil, zero value otherwise.

### GetPdSizeOk

`func (o *OsgWanStatusVO) GetPdSizeOk() (*int32, bool)`

GetPdSizeOk returns a tuple with the PdSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdSize

`func (o *OsgWanStatusVO) SetPdSize(v int32)`

SetPdSize sets PdSize field to given value.

### HasPdSize

`func (o *OsgWanStatusVO) HasPdSize() bool`

HasPdSize returns a boolean if a field has been set.

### GetPortName

`func (o *OsgWanStatusVO) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *OsgWanStatusVO) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *OsgWanStatusVO) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *OsgWanStatusVO) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPortUuid

`func (o *OsgWanStatusVO) GetPortUuid() string`

GetPortUuid returns the PortUuid field if non-nil, zero value otherwise.

### GetPortUuidOk

`func (o *OsgWanStatusVO) GetPortUuidOk() (*string, bool)`

GetPortUuidOk returns a tuple with the PortUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortUuid

`func (o *OsgWanStatusVO) SetPortUuid(v string)`

SetPortUuid sets PortUuid field to given value.

### HasPortUuid

`func (o *OsgWanStatusVO) HasPortUuid() bool`

HasPortUuid returns a boolean if a field has been set.

### GetPrefix

`func (o *OsgWanStatusVO) GetPrefix() string`

GetPrefix returns the Prefix field if non-nil, zero value otherwise.

### GetPrefixOk

`func (o *OsgWanStatusVO) GetPrefixOk() (*string, bool)`

GetPrefixOk returns a tuple with the Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrefix

`func (o *OsgWanStatusVO) SetPrefix(v string)`

SetPrefix sets Prefix field to given value.

### HasPrefix

`func (o *OsgWanStatusVO) HasPrefix() bool`

HasPrefix returns a boolean if a field has been set.

### GetProto

`func (o *OsgWanStatusVO) GetProto() string`

GetProto returns the Proto field if non-nil, zero value otherwise.

### GetProtoOk

`func (o *OsgWanStatusVO) GetProtoOk() (*string, bool)`

GetProtoOk returns a tuple with the Proto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProto

`func (o *OsgWanStatusVO) SetProto(v string)`

SetProto sets Proto field to given value.

### HasProto

`func (o *OsgWanStatusVO) HasProto() bool`

HasProto returns a boolean if a field has been set.

### GetType

`func (o *OsgWanStatusVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OsgWanStatusVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OsgWanStatusVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *OsgWanStatusVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


