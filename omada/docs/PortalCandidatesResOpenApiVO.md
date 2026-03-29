# PortalCandidatesResOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkList** | Pointer to [**[]PortalNetworkOpenApiVO**](PortalNetworkOpenApiVO.md) | Network list of the portals. | [optional] 
**WlanList** | Pointer to [**[]PortalWlanOpenApiVO**](PortalWlanOpenApiVO.md) | WLAN list of the portals. | [optional] 

## Methods

### NewPortalCandidatesResOpenApiVO

`func NewPortalCandidatesResOpenApiVO() *PortalCandidatesResOpenApiVO`

NewPortalCandidatesResOpenApiVO instantiates a new PortalCandidatesResOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalCandidatesResOpenApiVOWithDefaults

`func NewPortalCandidatesResOpenApiVOWithDefaults() *PortalCandidatesResOpenApiVO`

NewPortalCandidatesResOpenApiVOWithDefaults instantiates a new PortalCandidatesResOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkList

`func (o *PortalCandidatesResOpenApiVO) GetNetworkList() []PortalNetworkOpenApiVO`

GetNetworkList returns the NetworkList field if non-nil, zero value otherwise.

### GetNetworkListOk

`func (o *PortalCandidatesResOpenApiVO) GetNetworkListOk() (*[]PortalNetworkOpenApiVO, bool)`

GetNetworkListOk returns a tuple with the NetworkList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkList

`func (o *PortalCandidatesResOpenApiVO) SetNetworkList(v []PortalNetworkOpenApiVO)`

SetNetworkList sets NetworkList field to given value.

### HasNetworkList

`func (o *PortalCandidatesResOpenApiVO) HasNetworkList() bool`

HasNetworkList returns a boolean if a field has been set.

### GetWlanList

`func (o *PortalCandidatesResOpenApiVO) GetWlanList() []PortalWlanOpenApiVO`

GetWlanList returns the WlanList field if non-nil, zero value otherwise.

### GetWlanListOk

`func (o *PortalCandidatesResOpenApiVO) GetWlanListOk() (*[]PortalWlanOpenApiVO, bool)`

GetWlanListOk returns a tuple with the WlanList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWlanList

`func (o *PortalCandidatesResOpenApiVO) SetWlanList(v []PortalWlanOpenApiVO)`

SetWlanList sets WlanList field to given value.

### HasWlanList

`func (o *PortalCandidatesResOpenApiVO) HasWlanList() bool`

HasWlanList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


