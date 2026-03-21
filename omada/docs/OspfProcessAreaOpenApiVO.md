# OspfProcessAreaOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AreaId** | **string** | The 32 bit unsigned integer that uniquely identifies the area. It can be in decimal format or dotted decimal format. | 
**AreaType** | **int32** | OSPF area type should be a value as follows: 0: Normal, 1: Stub, or 2: NSSA. | 
**NetworkList** | Pointer to [**[]OspfProcessAreaNetworkOpenApiVO**](OspfProcessAreaNetworkOpenApiVO.md) | Up to 16 entries are allowed for the networkList. | [optional] 

## Methods

### NewOspfProcessAreaOpenApiVO

`func NewOspfProcessAreaOpenApiVO(areaId string, areaType int32, ) *OspfProcessAreaOpenApiVO`

NewOspfProcessAreaOpenApiVO instantiates a new OspfProcessAreaOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOspfProcessAreaOpenApiVOWithDefaults

`func NewOspfProcessAreaOpenApiVOWithDefaults() *OspfProcessAreaOpenApiVO`

NewOspfProcessAreaOpenApiVOWithDefaults instantiates a new OspfProcessAreaOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAreaId

`func (o *OspfProcessAreaOpenApiVO) GetAreaId() string`

GetAreaId returns the AreaId field if non-nil, zero value otherwise.

### GetAreaIdOk

`func (o *OspfProcessAreaOpenApiVO) GetAreaIdOk() (*string, bool)`

GetAreaIdOk returns a tuple with the AreaId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaId

`func (o *OspfProcessAreaOpenApiVO) SetAreaId(v string)`

SetAreaId sets AreaId field to given value.


### GetAreaType

`func (o *OspfProcessAreaOpenApiVO) GetAreaType() int32`

GetAreaType returns the AreaType field if non-nil, zero value otherwise.

### GetAreaTypeOk

`func (o *OspfProcessAreaOpenApiVO) GetAreaTypeOk() (*int32, bool)`

GetAreaTypeOk returns a tuple with the AreaType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAreaType

`func (o *OspfProcessAreaOpenApiVO) SetAreaType(v int32)`

SetAreaType sets AreaType field to given value.


### GetNetworkList

`func (o *OspfProcessAreaOpenApiVO) GetNetworkList() []OspfProcessAreaNetworkOpenApiVO`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *OspfProcessAreaOpenApiVO) GetNetworkListOk() (*[]OspfProcessAreaNetworkOpenApiVO, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *OspfProcessAreaOpenApiVO) SetNetworkList(v []OspfProcessAreaNetworkOpenApiVO)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *OspfProcessAreaOpenApiVO) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


