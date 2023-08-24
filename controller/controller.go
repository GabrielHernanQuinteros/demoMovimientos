package controller

import (
	mytools "github.com/GabrielHernanQuinteros/demoCommon"
	myvars "github.com/GabrielHernanQuinteros/demoPersonas/vars" //Modificar
)

func CrearRegistroSQL(registro myvars.EstrucReg) error {

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return err
	}

	_, err = bd.Exec("INSERT INTO personas (documento, nombre, direccion) VALUES (?, ?, ?)", registro.Documento, registro.Nombre, registro.Direccion) //Modificar

	return err

}

func BorrarRegistroSQL(id int64) error {

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return err
	}

	_, err = bd.Exec("DELETE FROM personas WHERE id = ?", id) //Modificar

	return err
}

func ModificarRegistroSQL(registro myvars.EstrucReg) error {

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return err
	}

	_, err = bd.Exec("UPDATE personas SET documento = ?, nombre = ?, direccion = ? WHERE id = ?", registro.Documento, registro.Nombre, registro.Direccion, registro.Id) //Modificar

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
	rows, err := bd.Query("SELECT * FROM personas") //Modificar

	if err != nil {
		return arrRegistros, err
	}

	// Iterate rows...
	for rows.Next() {
		// In each step, scan one row
		var registro myvars.EstrucReg

		err = rows.Scan(&registro.Id, &registro.Documento, &registro.Nombre, &registro.Direccion) //Modificar

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

	row := bd.QueryRow("SELECT * FROM personas WHERE id = ?", id) //Modificar

	err = row.Scan(&registro.Id, &registro.Documento, &registro.Nombre, &registro.Direccion) //Modificar

	if err != nil {
		return registro, err
	}

	// Success!
	return registro, nil

}

func TraerRegistroPorNombreSQL(parNombre string) (myvars.EstrucReg, error) {

	var registro myvars.EstrucReg

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return registro, err
	}

	row := bd.QueryRow("SELECT * FROM personas WHERE nombre = ?", parNombre) //Modificar

	err = row.Scan(&registro.Id, &registro.Documento, &registro.Nombre, &registro.Direccion) //Modificar

	if err != nil {
		return registro, err
	}

	// Success!
	return registro, nil

}
