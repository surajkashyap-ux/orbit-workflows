package handler
import("net/http";"encoding/json";"time")
type wfCreateReq struct{Name string `json:"name"`;Spec map[string]interface{} `json:"spec"`}
func Handler(w http.ResponseWriter,r *http.Request){
if r.Method==http.MethodPost{
var req wfCreateReq
json.NewDecoder(r.Body).Decode(&req)
if req.Name==""{http.Error(w,"name required",400);return}
wf:=map[string]interface{}{"id":"wf-demo-"+time.Now().UTC().Format("150405"),"name":req.Name,"spec":req.Spec,"created_at":time.Now().UTC()}
storeLock.Lock();demoWorkflows=append(demoWorkflows,wf);storeLock.Unlock()
w.WriteHeader(201);json.NewEncoder(w).Encode(wf);return}
json.NewEncoder(w).Encode(demoWorkflows)
}
