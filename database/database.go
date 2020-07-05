package database

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/karenirenecano/go-handlers/utils"
)

//DbConn globally declared
var DbConn *sql.DB

//SetupDatabase for referencing db connection
func SetupDatabase() {
	var err error
	rootCertPool := x509.NewCertPool()
	caPemKeyFile, err := utils.GetCWD("/ssl_certs/ca.pem")
	fmt.Println(caPemKeyFile)
	if err != nil {
		log.Fatal(err)
	}
	pem, err := ioutil.ReadFile(caPemKeyFile)
	if err != nil {
		log.Fatal(err)
	}
	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}
	clientCert := make([]tls.Certificate, 0, 1)
	clientCertFile, err := utils.GetCWD("/ssl_certs/client-cert.pem")
	fmt.Println(clientCertFile)

	if err != nil {
		log.Fatal(err)
	}
	clientKeyFile, err := utils.GetCWD("/ssl_certs/client-key.pem")
	fmt.Println(clientKeyFile)

	if err != nil {
		log.Fatal(err)
	}
	certs, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		log.Fatal(err)
	}
	clientCert = append(clientCert, certs)
	mysql.RegisterTLSConfig("custom", &tls.Config{
		RootCAs:            rootCertPool,
		Certificates:       clientCert,
		InsecureSkipVerify: true,
	})
	configSettings := mysql.Config{
		User:                 strings.TrimSpace(os.Getenv("DB_USERNAME")),
		Passwd:               strings.TrimSpace(os.Getenv("DB_PASSWORD")),
		Addr:                 strings.TrimSpace(os.Getenv("DB_ADDRESS")), //IP:PORT
		Net:                  "tcp",
		DBName:               strings.TrimSpace(os.Getenv("DB_NAME")),
		AllowNativePasswords: true,
		TLSConfig:            "custom",
	}
	connectionString := configSettings.FormatDSN()
	DbConn, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}
