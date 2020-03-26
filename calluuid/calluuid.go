// package to send request to uuid IP service and unpack annotation results.
package calluuid

import (
	"context"
	"io/ioutil"
	"log"
	"net"

	"github.com/m-lab/uuid-annotator/ipservice"
)

func CreateClient(ip string) {
	d, err := ioutil.TempDir("", "TestServerAndClientE2E")
	sock := d + "/annotator.sock"
	c := ipservice.NewClient(sock)
	ctx := context.Background()
	got, err := c.Annotate(ctx, net.ParseIP(ip))

	log.Println(got)
	log.Println(err)
}
