# OswInterfaceBriefVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID. | [optional] 
**InterfaceId** | Pointer to **int32** | Interface ID. For loopback interface, the value is loopbackId. For vlan interface, the value is vlan. | [optional] 
**Name** | Pointer to **string** | Interface Name. Not Null when type is 1. | [optional] 
**Type** | Pointer to **int32** | Interface type, 0: Loopback interface; 1: VLAN interface | [optional] 

## Methods

### NewOswInterfaceBriefVO

`func NewOswInterfaceBriefVO() *OswInterfaceBriefVO`

NewOswInterfaceBriefVO instantiates a new OswInterfaceBriefVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswInterfaceBriefVOWithDefaults

`func NewOswInterfaceBriefVOWithDefaults() *OswInterfaceBriefVO`

NewOswInterfaceBriefVOWithDefaults instantiates a new OswInterfaceBriefVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OswInterfaceBriefVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OswInterfaceBriefVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OswInterfaceBriefVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OswInterfaceBriefVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *OswInterfaceBriefVO) GetInterfaceId() int32`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *OswInterfaceBriefVO) GetInterfaceIdOk() (*int32, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *OswInterfaceBriefVO) SetInterfaceId(v int32)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *OswInterfaceBriefVO) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetName

`func (o *OswInterfaceBriefVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswInterfaceBriefVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswInterfaceBriefVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswInterfaceBriefVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *OswInterfaceBriefVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswInterfaceBriefVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswInterfaceBriefVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *OswInterfaceBriefVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


