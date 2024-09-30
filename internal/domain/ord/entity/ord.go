package entity

// Основная структура заказа
type Order struct {
	Main                     Main                    `json:"main"`
	SBCity                   SBCity                  `json:"sb_city"`
	Included                 []Included              `json:"included"`
	SBMerchantProduct        MerchantProduct         `json:"sb_merchant_product"`
	SBDeliveryPointOfService *DeliveryPointOfService `json:"sb_delivery_point_of_service,omitempty"`
}

// Структура для Main (основная информация о заказе)
type Main struct {
	Code                  string         `json:"code"`
	State                 string         `json:"state"`
	Status                string         `json:"status"`
	Customer              Customer       `json:"customer"`
	PreOrder              bool           `json:"preOrder"`
	Assembled             bool           `json:"assembled,omitempty"`
	CreditTerm            int            `json:"creditTerm,omitempty"`
	TotalPrice            float64        `json:"totalPrice"`
	PaymentMode           string         `json:"paymentMode"`
	CreationDate          int64          `json:"creationDate"`
	DeliveryCost          float64        `json:"deliveryCost"`
	DeliveryMode          string         `json:"deliveryMode"`
	KaspiDelivery         *KaspiDelivery `json:"kaspiDelivery,omitempty"`
	OriginAddress         *Address       `json:"originAddress,omitempty"`
	DeliveryAddress       Address        `json:"deliveryAddress"`
	DeliverySlot          *DeliverySlot  `json:"deliverySlot,omitempty"`
	PickupPointID         string         `json:"pickupPointId"`
	IsKaspiDelivery       bool           `json:"isKaspiDelivery"`
	SignatureRequired     bool           `json:"signatureRequired"`
	ApprovedByBankDate    int64          `json:"approvedByBankDate"`
	PlannedDeliveryDate   int64          `json:"plannedDeliveryDate"`
	DeliveryCostForSeller float64        `json:"deliveryCostForSeller"`
}

// Структура для клиента
type Customer struct {
	ID        string `json:"id"`
	Name      string `json:"name,omitempty"`
	LastName  string `json:"lastName"`
	CellPhone string `json:"cellPhone"`
	FirstName string `json:"firstName"`
}

// Структура для адреса доставки
type Address struct {
	Town             string   `json:"town"`
	Building         string   `json:"building,omitempty"`
	District         string   `json:"district,omitempty"`
	Latitude         *float64 `json:"latitude,omitempty"`
	Longitude        *float64 `json:"longitude,omitempty"`
	Apartment        string   `json:"apartment,omitempty"`
	StreetName       string   `json:"streetName"`
	StreetNumber     string   `json:"streetNumber"`
	FormattedAddress string   `json:"formattedAddress"`
}

// Структура для доставки Kaspi
type KaspiDelivery struct {
	Express                         bool   `json:"express"`
	Waybill                         string `json:"waybill,omitempty"`
	WaybillNumber                   string `json:"waybillNumber,omitempty"`
	FirstMileCourier                string `json:"firstMileCourier,omitempty"`
	ReturnedToWarehouse             bool   `json:"returnedToWarehouse,omitempty"`
	CourierTransmissionDate         *int64 `json:"courierTransmissionDate,omitempty"`
	CourierTransmissionPlanningDate int64  `json:"courierTransmissionPlanningDate"`
}

// Структура для времени доставки (слот времени)
type DeliverySlot struct {
	To   string `json:"to"`
	From string `json:"from"`
}

// Структура для города
type SBCity struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Active bool   `json:"active"`
}

// Вложенная структура для товаров
type Included struct {
	ID            string        `json:"id"`
	Type          string        `json:"type"`
	Links         Links         `json:"links"`
	Attributes    Attributes    `json:"attributes"`
	Relationships Relationships `json:"relationships"`
}

// Структура для атрибутов товаров
type Attributes struct {
	Offer        Offer    `json:"offer"`
	Weight       float64  `json:"weight,omitempty"`
	Category     Category `json:"category"`
	Quantity     int      `json:"quantity"`
	UnitType     string   `json:"unitType"`
	BasePrice    float64  `json:"basePrice"`
	TotalPrice   float64  `json:"totalPrice"`
	EntryNumber  int      `json:"entryNumber"`
	DeliveryCost float64  `json:"deliveryCost"`
}

// Структура для предложения товара
type Offer struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Структура для категории товара
type Category struct {
	Code  string `json:"code"`
	Title string `json:"title"`
}

// Структура для связей (Relationships)
type Relationships struct {
	Product                RelationshipData `json:"product"`
	DeliveryPointOfService RelationshipData `json:"deliveryPointOfService,omitempty"`
}

// Структура для RelationshipData
type RelationshipData struct {
	Data  RelationshipDataAttributes `json:"data"`
	Links Links                      `json:"links"`
}

// Структура для атрибутов RelationshipData
type RelationshipDataAttributes struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

// Структура для ссылок (Links)
type Links struct {
	Self    string `json:"self"`
	Related string `json:"related"`
}

// Структура для товара (merchant product)
type MerchantProduct struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// Структура для пункта выдачи доставки (DeliveryPointOfService)
type DeliveryPointOfService struct {
	Address     Address `json:"address"`
	DisplayName string  `json:"displayName"`
}
