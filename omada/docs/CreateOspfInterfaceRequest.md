# CreateOspfInterfaceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserInfoBriefDTO** | Pointer to [**UserInfoBriefDTO**](UserInfoBriefDTO.md) |  | [optional] 
**Vo** | Pointer to [**OspfInterfaceConfigOpenApiVO**](OspfInterfaceConfigOpenApiVO.md) |  | [optional] 

## Methods

### NewCreateOspfInterfaceRequest

`func NewCreateOspfInterfaceRequest() *CreateOspfInterfaceRequest`

NewCreateOspfInterfaceRequest instantiates a new CreateOspfInterfaceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOspfInterfaceRequestWithDefaults

`func NewCreateOspfInterfaceRequestWithDefaults() *CreateOspfInterfaceRequest`

NewCreateOspfInterfaceRequestWithDefaults instantiates a new CreateOspfInterfaceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserInfoBriefDTO

`func (o *CreateOspfInterfaceRequest) GetUserInfoBriefDTO() UserInfoBriefDTO`

GetUserInfoBriefDTO returns the UserInfoBriefDTO field if non-nil, zero value otherwise.

### GetUserInfoBriefDTOOk

`func (o *CreateOspfInterfaceRequest) GetUserInfoBriefDTOOk() (*UserInfoBriefDTO, bool)`

GetUserInfoBriefDTOOk returns a tuple with the UserInfoBriefDTO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserInfoBriefDTO

`func (o *CreateOspfInterfaceRequest) SetUserInfoBriefDTO(v UserInfoBriefDTO)`

SetUserInfoBriefDTO sets UserInfoBriefDTO field to given value.

### HasUserInfoBriefDTO

`func (o *CreateOspfInterfaceRequest) HasUserInfoBriefDTO() bool`

HasUserInfoBriefDTO returns a boolean if a field has been set.

### GetVo

`func (o *CreateOspfInterfaceRequest) GetVo() OspfInterfaceConfigOpenApiVO`

GetVo returns the Vo field if non-nil, zero value otherwise.

### GetVoOk

`func (o *CreateOspfInterfaceRequest) GetVoOk() (*OspfInterfaceConfigOpenApiVO, bool)`

GetVoOk returns a tuple with the Vo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVo

`func (o *CreateOspfInterfaceRequest) SetVo(v OspfInterfaceConfigOpenApiVO)`

SetVo sets Vo field to given value.

### HasVo

`func (o *CreateOspfInterfaceRequest) HasVo() bool`

HasVo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


