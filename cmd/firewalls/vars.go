package cmd

// create
var (
	createFirewallName        string
	createFirewallStatus      string
	createFirewallDescription string
	createFirewallRules       string
)

// list
var (
	listFirewallNameFilter      string
	listFirewallInstancesFilter string
	listFirewallIdFilter        string
	listFirewallStatusFilter    string
)

// get
var (
	getFirewallId string
)

// get-rules
var (
	getFirewallRulesId string
)

// get-instances
var getFirewallInstancesId string

// delete
var (
	deleteFirewallId string
)

// get
var (
	setFirewallId string
)

// edit
var (
	editFirewallId string
)

// update
var (
	updateFirewallId          string
	updateFirewallName        string
	updateFirewallStatus      string
	updateFirewallDescription string
)

// add rules

var (
	addRulesFirewallId    string
	addRulesFirewallRules string
)

// assign
var (
	assignFirewallId string
	assignInstanceId int64
)

// unassign
var (
	unassignFirewallId string
	unassignInstanceId int64
)

// history
var (
	historyfirewallIdFilter string
)
