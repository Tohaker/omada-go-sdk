# MdnsRuleTemplateOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | Pointer to [**ApMdnsRuleOpenApiVO**](ApMdnsRuleOpenApiVO.md) |  | [optional] 
**Id** | Pointer to **string** | MDNS rule ID | [optional] 
**Name** | Pointer to **string** | MDNS rule name | [optional] 
**Osg** | Pointer to [**OsgMdnsRuleTemplateOpenApiVO**](OsgMdnsRuleTemplateOpenApiVO.md) |  | [optional] 
**ProfileIds** | Pointer to **[]string** | ID list of selected Bonjour Service Profile | [optional] 
**Status** | Pointer to **bool** | MDNS rule enable status | [optional] 
**Type** | Pointer to **int32** | MDNS rule type, 0: AP, 1: Gateway | [optional] 

## Methods

### NewMdnsRuleTemplateOpenApiVO

`func NewMdnsRuleTemplateOpenApiVO() *MdnsRuleTemplateOpenApiVO`

NewMdnsRuleTemplateOpenApiVO instantiates a new MdnsRuleTemplateOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMdnsRuleTemplateOpenApiVOWithDefaults

`func NewMdnsRuleTemplateOpenApiVOWithDefaults() *MdnsRuleTemplateOpenApiVO`

NewMdnsRuleTemplateOpenApiVOWithDefaults instantiates a new MdnsRuleTemplateOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *MdnsRuleTemplateOpenApiVO) GetAp() ApMdnsRuleOpenApiVO`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *MdnsRuleTemplateOpenApiVO) GetApOk() (*ApMdnsRuleOpenApiVO, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *MdnsRuleTemplateOpenApiVO) SetAp(v ApMdnsRuleOpenApiVO)`

SetAp sets Ap field to given value.

### HasAp

`func (o *MdnsRuleTemplateOpenApiVO) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetId

`func (o *MdnsRuleTemplateOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MdnsRuleTemplateOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MdnsRuleTemplateOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MdnsRuleTemplateOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MdnsRuleTemplateOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MdnsRuleTemplateOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MdnsRuleTemplateOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MdnsRuleTemplateOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOsg

`func (o *MdnsRuleTemplateOpenApiVO) GetOsg() OsgMdnsRuleTemplateOpenApiVO`

GetOsg returns the Osg field if non-nil, zero value otherwise.

### GetOsgOk

`func (o *MdnsRuleTemplateOpenApiVO) GetOsgOk() (*OsgMdnsRuleTemplateOpenApiVO, bool)`

GetOsgOk returns a tuple with the Osg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsg

`func (o *MdnsRuleTemplateOpenApiVO) SetOsg(v OsgMdnsRuleTemplateOpenApiVO)`

SetOsg sets Osg field to given value.

### HasOsg

`func (o *MdnsRuleTemplateOpenApiVO) HasOsg() bool`

HasOsg returns a boolean if a field has been set.

### GetProfileIds

`func (o *MdnsRuleTemplateOpenApiVO) GetProfileIds() []string`

GetProfileIds returns the ProfileIds field if non-nil, zero value otherwise.

### GetProfileIdsOk

`func (o *MdnsRuleTemplateOpenApiVO) GetProfileIdsOk() (*[]string, bool)`

GetProfileIdsOk returns a tuple with the ProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIds

`func (o *MdnsRuleTemplateOpenApiVO) SetProfileIds(v []string)`

SetProfileIds sets ProfileIds field to given value.

### HasProfileIds

`func (o *MdnsRuleTemplateOpenApiVO) HasProfileIds() bool`

HasProfileIds returns a boolean if a field has been set.

### GetStatus

`func (o *MdnsRuleTemplateOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MdnsRuleTemplateOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MdnsRuleTemplateOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MdnsRuleTemplateOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *MdnsRuleTemplateOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MdnsRuleTemplateOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MdnsRuleTemplateOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *MdnsRuleTemplateOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


