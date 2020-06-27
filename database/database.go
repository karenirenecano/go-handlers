package database

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)

//DbConn globally declared
var DbConn *sql.DB

//SetupDatabase for referencing db connection
func SetupDatabase() {
	var err error

	rootCertPool := x509.NewCertPool()
	caPemKeyFile, err := getCWD("/ssl_certs/ca.pem")
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
	clientCertFile, err := getCWD("/ssl_certs/client-cert.pem")
	fmt.Println(clientCertFile)

	if err != nil {
		log.Fatal(err)
	}
	clientKeyFile, err := getCWD("/ssl_certs/client-key.pem")
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
		User:                 "root",
		Passwd:               "root",
		Addr:                 "localhost.co:3306", //IP:PORT
		Net:                  "tcp",
		DBName:               "inventorydb",
		AllowNativePasswords: true,
		TLSConfig:            "custom",
	}
	connectionString := configSettings.FormatDSN()
	DbConn, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}
}

func getCWD(file string) (certDir string, err error) {
	path, errorNotFound := os.Getwd()
	if errorNotFound != nil {
		log.Fatal(errorNotFound)
	}
	fileName := path + file
	_, errorMessage := os.Stat(fileName)
	if os.IsNotExist(errorMessage) {
		return "Not existing", fmt.Errorf("file [%s] does not exist", fileName)
	}

	return fileName, nil
}
