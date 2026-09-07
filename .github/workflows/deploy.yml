package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

var startTime = time.Now()

func statusHandler(w http.ResponseWriter, r *http.Request) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)

	uptime := time.Since(startTime).Round(time.Second)

	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Production Microservice Status</title>
    <style>
        body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif; background: #0f172a; color: #f8fafc; margin: 0; padding: 40px 20px; display: flex; justify-content: center; align-items: center; min-height: 80vh; }
        .container { max-width: 680px; width: 100%%; background: #1e293b; border-radius: 16px; padding: 32px; box-shadow: 0 20px 35px rgba(0,0,0,0.4); border: 1px solid #334155; }
        .badge { background: #10b981; color: #022c22; font-weight: bold; padding: 6px 14px; border-radius: 9999px; font-size: 13px; display: inline-block; margin-bottom: 18px; }
        h1 { margin: 0 0 8px 0; font-size: 24px; color: #ffffff; }
        p.desc { color: #94a3b8; margin: 0 0 24px 0; font-size: 14px; }
        .grid { display: grid; grid-template-columns: repeat(auto-fit, minmax(200px, 1fr)); gap: 14px; margin-bottom: 24px; }
        .card { background: #0f172a; padding: 16px; border-radius: 10px; border: 1px solid #334155; }
        .card span { font-size: 11px; color: #64748b; text-transform: uppercase; font-weight: 700; display: block; margin-bottom: 6px; }
        .card strong { font-size: 15px; color: #38bdf8; }
        .footer { border-top: 1px solid #334155; padding-top: 16px; font-size: 13px; color: #94a3b8; text-align: center; }
    </style>
</head>
<body>
    <div class="container">
        <div class="badge">● PRODUCTION LIVE & HEALTHY</div>
        <h1>Enterprise Cloud Microservice</h1>
        <p class="desc">Automated CI/CD Delivery Pipeline via GitHub Actions & AWS EC2</p>
        <div class="grid">
            <div class="card"><span>Host Environment</span><strong>AWS EC2 (Ubuntu 24.04)</strong></div>
            <div class="card"><span>Container Engine</span><strong>Docker Multi-Stage</strong></div>
            <div class="card"><span>System Uptime</span><strong>%s</strong></div>
            <div class="card"><span>Allocated Memory</span><strong>%d KB</strong></div>
            <div class="card"><span>Active Goroutines</span><strong>%d</strong></div>
            <div class="card"><span>Architecture Target</span><strong>x86_64 Alpine Scratch</strong></div>
        </div>
        <div class="footer">
            Engineered by Mohammad Faiz Ansari • DevOps & Cloud Infrastructure
        </div>
    </div>
</body>
</html>`, uptime.String(), mem.Alloc/1024, runtime.NumGoroutine())

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/", statusHandler)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
