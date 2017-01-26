// Package tethering provides the Chrome Debugging Protocol
// commands, types, and events for the Chrome Tethering domain.
//
// The Tethering domain defines methods and events for browser port binding.
//
// Generated by the chromedp-gen command.
package tethering

// AUTOGENERATED. DO NOT EDIT.

import (
	"context"

	cdp "github.com/knq/chromedp/cdp"
	"github.com/mailru/easyjson"
)

// BindParams request browser port binding.
type BindParams struct {
	Port int64 `json:"port"` // Port number to bind.
}

// Bind request browser port binding.
//
// parameters:
//   port - Port number to bind.
func Bind(port int64) *BindParams {
	return &BindParams{
		Port: port,
	}
}

// Do executes Tethering.bind.
func (p *BindParams) Do(ctxt context.Context, h cdp.FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandTetheringBind, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return cdp.ErrContextDone
	}

	return cdp.ErrUnknownResult
}

// UnbindParams request browser port unbinding.
type UnbindParams struct {
	Port int64 `json:"port"` // Port number to unbind.
}

// Unbind request browser port unbinding.
//
// parameters:
//   port - Port number to unbind.
func Unbind(port int64) *UnbindParams {
	return &UnbindParams{
		Port: port,
	}
}

// Do executes Tethering.unbind.
func (p *UnbindParams) Do(ctxt context.Context, h cdp.FrameHandler) (err error) {
	if ctxt == nil {
		ctxt = context.Background()
	}

	// marshal
	buf, err := easyjson.Marshal(p)
	if err != nil {
		return err
	}

	// execute
	ch := h.Execute(ctxt, cdp.CommandTetheringUnbind, easyjson.RawMessage(buf))

	// read response
	select {
	case res := <-ch:
		if res == nil {
			return cdp.ErrChannelClosed
		}

		switch v := res.(type) {
		case easyjson.RawMessage:
			return nil

		case error:
			return v
		}

	case <-ctxt.Done():
		return cdp.ErrContextDone
	}

	return cdp.ErrUnknownResult
}
