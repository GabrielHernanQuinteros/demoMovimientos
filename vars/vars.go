package vars

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type EstrucReg struct {
	Id             int64  `json:"id"`
	IdPersona      int64  `json:"idpersona"`
	NombrePersona  string `json:"nombrepersona"`
	IdArticulo     int64  `json:"idarticulo"`
	NombreArticulo string `json:"nombrearticulo"`
	Tipo           string `json:"tipo"`
	Cantidad       int64  `json:"cantidad"` //Modificar
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

var (
	ConnectionStringPersonas = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		"personas")
)

var (
	ConnectionStringArticulos = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		"articulos")
)

const AllowedCORSDomain = "http://localhost"

var Port = ":" + os.Getenv("puerto_propio")

const NombreRuta = "movimientos" //Modificar

var IpApiArticulos = os.Getenv("ip_api_articulos")
var IpApiPersonas = os.Getenv("ip_api_personas")
