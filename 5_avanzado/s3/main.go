package main
import (
	"github.com/culturadevops/s3/libs"

)

func main() {
	libs.S3.Region="us-east-1"
	libs.S3.NewSession(libs.S3.Region)
	libs.S3.Ls()
	libs.S3.Upload("subeme.txt","golandia","subido.txt")
	libs.S3.GenerateUrl("golandia","subido.txt")
}