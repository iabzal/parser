package daemon

import (
	"github.com/go-co-op/gocron"
	"github.com/iabzal/parser/config"
	"time"
)

type daemon struct {
	c   config.DaemonConfiguration
	p   config.ParseConfiguration
	tws *gocron.Scheduler
	ths *gocron.Scheduler
	hs  *gocron.Scheduler
}

func NewDaemon(configuration config.DaemonConfiguration, parseConfig config.ParseConfiguration) ParserDaemon {
	return &daemon{configuration, parseConfig, gocron.NewScheduler(time.Local), gocron.NewScheduler(time.Local), gocron.NewScheduler(time.Local)}
}

type ParserDaemon interface {
	Start()
	Stop()
}

func (d *daemon) Start() {
	go func() {
		d.tws.Every(uint64(d.c.SendInterval)).Minutes().Do(d.searchTwoRoom)
		<-d.tws.StartAsync()
	}()

	go func() {
		d.ths.Every(uint64(d.c.SendInterval)).Minutes().Do(d.searchThreeRoom)
		<-d.ths.StartAsync()
	}()

	go func() {
		d.hs.Every(uint64(d.c.SendInterval)).Minutes().Do(d.searchHome)
		<-d.hs.StartAsync()
	}()
}

func (d *daemon) Stop() {
	d.tws.Clear()
	d.ths.Clear()
	d.hs.Clear()
}

func (d *daemon) searchTwoRoom() {
	SearchTwoRoom(d.p.UrlTwoRoom)
}

func (d *daemon) searchThreeRoom() {
	SearchThreeRoom(d.p.UrlThreeRoom)

}

func (d *daemon) searchHome() {
	SearchHome(d.p.UrlHome)
}
