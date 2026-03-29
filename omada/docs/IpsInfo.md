# IpsInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomCategories** | Pointer to **[]int32** | Custom IDS/IPS categories list, if parameter[Dplevel] is 3, customCategories is needed.CustomCategories should be a list as follow: 1: Botcc, 2: Worm, 3: Malware, 4: Mobile_Malware, 6: P2P, 7: Tor, 8: Exploit, 9: Shellcode, 14: Activex, 15: DNS, 18: User Agents, 24: DShield | [optional] 
**DpLevel** | Pointer to **int32** | DpLevel should be a value as follows: 0: Low; 1: Medium; 2: High; 3: Custom | [optional] 
**Enable** | **bool** | Whether to enable IDS/IPS config. If parameter[enable] is true, parameter[ipsMode] and parameter[dplevel] are needed | 
**GeoEnable** | Pointer to **bool** | Whether to enable identifying the source country and destination country of attack ip addresses. | [optional] 
**IpsMode** | Pointer to **int32** | IpsMode should be a value as follows: 0: detect only; 1: detect and block | [optional] 
**TimeRangeId** | Pointer to **string** | This field represents Time Range ID. Time Range can be created using &#39;Create time range profile&#39; interface, and Time Range ID can be obtained from &#39;Get time range profile list&#39; interface. If parameter[timeRangeId] is null, IDS/IPS will be effective in all time ranges. | [optional] 

## Methods

### NewIpsInfo

`func NewIpsInfo(enable bool, ) *IpsInfo`

NewIpsInfo instantiates a new IpsInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpsInfoWithDefaults

`func NewIpsInfoWithDefaults() *IpsInfo`

NewIpsInfoWithDefaults instantiates a new IpsInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomCategories

`func (o *IpsInfo) GetCustomCategories() []int32`

GetCustomCategories returns the CustomCategories field if non-nil, zero value otherwise.

### GetCustomCategoriesOk

`func (o *IpsInfo) GetCustomCategoriesOk() (*[]int32, bool)`

GetCustomCategoriesOk returns a tuple with the CustomCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCategories

`func (o *IpsInfo) SetCustomCategories(v []int32)`

SetCustomCategories sets CustomCategories field to given value.

### HasCustomCategories

`func (o *IpsInfo) HasCustomCategories() bool`

HasCustomCategories returns a boolean if a field has been set.

### GetDpLevel

`func (o *IpsInfo) GetDpLevel() int32`

GetDpLevel returns the DpLevel field if non-nil, zero value otherwise.

### GetDpLevelOk

`func (o *IpsInfo) GetDpLevelOk() (*int32, bool)`

GetDpLevelOk returns a tuple with the DpLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDpLevel

`func (o *IpsInfo) SetDpLevel(v int32)`

SetDpLevel sets DpLevel field to given value.

### HasDpLevel

`func (o *IpsInfo) HasDpLevel() bool`

HasDpLevel returns a boolean if a field has been set.

### GetEnable

`func (o *IpsInfo) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *IpsInfo) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *IpsInfo) SetEnable(v bool)`

SetEnable sets Enable field to given value.


### GetGeoEnable

`func (o *IpsInfo) GetGeoEnable() bool`

GetGeoEnable returns the GeoEnable field if non-nil, zero value otherwise.

### GetGeoEnableOk

`func (o *IpsInfo) GetGeoEnableOk() (*bool, bool)`

GetGeoEnableOk returns a tuple with the GeoEnable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeoEnable

`func (o *IpsInfo) SetGeoEnable(v bool)`

SetGeoEnable sets GeoEnable field to given value.

### HasGeoEnable

`func (o *IpsInfo) HasGeoEnable() bool`

HasGeoEnable returns a boolean if a field has been set.

### GetIpsMode

`func (o *IpsInfo) GetIpsMode() int32`

GetIpsMode returns the IpsMode field if non-nil, zero value otherwise.

### GetIpsModeOk

`func (o *IpsInfo) GetIpsModeOk() (*int32, bool)`

GetIpsModeOk returns a tuple with the IpsMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpsMode

`func (o *IpsInfo) SetIpsMode(v int32)`

SetIpsMode sets IpsMode field to given value.

### HasIpsMode

`func (o *IpsInfo) HasIpsMode() bool`

HasIpsMode returns a boolean if a field has been set.

### GetTimeRangeId

`func (o *IpsInfo) GetTimeRangeId() string`

GetTimeRangeId returns the TimeRangeId field if non-nil, zero value otherwise.

### GetTimeRangeIdOk

`func (o *IpsInfo) GetTimeRangeIdOk() (*string, bool)`

GetTimeRangeIdOk returns a tuple with the TimeRangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeRangeId

`func (o *IpsInfo) SetTimeRangeId(v string)`

SetTimeRangeId sets TimeRangeId field to given value.

### HasTimeRangeId

`func (o *IpsInfo) HasTimeRangeId() bool`

HasTimeRangeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


