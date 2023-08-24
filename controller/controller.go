package controller

import (
	mytools "github.com/GabrielHernanQuinteros/demoCommon"
	myvars "github.com/GabrielHernanQuinteros/demoMovimientos/vars" //Modificar
)

func CrearRegistroSQL(registro myvars.EstrucReg) error {

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return err
	}

	//==========================================================================================================

	var auxIdPersona int64

	err = bd.QueryRow("SELECT id FROM personas WHERE nombre = ?", registro.NombrePersona).Scan(&auxIdPersona)

	if err != nil {
		return err
	}

	registro.IdPersona = auxIdPersona

	//==========================================================================================================

	var auxIdArticulo int64

	err = bd.QueryRow("SELECT id FROM personas WHERE nombre = ?", registro.NombreArticulo).Scan(&auxIdArticulo)

	if err != nil {
		return err
	}

	registro.IdArticulo = auxIdArticulo

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

func TraerRegistroPorNombreSQL(parNombre string) (myvars.EstrucReg, error) {

	var registro myvars.EstrucReg

	bd, err := mytools.ConectarDB(myvars.ConnectionString)

	if err != nil {
		return registro, err
	}

	row := bd.QueryRow("SELECT * FROM movimientos WHERE nombre = ?", parNombre) //Modificar

	err = row.Scan(&registro.Id, &registro.IdPersona, &registro.IdArticulo, &registro.Tipo, &registro.Cantidad) //Modificar

	if err != nil {
		return registro, err
	}

	// Success!
	return registro, nil

}
