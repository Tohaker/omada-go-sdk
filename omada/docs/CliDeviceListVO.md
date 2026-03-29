# CliDeviceListVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceSeriesType** | Pointer to **int32** | DeviceSeriesType should be a value as follows: 0:advanced; 1:pro | [optional] 
**Ip** | Pointer to **string** | Ip address | [optional] 
**Mac** | Pointer to **string** | Device mac | [optional] 
**MasterMac** | Pointer to **string** | Master mac. If the device is not a member of stack, masterMac should be null | [optional] 
**Model** | Pointer to **string** | Model of device | [optional] 
**ModelVersion** | Pointer to **string** | Model version of device | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**StackId** | Pointer to **string** | The ID of the associated stack. If the device is not a member of stack, stackId should be null | [optional] 
**StackName** | Pointer to **string** | The name of the associated stack. If the device is not a member of stack, stackName should be null | [optional] 
**Status** | Pointer to **int32** | Status of device.Status should be a value as follows: 0:Disconnected; 1:Disconnected(Migrating); 10:Provisioning; 11:Configuring; 12:Upgrading; 13:Rebooting; 14:Connected; 15:Connected(Wireless); 16:Connected(Migrating); 17:Connected(Wireless,Migrating); 20:Pending; 21:Pending(Wireless); 22:Adopting; 23:Adopting(Wireless); 24:Adopt Failed; 25:Adopt Failed(Wireless); 26:Managed By Others; 27:Managed By Others(Wireless); 30:Heartbeat Missed; 31:Heartbeat Missed(Wireless); 32:Heartbeat Missed(Migrating); 33:Heartbeat Missed(Wireless,Migrating); 40:Isolated; 41:Isolated(Migrating); 50:Slice Configuring | [optional] 
**SupportCli** | Pointer to **bool** | Whether the device supports Cli | [optional] 
**Type** | Pointer to **string** | Type of device should be a value as follows: switch | [optional] 

## Methods

### NewCliDeviceListVO

`func NewCliDeviceListVO() *CliDeviceListVO`

NewCliDeviceListVO instantiates a new CliDeviceListVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCliDeviceListVOWithDefaults

`func NewCliDeviceListVOWithDefaults() *CliDeviceListVO`

NewCliDeviceListVOWithDefaults instantiates a new CliDeviceListVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceSeriesType

`func (o *CliDeviceListVO) GetDeviceSeriesType() int32`

GetDeviceSeriesType returns the DeviceSeriesType field if non-nil, zero value otherwise.

### GetDeviceSeriesTypeOk

`func (o *CliDeviceListVO) GetDeviceSeriesTypeOk() (*int32, bool)`

GetDeviceSeriesTypeOk returns a tuple with the DeviceSeriesType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSeriesType

`func (o *CliDeviceListVO) SetDeviceSeriesType(v int32)`

SetDeviceSeriesType sets DeviceSeriesType field to given value.

### HasDeviceSeriesType

`func (o *CliDeviceListVO) HasDeviceSeriesType() bool`

HasDeviceSeriesType returns a boolean if a field has been set.

### GetIp

`func (o *CliDeviceListVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *CliDeviceListVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *CliDeviceListVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *CliDeviceListVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *CliDeviceListVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *CliDeviceListVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *CliDeviceListVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *CliDeviceListVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetMasterMac

`func (o *CliDeviceListVO) GetMasterMac() string`

GetMasterMac returns the MasterMac field if non-nil, zero value otherwise.

### GetMasterMacOk

`func (o *CliDeviceListVO) GetMasterMacOk() (*string, bool)`

GetMasterMacOk returns a tuple with the MasterMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterMac

`func (o *CliDeviceListVO) SetMasterMac(v string)`

SetMasterMac sets MasterMac field to given value.

### HasMasterMac

`func (o *CliDeviceListVO) HasMasterMac() bool`

HasMasterMac returns a boolean if a field has been set.

### GetModel

`func (o *CliDeviceListVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CliDeviceListVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CliDeviceListVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CliDeviceListVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *CliDeviceListVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *CliDeviceListVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *CliDeviceListVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *CliDeviceListVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *CliDeviceListVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CliDeviceListVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CliDeviceListVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CliDeviceListVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStackId

`func (o *CliDeviceListVO) GetStackId() string`

GetStackId returns the StackId field if non-nil, zero value otherwise.

### GetStackIdOk

`func (o *CliDeviceListVO) GetStackIdOk() (*string, bool)`

GetStackIdOk returns a tuple with the StackId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackId

`func (o *CliDeviceListVO) SetStackId(v string)`

SetStackId sets StackId field to given value.

### HasStackId

`func (o *CliDeviceListVO) HasStackId() bool`

HasStackId returns a boolean if a field has been set.

### GetStackName

`func (o *CliDeviceListVO) GetStackName() string`

GetStackName returns the StackName field if non-nil, zero value otherwise.

### GetStackNameOk

`func (o *CliDeviceListVO) GetStackNameOk() (*string, bool)`

GetStackNameOk returns a tuple with the StackName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackName

`func (o *CliDeviceListVO) SetStackName(v string)`

SetStackName sets StackName field to given value.

### HasStackName

`func (o *CliDeviceListVO) HasStackName() bool`

HasStackName returns a boolean if a field has been set.

### GetStatus

`func (o *CliDeviceListVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CliDeviceListVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CliDeviceListVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CliDeviceListVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportCli

`func (o *CliDeviceListVO) GetSupportCli() bool`

GetSupportCli returns the SupportCli field if non-nil, zero value otherwise.

### GetSupportCliOk

`func (o *CliDeviceListVO) GetSupportCliOk() (*bool, bool)`

GetSupportCliOk returns a tuple with the SupportCli field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportCli

`func (o *CliDeviceListVO) SetSupportCli(v bool)`

SetSupportCli sets SupportCli field to given value.

### HasSupportCli

`func (o *CliDeviceListVO) HasSupportCli() bool`

HasSupportCli returns a boolean if a field has been set.

### GetType

`func (o *CliDeviceListVO) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CliDeviceListVO) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CliDeviceListVO) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CliDeviceListVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


