package processor

import (
	"github.com/acore2026/ausf/internal/sbi/consumer"
	"github.com/acore2026/ausf/pkg/app"
)

type ProcessorAusf interface {
	app.App

	Consumer() *consumer.Consumer
}

type Processor struct {
	ProcessorAusf
}

func NewProcessor(ausf ProcessorAusf) (*Processor, error) {
	p := &Processor{
		ProcessorAusf: ausf,
	}
	return p, nil
}
