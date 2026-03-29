# EditUploadFirmwareInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of firmware | [optional] 
**TargetEnable** | Pointer to **bool** | Do the sites set up specified firmware, it should be a value as follows: true, false | [optional] 
**TargetSites** | Pointer to **[]string** | Target sites ID, it exists when \&quot;targetEnable\&quot; is true | [optional] 

## Methods

### NewEditUploadFirmwareInfo

`func NewEditUploadFirmwareInfo() *EditUploadFirmwareInfo`

NewEditUploadFirmwareInfo instantiates a new EditUploadFirmwareInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEditUploadFirmwareInfoWithDefaults

`func NewEditUploadFirmwareInfoWithDefaults() *EditUploadFirmwareInfo`

NewEditUploadFirmwareInfoWithDefaults instantiates a new EditUploadFirmwareInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *EditUploadFirmwareInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *EditUploadFirmwareInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *EditUploadFirmwareInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *EditUploadFirmwareInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTargetEnable

`func (o *EditUploadFirmwareInfo) GetTargetEnable() bool`

GetTargetEnable returns the TargetEnable field if non-nil, zero value otherwise.

### GetTargetEnableOk

`func (o *EditUploadFirmwareInfo) GetTargetEnableOk() (*bool, bool)`

GetTargetEnableOk returns a tuple with the TargetEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetEnable

`func (o *EditUploadFirmwareInfo) SetTargetEnable(v bool)`

SetTargetEnable sets TargetEnable field to given value.

### HasTargetEnable

`func (o *EditUploadFirmwareInfo) HasTargetEnable() bool`

HasTargetEnable returns a boolean if a field has been set.

### GetTargetSites

`func (o *EditUploadFirmwareInfo) GetTargetSites() []string`

GetTargetSites returns the TargetSites field if non-nil, zero value otherwise.

### GetTargetSitesOk

`func (o *EditUploadFirmwareInfo) GetTargetSitesOk() (*[]string, bool)`

GetTargetSitesOk returns a tuple with the TargetSites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetSites

`func (o *EditUploadFirmwareInfo) SetTargetSites(v []string)`

SetTargetSites sets TargetSites field to given value.

### HasTargetSites

`func (o *EditUploadFirmwareInfo) HasTargetSites() bool`

HasTargetSites returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


