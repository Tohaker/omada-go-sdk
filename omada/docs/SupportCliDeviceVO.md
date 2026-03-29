# SupportCliDeviceVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedInAdvanced** | Pointer to **bool** | Whether the device is added offline in advance. | [optional] 
**DeviceSeriesType** | Pointer to **int32** | DeviceSeriesType should be a value as follows: 0:advanced; 1:pro | [optional] 
**Ip** | Pointer to **string** | Ip address | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**MasterMac** | Pointer to **string** | Master mac. If the device is not a member of stack, masterMac should be null | [optional] 
**Model** | Pointer to **string** | Model of device | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**StackId** | Pointer to **string** | The ID of the associated stack. If the device is not a member of stack, stackId should be null | [optional] 
**StackName** | Pointer to **string** | The name of the associated stack. If the device is not a member of stack, stackName should be null | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected; 1:Disconnected(Migrating); 10:Provisioning; 11:Configuring; 12:Upgrading; 13:Rebooting; 14:Connected; 15:Connected(Wireless); 16:Connected(Migrating); 17:Connected(Wireless,Migrating); 20:Pending; 21:Pending(Wireless); 22:Adopting; 23:Adopting(Wireless); 24:Adopt Failed; 25:Adopt Failed(Wireless); 26:Managed By Others; 27:Managed By Others(Wireless); 30:Heartbeat Missed; 31:Heartbeat Missed(Wireless); 32:Heartbeat Missed(Migrating); 33:Heartbeat Missed(Wireless,Migrating); 40:Isolated; 41:Isolated(Migrating); 50:Slice Configuring | [optional] 
**SupportCli** | Pointer to **bool** | Whether the device supports Cli | [optional] 
**Type** | Pointer to **string** | Type of device should be a value as follows: switch | [optional] 

## Methods

### NewSupportCliDeviceVO

`func NewSupportCliDeviceVO() *SupportCliDeviceVO`

NewSupportCliDeviceVO instantiates a new SupportCliDeviceVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportCliDeviceVOWithDefaults

`func NewSupportCliDeviceVOWithDefaults() *SupportCliDeviceVO`

NewSupportCliDeviceVOWithDefaults instantiates a new SupportCliDeviceVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedInAdvanced

`func (o *SupportCliDeviceVO) GetAddedInAdvanced() bool`

GetAddedInAdvanced returns the AddedInAdvanced field if non-nil, zero value otherwise.

### GetAddedInAdvancedOk

`func (o *SupportCliDeviceVO) GetAddedInAdvancedOk() (*bool, bool)`

GetAddedInAdvancedOk returns a tuple with the AddedInAdvanced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedInAdvanced

`func (o *SupportCliDeviceVO) SetAddedInAdvanced(v bool)`

SetAddedInAdvanced sets AddedInAdvanced field to given value.

### HasAddedInAdvanced

`func (o *SupportCliDeviceVO) HasAddedInAdvanced() bool`

HasAddedInAdvanced returns a boolean if a field has been set.

### GetDeviceSeriesType

`func (o *SupportCliDeviceVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *SupportCliDeviceVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *SupportCliDeviceVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *SupportCliDeviceVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetIp

`func (o *SupportCliDeviceVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SupportCliDeviceVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SupportCliDeviceVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SupportCliDeviceVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *SupportCliDeviceVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SupportCliDeviceVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SupportCliDeviceVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SupportCliDeviceVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMasterMac

`func (o *SupportCliDeviceVO) GetMasterMac() string`

GetMasterMac returns the MasterMac field if non-nil, zero value otherwise.

### GetMasterMacOk

`func (o *SupportCliDeviceVO) GetMasterMacOk() (*string, bool)`

GetMasterMacOk returns a tuple with the MasterMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterMac

`func (o *SupportCliDeviceVO) SetMasterMac(v string)`

SetMasterMac sets MasterMac field to given value.

### HasMasterMac

`func (o *SupportCliDeviceVO) HasMasterMac() bool`

HasMasterMac returns a boolean if a field has been set.

### GetModel

`func (o *SupportCliDeviceVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SupportCliDeviceVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SupportCliDeviceVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SupportCliDeviceVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *SupportCliDeviceVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *SupportCliDeviceVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *SupportCliDeviceVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *SupportCliDeviceVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *SupportCliDeviceVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SupportCliDeviceVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SupportCliDeviceVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SupportCliDeviceVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStackId

`func (o *SupportCliDeviceVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *SupportCliDeviceVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *SupportCliDeviceVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *SupportCliDeviceVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *SupportCliDeviceVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *SupportCliDeviceVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *SupportCliDeviceVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *SupportCliDeviceVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStatus

`func (o *SupportCliDeviceVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SupportCliDeviceVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SupportCliDeviceVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SupportCliDeviceVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportCli

`func (o *SupportCliDeviceVO) GetSupportCli() bool`

GetSupportCli returns the SupportCli field if non-nil, zero value otherwise.

### GetSupportCliOk

`func (o *SupportCliDeviceVO) GetSupportCliOk() (*bool, bool)`

GetSupportCliOk returns a tuple with the SupportCli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCli

`func (o *SupportCliDeviceVO) SetSupportCli(v bool)`

SetSupportCli sets SupportCli field to given value.

### HasSupportCli

`func (o *SupportCliDeviceVO) HasSupportCli() bool`

HasSupportCli returns a boolean if a field has been set.

### GetType

`func (o *SupportCliDeviceVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SupportCliDeviceVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SupportCliDeviceVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SupportCliDeviceVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


