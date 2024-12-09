package types

type DeviceDetail struct {
	UserID          string
	DeviceName      string
	DevicePassword  string
	ClientAdded     bool
	Online          bool
	ScheduledAction bool
	Command         string
	Timer           string
	OnlineSince     uint
}