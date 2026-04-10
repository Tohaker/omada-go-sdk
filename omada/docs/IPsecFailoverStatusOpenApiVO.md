# IPsecFailoverStatusOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupWan** | Pointer to **string** | Backup WAN of the IPSec failover. | [optional] 
**Status** | Pointer to **bool** | Backup VPN status. | [optional] 

## Methods

### NewIPsecFailoverStatusOpenApiVO

`func NewIPsecFailoverStatusOpenApiVO() *IPsecFailoverStatusOpenApiVO`

NewIPsecFailoverStatusOpenApiVO instantiates a new IPsecFailoverStatusOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIPsecFailoverStatusOpenApiVOWithDefaults

`func NewIPsecFailoverStatusOpenApiVOWithDefaults() *IPsecFailoverStatusOpenApiVO`

NewIPsecFailoverStatusOpenApiVOWithDefaults instantiates a new IPsecFailoverStatusOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupWan

`func (o *IPsecFailoverStatusOpenApiVO) GetBackupWan() string`

GetBackupWan returns the BackupWan field if non-nil, zero value otherwise.

### GetBackupWanOk

`func (o *IPsecFailoverStatusOpenApiVO) GetBackupWanOk() (*string, bool)`

GetBackupWanOk returns a tuple with the BackupWan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupWan

`func (o *IPsecFailoverStatusOpenApiVO) SetBackupWan(v string)`

SetBackupWan sets BackupWan field to given value.

### HasBackupWan

`func (o *IPsecFailoverStatusOpenApiVO) HasBackupWan() bool`

HasBackupWan returns a boolean if a field has been set.

### GetStatus

`func (o *IPsecFailoverStatusOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IPsecFailoverStatusOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IPsecFailoverStatusOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IPsecFailoverStatusOpenApiVO) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


