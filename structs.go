package nexmo

type Links struct {
	Self  Link `json:"self,omitempty"`
	Next  Link `json:"next,omitempty"`
	Prev  Link `json:"prev,omitempty"`
	First Link `json:"first,omitempty"`
	Last  Link `json:"last,omitempty"`
}

type Link struct {
	Href string `json:"href,omitempty"`
}

type PhoneCallEndpoint struct {
	Type       string `json:"type"`
	Number     string `json:"number"`
	Dtmfanswer string `json:"dtmfAnswer,omitempty"`
}

type WebSocketCallEndpoint struct {
	Type        string      `json:"type"`
	URI         string      `json:"uri"`
	ContentType string      `json:"content-type,omitempty"`
	Headers     interface{} `json:"headers,omitempty"`
}

type SIPCallEndpoint struct {
	Type string `json:"type"`
	URI  string `json:"uri"`
}

type SendSMSRequest struct {
	APIKey          string `json:"api_key"`
	APISecret       string `json:"api_secret"`
	From            string `json:"from,omitempty"`
	To              string `json:"to,omitempty"`
	Type            string `json:"type,omitempty"`
	Text            string `json:"text,omitempty"`
	StatusReportReq int64  `json:"status-report-req,omitempty"`
	ClientRef       string `json:"client-ref,omitempty"`
	Vcard           string `json:"vcard,omitempty"`
	Vcal            string `json:"vcal,omitempty"`
	TTL             int64  `json:"ttl,omitempty"`
	Callback        string `json:"callback,omitempty"`
	MessageClass    int64  `json:"message-class,omitempty"`
	Udh             string `json:"udh,omitempty"`
	ProtocolID      int64  `json:"protocol-id,omitempty"`
	Body            string `json:"body,omitempty"`
	Title           string `json:"title,omitempty"`
	URL             string `json:"url,omitempty"`
	Validity        int64  `json:"validity,omitempty"`
}

func (r *SendSMSRequest) setApiCredentials(apiKey, apiSecret string) {
	r.APIKey = apiKey
	r.APISecret = apiSecret
}

type SendSMSResponse struct {
	MessageCount string                   `json:"message-count,omitempty"`
	Messages     []SendSMSResponseMessage `json:"messages,omitempty"`
}

type SendSMSResponseMessage struct {
	Status           string `json:"status,omitempty"`
	MessageID        string `json:"message-id,omitempty"`
	To               string `json:"to,omitempty"`
	ClientRef        string `json:"client-ref,omitempty"`
	RemainingBalance string `json:"remaining-balance,omitempty"`
	MessagePrice     string `json:"message-price,omitempty"`
	Network          string `json:"network,omitempty"`
	ErrorText        string `json:"error-text,omitempty"`
}

type CreateCallRequest struct {
	To               []interface{}       `json:"to"`
	From             map[string]string   `json:"from"`
	Ncco             []map[string]string `json:"ncco"`
	AnswerURL        []string            `json:"answer_url"`
	AnswerMethod     string              `json:"answer_method,omitempty"`
	EventURL         []string            `json:"event_url,omitempty"`
	EventMethod      string              `json:"event_method,omitempty"`
	MachineDetection string              `json:"machine_detection,omitempty"`
	LengthTimer      int64               `json:"length_timer,omitempty"`
	RingingTimer     int64               `json:"ringing_timer,omitempty"`
}

type CreateCallResponse struct {
	UUID             string `json:"uuid"`
	ConversationUUID string `json:"conversation_uuid"`
	Direction        string `json:"direction"`
	Status           string `json:"status"`
}

type SearchCallsRequest struct {
	Status           string `url:"status,omitempty"`
	DateStart        string `url:"date_start,omitempty"`
	DateEnd          string `url:"date_end,omitempty"`
	PageSize         int64  `url:"page_size,omitempty"`
	RecordIndex      int64  `url:"record_index,omitempty"`
	Order            string `url:"order,omitempty"`
	ConversationUUID string `url:"conversation_uuid,omitempty"`
}

type SearchCallsResponse struct {
	Count       int64         `json:"count,omitempty"`
	PageSize    int64         `json:"page_size,omitempty"`
	RecordIndex int64         `json:"record_index,omitempty"`
	Links       Links         `json:"_links,omitempty"`
	Embedded    EmbeddedCalls `json:"_embedded,omitempty"`
}

type EmbeddedCalls struct {
	Calls []CallInfo `json:"calls,omitempty"`
}

type CallInfo struct {
	UUID             string      `json:"uuid,omitempty"`
	ConversationUUID string      `json:"conversation_uuid,omitempty"`
	To               interface{} `json:"to,omitempty"`
	From             interface{} `json:"from,omitempty"`
	Status           string      `json:"status,omitempty"`
	Direction        string      `json:"direction,omitempty"`
	Rate             string      `json:"rate,omitempty"`
	Price            string      `json:"price,omitempty"`
	Duration         string      `json:"duration,omitempty"`
	Network          string      `json:"network,omitempty"`
	StartTime        string      `json:"start_time,omitempty"`
	EndTime          string      `json:"end_time,omitempty"`
}

type SimpleModifyCallRequest struct {
	Action string `json:"action,omitempty"`
}

type TransferCallRequest struct {
	Action      string              `json:"action,omitempty"`
	Destination TransferDestination `json:"destination,omitempty"`
}

type ModifyCallResponse struct {
	Message string `json:"message,omitempty"`
	UUID    string `json:"uuid,omitempty"`
}

type TransferDestination struct {
	Type string   `json:"type,omitempty"`
	URL  []string `json:"url,omitempty"`
}

type CallErrorResponse struct {
	Type       string `json:"type,omitempty"`
	ErrorTitle string `json:"error_title,omitempty"`
}

type StreamRequest struct {
	StreamURL []string `json:"stream_url"`
	Loop      int64    `json:"loop,omitempty"`
}

type TalkRequest struct {
	Text      string `json:"text"`
	VoiceName string `json:"voice_name,omitempty"`
	Loop      int64  `json:"loop,omitempty"`
}

type DTMFRequest struct {
	Digits string `json:"digits"`
}

type GetBalanceResponse struct {
	Value float64 `json:"value,omitempty"`
}

type GetPhoneOutboundPricingResponse struct {
	Network     string `json:"network,omitempty"`
	Phone       string `json:"phone,omitempty"`
	CountryCode string `json:"country-code,omitempty"`
	Price       string `json:"price,omitempty"`
}

// Reports api
type ReportRecordResponse interface {
	GetAccountId() string
	GetPrice() string
}

type CallRecordResponse struct {
	AccountID         string `json:"account_id"`
	CallID            string `json:"call_id"`
	Direction         string `json:"direction"`
	From              string `json:"from"`
	To                string `json:"to"`
	Network           string `json:"network"`
	NetworkName       string `json:"network_name"`
	Country           string `json:"country"`
	CountryName       string `json:"country_name"`
	DateStart         string `json:"date_start"`
	DateEnd           string `json:"date_end"`
	Duration          string `json:"duration"`
	Status            string `json:"status"`
	StatusDescription string `json:"status_description"`
	Currency          string `json:"currency"`
	Price             string `json:"price"`
	TotalPrice        string `json:"total_price"`
}

func (r *CallRecordResponse) GetAccountId() string { return r.AccountID }
func (r *CallRecordResponse) GetPrice() string     { return r.Price }

type SMSRecordResponse struct {
	AccountID            string `json:"account_id"`
	MessageID            string `json:"message_id"`
	ClientRef            string `json:"client_ref"`
	Direction            string `json:"direction"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	Network              string `json:"network"`
	NetworkName          string `json:"network_name"`
	Country              string `json:"country"`
	CountryName          string `json:"country_name"`
	DateReceived         string `json:"date_received"`
	DateFinalized        string `json:"date_finalized"`
	Latency              string `json:"latency"`
	Status               string `json:"status"`
	ErrorCode            string `json:"error_code"`
	ErrorCodeDescription string `json:"error_code_description"`
	Currency             string `json:"currency"`
	TotalPrice           string `json:"total_price"`
	MessageBody          string `json:"message_body"`
}

func (r *SMSRecordResponse) GetAccountId() string { return r.AccountID }
func (r *SMSRecordResponse) GetPrice() string     { return r.TotalPrice }

type LoadRecordsRequest struct {
	AccountID      string `json:"account_id"`
	ID             string `json:"id"`        // The UUID of the message or call to be searched for.
	Product        string `json:"product"`   // Must be one of: SMS, VOICE-CALL, WEBSOCKET-CALL, VERIFY-API, NUMBER-INSIGHT, MESSAGES, CONVERSATIONS or ASR
	Direction      string `json:"direction"` // Must be one of: inbound or outbound
	DateStart      string `json:"date_start"`
	DateEnd        string `json:"date_end"`
	IncludeMessage bool   `json:"include_message"`
	Type           string `json:"type"`
	Status         string `json:"status"` // Must be one of: delivered, expired, failed, rejected, accepted, buffered, unknown or deleted
}

type LoadRecordsResponse struct {
	Links          Links                  `json:"_links,omitempty"`
	RequestID      string                 `json:"request_id"`
	RequestStatus  string                 `json:"request_status"`
	ReceivedAt     string                 `json:"received_at"`
	Price          string                 `json:"price"`
	Currency       string                 `json:"currency"`
	AccountID      string                 `json:"account_id"`
	Product        string                 `json:"product"`
	Direction      string                 `json:"direction"`
	IncludeMessage string                 `json:"include_message"`
	ItemsCount     int                    `json:"items_count"`
	Records        []ReportRecordResponse `json:"records"`
}

type ReportErrorResponse struct {
	Type       string `json:"type,omitempty"`
	ErrorTitle string `json:"error_title,omitempty"`
}
