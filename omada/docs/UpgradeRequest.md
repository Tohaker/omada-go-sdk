# UpgradeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FileId** | Pointer to **string** | Firmware ID for upgrade | [optional] 
**Macs** | Pointer to **[]string** | MAC list of devices | [optional] 

## Methods

### NewUpgradeRequest

`func NewUpgradeRequest() *UpgradeRequest`

NewUpgradeRequest instantiates a new UpgradeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpgradeRequestWithDefaults

`func NewUpgradeRequestWithDefaults() *UpgradeRequest`

NewUpgradeRequestWithDefaults instantiates a new UpgradeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFileId

`func (o *UpgradeRequest) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *UpgradeRequest) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *UpgradeRequest) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *UpgradeRequest) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetMacs

`func (o *UpgradeRequest) GetMacs() []string`

GetMacs returns the Macs field if non-nil, zero value otherwise.

### GetMacsOk

`func (o *UpgradeRequest) GetMacsOk() (*[]string, bool)`

GetMacsOk returns a tuple with the Macs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacs

`func (o *UpgradeRequest) SetMacs(v []string)`

SetMacs sets Macs field to given value.

### HasMacs

`func (o *UpgradeRequest) HasMacs() bool`

HasMacs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


