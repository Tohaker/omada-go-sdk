# OswInterfaceOpenApi

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceId** | Pointer to **int32** | Interface ID. | [optional] 
**InterfaceType** | **int32** | Interface type. 0: Loopback; 1: VLAN. | 
**LoopbackInterface** | Pointer to [**OswLoopbackInterfaceVO**](OswLoopbackInterfaceVO.md) |  | [optional] 
**Name** | Pointer to **string** | Interface name. | [optional] 
**Status** | Pointer to **int32** | Interface status. 0: disable, 1: enable. | [optional] 
**VlanInterface** | Pointer to [**OswNetworkVO**](OswNetworkVO.md) |  | [optional] 

## Methods

### NewOswInterfaceOpenApi

`func NewOswInterfaceOpenApi(interfaceType int32, ) *OswInterfaceOpenApi`

NewOswInterfaceOpenApi instantiates a new OswInterfaceOpenApi object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswInterfaceOpenApiWithDefaults

`func NewOswInterfaceOpenApiWithDefaults() *OswInterfaceOpenApi`

NewOswInterfaceOpenApiWithDefaults instantiates a new OswInterfaceOpenApi object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceId

`func (o *OswInterfaceOpenApi) GetInterfaceId() int32`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *OswInterfaceOpenApi) GetInterfaceIdOk() (*int32, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *OswInterfaceOpenApi) SetInterfaceId(v int32)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *OswInterfaceOpenApi) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceType

`func (o *OswInterfaceOpenApi) GetInterfaceType() int32`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *OswInterfaceOpenApi) GetInterfaceTypeOk() (*int32, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *OswInterfaceOpenApi) SetInterfaceType(v int32)`

SetInterfaceType sets InterfaceType field to given value.


### GetLoopbackInterface

`func (o *OswInterfaceOpenApi) GetLoopbackInterface() OswLoopbackInterfaceVO`

GetLoopbackInterface returns the LoopbackInterface field if non-nil, zero value otherwise.

### GetLoopbackInterfaceOk

`func (o *OswInterfaceOpenApi) GetLoopbackInterfaceOk() (*OswLoopbackInterfaceVO, bool)`

GetLoopbackInterfaceOk returns a tuple with the LoopbackInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoopbackInterface

`func (o *OswInterfaceOpenApi) SetLoopbackInterface(v OswLoopbackInterfaceVO)`

SetLoopbackInterface sets LoopbackInterface field to given value.

### HasLoopbackInterface

`func (o *OswInterfaceOpenApi) HasLoopbackInterface() bool`

HasLoopbackInterface returns a boolean if a field has been set.

### GetName

`func (o *OswInterfaceOpenApi) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswInterfaceOpenApi) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswInterfaceOpenApi) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswInterfaceOpenApi) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *OswInterfaceOpenApi) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswInterfaceOpenApi) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswInterfaceOpenApi) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswInterfaceOpenApi) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetVlanInterface

`func (o *OswInterfaceOpenApi) GetVlanInterface() OswNetworkVO`

GetVlanInterface returns the VlanInterface field if non-nil, zero value otherwise.

### GetVlanInterfaceOk

`func (o *OswInterfaceOpenApi) GetVlanInterfaceOk() (*OswNetworkVO, bool)`

GetVlanInterfaceOk returns a tuple with the VlanInterface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanInterface

`func (o *OswInterfaceOpenApi) SetVlanInterface(v OswNetworkVO)`

SetVlanInterface sets VlanInterface field to given value.

### HasVlanInterface

`func (o *OswInterfaceOpenApi) HasVlanInterface() bool`

HasVlanInterface returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


