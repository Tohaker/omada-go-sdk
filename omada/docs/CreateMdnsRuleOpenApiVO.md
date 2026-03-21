# CreateMdnsRuleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ap** | Pointer to [**ApMdnsRuleOpenApiVO**](ApMdnsRuleOpenApiVO.md) |  | [optional] 
**Name** | **string** | MDNS rule name. Name should contain 1 to 64 characters | 
**Osg** | Pointer to [**OsgMdnsRuleOpenApiVO**](OsgMdnsRuleOpenApiVO.md) |  | [optional] 
**ProfileIds** | **[]string** | This field represents Bonjour service profile ID. Bonjour Service Profile can be created using &#39;Create new Bonjour Service&#39; interface, and Bonjour service profile ID can be obtained from &#39;Get Bonjour Service list&#39; interface | 
**Status** | **bool** | MDNS rule enable status | 
**Type** | Pointer to **int32** | MDNS rule type. Type should be a value as follows: 0: AP, 1: Gateway | [optional] 

## Methods

### NewCreateMdnsRuleOpenApiVO

`func NewCreateMdnsRuleOpenApiVO(name string, profileIds []string, status bool, ) *CreateMdnsRuleOpenApiVO`

NewCreateMdnsRuleOpenApiVO instantiates a new CreateMdnsRuleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMdnsRuleOpenApiVOWithDefaults

`func NewCreateMdnsRuleOpenApiVOWithDefaults() *CreateMdnsRuleOpenApiVO`

NewCreateMdnsRuleOpenApiVOWithDefaults instantiates a new CreateMdnsRuleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAp

`func (o *CreateMdnsRuleOpenApiVO) GetAp() ApMdnsRuleOpenApiVO`

GetAp returns the Ap field if non-nil, zero value otherwise.

### GetApOk

`func (o *CreateMdnsRuleOpenApiVO) GetApOk() (*ApMdnsRuleOpenApiVO, bool)`

GetApOk returns a tuple with the Ap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAp

`func (o *CreateMdnsRuleOpenApiVO) SetAp(v ApMdnsRuleOpenApiVO)`

SetAp sets Ap field to given value.

### HasAp

`func (o *CreateMdnsRuleOpenApiVO) HasAp() bool`

HasAp returns a boolean if a field has been set.

### GetName

`func (o *CreateMdnsRuleOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateMdnsRuleOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateMdnsRuleOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetOsg

`func (o *CreateMdnsRuleOpenApiVO) GetOsg() OsgMdnsRuleOpenApiVO`

GetOsg returns the Osg field if non-nil, zero value otherwise.

### GetOsgOk

`func (o *CreateMdnsRuleOpenApiVO) GetOsgOk() (*OsgMdnsRuleOpenApiVO, bool)`

GetOsgOk returns a tuple with the Osg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsg

`func (o *CreateMdnsRuleOpenApiVO) SetOsg(v OsgMdnsRuleOpenApiVO)`

SetOsg sets Osg field to given value.

### HasOsg

`func (o *CreateMdnsRuleOpenApiVO) HasOsg() bool`

HasOsg returns a boolean if a field has been set.

### GetProfileIds

`func (o *CreateMdnsRuleOpenApiVO) GetProfileIds() []string`

GetProfileIds returns the ProfileIds field if non-nil, zero value otherwise.

### GetProfileIdsOk

`func (o *CreateMdnsRuleOpenApiVO) GetProfileIdsOk() (*[]string, bool)`

GetProfileIdsOk returns a tuple with the ProfileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileIds

`func (o *CreateMdnsRuleOpenApiVO) SetProfileIds(v []string)`

SetProfileIds sets ProfileIds field to given value.


### GetStatus

`func (o *CreateMdnsRuleOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateMdnsRuleOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateMdnsRuleOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetType

`func (o *CreateMdnsRuleOpenApiVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateMdnsRuleOpenApiVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateMdnsRuleOpenApiVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *CreateMdnsRuleOpenApiVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


