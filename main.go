package main

import (
	"fmt"

	"github.com/jpia/engineadapter/internal/engine"
	"github.com/jpia/engineadapter/internal/engines/black"
	"github.com/jpia/engineadapter/internal/engines/red"
	"github.com/jpia/engineadapter/pkg/models/amodel"
	"github.com/jpia/engineadapter/pkg/models/bmodel"
	"github.com/jpia/engineadapter/pkg/models/cmodel"
)

func main() {

	redEngine := red.NewRedEngine("RedEngine")
	blackEngine := black.NewBlackEngine("BlackEngine")

	singleAdapter := engine.NewEngineAdapter(redEngine.GetName(), false)
	singleAdapter.AddEngine(redEngine)
	singleAdapter.AddEngine(blackEngine)

	multiAdapter := engine.NewEngineAdapter(blackEngine.GetName(), true)
	multiAdapter.AddEngine(redEngine)
	multiAdapter.AddEngine(blackEngine)

	aRequest := amodel.AModel{
		Afield1: "Request A",
		Afield2: 42,
	}
	bRequest := bmodel.BModel{
		Bfield1: "Request B",
		Bfield2: 84,
	}
	cRequest := cmodel.CModel{
		Cfield1: "Request C",
		Cfield2: 168,
	}

	// Example usage of single mode

	response, err := singleAdapter.CallFunc("AFunc", aRequest)
	fmt.Println("Single Mode AFunc Response:", response, "Error:", err)

	response, err = singleAdapter.CallFunc("BFunc", bRequest)
	fmt.Println("Single Mode BFunc Response:", response, "Error:", err)

	response, err = singleAdapter.CallFunc("CFunc", cRequest)
	fmt.Println("Single Mode CFunc Response:", response, "Error:", err)

	// Example usage of multi mode
	response, err = multiAdapter.CallFunc("AFunc", aRequest)
	fmt.Println("Multi Mode AFunc Response:", response, "Error:", err)
	response, err = multiAdapter.CallFunc("BFunc", bRequest)
	fmt.Println("Multi Mode BFunc Response:", response, "Error:", err)
	response, err = multiAdapter.CallFunc("CFunc", cRequest)
	fmt.Println("Multi Mode CFunc Response:", response, "Error:", err)
}
