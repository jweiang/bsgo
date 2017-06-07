package main

import (
"log"
   "fmt"
   "net/http"
   "os/exec"
   "io"
)
var cmd *exec.Cmd
const html = `
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>log-view</title>
    <link href="//cdn.bootcss.com/bootstrap/3.3.6/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>
    <div class="container">
        <form method="post" style="margin-top: 20px">
            <div class="form-group">
                <input type="text" class="form-control" name="Value1" placeholder="btoc yizhen btob" required>
            </div>
            <div class="form-group">
                <input type="text" class="form-control" name="Value2" placeholder="server_ip" required>
            </div>
            <div class="form-group">
                <input type="text" class="form-control" name="Value3" placeholder="log path" required>
            </div>
            <button type="submit" class="btn btn-primary center-block">确定</button>
        </form>
    </div>
</body>
</html>
`

func main() {
   http.HandleFunc("/jltg", handler)

   log.Fatal(http.ListenAndServe(":8290", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
if r.Method == "POST" {
      v1 := r.FormValue("Value1")
      v2 := r.FormValue("Value2")
      v3 := r.FormValue("Value3")
switch  {
case v1 == "yizhen":
            cmd = exec.Command("/bin/bash", "/app/shell/yyw_op_logyizhen.sh", v1, v2, v3)
case v1 == "btoc":
            cmd = exec.Command("/bin/bash", "/app/shell/yyw_op_logbtoc.sh", v1, v2,v3)
case v1 == "btob":
            cmd = exec.Command("/bin/bash", "/app/shell/yyw_op_logbtob.sh", v1, v2,v3)
default:
         fmt.Printf("There is no parameters")
      }
      out, err := cmd.CombinedOutput()
if err != nil {
         io.WriteString(w, fmt.Sprintf("Exec command error: %s\nOutput:%s", err.Error(), string(out)))
      } else {
         io.WriteString(w, "Exec command success, Output: " + string(out))
      }
   } else {
      io.WriteString(w, html)
   }
}
