package handler
import("net/http";"encoding/json")
func Handler(w http.ResponseWriter,r *http.Request){
id:=r.URL.Query().Get("id")
for _,rr:=range demoRuns{if rr["id"]==id{json.NewEncoder(w).Encode(rr);return}}
http.Error(w,"notfound",404)
}
