# OswStatPortVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigMlagDad** | Pointer to **bool** |  | [optional] 
**ConfigMlagPeerLink** | Pointer to **bool** | Standard port information | [optional] 
**ConfigStack** | Pointer to **bool** | Indicates whether the current port is configured as a stack port (joined a stack aggregation group) | [optional] 
**Disable** | Pointer to **bool** | Whether the port is disabled | [optional] 
**Downlink** | Pointer to [**OswStatDownLinkVO**](OswStatDownLinkVO.md) |  | [optional] 
**MadUsed** | Pointer to **bool** |  | [optional] 
**MaxSpeed** | Pointer to **int32** | Max speed supported by the port | [optional] 
**Name** | Pointer to **string** | Port name | [optional] 
**Operation** | Pointer to **string** | Operation should be a value as follows: SWITCHING; MIRRORING; AGGREGATING | [optional] 
**OswStandPort** | Pointer to [**OswStandPortVO**](OswStandPortVO.md) |  | [optional] 
**Port** | Pointer to **int32** | Port id | [optional] 
**PortStatus** | Pointer to [**OswStatPortStatusVO**](OswStatPortStatusVO.md) |  | [optional] 
**StackPortsGroupIndex** | Pointer to **int32** | Number of the stacking port aggregation group to join | [optional] 
**Type** | Pointer to **int32** | Type should be a value as follows: 1:Copper; 2:Combo; 3:SFP | [optional] 

## Methods

### NewOswStatPortVO

`func NewOswStatPortVO() *OswStatPortVO`

NewOswStatPortVO instantiates a new OswStatPortVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStatPortVOWithDefaults

`func NewOswStatPortVOWithDefaults() *OswStatPortVO`

NewOswStatPortVOWithDefaults instantiates a new OswStatPortVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigMlagDad

`func (o *OswStatPortVO) GetConfigMlagDad() bool`

GetConfigMlagDad returns the ConfigMlagDad field if non-nil, zero value otherwise.

### GetConfigMlagDadOk

`func (o *OswStatPortVO) GetConfigMlagDadOk() (*bool, bool)`

GetConfigMlagDadOk returns a tuple with the ConfigMlagDad field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMlagDad

`func (o *OswStatPortVO) SetConfigMlagDad(v bool)`

SetConfigMlagDad sets ConfigMlagDad field to given value.

### HasConfigMlagDad

`func (o *OswStatPortVO) HasConfigMlagDad() bool`

HasConfigMlagDad returns a boolean if a field has been set.

### GetConfigMlagPeerLink

`func (o *OswStatPortVO) GetConfigMlagPeerLink() bool`

GetConfigMlagPeerLink returns the ConfigMlagPeerLink field if non-nil, zero value otherwise.

### GetConfigMlagPeerLinkOk

`func (o *OswStatPortVO) GetConfigMlagPeerLinkOk() (*bool, bool)`

GetConfigMlagPeerLinkOk returns a tuple with the ConfigMlagPeerLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigMlagPeerLink

`func (o *OswStatPortVO) SetConfigMlagPeerLink(v bool)`

SetConfigMlagPeerLink sets ConfigMlagPeerLink field to given value.

### HasConfigMlagPeerLink

`func (o *OswStatPortVO) HasConfigMlagPeerLink() bool`

HasConfigMlagPeerLink returns a boolean if a field has been set.

### GetConfigStack

`func (o *OswStatPortVO) GetConfigStack() bool`

GetConfigStack returns the ConfigStack field if non-nil, zero value otherwise.

### GetConfigStackOk

`func (o *OswStatPortVO) GetConfigStackOk() (*bool, bool)`

GetConfigStackOk returns a tuple with the ConfigStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigStack

`func (o *OswStatPortVO) SetConfigStack(v bool)`

SetConfigStack sets ConfigStack field to given value.

### HasConfigStack

`func (o *OswStatPortVO) HasConfigStack() bool`

HasConfigStack returns a boolean if a field has been set.

### GetDisable

`func (o *OswStatPortVO) GetDisable() bool`

GetDisable returns the Disable field if non-nil, zero value otherwise.

### GetDisableOk

`func (o *OswStatPortVO) GetDisableOk() (*bool, bool)`

GetDisableOk returns a tuple with the Disable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisable

`func (o *OswStatPortVO) SetDisable(v bool)`

SetDisable sets Disable field to given value.

### HasDisable

`func (o *OswStatPortVO) HasDisable() bool`

HasDisable returns a boolean if a field has been set.

### GetDownlink

`func (o *OswStatPortVO) GetDownlink() OswStatDownLinkVO`

GetDownlink returns the Downlink field if non-nil, zero value otherwise.

### GetDownlinkOk

`func (o *OswStatPortVO) GetDownlinkOk() (*OswStatDownLinkVO, bool)`

GetDownlinkOk returns a tuple with the Downlink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlink

`func (o *OswStatPortVO) SetDownlink(v OswStatDownLinkVO)`

SetDownlink sets Downlink field to given value.

### HasDownlink

`func (o *OswStatPortVO) HasDownlink() bool`

HasDownlink returns a boolean if a field has been set.

### GetMadUsed

`func (o *OswStatPortVO) GetMadUsed() bool`

GetMadUsed returns the MadUsed field if non-nil, zero value otherwise.

### GetMadUsedOk

`func (o *OswStatPortVO) GetMadUsedOk() (*bool, bool)`

GetMadUsedOk returns a tuple with the MadUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMadUsed

`func (o *OswStatPortVO) SetMadUsed(v bool)`

SetMadUsed sets MadUsed field to given value.

### HasMadUsed

`func (o *OswStatPortVO) HasMadUsed() bool`

HasMadUsed returns a boolean if a field has been set.

### GetMaxSpeed

`func (o *OswStatPortVO) GetMaxSpeed() int32`

GetMaxSpeed returns the MaxSpeed field if non-nil, zero value otherwise.

### GetMaxSpeedOk

`func (o *OswStatPortVO) GetMaxSpeedOk() (*int32, bool)`

GetMaxSpeedOk returns a tuple with the MaxSpeed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSpeed

`func (o *OswStatPortVO) SetMaxSpeed(v int32)`

SetMaxSpeed sets MaxSpeed field to given value.

### HasMaxSpeed

`func (o *OswStatPortVO) HasMaxSpeed() bool`

HasMaxSpeed returns a boolean if a field has been set.

### GetName

`func (o *OswStatPortVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswStatPortVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswStatPortVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OswStatPortVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperation

`func (o *OswStatPortVO) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *OswStatPortVO) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *OswStatPortVO) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *OswStatPortVO) HasOperation() bool`

HasOperation returns a boolean if a field has been set.

### GetOswStandPort

`func (o *OswStatPortVO) GetOswStandPort() OswStandPortVO`

GetOswStandPort returns the OswStandPort field if non-nil, zero value otherwise.

### GetOswStandPortOk

`func (o *OswStatPortVO) GetOswStandPortOk() (*OswStandPortVO, bool)`

GetOswStandPortOk returns a tuple with the OswStandPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOswStandPort

`func (o *OswStatPortVO) SetOswStandPort(v OswStandPortVO)`

SetOswStandPort sets OswStandPort field to given value.

### HasOswStandPort

`func (o *OswStatPortVO) HasOswStandPort() bool`

HasOswStandPort returns a boolean if a field has been set.

### GetPort

`func (o *OswStatPortVO) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *OswStatPortVO) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *OswStatPortVO) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *OswStatPortVO) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPortStatus

`func (o *OswStatPortVO) GetPortStatus() OswStatPortStatusVO`

GetPortStatus returns the PortStatus field if non-nil, zero value otherwise.

### GetPortStatusOk

`func (o *OswStatPortVO) GetPortStatusOk() (*OswStatPortStatusVO, bool)`

GetPortStatusOk returns a tuple with the PortStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortStatus

`func (o *OswStatPortVO) SetPortStatus(v OswStatPortStatusVO)`

SetPortStatus sets PortStatus field to given value.

### HasPortStatus

`func (o *OswStatPortVO) HasPortStatus() bool`

HasPortStatus returns a boolean if a field has been set.

### GetStackPortsGroupIndex

`func (o *OswStatPortVO) GetStackPortsGroupIndex() int32`

GetStackPortsGroupIndex returns the StackPortsGroupIndex field if non-nil, zero value otherwise.

### GetStackPortsGroupIndexOk

`func (o *OswStatPortVO) GetStackPortsGroupIndexOk() (*int32, bool)`

GetStackPortsGroupIndexOk returns a tuple with the StackPortsGroupIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStackPortsGroupIndex

`func (o *OswStatPortVO) SetStackPortsGroupIndex(v int32)`

SetStackPortsGroupIndex sets StackPortsGroupIndex field to given value.

### HasStackPortsGroupIndex

`func (o *OswStatPortVO) HasStackPortsGroupIndex() bool`

HasStackPortsGroupIndex returns a boolean if a field has been set.

### GetType

`func (o *OswStatPortVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OswStatPortVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OswStatPortVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *OswStatPortVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


