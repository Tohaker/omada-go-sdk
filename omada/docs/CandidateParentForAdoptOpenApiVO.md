# CandidateParentForAdoptOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Childsnum** | Pointer to **int32** | The number of child aps already connected to the candidate parent ap | [optional] 
**Hop** | Pointer to **int32** | The number of hops to mesh with the candidate parent AP | [optional] 
**LinkStatus** | Pointer to **int32** | 0-init; 1-linking; 2-linked; 3-link fail; 4-offline | [optional] 
**Mac** | Pointer to **string** | candidate parent ap mac | [optional] 
**MultiBandInfo** | Pointer to [**[]MultiBandInfoOpenApiVO**](MultiBandInfoOpenApiVO.md) | multi band info of candidate parent ap | [optional] 
**Name** | Pointer to **string** | candidate parent ap name | [optional] 
**Recommend** | Pointer to **bool** | Whether the ap is recommended by the algorithm for connection | [optional] 
**RecommendRadioId** | Pointer to **int32** | The radio ID for mesh recommended by the algorithm | [optional] 

## Methods

### NewCandidateParentForAdoptOpenApiVO

`func NewCandidateParentForAdoptOpenApiVO() *CandidateParentForAdoptOpenApiVO`

NewCandidateParentForAdoptOpenApiVO instantiates a new CandidateParentForAdoptOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCandidateParentForAdoptOpenApiVOWithDefaults

`func NewCandidateParentForAdoptOpenApiVOWithDefaults() *CandidateParentForAdoptOpenApiVO`

NewCandidateParentForAdoptOpenApiVOWithDefaults instantiates a new CandidateParentForAdoptOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildsnum

`func (o *CandidateParentForAdoptOpenApiVO) GetChildsnum() int32`

GetChildsnum returns the Childsnum field if non-nil, zero value otherwise.

### GetChildsnumOk

`func (o *CandidateParentForAdoptOpenApiVO) GetChildsnumOk() (*int32, bool)`

GetChildsnumOk returns a tuple with the Childsnum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildsnum

`func (o *CandidateParentForAdoptOpenApiVO) SetChildsnum(v int32)`

SetChildsnum sets Childsnum field to given value.

### HasChildsnum

`func (o *CandidateParentForAdoptOpenApiVO) HasChildsnum() bool`

HasChildsnum returns a boolean if a field has been set.

### GetHop

`func (o *CandidateParentForAdoptOpenApiVO) GetHop() int32`

GetHop returns the Hop field if non-nil, zero value otherwise.

### GetHopOk

`func (o *CandidateParentForAdoptOpenApiVO) GetHopOk() (*int32, bool)`

GetHopOk returns a tuple with the Hop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHop

`func (o *CandidateParentForAdoptOpenApiVO) SetHop(v int32)`

SetHop sets Hop field to given value.

### HasHop

`func (o *CandidateParentForAdoptOpenApiVO) HasHop() bool`

HasHop returns a boolean if a field has been set.

### GetLinkStatus

`func (o *CandidateParentForAdoptOpenApiVO) GetLinkStatus() int32`

GetLinkStatus returns the LinkStatus field if non-nil, zero value otherwise.

### GetLinkStatusOk

`func (o *CandidateParentForAdoptOpenApiVO) GetLinkStatusOk() (*int32, bool)`

GetLinkStatusOk returns a tuple with the LinkStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkStatus

`func (o *CandidateParentForAdoptOpenApiVO) SetLinkStatus(v int32)`

SetLinkStatus sets LinkStatus field to given value.

### HasLinkStatus

`func (o *CandidateParentForAdoptOpenApiVO) HasLinkStatus() bool`

HasLinkStatus returns a boolean if a field has been set.

### GetMac

`func (o *CandidateParentForAdoptOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *CandidateParentForAdoptOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *CandidateParentForAdoptOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *CandidateParentForAdoptOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMultiBandInfo

`func (o *CandidateParentForAdoptOpenApiVO) GetMultiBandInfo() []MultiBandInfoOpenApiVO`

GetMultiBandInfo returns the MultiBandInfo field if non-nil, zero value otherwise.

### GetMultiBandInfoOk

`func (o *CandidateParentForAdoptOpenApiVO) GetMultiBandInfoOk() (*[]MultiBandInfoOpenApiVO, bool)`

GetMultiBandInfoOk returns a tuple with the MultiBandInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiBandInfo

`func (o *CandidateParentForAdoptOpenApiVO) SetMultiBandInfo(v []MultiBandInfoOpenApiVO)`

SetMultiBandInfo sets MultiBandInfo field to given value.

### HasMultiBandInfo

`func (o *CandidateParentForAdoptOpenApiVO) HasMultiBandInfo() bool`

HasMultiBandInfo returns a boolean if a field has been set.

### GetName

`func (o *CandidateParentForAdoptOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CandidateParentForAdoptOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CandidateParentForAdoptOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CandidateParentForAdoptOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRecommend

`func (o *CandidateParentForAdoptOpenApiVO) GetRecommend() bool`

GetRecommend returns the Recommend field if non-nil, zero value otherwise.

### GetRecommendOk

`func (o *CandidateParentForAdoptOpenApiVO) GetRecommendOk() (*bool, bool)`

GetRecommendOk returns a tuple with the Recommend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommend

`func (o *CandidateParentForAdoptOpenApiVO) SetRecommend(v bool)`

SetRecommend sets Recommend field to given value.

### HasRecommend

`func (o *CandidateParentForAdoptOpenApiVO) HasRecommend() bool`

HasRecommend returns a boolean if a field has been set.

### GetRecommendRadioId

`func (o *CandidateParentForAdoptOpenApiVO) GetRecommendRadioId() int32`

GetRecommendRadioId returns the RecommendRadioId field if non-nil, zero value otherwise.

### GetRecommendRadioIdOk

`func (o *CandidateParentForAdoptOpenApiVO) GetRecommendRadioIdOk() (*int32, bool)`

GetRecommendRadioIdOk returns a tuple with the RecommendRadioId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecommendRadioId

`func (o *CandidateParentForAdoptOpenApiVO) SetRecommendRadioId(v int32)`

SetRecommendRadioId sets RecommendRadioId field to given value.

### HasRecommendRadioId

`func (o *CandidateParentForAdoptOpenApiVO) HasRecommendRadioId() bool`

HasRecommendRadioId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


