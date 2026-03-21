# BonjourServiceOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of Bonjour Service should contain 1 to 64 characters. | 
**ServiceIds** | **[]string** | The Service ID list of Bonjour Service. Service ID is a string in \&quot;_A._B.local\&quot; format, where \&quot;A\&quot; can be lowercase letters/numbers/hyphens (-)/underscore (_), and \&quot;B\&quot; should be lowercase letters. For example: _a1._b.local.Up to 3 entries are allowed for the serviceIds list. | 

## Methods

### NewBonjourServiceOpenApiVO

`func NewBonjourServiceOpenApiVO(name string, serviceIds []string, ) *BonjourServiceOpenApiVO`

NewBonjourServiceOpenApiVO instantiates a new BonjourServiceOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBonjourServiceOpenApiVOWithDefaults

`func NewBonjourServiceOpenApiVOWithDefaults() *BonjourServiceOpenApiVO`

NewBonjourServiceOpenApiVOWithDefaults instantiates a new BonjourServiceOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *BonjourServiceOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BonjourServiceOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BonjourServiceOpenApiVO) SetName(v string)`

SetName sets Name field to given value.


### GetServiceIds

`func (o *BonjourServiceOpenApiVO) GetServiceIds() []string`

GetServiceIds returns the ServiceIds field if non-nil, zero value otherwise.

### GetServiceIdsOk

`func (o *BonjourServiceOpenApiVO) GetServiceIdsOk() (*[]string, bool)`

GetServiceIdsOk returns a tuple with the ServiceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceIds

`func (o *BonjourServiceOpenApiVO) SetServiceIds(v []string)`

SetServiceIds sets ServiceIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


