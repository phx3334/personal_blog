package task

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"go_blog/server/global"
)

func RegisterScheduledTasks(c *cron.Cron) error {
	if _, err := c.AddFunc("@hourly", func() {
		if err := UpdateArticleViewsSyncTask(); err != nil {
			global.Log.Error("Failed to update article views:", zap.Error(err))
		}
	}); err != nil {
		return err
	}
	return nil
}