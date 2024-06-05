package leads

type Lead struct {
	Id             uint64
	AdId           uint64
	Title          string
	Description    string
	Price          uint
	Status         string
	Responsible    uint64
	CompanyId      uint64
	ClientCompany  uint64
	ClientEmployee uint64
	CreationDate   int64
	IsDeleted      bool
}

type Filter struct {
	Limit  uint
	Offset uint

	Status   string
	ByStatus bool

	Responsible   uint64
	ByResponsible bool
}

type UpdateLead struct {
	Title       string
	Description string
	Price       uint
	Status      string
	Responsible uint64
}
