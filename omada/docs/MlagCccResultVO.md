# MlagCccResultVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CccRslt** | Pointer to **string** |  | [optional] 
**CfgName** | Pointer to **string** | M-LAG group member configuration name | [optional] 
**LocalVal** | Pointer to **string** |  | [optional] 
**PathType** | Pointer to **int32** | M-LAG group member configuration path. It should be a value as follows:-1 : Settings -&gt; CLI Configuration0 : Device configuration page -&gt; config -&gt; General -&gt; Hash Algorithm1 : Device configuration page -&gt; ports -&gt; Profile Overrides -&gt; Operation(Aggregating)2 : Device configuration page -&gt; config -&gt; Services -&gt; Loopback Control -&gt; Loopback Detection3 : Device configuration page -&gt; ports -&gt; Profile Overrides -&gt; Loopback Control4 : Device configuration page -&gt; config -&gt; Services -&gt; Loopback Control -&gt; Spanning Tree5 : Device configuration page -&gt; config -&gt; Services -&gt; Loopback Control -&gt; CIST Priority6 : Device configuration page -&gt; config -&gt; Services -&gt; Loopback Control -&gt; MSTP Instance config7 : Device configuration page -&gt; ports -&gt; Profile Overrides -&gt; Spanning Tree Config8 : Device configuration page -&gt; ports -&gt; Profile Overrides -&gt; Loopback Control9 : Device configuration page -&gt; config -&gt; VLAN Interface10 : Settings -&gt; Transmission -&gt; VRRP -&gt; Optional Settings11 : Settings -&gt; Transmission -&gt; VRRP12 : Settings -&gt; Wired&amp;Wireless Networks -&gt; LAN -&gt; Networks -&gt; DHCP L2 Relay13 : Device configuration page -&gt; ports -&gt; LAG -&gt; Profile14 : Settings -&gt; Wired&amp;Wireless Networks -&gt; LAN -&gt; Networks -&gt; IGMP Snooping15 : Settings -&gt; Wired&amp;Wireless Networks -&gt; LAN -&gt; Networks -&gt; MLD Snooping | [optional] 
**PeerVal** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **int32** | M-LAG group member configuration type. It should be a value as follows:1 : Critical2 : Significant | [optional] 

## Methods

### NewMlagCccResultVO

`func NewMlagCccResultVO() *MlagCccResultVO`

NewMlagCccResultVO instantiates a new MlagCccResultVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMlagCccResultVOWithDefaults

`func NewMlagCccResultVOWithDefaults() *MlagCccResultVO`

NewMlagCccResultVOWithDefaults instantiates a new MlagCccResultVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCccRslt

`func (o *MlagCccResultVO) GetCccRslt() string`

GetCccRslt returns the CccRslt field if non-nil, zero value otherwise.

### GetCccRsltOk

`func (o *MlagCccResultVO) GetCccRsltOk() (*string, bool)`

GetCccRsltOk returns a tuple with the CccRslt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCccRslt

`func (o *MlagCccResultVO) SetCccRslt(v string)`

SetCccRslt sets CccRslt field to given value.

### HasCccRslt

`func (o *MlagCccResultVO) HasCccRslt() bool`

HasCccRslt returns a boolean if a field has been set.

### GetCfgName

`func (o *MlagCccResultVO) GetCfgName() string`

GetCfgName returns the CfgName field if non-nil, zero value otherwise.

### GetCfgNameOk

`func (o *MlagCccResultVO) GetCfgNameOk() (*string, bool)`

GetCfgNameOk returns a tuple with the CfgName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfgName

`func (o *MlagCccResultVO) SetCfgName(v string)`

SetCfgName sets CfgName field to given value.

### HasCfgName

`func (o *MlagCccResultVO) HasCfgName() bool`

HasCfgName returns a boolean if a field has been set.

### GetLocalVal

`func (o *MlagCccResultVO) GetLocalVal() string`

GetLocalVal returns the LocalVal field if non-nil, zero value otherwise.

### GetLocalValOk

`func (o *MlagCccResultVO) GetLocalValOk() (*string, bool)`

GetLocalValOk returns a tuple with the LocalVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalVal

`func (o *MlagCccResultVO) SetLocalVal(v string)`

SetLocalVal sets LocalVal field to given value.

### HasLocalVal

`func (o *MlagCccResultVO) HasLocalVal() bool`

HasLocalVal returns a boolean if a field has been set.

### GetPathType

`func (o *MlagCccResultVO) GetPathType() int32`

GetPathType returns the PathType field if non-nil, zero value otherwise.

### GetPathTypeOk

`func (o *MlagCccResultVO) GetPathTypeOk() (*int32, bool)`

GetPathTypeOk returns a tuple with the PathType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathType

`func (o *MlagCccResultVO) SetPathType(v int32)`

SetPathType sets PathType field to given value.

### HasPathType

`func (o *MlagCccResultVO) HasPathType() bool`

HasPathType returns a boolean if a field has been set.

### GetPeerVal

`func (o *MlagCccResultVO) GetPeerVal() string`

GetPeerVal returns the PeerVal field if non-nil, zero value otherwise.

### GetPeerValOk

`func (o *MlagCccResultVO) GetPeerValOk() (*string, bool)`

GetPeerValOk returns a tuple with the PeerVal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerVal

`func (o *MlagCccResultVO) SetPeerVal(v string)`

SetPeerVal sets PeerVal field to given value.

### HasPeerVal

`func (o *MlagCccResultVO) HasPeerVal() bool`

HasPeerVal returns a boolean if a field has been set.

### GetType

`func (o *MlagCccResultVO) GetType() int32`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MlagCccResultVO) GetTypeOk() (*int32, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MlagCccResultVO) SetType(v int32)`

SetType sets Type field to given value.

### HasType

`func (o *MlagCccResultVO) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


