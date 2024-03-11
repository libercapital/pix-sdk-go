### Installing
Use `go get` to retrieve the SDK to add it to your `GOPATH` workspace, or
project's Go module dependencies.

	go get github.com/libercapital/pix-sdk-go

### Examples

```
import (
	"log"
	"time"

	pixsdk "github.com/libercapital/pix-sdk-go"
	"github.com/libercapital/pix-sdk-go/services/pix"
)

func main() {
	var pixService = pixsdk.NewService(pixsdk.Config{
		AuthURL:      "https://api.pix.com.br/oauth/token",
		BaseURL:      "https://api.pix.com.br/v2",
		ClientId:     "client-id",
		ClientSecret: "client-secret",
		//Certificate: use 	tls.LoadX509KeyPair() for load certificate if you need,
	})

	// Find Pix by EndToEndId

	px, err := pixService.FindPix("E879247924E724")
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("E2E: %s, TxId: %s, Data: %s, Valor: %f\n", px.E2EId, px.TxId, px.Time.Format(time.RFC3339), px.Value)

	// List Pix by between date
	listPix, err := pixService.ListPix(pix.ListPix{
		pix.StartDate(time.Now().AddDate(0, 0, -1)),
		pix.EndDate(time.Now()),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	for _, px := range listPix.Pix {
		log.Printf("E2E: %s, TxId: %s, Data: %s, Valor: %f\n", px.Pix.E2EId, px.Pix.TxId, px.Pix.Time.Format(time.RFC3339), px.Pix.Value)
	}
}

```