package cmd

import (
	"github.com/iabzal/parser/config"
	"git.osulta.kz/asia-freight/mailer/pkg/daemon"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
)

func main() {
	// Initialize configs
	viper.SetConfigName("config")
	viper.AddConfigPath("/var/www/microservices/mailer/")
	//viper.AddConfigPath(".")
	var c config.Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	// Initialize logger
	f, err := os.OpenFile("/var/www/microservices/mailer/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//f, err := os.OpenFile("error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)


	//Initialize and start daemon
	d := daemon.NewDaemon(c.Daemon, c.Parse)
	d.Start()
	defer d.Stop()
}
