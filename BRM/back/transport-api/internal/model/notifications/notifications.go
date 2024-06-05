package notifications

type Notification struct {
	Id        uint64
	CompanyId uint64
	Type      string
	Date      int64
	Viewed    bool

	*NewLead
	*ClosedLead
}

type ClosedLead struct {
	AdId            uint64
	ProducerCompany uint64
	Answered        bool
}

type NewLead struct {
	LeadId        uint64
	ClientCompany uint64
}
