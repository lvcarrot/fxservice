package handler

import (
	"fxlibraries/httpserver"
	"fxlibraries/loggers"
	"fxservice/service/logcenter/adapter"
	versionCmp "github.com/hashicorp/go-version"
)

func PDF2WordVersion(r *httpserver.Request) *httpserver.Response {
	deviceID := r.QueryParams.Get("device_id")
	version := r.QueryParams.Get("version")
	os := r.QueryParams.Get("os")

	loggers.Info.Printf("PDF2Word plugin check update deviceID:%s version:%s os:%s", deviceID, version, os)
	latestVersion, err := adapter.GetLatestPDF2WordVersion()
	if err != nil {
		loggers.Error.Printf("CheckUpdate GetLatestVersion error:%s", err.Error())
		return noNeedUpdate()
	}

	curVer, err := versionCmp.NewVersion(version)
	if err != nil {
		loggers.Error.Printf("CheckUpdate invalid current version:%s", version)
		resp := httpserver.NewResponse()
		resp.Data = latestVersion
		return resp
	}
	latestVer, err := versionCmp.NewVersion(latestVersion.Version)
	if err != nil {
		loggers.Error.Printf("CheckUpdate invalid latest version:%s", latestVersion.Version)
		return noNeedUpdate()
	}

	if latestVer.GreaterThan(curVer) {
		resp := httpserver.NewResponse()
		resp.Data = latestVersion
		return resp
	}

	return noNeedUpdate()
}

func noNeedUpdate() *httpserver.Response {
	resp := httpserver.NewResponse()
	type Resp struct {
		NeedUpdate int `json:"need_update"`
	}
	resp.Data = Resp{0}
	return resp
}