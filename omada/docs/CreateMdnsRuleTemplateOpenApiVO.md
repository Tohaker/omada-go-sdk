# CreateMdnsRuleTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | Pointer to [**ApMdnsRuleOpenApiVO**](ApMdnsRuleOpenApiVO.md) |  | [optional] 
**Name** | **string** | MDNS rule name. Name should contain 1 to 64 characters | 
**Osg** | Pointer to [**OsgMdnsRuleTemplateOpenApiVO**](OsgMdnsRuleTemplateOpenApiVO.md) |  | [optional] 
**ProfileIds** | **[]string** | This field represents Bonjour service profile ID. Bonjour Service Profile can be created using &#39;Create new Bonjour Service&#39; interface, and Bonjour service profile ID can be obtained from &#39;Get Bonjour Service list&#39; interface | 
**Status** | **bool** | MDNS rule enable status | 
**Type** | Pointer to **int32** | MDNS rule type. Type should be a value as follows: 0: AP, 1: Gateway | [optional] 

## Methods

### NewCreateMdnsRuleTemplateOpenApiVO

`func NewCreateMdnsRuleTemplateOpenApiVO(name string, profileIds []string, status bool, ) *CreateMdnsRuleTemplateOpenApiVO`

NewCreateMdnsRuleTemplateOpenApiVO instantiates a new CreateMdnsRuleTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMdnsRuleTemplateOpenApiVOWithDefaults

`func NewCreateMdnsRuleTemplateOpenApiVOWithDefaults() *CreateMdnsRuleTemplateOpenApiVO`

NewCreateMdnsRuleTemplateOpenApiVOWithDefaults instantiates a new CreateMdnsRuleTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *CreateMdnsRuleTemplateOpenApiVO) GetAp() ApMdnsRuleOpenApiVO`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *CreateMdnsRuleTemplateOpenApiVO) GetApOk() (*ApMdnsRuleOpenApiVO, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *CreateMdnsRuleTemplateOpenApiVO) SetAp(v ApMdnsRuleOpenApiVO)`

SetAp sets Ap field to given value.

### HasAp

`func (o *CreateMdnsRuleTemplateOpenApiVO) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetName

`func (o *CreateMdnsRuleTemplateOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMdnsRuleTemplateOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMdnsRuleTemplateOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetOsg

`func (o *CreateMdnsRuleTemplateOpenApiVO) GetOsg() OsgMdnsRuleTemplateOpenApiVO`

GetOsg returns the Osg field if non-nil, zero value otherwise.

### GetOsgOk

`func (o *CreateMdnsRuleTemplateOpenApiVO) GetOsgOk() (*OsgMdnsRuleTemplateOpenApiVO, bool)`

GetOsgOk returns a tuple with the Osg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsg

`func (o *CreateMdnsRuleTemplateOpenApiVO) SetOsg(v OsgMdnsRuleTemplateOpenApiVO)`

SetOsg sets Osg field to given value.

### HasOsg

`func (o *CreateMdnsRuleTemplateOpenApiVO) HasOsg() bool`

HasOsg returns a boolean if a field has been set.

### GetProfileIds

`func (o *CreateMdnsRuleTemplateOpenApiVO) GetProfileIds() []string`

GetProfileIds returns the ProfileIds field if non-nil, zero value otherwise.

### GetProfileIdsOk

`func (o *CreateMdnsRuleTemplateOpenApiVO) GetProfileIdsOk() (*[]string, bool)`

GetProfileIdsOk returns a tuple with the ProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIds

`func (o *CreateMdnsRuleTemplateOpenApiVO) SetProfileIds(v []string)`

SetProfileIds sets ProfileIds field to given value.


### GetStatus

`func (o *CreateMdnsRuleTemplateOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateMdnsRuleTemplateOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateMdnsRuleTemplateOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetType

`func (o *CreateMdnsRuleTemplateOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateMdnsRuleTemplateOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateMdnsRuleTemplateOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CreateMdnsRuleTemplateOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


