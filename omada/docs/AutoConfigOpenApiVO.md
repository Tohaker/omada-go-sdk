# AutoConfigOpenApiVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | **int32** | Location index, which refers to \&quot;Get location and ISP info\&quot;. | 
**MobileISP** | **int32** | MobileISP index, which refers to \&quot;Get location and ISP info\&quot;. | 

## Methods

### NewAutoConfigOpenApiVO

`func NewAutoConfigOpenApiVO(location int32, mobileISP int32, ) *AutoConfigOpenApiVO`

NewAutoConfigOpenApiVO instantiates a new AutoConfigOpenApiVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutoConfigOpenApiVOWithDefaults

`func NewAutoConfigOpenApiVOWithDefaults() *AutoConfigOpenApiVO`

NewAutoConfigOpenApiVOWithDefaults instantiates a new AutoConfigOpenApiVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *AutoConfigOpenApiVO) GetLocation() int32`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AutoConfigOpenApiVO) GetLocationOk() (*int32, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AutoConfigOpenApiVO) SetLocation(v int32)`

SetLocation sets Location field to given value.


### GetMobileISP

`func (o *AutoConfigOpenApiVO) GetMobileISP() int32`

GetMobileISP returns the MobileISP field if non-nil, zero value otherwise.

### GetMobileISPOk

`func (o *AutoConfigOpenApiVO) GetMobileISPOk() (*int32, bool)`

GetMobileISPOk returns a tuple with the MobileISP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileISP

`func (o *AutoConfigOpenApiVO) SetMobileISP(v int32)`

SetMobileISP sets MobileISP field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


