package handler
import("sync";"time")
var demoWorkflows=[]map[string]interface{}{{"id":"wf-demo-1","name":"Sample Workflow","created_at":time.Now().UTC()}}
var demoRuns=[]map[string]interface{}{}
var demoPlugins=[]map[string]interface{}{{"id":"plugin-1","name":"Demo Plugin","url":"https://example.com"}}
var storeLock sync.Mutex
