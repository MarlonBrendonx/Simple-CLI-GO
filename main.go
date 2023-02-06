package main


import (
  "cli-app/app"
  "os"
  "log"
)

func main(){
  app := app.Generate()

  if erro := app.Run(os.Args); erro != nil{

    log.Fatal(erro)
  
  }

}
