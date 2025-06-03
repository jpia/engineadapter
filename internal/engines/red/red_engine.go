package red

import (
	"github.com/jpia/engineadapter/internal/engine"
	"github.com/jpia/engineadapter/pkg/models/amodel"
	"github.com/jpia/engineadapter/pkg/models/bmodel"
	"github.com/jpia/engineadapter/pkg/models/cmodel"
)

type RedEngine struct {
	name string
}

func NewRedEngine(name string) *RedEngine {
	return &RedEngine{name: name}
}

func (e *RedEngine) GetName() string {
	return e.name
}
func (e *RedEngine) AFunc(request amodel.AModel) (engine.EngineResponse, error) {
	return engine.EngineResponse{
		EnginesUsed: []string{e.name},
		Message:     "RedEngine AFunc executed",
		Data:        request,
	}, nil
}

func (e *RedEngine) BFunc(request bmodel.BModel) (engine.EngineResponse, error) {
	return engine.EngineResponse{
		EnginesUsed: []string{e.name},
		Message:     "RedEngine BFunc executed",
		Data:        request,
	}, nil
}

func (e *RedEngine) CFunc(request cmodel.CModel) (engine.EngineResponse, error) {
	return engine.EngineResponse{
		EnginesUsed: []string{e.name},
		Message:     "RedEngine CFunc executed",
		Data:        request,
	}, nil
}
