package zabbix

// TriggerObject represents a Zabbix trigger object
type TriggerObject struct {
	TriggerId          string       `json:"triggerid"`
	Description        string       `json:"description"`
	Expression         string       `json:"expression"`
	Comments           string       `json:"comments,omitempty"`
	Error              string       `json:"error,omitempty"`
	Flags              int          `json:"flags,omitempty"`
	LastChange         int64        `json:"lastchange,omitempty"`
	Priority           int          `json:"priority"`
	Status             int          `json:"status"`
	TemplateId         string       `json:"templateid,omitempty"`
	Type               int          `json:"type"`
	Url                string       `json:"url,omitempty"`
	Value              int          `json:"value"`
	State              int          `json:"state"`
	RecoveryMode       int          `json:"recovery_mode,omitempty"`
	RecoveryExpression string       `json:"recovery_expression,omitempty"`
	CorrelationMode    int          `json:"correlation_mode,omitempty"`
	CorrelationTag     string       `json:"correlation_tag,omitempty"`
	ManualClose        int          `json:"manual_close,omitempty"`
	OpData             string       `json:"opdata,omitempty"`
	Hosts              []Host       `json:"hosts,omitempty"`
	Functions          []Function   `json:"functions,omitempty"`
	Dependencies       []Dependency `json:"dependencies,omitempty"`
	Items              []Item       `json:"items,omitempty"`
}

// Host represents a simplified host object for a trigger
type Host struct {
	HostId string `json:"hostid"`
	Host   string `json:"host"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}

// Function represents a simplified function object for a trigger
type Function struct {
	FunctionId string `json:"functionid"`
	ItemId     string `json:"itemid"`
	Function   string `json:"function"`
	Parameter  string `json:"parameter"`
}

// Dependency represents a simplified dependency object for a trigger
type Dependency struct {
	TriggerId   string `json:"triggerid"`
	Description string `json:"description"`
}

// Item represents a simplified item object for a trigger
type Item struct {
	ItemId    string `json:"itemid"`
	HostId    string `json:"hostid"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	LastValue string `json:"lastvalue"`
	LastClock int64  `json:"lastclock"`
	State     int    `json:"state"`
	Status    int    `json:"status"`
}

// TriggerGetParameters struct is used for item get requests
//
// see: https://www.zabbix.com/documentation/5.0/en/manual/api/reference/trigger/get#Parameters
type TriggerGetParameters struct {
	GetParameters
	Active             bool     `json:"active,omitempty"`
	ApplicationIds     []string `json:"applicationids,omitempty"`
	Description        string   `json:"description,omitempty"`
	ExpandComment      bool     `json:"expandComment,omitempty"`
	ExpandExpression   bool     `json:"expandExpression,omitempty"`
	ExpandDescription  bool     `json:"expandDescription,omitempty"`
	ExpandRecovery     bool     `json:"expandRecovery,omitempty"`
	ExpandTrigger      bool     `json:"expandTrigger,omitempty"`
	GroupIds           []string `json:"groupids,omitempty"`
	HostIds            []string `json:"hostids,omitempty"`
	Inherited          bool     `json:"inherited,omitempty"`
	ItemIds            []string `json:"itemids,omitempty"`
	LastChangeSince    int64    `json:"lastChangeSince,omitempty"`
	LastChangeTill     int64    `json:"lastChangeTill,omitempty"`
	Maintenance        bool     `json:"maintenance,omitempty"`
	MinSeverity        int      `json:"min_severity,omitempty"`
	Monitored          bool     `json:"monitored,omitempty"`
	OnlyTrue           bool     `json:"only_true,omitempty"`
	Output             []string `json:"output,omitempty"` // Output might differ, could be SelectQuery
	Recovery           bool     `json:"recovery,omitempty"`
	SkipDependent      bool     `json:"skipDependent,omitempty"`
	TemplateIds        []string `json:"templateids,omitempty"`
	TriggerIds         []string `json:"triggerids,omitempty"`
	WithAcknowledges   bool     `json:"withAcknowledges,omitempty"`
	WithLastEventUnack bool     `json:"withLastEventUnack,omitempty"`
	WithRecovery       bool     `json:"withRecovery,omitempty"`
	WithUnacknowledged bool     `json:"withUnacknowledged,omitempty"`
	// Add other specific parameters if any
}

func (z *Context) TriggerGet(params TriggerGetParameters) ([]TriggerObject, int, error) {
	var result []TriggerObject

	status, err := z.request("trigger.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}
