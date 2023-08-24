package routes

import (
	"encoding/json"
	"net/http"

	mytools "github.com/GabrielHernanQuinteros/demoCommon"
	myctrllr "github.com/GabrielHernanQuinteros/demoPersonas/controller" //Modificar
	myvars "github.com/GabrielHernanQuinteros/demoPersonas/vars"         //Modificar
	"github.com/gorilla/mux"
)

//===================================================================================================
// Funciones de ROUTERS

func TraerRegistros(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxRegistros, err := myctrllr.TraerRegistrosSQL()

	if err == nil {
		mytools.RespondWithSuccess(auxRegistros, parWriter)
	} else {
		mytools.RespondWithError(err, parWriter)
	}

}

func TraerRegistroPorId(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxIdString := mux.Vars(parRequest)["id"]
	auxId, err := mytools.StringToInt64(auxIdString)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
		return
	}

	auxRegistro, err := myctrllr.TraerRegistroPorIdSQL(auxId)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
	} else {
		mytools.RespondWithSuccess(auxRegistro, parWriter)
	}

}

func CrearRegistro(parWriter http.ResponseWriter, parRequest *http.Request) {

	var auxRegistro myvars.EstrucReg
	err := json.NewDecoder(parRequest.Body).Decode(&auxRegistro)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
	} else {
		err := myctrllr.CrearRegistroSQL(auxRegistro)

		if err != nil {
			mytools.RespondWithError(err, parWriter)
		} else {
			mytools.RespondWithSuccess(true, parWriter)
		}

	}

}

func ModificarRegistro(parWriter http.ResponseWriter, parRequest *http.Request) {

	var auxRegistro myvars.EstrucReg
	err := json.NewDecoder(parRequest.Body).Decode(&auxRegistro)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
	} else {
		err := myctrllr.ModificarRegistroSQL(auxRegistro)

		if err != nil {
			mytools.RespondWithError(err, parWriter)
		} else {
			mytools.RespondWithSuccess(true, parWriter)
		}

	}

}

func BorrarRegistro(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxIdString := mux.Vars(parRequest)["id"]
	auxId, err := mytools.StringToInt64(auxIdString)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
		return
	}

	err = myctrllr.BorrarRegistroSQL(auxId)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
	} else {
		mytools.RespondWithSuccess(true, parWriter)
	}

}

func TraerRegistroPorNombre(parWriter http.ResponseWriter, parRequest *http.Request) {

	auxNombre := mux.Vars(parRequest)["nombre"]

	auxRegistro, err := myctrllr.TraerRegistroPorNombreSQL(auxNombre)

	if err != nil {
		mytools.RespondWithError(err, parWriter)
	} else {
		mytools.RespondWithSuccess(auxRegistro, parWriter)
	}

}
