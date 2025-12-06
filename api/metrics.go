package handler
import("net/http")
func Handler(w http.ResponseWriter,r *http.Request){w.Write([]byte("# demo metrics\napi_requests_total 42"))}
