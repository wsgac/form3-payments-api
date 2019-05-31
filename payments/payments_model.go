package payments

import "github.com/globalsign/mgo/bson"

type Payment struct {
	Type           string        `bson:"type" json:"type"`
	ID             bson.ObjectId `bson:"_id" json:"id"`
	Version        int           `bson:"version" json:"version"`
	OrganisationID string        `bson:"organisation_id" json:"organisation_id"`
	Attributes     Attributes    `bson:"attributes" json:"attributes"`
}

type Attributes struct {
	Amount               string             `bson:"amount" json:"amount"`
	BeneficiaryParty     Party              `bson:"beneficiary_party" json:"beneficiary_party"`
	ChargesInformation   ChargesInformation `bson:"charges_information" json:"charges_information"`
	Currency             string             `bson:"currency" json:"currency"`
	DebtorParty          Party              `bson:"debtor_party" json:"debtor_party"`
	EndToEndReference    string             `bson:"end_to_end_reference" json:"end_to_end_reference"`
	Fx                   Fx                 `bson:"fx" json:"fx"`
	NumericReference     string             `bson:"numeric_reference" json:"numeric_reference"`
	PaymentID            string             `bson:"payment_id" json:"payment_id"`
	PaymentPurpose       string             `bson:"payment_purpose" json:"payment_purpose"`
	PaymentScheme        string             `bson:"payment_scheme" json:"payment_scheme"`
	PaymentType          string             `bson:"payment_type" json:"payment_type"`
	ProcessingDate       string             `bson:"processing_date" json:"processing_date"`
	Reference            string             `bson:"reference" json:"reference"`
	SchemePaymentSubType string             `bson:"scheme_payment_sub_type" json:"scheme_payment_sub_type"`
	SchemePaymentType    string             `bson:"scheme_payment_type" json:"scheme_payment_type"`
	SponsorParty         Party              `bson:"sponsor_party" json:"sponsor_party"`
}

type Party struct {
	AccountName       string `bson:"account_name" json:"account_name"`
	AccountNumber     string `bson:"account_number" json:"account_number"`
	AccountNumberCode string `bson:"account_number_code" json:"account_number_code"`
	AccountType       int    `bson:"account_type" json:"account_type"`
	Address           string `bson:"address" json:"address"`
	BankID            string `bson:"bank_id" json:"bank_id"`
	BankIDCode        string `bson:"bank_id_code" json:"bank_id_code"`
	Name              string `bson:"name" json:"name"`
}

type Fx struct {
	ContractReference string `bson:"contract_reference" json:"contract_reference"`
	ExchangeRate      string `bson:"exchange_rate" json:"exchange_rate"`
	OriginalAmount    string `bson:"original_amount" json:"original_amount"`
	OriginalCurrency  string `bson:"original_currency" json:"original_currency"`
}

type ChargesInformation struct {
	BearerCode              string   `bson:"bearer_code" json:"bearer_code"`
	SenderCharges           []Charge `bson:"sender_charges" json:"sender_charges"`
	ReceiverChargesAmount   string   `bson:"receiver_charges_amount" json:"receiver_charges_amount"`
	ReceiverChargesCurrency string   `bson:"receiver_charges_currency" json:"receiver_charges_currency"`
}

type Charge struct {
	Amount  string `bson:"amount" json:"amount"`
	Curency string `bson:"currency" json:"currency"`
}
