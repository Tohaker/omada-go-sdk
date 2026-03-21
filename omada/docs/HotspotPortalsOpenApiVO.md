# HotspotPortalsOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplyToAllPortals** | Pointer to **bool** | Is the localuser effective for all portals, including all newly created portals | [optional] 
**Portals** | **[]string** | Bound portal ID list. | 

## Methods

### NewHotspotPortalsOpenApiVO

`func NewHotspotPortalsOpenApiVO(portals []string, ) *HotspotPortalsOpenApiVO`

NewHotspotPortalsOpenApiVO instantiates a new HotspotPortalsOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHotspotPortalsOpenApiVOWithDefaults

`func NewHotspotPortalsOpenApiVOWithDefaults() *HotspotPortalsOpenApiVO`

NewHotspotPortalsOpenApiVOWithDefaults instantiates a new HotspotPortalsOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplyToAllPortals

`func (o *HotspotPortalsOpenApiVO) GetApplyToAllPortals() bool`

GetApplyToAllPortals returns the ApplyToAllPortals field if non-nil, zero value otherwise.

### GetApplyToAllPortalsOk

`func (o *HotspotPortalsOpenApiVO) GetApplyToAllPortalsOk() (*bool, bool)`

GetApplyToAllPortalsOk returns a tuple with the ApplyToAllPortals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplyToAllPortals

`func (o *HotspotPortalsOpenApiVO) SetApplyToAllPortals(v bool)`

SetApplyToAllPortals sets ApplyToAllPortals field to given value.

### HasApplyToAllPortals

`func (o *HotspotPortalsOpenApiVO) HasApplyToAllPortals() bool`

HasApplyToAllPortals returns a boolean if a field has been set.

### GetPortals

`func (o *HotspotPortalsOpenApiVO) GetPortals() []string`

GetPortals returns the Portals field if non-nil, zero value otherwise.

### GetPortalsOk

`func (o *HotspotPortalsOpenApiVO) GetPortalsOk() (*[]string, bool)`

GetPortalsOk returns a tuple with the Portals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortals

`func (o *HotspotPortalsOpenApiVO) SetPortals(v []string)`

SetPortals sets Portals field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


