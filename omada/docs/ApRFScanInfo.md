# ApRFScanInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Channel2g** | Pointer to [**[]RFScanRadio**](RFScanRadio.md) | Channel 2g | [optional] 
**Channel5g** | Pointer to [**[]RFScanRadio**](RFScanRadio.md) | Channel 5g | [optional] 
**Channel5g2** | Pointer to [**[]RFScanRadio**](RFScanRadio.md) | Channel 5g2 | [optional] 
**Channel6g** | Pointer to [**[]RFScanRadio**](RFScanRadio.md) | Channel 6g | [optional] 
**CurrentChan2g** | Pointer to **string** | The 2G channel of AP, such as 1,6,11,13. It should be within the range of 1–13. | [optional] 
**CurrentChan5g** | Pointer to **string** | The 5G channel of AP, such as 36,161. It should be within the range of 36–161. | [optional] 
**CurrentChan5g2** | Pointer to **string** | The 5G2 channel of AP,such as 36,161. It should be within the range of 36–161. | [optional] 
**CurrentChan6g** | Pointer to **string** | The 6G channel of AP,such as 36,161. It should be within the range of 36–161. | [optional] 
**CurrentChanW2g** | Pointer to **int32** | The 2g channel bandwidth of the AP. It should be a value as follows: 2:20MHz, 3: 40MHz | [optional] 
**CurrentChanW5g** | Pointer to **int32** | The 5g channel bandwidth of the AP. It should be a value as follows: 2:20MHz, 3: 40MHz, 5: 80MHz2 menas 20MHz, 3 means 40MHz, 5 means 80MHz | [optional] 
**CurrentChanW5g2** | Pointer to **int32** | The 5g2 channel bandwidth of the AP. It should be a value as follows: 2:20MHz, 3: 40MHz, 5: 80MHz | [optional] 
**CurrentChanW6g** | Pointer to **int32** | The 6g channel bandwidth of the AP. It should be a value as follows: 2:20MHz, 3: 40MHz, 5: 80MHz | [optional] 
**Status** | Pointer to **int32** | Status should be a value as follows: 0: the scan result is displayed; 1: no scan result; 2: Scanning | [optional] 
**Time** | Pointer to **int64** | The scan time(13 bits), Unit: ms | [optional] 

## Methods

### NewApRFScanInfo

`func NewApRFScanInfo() *ApRFScanInfo`

NewApRFScanInfo instantiates a new ApRFScanInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApRFScanInfoWithDefaults

`func NewApRFScanInfoWithDefaults() *ApRFScanInfo`

NewApRFScanInfoWithDefaults instantiates a new ApRFScanInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChannel2g

`func (o *ApRFScanInfo) GetChannel2g() []RFScanRadio`

GetChannel2g returns the Channel2g field if non-nil, zero value otherwise.

### GetChannel2gOk

`func (o *ApRFScanInfo) GetChannel2gOk() (*[]RFScanRadio, bool)`

GetChannel2gOk returns a tuple with the Channel2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel2g

`func (o *ApRFScanInfo) SetChannel2g(v []RFScanRadio)`

SetChannel2g sets Channel2g field to given value.

### HasChannel2g

`func (o *ApRFScanInfo) HasChannel2g() bool`

HasChannel2g returns a boolean if a field has been set.

### GetChannel5g

`func (o *ApRFScanInfo) GetChannel5g() []RFScanRadio`

GetChannel5g returns the Channel5g field if non-nil, zero value otherwise.

### GetChannel5gOk

`func (o *ApRFScanInfo) GetChannel5gOk() (*[]RFScanRadio, bool)`

GetChannel5gOk returns a tuple with the Channel5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel5g

`func (o *ApRFScanInfo) SetChannel5g(v []RFScanRadio)`

SetChannel5g sets Channel5g field to given value.

### HasChannel5g

`func (o *ApRFScanInfo) HasChannel5g() bool`

HasChannel5g returns a boolean if a field has been set.

### GetChannel5g2

`func (o *ApRFScanInfo) GetChannel5g2() []RFScanRadio`

GetChannel5g2 returns the Channel5g2 field if non-nil, zero value otherwise.

### GetChannel5g2Ok

`func (o *ApRFScanInfo) GetChannel5g2Ok() (*[]RFScanRadio, bool)`

GetChannel5g2Ok returns a tuple with the Channel5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel5g2

`func (o *ApRFScanInfo) SetChannel5g2(v []RFScanRadio)`

SetChannel5g2 sets Channel5g2 field to given value.

### HasChannel5g2

`func (o *ApRFScanInfo) HasChannel5g2() bool`

HasChannel5g2 returns a boolean if a field has been set.

### GetChannel6g

`func (o *ApRFScanInfo) GetChannel6g() []RFScanRadio`

GetChannel6g returns the Channel6g field if non-nil, zero value otherwise.

### GetChannel6gOk

`func (o *ApRFScanInfo) GetChannel6gOk() (*[]RFScanRadio, bool)`

GetChannel6gOk returns a tuple with the Channel6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannel6g

`func (o *ApRFScanInfo) SetChannel6g(v []RFScanRadio)`

SetChannel6g sets Channel6g field to given value.

### HasChannel6g

`func (o *ApRFScanInfo) HasChannel6g() bool`

HasChannel6g returns a boolean if a field has been set.

### GetCurrentChan2g

`func (o *ApRFScanInfo) GetCurrentChan2g() string`

GetCurrentChan2g returns the CurrentChan2g field if non-nil, zero value otherwise.

### GetCurrentChan2gOk

`func (o *ApRFScanInfo) GetCurrentChan2gOk() (*string, bool)`

GetCurrentChan2gOk returns a tuple with the CurrentChan2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChan2g

`func (o *ApRFScanInfo) SetCurrentChan2g(v string)`

SetCurrentChan2g sets CurrentChan2g field to given value.

### HasCurrentChan2g

`func (o *ApRFScanInfo) HasCurrentChan2g() bool`

HasCurrentChan2g returns a boolean if a field has been set.

### GetCurrentChan5g

`func (o *ApRFScanInfo) GetCurrentChan5g() string`

GetCurrentChan5g returns the CurrentChan5g field if non-nil, zero value otherwise.

### GetCurrentChan5gOk

`func (o *ApRFScanInfo) GetCurrentChan5gOk() (*string, bool)`

GetCurrentChan5gOk returns a tuple with the CurrentChan5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChan5g

`func (o *ApRFScanInfo) SetCurrentChan5g(v string)`

SetCurrentChan5g sets CurrentChan5g field to given value.

### HasCurrentChan5g

`func (o *ApRFScanInfo) HasCurrentChan5g() bool`

HasCurrentChan5g returns a boolean if a field has been set.

### GetCurrentChan5g2

`func (o *ApRFScanInfo) GetCurrentChan5g2() string`

GetCurrentChan5g2 returns the CurrentChan5g2 field if non-nil, zero value otherwise.

### GetCurrentChan5g2Ok

`func (o *ApRFScanInfo) GetCurrentChan5g2Ok() (*string, bool)`

GetCurrentChan5g2Ok returns a tuple with the CurrentChan5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChan5g2

`func (o *ApRFScanInfo) SetCurrentChan5g2(v string)`

SetCurrentChan5g2 sets CurrentChan5g2 field to given value.

### HasCurrentChan5g2

`func (o *ApRFScanInfo) HasCurrentChan5g2() bool`

HasCurrentChan5g2 returns a boolean if a field has been set.

### GetCurrentChan6g

`func (o *ApRFScanInfo) GetCurrentChan6g() string`

GetCurrentChan6g returns the CurrentChan6g field if non-nil, zero value otherwise.

### GetCurrentChan6gOk

`func (o *ApRFScanInfo) GetCurrentChan6gOk() (*string, bool)`

GetCurrentChan6gOk returns a tuple with the CurrentChan6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChan6g

`func (o *ApRFScanInfo) SetCurrentChan6g(v string)`

SetCurrentChan6g sets CurrentChan6g field to given value.

### HasCurrentChan6g

`func (o *ApRFScanInfo) HasCurrentChan6g() bool`

HasCurrentChan6g returns a boolean if a field has been set.

### GetCurrentChanW2g

`func (o *ApRFScanInfo) GetCurrentChanW2g() int32`

GetCurrentChanW2g returns the CurrentChanW2g field if non-nil, zero value otherwise.

### GetCurrentChanW2gOk

`func (o *ApRFScanInfo) GetCurrentChanW2gOk() (*int32, bool)`

GetCurrentChanW2gOk returns a tuple with the CurrentChanW2g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChanW2g

`func (o *ApRFScanInfo) SetCurrentChanW2g(v int32)`

SetCurrentChanW2g sets CurrentChanW2g field to given value.

### HasCurrentChanW2g

`func (o *ApRFScanInfo) HasCurrentChanW2g() bool`

HasCurrentChanW2g returns a boolean if a field has been set.

### GetCurrentChanW5g

`func (o *ApRFScanInfo) GetCurrentChanW5g() int32`

GetCurrentChanW5g returns the CurrentChanW5g field if non-nil, zero value otherwise.

### GetCurrentChanW5gOk

`func (o *ApRFScanInfo) GetCurrentChanW5gOk() (*int32, bool)`

GetCurrentChanW5gOk returns a tuple with the CurrentChanW5g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChanW5g

`func (o *ApRFScanInfo) SetCurrentChanW5g(v int32)`

SetCurrentChanW5g sets CurrentChanW5g field to given value.

### HasCurrentChanW5g

`func (o *ApRFScanInfo) HasCurrentChanW5g() bool`

HasCurrentChanW5g returns a boolean if a field has been set.

### GetCurrentChanW5g2

`func (o *ApRFScanInfo) GetCurrentChanW5g2() int32`

GetCurrentChanW5g2 returns the CurrentChanW5g2 field if non-nil, zero value otherwise.

### GetCurrentChanW5g2Ok

`func (o *ApRFScanInfo) GetCurrentChanW5g2Ok() (*int32, bool)`

GetCurrentChanW5g2Ok returns a tuple with the CurrentChanW5g2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChanW5g2

`func (o *ApRFScanInfo) SetCurrentChanW5g2(v int32)`

SetCurrentChanW5g2 sets CurrentChanW5g2 field to given value.

### HasCurrentChanW5g2

`func (o *ApRFScanInfo) HasCurrentChanW5g2() bool`

HasCurrentChanW5g2 returns a boolean if a field has been set.

### GetCurrentChanW6g

`func (o *ApRFScanInfo) GetCurrentChanW6g() int32`

GetCurrentChanW6g returns the CurrentChanW6g field if non-nil, zero value otherwise.

### GetCurrentChanW6gOk

`func (o *ApRFScanInfo) GetCurrentChanW6gOk() (*int32, bool)`

GetCurrentChanW6gOk returns a tuple with the CurrentChanW6g field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentChanW6g

`func (o *ApRFScanInfo) SetCurrentChanW6g(v int32)`

SetCurrentChanW6g sets CurrentChanW6g field to given value.

### HasCurrentChanW6g

`func (o *ApRFScanInfo) HasCurrentChanW6g() bool`

HasCurrentChanW6g returns a boolean if a field has been set.

### GetStatus

`func (o *ApRFScanInfo) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApRFScanInfo) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApRFScanInfo) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ApRFScanInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTime

`func (o *ApRFScanInfo) GetTime() int64`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ApRFScanInfo) GetTimeOk() (*int64, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ApRFScanInfo) SetTime(v int64)`

SetTime sets Time field to given value.

### HasTime

`func (o *ApRFScanInfo) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


