# VpnAdvancedSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dpd** | Pointer to **bool** | DPD of the VPN advanced setting. | [optional] 
**DpdInterval** | Pointer to **int32** | DPD interval of the VPN advanced setting. | [optional] 
**EncapsulationMode** | Pointer to **int32** | Encapsulation mode should be a value as follows: 0: Tunnel Mode; 1: Transport Mode. | [optional] 
**ExchangeMode** | Pointer to **int32** | Exchange mode should be a value as follows: 0: Main Mode; 1: Aggressive Mode. | [optional] 
**KeyExchangeVersion** | Pointer to **int32** | Key exchange version should be a value as follows: 0: IKEv1; 1: IKEv2. | [optional] 
**LocalIdType** | Pointer to **int32** | Local ID type should be a value as follows: 0: IP Address; 1: Name. | [optional] 
**LocalName** | Pointer to **string** | Local name of the VPN advanced setting. | [optional] 
**NegotiationMode** | Pointer to **int32** | Negotiation mode should be a value as follows: 0: Initiator; 1: Aggressive Mode. | [optional] 
**Pfs** | Pointer to **int32** | PFS should be a value as follows: 0: None; 1: dh1; 2: dh2; 3: dh5; 14: dh14; 15: dh15. | [optional] 
**Phase1Proposal1** | Pointer to **int32** | Phase1 proposal1 should be a value as follows: 0: MD5; 1: SHA1. | [optional] 
**Phase1Proposal2** | Pointer to **int32** | Phase1 proposal2 should be a value as follows: 0: DES; 1: 3DES; 2: AES128; 3: AES192; 4: AES256. | [optional] 
**Phase1Proposal3** | Pointer to **int32** | Phase1 proposal3 should be a value as follows: 0: DH1; 1: DH2; 2: DH5; 3: DH14; 4: DH15; 5: DH16; 6: DH19; 7: DH20; 8: DH21; 9: DH25; 10: DH26. | [optional] 
**Phase2Proposal1** | Pointer to **int32** | Phase2 proposal1 should be a value as follows: 0: AH; 1: ESP. | [optional] 
**Phase2Proposal2** | Pointer to **int32** | Phase2 proposal2 should be a value as follows: 0: MD; 1: SHA1. | [optional] 
**Phase2Proposal3** | Pointer to **int32** | Phase2 proposal3 should be a value as follows: 0: DES; 1: 3DES; 2: AES128; 3: AES192; 4: AES256. | [optional] 
**RemoteIdType** | Pointer to **int32** | Remote ID type should be a value as follows: 0: IP Address; 1: Name. | [optional] 
**RemoteName** | Pointer to **string** | Remote name of the VPN advanced setting. | [optional] 
**SaLifetime** | Pointer to **int32** | SA lifetime of the VPN advanced setting. | [optional] 
**SaLifetime2** | Pointer to **int32** | SA lifetime2 of the VPN advanced setting. | [optional] 

## Methods

### NewVpnAdvancedSettingOpenApiVO

`func NewVpnAdvancedSettingOpenApiVO() *VpnAdvancedSettingOpenApiVO`

NewVpnAdvancedSettingOpenApiVO instantiates a new VpnAdvancedSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVpnAdvancedSettingOpenApiVOWithDefaults

`func NewVpnAdvancedSettingOpenApiVOWithDefaults() *VpnAdvancedSettingOpenApiVO`

NewVpnAdvancedSettingOpenApiVOWithDefaults instantiates a new VpnAdvancedSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDpd

`func (o *VpnAdvancedSettingOpenApiVO) GetDpd() bool`

GetDpd returns the Dpd field if non-nil, zero value otherwise.

### GetDpdOk

`func (o *VpnAdvancedSettingOpenApiVO) GetDpdOk() (*bool, bool)`

GetDpdOk returns a tuple with the Dpd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpd

`func (o *VpnAdvancedSettingOpenApiVO) SetDpd(v bool)`

SetDpd sets Dpd field to given value.

### HasDpd

`func (o *VpnAdvancedSettingOpenApiVO) HasDpd() bool`

HasDpd returns a boolean if a field has been set.

### GetDpdInterval

`func (o *VpnAdvancedSettingOpenApiVO) GetDpdInterval() int32`

GetDpdInterval returns the DpdInterval field if non-nil, zero value otherwise.

### GetDpdIntervalOk

`func (o *VpnAdvancedSettingOpenApiVO) GetDpdIntervalOk() (*int32, bool)`

GetDpdIntervalOk returns a tuple with the DpdInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpdInterval

`func (o *VpnAdvancedSettingOpenApiVO) SetDpdInterval(v int32)`

SetDpdInterval sets DpdInterval field to given value.

### HasDpdInterval

`func (o *VpnAdvancedSettingOpenApiVO) HasDpdInterval() bool`

HasDpdInterval returns a boolean if a field has been set.

### GetEncapsulationMode

`func (o *VpnAdvancedSettingOpenApiVO) GetEncapsulationMode() int32`

GetEncapsulationMode returns the EncapsulationMode field if non-nil, zero value otherwise.

### GetEncapsulationModeOk

`func (o *VpnAdvancedSettingOpenApiVO) GetEncapsulationModeOk() (*int32, bool)`

GetEncapsulationModeOk returns a tuple with the EncapsulationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncapsulationMode

`func (o *VpnAdvancedSettingOpenApiVO) SetEncapsulationMode(v int32)`

SetEncapsulationMode sets EncapsulationMode field to given value.

### HasEncapsulationMode

`func (o *VpnAdvancedSettingOpenApiVO) HasEncapsulationMode() bool`

HasEncapsulationMode returns a boolean if a field has been set.

### GetExchangeMode

`func (o *VpnAdvancedSettingOpenApiVO) GetExchangeMode() int32`

GetExchangeMode returns the ExchangeMode field if non-nil, zero value otherwise.

### GetExchangeModeOk

`func (o *VpnAdvancedSettingOpenApiVO) GetExchangeModeOk() (*int32, bool)`

GetExchangeModeOk returns a tuple with the ExchangeMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeMode

`func (o *VpnAdvancedSettingOpenApiVO) SetExchangeMode(v int32)`

SetExchangeMode sets ExchangeMode field to given value.

### HasExchangeMode

`func (o *VpnAdvancedSettingOpenApiVO) HasExchangeMode() bool`

HasExchangeMode returns a boolean if a field has been set.

### GetKeyExchangeVersion

`func (o *VpnAdvancedSettingOpenApiVO) GetKeyExchangeVersion() int32`

GetKeyExchangeVersion returns the KeyExchangeVersion field if non-nil, zero value otherwise.

### GetKeyExchangeVersionOk

`func (o *VpnAdvancedSettingOpenApiVO) GetKeyExchangeVersionOk() (*int32, bool)`

GetKeyExchangeVersionOk returns a tuple with the KeyExchangeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyExchangeVersion

`func (o *VpnAdvancedSettingOpenApiVO) SetKeyExchangeVersion(v int32)`

SetKeyExchangeVersion sets KeyExchangeVersion field to given value.

### HasKeyExchangeVersion

`func (o *VpnAdvancedSettingOpenApiVO) HasKeyExchangeVersion() bool`

HasKeyExchangeVersion returns a boolean if a field has been set.

### GetLocalIdType

`func (o *VpnAdvancedSettingOpenApiVO) GetLocalIdType() int32`

GetLocalIdType returns the LocalIdType field if non-nil, zero value otherwise.

### GetLocalIdTypeOk

`func (o *VpnAdvancedSettingOpenApiVO) GetLocalIdTypeOk() (*int32, bool)`

GetLocalIdTypeOk returns a tuple with the LocalIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalIdType

`func (o *VpnAdvancedSettingOpenApiVO) SetLocalIdType(v int32)`

SetLocalIdType sets LocalIdType field to given value.

### HasLocalIdType

`func (o *VpnAdvancedSettingOpenApiVO) HasLocalIdType() bool`

HasLocalIdType returns a boolean if a field has been set.

### GetLocalName

`func (o *VpnAdvancedSettingOpenApiVO) GetLocalName() string`

GetLocalName returns the LocalName field if non-nil, zero value otherwise.

### GetLocalNameOk

`func (o *VpnAdvancedSettingOpenApiVO) GetLocalNameOk() (*string, bool)`

GetLocalNameOk returns a tuple with the LocalName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalName

`func (o *VpnAdvancedSettingOpenApiVO) SetLocalName(v string)`

SetLocalName sets LocalName field to given value.

### HasLocalName

`func (o *VpnAdvancedSettingOpenApiVO) HasLocalName() bool`

HasLocalName returns a boolean if a field has been set.

### GetNegotiationMode

`func (o *VpnAdvancedSettingOpenApiVO) GetNegotiationMode() int32`

GetNegotiationMode returns the NegotiationMode field if non-nil, zero value otherwise.

### GetNegotiationModeOk

`func (o *VpnAdvancedSettingOpenApiVO) GetNegotiationModeOk() (*int32, bool)`

GetNegotiationModeOk returns a tuple with the NegotiationMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegotiationMode

`func (o *VpnAdvancedSettingOpenApiVO) SetNegotiationMode(v int32)`

SetNegotiationMode sets NegotiationMode field to given value.

### HasNegotiationMode

`func (o *VpnAdvancedSettingOpenApiVO) HasNegotiationMode() bool`

HasNegotiationMode returns a boolean if a field has been set.

### GetPfs

`func (o *VpnAdvancedSettingOpenApiVO) GetPfs() int32`

GetPfs returns the Pfs field if non-nil, zero value otherwise.

### GetPfsOk

`func (o *VpnAdvancedSettingOpenApiVO) GetPfsOk() (*int32, bool)`

GetPfsOk returns a tuple with the Pfs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfs

`func (o *VpnAdvancedSettingOpenApiVO) SetPfs(v int32)`

SetPfs sets Pfs field to given value.

### HasPfs

`func (o *VpnAdvancedSettingOpenApiVO) HasPfs() bool`

HasPfs returns a boolean if a field has been set.

### GetPhase1Proposal1

`func (o *VpnAdvancedSettingOpenApiVO) GetPhase1Proposal1() int32`

GetPhase1Proposal1 returns the Phase1Proposal1 field if non-nil, zero value otherwise.

### GetPhase1Proposal1Ok

`func (o *VpnAdvancedSettingOpenApiVO) GetPhase1Proposal1Ok() (*int32, bool)`

GetPhase1Proposal1Ok returns a tuple with the Phase1Proposal1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1Proposal1

`func (o *VpnAdvancedSettingOpenApiVO) SetPhase1Proposal1(v int32)`

SetPhase1Proposal1 sets Phase1Proposal1 field to given value.

### HasPhase1Proposal1

`func (o *VpnAdvancedSettingOpenApiVO) HasPhase1Proposal1() bool`

HasPhase1Proposal1 returns a boolean if a field has been set.

### GetPhase1Proposal2

`func (o *VpnAdvancedSettingOpenApiVO) GetPhase1Proposal2() int32`

GetPhase1Proposal2 returns the Phase1Proposal2 field if non-nil, zero value otherwise.

### GetPhase1Proposal2Ok

`func (o *VpnAdvancedSettingOpenApiVO) GetPhase1Proposal2Ok() (*int32, bool)`

GetPhase1Proposal2Ok returns a tuple with the Phase1Proposal2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1Proposal2

`func (o *VpnAdvancedSettingOpenApiVO) SetPhase1Proposal2(v int32)`

SetPhase1Proposal2 sets Phase1Proposal2 field to given value.

### HasPhase1Proposal2

`func (o *VpnAdvancedSettingOpenApiVO) HasPhase1Proposal2() bool`

HasPhase1Proposal2 returns a boolean if a field has been set.

### GetPhase1Proposal3

`func (o *VpnAdvancedSettingOpenApiVO) GetPhase1Proposal3() int32`

GetPhase1Proposal3 returns the Phase1Proposal3 field if non-nil, zero value otherwise.

### GetPhase1Proposal3Ok

`func (o *VpnAdvancedSettingOpenApiVO) GetPhase1Proposal3Ok() (*int32, bool)`

GetPhase1Proposal3Ok returns a tuple with the Phase1Proposal3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase1Proposal3

`func (o *VpnAdvancedSettingOpenApiVO) SetPhase1Proposal3(v int32)`

SetPhase1Proposal3 sets Phase1Proposal3 field to given value.

### HasPhase1Proposal3

`func (o *VpnAdvancedSettingOpenApiVO) HasPhase1Proposal3() bool`

HasPhase1Proposal3 returns a boolean if a field has been set.

### GetPhase2Proposal1

`func (o *VpnAdvancedSettingOpenApiVO) GetPhase2Proposal1() int32`

GetPhase2Proposal1 returns the Phase2Proposal1 field if non-nil, zero value otherwise.

### GetPhase2Proposal1Ok

`func (o *VpnAdvancedSettingOpenApiVO) GetPhase2Proposal1Ok() (*int32, bool)`

GetPhase2Proposal1Ok returns a tuple with the Phase2Proposal1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Proposal1

`func (o *VpnAdvancedSettingOpenApiVO) SetPhase2Proposal1(v int32)`

SetPhase2Proposal1 sets Phase2Proposal1 field to given value.

### HasPhase2Proposal1

`func (o *VpnAdvancedSettingOpenApiVO) HasPhase2Proposal1() bool`

HasPhase2Proposal1 returns a boolean if a field has been set.

### GetPhase2Proposal2

`func (o *VpnAdvancedSettingOpenApiVO) GetPhase2Proposal2() int32`

GetPhase2Proposal2 returns the Phase2Proposal2 field if non-nil, zero value otherwise.

### GetPhase2Proposal2Ok

`func (o *VpnAdvancedSettingOpenApiVO) GetPhase2Proposal2Ok() (*int32, bool)`

GetPhase2Proposal2Ok returns a tuple with the Phase2Proposal2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Proposal2

`func (o *VpnAdvancedSettingOpenApiVO) SetPhase2Proposal2(v int32)`

SetPhase2Proposal2 sets Phase2Proposal2 field to given value.

### HasPhase2Proposal2

`func (o *VpnAdvancedSettingOpenApiVO) HasPhase2Proposal2() bool`

HasPhase2Proposal2 returns a boolean if a field has been set.

### GetPhase2Proposal3

`func (o *VpnAdvancedSettingOpenApiVO) GetPhase2Proposal3() int32`

GetPhase2Proposal3 returns the Phase2Proposal3 field if non-nil, zero value otherwise.

### GetPhase2Proposal3Ok

`func (o *VpnAdvancedSettingOpenApiVO) GetPhase2Proposal3Ok() (*int32, bool)`

GetPhase2Proposal3Ok returns a tuple with the Phase2Proposal3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhase2Proposal3

`func (o *VpnAdvancedSettingOpenApiVO) SetPhase2Proposal3(v int32)`

SetPhase2Proposal3 sets Phase2Proposal3 field to given value.

### HasPhase2Proposal3

`func (o *VpnAdvancedSettingOpenApiVO) HasPhase2Proposal3() bool`

HasPhase2Proposal3 returns a boolean if a field has been set.

### GetRemoteIdType

`func (o *VpnAdvancedSettingOpenApiVO) GetRemoteIdType() int32`

GetRemoteIdType returns the RemoteIdType field if non-nil, zero value otherwise.

### GetRemoteIdTypeOk

`func (o *VpnAdvancedSettingOpenApiVO) GetRemoteIdTypeOk() (*int32, bool)`

GetRemoteIdTypeOk returns a tuple with the RemoteIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteIdType

`func (o *VpnAdvancedSettingOpenApiVO) SetRemoteIdType(v int32)`

SetRemoteIdType sets RemoteIdType field to given value.

### HasRemoteIdType

`func (o *VpnAdvancedSettingOpenApiVO) HasRemoteIdType() bool`

HasRemoteIdType returns a boolean if a field has been set.

### GetRemoteName

`func (o *VpnAdvancedSettingOpenApiVO) GetRemoteName() string`

GetRemoteName returns the RemoteName field if non-nil, zero value otherwise.

### GetRemoteNameOk

`func (o *VpnAdvancedSettingOpenApiVO) GetRemoteNameOk() (*string, bool)`

GetRemoteNameOk returns a tuple with the RemoteName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteName

`func (o *VpnAdvancedSettingOpenApiVO) SetRemoteName(v string)`

SetRemoteName sets RemoteName field to given value.

### HasRemoteName

`func (o *VpnAdvancedSettingOpenApiVO) HasRemoteName() bool`

HasRemoteName returns a boolean if a field has been set.

### GetSaLifetime

`func (o *VpnAdvancedSettingOpenApiVO) GetSaLifetime() int32`

GetSaLifetime returns the SaLifetime field if non-nil, zero value otherwise.

### GetSaLifetimeOk

`func (o *VpnAdvancedSettingOpenApiVO) GetSaLifetimeOk() (*int32, bool)`

GetSaLifetimeOk returns a tuple with the SaLifetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetime

`func (o *VpnAdvancedSettingOpenApiVO) SetSaLifetime(v int32)`

SetSaLifetime sets SaLifetime field to given value.

### HasSaLifetime

`func (o *VpnAdvancedSettingOpenApiVO) HasSaLifetime() bool`

HasSaLifetime returns a boolean if a field has been set.

### GetSaLifetime2

`func (o *VpnAdvancedSettingOpenApiVO) GetSaLifetime2() int32`

GetSaLifetime2 returns the SaLifetime2 field if non-nil, zero value otherwise.

### GetSaLifetime2Ok

`func (o *VpnAdvancedSettingOpenApiVO) GetSaLifetime2Ok() (*int32, bool)`

GetSaLifetime2Ok returns a tuple with the SaLifetime2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaLifetime2

`func (o *VpnAdvancedSettingOpenApiVO) SetSaLifetime2(v int32)`

SetSaLifetime2 sets SaLifetime2 field to given value.

### HasSaLifetime2

`func (o *VpnAdvancedSettingOpenApiVO) HasSaLifetime2() bool`

HasSaLifetime2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


