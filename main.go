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

	uptime := time.Since(startTime).Round(time.Second).String()
	allocMem := mem.Alloc / 1024
	goroutines := runtime.NumGoroutine()

	html := fmt.Sprintf(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Production Microservice Dashboard</title>
    <style>
        * { box-sizing: border-box; margin: 0; padding: 0; }
        body { font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif; background: #0b0f19; color: #f1f5f9; min-height: 100vh; display: flex; align-items: center; justify-content: center; padding: 20px; }
        .card { background: #111827; border: 1px solid #1f2937; border-radius: 16px; padding: 36px; max-width: 650px; width: 100%%; box-shadow: 0 25px 50px -12px rgba(0, 0, 0, 0.7); }
        .status-badge { display: inline-flex; align-items: center; gap: 8px; background: rgba(16, 185, 129, 0.1); color: #10b981; border: 1px solid rgba(16, 185, 129, 0.3); padding: 6px 14px; border-radius: 9999px; font-size: 13px; font-weight: 600; margin-bottom: 20px; }
        .status-dot { width: 8px; height: 8px; background: #10b981; border-radius: 50%%; box-shadow: 0 0 10px #10b981; }
        h1 { font-size: 26px; font-weight: 700; color: #ffffff; margin-bottom: 8px; }
        p.subtitle { color: #94a3b8; font-size: 14px; margin-bottom: 28px; line-height: 1.5; }
        .grid { display: grid; grid-template-columns: repeat(2, 1fr); gap: 14px; margin-bottom: 28px; }
        .metric-box { background: #1f2937; border: 1px solid #374151; padding: 16px; border-radius: 10px; }
        .metric-label { font-size: 11px; text-transform: uppercase; letter-spacing: 0.05em; color: #9ca3af; margin-bottom: 6px; font-weight: 600; }
        .metric-value { font-size: 16px; font-weight: 600; color: #38bdf8; }
        .footer { border-top: 1px solid #1f2937; padding-top: 20px; font-size: 13px; color: #6b7280; text-align: center; }
        .footer strong { color: #e5e7eb; }
    </style>
</head>
<body>
    <div class="card">
        <div class="status-badge">
            <span class="status-dot"></span>
            PRODUCTION SYSTEM ACTIVE & HEALTHY
        </div>
        <h1>Enterprise Cloud Microservice</h1>
        <p class="subtitle">Continuous Delivery Architecture via GitHub Actions, Docker Scratch & AWS EC2</p>
        
        <div class="grid">
            <div class="card-item metric-box">
                <div class="metric-label">Target Cloud Host</div>
                <div class="metric-value">AWS EC2 (Ubuntu 24.04)</div>
            </div>
            <div class="card-item metric-box">
                <div class="metric-label">Container Engine</div>
                <div class="metric-value">Docker Multi-Stage</div>
            </div>
            <div class="card-item metric-box">
                <div class="metric-label">System Uptime</div>
                <div class="metric-value">%s</div>
            </div>
            <div class="card-item metric-box">
                <div class="metric-label">Memory Allocation</div>
                <div class="metric-value">%d KB</div>
            </div>
            <div class="card-item metric-box">
                <div class="metric-label">Active Goroutines</div>
                <div class="metric-value">%d</div>
            </div>
            <div class="card-item metric-box">
                <div class="metric-label">Base OS Image</div>
                <div class="metric-value">Alpine Scratch (Zero Bloat)</div>
            </div>
        </div>

        <div class="footer">
            Engineered & Deployed by <strong>Mohammad Faiz Ansari</strong> • DevOps & Cloud Infrastructure
        </div>
    </div>
</body>
</html>`, uptime, allocMem, goroutines)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(html))
}

func main() {
	http.HandleFunc("/", statusHandler)
	fmt.Println("Server running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
