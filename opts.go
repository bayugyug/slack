package slack

import (
	"errors"
	"strings"
)

// Options ...
type Options struct {
	Host string `json:"host,omitempty"`
	Path string `json:"path,omitempty"`
}

// Payload ...
type Payload struct {
	Title       string       `json:"title,omitempty"`
	Message     string       `json:"message,omitempty"`
	WithDivider bool         `json:"with_divider,omitempty"`
	IsHTML      bool         `json:"is_html,omitempty"`
	Priority    PriorityType `json:"priority,omitempty"`
	Icon        IconType     `json:"icon,omitempty"`
	Href        Href         `json:"href,omitempty"`
	Here        string       `json:"here,omitempty"`
}

const (
	// DividerBlockType ...
	DividerBlockType = "divider"
	// SectionBlockType ...
	SectionBlockType = "section"
	// MarkDownBlockType ...
	MarkDownBlockType = "mrkdwn"
	// ContextBlockType ...
	ContextBlockType = "context"
	// ImageBlockType ...
	ImageBlockType = "image"
	// PlainTextBlockType ...
	PlainTextBlockType = "plain_text"
)

// Block ...
type Block struct {
	BlockID   string     `json:"block_id,omitempty"`
	Type      string     `json:"type,omitempty"`
	Text      *Text      `json:"text,omitempty"`
	Title     *Text      `json:"title,omitempty"`
	Fields    []*Text    `json:"fields,omitempty"`
	Elements  []*Text    `json:"elements,omitempty"`
	Emoji     bool       `json:"emoji,omitempty"`
	Accessory *Accessory `json:"accessory,omitempty"`
}

// Blocks ...
type Blocks struct {
	Channel string   `json:"channel,omitempty"`
	Blocks  []*Block `json:"blocks,omitempty"`
}

// Href ...
type Href struct {
	Link string `json:"link,omitempty"`
	Text string `json:"text,omitempty"`
}

// Accessory ...
type Accessory struct {
	Type     string `json:"type,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	AltText  string `json:"alt_text,omitempty"`
}

// Text ...
type Text struct {
	Type string `json:"type,omitempty"`
	Text string `json:"text,omitempty"`
}

const (
	defaultTimeout = 20
	// PinImageURL  ...
	PinImageURL = `https://image.freepik.com/free-photo/red-drawing-pin_1156-445.jpg`
	// GoogleSearchPath ...
	GoogleSearchPath = `https://www.google.com/search?q=date+and+time+today`
	// AtHere ...
	AtHere = `@here`
)

var (
	// ErrMissingParams ...
	ErrMissingParams = errors.New("missing required parameters")
	// ServiceEndPoint ...
	ServiceEndPoint = `https://hooks.slack.com`
)

// PriorityType type
type PriorityType int

const (
	// PriorityLow  ...
	PriorityLow PriorityType = iota
	// PriorityMedium  ...
	PriorityMedium
	// PriorityHigh  ...
	PriorityHigh
	// PriorityCritical  ...
	PriorityCritical
)

var priorityList = [...]string{"low", "medium", "high", "critical"}

// String mapping
func (s PriorityType) String() string {
	return strings.TrimSpace(priorityList[s])
}

// IconType type
type IconType int

const (
	// IconTypeSpeechBalloon ...
	IconTypeSpeechBalloon IconType = iota
	// IconTypeOkHand  ...
	IconTypeOkHand
	// IconTypeThumbsUp  ...
	IconTypeThumbsUp
	// IconTypeStar ...
	IconTypeStar
	// IconTypeHeart ...
	IconTypeHeart
	// IconTypeWarning ...
	IconTypeWarning
	// IconTypeBug  ...
	IconTypeBug
	// IconTypeCritical ...
	IconTypeCritical
)

var iconList = [...]string{
	":speech_balloon:",
	":ok_hand:",
	":thumbsup:",
	":star:",
	":heart:",
	":warning:",
	":bug:",
	":helmet_with_white_cross:",
}

// String mapping
func (s IconType) String() string {
	return strings.TrimSpace(iconList[s])
}
