package order

type Status string

const (
	Pending   Status = "PENDING"
	Confirmed Status = "CONFIRMED"
	Cancelled Status = "CANCELLED"
	Completed Status = "COMPLETED"
)

func (s Status) String() string {
	return string(s)
}
