package zabbix

// ItemObject struct is used to store item operations result
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/item/object
type ItemObject struct {
	ItemID               string              `json:"itemid,omitempty"`
	Type                 int                 `json:"type"`
	SNMPCommunity        string              `json:"snmp_community,omitempty"`
	SNMPOID              string              `json:"snmp_oid,omitempty"`
	HostID               string              `json:"hostid"`
	Name                 string              `json:"name"`
	Key                  string              `json:"key_"`
	Delay                string              `json:"delay"`
	History              string              `json:"history,omitempty"`
	Trends               string              `json:"trends,omitempty"`
	Status               int                 `json:"status"`
	ValueType            int                 `json:"value_type"`
	AllowedHosts         string              `json:"allowed_hosts,omitempty"`
	Units                string              `json:"units,omitempty"`
	Delta                int                 `json:"delta"`
	Snmpv3Contextname    string              `json:"snmpv3_contextname,omitempty"`
	Snmpv3SecurityName   string              `json:"snmpv3_securityname,omitempty"`
	Snmpv3SecurityLevel  int                 `json:"snmpv3_securitylevel,omitempty"`
	Snmpv3AuthPassphrase string              `json:"snmpv3_authpassphrase,omitempty"`
	Snmpv3PrivPassphrase string              `json:"snmpv3_privpassphrase,omitempty"`
	Params               string              `json:"params,omitempty"`
	IPMIAuthType         int                 `json:"ipmi_auth_type,omitempty"`
	IPMIPrivilege        int                 `json:"ipmi_privilege,omitempty"`
	IPMIUsername         string              `json:"ipmi_username,omitempty"`
	IPMIPassword         string              `json:"ipmi_password,omitempty"`
	DelayFlex            string              `json:"delay_flex,omitempty"`
	ParamsEscaped        string              `json:"params_escaped,omitempty"`
	InventoryLink        int                 `json:"inventory_link,omitempty"`
	Description          string              `json:"description,omitempty"`
	Lifetime             string              `json:"lifetime,omitempty"`
	Filter               string              `json:"filter,omitempty"`
	DataType             int                 `json:"data_type,omitempty"`
	ValueMapID           string              `json:"valuemapid,omitempty"`
	LogTimeFormat        string              `json:"logtimefmt,omitempty"`
	JmxEndpoint          string              `json:"jmx_endpoint,omitempty"`
	MasterItemID         string              `json:"master_itemid,omitempty"`
	Timeout              string              `json:"timeout,omitempty"`
	URL                  string              `json:"url,omitempty"`
	QueryFields          []QueryField        `json:"query_fields,omitempty"`
	Posts                string              `json:"posts,omitempty"`
	StatusCodes          string              `json:"status_codes,omitempty"`
	FollowRedirects      int                 `json:"follow_redirects,omitempty"`
	PostType             int                 `json:"post_type,omitempty"`
	HTTPProxy            string              `json:"http_proxy,omitempty"`
	Headers              []Header            `json:"headers,omitempty"`
	RetrieveMode         int                 `json:"retrieve_mode,omitempty"`
	RequestMethod        string              `json:"request_method,omitempty"`
	OutputFormat         string              `json:"output_format,omitempty"`
	SSLVerifyPeer        int                 `json:"ssl_verify_peer,omitempty"`
	SSLVerifyHost        int                 `json:"ssl_verify_host,omitempty"`
	VerifyHost           int                 `json:"verify_host,omitempty"`
	SSLKeyFile           string              `json:"ssl_key_file,omitempty"`
	SSLCertFile          string              `json:"ssl_cert_file,omitempty"`
	SSLCAFile            string              `json:"ssl_ca_file,omitempty"`
	SNMPTrapType         int                 `json:"snmptrap_type,omitempty"`
	SNMPTrapPort         int                 `json:"snmptrap_port,omitempty"`
	SNMPCommunity2       string              `json:"snmp_community2,omitempty"`
	SNMPTrapOid          string              `json:"snmptrap_oid,omitempty"`
	SNMPTrapComm2        string              `json:"snmptrap_comm2,omitempty"`
	SNMPTrapOID2         string              `json:"snmptrap_oid2,omitempty"`
	SNMPTrapV1Comm2      string              `json:"snmptrap_v1_comm2,omitempty"`
	SNMPTrapV1OID2       string              `json:"snmptrap_v1_oid2,omitempty"`
	SNMPTrapV1Comm1      string              `json:"snmptrap_v1_comm1,omitempty"`
	SNMPTrapV1OID1       string              `json:"snmptrap_v1_oid1,omitempty"`
	PreprocessingSteps   []PreprocessingStep `json:"preprocessing_steps,omitempty"`
	IsItemPrototype      bool                `json:"is_item_prototype,omitempty"`
}

// QueryField represents a query field object.
type QueryField struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Header represents an HTTP header object.
type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// PreprocessingStep represents a preprocessing step object.
type PreprocessingStep struct {
	Type               int    `json:"type"`
	Params             string `json:"params"`
	ErrorHandler       int    `json:"error_handler"`
	ErrorHandlerParams string `json:"error_handler_params"`
}

// ItemGetParams struct is used for item get requests
//
// see: https://www.zabbix.com/documentation/5.0/manual/api/reference/item/get#parameters
type ItemGetParameters struct {
	GetParameters
	ItemIDs             []string    `json:"itemids,omitempty"`
	GroupIDs            []string    `json:"groupids,omitempty"`
	TemplateIDs         []string    `json:"templateids,omitempty"`
	HostIDs             []string    `json:"hostids,omitempty"`
	ProxyIDs            []string    `json:"proxyids,omitempty"`
	InterfaceIDs        []string    `json:"interfaceids,omitempty"`
	GraphIDs            []string    `json:"graphids,omitempty"`
	TriggerIDs          []string    `json:"triggerids,omitempty"`
	ApplicationIDs      []string    `json:"applicationids,omitempty"`
	WebItems            bool        `json:"webitems,omitempty"`
	Inherited           bool        `json:"inherited,omitempty"`
	Templated           bool        `json:"templated,omitempty"`
	Monitored           bool        `json:"monitored,omitempty"`
	Group               string      `json:"group,omitempty"`
	Host                string      `json:"host,omitempty"`
	Application         string      `json:"application,omitempty"`
	WithTriggers        bool        `json:"with_triggers,omitempty"`
	SelectHosts         SelectQuery `json:"selectHosts,omitempty"`
	SelectInterfaces    SelectQuery `json:"selectInterfaces,omitempty"`
	SelectTriggers      SelectQuery `json:"selectTriggers,omitempty"`
	SelectGraphs        SelectQuery `json:"selectGraphs,omitempty"`
	SelectApplications  SelectQuery `json:"selectApplications,omitempty"`
	SelectDiscoveryRule SelectQuery `json:"selectDiscoveryRule,omitempty"`
	SelectItemDiscovery SelectQuery `json:"selectItemDiscovery,omitempty"`
	SelectPreprocessing SelectQuery `json:"selectPreprocessing,omitempty"`
	LimitSelects        int         `json:"limitSelects,omitempty"`
}

func (z *Context) ItemGet(params ItemGetParameters) ([]ItemObject, int, error) {
	var result []ItemObject

	status, err := z.request("item.get", params, &result)
	if err != nil {
		return nil, status, err
	}

	return result, status, nil
}
