# SsidDeviceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApGroupId** | Pointer to **string** | apGroupId | [optional] 
**ApGroupName** | Pointer to **string** | AP Group Name | [optional] 
**ClientNum** | Pointer to **int64** | number of clients connected to Device | [optional] 
**DeviceType** | Pointer to **string** | Device type, such as EAP, Gateway | [optional] 
**Id** | Pointer to **string** | Device ID | [optional] 
**Ip** | Pointer to **string** | Ip address,such as 192.168.0.105 | [optional] 
**Mac** | Pointer to **string** | Device MAC address | [optional] 
**Model** | Pointer to **string** | Device model | [optional] 
**ModelVersion** | Pointer to **string** | Device model version | [optional] 
**Name** | Pointer to **string** | Device name | [optional] 
**OverrideNum** | Pointer to **int32** | Override ssid number for AP | [optional] 
**ShowModel** | Pointer to **string** | Model complex shown in the front end.Ap：model+(country)+modelVersion,EAP225(EU) v3.0  Gateway/Switch：model+modelVersion,Osg v3.0 | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected;1:Disconnected(Migrating);10:Provisioning;11:Configuring;12:Upgrading;13:Rebooting;14:Connected;15:Connected(Wireless);16:Connected(Migrating);17:Connected(Wireless,Migrating);20:Pending;21:Pending(Wireless);22:Adopting;23:Adopting(Wireless);24:Adopt Failed;25:Adopt Failed(Wireless);26:Managed By Others;27:Managed By Others(Wireless);30:Heartbeat Missed;31:Heartbeat Missed(Wireless);32:Heartbeat Missed(Migrating);33:Heartbeat Missed(Wireless,Migrating);40:Isolated;41:Isolated(Migrating);50:Slice Configuring | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected;1:Connected;2:Pending;3:Heartbeat Missed;4:Isolated | [optional] 
**Traffic** | Pointer to **int64** | total traffic of Device | [optional] 

## Methods

### NewSsidDeviceOpenApiVO

`func NewSsidDeviceOpenApiVO() *SsidDeviceOpenApiVO`

NewSsidDeviceOpenApiVO instantiates a new SsidDeviceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsidDeviceOpenApiVOWithDefaults

`func NewSsidDeviceOpenApiVOWithDefaults() *SsidDeviceOpenApiVO`

NewSsidDeviceOpenApiVOWithDefaults instantiates a new SsidDeviceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApGroupId

`func (o *SsidDeviceOpenApiVO) GetApGroupId() string`

GetApGroupId returns the ApGroupId field if non-nil, zero value otherwise.

### GetApGroupIdOk

`func (o *SsidDeviceOpenApiVO) GetApGroupIdOk() (*string, bool)`

GetApGroupIdOk returns a tuple with the ApGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApGroupId

`func (o *SsidDeviceOpenApiVO) SetApGroupId(v string)`

SetApGroupId sets ApGroupId field to given value.

### HasApGroupId

`func (o *SsidDeviceOpenApiVO) HasApGroupId() bool`

HasApGroupId returns a boolean if a field has been set.

### GetApGroupName

`func (o *SsidDeviceOpenApiVO) GetApGroupName() string`

GetApGroupName returns the ApGroupName field if non-nil, zero value otherwise.

### GetApGroupNameOk

`func (o *SsidDeviceOpenApiVO) GetApGroupNameOk() (*string, bool)`

GetApGroupNameOk returns a tuple with the ApGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApGroupName

`func (o *SsidDeviceOpenApiVO) SetApGroupName(v string)`

SetApGroupName sets ApGroupName field to given value.

### HasApGroupName

`func (o *SsidDeviceOpenApiVO) HasApGroupName() bool`

HasApGroupName returns a boolean if a field has been set.

### GetClientNum

`func (o *SsidDeviceOpenApiVO) GetClientNum() int64`

GetClientNum returns the ClientNum field if non-nil, zero value otherwise.

### GetClientNumOk

`func (o *SsidDeviceOpenApiVO) GetClientNumOk() (*int64, bool)`

GetClientNumOk returns a tuple with the ClientNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientNum

`func (o *SsidDeviceOpenApiVO) SetClientNum(v int64)`

SetClientNum sets ClientNum field to given value.

### HasClientNum

`func (o *SsidDeviceOpenApiVO) HasClientNum() bool`

HasClientNum returns a boolean if a field has been set.

### GetDeviceType

`func (o *SsidDeviceOpenApiVO) GetDeviceType() string`

GetDeviceType returns the DeviceType field if non-nil, zero value otherwise.

### GetDeviceTypeOk

`func (o *SsidDeviceOpenApiVO) GetDeviceTypeOk() (*string, bool)`

GetDeviceTypeOk returns a tuple with the DeviceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceType

`func (o *SsidDeviceOpenApiVO) SetDeviceType(v string)`

SetDeviceType sets DeviceType field to given value.

### HasDeviceType

`func (o *SsidDeviceOpenApiVO) HasDeviceType() bool`

HasDeviceType returns a boolean if a field has been set.

### GetId

`func (o *SsidDeviceOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SsidDeviceOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SsidDeviceOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SsidDeviceOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIp

`func (o *SsidDeviceOpenApiVO) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SsidDeviceOpenApiVO) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SsidDeviceOpenApiVO) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SsidDeviceOpenApiVO) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetMac

`func (o *SsidDeviceOpenApiVO) GetMac() string`

GetMac returns the Mac field if non-nil, zero value otherwise.

### GetMacOk

`func (o *SsidDeviceOpenApiVO) GetMacOk() (*string, bool)`

GetMacOk returns a tuple with the Mac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMac

`func (o *SsidDeviceOpenApiVO) SetMac(v string)`

SetMac sets Mac field to given value.

### HasMac

`func (o *SsidDeviceOpenApiVO) HasMac() bool`

HasMac returns a boolean if a field has been set.

### GetModel

`func (o *SsidDeviceOpenApiVO) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SsidDeviceOpenApiVO) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SsidDeviceOpenApiVO) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SsidDeviceOpenApiVO) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetModelVersion

`func (o *SsidDeviceOpenApiVO) GetModelVersion() string`

GetModelVersion returns the ModelVersion field if non-nil, zero value otherwise.

### GetModelVersionOk

`func (o *SsidDeviceOpenApiVO) GetModelVersionOk() (*string, bool)`

GetModelVersionOk returns a tuple with the ModelVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelVersion

`func (o *SsidDeviceOpenApiVO) SetModelVersion(v string)`

SetModelVersion sets ModelVersion field to given value.

### HasModelVersion

`func (o *SsidDeviceOpenApiVO) HasModelVersion() bool`

HasModelVersion returns a boolean if a field has been set.

### GetName

`func (o *SsidDeviceOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SsidDeviceOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SsidDeviceOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SsidDeviceOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOverrideNum

`func (o *SsidDeviceOpenApiVO) GetOverrideNum() int32`

GetOverrideNum returns the OverrideNum field if non-nil, zero value otherwise.

### GetOverrideNumOk

`func (o *SsidDeviceOpenApiVO) GetOverrideNumOk() (*int32, bool)`

GetOverrideNumOk returns a tuple with the OverrideNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideNum

`func (o *SsidDeviceOpenApiVO) SetOverrideNum(v int32)`

SetOverrideNum sets OverrideNum field to given value.

### HasOverrideNum

`func (o *SsidDeviceOpenApiVO) HasOverrideNum() bool`

HasOverrideNum returns a boolean if a field has been set.

### GetShowModel

`func (o *SsidDeviceOpenApiVO) GetShowModel() string`

GetShowModel returns the ShowModel field if non-nil, zero value otherwise.

### GetShowModelOk

`func (o *SsidDeviceOpenApiVO) GetShowModelOk() (*string, bool)`

GetShowModelOk returns a tuple with the ShowModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowModel

`func (o *SsidDeviceOpenApiVO) SetShowModel(v string)`

SetShowModel sets ShowModel field to given value.

### HasShowModel

`func (o *SsidDeviceOpenApiVO) HasShowModel() bool`

HasShowModel returns a boolean if a field has been set.

### GetStatus

`func (o *SsidDeviceOpenApiVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SsidDeviceOpenApiVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SsidDeviceOpenApiVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SsidDeviceOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *SsidDeviceOpenApiVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *SsidDeviceOpenApiVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *SsidDeviceOpenApiVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *SsidDeviceOpenApiVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetTraffic

`func (o *SsidDeviceOpenApiVO) GetTraffic() int64`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *SsidDeviceOpenApiVO) GetTrafficOk() (*int64, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *SsidDeviceOpenApiVO) SetTraffic(v int64)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *SsidDeviceOpenApiVO) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


