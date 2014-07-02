package controllers

import (
	"github.com/revel/revel"
	"github.com/scy/doohan"
)

type App struct {
	*revel.Controller
}

type stringMap map[string]interface{}

func entryToJson(entry doohan.Entry) stringMap {
	json := make(stringMap, 10)
	json["id"] = entry.ID
	json["start"] = entry.Start.UnixNano()
	if entry.Stop.Valid {
		json["stop"] = entry.Stop.Time.UnixNano()
		json["duration"] = entry.Stop.Time.Sub(entry.Start).Nanoseconds()
	}
	json["running"], json["description"] = entry.Running, entry.Description
	return json
}

func entriesToJson(entries []doohan.Entry) []stringMap {
	var jsons []stringMap
	for _, entry := range entries {
		jsons = append(jsons, entryToJson(entry))
	}
	return jsons
}

func (c App) Index() revel.Result {
	return c.RenderJson(entriesToJson(doohan.FetchEntries()))
}
