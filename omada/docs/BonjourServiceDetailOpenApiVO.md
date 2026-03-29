# BonjourServiceDetailOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultProfile** | Pointer to **bool** | Indicating that this profile is default which can not be modified. | [optional] 
**Id** | Pointer to **string** | The id of Bonjour Service. | [optional] 
**Name** | Pointer to **string** | The name of Bonjour Service should contain 1 to 64 characters. | [optional] 
**ServiceIds** | Pointer to **[]string** | The Service ID list of Bonjour Service. Service ID is a string in \&quot;_A._B.local\&quot; format, where \&quot;A\&quot; can be lowercase letters/numbers/hyphens (-)/underscore (_), and \&quot;B\&quot; should be lowercase letters. For example: _a1._b.local.Up to 3 entries are allowed for the serviceIds list. | [optional] 

## Methods

### NewBonjourServiceDetailOpenApiVO

`func NewBonjourServiceDetailOpenApiVO() *BonjourServiceDetailOpenApiVO`

NewBonjourServiceDetailOpenApiVO instantiates a new BonjourServiceDetailOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBonjourServiceDetailOpenApiVOWithDefaults

`func NewBonjourServiceDetailOpenApiVOWithDefaults() *BonjourServiceDetailOpenApiVO`

NewBonjourServiceDetailOpenApiVOWithDefaults instantiates a new BonjourServiceDetailOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultProfile

`func (o *BonjourServiceDetailOpenApiVO) GetDefaultProfile() bool`

GetDefaultProfile returns the DefaultProfile field if non-nil, zero value otherwise.

### GetDefaultProfileOk

`func (o *BonjourServiceDetailOpenApiVO) GetDefaultProfileOk() (*bool, bool)`

GetDefaultProfileOk returns a tuple with the DefaultProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultProfile

`func (o *BonjourServiceDetailOpenApiVO) SetDefaultProfile(v bool)`

SetDefaultProfile sets DefaultProfile field to given value.

### HasDefaultProfile

`func (o *BonjourServiceDetailOpenApiVO) HasDefaultProfile() bool`

HasDefaultProfile returns a boolean if a field has been set.

### GetId

`func (o *BonjourServiceDetailOpenApiVO) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BonjourServiceDetailOpenApiVO) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BonjourServiceDetailOpenApiVO) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BonjourServiceDetailOpenApiVO) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *BonjourServiceDetailOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BonjourServiceDetailOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BonjourServiceDetailOpenApiVO) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BonjourServiceDetailOpenApiVO) HasName() bool`

HasName returns a boolean if a field has been set.

### GetServiceIds

`func (o *BonjourServiceDetailOpenApiVO) GetServiceIds() []string`

GetServiceIds returns the ServiceIds field if non-nil, zero value otherwise.

### GetServiceIdsOk

`func (o *BonjourServiceDetailOpenApiVO) GetServiceIdsOk() (*[]string, bool)`

GetServiceIdsOk returns a tuple with the ServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIds

`func (o *BonjourServiceDetailOpenApiVO) SetServiceIds(v []string)`

SetServiceIds sets ServiceIds field to given value.

### HasServiceIds

`func (o *BonjourServiceDetailOpenApiVO) HasServiceIds() bool`

HasServiceIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


