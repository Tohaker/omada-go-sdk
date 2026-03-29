# CandidateParentOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Childsnum** | Pointer to **int32** | The number of child aps already connected to the candidate parent ap | [optional] 
**Hop** | Pointer to **int32** | The number of hops to mesh with the candidate parent AP | [optional] 
**LinkStatus** | Pointer to **int32** | 0-init; 1-linking; 2-linked; 3-link fail; 4-offline | [optional] 
**Mac** | Pointer to **string** | candidate parent ap mac | [optional] 
**Model** | Pointer to **string** | Model | [optional] 
**ModelVersion** | Pointer to **string** | Model version | [optional] 
**MultiBandInfo** | Pointer to [**[]MultiBandInfoOpenApiVO**](MultiBandInfoOpenApiVO.md) | multi band info of candidate parent ap | [optional] 
**Name** | Pointer to **string** | candidate parent ap name | [optional] 
**PriorityStatus** | Pointer to **int32** | priority parent ap status | [optional] 
**Support5gMultiBand** | Pointer to **bool** | 5G support 5G1 &amp; 5G2 multi band | [optional] 

## Methods

### NewCandidateParentOpenApiVO

`func NewCandidateParentOpenApiVO() *CandidateParentOpenApiVO`

NewCandidateParentOpenApiVO instantiates a new CandidateParentOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCandidateParentOpenApiVOWithDefaults

`func NewCandidateParentOpenApiVOWithDefaults() *CandidateParentOpenApiVO`

NewCandidateParentOpenApiVOWithDefaults instantiates a new CandidateParentOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildsnum

`func (o *CandidateParentOpenApiVO) GetChildsnum() int32`

GetChildsnum returns the Childsnum field if non-nil, zero value otherwise.

### GetChildsnumOk

`func (o *CandidateParentOpenApiVO) GetChildsnumOk() (*int32, bool)`

GetChildsnumOk returns a tuple with the Childsnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildsnum

`func (o *CandidateParentOpenApiVO) SetChildsnum(v int32)`

SetChildsnum sets Childsnum field to given value.

### HasChildsnum

`func (o *CandidateParentOpenApiVO) HasChildsnum() bool`

HasChildsnum returns a boolean if a field has been set.

### GetHop

`func (o *CandidateParentOpenApiVO) GetHop() int32`

GetHop returns the Hop field if non-nil, zero value otherwise.

### GetHopOk

`func (o *CandidateParentOpenApiVO) GetHopOk() (*int32, bool)`

GetHopOk returns a tuple with the Hop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHop

`func (o *CandidateParentOpenApiVO) SetHop(v int32)`

SetHop sets Hop field to given value.

### HasHop

`func (o *CandidateParentOpenApiVO) HasHop() bool`

HasHop returns a boolean if a field has been set.

### GetLinkStatus

`func (o *CandidateParentOpenApiVO) GetLinkStatus() int32`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *CandidateParentOpenApiVO) GetLinkStatusOk() (*int32, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *CandidateParentOpenApiVO) SetLinkStatus(v int32)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *CandidateParentOpenApiVO) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetMac

`func (o *CandidateParentOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *CandidateParentOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *CandidateParentOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *CandidateParentOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *CandidateParentOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CandidateParentOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CandidateParentOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CandidateParentOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *CandidateParentOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *CandidateParentOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *CandidateParentOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *CandidateParentOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetMultiBandInfo

`func (o *CandidateParentOpenApiVO) GetMultiBandInfo() []MultiBandInfoOpenApiVO`

GetMultiBandInfo returns the MultiBandInfo field if non-nil, zero value otherwise.

### GetMultiBandInfoOk

`func (o *CandidateParentOpenApiVO) GetMultiBandInfoOk() (*[]MultiBandInfoOpenApiVO, bool)`

GetMultiBandInfoOk returns a tuple with the MultiBandInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiBandInfo

`func (o *CandidateParentOpenApiVO) SetMultiBandInfo(v []MultiBandInfoOpenApiVO)`

SetMultiBandInfo sets MultiBandInfo field to given value.

### HasMultiBandInfo

`func (o *CandidateParentOpenApiVO) HasMultiBandInfo() bool`

HasMultiBandInfo returns a boolean if a field has been set.

### GetName

`func (o *CandidateParentOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CandidateParentOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CandidateParentOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CandidateParentOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPriorityStatus

`func (o *CandidateParentOpenApiVO) GetPriorityStatus() int32`

GetPriorityStatus returns the PriorityStatus field if non-nil, zero value otherwise.

### GetPriorityStatusOk

`func (o *CandidateParentOpenApiVO) GetPriorityStatusOk() (*int32, bool)`

GetPriorityStatusOk returns a tuple with the PriorityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityStatus

`func (o *CandidateParentOpenApiVO) SetPriorityStatus(v int32)`

SetPriorityStatus sets PriorityStatus field to given value.

### HasPriorityStatus

`func (o *CandidateParentOpenApiVO) HasPriorityStatus() bool`

HasPriorityStatus returns a boolean if a field has been set.

### GetSupport5gMultiBand

`func (o *CandidateParentOpenApiVO) GetSupport5gMultiBand() bool`

GetSupport5gMultiBand returns the Support5gMultiBand field if non-nil, zero value otherwise.

### GetSupport5gMultiBandOk

`func (o *CandidateParentOpenApiVO) GetSupport5gMultiBandOk() (*bool, bool)`

GetSupport5gMultiBandOk returns a tuple with the Support5gMultiBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupport5gMultiBand

`func (o *CandidateParentOpenApiVO) SetSupport5gMultiBand(v bool)`

SetSupport5gMultiBand sets Support5gMultiBand field to given value.

### HasSupport5gMultiBand

`func (o *CandidateParentOpenApiVO) HasSupport5gMultiBand() bool`

HasSupport5gMultiBand returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


