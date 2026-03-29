# OswProfileOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mac** | Pointer to **string** | Switch MAC | [optional] 
**Name** | Pointer to **string** | Switch name | [optional] 
**Profiles** | Pointer to **map[string]string** | Switch port profiles | [optional] 

## Methods

### NewOswProfileOpenApiVO

`func NewOswProfileOpenApiVO() *OswProfileOpenApiVO`

NewOswProfileOpenApiVO instantiates a new OswProfileOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswProfileOpenApiVOWithDefaults

`func NewOswProfileOpenApiVOWithDefaults() *OswProfileOpenApiVO`

NewOswProfileOpenApiVOWithDefaults instantiates a new OswProfileOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMac

`func (o *OswProfileOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *OswProfileOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *OswProfileOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *OswProfileOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetName

`func (o *OswProfileOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswProfileOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswProfileOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswProfileOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfiles

`func (o *OswProfileOpenApiVO) GetProfiles() map[string]string`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *OswProfileOpenApiVO) GetProfilesOk() (*map[string]string, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *OswProfileOpenApiVO) SetProfiles(v map[string]string)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *OswProfileOpenApiVO) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


