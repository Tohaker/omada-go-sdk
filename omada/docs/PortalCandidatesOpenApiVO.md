# PortalCandidatesOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllPortal** | **bool** | Whether portal privilege is true | 
**PortalIds** | Pointer to **[]string** | Portal IDs of the requested SSIDs and networks. | [optional] 

## Methods

### NewPortalCandidatesOpenApiVO

`func NewPortalCandidatesOpenApiVO(allPortal bool, ) *PortalCandidatesOpenApiVO`

NewPortalCandidatesOpenApiVO instantiates a new PortalCandidatesOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortalCandidatesOpenApiVOWithDefaults

`func NewPortalCandidatesOpenApiVOWithDefaults() *PortalCandidatesOpenApiVO`

NewPortalCandidatesOpenApiVOWithDefaults instantiates a new PortalCandidatesOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllPortal

`func (o *PortalCandidatesOpenApiVO) GetAllPortal() bool`

GetAllPortal returns the AllPortal field if non-nil, zero value otherwise.

### GetAllPortalOk

`func (o *PortalCandidatesOpenApiVO) GetAllPortalOk() (*bool, bool)`

GetAllPortalOk returns a tuple with the AllPortal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllPortal

`func (o *PortalCandidatesOpenApiVO) SetAllPortal(v bool)`

SetAllPortal sets AllPortal field to given value.


### GetPortalIds

`func (o *PortalCandidatesOpenApiVO) GetPortalIds() []string`

GetPortalIds returns the PortalIds field if non-nil, zero value otherwise.

### GetPortalIdsOk

`func (o *PortalCandidatesOpenApiVO) GetPortalIdsOk() (*[]string, bool)`

GetPortalIdsOk returns a tuple with the PortalIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalIds

`func (o *PortalCandidatesOpenApiVO) SetPortalIds(v []string)`

SetPortalIds sets PortalIds field to given value.

### HasPortalIds

`func (o *PortalCandidatesOpenApiVO) HasPortalIds() bool`

HasPortalIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


