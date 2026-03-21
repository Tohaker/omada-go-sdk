# OswStackConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Member** | [**[]OswStackMemberVO**](OswStackMemberVO.md) | Stack member list | 
**Name** | **string** | Stack Name | 

## Methods

### NewOswStackConfigOpenApiVO

`func NewOswStackConfigOpenApiVO(member []OswStackMemberVO, name string, ) *OswStackConfigOpenApiVO`

NewOswStackConfigOpenApiVO instantiates a new OswStackConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOswStackConfigOpenApiVOWithDefaults

`func NewOswStackConfigOpenApiVOWithDefaults() *OswStackConfigOpenApiVO`

NewOswStackConfigOpenApiVOWithDefaults instantiates a new OswStackConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMember

`func (o *OswStackConfigOpenApiVO) GetMember() []OswStackMemberVO`

GetMember returns the Member field if non-nil, zero value otherwise.

### GetMemberOk

`func (o *OswStackConfigOpenApiVO) GetMemberOk() (*[]OswStackMemberVO, bool)`

GetMemberOk returns a tuple with the Member field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMember

`func (o *OswStackConfigOpenApiVO) SetMember(v []OswStackMemberVO)`

SetMember sets Member field to given value.


### GetName

`func (o *OswStackConfigOpenApiVO) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OswStackConfigOpenApiVO) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OswStackConfigOpenApiVO) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


