# ProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Badge** | Pointer to **string** |  | [optional] 
**BannerColor** | Pointer to **string** | optional; when present: \&quot;success\&quot; | \&quot;brand\&quot; — UI accent hint, not localized. Independent of bannerText | [optional] 
**BannerText** | Pointer to **string** |  | [optional] 
**BannerTooltip** | Pointer to **string** | optional; when present: \&quot;tiersMenu\&quot; — UI tooltip hint, not localized. Independent of bannerText | [optional] 
**BillingModel** | [**BillingModel**](BillingModel.md) |  | 
**Currency** | Pointer to **string** | ISO currency code, e.g. \&quot;eur\&quot; | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Disabled** | Pointer to **bool** |  | [optional] 
**DisplayOrder** | Pointer to **int32** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**I18n** | Pointer to [**map[string]ProductI18nResponse**](ProductI18nResponse.md) | per-language variants; keys match additionalLanguages in products_i18n.go, always present (serializes to {} when empty) | [optional] 
**Id** | **string** |  | 
**MaxSeats** | Pointer to **int32** |  | [optional] 
**MinSeats** | Pointer to **int32** |  | [optional] 
**Name** | **string** |  | 
**Price** | Pointer to **int32** | nil for tiered products; non-nil for flat (may be 0 for free) | [optional] 
**PriceSupplement** | Pointer to **string** |  | [optional] 
**Tiers** | Pointer to [**[]PriceTierResponse**](PriceTierResponse.md) | nil/omitted for flat-price products | [optional] 
**TrialDays** | Pointer to **int32** |  | [optional] 
**Type** | [**ProductType**](ProductType.md) |  | 

## Methods

### NewProductResponse

`func NewProductResponse(billingModel BillingModel, id string, name string, type_ ProductType, ) *ProductResponse`

NewProductResponse instantiates a new ProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProductResponseWithDefaults

`func NewProductResponseWithDefaults() *ProductResponse`

NewProductResponseWithDefaults instantiates a new ProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadge

`func (o *ProductResponse) GetBadge() string`

GetBadge returns the Badge field if non-nil, zero value otherwise.

### GetBadgeOk

`func (o *ProductResponse) GetBadgeOk() (*string, bool)`

GetBadgeOk returns a tuple with the Badge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadge

`func (o *ProductResponse) SetBadge(v string)`

SetBadge sets Badge field to given value.

### HasBadge

`func (o *ProductResponse) HasBadge() bool`

HasBadge returns a boolean if a field has been set.

### GetBannerColor

`func (o *ProductResponse) GetBannerColor() string`

GetBannerColor returns the BannerColor field if non-nil, zero value otherwise.

### GetBannerColorOk

`func (o *ProductResponse) GetBannerColorOk() (*string, bool)`

GetBannerColorOk returns a tuple with the BannerColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerColor

`func (o *ProductResponse) SetBannerColor(v string)`

SetBannerColor sets BannerColor field to given value.

### HasBannerColor

`func (o *ProductResponse) HasBannerColor() bool`

HasBannerColor returns a boolean if a field has been set.

### GetBannerText

`func (o *ProductResponse) GetBannerText() string`

GetBannerText returns the BannerText field if non-nil, zero value otherwise.

### GetBannerTextOk

`func (o *ProductResponse) GetBannerTextOk() (*string, bool)`

GetBannerTextOk returns a tuple with the BannerText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerText

`func (o *ProductResponse) SetBannerText(v string)`

SetBannerText sets BannerText field to given value.

### HasBannerText

`func (o *ProductResponse) HasBannerText() bool`

HasBannerText returns a boolean if a field has been set.

### GetBannerTooltip

`func (o *ProductResponse) GetBannerTooltip() string`

GetBannerTooltip returns the BannerTooltip field if non-nil, zero value otherwise.

### GetBannerTooltipOk

`func (o *ProductResponse) GetBannerTooltipOk() (*string, bool)`

GetBannerTooltipOk returns a tuple with the BannerTooltip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBannerTooltip

`func (o *ProductResponse) SetBannerTooltip(v string)`

SetBannerTooltip sets BannerTooltip field to given value.

### HasBannerTooltip

`func (o *ProductResponse) HasBannerTooltip() bool`

HasBannerTooltip returns a boolean if a field has been set.

### GetBillingModel

`func (o *ProductResponse) GetBillingModel() BillingModel`

GetBillingModel returns the BillingModel field if non-nil, zero value otherwise.

### GetBillingModelOk

`func (o *ProductResponse) GetBillingModelOk() (*BillingModel, bool)`

GetBillingModelOk returns a tuple with the BillingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingModel

`func (o *ProductResponse) SetBillingModel(v BillingModel)`

SetBillingModel sets BillingModel field to given value.


### GetCurrency

`func (o *ProductResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ProductResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ProductResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ProductResponse) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDescription

`func (o *ProductResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProductResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProductResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ProductResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDisabled

`func (o *ProductResponse) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *ProductResponse) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *ProductResponse) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *ProductResponse) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDisplayOrder

`func (o *ProductResponse) GetDisplayOrder() int32`

GetDisplayOrder returns the DisplayOrder field if non-nil, zero value otherwise.

### GetDisplayOrderOk

`func (o *ProductResponse) GetDisplayOrderOk() (*int32, bool)`

GetDisplayOrderOk returns a tuple with the DisplayOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayOrder

`func (o *ProductResponse) SetDisplayOrder(v int32)`

SetDisplayOrder sets DisplayOrder field to given value.

### HasDisplayOrder

`func (o *ProductResponse) HasDisplayOrder() bool`

HasDisplayOrder returns a boolean if a field has been set.

### GetFeatures

`func (o *ProductResponse) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *ProductResponse) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *ProductResponse) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *ProductResponse) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetI18n

`func (o *ProductResponse) GetI18n() map[string]ProductI18nResponse`

GetI18n returns the I18n field if non-nil, zero value otherwise.

### GetI18nOk

`func (o *ProductResponse) GetI18nOk() (*map[string]ProductI18nResponse, bool)`

GetI18nOk returns a tuple with the I18n field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetI18n

`func (o *ProductResponse) SetI18n(v map[string]ProductI18nResponse)`

SetI18n sets I18n field to given value.

### HasI18n

`func (o *ProductResponse) HasI18n() bool`

HasI18n returns a boolean if a field has been set.

### GetId

`func (o *ProductResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProductResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProductResponse) SetId(v string)`

SetId sets Id field to given value.


### GetMaxSeats

`func (o *ProductResponse) GetMaxSeats() int32`

GetMaxSeats returns the MaxSeats field if non-nil, zero value otherwise.

### GetMaxSeatsOk

`func (o *ProductResponse) GetMaxSeatsOk() (*int32, bool)`

GetMaxSeatsOk returns a tuple with the MaxSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSeats

`func (o *ProductResponse) SetMaxSeats(v int32)`

SetMaxSeats sets MaxSeats field to given value.

### HasMaxSeats

`func (o *ProductResponse) HasMaxSeats() bool`

HasMaxSeats returns a boolean if a field has been set.

### GetMinSeats

`func (o *ProductResponse) GetMinSeats() int32`

GetMinSeats returns the MinSeats field if non-nil, zero value otherwise.

### GetMinSeatsOk

`func (o *ProductResponse) GetMinSeatsOk() (*int32, bool)`

GetMinSeatsOk returns a tuple with the MinSeats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSeats

`func (o *ProductResponse) SetMinSeats(v int32)`

SetMinSeats sets MinSeats field to given value.

### HasMinSeats

`func (o *ProductResponse) HasMinSeats() bool`

HasMinSeats returns a boolean if a field has been set.

### GetName

`func (o *ProductResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProductResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProductResponse) SetName(v string)`

SetName sets Name field to given value.


### GetPrice

`func (o *ProductResponse) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ProductResponse) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ProductResponse) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ProductResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceSupplement

`func (o *ProductResponse) GetPriceSupplement() string`

GetPriceSupplement returns the PriceSupplement field if non-nil, zero value otherwise.

### GetPriceSupplementOk

`func (o *ProductResponse) GetPriceSupplementOk() (*string, bool)`

GetPriceSupplementOk returns a tuple with the PriceSupplement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSupplement

`func (o *ProductResponse) SetPriceSupplement(v string)`

SetPriceSupplement sets PriceSupplement field to given value.

### HasPriceSupplement

`func (o *ProductResponse) HasPriceSupplement() bool`

HasPriceSupplement returns a boolean if a field has been set.

### GetTiers

`func (o *ProductResponse) GetTiers() []PriceTierResponse`

GetTiers returns the Tiers field if non-nil, zero value otherwise.

### GetTiersOk

`func (o *ProductResponse) GetTiersOk() (*[]PriceTierResponse, bool)`

GetTiersOk returns a tuple with the Tiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTiers

`func (o *ProductResponse) SetTiers(v []PriceTierResponse)`

SetTiers sets Tiers field to given value.

### HasTiers

`func (o *ProductResponse) HasTiers() bool`

HasTiers returns a boolean if a field has been set.

### GetTrialDays

`func (o *ProductResponse) GetTrialDays() int32`

GetTrialDays returns the TrialDays field if non-nil, zero value otherwise.

### GetTrialDaysOk

`func (o *ProductResponse) GetTrialDaysOk() (*int32, bool)`

GetTrialDaysOk returns a tuple with the TrialDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrialDays

`func (o *ProductResponse) SetTrialDays(v int32)`

SetTrialDays sets TrialDays field to given value.

### HasTrialDays

`func (o *ProductResponse) HasTrialDays() bool`

HasTrialDays returns a boolean if a field has been set.

### GetType

`func (o *ProductResponse) GetType() ProductType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProductResponse) GetTypeOk() (*ProductType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProductResponse) SetType(v ProductType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


