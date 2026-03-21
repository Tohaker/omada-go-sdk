# CreateOspfProcessRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserInfoBriefDTO** | Pointer to [**UserInfoBriefDTO**](UserInfoBriefDTO.md) |  | [optional] 
**Vo** | Pointer to [**OspfProcessConfigOpenApiVO**](OspfProcessConfigOpenApiVO.md) |  | [optional] 

## Methods

### NewCreateOspfProcessRequest

`func NewCreateOspfProcessRequest() *CreateOspfProcessRequest`

NewCreateOspfProcessRequest instantiates a new CreateOspfProcessRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOspfProcessRequestWithDefaults

`func NewCreateOspfProcessRequestWithDefaults() *CreateOspfProcessRequest`

NewCreateOspfProcessRequestWithDefaults instantiates a new CreateOspfProcessRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserInfoBriefDTO

`func (o *CreateOspfProcessRequest) GetUserInfoBriefDTO() UserInfoBriefDTO`

GetUserInfoBriefDTO returns the UserInfoBriefDTO field if non-nil, zero value otherwise.

### GetUserInfoBriefDTOOk

`func (o *CreateOspfProcessRequest) GetUserInfoBriefDTOOk() (*UserInfoBriefDTO, bool)`

GetUserInfoBriefDTOOk returns a tuple with the UserInfoBriefDTO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoBriefDTO

`func (o *CreateOspfProcessRequest) SetUserInfoBriefDTO(v UserInfoBriefDTO)`

SetUserInfoBriefDTO sets UserInfoBriefDTO field to given value.

### HasUserInfoBriefDTO

`func (o *CreateOspfProcessRequest) HasUserInfoBriefDTO() bool`

HasUserInfoBriefDTO returns a boolean if a field has been set.

### GetVo

`func (o *CreateOspfProcessRequest) GetVo() OspfProcessConfigOpenApiVO`

GetVo returns the Vo field if non-nil, zero value otherwise.

### GetVoOk

`func (o *CreateOspfProcessRequest) GetVoOk() (*OspfProcessConfigOpenApiVO, bool)`

GetVoOk returns a tuple with the Vo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVo

`func (o *CreateOspfProcessRequest) SetVo(v OspfProcessConfigOpenApiVO)`

SetVo sets Vo field to given value.

### HasVo

`func (o *CreateOspfProcessRequest) HasVo() bool`

HasVo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


