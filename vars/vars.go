package vars

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EstrucReg struct {
	Id        int64  `json:"id"`
	Documento int64  `json:"documento"`
	Nombre    string `json:"nombre"`
	Direccion string `json:"direccion"` //Modificar
}

var _ = godotenv.Load(".env") // Cargar del archivo llamado ".env"
var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))
)

const AllowedCORSDomain = "http://localhost"

const Port = ":8001" //Modificar

const NombreRuta = "personas" //Modificar
