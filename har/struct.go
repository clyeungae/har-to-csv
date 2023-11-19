package har

type Har struct {
	Log Log
}

type Log struct {
	Version string
	Creator Creator
	Page    string
	Entries []Entry
}

type Creator struct {
	Name    string
	Version string
}

type Entry struct {
	Initiator       Initiator `json:"_initiator"`
	Priority        string    `json:"_priority"`
	ResourceType    string    `json:"_resourceType"`
	Cache           Cache
	Connection      string
	Request         Request
	Response        Response
	ServerIPAddress string
	StartedDateTime string
}

type Cache struct {
}

type Initiator struct {
	Type  string
	Stack Stack
}

type Stack struct {
	CallFrames []CallFrame
	Parent     Parent
}

type CallFrame struct {
	FunctionName string
	ScriptId     string
	Url          string
	LineNumber   int
	ColumnNumber int
}

type Parent struct {
	Description string
	CallFrames  []CallFrame
	Parent      *Parent
}

type Request struct {
	Method      string
	URL         string
	HttpVersion string
	Headers     []Header
	QueryString []QueryString
	Cookies     []Cookie
	HeaderSize  int
	BodySize    int
}

type QueryString struct {
}

type Header struct {
	Name  string
	Value string
}

type Cookie struct {
	Name     string
	Value    string
	Path     string
	Domain   string
	Expires  string
	HttpOnly bool
	Secure   bool
}

type Response struct {
	Status       int
	StatusText   string
	HttpVersion  string
	Headers      []Header
	Cookies      []Cookie
	Content      Content
	RedirectURL  string
	HeaderSize   int
	BodySize     int
	TransferSize int    `json:"_transferSize"`
	Error        string `json:"_error"`
}

type Content struct {
	Size        int
	MimeType    string
	Compression int
	Text        string
}
