package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"

	"github.com/aalmazanarbs/memstatsbeat/config"
)

type Memstatsbeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
	beat_config := config.DefaultConfig
	if err := cfg.Unpack(&beat_config); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	bt := &Memstatsbeat {
		done:   make(chan struct{}),
		config: beat_config,
	}
	return bt, nil
}

func (bt *Memstatsbeat) Run (b *beat.Beat) error {
	logp.Info("memstatsbeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)

	for {

		m, errorV := mem.VirtualMemory()
		c, errorC := cpu.Info()
		d, errorD := disk.Usage(bt.config.VolumePath)

		if errorV != nil || errorC != nil || errorD != nil || len(c) != 1 {
			return fmt.Errorf("error getting stats: %v %v %v", errorV, errorC, errorD)
		}

		event := common.MapStr {
			"type":            b.Name,
			"@timestamp":      common.Time(time.Now()),
			"cpuModel":        c[0].ModelName,
			"cpuCacheL2GB":    prettyKilobytesNumber(c[0].CacheSize),
			"memAvailableGB":  prettyBytesNumber(m.Available),
			"memUsedGB":       prettyBytesNumber(m.Used),
			"memTotalGB":      prettyBytesNumber(m.Total),
			"diskAvailableGB": prettyBytesNumber(d.Free),
			"diskUsedGB":      prettyBytesNumber(d.Used),
			"diskTotalGB":     prettyBytesNumber(d.Total),
		}

		bt.client.PublishEvent(event)

		select {
			case <-bt.done:
				return nil
			case <-ticker.C:
		}
	}
}

func (bt *Memstatsbeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
