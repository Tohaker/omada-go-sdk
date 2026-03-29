# OswStpRegionConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**RebootDelay** | Pointer to **int64** |  | [optional] 
**Revision** | **int32** | Revision | 

## Methods

### NewOswStpRegionConfig

`func NewOswStpRegionConfig(revision int32, ) *OswStpRegionConfig`

NewOswStpRegionConfig instantiates a new OswStpRegionConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStpRegionConfigWithDefaults

`func NewOswStpRegionConfigWithDefaults() *OswStpRegionConfig`

NewOswStpRegionConfigWithDefaults instantiates a new OswStpRegionConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *OswStpRegionConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswStpRegionConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswStpRegionConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswStpRegionConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRebootDelay

`func (o *OswStpRegionConfig) GetRebootDelay() int64`

GetRebootDelay returns the RebootDelay field if non-nil, zero value otherwise.

### GetRebootDelayOk

`func (o *OswStpRegionConfig) GetRebootDelayOk() (*int64, bool)`

GetRebootDelayOk returns a tuple with the RebootDelay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebootDelay

`func (o *OswStpRegionConfig) SetRebootDelay(v int64)`

SetRebootDelay sets RebootDelay field to given value.

### HasRebootDelay

`func (o *OswStpRegionConfig) HasRebootDelay() bool`

HasRebootDelay returns a boolean if a field has been set.

### GetRevision

`func (o *OswStpRegionConfig) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *OswStpRegionConfig) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *OswStpRegionConfig) SetRevision(v int32)`

SetRevision sets Revision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


