# PackageCaptureConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArChannel** | Pointer to **int32** | Channel. | [optional] 
**CaptureMode** | Pointer to **int32** | Capture mode. 0: Local packet capture; 1: Flow-mode packet capture. | [optional] 
**Channel** | Pointer to **int32** | (Wireless)  0: 2.4GHz  1: 5GHz-1  2:5GHz-2 3: 6GHz | [optional] 
**DestIp** | Pointer to **string** | Destination IP address of the packet. | [optional] 
**DestMac** | Pointer to **string** | Destination MAC address of the packet. | [optional] 
**DestPort** | Pointer to **int32** | Destination L4 port of the packet. | [optional] 
**Duration** | Pointer to **int32** | Packet capture duration, in seconds. | [optional] 
**FilterRules** | Pointer to **string** | Filter rules. | [optional] 
**InterfaceId** | Pointer to **string** | Interface ID, for example: if interfaceType is network, interfaceId should be LAN network ID. LAN Network can be created using &#39;Create LAN network&#39; interface, and LAN Network ID can be obtained from &#39;Get LAN network list&#39; interface. | [optional] 
**InterfaceName** | Pointer to **string** | Interface name. | [optional] 
**InterfaceType** | Pointer to **int32** | Interface type. 0: Wired; 1: Wireless; 2: LAN Network. | [optional] 
**OtaFilterRules** | Pointer to **string** | Filter rules of air interface packet capture. | [optional] 
**Protocol** | Pointer to **int32** | Packet type/protocol. It&#39;s required when a Switch captures packets. | [optional] 
**SinglePackageSize** | Pointer to **int64** | Single package size, in bytes. | [optional] 
**SrcIp** | Pointer to **string** | IP address of the message sender. | [optional] 
**SrcMac** | Pointer to **int32** | Source L4 port of the packet. | [optional] 
**Stack** | Pointer to **bool** | Whether the device supports stacking. | [optional] 
**Unit** | Pointer to **int32** | Equipment unit ID. | [optional] 

## Methods

### NewPackageCaptureConfig

`func NewPackageCaptureConfig() *PackageCaptureConfig`

NewPackageCaptureConfig instantiates a new PackageCaptureConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPackageCaptureConfigWithDefaults

`func NewPackageCaptureConfigWithDefaults() *PackageCaptureConfig`

NewPackageCaptureConfigWithDefaults instantiates a new PackageCaptureConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArChannel

`func (o *PackageCaptureConfig) GetArChannel() int32`

GetArChannel returns the ArChannel field if non-nil, zero value otherwise.

### GetArChannelOk

`func (o *PackageCaptureConfig) GetArChannelOk() (*int32, bool)`

GetArChannelOk returns a tuple with the ArChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArChannel

`func (o *PackageCaptureConfig) SetArChannel(v int32)`

SetArChannel sets ArChannel field to given value.

### HasArChannel

`func (o *PackageCaptureConfig) HasArChannel() bool`

HasArChannel returns a boolean if a field has been set.

### GetCaptureMode

`func (o *PackageCaptureConfig) GetCaptureMode() int32`

GetCaptureMode returns the CaptureMode field if non-nil, zero value otherwise.

### GetCaptureModeOk

`func (o *PackageCaptureConfig) GetCaptureModeOk() (*int32, bool)`

GetCaptureModeOk returns a tuple with the CaptureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureMode

`func (o *PackageCaptureConfig) SetCaptureMode(v int32)`

SetCaptureMode sets CaptureMode field to given value.

### HasCaptureMode

`func (o *PackageCaptureConfig) HasCaptureMode() bool`

HasCaptureMode returns a boolean if a field has been set.

### GetChannel

`func (o *PackageCaptureConfig) GetChannel() int32`

GetChannel returns the Channel field if non-nil, zero value otherwise.

### GetChannelOk

`func (o *PackageCaptureConfig) GetChannelOk() (*int32, bool)`

GetChannelOk returns a tuple with the Channel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel

`func (o *PackageCaptureConfig) SetChannel(v int32)`

SetChannel sets Channel field to given value.

### HasChannel

`func (o *PackageCaptureConfig) HasChannel() bool`

HasChannel returns a boolean if a field has been set.

### GetDestIp

`func (o *PackageCaptureConfig) GetDestIp() string`

GetDestIp returns the DestIp field if non-nil, zero value otherwise.

### GetDestIpOk

`func (o *PackageCaptureConfig) GetDestIpOk() (*string, bool)`

GetDestIpOk returns a tuple with the DestIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestIp

`func (o *PackageCaptureConfig) SetDestIp(v string)`

SetDestIp sets DestIp field to given value.

### HasDestIp

`func (o *PackageCaptureConfig) HasDestIp() bool`

HasDestIp returns a boolean if a field has been set.

### GetDestMac

`func (o *PackageCaptureConfig) GetDestMac() string`

GetDestMac returns the DestMac field if non-nil, zero value otherwise.

### GetDestMacOk

`func (o *PackageCaptureConfig) GetDestMacOk() (*string, bool)`

GetDestMacOk returns a tuple with the DestMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestMac

`func (o *PackageCaptureConfig) SetDestMac(v string)`

SetDestMac sets DestMac field to given value.

### HasDestMac

`func (o *PackageCaptureConfig) HasDestMac() bool`

HasDestMac returns a boolean if a field has been set.

### GetDestPort

`func (o *PackageCaptureConfig) GetDestPort() int32`

GetDestPort returns the DestPort field if non-nil, zero value otherwise.

### GetDestPortOk

`func (o *PackageCaptureConfig) GetDestPortOk() (*int32, bool)`

GetDestPortOk returns a tuple with the DestPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestPort

`func (o *PackageCaptureConfig) SetDestPort(v int32)`

SetDestPort sets DestPort field to given value.

### HasDestPort

`func (o *PackageCaptureConfig) HasDestPort() bool`

HasDestPort returns a boolean if a field has been set.

### GetDuration

`func (o *PackageCaptureConfig) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *PackageCaptureConfig) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *PackageCaptureConfig) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *PackageCaptureConfig) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFilterRules

`func (o *PackageCaptureConfig) GetFilterRules() string`

GetFilterRules returns the FilterRules field if non-nil, zero value otherwise.

### GetFilterRulesOk

`func (o *PackageCaptureConfig) GetFilterRulesOk() (*string, bool)`

GetFilterRulesOk returns a tuple with the FilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterRules

`func (o *PackageCaptureConfig) SetFilterRules(v string)`

SetFilterRules sets FilterRules field to given value.

### HasFilterRules

`func (o *PackageCaptureConfig) HasFilterRules() bool`

HasFilterRules returns a boolean if a field has been set.

### GetInterfaceId

`func (o *PackageCaptureConfig) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *PackageCaptureConfig) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *PackageCaptureConfig) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *PackageCaptureConfig) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceName

`func (o *PackageCaptureConfig) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *PackageCaptureConfig) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *PackageCaptureConfig) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *PackageCaptureConfig) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetInterfaceType

`func (o *PackageCaptureConfig) GetInterfaceType() int32`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *PackageCaptureConfig) GetInterfaceTypeOk() (*int32, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *PackageCaptureConfig) SetInterfaceType(v int32)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *PackageCaptureConfig) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetOtaFilterRules

`func (o *PackageCaptureConfig) GetOtaFilterRules() string`

GetOtaFilterRules returns the OtaFilterRules field if non-nil, zero value otherwise.

### GetOtaFilterRulesOk

`func (o *PackageCaptureConfig) GetOtaFilterRulesOk() (*string, bool)`

GetOtaFilterRulesOk returns a tuple with the OtaFilterRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtaFilterRules

`func (o *PackageCaptureConfig) SetOtaFilterRules(v string)`

SetOtaFilterRules sets OtaFilterRules field to given value.

### HasOtaFilterRules

`func (o *PackageCaptureConfig) HasOtaFilterRules() bool`

HasOtaFilterRules returns a boolean if a field has been set.

### GetProtocol

`func (o *PackageCaptureConfig) GetProtocol() int32`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *PackageCaptureConfig) GetProtocolOk() (*int32, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *PackageCaptureConfig) SetProtocol(v int32)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *PackageCaptureConfig) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetSinglePackageSize

`func (o *PackageCaptureConfig) GetSinglePackageSize() int64`

GetSinglePackageSize returns the SinglePackageSize field if non-nil, zero value otherwise.

### GetSinglePackageSizeOk

`func (o *PackageCaptureConfig) GetSinglePackageSizeOk() (*int64, bool)`

GetSinglePackageSizeOk returns a tuple with the SinglePackageSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSinglePackageSize

`func (o *PackageCaptureConfig) SetSinglePackageSize(v int64)`

SetSinglePackageSize sets SinglePackageSize field to given value.

### HasSinglePackageSize

`func (o *PackageCaptureConfig) HasSinglePackageSize() bool`

HasSinglePackageSize returns a boolean if a field has been set.

### GetSrcIp

`func (o *PackageCaptureConfig) GetSrcIp() string`

GetSrcIp returns the SrcIp field if non-nil, zero value otherwise.

### GetSrcIpOk

`func (o *PackageCaptureConfig) GetSrcIpOk() (*string, bool)`

GetSrcIpOk returns a tuple with the SrcIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcIp

`func (o *PackageCaptureConfig) SetSrcIp(v string)`

SetSrcIp sets SrcIp field to given value.

### HasSrcIp

`func (o *PackageCaptureConfig) HasSrcIp() bool`

HasSrcIp returns a boolean if a field has been set.

### GetSrcMac

`func (o *PackageCaptureConfig) GetSrcMac() int32`

GetSrcMac returns the SrcMac field if non-nil, zero value otherwise.

### GetSrcMacOk

`func (o *PackageCaptureConfig) GetSrcMacOk() (*int32, bool)`

GetSrcMacOk returns a tuple with the SrcMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSrcMac

`func (o *PackageCaptureConfig) SetSrcMac(v int32)`

SetSrcMac sets SrcMac field to given value.

### HasSrcMac

`func (o *PackageCaptureConfig) HasSrcMac() bool`

HasSrcMac returns a boolean if a field has been set.

### GetStack

`func (o *PackageCaptureConfig) GetStack() bool`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *PackageCaptureConfig) GetStackOk() (*bool, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *PackageCaptureConfig) SetStack(v bool)`

SetStack sets Stack field to given value.

### HasStack

`func (o *PackageCaptureConfig) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetUnit

`func (o *PackageCaptureConfig) GetUnit() int32`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *PackageCaptureConfig) GetUnitOk() (*int32, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *PackageCaptureConfig) SetUnit(v int32)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *PackageCaptureConfig) HasUnit() bool`

HasUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


