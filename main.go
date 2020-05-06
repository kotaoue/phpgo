package main 

import (
        "fmt"
        "flag"
)

func main(){
        f := flag.String("s", "hoge", "Search for php open tag '<?php' at other than head of file.")
        flag.Parse()

        fmt.Println("Hello %s", *f)
}
