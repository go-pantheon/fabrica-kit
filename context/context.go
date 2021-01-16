package context

import (
	"context"
)

// Context is the context of the game
// Use the custom type for your constants
const (
	CtxSID         = "x-md-global-sid"
	CtxUID         = "x-md-global-uid"    
	CtxOID         = "x-md-global-oid"    
	CtxColor       = "x-md-global-color"  
	CtxStatus      = "x-md-global-status" 
	CtxReferer     = "x-md-global-referer"
	CtxClientIP    = "x-md-global-client-ip"
	CtxGateReferer = "x-md-global-gate-referer"
)

var Keys = []string{CtxSID, CtxUID, CtxOID, CtxStatus, CtxColor, CtxReferer, CtxClientIP, CtxGateReferer}
