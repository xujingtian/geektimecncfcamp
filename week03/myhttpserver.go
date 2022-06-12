package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/header", headers)
	mux.HandleFunc("/healthz", healthz)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		s := <-c
		fmt.Printf("System Single call: %+v\n", s)
		cancel()
	}()

	go func() {
		fmt.Println("Server start")
		// When Shutdown is called, Serve, ListenAndServe, and
		// ListenAndServeTLS immediately return ErrServerClosed. Make sure the
		// program doesn't exit and waits instead for Shutdown to return.
		err := server.ListenAndServe()
		fmt.Printf("Server close: %+v\n", err)
	}()

	// 优雅关闭
	// ListenAndServe 会直接返回 所以 我们主程序要改由 Shutdown 阻塞
	<-ctx.Done()
	fmt.Println("Received close single.")
	timeoutCtx, timeoutCancel := context.WithTimeout(context.Background(), time.Second*30)
	defer func() {
		timeoutCancel()
	}()

	if err := server.Shutdown(timeoutCtx); err != nil {
		fmt.Printf("Server Shutdown err; %+v\n", err)
		return
	}
	fmt.Println("Server closed gracefully")

}

func healthz(w http.ResponseWriter, r *http.Request) {
	httpState := http.StatusOK
	fmt.Fprintf(w, "http 200")

	defer func() {
		requestInfo(httpState, r)
	}()
}

func headers(w http.ResponseWriter, r *http.Request) {
	httpState := http.StatusOK
	defer func() {
		requestInfo(httpState, r)
	}()
	v := runtime.Version() //os.Getenv("GOVERSION")

	fmt.Println(v)
	w.Header().Set("GOVERSION", v)
	for name, headers := range r.Header {
		for _, header := range headers {
			w.Header().Add(name, header)
		}
	}
}

func requestInfo(httpState int, r *http.Request) {
	cIP := clientIP(r)
	path := r.URL
	fmt.Printf("[%s] [%d] [%s] client ip: %s\n", time.Now(), httpState, path, cIP)
}

//  参照 gin.clientIP()
func clientIP(r *http.Request) string {
	cIP := r.Header.Get("X-Forwarded-For")
	cIP = strings.TrimSpace(strings.Split(cIP, ",")[0])
	if cIP == "" {
		cIP = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	}
	if cIP != "" {
		return cIP
	}

	if addr := r.Header.Get("X-Appengine-Remote-Addr"); addr != "" {
		return addr
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
