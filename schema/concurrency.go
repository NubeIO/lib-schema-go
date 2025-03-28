package schema

type EnableConcurrency struct {
	Type     string `json:"type" default:"boolean"`
	Title    string `json:"title" default:"Enable Concurrency"`
	Default  bool   `json:"default" default:"false"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}

type ConcurrencyLimit struct {
	Type     string `json:"type" default:"number"`
	Title    string `json:"title" default:"Concurrency Limit"`
	Default  int    `json:"default" default:"0"`
	ReadOnly bool   `json:"readOnly" default:"false"`
}
