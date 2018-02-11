// Code generated by cdpgen. DO NOT EDIT.

package page

import (
	"github.com/mafredri/cdp/protocol/network"
	"github.com/mafredri/cdp/rpcc"
)

// DOMContentEventFiredClient is a client for DOMContentEventFired events.
type DOMContentEventFiredClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*DOMContentEventFiredReply, error)
	rpcc.Stream
}

// DOMContentEventFiredReply is the reply for DOMContentEventFired events.
type DOMContentEventFiredReply struct {
	Timestamp network.MonotonicTime `json:"timestamp"` // No description.
}

// FrameAttachedClient is a client for FrameAttached events. Fired when frame
// has been attached to its parent.
type FrameAttachedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*FrameAttachedReply, error)
	rpcc.Stream
}

// FrameNavigatedReply is the reply for FrameNavigated events.
type FrameNavigatedReply struct {
	Frame Frame `json:"frame"` // Frame object.
}

// FrameResizedClient is a client for FrameResized events.
type FrameResizedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*FrameResizedReply, error)
	rpcc.Stream
}

// FrameResizedReply is the reply for FrameResized events.
type FrameResizedReply struct {
}

// FrameScheduledNavigationClient is a client for FrameScheduledNavigation events.
// Fired when frame schedules a potential navigation.
type FrameScheduledNavigationClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*FrameScheduledNavigationReply, error)
	rpcc.Stream
}

// InterstitialHiddenReply is the reply for InterstitialHidden events.
type InterstitialHiddenReply struct {
}

// InterstitialShownClient is a client for InterstitialShown events. Fired
// when interstitial page was shown
type InterstitialShownClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*InterstitialShownReply, error)
	rpcc.Stream
}

// InterstitialShownReply is the reply for InterstitialShown events.
type InterstitialShownReply struct {
}

// JavascriptDialogClosedClient is a client for JavascriptDialogClosed events.
// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or
// onbeforeunload) has been closed.
type JavascriptDialogClosedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*JavascriptDialogClosedReply, error)
	rpcc.Stream
}

// JavascriptDialogClosedReply is the reply for JavascriptDialogClosed events.
type JavascriptDialogClosedReply struct {
	Result    bool   `json:"result"`    // Whether dialog was confirmed.
	UserInput string `json:"userInput"` // User input in case of prompt.
}

// JavascriptDialogOpeningClient is a client for JavascriptDialogOpening events.
// Fired when a JavaScript initiated dialog (alert, confirm, prompt, or
// onbeforeunload) is about to open.
type JavascriptDialogOpeningClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*JavascriptDialogOpeningReply, error)
	rpcc.Stream
}

// JavascriptDialogOpeningReply is the reply for JavascriptDialogOpening events.
type JavascriptDialogOpeningReply struct {
	URL           string     `json:"url"`                     // Frame url.
	Message       string     `json:"message"`                 // Message that will be displayed by the dialog.
	Type          DialogType `json:"type"`                    // Dialog type.
	DefaultPrompt *string    `json:"defaultPrompt,omitempty"` // Default dialog prompt.
}

// LifecycleEventClient is a client for LifecycleEvent events. Fired for top
// level page lifecycle events such as navigation, load, paint, etc.
type LifecycleEventClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*LifecycleEventReply, error)
	rpcc.Stream
}

// LoadEventFiredReply is the reply for LoadEventFired events.
type LoadEventFiredReply struct {
	Timestamp network.MonotonicTime `json:"timestamp"` // No description.
}

// ScreencastFrameClient is a client for ScreencastFrame events. Compressed
// image data requested by the `startScreencast`.
type ScreencastFrameClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ScreencastFrameReply, error)
	rpcc.Stream
}

// ScreencastFrameReply is the reply for ScreencastFrame events.
type ScreencastFrameReply struct {
	Data      []byte                  `json:"data"`      // Base64-encoded compressed image.
	Metadata  ScreencastFrameMetadata `json:"metadata"`  // Screencast frame metadata.
	SessionID int                     `json:"sessionId"` // Frame number.
}

// ScreencastVisibilityChangedClient is a client for ScreencastVisibilityChanged events.
// Fired when the page with currently enabled screencast was shown or hidden `.
type ScreencastVisibilityChangedClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*ScreencastVisibilityChangedReply, error)
	rpcc.Stream
}

// ScreencastVisibilityChangedReply is the reply for ScreencastVisibilityChanged events.
type ScreencastVisibilityChangedReply struct {
	Visible bool `json:"visible"` // True if the page is visible.
}

// WindowOpenClient is a client for WindowOpen events. Fired when a new window
// is going to be opened, via window.open(), link click, form submission, etc.
type WindowOpenClient interface {
	// Recv calls RecvMsg on rpcc.Stream, blocks until the event is
	// triggered, context canceled or connection closed.
	Recv() (*WindowOpenReply, error)
	rpcc.Stream
}

// WindowOpenReply is the reply for WindowOpen events.
type WindowOpenReply struct {
	URL            string   `json:"url"`            // The URL for the new window.
	WindowName     string   `json:"windowName"`     // Window name.
	WindowFeatures []string `json:"windowFeatures"` // An array of enabled window features.
	UserGesture    bool     `json:"userGesture"`    // Whether or not it was triggered by user gesture.
}
