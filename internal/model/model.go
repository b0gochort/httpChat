package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type NewChatReq struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type NewChatItem struct {
	ID              int64       `json:"id" reindex:"id,hash,pk"`
	Time            int64       `json:"time" reindex:"time,tree"`
	Message         MessageType `json:"message"`
	LastHostStaff   bool        `json:"last_host_staff" reindex:"last_host_staff,-"`
	UID             int64       `json:"uid" reindex:"uid,hash"`
	IP              string      `json:"ip" reindex:"ip,hash"`
	Category        string      `json:"category" reindex:"category,hash"`
	RequestTime     float64     `json:"request_time" reindex:"request_time,tree"`
	Priority        int64       `json:"priority" reindex:"priority,hash"`
	CategoryGuessed string      `json:"category_guessed" reindex:"category_guessed,hash"`
	MID             int64       `json:"mid" reindex:"mid,hash"`
	IsEnded         bool        `json:"is_ended" reindex:"is_ended,-"`
}

type MessageType struct {
	NumberOfUnread int    `json:"number_of_unread" reindex:"number_of_unread,-"`
	LastMessage    string `json:"last_message" reindex:"last_message,-"`
}

type NewChatRes struct {
	ID              int64       `json:"id"`
	Time            int64       `json:"time"`
	Message         MessageType `json:"message"`
	LastHostStaff   bool        `json:"last_host_staff"`
	UID             int64       `json:"uid"`
	IP              string      `json:"ip"`
	Category        string      `json:"category"`
	Name            string      `json:"name"`
	Surname         string      `json:"surname"`
	RequestTime     float64     `json:"request_time"`
	Priority        int64       `json:"priority" `
	CategoryGuessed string      `json:"category_guessed"`
	MID             int64       `json:"mid"`
	IsEnded         bool        `json:"is_ended"`
}

type NewMessageReq struct {
	ChatID  int64  `json:"chat_id"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type MessageItem struct {
	ID     int64    `json:"id" reindex:"id,hash,pk"`
	ChatId int64    `json:"chat_id" reindex:"chat_id,hash"`
	Time   int64    `json:"time" reindex:"time,tree"`
	Host   HostType `json:"host"`
	Text   string   `json:"text" reindex:"text,-"`
}

type HostType struct {
	UserId int64  `json:"user_id" reindex:"user_id,hash"`
	Sub    string `json:"sub" reindex:"sub,hash"`
}

type NewMessageRes struct {
	ID int64 `json:"id" `
}

type Message struct {
	ID     int64    `json:"id"`
	ChatId int64    `json:"chat_id"`
	Time   int64    `json:"time"`
	Host   HostType `json:"host"`
	Text   string   `json:"text"`
}

type JWTCustomClaims struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	jwt.RegisteredClaims
}

type CompetentItem struct {
	ID                     int64 `json:"id" reindex:"id,hash,pk"`
	UID                    int64 `json:"uid" reindex:"uid,hash"`
	PaymentIssue           int32 `json:"payment_issue" reindex:"payment_issue,hash"`
	CreateAccount          int32 `json:"create_account" reindex:"create_account,hash"`
	ContactCustomerService int32 `json:"contact_customer_service" reindex:"contact_customer_service,hash"`
	GetInvoice             int32 `json:"get_invoice" reindex:"get_invoice,hash"`
	TrackOrder             int32 `json:"track_order" reindex:"track_order,hash"`
	GetRefund              int32 `json:"get_refund" reindex:"get_refund,hash"`
	ContactHumanObject     int32 `json:"contact_human_object" reindex:"contact_human_object,hash"`
	RecoverPassword        int32 `json:"recover_password" reindex:"recover_password,hash"`
	ChangeOrder            int32 `json:"change_order" reindex:"change_order,hash"`
	DeleteAccount          int32 `json:"delete_account" reindex:"delete_account,hash"`
	Complaint              int32 `json:"complaint" reindex:"complaint,hash"`
	CheckInvoices          int32 `json:"check_invoices" reindex:"check_invoices,hash"`
	Review                 int32 `json:"review" reindex:"review,hash"`
	CheckRefundPolicy      int32 `json:"check_refund_policy" reindex:"check_refund_policy,hash"`
	DeliveryOptions        int32 `json:"delivery_options" reindex:"delivery_options,hash"`
	CheckCancellationFee   int32 `json:"check_cancellation_fee" reindex:"check_cancellation_fee,hash"`
	TrackRefund            int32 `json:"track_refund" reindex:"track_refund,hash"`
	CheckPaymentMethods    int32 `json:"check_payment_methods" reindex:"check_payment_methods,hash"`
	SwitchAccount          int32 `json:"switch_account" reindex:"switch_account,hash"`
	NewsletterSubscription int32 `json:"newsletter_subscription" reindex:"newsletter_subscription,hash"`
	DeliveryPeriod         int32 `json:"delivery_period" reindex:"delivery_period,hash"`
	EditAccount            int32 `json:"edit_account" reindex:"edit_account,hash"`
	RegistrationProblems   int32 `json:"registration_problems" reindex:"registration_problems,hash"`
	ChangeShippingAddress  int32 `json:"change_shipping_address" reindex:"change_shipping_address,hash"`
	SetUpShippingAddress   int32 `json:"set_up_shipping_address" reindex:"set_up_shipping_address,hash"`
	PlaceOrder             int32 `json:"place_order" reindex:"place_order,hash"`
	CancelOrder            int32 `json:"cancel_order" reindex:"cancel_order,hash"`
	CheckInvoice           int32 `json:"check_invoice" reindex:"check_invoice,hash"`
}

type NeuralService struct {
	Category string `json:"category" `
	Priority string `json:"priority" `
}
