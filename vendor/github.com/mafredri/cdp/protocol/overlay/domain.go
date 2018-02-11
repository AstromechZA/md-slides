// Code generated by cdpgen. DO NOT EDIT.

// Package overlay implements the Overlay domain. This domain provides various
// functionality related to drawing atop the inspected page.
package overlay

import (
	"context"

	"github.com/mafredri/cdp/protocol/internal"
	"github.com/mafredri/cdp/rpcc"
)

// domainClient is a client for the Overlay domain. This domain provides
// various functionality related to drawing atop the inspected page.
type domainClient struct{ conn *rpcc.Conn }

// NewClient returns a client for the Overlay domain with the connection set to conn.
func NewClient(conn *rpcc.Conn) *domainClient {
	return &domainClient{conn: conn}
}

// Disable invokes the Overlay method. Disables domain notifications.
func (d *domainClient) Disable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Overlay.disable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "Disable", Err: err}
	}
	return
}

// Enable invokes the Overlay method. Enables domain notifications.
func (d *domainClient) Enable(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Overlay.enable", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "Enable", Err: err}
	}
	return
}

// GetHighlightObjectForTest invokes the Overlay method. For testing.
func (d *domainClient) GetHighlightObjectForTest(ctx context.Context, args *GetHighlightObjectForTestArgs) (reply *GetHighlightObjectForTestReply, err error) {
	reply = new(GetHighlightObjectForTestReply)
	if args != nil {
		err = rpcc.Invoke(ctx, "Overlay.getHighlightObjectForTest", args, reply, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Overlay.getHighlightObjectForTest", nil, reply, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "GetHighlightObjectForTest", Err: err}
	}
	return
}

// HideHighlight invokes the Overlay method. Hides any highlight.
func (d *domainClient) HideHighlight(ctx context.Context) (err error) {
	err = rpcc.Invoke(ctx, "Overlay.hideHighlight", nil, nil, d.conn)
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "HideHighlight", Err: err}
	}
	return
}

// HighlightFrame invokes the Overlay method. Highlights owner element of the
// frame with given id.
func (d *domainClient) HighlightFrame(ctx context.Context, args *HighlightFrameArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Overlay.highlightFrame", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Overlay.highlightFrame", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "HighlightFrame", Err: err}
	}
	return
}

// HighlightNode invokes the Overlay method. Highlights DOM node with given id
// or with the given JavaScript object wrapper. Either nodeId or objectId must
// be specified.
func (d *domainClient) HighlightNode(ctx context.Context, args *HighlightNodeArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Overlay.highlightNode", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Overlay.highlightNode", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "HighlightNode", Err: err}
	}
	return
}

// HighlightQuad invokes the Overlay method. Highlights given quad.
// Coordinates are absolute with respect to the main frame viewport.
func (d *domainClient) HighlightQuad(ctx context.Context, args *HighlightQuadArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Overlay.highlightQuad", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Overlay.highlightQuad", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "HighlightQuad", Err: err}
	}
	return
}

// HighlightRect invokes the Overlay method. Highlights given rectangle.
// Coordinates are absolute with respect to the main frame viewport.
func (d *domainClient) HighlightRect(ctx context.Context, args *HighlightRectArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Overlay.highlightRect", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Overlay.highlightRect", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "HighlightRect", Err: err}
	}
	return
}

// SetInspectMode invokes the Overlay method. Enters the 'inspect' mode. In
// this mode, elements that user is hovering over are highlighted. Backend then
// generates 'inspectNodeRequested' event upon element selection.
func (d *domainClient) SetInspectMode(ctx context.Context, args *SetInspectModeArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Overlay.setInspectMode", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Overlay.setInspectMode", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "SetInspectMode", Err: err}
	}
	return
}

// SetPausedInDebuggerMessage invokes the Overlay method.
func (d *domainClient) SetPausedInDebuggerMessage(ctx context.Context, args *SetPausedInDebuggerMessageArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Overlay.setPausedInDebuggerMessage", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Overlay.setPausedInDebuggerMessage", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "SetPausedInDebuggerMessage", Err: err}
	}
	return
}

// SetShowDebugBorders invokes the Overlay method. Requests that backend shows
// debug borders on layers
func (d *domainClient) SetShowDebugBorders(ctx context.Context, args *SetShowDebugBordersArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Overlay.setShowDebugBorders", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Overlay.setShowDebugBorders", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "SetShowDebugBorders", Err: err}
	}
	return
}

// SetShowFPSCounter invokes the Overlay method. Requests that backend shows
// the FPS counter
func (d *domainClient) SetShowFPSCounter(ctx context.Context, args *SetShowFPSCounterArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Overlay.setShowFPSCounter", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Overlay.setShowFPSCounter", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "SetShowFPSCounter", Err: err}
	}
	return
}

// SetShowPaintRects invokes the Overlay method. Requests that backend shows
// paint rectangles
func (d *domainClient) SetShowPaintRects(ctx context.Context, args *SetShowPaintRectsArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Overlay.setShowPaintRects", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Overlay.setShowPaintRects", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "SetShowPaintRects", Err: err}
	}
	return
}

// SetShowScrollBottleneckRects invokes the Overlay method. Requests that
// backend shows scroll bottleneck rects
func (d *domainClient) SetShowScrollBottleneckRects(ctx context.Context, args *SetShowScrollBottleneckRectsArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Overlay.setShowScrollBottleneckRects", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Overlay.setShowScrollBottleneckRects", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "SetShowScrollBottleneckRects", Err: err}
	}
	return
}

// SetShowViewportSizeOnResize invokes the Overlay method. Paints viewport
// size upon main frame resize.
func (d *domainClient) SetShowViewportSizeOnResize(ctx context.Context, args *SetShowViewportSizeOnResizeArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Overlay.setShowViewportSizeOnResize", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Overlay.setShowViewportSizeOnResize", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "SetShowViewportSizeOnResize", Err: err}
	}
	return
}

// SetSuspended invokes the Overlay method.
func (d *domainClient) SetSuspended(ctx context.Context, args *SetSuspendedArgs) (err error) {
	if args != nil {
		err = rpcc.Invoke(ctx, "Overlay.setSuspended", args, nil, d.conn)
	} else {
		err = rpcc.Invoke(ctx, "Overlay.setSuspended", nil, nil, d.conn)
	}
	if err != nil {
		err = &internal.OpError{Domain: "Overlay", Op: "SetSuspended", Err: err}
	}
	return
}

func (d *domainClient) InspectNodeRequested(ctx context.Context) (InspectNodeRequestedClient, error) {
	s, err := rpcc.NewStream(ctx, "Overlay.inspectNodeRequested", d.conn)
	if err != nil {
		return nil, err
	}
	return &inspectNodeRequestedClient{Stream: s}, nil
}

type inspectNodeRequestedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *inspectNodeRequestedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *inspectNodeRequestedClient) Recv() (*InspectNodeRequestedReply, error) {
	event := new(InspectNodeRequestedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Overlay", Op: "InspectNodeRequested Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) NodeHighlightRequested(ctx context.Context) (NodeHighlightRequestedClient, error) {
	s, err := rpcc.NewStream(ctx, "Overlay.nodeHighlightRequested", d.conn)
	if err != nil {
		return nil, err
	}
	return &nodeHighlightRequestedClient{Stream: s}, nil
}

type nodeHighlightRequestedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *nodeHighlightRequestedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *nodeHighlightRequestedClient) Recv() (*NodeHighlightRequestedReply, error) {
	event := new(NodeHighlightRequestedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Overlay", Op: "NodeHighlightRequested Recv", Err: err}
	}
	return event, nil
}

func (d *domainClient) ScreenshotRequested(ctx context.Context) (ScreenshotRequestedClient, error) {
	s, err := rpcc.NewStream(ctx, "Overlay.screenshotRequested", d.conn)
	if err != nil {
		return nil, err
	}
	return &screenshotRequestedClient{Stream: s}, nil
}

type screenshotRequestedClient struct{ rpcc.Stream }

// GetStream returns the original Stream for use with cdp.Sync.
func (c *screenshotRequestedClient) GetStream() rpcc.Stream { return c.Stream }

func (c *screenshotRequestedClient) Recv() (*ScreenshotRequestedReply, error) {
	event := new(ScreenshotRequestedReply)
	if err := c.RecvMsg(event); err != nil {
		return nil, &internal.OpError{Domain: "Overlay", Op: "ScreenshotRequested Recv", Err: err}
	}
	return event, nil
}
