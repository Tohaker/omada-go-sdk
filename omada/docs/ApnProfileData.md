# ApnProfileData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]ApnProfile**](ApnProfile.md) | Apn profile list | [optional] 
**SupportDualSim** | Pointer to **int32** | Support dual sim card or not | [optional] 

## Methods

### NewApnProfileData

`func NewApnProfileData() *ApnProfileData`

NewApnProfileData instantiates a new ApnProfileData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApnProfileDataWithDefaults

`func NewApnProfileDataWithDefaults() *ApnProfileData`

NewApnProfileDataWithDefaults instantiates a new ApnProfileData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ApnProfileData) GetData() []ApnProfile`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ApnProfileData) GetDataOk() (*[]ApnProfile, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ApnProfileData) SetData(v []ApnProfile)`

SetData sets Data field to given value.

### HasData

`func (o *ApnProfileData) HasData() bool`

HasData returns a boolean if a field has been set.

### GetSupportDualSim

`func (o *ApnProfileData) GetSupportDualSim() int32`

GetSupportDualSim returns the SupportDualSim field if non-nil, zero value otherwise.

### GetSupportDualSimOk

`func (o *ApnProfileData) GetSupportDualSimOk() (*int32, bool)`

GetSupportDualSimOk returns a tuple with the SupportDualSim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportDualSim

`func (o *ApnProfileData) SetSupportDualSim(v int32)`

SetSupportDualSim sets SupportDualSim field to given value.

### HasSupportDualSim

`func (o *ApnProfileData) HasSupportDualSim() bool`

HasSupportDualSim returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


