package webscraping

import (
	"fmt"
	"math/rand"

	"github.com/golang-jwt/jwt/v4"
)

type loginResponse struct {
	Status                   int    `json:"status,omitempty"`
	SessionJwtToken          string `json:"sessionJwtToken,omitempty"`
	ProductCodesLicenseCodes []struct {
		ProductCode string `json:"productCode,omitempty"`
		LicenseCode string `json:"licenseCode,omitempty"`
	} `json:"productCodesLicenseCodes,omitempty"`
	LocalSessionPingInterval int `json:"localSessionPingInterval,omitempty"`
}

type SBSEPC5S struct {
	jwt.RegisteredClaims
	Sid string `json:"SID,omitempty"`
	Ts  string `json:"TS,omitempty"`
	Pk  string `json:"PK,omitempty"`
	Rd  string `json:"RD,omitempty"`
}

type SBSEPC5CS struct {
	jwt.RegisteredClaims
	Sid string `json:"SID,omitempty,omitempty"`
	Rd  string `json:"RD,omitempty"`
	Ts  string `json:"TS,omitempty"`
	Pk  string `json:"PK,omitempty"`
}

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func (b *BotSubaru) getDatasetId(datasetName string) string {
	for _, v := range b.UserBot.DatasetSettings {
		if v.DatasetName == datasetName {
			return v.DatasetID
		}
	}
	return ""
}

func (b *BotSubaru) returnFirma() string {
	return "jobId=1|dataSetId=" + b.getDatasetId("Subaru ATV") + "|locale=en-US|busReg=SUZ|LA|dataSetId=" + b.getDatasetId("Subaru Automotive") + "|locale=en-US|busReg=SUZ|LA|dataSetId=" + b.getDatasetId("Subaru Marine") + "|locale=en-US|busReg=SUZ|LA|dataSetId=" + b.getDatasetId("Subaru Motorcycle") + "|locale=en-US|busReg=SUZ|LA|userId=" + b.AccountBot.UserDetails.UserID
}

func (rb *BotSubaru) getCookieString() string {
	list := ""
	for key, cookie := range rb.Cookies {
		item := key + "=" + cookie.Value + ";"
		list = list + item
	}
	//list = list + rb.Csrf + ";" + rb.XfSession + ";xf_from_search=google; xf_csrf=5VOMNgavdKeCMhB8; _ga=GA1.2.864175125.1673533370; _gid=GA1.2.2061440800.1673533370;_ga=GA1.2.2005128399.1673553711; _gid=GA1.2.1758049683.1673553711; _gat_gtag_UA_62047822_1=1; xf_csrf=rQEBXUK_F0KBwmy7; __cf_bm=rBmeY2Kvbe5.uzpe9R12wH4Dwvz_PBbcut.9Gvb3T4c-1673553710-0-AVqClLrNn63m3XgpndbFY//mx/pb6owgX2jTPt/E0fxd7fUIB/OLMSImzDY3kdAu9wRmlRHX+u3Qv2KLitSsyl6+m67tqHnfUeBdYcveclgSN6nen+Ly6v9POi//bFkLAYP9eurO2dbOxC3np0cZNaM=;"
	fmt.Println("SETEANDO COOKIES " + list)
	return list
}

func (b *BotSubaru) RandomString(length int) string {
	ll := len(chars)
	bit := make([]byte, length)
	rand.Read(bit) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		bit[i] = chars[int(bit[i])%ll]
	}
	return string(bit)
}

// https://snaponepc.com/epc-services/equipment/search
type SearchVIN struct {
	VinSearchResults []struct {
		DatasetName string `json:"datasetName,omitempty"`
		Vins        []VIN  `json:"vins,omitempty"`
		Columns     []struct {
			Field  string `json:"field,omitempty"`
			Header string `json:"header,omitempty"`
		} `json:"columns,omitempty"`
	} `json:"vinSearchResults,omitempty"`
}
type VIN struct {
	DatasetID              string        `json:"datasetId,omitempty"`
	SerializedPath         string        `json:"serializedPath,omitempty"`
	DatasetName            string        `json:"datasetName,omitempty"`
	ModelName              string        `json:"modelName,omitempty"`
	ModelQualifierName     string        `json:"modelQualifierName,omitempty"`
	EquipmentName          string        `json:"equipmentName,omitempty"`
	ID                     string        `json:"id,omitempty"`
	Vin                    string        `json:"vin,omitempty"`
	FormattedVin           string        `json:"formattedVin,omitempty"`
	BusinessRegion         int           `json:"businessRegion,omitempty"`
	BusinessRegionName     string        `json:"businessRegionName,omitempty"`
	ExternalBusinessRegion bool          `json:"externalBusinessRegion,omitempty"`
	VinNote                []interface{} `json:"vinNote,omitempty"`
	EquipmentRefID         string        `json:"equipmentRefId,omitempty"`
	EquipmentKey           string        `json:"equipmentKey,omitempty"`
	RangeLookup            bool          `json:"rangeLookup,omitempty"`
	VinResolution          string        `json:"vinResolution,omitempty"`
	EinID                  string        `json:"einId,omitempty"`
	ValidationString       string        `json:"validationString,omitempty"`
	HasUserVinNote         bool          `json:"hasUserVinNote,omitempty"`
	VinRecalled            bool          `json:"vinRecalled,omitempty"`
	Categories             []Category    `json:"categories,omitempty"`
}

// https://snaponepc.com/epc-services/auth/account
type Account struct {
	UserDetails struct {
		UserName     string `json:"userName,omitempty"`
		UserID       string `json:"userId,omitempty"`
		LastAccess   int64  `json:"lastAccess,omitempty"`
		FirstName    string `json:"firstName,omitempty"`
		LastName     string `json:"lastName,omitempty"`
		EmailAddress string `json:"emailAddress,omitempty"`
	} `json:"userDetails,omitempty"`
	DealerDetails struct {
		Name         string `json:"name,omitempty"`
		Address1     string `json:"address1,omitempty"`
		Address2     string `json:"address2,omitempty"`
		Address3     string `json:"address3,omitempty"`
		City         string `json:"city,omitempty"`
		PostalCode   string `json:"postalCode,omitempty"`
		Country      string `json:"country,omitempty"`
		EmailAddress string `json:"emailAddress,omitempty"`
		Phone1       string `json:"phone1,omitempty"`
		Fax          string `json:"fax,omitempty"`
	} `json:"dealerDetails,omitempty"`
}

// https://snaponepc.com/epc-services/settings/user/
type User struct {
	ApplicationSettings struct {
		Locale                          string  `json:"locale,omitempty"`
		DateFormat                      string  `json:"dateFormat,omitempty"`
		TimeFormat                      string  `json:"timeFormat,omitempty"`
		NavigationStyle                 string  `json:"navigationStyle,omitempty"`
		PartsPanelPosition              string  `json:"partsPanelPosition,omitempty"`
		EulaVersionAccepted             float64 `json:"eulaVersionAccepted,omitempty"`
		ConfirmBeforeExitingApplication bool    `json:"confirmBeforeExitingApplication,omitempty"`
		ConfirmBeforeClearingPicklist   bool    `json:"confirmBeforeClearingPicklist,omitempty"`
		ConfirmBeforeClosingJob         bool    `json:"confirmBeforeClosingJob,omitempty"`
		WarnMeOnUpdatesAvailable        bool    `json:"warnMeOnUpdatesAvailable,omitempty"`
		SelectedPriceSource             struct {
		} `json:"selectedPriceSource,omitempty"`
		AddPartsToBottomOfPicklist  bool   `json:"addPartsToBottomOfPicklist,omitempty"`
		ShowPicklistByDefault       bool   `json:"showPicklistByDefault,omitempty"`
		SelectedPicklistPriceSource string `json:"selectedPicklistPriceSource,omitempty"`
		AutoClearPicklist           bool   `json:"autoClearPicklist,omitempty"`
		HideQtyPrompt               bool   `json:"hideQtyPrompt,omitempty"`
		DebugEnabled                bool   `json:"debugEnabled,omitempty"`
		ShowAllIndicators           bool   `json:"showAllIndicators,omitempty"`
		StatisticsEnabled           bool   `json:"statisticsEnabled,omitempty"`
	} `json:"applicationSettings,omitempty"`
	DatasetSettings []struct {
		DatasetID         string `json:"datasetId,omitempty"`
		DatasetName       string `json:"datasetName,omitempty"`
		Locale            string `json:"locale,omitempty"`
		BusinessRegionKey string `json:"businessRegionKey,omitempty"`
		BusinessRegion    int    `json:"businessRegion,omitempty"`
	} `json:"datasetSettings,omitempty"`
	EstimateSettings struct {
		Contact         string `json:"contact,omitempty"`
		PriceMultiplier string `json:"priceMultiplier,omitempty"`
		Currency        string `json:"currency,omitempty"`
		LaborRate       string `json:"laborRate,omitempty"`
		TaxRate         string `json:"taxRate,omitempty"`
		HidePartNumbers bool   `json:"hidePartNumbers,omitempty"`
		TaxLabor        bool   `json:"taxLabor,omitempty"`
	} `json:"estimateSettings,omitempty"`
}

type ResponseFilters struct {
	DatasetID       string `json:"datasetId,omitempty"`
	LevelColumnSort string `json:"levelColumnSort,omitempty"`
	Children        struct {
		ChildLevelTitle              string        `json:"childLevelTitle,omitempty"`
		ChildLevelType               string        `json:"childLevelType,omitempty"`
		ChildLevelSection            string        `json:"childLevelSection,omitempty"`
		ChildLevelIllustrated        bool          `json:"childLevelIllustrated,omitempty"`
		ChildLevelIllustrationWidth  int           `json:"childLevelIllustrationWidth,omitempty"`
		ChildLevelIllustrationHeight int           `json:"childLevelIllustrationHeight,omitempty"`
		ChildNodes                   []SubCategory `json:"childNodes,omitempty"`
	} `json:"children,omitempty"`
	Error bool `json:"error,omitempty"`
}

type SubCategory struct {
	ID             string         `json:"id,omitempty"`
	Name           string         `json:"name,omitempty"`
	HasNotes       bool           `json:"hasNotes,omitempty"`
	LeafNode       bool           `json:"leafNode,omitempty"`
	ImageID        string         `json:"imageId,omitempty"`
	SerializedPath string         `json:"serializedPath,omitempty"`
	Filtered       bool           `json:"filtered,omitempty"`
	Parts          ResponsePiezas `json:"parts,omitempty"`
}

type Category struct {
	ID             string        `json:"id,omitempty"`
	Name           string        `json:"name,omitempty"`
	HasNotes       bool          `json:"hasNotes,omitempty"`
	LeafNode       bool          `json:"leafNode,omitempty"`
	ImageID        string        `json:"imageId,omitempty"`
	SerializedPath string        `json:"serializedPath,omitempty"`
	Filtered       bool          `json:"filtered,omitempty"`
	SubCategory    []SubCategory `json:"subCategory,omitempty"`
}

type SuperSesion struct {
	PartItem PartItems `json:"partItem,omitempty"`
}

type ResponsePiezas struct {
	PageID            string           `json:"pageId,omitempty"`
	Illustrated       bool             `json:"illustrated,omitempty"`
	ImageID           string           `json:"imageId,omitempty"`
	PageCode          string           `json:"pageCode,omitempty"`
	PartItems         []PartItems      `json:"partItems,omitempty"`
	PageImages        []PageImages     `json:"pageImages,omitempty"`
	ColumnConfigs     []ColumnConfigs  `json:"columnConfigs,omitempty"`
	PageLimitExceeded bool             `json:"pageLimitExceeded,omitempty"`
	HasPageNotes      bool             `json:"hasPageNotes,omitempty"`
	SubParts          []ResponsePiezas `json:"subparts,omitempty,omitempty"`
}
type PartItems struct {
	PartID              string           `json:"partId,omitempty"`
	ParentPartID        string           `json:"parentPartId,omitempty"`
	SecondaryPartID     string           `json:"secondaryPartId,omitempty"`
	Manufacturer        string           `json:"manufacturer,omitempty"`
	PartNumber          string           `json:"partNumber,omitempty"`
	FormattedPartNumber string           `json:"formattedPartNumber,omitempty"`
	PartItemID          string           `json:"partItemId,omitempty"`
	ParentPartItemID    string           `json:"parentPartItemId,omitempty"`
	CalloutLabel        string           `json:"calloutLabel,omitempty"`
	PaddedCallout       string           `json:"paddedCallout,omitempty"`
	CrossCatKey         string           `json:"crossCatKey,omitempty"`
	Description         string           `json:"description,omitempty"`
	Quantity            string           `json:"quantity,omitempty"`
	PartType            string           `json:"partType,omitempty"`
	Indicators          []string         `json:"indicators,omitempty"`
	SuperSession        []SuperPartItems `json:"superSession,omitempty"`
	AlphaSort           string           `json:"alphaSort,omitempty"`
	AlphaSortSequence   string           `json:"alphaSortSequence,omitempty"`
	Filtered            bool             `json:"filtered,omitempty"`
	AddedManually       bool             `json:"addedManually,omitempty"`
	Remarks             string           `json:"remarks,omitempty,omitempty"`
}

type SuperPartItems struct {
	PartItem struct {
		PartID              string `json:"partId,omitempty"`
		ParentPartID        string `json:"parentPartId,omitempty"`
		SecondaryPartID     string `json:"secondaryPartId,omitempty"`
		Manufacturer        string `json:"manufacturer,omitempty"`
		PartNumber          string `json:"partNumber,omitempty"`
		FormattedPartNumber string `json:"formattedPartNumber,omitempty"`
		PartItemID          string `json:"partItemId,omitempty"`
		ParentPartItemID    string `json:"parentPartItemId,omitempty"`
		CalloutLabel        string `json:"calloutLabel,omitempty"`
		PaddedCallout       string `json:"paddedCallout,omitempty"`
		CrossCatKey         string `json:"crossCatKey,omitempty"`
		Description         string `json:"description,omitempty"`
		Quantity            string `json:"quantity,omitempty"`
		PartType            string `json:"partType,omitempty"`
		DynamicColumns      []struct {
			Code  string `json:"code,omitempty"`
			Value string `json:"value,omitempty"`
			Name  string `json:"name,omitempty"`
		} `json:"dynamicColumns,omitempty"`
		AlphaSort         string `json:"alphaSort,omitempty"`
		AlphaSortSequence string `json:"alphaSortSequence,omitempty"`
		Filtered          bool   `json:"filtered,omitempty"`
		AddedManually     bool   `json:"addedManually,omitempty"`
	} `json:"partItem,omitempty"`
	PriceBookSelected       bool  `json:"priceBookSelected,omitempty"`
	SupersessionAvailable   bool  `json:"supersessionAvailable,omitempty"`
	HistoryAvailable        bool  `json:"historyAvailable,omitempty"`
	AttachmentAvailable     bool  `json:"attachmentAvailable,omitempty"`
	AlternateAvailable      bool  `json:"alternateAvailable,omitempty"`
	PartAttributesAvailable bool  `json:"partAttributesAvailable,omitempty"`
	KitAvailable            bool  `json:"kitAvailable,omitempty"`
	IdentifiesAvailable     bool  `json:"identifiesAvailable,omitempty"`
	Images                  []any `json:"images,omitempty"`
}

type PageImages struct {
	ImageID    string `json:"imageId,omitempty"`
	PageID     string `json:"pageId,omitempty"`
	ImageTitle string `json:"imageTitle,omitempty"`
}
type ColumnConfigs struct {
	Key         string  `json:"key,omitempty"`
	Order       int     `json:"order,omitempty"`
	Width       int     `json:"width,omitempty"`
	PdfWidth    int     `json:"pdfWidth,omitempty"`
	MaxPdfWidth float64 `json:"maxPdfWidth,omitempty"`
	MinWidth    int     `json:"minWidth,omitempty"`
	MaxWidth    int     `json:"maxWidth,omitempty"`
	Resizable   bool    `json:"resizable,omitempty"`
	Title       string  `json:"title,omitempty"`
	Visible     bool    `json:"visible,omitempty"`
	Override    string  `json:"override,omitempty"`
}
