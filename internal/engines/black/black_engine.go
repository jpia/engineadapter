package black

import (
	"github.com/jpia/engineadapter/internal/engine"
	"github.com/jpia/engineadapter/pkg/models/amodel"
	"github.com/jpia/engineadapter/pkg/models/bmodel"
	"github.com/jpia/engineadapter/pkg/models/cmodel"
)

type BlackEngine struct {
	name string
}

func NewBlackEngine(name string) *BlackEngine {
	return &BlackEngine{
		name: name,
	}
}

func (b *BlackEngine) GetName() string {
	return b.name
}

func (b *BlackEngine) AFunc(request amodel.AModel) (engine.EngineResponse, error) {
	// Implement the logic for AFunc
	return engine.EngineResponse{
		EnginesUsed: []string{b.name},
		Message:     "AFunc executed",
		Data:        request,
	}, nil
}

func (b *BlackEngine) BFunc(request bmodel.BModel) (engine.EngineResponse, error) {
	// Implement the logic for BFunc
	return engine.EngineResponse{
		EnginesUsed: []string{b.name},
		Message:     "BFunc executed",
		Data:        request,
	}, nil
}

func (b *BlackEngine) CFunc(request cmodel.CModel) (engine.EngineResponse, error) {
	// Implement the logic for CFunc
	return engine.EngineResponse{
		EnginesUsed: []string{b.name},
		Message:     "CFunc executed",
		Data:        request,
	}, nil
}
