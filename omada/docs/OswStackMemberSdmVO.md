# OswStackMemberSdmVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceMac** | Pointer to **string** | Mac of device. | [optional] 
**IsStackMaster** | Pointer to **bool** | Whether the device is the master unit in a stack. | [optional] 
**Model** | Pointer to **string** | Device model. | [optional] 
**ModelVersion** | Pointer to **string** | Device model Version. | [optional] 
**SdmResourceUsageList** | Pointer to [**[]SdmResourceUsage**](SdmResourceUsage.md) | The list contains usage details of SDM resources on the device, including used and available resources. | [optional] 
**Status** | Pointer to **int32** | Device status. | [optional] 

## Methods

### NewOswStackMemberSdmVO

`func NewOswStackMemberSdmVO() *OswStackMemberSdmVO`

NewOswStackMemberSdmVO instantiates a new OswStackMemberSdmVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackMemberSdmVOWithDefaults

`func NewOswStackMemberSdmVOWithDefaults() *OswStackMemberSdmVO`

NewOswStackMemberSdmVOWithDefaults instantiates a new OswStackMemberSdmVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceMac

`func (o *OswStackMemberSdmVO) GetDeviceMac() string`

GetDeviceMac returns the DeviceMac field if non-nil, zero value otherwise.

### GetDeviceMacOk

`func (o *OswStackMemberSdmVO) GetDeviceMacOk() (*string, bool)`

GetDeviceMacOk returns a tuple with the DeviceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceMac

`func (o *OswStackMemberSdmVO) SetDeviceMac(v string)`

SetDeviceMac sets DeviceMac field to given value.

### HasDeviceMac

`func (o *OswStackMemberSdmVO) HasDeviceMac() bool`

HasDeviceMac returns a boolean if a field has been set.

### GetIsStackMaster

`func (o *OswStackMemberSdmVO) GetIsStackMaster() bool`

GetIsStackMaster returns the IsStackMaster field if non-nil, zero value otherwise.

### GetIsStackMasterOk

`func (o *OswStackMemberSdmVO) GetIsStackMasterOk() (*bool, bool)`

GetIsStackMasterOk returns a tuple with the IsStackMaster field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStackMaster

`func (o *OswStackMemberSdmVO) SetIsStackMaster(v bool)`

SetIsStackMaster sets IsStackMaster field to given value.

### HasIsStackMaster

`func (o *OswStackMemberSdmVO) HasIsStackMaster() bool`

HasIsStackMaster returns a boolean if a field has been set.

### GetModel

`func (o *OswStackMemberSdmVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OswStackMemberSdmVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OswStackMemberSdmVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OswStackMemberSdmVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *OswStackMemberSdmVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *OswStackMemberSdmVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *OswStackMemberSdmVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *OswStackMemberSdmVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetSdmResourceUsageList

`func (o *OswStackMemberSdmVO) GetSdmResourceUsageList() []SdmResourceUsage`

GetSdmResourceUsageList returns the SdmResourceUsageList field if non-nil, zero value otherwise.

### GetSdmResourceUsageListOk

`func (o *OswStackMemberSdmVO) GetSdmResourceUsageListOk() (*[]SdmResourceUsage, bool)`

GetSdmResourceUsageListOk returns a tuple with the SdmResourceUsageList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSdmResourceUsageList

`func (o *OswStackMemberSdmVO) SetSdmResourceUsageList(v []SdmResourceUsage)`

SetSdmResourceUsageList sets SdmResourceUsageList field to given value.

### HasSdmResourceUsageList

`func (o *OswStackMemberSdmVO) HasSdmResourceUsageList() bool`

HasSdmResourceUsageList returns a boolean if a field has been set.

### GetStatus

`func (o *OswStackMemberSdmVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OswStackMemberSdmVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OswStackMemberSdmVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OswStackMemberSdmVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


