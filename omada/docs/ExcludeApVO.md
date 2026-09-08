# ExcludeApVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedInAdvanced** | Pointer to **bool** | Whether the device is added in advanced. | [optional] 
**Ip** | Pointer to **string** | IP adress. | [optional] 
**Mac** | Pointer to **string** | Device MAC. | [optional] 
**Model** | Pointer to **string** | Device model. | [optional] 
**ModelVersion** | Pointer to **string** | Model version. | [optional] 
**Name** | Pointer to **string** | Device name. | [optional] 
**ShowModel** | Pointer to **string** | The model name of device. | [optional] 
**Status** | Pointer to **int32** | Detailed status.  | [optional] 
**StatusCategory** | Pointer to **int32** | Status category. 0 : Disconnected. 1: Connected. 2: Pending. 3: Heartbeat Missed. 4: Isolated. | [optional] 
**Type** | Pointer to **string** | Device type. | [optional] 
**UnsupportType** | Pointer to **int32** | 0: Not support mesh(cannot be excluded). 1: Not support channel deployment. 2: Not support power deployment. | [optional] 

## Methods

### NewExcludeApVO

`func NewExcludeApVO() *ExcludeApVO`

NewExcludeApVO instantiates a new ExcludeApVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExcludeApVOWithDefaults

`func NewExcludeApVOWithDefaults() *ExcludeApVO`

NewExcludeApVOWithDefaults instantiates a new ExcludeApVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedInAdvanced

`func (o *ExcludeApVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *ExcludeApVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *ExcludeApVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *ExcludeApVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetIp

`func (o *ExcludeApVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *ExcludeApVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *ExcludeApVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *ExcludeApVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *ExcludeApVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *ExcludeApVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *ExcludeApVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *ExcludeApVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *ExcludeApVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ExcludeApVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ExcludeApVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ExcludeApVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *ExcludeApVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *ExcludeApVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *ExcludeApVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *ExcludeApVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *ExcludeApVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExcludeApVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExcludeApVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ExcludeApVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetShowModel

`func (o *ExcludeApVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *ExcludeApVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *ExcludeApVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *ExcludeApVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetStatus

`func (o *ExcludeApVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExcludeApVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExcludeApVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExcludeApVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *ExcludeApVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *ExcludeApVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *ExcludeApVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *ExcludeApVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetType

`func (o *ExcludeApVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExcludeApVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExcludeApVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExcludeApVO) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnsupportType

`func (o *ExcludeApVO) GetUnsupportType() int32`

GetUnsupportType returns the UnsupportType field if non-nil, zero value otherwise.

### GetUnsupportTypeOk

`func (o *ExcludeApVO) GetUnsupportTypeOk() (*int32, bool)`

GetUnsupportTypeOk returns a tuple with the UnsupportType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsupportType

`func (o *ExcludeApVO) SetUnsupportType(v int32)`

SetUnsupportType sets UnsupportType field to given value.

### HasUnsupportType

`func (o *ExcludeApVO) HasUnsupportType() bool`

HasUnsupportType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


