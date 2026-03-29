# RadiusUserInfoOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | Build-in Radius profile user password. Password should contain 1 to 64 characters | 
**Username** | **string** | Build-in Radius profile user name. Username should contain 1 to 64 characters | 

## Methods

### NewRadiusUserInfoOpenApiVO

`func NewRadiusUserInfoOpenApiVO(password string, username string, ) *RadiusUserInfoOpenApiVO`

NewRadiusUserInfoOpenApiVO instantiates a new RadiusUserInfoOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRadiusUserInfoOpenApiVOWithDefaults

`func NewRadiusUserInfoOpenApiVOWithDefaults() *RadiusUserInfoOpenApiVO`

NewRadiusUserInfoOpenApiVOWithDefaults instantiates a new RadiusUserInfoOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *RadiusUserInfoOpenApiVO) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *RadiusUserInfoOpenApiVO) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *RadiusUserInfoOpenApiVO) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetUsername

`func (o *RadiusUserInfoOpenApiVO) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *RadiusUserInfoOpenApiVO) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *RadiusUserInfoOpenApiVO) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


