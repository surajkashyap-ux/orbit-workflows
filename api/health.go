package handler
import("net/http")
func Handler(w http.ResponseWriter,r *http.Request){w.WriteHeader(200);w.Write([]byte("demo_ok"))}
