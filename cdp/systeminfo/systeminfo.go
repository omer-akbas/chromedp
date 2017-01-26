// Package systeminfo provides the Chrome Debugging Protocol
// commands, types, and events for the Chrome SystemInfo domain.
//
// The SystemInfo domain defines methods and events for querying low-level
// system information.
//
// Generated by the chromedp-gen command.
package systeminfo

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
)

// GetInfoParams returns information about the system.
type GetInfoParams struct{}

// GetInfo returns information about the system.
func GetInfo() *GetInfoParams {
	return &GetInfoParams{}
}

// GetInfoReturns return values.
type GetInfoReturns struct {
	Gpu          *GPUInfo `json:"gpu,omitempty"`          // Information about the GPUs on the system.
	ModelName    string   `json:"modelName,omitempty"`    // A platform-dependent description of the model of the machine. On Mac OS, this is, for example, 'MacBookPro'. Will be the empty string if not supported.
	ModelVersion string   `json:"modelVersion,omitempty"` // A platform-dependent description of the version of the machine. On Mac OS, this is, for example, '10.1'. Will be the empty string if not supported.
}

// Do executes SystemInfo.getInfo.
//
// returns:
//   gpu - Information about the GPUs on the system.
//   modelName - A platform-dependent description of the model of the machine. On Mac OS, this is, for example, 'MacBookPro'. Will be the empty string if not supported.
//   modelVersion - A platform-dependent description of the version of the machine. On Mac OS, this is, for example, '10.1'. Will be the empty string if not supported.
func (p *GetInfoParams) Do(ctxt context.Context, h cdp.FrameHandler) (gpu *GPUInfo, modelName string, modelVersion string, err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandSystemInfoGetInfo, cdp.Empty)

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return nil, "", "", cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			// unmarshal
			var r GetInfoReturns
			err = easyjson.Unmarshal(v, &r)
			if err != nil {
				return nil, "", "", cdp.ErrInvalidResult
			}

			return r.Gpu, r.ModelName, r.ModelVersion, nil

		case error:
			return nil, "", "", v
		}

	case <-ctxt.Done():
		return nil, "", "", cdp.ErrContextDone
	}

	return nil, "", "", cdp.ErrUnknownResult
}
