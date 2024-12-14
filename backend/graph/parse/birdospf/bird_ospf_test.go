package birdospf

import (
	_ "embed"
	"testing"

	"github.com/BaiMeow/NetworkMonitor/graph/parse"
)

func TestBirdParser(t *testing.T) {
	var p BirdOSPF
	p.asn = 4242424242
	p.Init([]byte(birdOutput))
	var drawing parse.Drawing
	drawing.OSPF = make(map[uint32]*parse.OSPF)
	err := p.ParseAndMerge(&drawing)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(drawing.OSPF[4242424242])
}

var testcase = "area 0.0.0.0\n\n\trouter 172.16.3.91\n\t\tdistance 30\n\t\trouter 172.16.3.254 metric 20\n\n\trouter 172.16.3.120\n\t\tdistance 0\n\t\tnetwork 172.16.3.0/24 metric 10\n\n\trouter 172.16.3.254\n\t\tdistance 10\n\t\trouter 172.16.3.91 metric 20\n\t\tnetwork 172.16.3.0/24 metric 10\n\t\tstubnet 172.16.3.53/32 metric 0\n\t\tstubnet 172.16.255.53/32 metric 0\n\t\tstubnet 172.16.3.13/32 metric 0\n\n\tnetwork 172.16.3.0/24\n\t\tdr 172.16.3.120\n\t\tdistance 10\n\t\trouter 172.16.3.120\n\t\trouter 172.16.3.254\n"

//go:embed testdata/birdospf.txt
var birdOutput string
