# OltConfigModifyDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**DeviceLocationDetailVO**](DeviceLocationDetailVO.md) |  | [optional] 
**Name** | Pointer to **string** | Device name,default value is the mac address of device | [optional] 
**TagIds** | Pointer to **[]string** | Tag ID list | [optional] 

## Methods

### NewOltConfigModifyDTO

`func NewOltConfigModifyDTO() *OltConfigModifyDTO`

NewOltConfigModifyDTO instantiates a new OltConfigModifyDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOltConfigModifyDTOWithDefaults

`func NewOltConfigModifyDTOWithDefaults() *OltConfigModifyDTO`

NewOltConfigModifyDTOWithDefaults instantiates a new OltConfigModifyDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *OltConfigModifyDTO) GetLocation() DeviceLocationDetailVO`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *OltConfigModifyDTO) GetLocationOk() (*DeviceLocationDetailVO, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *OltConfigModifyDTO) SetLocation(v DeviceLocationDetailVO)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *OltConfigModifyDTO) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *OltConfigModifyDTO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OltConfigModifyDTO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OltConfigModifyDTO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OltConfigModifyDTO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTagIds

`func (o *OltConfigModifyDTO) GetTagIds() []string`

GetTagIds returns the TagIds field if non-nil, zero value otherwise.

### GetTagIdsOk

`func (o *OltConfigModifyDTO) GetTagIdsOk() (*[]string, bool)`

GetTagIdsOk returns a tuple with the TagIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagIds

`func (o *OltConfigModifyDTO) SetTagIds(v []string)`

SetTagIds sets TagIds field to given value.

### HasTagIds

`func (o *OltConfigModifyDTO) HasTagIds() bool`

HasTagIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


