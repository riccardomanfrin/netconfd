// Copyright (c) 2021, Athonet S.r.l. All rights reserved.
// riccardo.manfrin@athonet.com

package nc

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"runtime/debug"
	"syscall"
)

//ErrorCode describes the error type via enumeration
type ErrorCode int

const (
	//CONFLICT error type (inconsistency with respect to the existing state)
	CONFLICT ErrorCode = iota
	//NOT_FOUND error types encodes a restful resource not found by its ID
	NOT_FOUND
	//SEMANTIC error type of the requested operation in the syntax or logical content
	SEMANTIC
	//SYNTAX error type is for synctactical errors
	SYNTAX
	//UNKNOWN_TYPE error type (the value type is not recognized/supported)
	UNKNOWN_TYPE
	//UNEXPECTED_CORNER_CASE error type describes an error that was not meant to appear
	UNEXPECTED_CORNER_CASE
	//UNSUPPORTED error type describes an error where a part of the implementation is missing
	UNSUPPORTED
	//RESERVED can be used for outer error enum cohexistence
	RESERVED = 1000
)

var errorCodeToString = map[ErrorCode]string{
	CONFLICT:     "Conflict Error",
	SEMANTIC:     "Semantic Error",
	SYNTAX:       "Syntax Error",
	UNKNOWN_TYPE: "UnknownType Error",
}
var NetconfdDebugTrace = false

//InitErrorsLogsTracing Initializes error logs with tracing
func InitErrorsLogsTracing() {
	netconfdDebugTraceStr, netconfdDebugTraceFound := os.LookupEnv("APP_DEBUG_TRACE")
	if !netconfdDebugTraceFound {
		NetconfdDebugTrace = false
	}
	if netconfdDebugTraceStr == "true" {
		NetconfdDebugTrace = true
	}
}

//Can also just NOT be an error!
func mapNetlinkError(err error, r Resource) error {
	if err != nil {
		switch err.(type) {
		case syscall.Errno:
			if err.(syscall.Errno) == syscall.EINVAL {
				return NewEINVALError()
			} else if err.(syscall.Errno) == syscall.EPERM {
				return NewEPERMError(r)
			} else if err.(syscall.Errno) == syscall.ENETUNREACH {
				return NewENETUNREACHError(r)
			} else if err.(syscall.Errno) == syscall.EEXIST {
				return NewEEXISTError(r)
			} else if err.(syscall.Errno) == syscall.ERANGE {
				return NewERANGEError(r)
			} else if err.(syscall.Errno) == syscall.EACCES {
				return NewEACCESError(r)
			}
		}
		debug.PrintStack()
		return NewGenericError(err)
	}
	return err
}

//GenericError describes a generic error of the library
type GenericError struct {
	//code error type
	Code ErrorCode `json:"code"`
	//reason describes the specific reason for the error
	Reason string `json:"reason"`
}

func (e *GenericError) Error() string {
	strerr, _ := json.Marshal(*e)
	return string(strerr)
}

//NewGenericError returns a generic error
func NewGenericError(err error) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &GenericError{Code: UNKNOWN_TYPE, Reason: fmt.Sprintf("Generic uncharted error: %v", err.Error())}
}

//NewGenericErrorWithReason returns a generic semantic error
func NewGenericErrorWithReason(reason string) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &GenericError{Code: UNKNOWN_TYPE, Reason: reason}
}

//SemanticError is a logical error on the content of the operation requested to be performed
type SemanticError GenericError

func (e *SemanticError) Error() string {
	strerr, _ := json.Marshal(*e)
	return string(strerr)
}

//NewGenericSemanticError returns a generic semantic error
func NewGenericSemanticError() error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: "Generic Semantic Error"}
}

//NewUnknownLinkKindError returns a SemanticError error on link layer type interfaces
func NewUnknownLinkKindError(linkKind string) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: "LinkKind " + string(linkKind) + " not known"}
}

//NewUnsupportedSlaveKindError returns a SemanticError error un unsupported kind of device enslavement
func NewUnsupportedSlaveKindError(infoSlaveKind string) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: fmt.Sprintf("Unsupported/Unknown Slave kind %v", infoSlaveKind)}
}

//NewBadAddressError returns a bad address error on link layer interfaces
func NewBadAddressError(c CIDRAddr) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: "Bad IP address " + c.String()}
}

//NewInvalidPrefixLenForIPv4AddrError returns a bad address error on link layer interfaces
func NewInvalidPrefixLenForIPv4AddrError(len int) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: fmt.Sprintf("Invalid network Prefix length %v for IPv4 address", len)}
}

//NewInvalidPrefixLenForIPv6AddrError returns a bad address error on link layer interfaces
func NewInvalidPrefixLenForIPv6AddrError(len int) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: fmt.Sprintf("Invalid network Prefix length %v for IPv6 address", len)}
}

//NewInvalidUnmanagedResourceTypeError returns a bad address error on link layer interfaces
func NewInvalidUnmanagedResourceTypeError(t Type) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: fmt.Sprintf("Invalid unmanaged resource type %v", t)}
}

//NewEINVALError returns a bad address error on link layer interfaces
func NewEINVALError() error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: "Got EINVAL error: (check dmesg)"}
}

//NewActiveSlaveIfaceNotFoundForActiveBackupBondError Returns an error if an active interface is not found for an Active-Backup type bond
func NewActiveSlaveIfaceNotFoundForActiveBackupBondError(bondIfname LinkID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: "Active Slave Iface not found for Active-Backup type bond " + string(bondIfname)}
}

//NewParentLinkNotFoundForVlan returns a Not found error on link layer interfaces
func NewParentLinkNotFoundForVlan(ifname LinkID, parentIfname LinkID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: "Parent Link " + string(parentIfname) + " not found for Vlan Link " + string(ifname)}
}

//NewMultipleActiveSlaveIfacesFoundForActiveBackupBondError Returns an error if an active interface is not found for an Active-Backup type bond
func NewMultipleActiveSlaveIfacesFoundForActiveBackupBondError(bondIfname LinkID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: "Multiple Active Slave Ifaces found for Active-Backup type bond " + string(bondIfname)}
}

//NewBackupSlaveIfaceFoundForNonActiveBackupBondError Returns an error if a backup interface is found for a non Active-Backup type bond
func NewBackupSlaveIfaceFoundForNonActiveBackupBondError(backupIfname LinkID, bondIfname LinkID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: "Backup Slave Iface " + string(backupIfname) + " found for non Active-Backup type bond " + string(bondIfname)}
}

//NewRouteLinkDeviceNotFoundError describes a link device not found for a route to create
func NewRouteLinkDeviceNotFoundError(routeID RouteID, linkID LinkID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: "Route " + string(routeID) + " Link Device " + string(linkID) + " not found"}
}

//NewENETUNREACHError returns a network unreachable error
func NewENETUNREACHError(r Resource) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: fmt.Sprintf("Got ENETUNREACH error: network is not reachable for route %v", r.Print())}
}

//NewEEXISTError returns a conflict error
func NewEEXISTError(r Resource) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: fmt.Sprintf("Got EEXIST error: route exists %v", r.Print())}
}

//NewERANGEError returns an out of range error
func NewERANGEError(r Resource) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: fmt.Sprintf("Got ERANGE error: parameter value out of range for resource %+v", r.Print())}
}

//NewEACCESError returns an out of range error
func NewEACCESError(r Resource) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: fmt.Sprintf("Got EACCES error: incompatible param in resource %+v (check dmsg)", r.Print())}
}

//NewTooManyDNSServersError describes an error on the number of requested DNS servers
func NewTooManyDNSServersError() error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: "More than two entries found in DNS servers config"}
}

//NewDuplicateDNSServersIDsError describes a config error on the DNS serves IDs
func NewDuplicateDNSServersIDsError(dnsid1 DnsID, dnsid2 DnsID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: fmt.Sprintf("Duplicate DNS server IDs %v, %v config", string(dnsid1), string(dnsid2))}
}

//NewUnknownUnsupportedDNSServersIDsError describes a config error on the DNS serves IDs
func NewUnknownUnsupportedDNSServersIDsError(dnsid1 DnsID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SemanticError{Code: SEMANTIC, Reason: fmt.Sprintf("Unknown/Unsupported DNS server IDs %v", string(dnsid1))}
}

//SyntaxError is a logical error on the content of the operation requested to be performed
type SyntaxError GenericError

func (e *SyntaxError) Error() string {
	strerr, _ := json.Marshal(*e)
	return string(strerr)
}

//NewInvalidIPAddressError Returns an error if a backup interface is found for a non Active-Backup type bond
func NewInvalidIPAddressError(addr string) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &SyntaxError{Code: SYNTAX, Reason: "Invalid IP Address/Network  " + addr}
}

//UnknownTypeError is a logical error on the content of the operation requested to be performed
type UnknownTypeError GenericError

func (e *UnknownTypeError) Error() string {
	strerr, _ := json.Marshal(*e)
	return string(strerr)
}

//NewLinkUnknownFlagTypeError returns a Conflict error on link layer interfaces
func NewLinkUnknownFlagTypeError(flag LinkFlag) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &UnknownTypeError{Code: UNKNOWN_TYPE, Reason: "Link Flag Type" + string(flag) + " unknown/unsupported"}
}

//ConflictError describes a conflict with the network state and requested changes
type ConflictError GenericError

func (e *ConflictError) Error() string {
	strerr, _ := json.Marshal(*e)
	return string(strerr)
}

//NewLinkExistsConflictError returns a Conflict error on link layer interfaces
func NewLinkExistsConflictError(linkID LinkID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: "Link " + string(linkID) + " exists"}
}

//NewLinkDeviceDoesNotExistError returns a Conflict error on link layer interfaces
func NewLinkDeviceDoesNotExistError(linkID LinkID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: "Link device " + string(linkID) + " does not exist"}
}

//NewNonBondMasterLinkTypeError returns an error for non bond master link type
func NewNonBondMasterLinkTypeError(ifname LinkID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: "Master link interface " + string(ifname) + " is not a bond"}
}

//NewCannotStopDHCPError returns an error for DHCP related stop errors
func NewCannotStopDHCPError(ifname LinkID, e error) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: fmt.Sprintf("Failed to stop DHCP for interface %v: %v", string(ifname), e)}
}

//NewUnmanagedLinkCannotBeModifiedError returns an error for unmanaged links which are requested to be configured
func NewUnmanagedLinkCannotBeModifiedError(ifname LinkID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: fmt.Sprintf("Cannot modify unmanaged interface %v", string(ifname))}
}

//NewUnmanagedLinkRouteCannotBeModifiedError returns an error for unmanaged links routes which are requested to be configured
func NewUnmanagedLinkRouteCannotBeModifiedError(r Route) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: fmt.Sprintf("Cannot modify unmanaged link %v route %v", r.Dev, r.Print())}
}

//NewUnmanagedLinkRuleCannotBeModifiedError returns an error for unmanaged links rules which are requested to be configured
func NewUnmanagedLinkRuleCannotBeModifiedError(r Rule) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: fmt.Sprintf("Cannot modify unmanaged link %v/%v rule %v", r.IifName, r.OifName, r.Print())}
}

//NewUnmanagedLinkDHCPCannotBeModifiedError returns an error for unmanaged links which are requested to be configured
func NewUnmanagedLinkDHCPCannotBeModifiedError(ifname LinkID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: fmt.Sprintf("Cannot modify unmanaged link %v DHCP config")}
}

//NewCannotStartDHCPError returns an error for DHCP related stop errors
func NewCannotStartDHCPError(ifname LinkID, e error) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: fmt.Sprintf("Failed to start DHCP for interface %v: %v", string(ifname), e)}
}

//NewCannotStatusDHCPError returns an error for DHCP related status errors
func NewCannotStatusDHCPError(ifname LinkID, e error) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: fmt.Sprintf("Failed to get DHCP status for interface %v: %v", string(ifname), e)}
}

//NewDHCPAlreadyRunningConflictError returns an error for DHCP that is requested for an interface where it's already running
func NewDHCPAlreadyRunningConflictError(ifname LinkID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: fmt.Sprintf("DHCP is alreay running for interface %v", string(ifname))}
}

//NewDNSServerExistsConflictError returns an error for DHCP that is requested for an interface where it's already running
func NewDNSServerExistsConflictError(dnsid DnsID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: fmt.Sprintf("DNS server %v already exists", string(dnsid))}
}

//NewUnknownLinkDeviceLabel
func NewUnknownLinkDeviceLabel(label string) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: fmt.Sprintf("Unknown Link Device Label %v in renaming procedure", label)}
}

//NewEPERMError returns a missing permissions error
func NewEPERMError(r Resource) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{
		Code: CONFLICT,
		Reason: fmt.Sprintf("Got EPERM error: insufficient permissions to perform action on %v: %v",
			reflect.TypeOf(r),
			r.Print())}
}

//NewRouteExistsConflictError returns a Conflict error on link layer interfaces
func NewRouteExistsConflictError(routeID RouteID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: "Route " + string(routeID) + " exists"}
}

//NewruleExistsConflictError returns a Conflict error on link layer interfaces
func NewRuleExistsConflictError(ruleID RuleID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &ConflictError{Code: CONFLICT, Reason: "Rule " + string(ruleID) + " exists"}
}

//NotFoundError is a logical error on the content of the operation requested to be performed
type NotFoundError ConflictError

func (e *NotFoundError) Error() string {
	strerr, _ := json.Marshal(*e)
	return string(strerr)
}

//NewLinkNotFoundError returns a Not found error on link layer interfaces
func NewLinkNotFoundError(linkID LinkID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &NotFoundError{Code: NOT_FOUND, Reason: "Link " + string(linkID) + " not found"}
}

//NewRouteByIDNotFoundError returns a Not found error on link layer interfaces
func NewRouteByIDNotFoundError(routeid RouteID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &NotFoundError{Code: NOT_FOUND, Reason: "Route ID " + string(routeid) + " not found"}
}

//NewDHCPRunningNotFoundError returns a Not found error on link layer interfaces not managed by DHCP
func NewDHCPRunningNotFoundError(linkID LinkID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &NotFoundError{Code: NOT_FOUND, Reason: "DHCP for Link ID " + string(linkID) + " not found"}
}

//NewDNSServerNotFoundError returns a Not found error on DNS not found by ID
func NewDNSServerNotFoundError(dnsID DnsID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &NotFoundError{Code: NOT_FOUND, Reason: "DNS ID " + string(dnsID) + " not found"}
}

//NewUnmanagedResourceNotFoundError returns a Not found error on unmanaged resource not found
func NewUnmanagedResourceNotFoundError(id UnmanagedID) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &NotFoundError{Code: NOT_FOUND, Reason: "Unmanaged Resource ID " + string(id) + " not found"}
}

//UnexpetecdCornerCaseError is fundamentally an implementation error catch exception
//It makes explitic to developer that he did not think of a case that instead happened
type UnexpectedCornerCaseError GenericError

func (e *UnexpectedCornerCaseError) Error() string {
	strerr, _ := json.Marshal(*e)
	return string(strerr)
}

//NewUnexpectedCornerCaseError returns a Conflict error on link layer interfaces
func NewUnexpectedCornerCaseError(reason string) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &UnexpectedCornerCaseError{Code: UNEXPECTED_CORNER_CASE, Reason: reason}
}

//Unsupported describes an error about a part of implementation which is missing
type UnsupportedError GenericError

func (e *UnsupportedError) Error() string {
	strerr, _ := json.Marshal(*e)
	return string(strerr)
}

//NewUnsupportedError returns a Conflict error on link layer interfaces
func NewUnsupportedError(reason string) error {
	if NetconfdDebugTrace {
		debug.PrintStack()
	}
	return &UnsupportedError{Code: UNSUPPORTED, Reason: reason}
}
