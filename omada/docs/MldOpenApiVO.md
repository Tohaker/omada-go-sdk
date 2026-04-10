# MldOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enable** | **bool** |  | 
**Version** | **int32** | Version should be one of the following values: 1: v1; 2: v2. | 
**WanPortId** | Pointer to **string** | WAN port ID can be obtained from &#39;Get internet basic info&#39; interface. At least one of the WAN Port IDs should not be null. Only IPv6-enabled WAN ports can be selected as MLD Interface. MLD does not support the 6to4 Tunnel and Pass-Through(Bridge) IPv6 dial-up modes. | [optional] 

## Methods

### NewMldOpenApiVO

`func NewMldOpenApiVO(enable bool, version int32, ) *MldOpenApiVO`

NewMldOpenApiVO instantiates a new MldOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMldOpenApiVOWithDefaults

`func NewMldOpenApiVOWithDefaults() *MldOpenApiVO`

NewMldOpenApiVOWithDefaults instantiates a new MldOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnable

`func (o *MldOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *MldOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *MldOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetVersion

`func (o *MldOpenApiVO) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *MldOpenApiVO) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *MldOpenApiVO) SetVersion(v int32)`

SetVersion sets Version field to given value.


### GetWanPortId

`func (o *MldOpenApiVO) GetWanPortId() string`

GetWanPortId returns the WanPortId field if non-nil, zero value otherwise.

### GetWanPortIdOk

`func (o *MldOpenApiVO) GetWanPortIdOk() (*string, bool)`

GetWanPortIdOk returns a tuple with the WanPortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWanPortId

`func (o *MldOpenApiVO) SetWanPortId(v string)`

SetWanPortId sets WanPortId field to given value.

### HasWanPortId

`func (o *MldOpenApiVO) HasWanPortId() bool`

HasWanPortId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


