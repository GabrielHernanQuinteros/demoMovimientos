package main

import (
	"log"
	"net/http"
	"time"

	mytools "github.com/GabrielHernanQuinteros/demoCommon"
	myroute "github.com/GabrielHernanQuinteros/demoPersonas/routes" //Modificar
	myvars "github.com/GabrielHernanQuinteros/demoPersonas/vars"    //Modificar
	"github.com/gorilla/mux"
)

func main() {

	//Se conecta a la BD para hacer un ping y probar la existencia y conectividad
	auxBaseDatos, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {

		log.Printf("Error con la base de datos" + err.Error())
		return

	} else {

		err = auxBaseDatos.Ping()

		if err != nil {
			log.Printf("Error estableciendo la conexión con la base de datos. Por favor verifique sus credenciales. El error es: " + err.Error())
			return
		}
	}

	//===================================================================================================
	// Definicion de las rutas

	auxRouter := mux.NewRouter()

	mytools.EnableCORS(auxRouter)

	auxRouter.HandleFunc("/"+myvars.NombreRuta, myroute.TraerRegistros).Methods(http.MethodGet)
	auxRouter.HandleFunc("/"+myvars.NombreRuta+"/{id}", myroute.TraerRegistroPorId).Methods(http.MethodGet)
	auxRouter.HandleFunc("/"+myvars.NombreRuta, myroute.CrearRegistro).Methods(http.MethodPost)
	auxRouter.HandleFunc("/"+myvars.NombreRuta, myroute.ModificarRegistro).Methods(http.MethodPut)
	auxRouter.HandleFunc("/"+myvars.NombreRuta+"/{id}", myroute.BorrarRegistro).Methods(http.MethodDelete)
	auxRouter.HandleFunc("/"+myvars.NombreRuta+"PorNombre/{nombre}", myroute.TraerRegistroPorNombre).Methods(http.MethodGet)

	//===================================================================================================
	// Setear y levantar el servidor

	server := &http.Server{
		Handler:      auxRouter,
		Addr:         myvars.Port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("Server started at %s", myvars.Port)
	log.Fatal(server.ListenAndServe())
}
