package handler
import("net/http";"encoding/json";"time")
func Handler(w http.ResponseWriter,r *http.Request){
if r.Method!=http.MethodPost{http.Error(w,"method",405);return}
id:=r.URL.Query().Get("id")
run:=map[string]interface{}{"id":"run-"+time.Now().UTC().Format("150405"),"workflow_id":id,"status":"completed","created_at":time.Now().UTC(),"result":map[string]interface{}{"message":"done"}}
storeLock.Lock();demoRuns=append(demoRuns,run);storeLock.Unlock()
w.WriteHeader(202);json.NewEncoder(w).Encode(run)
}
