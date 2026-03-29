# PpskProfileBriefInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | PPSK Profile ID | [optional] 
**ProfileName** | **string** | PPSK Profile Name, should contain 1 to 64 characters. | 
**Ssid** | Pointer to **[]string** | SSIDs Bound With PPSK Profile | [optional] 

## Methods

### NewPpskProfileBriefInfo

`func NewPpskProfileBriefInfo(profileName string, ) *PpskProfileBriefInfo`

NewPpskProfileBriefInfo instantiates a new PpskProfileBriefInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPpskProfileBriefInfoWithDefaults

`func NewPpskProfileBriefInfoWithDefaults() *PpskProfileBriefInfo`

NewPpskProfileBriefInfoWithDefaults instantiates a new PpskProfileBriefInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PpskProfileBriefInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PpskProfileBriefInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PpskProfileBriefInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PpskProfileBriefInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProfileName

`func (o *PpskProfileBriefInfo) GetProfileName() string`

GetProfileName returns the ProfileName field if non-nil, zero value otherwise.

### GetProfileNameOk

`func (o *PpskProfileBriefInfo) GetProfileNameOk() (*string, bool)`

GetProfileNameOk returns a tuple with the ProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfileName

`func (o *PpskProfileBriefInfo) SetProfileName(v string)`

SetProfileName sets ProfileName field to given value.


### GetSsid

`func (o *PpskProfileBriefInfo) GetSsid() []string`

GetSsid returns the Ssid field if non-nil, zero value otherwise.

### GetSsidOk

`func (o *PpskProfileBriefInfo) GetSsidOk() (*[]string, bool)`

GetSsidOk returns a tuple with the Ssid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSsid

`func (o *PpskProfileBriefInfo) SetSsid(v []string)`

SetSsid sets Ssid field to given value.

### HasSsid

`func (o *PpskProfileBriefInfo) HasSsid() bool`

HasSsid returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


