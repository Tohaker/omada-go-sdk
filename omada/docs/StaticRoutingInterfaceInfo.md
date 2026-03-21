# StaticRoutingInterfaceInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Interface ID. | [optional] 
**InterfaceType** | Pointer to **int32** | Enter a value as follows: 0: WAN; 1: LAN; 2: L2TP, 3: PPTP, 5: Virtual WAN | [optional] 
**Name** | Pointer to **string** | Interface name. | [optional] 
**NeedNextHop** | Pointer to **bool** | Parameter [needNextHop] will be true when [interfaceType] is WAN or virtual WAN, and proto is DHCP or static. | [optional] 

## Methods

### NewStaticRoutingInterfaceInfo

`func NewStaticRoutingInterfaceInfo() *StaticRoutingInterfaceInfo`

NewStaticRoutingInterfaceInfo instantiates a new StaticRoutingInterfaceInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticRoutingInterfaceInfoWithDefaults

`func NewStaticRoutingInterfaceInfoWithDefaults() *StaticRoutingInterfaceInfo`

NewStaticRoutingInterfaceInfoWithDefaults instantiates a new StaticRoutingInterfaceInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StaticRoutingInterfaceInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StaticRoutingInterfaceInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StaticRoutingInterfaceInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StaticRoutingInterfaceInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterfaceType

`func (o *StaticRoutingInterfaceInfo) GetInterfaceType() int32`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *StaticRoutingInterfaceInfo) GetInterfaceTypeOk() (*int32, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *StaticRoutingInterfaceInfo) SetInterfaceType(v int32)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *StaticRoutingInterfaceInfo) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetName

`func (o *StaticRoutingInterfaceInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StaticRoutingInterfaceInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StaticRoutingInterfaceInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StaticRoutingInterfaceInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedNextHop

`func (o *StaticRoutingInterfaceInfo) GetNeedNextHop() bool`

GetNeedNextHop returns the NeedNextHop field if non-nil, zero value otherwise.

### GetNeedNextHopOk

`func (o *StaticRoutingInterfaceInfo) GetNeedNextHopOk() (*bool, bool)`

GetNeedNextHopOk returns a tuple with the NeedNextHop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedNextHop

`func (o *StaticRoutingInterfaceInfo) SetNeedNextHop(v bool)`

SetNeedNextHop sets NeedNextHop field to given value.

### HasNeedNextHop

`func (o *StaticRoutingInterfaceInfo) HasNeedNextHop() bool`

HasNeedNextHop returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


