# StaticRoutingInterfaceResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasWanEnable** | Pointer to **bool** | Internet has enable WAN port(s). | [optional] 
**InterfaceInfoList** | Pointer to [**[]StaticRoutingInterfaceInfo**](StaticRoutingInterfaceInfo.md) | Interface information list. | [optional] 

## Methods

### NewStaticRoutingInterfaceResult

`func NewStaticRoutingInterfaceResult() *StaticRoutingInterfaceResult`

NewStaticRoutingInterfaceResult instantiates a new StaticRoutingInterfaceResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStaticRoutingInterfaceResultWithDefaults

`func NewStaticRoutingInterfaceResultWithDefaults() *StaticRoutingInterfaceResult`

NewStaticRoutingInterfaceResultWithDefaults instantiates a new StaticRoutingInterfaceResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasWanEnable

`func (o *StaticRoutingInterfaceResult) GetHasWanEnable() bool`

GetHasWanEnable returns the HasWanEnable field if non-nil, zero value otherwise.

### GetHasWanEnableOk

`func (o *StaticRoutingInterfaceResult) GetHasWanEnableOk() (*bool, bool)`

GetHasWanEnableOk returns a tuple with the HasWanEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasWanEnable

`func (o *StaticRoutingInterfaceResult) SetHasWanEnable(v bool)`

SetHasWanEnable sets HasWanEnable field to given value.

### HasHasWanEnable

`func (o *StaticRoutingInterfaceResult) HasHasWanEnable() bool`

HasHasWanEnable returns a boolean if a field has been set.

### GetInterfaceInfoList

`func (o *StaticRoutingInterfaceResult) GetInterfaceInfoList() []StaticRoutingInterfaceInfo`

GetInterfaceInfoList returns the InterfaceInfoList field if non-nil, zero value otherwise.

### GetInterfaceInfoListOk

`func (o *StaticRoutingInterfaceResult) GetInterfaceInfoListOk() (*[]StaticRoutingInterfaceInfo, bool)`

GetInterfaceInfoListOk returns a tuple with the InterfaceInfoList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceInfoList

`func (o *StaticRoutingInterfaceResult) SetInterfaceInfoList(v []StaticRoutingInterfaceInfo)`

SetInterfaceInfoList sets InterfaceInfoList field to given value.

### HasInterfaceInfoList

`func (o *StaticRoutingInterfaceResult) HasInterfaceInfoList() bool`

HasInterfaceInfoList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


