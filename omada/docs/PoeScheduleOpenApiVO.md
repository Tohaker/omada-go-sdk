# PoeScheduleOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | PoE Schedule Name should contain 1 to 128 characters. | 
**PoePortsMap** | **map[string][]int32** | Key:Mac(\&quot;String\&quot;), Value:Set of Ports(\&quot;Integer\&quot;) | 
**SiteId** | Pointer to **string** | Site ID | [optional] 
**Status** | **bool** | PoE Schedule Status. | 
**TurnOnTime** | **string** | Time Range ID, cannot be empty. | 

## Methods

### NewPoeScheduleOpenApiVO

`func NewPoeScheduleOpenApiVO(name string, poePortsMap map[string][]int32, status bool, turnOnTime string, ) *PoeScheduleOpenApiVO`

NewPoeScheduleOpenApiVO instantiates a new PoeScheduleOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoeScheduleOpenApiVOWithDefaults

`func NewPoeScheduleOpenApiVOWithDefaults() *PoeScheduleOpenApiVO`

NewPoeScheduleOpenApiVOWithDefaults instantiates a new PoeScheduleOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PoeScheduleOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PoeScheduleOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PoeScheduleOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetPoePortsMap

`func (o *PoeScheduleOpenApiVO) GetPoePortsMap() map[string][]int32`

GetPoePortsMap returns the PoePortsMap field if non-nil, zero value otherwise.

### GetPoePortsMapOk

`func (o *PoeScheduleOpenApiVO) GetPoePortsMapOk() (*map[string][]int32, bool)`

GetPoePortsMapOk returns a tuple with the PoePortsMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoePortsMap

`func (o *PoeScheduleOpenApiVO) SetPoePortsMap(v map[string][]int32)`

SetPoePortsMap sets PoePortsMap field to given value.


### GetSiteId

`func (o *PoeScheduleOpenApiVO) GetSiteId() string`

GetSiteId returns the SiteId field if non-nil, zero value otherwise.

### GetSiteIdOk

`func (o *PoeScheduleOpenApiVO) GetSiteIdOk() (*string, bool)`

GetSiteIdOk returns a tuple with the SiteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteId

`func (o *PoeScheduleOpenApiVO) SetSiteId(v string)`

SetSiteId sets SiteId field to given value.

### HasSiteId

`func (o *PoeScheduleOpenApiVO) HasSiteId() bool`

HasSiteId returns a boolean if a field has been set.

### GetStatus

`func (o *PoeScheduleOpenApiVO) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PoeScheduleOpenApiVO) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PoeScheduleOpenApiVO) SetStatus(v bool)`

SetStatus sets Status field to given value.


### GetTurnOnTime

`func (o *PoeScheduleOpenApiVO) GetTurnOnTime() string`

GetTurnOnTime returns the TurnOnTime field if non-nil, zero value otherwise.

### GetTurnOnTimeOk

`func (o *PoeScheduleOpenApiVO) GetTurnOnTimeOk() (*string, bool)`

GetTurnOnTimeOk returns a tuple with the TurnOnTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTurnOnTime

`func (o *PoeScheduleOpenApiVO) SetTurnOnTime(v string)`

SetTurnOnTime sets TurnOnTime field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


