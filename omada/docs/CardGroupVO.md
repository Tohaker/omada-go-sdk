# CardGroupVO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardGroupName** | Pointer to **string** |  | [optional] 
**CardList** | Pointer to **[]string** |  | [optional] 

## Methods

### NewCardGroupVO

`func NewCardGroupVO() *CardGroupVO`

NewCardGroupVO instantiates a new CardGroupVO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCardGroupVOWithDefaults

`func NewCardGroupVOWithDefaults() *CardGroupVO`

NewCardGroupVOWithDefaults instantiates a new CardGroupVO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardGroupName

`func (o *CardGroupVO) GetCardGroupName() string`

GetCardGroupName returns the CardGroupName field if non-nil, zero value otherwise.

### GetCardGroupNameOk

`func (o *CardGroupVO) GetCardGroupNameOk() (*string, bool)`

GetCardGroupNameOk returns a tuple with the CardGroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardGroupName

`func (o *CardGroupVO) SetCardGroupName(v string)`

SetCardGroupName sets CardGroupName field to given value.

### HasCardGroupName

`func (o *CardGroupVO) HasCardGroupName() bool`

HasCardGroupName returns a boolean if a field has been set.

### GetCardList

`func (o *CardGroupVO) GetCardList() []string`

GetCardList returns the CardList field if non-nil, zero value otherwise.

### GetCardListOk

`func (o *CardGroupVO) GetCardListOk() (*[]string, bool)`

GetCardListOk returns a tuple with the CardList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardList

`func (o *CardGroupVO) SetCardList(v []string)`

SetCardList sets CardList field to given value.

### HasCardList

`func (o *CardGroupVO) HasCardList() bool`

HasCardList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


