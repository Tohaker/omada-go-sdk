# PortScheduleQueryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Port Schedule individual ID. | [optional] 
**Name** | **string** | Port Schedule Name should contain 1 to 128 characters. | 
**NextExecute** | Pointer to **int64** | Show Poe Schedule next execution timestamp. | [optional] 
**NotInDst** | Pointer to **bool** | When notInDst is true, Not displayed in DST. | [optional] 
**PortsMap** | **map[string][]int32** | Key:MAC(\&quot;String\&quot;), Value:Set of Ports(\&quot;Integer\&quot;) | 
**SiteId** | Pointer to **string** | Site ID | [optional] 
**Status** | **bool** | Port Schedule Status. | 
**TurnOnTime** | **string** | Time Range ID, cannot be empty. | 

## Methods

### NewPortScheduleQueryOpenApiVO

`func NewPortScheduleQueryOpenApiVO(name string, portsMap map[string][]int32, status bool, turnOnTime string, ) *PortScheduleQueryOpenApiVO`

NewPortScheduleQueryOpenApiVO instantiates a new PortScheduleQueryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortScheduleQueryOpenApiVOWithDefaults

`func NewPortScheduleQueryOpenApiVOWithDefaults() *PortScheduleQueryOpenApiVO`

NewPortScheduleQueryOpenApiVOWithDefaults instantiates a new PortScheduleQueryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortScheduleQueryOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortScheduleQueryOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortScheduleQueryOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortScheduleQueryOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PortScheduleQueryOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortScheduleQueryOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortScheduleQueryOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetNextExecute

`func (o *PortScheduleQueryOpenApiVO) GetNextExecute() int64`

GetNextExecute returns the NextExecute field if non-nil, zero value otherwise.

### GetNextExecuteOk

`func (o *PortScheduleQueryOpenApiVO) GetNextExecuteOk() (*int64, bool)`

GetNextExecuteOk returns a tuple with the NextExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecute

`func (o *PortScheduleQueryOpenApiVO) SetNextExecute(v int64)`

SetNextExecute sets NextExecute field to given value.

### HasNextExecute

`func (o *PortScheduleQueryOpenApiVO) HasNextExecute() bool`

HasNextExecute returns a boolean if a field has been set.

### GetNotInDst

`func (o *PortScheduleQueryOpenApiVO) GetNotInDst() bool`

GetNotInDst returns the NotInDst field if non-nil, zero value otherwise.

### GetNotInDstOk

`func (o *PortScheduleQueryOpenApiVO) GetNotInDstOk() (*bool, bool)`

GetNotInDstOk returns a tuple with the NotInDst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotInDst

`func (o *PortScheduleQueryOpenApiVO) SetNotInDst(v bool)`

SetNotInDst sets NotInDst field to given value.

### HasNotInDst

`func (o *PortScheduleQueryOpenApiVO) HasNotInDst() bool`

HasNotInDst returns a boolean if a field has been set.

### GetPortsMap

`func (o *PortScheduleQueryOpenApiVO) GetPortsMap() map[string][]int32`

GetPortsMap returns the PortsMap field if non-nil, zero value otherwise.

### GetPortsMapOk

`func (o *PortScheduleQueryOpenApiVO) GetPortsMapOk() (*map[string][]int32, bool)`

GetPortsMapOk returns a tuple with the PortsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortsMap

`func (o *PortScheduleQueryOpenApiVO) SetPortsMap(v map[string][]int32)`

SetPortsMap sets PortsMap field to given value.


### GetSiteId

`func (o *PortScheduleQueryOpenApiVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PortScheduleQueryOpenApiVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PortScheduleQueryOpenApiVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *PortScheduleQueryOpenApiVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStatus

`func (o *PortScheduleQueryOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortScheduleQueryOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortScheduleQueryOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTurnOnTime

`func (o *PortScheduleQueryOpenApiVO) GetTurnOnTime() string`

GetTurnOnTime returns the TurnOnTime field if non-nil, zero value otherwise.

### GetTurnOnTimeOk

`func (o *PortScheduleQueryOpenApiVO) GetTurnOnTimeOk() (*string, bool)`

GetTurnOnTimeOk returns a tuple with the TurnOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnOnTime

`func (o *PortScheduleQueryOpenApiVO) SetTurnOnTime(v string)`

SetTurnOnTime sets TurnOnTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


