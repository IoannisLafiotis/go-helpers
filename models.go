package helpersforgolang

import (
	"net/http"
	"time"
)

type Request struct {
	URL            string            `json:"url"`
	Method         string            `json:"method"`
	Headers        map[string]string `json:"headers"`
	RequestBody    interface{}       `json:"requestBody"`
	Authentication string            `json:"authentication"`
	Username       string            `json:"username"`
	Password       string            `json:"password"`
	OAuth2Token    string            `json:"oauth2Token"`
	APIToken       string            `json:"apiToken"`
	QueryParams    map[string]string `json:"queryParams"`
	PathParams     map[string]string `json:"pathParams"`
}
type CallParams struct {
	pathName           string
	method             string
	requestContentType string
}

// Response represents the response received from the API
type Response struct {
	Status  int         `json:"status"`
	Headers http.Header `json:"headers"`
	Body    []byte      `json:"body"`
}

type Flow struct {
	FlowNode []map[Component]interface{}
}

/*
type Component struct {
	Id                 string      `json:"id"`
	VirtualComponentId string      `json:"virtualComponentId"`
	AuthType           string      `json:"string"`
	ComponentId        string      `json:"componentId"`
	Cfg                interface{} `json:"cfg"`
	Msg                interface{} `json:"msg"`
}
*/

type Element struct {
	Data map[any]interface{}
}

type Snapshot struct {
	LastUpdated time.Time
}

type NodeSettingsI struct {
	SnapshotKey       string
	ArraySplittingKey string
	SyncParam         string
	SkipSnapshot      string
}

type Cfg struct {
	NodeSettings  NodeSettingsI
	Server        string
	OtherServer   string
	TriggerParams map[string]interface{}
}

type Component struct {
	Name       string `json:"name"`
	ID         string `json:"id"`
	TypeOfCall string `json:"typeOfCall"`
}

type FlowData struct {
	Components []Component `json:"components"`
}

func main() {}
