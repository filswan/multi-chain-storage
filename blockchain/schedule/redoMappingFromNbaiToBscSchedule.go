package schedule

import (
	"github.com/robfig/cron"
	"log"
	"payment-bridge/config"
	"payment-bridge/logs"
	"time"
)

func RedoMappingSchedule() {
	c := cron.New()
	c.AddFunc(config.GetConfig().ScheduleRule.Nbai2BscMappingRedoRule, func() {
		log.Println("nbai2Bsc redo schedule is running at " + time.Now().Format("2006-01-02 15:04:05"))
		err := RedoMapping()
		if err != nil {
			logs.GetLogger().Error(err)
		}
	})
	c.Start()
}
