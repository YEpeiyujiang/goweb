package common

import (
	"goweb/config"
	"goweb/models"
	"sync"
)

var Template models.HtmlTemplate

func LoadTemplate() {

	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		//耗时间
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}
