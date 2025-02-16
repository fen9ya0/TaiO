package response

type Search struct {
	Code     int
	Msg      string
	Hosts    string
	Ports    []string
	Services []Service
	Fingers  []Finger
}

type Service struct {
	Port        string
	ServiceName string
}

type Finger struct {
	Url    string
	App    string
	Server string
}
