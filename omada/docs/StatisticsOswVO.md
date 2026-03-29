# StatisticsOswVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MlagMsg** | Pointer to [**MlagMsgVO**](MlagMsgVO.md) |  | [optional] 
**Name** | Pointer to **string** | Device name,default value is the mac address of device | [optional] 
**PortNum** | Pointer to **int32** | Port number | [optional] 
**Ports** | Pointer to [**[]OswStatPortVO**](OswStatPortVO.md) | Ports information | [optional] 
**Speeds** | Pointer to **[]int32** | Supported rate list for all ports. Speeds should be a value as follows: 0:auto; 1:10M; 2:100M; 3:1000M; 4:2.5G; 5:10G; 6:5G; 7:25G; 8:100G; 9:40G; -1:error; no value:all rate supported | [optional] 
**Status** | Pointer to **int32** | Status of device,status should be a value as follows: 0:Disconnected; 1:Disconnected(Migrating); 10:Provisioning; 11:Configuring; 12:Upgrading; 13:Rebooting; 14:Connected; 15:Connected(Wireless); 16:Connected(Migrating); 17:Connected(Wireless,Migrating); 20:Pending; 21:Pending(Wireless); 22:Adopting; 23:Adopting(Wireless); 24:Adopt Failed; 25:Adopt Failed(Wireless); 26:Managed By Others; 27:Managed By Others(Wireless); 30:Heartbeat Missed; 31:Heartbeat Missed(Wireless); 32:Heartbeat Missed(Migrating); 33:Heartbeat Missed(Wireless,Migrating); 40:Isolated; 41:Isolated(Migrating); 50:Slice Configuring | [optional] 
**StatusCategory** | Pointer to **int32** | Category of device status,statusCategory should be a value as follows: 0:Disconnected; 1:Connected; 2:Pending; 3:Heartbeat Missed; 4:Isolated | [optional] 
**SupportPoe** | Pointer to **bool** | Whether the device supports POE function | [optional] 
**SupportSTP** | Pointer to **bool** | Whether device supports stp function | [optional] 
**SupportStack** | Pointer to **bool** | Whether device supports stack function | [optional] 
**Uplink** | Pointer to [**OswStatUplinkVO**](OswStatUplinkVO.md) |  | [optional] 

## Methods

### NewStatisticsOswVO

`func NewStatisticsOswVO() *StatisticsOswVO`

NewStatisticsOswVO instantiates a new StatisticsOswVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatisticsOswVOWithDefaults

`func NewStatisticsOswVOWithDefaults() *StatisticsOswVO`

NewStatisticsOswVOWithDefaults instantiates a new StatisticsOswVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMlagMsg

`func (o *StatisticsOswVO) GetMlagMsg() MlagMsgVO`

GetMlagMsg returns the MlagMsg field if non-nil, zero value otherwise.

### GetMlagMsgOk

`func (o *StatisticsOswVO) GetMlagMsgOk() (*MlagMsgVO, bool)`

GetMlagMsgOk returns a tuple with the MlagMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMlagMsg

`func (o *StatisticsOswVO) SetMlagMsg(v MlagMsgVO)`

SetMlagMsg sets MlagMsg field to given value.

### HasMlagMsg

`func (o *StatisticsOswVO) HasMlagMsg() bool`

HasMlagMsg returns a boolean if a field has been set.

### GetName

`func (o *StatisticsOswVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StatisticsOswVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StatisticsOswVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *StatisticsOswVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPortNum

`func (o *StatisticsOswVO) GetPortNum() int32`

GetPortNum returns the PortNum field if non-nil, zero value otherwise.

### GetPortNumOk

`func (o *StatisticsOswVO) GetPortNumOk() (*int32, bool)`

GetPortNumOk returns a tuple with the PortNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortNum

`func (o *StatisticsOswVO) SetPortNum(v int32)`

SetPortNum sets PortNum field to given value.

### HasPortNum

`func (o *StatisticsOswVO) HasPortNum() bool`

HasPortNum returns a boolean if a field has been set.

### GetPorts

`func (o *StatisticsOswVO) GetPorts() []OswStatPortVO`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *StatisticsOswVO) GetPortsOk() (*[]OswStatPortVO, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *StatisticsOswVO) SetPorts(v []OswStatPortVO)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *StatisticsOswVO) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetSpeeds

`func (o *StatisticsOswVO) GetSpeeds() []int32`

GetSpeeds returns the Speeds field if non-nil, zero value otherwise.

### GetSpeedsOk

`func (o *StatisticsOswVO) GetSpeedsOk() (*[]int32, bool)`

GetSpeedsOk returns a tuple with the Speeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeeds

`func (o *StatisticsOswVO) SetSpeeds(v []int32)`

SetSpeeds sets Speeds field to given value.

### HasSpeeds

`func (o *StatisticsOswVO) HasSpeeds() bool`

HasSpeeds returns a boolean if a field has been set.

### GetStatus

`func (o *StatisticsOswVO) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *StatisticsOswVO) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *StatisticsOswVO) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *StatisticsOswVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusCategory

`func (o *StatisticsOswVO) GetStatusCategory() int32`

GetStatusCategory returns the StatusCategory field if non-nil, zero value otherwise.

### GetStatusCategoryOk

`func (o *StatisticsOswVO) GetStatusCategoryOk() (*int32, bool)`

GetStatusCategoryOk returns a tuple with the StatusCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCategory

`func (o *StatisticsOswVO) SetStatusCategory(v int32)`

SetStatusCategory sets StatusCategory field to given value.

### HasStatusCategory

`func (o *StatisticsOswVO) HasStatusCategory() bool`

HasStatusCategory returns a boolean if a field has been set.

### GetSupportPoe

`func (o *StatisticsOswVO) GetSupportPoe() bool`

GetSupportPoe returns the SupportPoe field if non-nil, zero value otherwise.

### GetSupportPoeOk

`func (o *StatisticsOswVO) GetSupportPoeOk() (*bool, bool)`

GetSupportPoeOk returns a tuple with the SupportPoe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportPoe

`func (o *StatisticsOswVO) SetSupportPoe(v bool)`

SetSupportPoe sets SupportPoe field to given value.

### HasSupportPoe

`func (o *StatisticsOswVO) HasSupportPoe() bool`

HasSupportPoe returns a boolean if a field has been set.

### GetSupportSTP

`func (o *StatisticsOswVO) GetSupportSTP() bool`

GetSupportSTP returns the SupportSTP field if non-nil, zero value otherwise.

### GetSupportSTPOk

`func (o *StatisticsOswVO) GetSupportSTPOk() (*bool, bool)`

GetSupportSTPOk returns a tuple with the SupportSTP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportSTP

`func (o *StatisticsOswVO) SetSupportSTP(v bool)`

SetSupportSTP sets SupportSTP field to given value.

### HasSupportSTP

`func (o *StatisticsOswVO) HasSupportSTP() bool`

HasSupportSTP returns a boolean if a field has been set.

### GetSupportStack

`func (o *StatisticsOswVO) GetSupportStack() bool`

GetSupportStack returns the SupportStack field if non-nil, zero value otherwise.

### GetSupportStackOk

`func (o *StatisticsOswVO) GetSupportStackOk() (*bool, bool)`

GetSupportStackOk returns a tuple with the SupportStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportStack

`func (o *StatisticsOswVO) SetSupportStack(v bool)`

SetSupportStack sets SupportStack field to given value.

### HasSupportStack

`func (o *StatisticsOswVO) HasSupportStack() bool`

HasSupportStack returns a boolean if a field has been set.

### GetUplink

`func (o *StatisticsOswVO) GetUplink() OswStatUplinkVO`

GetUplink returns the Uplink field if non-nil, zero value otherwise.

### GetUplinkOk

`func (o *StatisticsOswVO) GetUplinkOk() (*OswStatUplinkVO, bool)`

GetUplinkOk returns a tuple with the Uplink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUplink

`func (o *StatisticsOswVO) SetUplink(v OswStatUplinkVO)`

SetUplink sets Uplink field to given value.

### HasUplink

`func (o *StatisticsOswVO) HasUplink() bool`

HasUplink returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


