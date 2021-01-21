package libs
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "fmt"
	"os"
	"log"
    "time"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)
var S3 *S3Client


func init(){
	S3=new(S3Client)
}
type S3Client struct {
	Region string
	Sess *session.Session
	Svc *s3.S3
}
func exitErrorf(msg string, args ...interface{}) {
    fmt.Fprintf(os.Stderr, msg+"\n", args...)
}
func (t *S3Client) NewSession(region string){
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region)},
	)

	if err != nil {
		exitErrorf("PROBLEMA DE SESSION CON S3, %v", err)
	}
	t.Sess=sess
	t.Svc = s3.New(t.Sess)
}
/*
lista todo los s3 que puedas ver en la cuenta de aws 
*/
func (t *S3Client) Ls(){
	result, err := t.Svc.ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}
	
	fmt.Println("Buckets:")
	
	for _, b := range result.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
/*
subes un archivo a s3 

filename string archivo local que deseas subir
myBucket string nombre del bucket en tu cuenta de s3
keyName string nombre del objeto final con la ruta completa pero sin el nombre del bucket 
*/

func (t *S3Client) Upload(filename string,myBucket string,keyName string){
	uploader := s3manager.NewUploader(t.Sess)
	f, err  := os.Open(filename)
	if err != nil {
	fmt.Println( fmt.Errorf("failed to open file %q, %v", filename, err))
	}

	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
	Bucket: aws.String(myBucket),
	Key:    aws.String(keyName),
	Body:   f,
	})
	if err != nil {
	fmt.Println( fmt.Errorf("failed to upload file, %v", err))
	}
	fmt.Println(result)
}
/*
	Genera URL publico del objeto que se encuentre en el bucket (myBucket) 
	y que tenga el nombre que se ponga en (keyName)

	myBucket string nombre del bucket en tu cuenta de s3

	keyName string nombre del objeto a descargar con la ruta 
				completa pero sin el nombre del bucket 
*/
func (t *S3Client) GenerateUrl(myBucket string,keyName string) string{
    req, _ := t.Svc.GetObjectRequest(&s3.GetObjectInput{
        Bucket: aws.String(myBucket),
        Key:    aws.String(keyName),
    })
    urlStr, err := req.Presign(15 * time.Minute)

    if err != nil {
        log.Println("Failed to sign request", err)
	}
	fmt.Println(urlStr)
	return urlStr
  
}
