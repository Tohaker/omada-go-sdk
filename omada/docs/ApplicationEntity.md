# ApplicationEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int32** | Application ID | [optional] 
**ApplicationName** | Pointer to **string** | Application name | [optional] 
**Description** | Pointer to **string** | Description of application | [optional] 
**Family** | Pointer to **string** | Family of application | [optional] 

## Methods

### NewApplicationEntity

`func NewApplicationEntity() *ApplicationEntity`

NewApplicationEntity instantiates a new ApplicationEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationEntityWithDefaults

`func NewApplicationEntityWithDefaults() *ApplicationEntity`

NewApplicationEntityWithDefaults instantiates a new ApplicationEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *ApplicationEntity) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *ApplicationEntity) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *ApplicationEntity) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *ApplicationEntity) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationName

`func (o *ApplicationEntity) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *ApplicationEntity) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *ApplicationEntity) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *ApplicationEntity) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetDescription

`func (o *ApplicationEntity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApplicationEntity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApplicationEntity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApplicationEntity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFamily

`func (o *ApplicationEntity) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *ApplicationEntity) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *ApplicationEntity) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *ApplicationEntity) HasFamily() bool`

HasFamily returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


