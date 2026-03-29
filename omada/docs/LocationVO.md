# LocationVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Angle** | Pointer to **int32** | Angle | [optional] 
**Height** | Pointer to **float64** | Height | [optional] 
**InstallType** | Pointer to **int32** | The installation method affects the default mounting height: CEILING (0), DESK (1), PANEL (2), WALL (3), POLE (4). | [optional] 
**Located** | Pointer to **bool** | Whether the device is located,false: indicates the disassociation of the device from the map. | [optional] 
**Locked** | Pointer to **bool** | Whether is locked | [optional] 
**MapId** | Pointer to **string** | Map id | [optional] 
**PosX** | Pointer to **float64** | X-axis position | [optional] 
**PosY** | Pointer to **float64** | Y-axis position | [optional] 

## Methods

### NewLocationVO

`func NewLocationVO() *LocationVO`

NewLocationVO instantiates a new LocationVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationVOWithDefaults

`func NewLocationVOWithDefaults() *LocationVO`

NewLocationVOWithDefaults instantiates a new LocationVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAngle

`func (o *LocationVO) GetAngle() int32`

GetAngle returns the Angle field if non-nil, zero value otherwise.

### GetAngleOk

`func (o *LocationVO) GetAngleOk() (*int32, bool)`

GetAngleOk returns a tuple with the Angle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAngle

`func (o *LocationVO) SetAngle(v int32)`

SetAngle sets Angle field to given value.

### HasAngle

`func (o *LocationVO) HasAngle() bool`

HasAngle returns a boolean if a field has been set.

### GetHeight

`func (o *LocationVO) GetHeight() float64`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *LocationVO) GetHeightOk() (*float64, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *LocationVO) SetHeight(v float64)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *LocationVO) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### GetInstallType

`func (o *LocationVO) GetInstallType() int32`

GetInstallType returns the InstallType field if non-nil, zero value otherwise.

### GetInstallTypeOk

`func (o *LocationVO) GetInstallTypeOk() (*int32, bool)`

GetInstallTypeOk returns a tuple with the InstallType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallType

`func (o *LocationVO) SetInstallType(v int32)`

SetInstallType sets InstallType field to given value.

### HasInstallType

`func (o *LocationVO) HasInstallType() bool`

HasInstallType returns a boolean if a field has been set.

### GetLocated

`func (o *LocationVO) GetLocated() bool`

GetLocated returns the Located field if non-nil, zero value otherwise.

### GetLocatedOk

`func (o *LocationVO) GetLocatedOk() (*bool, bool)`

GetLocatedOk returns a tuple with the Located field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocated

`func (o *LocationVO) SetLocated(v bool)`

SetLocated sets Located field to given value.

### HasLocated

`func (o *LocationVO) HasLocated() bool`

HasLocated returns a boolean if a field has been set.

### GetLocked

`func (o *LocationVO) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *LocationVO) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *LocationVO) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *LocationVO) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetMapId

`func (o *LocationVO) GetMapId() string`

GetMapId returns the MapId field if non-nil, zero value otherwise.

### GetMapIdOk

`func (o *LocationVO) GetMapIdOk() (*string, bool)`

GetMapIdOk returns a tuple with the MapId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMapId

`func (o *LocationVO) SetMapId(v string)`

SetMapId sets MapId field to given value.

### HasMapId

`func (o *LocationVO) HasMapId() bool`

HasMapId returns a boolean if a field has been set.

### GetPosX

`func (o *LocationVO) GetPosX() float64`

GetPosX returns the PosX field if non-nil, zero value otherwise.

### GetPosXOk

`func (o *LocationVO) GetPosXOk() (*float64, bool)`

GetPosXOk returns a tuple with the PosX field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosX

`func (o *LocationVO) SetPosX(v float64)`

SetPosX sets PosX field to given value.

### HasPosX

`func (o *LocationVO) HasPosX() bool`

HasPosX returns a boolean if a field has been set.

### GetPosY

`func (o *LocationVO) GetPosY() float64`

GetPosY returns the PosY field if non-nil, zero value otherwise.

### GetPosYOk

`func (o *LocationVO) GetPosYOk() (*float64, bool)`

GetPosYOk returns a tuple with the PosY field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosY

`func (o *LocationVO) SetPosY(v float64)`

SetPosY sets PosY field to given value.

### HasPosY

`func (o *LocationVO) HasPosY() bool`

HasPosY returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


