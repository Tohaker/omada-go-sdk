# OswSdmBriefVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]OswSdmApplicationVO**](OswSdmApplicationVO.md) | Application  | [optional] 
**Name** | Pointer to **string** | Sdm template name. | [optional] 

## Methods

### NewOswSdmBriefVO

`func NewOswSdmBriefVO() *OswSdmBriefVO`

NewOswSdmBriefVO instantiates a new OswSdmBriefVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswSdmBriefVOWithDefaults

`func NewOswSdmBriefVOWithDefaults() *OswSdmBriefVO`

NewOswSdmBriefVOWithDefaults instantiates a new OswSdmBriefVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApps

`func (o *OswSdmBriefVO) GetApps() []OswSdmApplicationVO`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *OswSdmBriefVO) GetAppsOk() (*[]OswSdmApplicationVO, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *OswSdmBriefVO) SetApps(v []OswSdmApplicationVO)`

SetApps sets Apps field to given value.

### HasApps

`func (o *OswSdmBriefVO) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetName

`func (o *OswSdmBriefVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswSdmBriefVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswSdmBriefVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswSdmBriefVO) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


