package engine

import (
	"fmt"

	"github.com/jpia/engineadapter/pkg/models/amodel"
	"github.com/jpia/engineadapter/pkg/models/bmodel"
	"github.com/jpia/engineadapter/pkg/models/cmodel"
)

type Engine interface {
	GetName() string
	AFunc(request amodel.AModel) (EngineResponse, error)
	BFunc(request bmodel.BModel) (EngineResponse, error)
	CFunc(request cmodel.CModel) (EngineResponse, error)
}

type EngineResponse struct {
	EnginesUsed []string // List of engines used for this response
	Message     string
	Data        any
}

type EngineAdapter struct {
	PrimaryEngine string
	EngineMap     map[string]Engine
	MultiMode     bool
}

func NewEngineAdapter(primaryEngine string, multiMode bool) *EngineAdapter {
	return &EngineAdapter{
		PrimaryEngine: primaryEngine,
		EngineMap:     make(map[string]Engine),
		MultiMode:     multiMode,
	}
}

func (ea *EngineAdapter) GetPrimaryEngine() Engine {
	if engine, exists := ea.EngineMap[ea.PrimaryEngine]; exists {
		return engine
	}
	return nil
}

func (ea *EngineAdapter) GetEngine(name string) Engine {
	if engine, exists := ea.EngineMap[name]; exists {
		return engine
	}
	return nil
}

func (ea *EngineAdapter) AddEngine(engine Engine) {
	if ea.EngineMap == nil {
		ea.EngineMap = make(map[string]Engine)
	}
	ea.EngineMap[engine.GetName()] = engine
}

func (ea *EngineAdapter) CallFuncByEngine(engineName string, funcName string, request any) (EngineResponse, error) {
	engine := ea.GetEngine(engineName)
	if engine == nil {
		return EngineResponse{}, fmt.Errorf("engine %s not found", engineName)
	}

	var response EngineResponse
	var err error

	switch funcName {
	case "AFunc":
		if req, ok := request.(amodel.AModel); ok {
			response, err = engine.AFunc(req)
		} else {
			return EngineResponse{}, fmt.Errorf("invalid request type for AFunc: expected amodel.AModel, got %T", request)
		}
	case "BFunc":
		if req, ok := request.(bmodel.BModel); ok {
			response, err = engine.BFunc(req)
		} else {
			return EngineResponse{}, fmt.Errorf("invalid request type for BFunc: expected bmodel.BModel, got %T", request)
		}
	case "CFunc":
		if req, ok := request.(cmodel.CModel); ok {
			response, err = engine.CFunc(req)
		} else {
			return EngineResponse{}, fmt.Errorf("invalid request type for CFunc: expected cmodel.CModel, got %T", request)
		}
	default:
		return EngineResponse{}, fmt.Errorf("function %s not found in engine %s", funcName, engineName)
	}

	return response, err
}

func (ea *EngineAdapter) CallFunc(funcName string, request any) (EngineResponse, error) {

	if ea.MultiMode {
		var enginesUsed []string
		var responseMap = make(map[string]EngineResponse)
		var errorMap = make(map[string]error)
		for _, engine := range ea.EngineMap {
			response, err := ea.CallFuncByEngine(engine.GetName(), funcName, request)
			if err != nil {
				errorMap[engine.GetName()] = err
				continue
			}
			responseMap[engine.GetName()] = response
			enginesUsed = append(enginesUsed, engine.GetName())
		}
		if len(errorMap) > 0 {
			return EngineResponse{}, fmt.Errorf("errors occurred in some engines: %v", errorMap)
		}
		return EngineResponse{
			EnginesUsed: enginesUsed,
			Message:     "Multi-mode response",
			Data:        responseMap[ea.PrimaryEngine].Data,
		}, nil
	} else {
		return ea.CallFuncByEngine(ea.PrimaryEngine, funcName, request)
	}
}
