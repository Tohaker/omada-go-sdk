# TagOutItemOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassType** | **int32** | Class type should be a value as follows: 1: class 1, 2: class 2, 3: class 3, 0: others. | 
**Dscp** | **string** | The DSCP value selected in the Precedence configuration should be a value as follows: 8: IP precedence 1; 16: IP precedence 2; 24: IP precedence 3; 32: IP precedence 4; 40: IP precedence 5; 48: IP precedence 6; 56: IP precedence 7; 10: AF Class 1 (Low Drop); 12: AF Class 1 (Medium Drop); 14: AF Class 1 (High Drop); 18: AF Class 2 (Low Drop); 20: AF Class 2 (Medium Drop); 22: AF Class 2 (High Drop); 26: AF Class 3 (Low Drop); 28: AF Class 3 (Medium Drop); 30: AF Class 3 (High Drop); 34: AF Class 4 (Low Drop); 36: AF Class 4 (Medium Drop); 38: AF Class 4 (High Drop); 46: EF Class. | 
**Enable** | **bool** | The status of Class Type, valid value is true or false. | 

## Methods

### NewTagOutItemOpenApiVO

`func NewTagOutItemOpenApiVO(classType int32, dscp string, enable bool, ) *TagOutItemOpenApiVO`

NewTagOutItemOpenApiVO instantiates a new TagOutItemOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTagOutItemOpenApiVOWithDefaults

`func NewTagOutItemOpenApiVOWithDefaults() *TagOutItemOpenApiVO`

NewTagOutItemOpenApiVOWithDefaults instantiates a new TagOutItemOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassType

`func (o *TagOutItemOpenApiVO) GetClassType() int32`

GetClassType returns the ClassType field if non-nil, zero value otherwise.

### GetClassTypeOk

`func (o *TagOutItemOpenApiVO) GetClassTypeOk() (*int32, bool)`

GetClassTypeOk returns a tuple with the ClassType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassType

`func (o *TagOutItemOpenApiVO) SetClassType(v int32)`

SetClassType sets ClassType field to given value.


### GetDscp

`func (o *TagOutItemOpenApiVO) GetDscp() string`

GetDscp returns the Dscp field if non-nil, zero value otherwise.

### GetDscpOk

`func (o *TagOutItemOpenApiVO) GetDscpOk() (*string, bool)`

GetDscpOk returns a tuple with the Dscp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDscp

`func (o *TagOutItemOpenApiVO) SetDscp(v string)`

SetDscp sets Dscp field to given value.


### GetEnable

`func (o *TagOutItemOpenApiVO) GetEnable() bool`

GetEnable returns the Enable field if non-nil, zero value otherwise.

### GetEnableOk

`func (o *TagOutItemOpenApiVO) GetEnableOk() (*bool, bool)`

GetEnableOk returns a tuple with the Enable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnable

`func (o *TagOutItemOpenApiVO) SetEnable(v bool)`

SetEnable sets Enable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


