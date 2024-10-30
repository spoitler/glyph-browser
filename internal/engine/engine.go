package engine

import (
	"context"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
)

type Engine struct {
	sigChan chan os.Signal
}

func New() *Engine {
	return &Engine{}
}

func (e Engine) Run() {
	e.registerFromOsSignal()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// get the worker init implementation i've created from work maybe
	// maybe a pipe need to be a struct. that make possible to have a context for each pipe. and each pipe based their context fromm the engine context.
	// if the engine cancel is ctx due to interrupt for example, it can cancel ctx for each pipe. each pipe manage and track their context.
	// pipe is representing a "tab"
	e.startNewPipe(ctx, cancel)
	// start a new pipe of goroutine for each "tab"

	<-e.sigChan
	log.Info().Msg("Shutting down...")
}

func (e Engine) registerFromOsSignal() {
	e.sigChan = make(chan os.Signal)
	signal.Notify(e.sigChan, os.Interrupt, os.Kill, syscall.SIGTERM)
}

func (e Engine) startNewPipe(context.Context, context.CancelFunc) {

	// create pipe builder

	// NETWORKING
	// used to resolve an uri and get the content of the page pass it to the renderer through a chan

	// RENDERER
	// instantiate a parser for each piece of a page content (HTML, CSS,....)
}
