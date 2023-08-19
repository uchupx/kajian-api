package internal

import "github.com/uchupx/kajian-api/internal/handler"

type Internal struct {
	kajianHandler *handler.KajianHandler
}

func (i *Internal) GetKajianHandler() *handler.KajianHandler {
	if i.kajianHandler == nil {
		i.kajianHandler = &handler.KajianHandler{}
	}

	return i.kajianHandler
}
