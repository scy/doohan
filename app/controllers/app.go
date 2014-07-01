package controllers

import (
	"github.com/revel/revel"
	"github.com/scy/doohan"
)

type App struct {
	*revel.Controller
}

type stringMap map[string]interface{}

func (c App) Index() revel.Result {
	entries := doohan.FetchEntries()
	var jsons []stringMap
	for _, entry := range entries {
		json := make(stringMap, 10)
		json["id"] = entry.ID
		json["start"] = entry.Start.UnixNano()
		if entry.Stop.Valid {
			json["stop"] = entry.Stop.Time.UnixNano()
			json["duration"] = entry.Stop.Time.Sub(entry.Start).Nanoseconds()
		}
		json["running"], json["description"] = entry.Running, entry.Description
		jsons = append(jsons, json)
	}
	return c.RenderJson(jsons)
}
