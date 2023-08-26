package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	mytools "github.com/GabrielHernanQuinteros/demoCommon"
	myvars "github.com/GabrielHernanQuinteros/demoMovimientos/vars" //Modificar
)

func CrearRegistroSQL(registro myvars.EstrucReg) error {

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return err
	}

	//==========================================================================================================

	auxPeticion := "http://" + IpApiPersonas + ":8001/personasPorNombre/" + registro.NombrePersona

	response, err := http.Get(auxPeticion)

	if err != nil {
		return fmt.Errorf("Error: API Personas")
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if strings.Contains(string(responseData), "sql: no rows in result set") {
		return fmt.Errorf("Error: Persona no encontrada")
	}

	var data map[string]interface{}

	// Deserializar la cadena JSON en la estructura
	err = json.Unmarshal([]byte(responseData), &data)

	if err != nil {
		return err
	}

	// Acceder a la clave "id" en la estructura, cambiando el TIPO
	registro.IdPersona, err = mytools.InterfaceToInt64(data["id"])

	if err != nil {
		return err
	}

	//==========================================================================================================

	auxPeticion = "http://" + IpApiArticulos + ":8000/articulosPorNombre/" + registro.NombreArticulo

	response, err = http.Get(auxPeticion)

	if err != nil {
		return fmt.Errorf("Error: API Articulos")
	}

	responseData, err = ioutil.ReadAll(response.Body)

	if strings.Contains(string(responseData), "sql: no rows in result set") {
		return fmt.Errorf("Error: Artículo no encontrado")
	}

	//var data map[string]interface{}

	// Deserializar la cadena JSON en la estructura
	err = json.Unmarshal([]byte(responseData), &data)

	if err != nil {
		return err
	}

	// Acceder a la clave "id" en la estructura, cambiando el TIPO
	registro.IdArticulo, err = mytools.InterfaceToInt64(data["id"])

	if err != nil {
		return err
	}

	// ************   UPDATE DEL STOCK EN EL ARTICULO   **************

	bd3, err := mytools.ConectarDB(myvars.ConnectionStringArticulos)

	if err != nil {
		return err
	}

	var auxCantidad int64

	if registro.Tipo == "compra" {
		auxCantidad = registro.Cantidad
	} else {

		if registro.Tipo == "venta" {
			auxCantidad = registro.Cantidad * -1
		} else {
			return fmt.Errorf("Error: El tipo debe ser 'compra' ó 'venta'")
		}

	}

	_, err = bd3.Exec("UPDATE articulos SET stock = stock + ? WHERE id = ?", auxCantidad, registro.IdArticulo)

	if err != nil {
		return err
	}

	//==========================================================================================================

	_, err = bd.Exec("INSERT INTO movimientos (idPersona, idArticulo, tipo, cantidad) VALUES (?, ?, ?, ?)", registro.IdPersona, registro.IdArticulo, registro.Tipo, registro.Cantidad) //Modificar

	return err

}

func BorrarRegistroSQL(id int64) error {

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return err
	}

	_, err = bd.Exec("DELETE FROM movimientos WHERE id = ?", id) //Modificar

	return err
}

func ModificarRegistroSQL(registro myvars.EstrucReg) error {

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return err
	}

	_, err = bd.Exec("UPDATE movimientos SET IdPersona = ?, IdArticulo = ?, Tipo = ?, Cantidad = ? WHERE id = ?", registro.IdPersona, registro.IdArticulo, registro.Tipo, registro.Cantidad, registro.Id) //Modificar

	return err
}

func TraerRegistrosSQL() ([]myvars.EstrucReg, error) {

	//Declare an array because if there's error, we return it empty
	arrRegistros := []myvars.EstrucReg{}

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return arrRegistros, err
	}

	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT * FROM movimientos") //Modificar

	if err != nil {
		return arrRegistros, err
	}

	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var registro myvars.EstrucReg

		err = rows.Scan(&registro.Id, &registro.IdPersona, &registro.IdArticulo, &registro.Tipo, &registro.Cantidad) //Modificar

		if err != nil {
			return arrRegistros, err
		}

		// and append it to the array
		arrRegistros = append(arrRegistros, registro)
	}

	return arrRegistros, nil

}

func TraerRegistroPorIdSQL(id int64) (myvars.EstrucReg, error) {

	var registro myvars.EstrucReg

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return registro, err
	}

	row := bd.QueryRow("SELECT * FROM movimientos WHERE id = ?", id) //Modificar

	err = row.Scan(&registro.Id, &registro.IdPersona, &registro.IdArticulo, &registro.Tipo, &registro.Cantidad) //Modificar

	if err != nil {
		return registro, err
	}

	// Success!
	return registro, nil

}

func TraerRegistroPorNombreSQL(parNombre string) ([]myvars.EstrucReg, error) {

	arrRegistros := []myvars.EstrucReg{}

	var auxNombreArticulo string

	//==========================================================================================================

	bd2, err := mytools.ConectarDB(myvars.ConnectionStringPersonas)

	var auxIdPersona int64

	err = bd2.QueryRow("SELECT id FROM personas WHERE nombre = ?", parNombre).Scan(&auxIdPersona)

	if err != nil {
		return arrRegistros, fmt.Errorf("Error: Persona no encontrada")
	}

	//==========================================================================================================

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	// Get rows so we can iterate them
	rows, err := bd.Query("SELECT * FROM movimientos WHERE idpersona = ?", auxIdPersona) //Modificar

	if err != nil {
		return arrRegistros, err
	}

	bd3, err := mytools.ConectarDB(myvars.ConnectionStringArticulos)

	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var registro myvars.EstrucReg

		err = rows.Scan(&registro.Id, &registro.IdPersona, &registro.IdArticulo, &registro.Tipo, &registro.Cantidad) //Modificar

		if err != nil {
			return arrRegistros, err
		}

		err = bd3.QueryRow("SELECT nombre FROM articulos WHERE id = ?", registro.IdArticulo).Scan(&auxNombreArticulo)

		if err != nil {
			registro.NombreArticulo = "Error: Articulo no encontrada"
		} else {
			registro.NombreArticulo = auxNombreArticulo
		}

		// and append it to the array
		arrRegistros = append(arrRegistros, registro)
	}

	return arrRegistros, nil

}
