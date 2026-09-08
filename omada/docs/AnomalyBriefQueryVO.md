# AnomalyBriefQueryVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnomalyCode** | **string** | For the values of Anomaly event code, refer to section 5.7.2.1 of the Open API Access | 
**Object** | **[]string** |  | 

## Methods

### NewAnomalyBriefQueryVO

`func NewAnomalyBriefQueryVO(anomalyCode string, object []string, ) *AnomalyBriefQueryVO`

NewAnomalyBriefQueryVO instantiates a new AnomalyBriefQueryVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnomalyBriefQueryVOWithDefaults

`func NewAnomalyBriefQueryVOWithDefaults() *AnomalyBriefQueryVO`

NewAnomalyBriefQueryVOWithDefaults instantiates a new AnomalyBriefQueryVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnomalyCode

`func (o *AnomalyBriefQueryVO) GetAnomalyCode() string`

GetAnomalyCode returns the AnomalyCode field if non-nil, zero value otherwise.

### GetAnomalyCodeOk

`func (o *AnomalyBriefQueryVO) GetAnomalyCodeOk() (*string, bool)`

GetAnomalyCodeOk returns a tuple with the AnomalyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnomalyCode

`func (o *AnomalyBriefQueryVO) SetAnomalyCode(v string)`

SetAnomalyCode sets AnomalyCode field to given value.


### GetObject

`func (o *AnomalyBriefQueryVO) GetObject() []string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AnomalyBriefQueryVO) GetObjectOk() (*[]string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AnomalyBriefQueryVO) SetObject(v []string)`

SetObject sets Object field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


