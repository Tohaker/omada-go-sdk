# IPsecFailoverSettingOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupWan** | Pointer to **string** | Backup WAN of the IPSec failover. | [optional] 
**Failback** | Pointer to **bool** | Failback of the IPSec failover. | [optional] 
**FailbackTime** | Pointer to **int32** | Failback time should be within the range of 10–3600s. | [optional] 
**Failover** | Pointer to **bool** | Indicates whether to use the IPSec failover. | [optional] 

## Methods

### NewIPsecFailoverSettingOpenApiVO

`func NewIPsecFailoverSettingOpenApiVO() *IPsecFailoverSettingOpenApiVO`

NewIPsecFailoverSettingOpenApiVO instantiates a new IPsecFailoverSettingOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPsecFailoverSettingOpenApiVOWithDefaults

`func NewIPsecFailoverSettingOpenApiVOWithDefaults() *IPsecFailoverSettingOpenApiVO`

NewIPsecFailoverSettingOpenApiVOWithDefaults instantiates a new IPsecFailoverSettingOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupWan

`func (o *IPsecFailoverSettingOpenApiVO) GetBackupWan() string`

GetBackupWan returns the BackupWan field if non-nil, zero value otherwise.

### GetBackupWanOk

`func (o *IPsecFailoverSettingOpenApiVO) GetBackupWanOk() (*string, bool)`

GetBackupWanOk returns a tuple with the BackupWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupWan

`func (o *IPsecFailoverSettingOpenApiVO) SetBackupWan(v string)`

SetBackupWan sets BackupWan field to given value.

### HasBackupWan

`func (o *IPsecFailoverSettingOpenApiVO) HasBackupWan() bool`

HasBackupWan returns a boolean if a field has been set.

### GetFailback

`func (o *IPsecFailoverSettingOpenApiVO) GetFailback() bool`

GetFailback returns the Failback field if non-nil, zero value otherwise.

### GetFailbackOk

`func (o *IPsecFailoverSettingOpenApiVO) GetFailbackOk() (*bool, bool)`

GetFailbackOk returns a tuple with the Failback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailback

`func (o *IPsecFailoverSettingOpenApiVO) SetFailback(v bool)`

SetFailback sets Failback field to given value.

### HasFailback

`func (o *IPsecFailoverSettingOpenApiVO) HasFailback() bool`

HasFailback returns a boolean if a field has been set.

### GetFailbackTime

`func (o *IPsecFailoverSettingOpenApiVO) GetFailbackTime() int32`

GetFailbackTime returns the FailbackTime field if non-nil, zero value otherwise.

### GetFailbackTimeOk

`func (o *IPsecFailoverSettingOpenApiVO) GetFailbackTimeOk() (*int32, bool)`

GetFailbackTimeOk returns a tuple with the FailbackTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailbackTime

`func (o *IPsecFailoverSettingOpenApiVO) SetFailbackTime(v int32)`

SetFailbackTime sets FailbackTime field to given value.

### HasFailbackTime

`func (o *IPsecFailoverSettingOpenApiVO) HasFailbackTime() bool`

HasFailbackTime returns a boolean if a field has been set.

### GetFailover

`func (o *IPsecFailoverSettingOpenApiVO) GetFailover() bool`

GetFailover returns the Failover field if non-nil, zero value otherwise.

### GetFailoverOk

`func (o *IPsecFailoverSettingOpenApiVO) GetFailoverOk() (*bool, bool)`

GetFailoverOk returns a tuple with the Failover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover

`func (o *IPsecFailoverSettingOpenApiVO) SetFailover(v bool)`

SetFailover sets Failover field to given value.

### HasFailover

`func (o *IPsecFailoverSettingOpenApiVO) HasFailover() bool`

HasFailover returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


