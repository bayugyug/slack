package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	ioutil "io"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"

	"github.com/bayugyug/commons"
)

// NotificationCreator  ...
//
//go:generate mockgen -destination ./mock/mock_notificationcreator.go -package mock github.com/bayugyug/slack NotificationCreator
type NotificationCreator interface {
	Notify(messages []*Payload, meta ...Block) error
	WithTimer(p bool)
}

// Notification  ...
type Notification struct {
	opts      *Options
	client    *http.Client
	re        *regexp.Regexp
	withTimer bool
}

// NewNotification ...
func NewNotification(path string) NotificationCreator {
	return &Notification{
		opts: &Options{
			Host: ServiceEndPoint,
			Path: `/` + strings.TrimLeft(path, "/"),
		},
		client: commons.DefaultHTTPClient(defaultTimeout),
		re:     regexp.MustCompile(`\\r\\n`),
	}
}

// Notify ...
func (s *Notification) Notify(messages []*Payload, meta ...Block) error {
	// sanity check
	if len(messages) <= 0 && len(meta) <= 0 {
		return ErrMissingParams
	}

	blocks := &Blocks{
		Blocks: make([]*Block, 0),
	}

	for _, val := range messages {
		blk := &Block{
			BlockID: uuid.New().String(),
			Type:    SectionBlockType,
			Text: &Text{
				Type: MarkDownBlockType,
			},
		}

		msg := strings.TrimSpace(val.Icon.String())

		if val.Title != "" {
			msg = strings.TrimSpace(fmt.Sprintf("%s  *%s*", msg, val.Title))
		}
		if val.Here != "" || val.Priority == PriorityCritical {
			msg = fmt.Sprintf("%s\n\n\n%s `%s`\n", msg, AtHere, val.Here)
		}
		if val.Message != "" {
			msg = fmt.Sprintf("%s\n\n\n%s\n", msg, val.Message)
		}

		// add the msg
		blk.Text.Text = msg

		// nice
		blocks.Blocks = append(blocks.Blocks, blk)

		if val.WithDivider {
			blocks.Blocks = append(blocks.Blocks, &Block{
				BlockID: uuid.New().String(),
				Type:    DividerBlockType,
			})
		}
	}

	// add extra
	for _, m := range meta {
		blocks.Blocks = append(blocks.Blocks, &m)
	}

	// last updated
	if s.withTimer {
		blocks.Blocks = append(blocks.Blocks, &Block{
			BlockID: uuid.New().String(),
			Type:    SectionBlockType,
			Text: &Text{
				Type: MarkDownBlockType,
				Text: fmt.Sprintf(":timer_clock: <%s|%s>",
					GoogleSearchPath,
					time.Now().Local().Format(time.RFC3339),
				),
			},
		})
	}

	bfr, err := json.Marshal(&blocks)
	if err != nil {
		log.WithError(err).Println("failed to marshal payload")
		return err
	}

	link := fmt.Sprintf("%s%s", s.opts.Host, s.opts.Path)

	req, err := http.NewRequest(http.MethodPost, link, bytes.NewReader(bfr))
	if err != nil {
		log.WithError(err).Println("failed to make req")
		return err
	}

	// add headers
	req.Header.Add("Accept", "application/json")

	// Make request
	ret, err := s.client.Do(req)
	if err != nil {
		log.WithError(err).Println("failed to send req")
		return err
	}

	defer func() {
		_ = ret.Body.Close()
	}()

	if ret.StatusCode != http.StatusOK {
		err = fmt.Errorf("failed response code: %v", ret.StatusCode)
		log.WithError(err).Println("status")
		return err
	}

	b, _ := ioutil.ReadAll(ret.Body)
	if tmp := string(b); tmp != `ok` {
		err = fmt.Errorf("failed response body not ok")
		log.WithError(err).Println("body")
		return err
	}

	// check params
	return nil
}

// WithTimer ...
func (s *Notification) WithTimer(p bool) {
	s.withTimer = p
}
