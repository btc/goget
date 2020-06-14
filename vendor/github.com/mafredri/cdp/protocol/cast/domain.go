// Code generated by cdpgen. DO NOT EDIT.

// Package cast implements the Cast domain. A domain for interacting with
// Cast, Presentation API, and Remote Playback API functionalities.
package cast

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Cast domain. A domain for interacting with
// Cast, Presentation API, and Remote Playback API functionalities.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Cast domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// Enable invokes the Cast method. Starts observing for sinks that can be used
// for tab mirroring, and if set, sinks compatible with |presentationUrl| as
// well. When sinks are found, a |sinksUpdated| event is fired. Also starts
// observing for issue messages. When an issue is added or removed, an
// |issueUpdated| event is fired.
func (d *domainClient) Enable(ctx context.Context, args *EnableArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Cast.enable", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Cast.enable", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Cast", Op: "Enable", Err: err}
	}
	return
}

// Disable invokes the Cast method. Stops observing for sinks and issues.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Cast.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Cast", Op: "Disable", Err: err}
	}
	return
}

// SetSinkToUse invokes the Cast method. Sets a sink to be used when the web
// page requests the browser to choose a sink via Presentation API, Remote
// Playback API, or Cast SDK.
func (d *domainClient) SetSinkToUse(ctx context.Context, args *SetSinkToUseArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Cast.setSinkToUse", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Cast.setSinkToUse", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Cast", Op: "SetSinkToUse", Err: err}
	}
	return
}

// StartTabMirroring invokes the Cast method. Starts mirroring the tab to the
// sink.
func (d *domainClient) StartTabMirroring(ctx context.Context, args *StartTabMirroringArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Cast.startTabMirroring", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Cast.startTabMirroring", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Cast", Op: "StartTabMirroring", Err: err}
	}
	return
}

// StopCasting invokes the Cast method. Stops the active Cast session on the
// sink.
func (d *domainClient) StopCasting(ctx context.Context, args *StopCastingArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Cast.stopCasting", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Cast.stopCasting", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Cast", Op: "StopCasting", Err: err}
	}
	return
}

func (d *domainClient) SinksUpdated(ctx context.Context) (SinksUpdatedClient, error) {
	s, err := rpcc.NewStream(ctx, "Cast.sinksUpdated", d.conn)
	if err != nil {
		return nil, err
	}
	return &sinksUpdatedClient{Stream: s}, nil
}

type sinksUpdatedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *sinksUpdatedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *sinksUpdatedClient) Recv() (*SinksUpdatedReply, error) {
	event := new(SinksUpdatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Cast", Op: "SinksUpdated Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) IssueUpdated(ctx context.Context) (IssueUpdatedClient, error) {
	s, err := rpcc.NewStream(ctx, "Cast.issueUpdated", d.conn)
	if err != nil {
		return nil, err
	}
	return &issueUpdatedClient{Stream: s}, nil
}

type issueUpdatedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *issueUpdatedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *issueUpdatedClient) Recv() (*IssueUpdatedReply, error) {
	event := new(IssueUpdatedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Cast", Op: "IssueUpdated Recv", Err: err}
	}
	return event, nil
}
