# RadiusServerBriefInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuiltInServer** | Pointer to **bool** | Is this Radius server a built-in server | [optional] 
**Id** | **string** | radius server id. | 
**Name** | **string** | radius server name. | 
**OverNumLimit** | Pointer to **bool** | Is the number of authentication/accounting Servers greater than 2. | [optional] 

## Methods

### NewRadiusServerBriefInfo

`func NewRadiusServerBriefInfo(id string, name string, ) *RadiusServerBriefInfo`

NewRadiusServerBriefInfo instantiates a new RadiusServerBriefInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusServerBriefInfoWithDefaults

`func NewRadiusServerBriefInfoWithDefaults() *RadiusServerBriefInfo`

NewRadiusServerBriefInfoWithDefaults instantiates a new RadiusServerBriefInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuiltInServer

`func (o *RadiusServerBriefInfo) GetBuiltInServer() bool`

GetBuiltInServer returns the BuiltInServer field if non-nil, zero value otherwise.

### GetBuiltInServerOk

`func (o *RadiusServerBriefInfo) GetBuiltInServerOk() (*bool, bool)`

GetBuiltInServerOk returns a tuple with the BuiltInServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltInServer

`func (o *RadiusServerBriefInfo) SetBuiltInServer(v bool)`

SetBuiltInServer sets BuiltInServer field to given value.

### HasBuiltInServer

`func (o *RadiusServerBriefInfo) HasBuiltInServer() bool`

HasBuiltInServer returns a boolean if a field has been set.

### GetId

`func (o *RadiusServerBriefInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RadiusServerBriefInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RadiusServerBriefInfo) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *RadiusServerBriefInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *RadiusServerBriefInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *RadiusServerBriefInfo) SetName(v string)`

SetName sets Name field to given value.


### GetOverNumLimit

`func (o *RadiusServerBriefInfo) GetOverNumLimit() bool`

GetOverNumLimit returns the OverNumLimit field if non-nil, zero value otherwise.

### GetOverNumLimitOk

`func (o *RadiusServerBriefInfo) GetOverNumLimitOk() (*bool, bool)`

GetOverNumLimitOk returns a tuple with the OverNumLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverNumLimit

`func (o *RadiusServerBriefInfo) SetOverNumLimit(v bool)`

SetOverNumLimit sets OverNumLimit field to given value.

### HasOverNumLimit

`func (o *RadiusServerBriefInfo) HasOverNumLimit() bool`

HasOverNumLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


