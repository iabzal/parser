package daemon

import (
	"github.com/iabzal/parser/config"
	"github.com/go-co-op/gocron"
	"time"
)

type daemon struct {
	c  config.DaemonConfiguration
	a  config.ParseConfiguration
	tws *gocron.Scheduler
	ths *gocron.Scheduler
	hs *gocron.Scheduler
}

func NewDaemon(configuration config.DaemonConfiguration, parseConfig config.ParseConfiguration) MailerDaemon {
	return &daemon{configuration,parseConfig, gocron.NewScheduler(time.Local), gocron.NewScheduler(time.Local), gocron.NewScheduler(time.Local)}
}

type MailerDaemon interface {
	Start()
	Stop()
	GetRequestForAmwayReport() error
}

func (d *daemon) Start() {
	go func() {
		d.tws.Every(uint64(d.c.SendInterval)).Minutes().Do(searchTwoRoom)
		<-d.tws.StartAsync()
	}()

	go func() {
		d.ths.Every(uint64(d.c.SendInterval)).Minutes().Do(searchThreeRoom)
		<-d.ths.StartAsync()
	}()

	go func() {
		d.hs.Every(uint64(d.c.SendInterval)).Minutes().Do(searchHome)
		<-d.hs.StartAsync()
	}()
}

func (d *daemon) Stop() {
	d.tws.Clear()
	d.ths.Clear()
	d.hs.Clear()
}

func searchTwoRoom()  {

}


func searchThreeRoom()  {

}


func searchHome()  {

}
