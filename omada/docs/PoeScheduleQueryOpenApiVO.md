# PoeScheduleQueryOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | PoE Schedule Individual ID. | [optional] 
**Name** | **string** | PoE Schedule Name should contain 1 to 128 characters. | 
**NextExecute** | Pointer to **int64** | Show PoE Schedule next execution timestamp. | [optional] 
**NotInDst** | Pointer to **bool** | When notInDst is true, not displayed in DST. | [optional] 
**PoePortsMap** | **map[string][]int32** | Key:Mac(\&quot;String\&quot;), Value:Set of Ports(\&quot;Integer\&quot;) | 
**SiteId** | Pointer to **string** | Site ID | [optional] 
**Status** | **bool** | PoE Schedule Status. | 
**TurnOnTime** | **string** | Time Range ID, cannot be empty. | 

## Methods

### NewPoeScheduleQueryOpenApiVO

`func NewPoeScheduleQueryOpenApiVO(name string, poePortsMap map[string][]int32, status bool, turnOnTime string, ) *PoeScheduleQueryOpenApiVO`

NewPoeScheduleQueryOpenApiVO instantiates a new PoeScheduleQueryOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoeScheduleQueryOpenApiVOWithDefaults

`func NewPoeScheduleQueryOpenApiVOWithDefaults() *PoeScheduleQueryOpenApiVO`

NewPoeScheduleQueryOpenApiVOWithDefaults instantiates a new PoeScheduleQueryOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PoeScheduleQueryOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PoeScheduleQueryOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PoeScheduleQueryOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PoeScheduleQueryOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PoeScheduleQueryOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoeScheduleQueryOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoeScheduleQueryOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetNextExecute

`func (o *PoeScheduleQueryOpenApiVO) GetNextExecute() int64`

GetNextExecute returns the NextExecute field if non-nil, zero value otherwise.

### GetNextExecuteOk

`func (o *PoeScheduleQueryOpenApiVO) GetNextExecuteOk() (*int64, bool)`

GetNextExecuteOk returns a tuple with the NextExecute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextExecute

`func (o *PoeScheduleQueryOpenApiVO) SetNextExecute(v int64)`

SetNextExecute sets NextExecute field to given value.

### HasNextExecute

`func (o *PoeScheduleQueryOpenApiVO) HasNextExecute() bool`

HasNextExecute returns a boolean if a field has been set.

### GetNotInDst

`func (o *PoeScheduleQueryOpenApiVO) GetNotInDst() bool`

GetNotInDst returns the NotInDst field if non-nil, zero value otherwise.

### GetNotInDstOk

`func (o *PoeScheduleQueryOpenApiVO) GetNotInDstOk() (*bool, bool)`

GetNotInDstOk returns a tuple with the NotInDst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotInDst

`func (o *PoeScheduleQueryOpenApiVO) SetNotInDst(v bool)`

SetNotInDst sets NotInDst field to given value.

### HasNotInDst

`func (o *PoeScheduleQueryOpenApiVO) HasNotInDst() bool`

HasNotInDst returns a boolean if a field has been set.

### GetPoePortsMap

`func (o *PoeScheduleQueryOpenApiVO) GetPoePortsMap() map[string][]int32`

GetPoePortsMap returns the PoePortsMap field if non-nil, zero value otherwise.

### GetPoePortsMapOk

`func (o *PoeScheduleQueryOpenApiVO) GetPoePortsMapOk() (*map[string][]int32, bool)`

GetPoePortsMapOk returns a tuple with the PoePortsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePortsMap

`func (o *PoeScheduleQueryOpenApiVO) SetPoePortsMap(v map[string][]int32)`

SetPoePortsMap sets PoePortsMap field to given value.


### GetSiteId

`func (o *PoeScheduleQueryOpenApiVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PoeScheduleQueryOpenApiVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PoeScheduleQueryOpenApiVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *PoeScheduleQueryOpenApiVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStatus

`func (o *PoeScheduleQueryOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PoeScheduleQueryOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PoeScheduleQueryOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTurnOnTime

`func (o *PoeScheduleQueryOpenApiVO) GetTurnOnTime() string`

GetTurnOnTime returns the TurnOnTime field if non-nil, zero value otherwise.

### GetTurnOnTimeOk

`func (o *PoeScheduleQueryOpenApiVO) GetTurnOnTimeOk() (*string, bool)`

GetTurnOnTimeOk returns a tuple with the TurnOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnOnTime

`func (o *PoeScheduleQueryOpenApiVO) SetTurnOnTime(v string)`

SetTurnOnTime sets TurnOnTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


