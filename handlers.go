package main

import (
	"net/http"
	"time"
)

func makeHTTPServer() *http.Server {
	mux := &http.ServeMux{}

	mux.HandleFunc("/", withAnalyticsLogging(handleMainPage))
	mux.HandleFunc("/favicon.ico", handleFavicon)
	mux.HandleFunc("/robots.txt", handleRobotsTxt)
	mux.HandleFunc("/contactme.html", withAnalyticsLogging(handleContactme))

	mux.HandleFunc("/app/crashsubmit", withAnalyticsLogging(handleCrashSubmit))
	mux.HandleFunc("/app/debug", withAnalyticsLogging(handleDebug))
	mux.HandleFunc("/atom.xml", withAnalyticsLogging(handleAtom))
	mux.HandleFunc("/atom-all.xml", withAnalyticsLogging(handleAtomAll))
	mux.HandleFunc("/archives.html", withAnalyticsLogging(handleArchives))
	mux.HandleFunc("/software", withAnalyticsLogging(handleSoftware))
	mux.HandleFunc("/software/", withAnalyticsLogging(handleSoftware))
	mux.HandleFunc("/extremeoptimizations/", withAnalyticsLogging(handleExtremeOpt))
	mux.HandleFunc("/article/", withAnalyticsLogging(handleArticle))
	mux.HandleFunc("/kb/", withAnalyticsLogging(handleArticle))
	mux.HandleFunc("/blog/", withAnalyticsLogging(handleArticle))
	mux.HandleFunc("/forum_sumatra/", withAnalyticsLogging(forumRedirect))
	mux.HandleFunc("/articles/", withAnalyticsLogging(handleArticles))
	mux.HandleFunc("/tag/", withAnalyticsLogging(handleTag))
	mux.HandleFunc("/static/", withAnalyticsLogging(handleStatic))
	mux.HandleFunc("/css/", handleCss)
	mux.HandleFunc("/js/", handleJs)
	mux.HandleFunc("/gfx/", handleGfx)
	mux.HandleFunc("/djs/", withAnalyticsLogging(handleDjs))
	if !flgProduction {
		mux.HandleFunc("/ws", serveWs)
	}

	// https://blog.gopheracademy.com/advent-2016/exposing-go-on-the-internet/
	srv := &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
	return srv
}
