package handler
import("net/http";"encoding/json";"time")
func Handler(w http.ResponseWriter,r *http.Request){
if r.Method==http.MethodPost{
var p map[string]interface{}
json.NewDecoder(r.Body).Decode(&p)
p["id"]="plugin-"+time.Now().UTC().Format("150405")
storeLock.Lock();demoPlugins=append(demoPlugins,p);storeLock.Unlock()
w.WriteHeader(201);json.NewEncoder(w).Encode(p);return}
json.NewEncoder(w).Encode(demoPlugins)
}
