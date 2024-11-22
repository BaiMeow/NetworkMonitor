package parse

import "testing"

func TestBirdParser(t *testing.T) {
	var p BirdOSPF
	p.asn = 4242424242
	p.Init([]byte(testcase))
	var drawing Drawing
	drawing.OSPF = make(map[uint32]*OSPF)
	err := p.ParseAndMerge(&drawing)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(drawing.OSPF[4242424242])
}

var testcase = "area 0.0.0.0\n\n\trouter 172.16.3.91\n\t\tdistance 30\n\t\trouter 172.16.3.254 metric 20\n\n\trouter 172.16.3.120\n\t\tdistance 0\n\t\tnetwork 172.16.3.0/24 metric 10\n\n\trouter 172.16.3.254\n\t\tdistance 10\n\t\trouter 172.16.3.91 metric 20\n\t\tnetwork 172.16.3.0/24 metric 10\n\t\tstubnet 172.16.3.53/32 metric 0\n\t\tstubnet 172.16.255.53/32 metric 0\n\t\tstubnet 172.16.3.13/32 metric 0\n\n\tnetwork 172.16.3.0/24\n\t\tdr 172.16.3.120\n\t\tdistance 10\n\t\trouter 172.16.3.120\n\t\trouter 172.16.3.254\n"

var birdOutput = `BIRD 2.0.12 ready.

area 0.0.0.0

        router 10.17.255.2
                distance 5
                router 10.17.255.4 metric 5
                router 10.17.255.3 metric 5
                stubnet 10.17.128.50/32 metric 0
                stubnet 10.17.128.51/32 metric 0
                stubnet 10.17.128.49/32 metric 0
                stubnet 10.17.128.52/32 metric 0
                xnetwork 10.17.2.0/24 metric 10
                xnetwork 10.17.128.53/32 metric 0
                external 10.17.128.2/32 metric2 10000
                external 10.17.128.18/32 metric2 10000
                external 10.17.128.35/32 metric2 10000
                external 10.17.128.54/32 metric2 10000
                external 10.17.128.66/32 metric2 10000

        router 10.17.255.3
                distance 5
                router 10.17.255.2 metric 5
                router 10.17.255.4 metric 5
                router 10.17.255.7 metric 5
                stubnet 10.17.128.67/32 metric 0
                stubnet 10.17.128.66/32 metric 0
                stubnet 10.17.128.64/32 metric 0
                stubnet 10.17.128.65/32 metric 0
                stubnet 10.17.128.68/32 metric 0
                stubnet 10.17.128.69/32 metric 0
                stubnet 10.17.3.222/32 metric 0
                xnetwork 10.17.3.0/24 metric 10

        router 10.17.255.4
                distance 0
                router 10.17.255.3 metric 5
                router 10.17.255.2 metric 5
                router 10.17.255.7 metric 5
                router 10.17.255.5 metric 5
                stubnet 10.17.128.21/32 metric 0
                stubnet 10.17.4.0/24 metric 1
                stubnet 10.17.128.19/32 metric 0
                stubnet 10.17.128.18/32 metric 0
                stubnet 10.17.128.20/32 metric 0
                stubnet 10.17.128.17/32 metric 0
                stubnet 10.17.128.16/32 metric 0
                external 10.255.1.1/32 metric2 10000
                external 10.17.128.2/32 metric2 10000
                external 10.17.128.32/32 metric2 10000
                external 10.17.128.50/32 metric2 10000
                external 10.17.128.64/32 metric2 10000
                external 10.17.128.81/32 metric2 10000
                external 10.17.128.96/32 metric2 10000

        router 10.17.255.5
                distance 5
                router 10.17.255.4 metric 5
                stubnet 10.17.128.34/32 metric 0
                stubnet 10.17.128.35/32 metric 0
                stubnet 10.17.128.33/32 metric 0
                stubnet 10.17.128.32/32 metric 0
                xnetwork 10.17.1.0/24 metric 10
                external 10.17.128.3/32 metric2 10000
                external 10.17.128.17/32 metric2 10000
                external 10.17.128.52/32 metric2 10000
                external 10.17.128.65/32 metric2 10000

        router 10.17.255.7
                distance 5
                router 10.17.255.3 metric 5
                router 10.17.255.4 metric 5
                stubnet 10.17.128.97/32 metric 0
                stubnet 10.17.128.96/32 metric 0
                xnetwork 10.17.7.0/24 metric 10
                external 10.17.128.2/32 metric2 10000
                external 10.17.128.69/32 metric2 10000

area 0.0.0.4

        router 10.17.255.4
                distance 0
                xnetwork 10.17.1.0/24 metric 15
                xnetwork 10.17.2.0/24 metric 15
                xnetwork 10.17.3.0/24 metric 15
                xnetwork 10.17.3.222/32 metric 5
                xnetwork 10.17.4.0/24 metric 1
                xnetwork 10.17.7.0/24 metric 15
                xnetwork 10.17.128.16/32 metric 0
                xnetwork 10.17.128.17/32 metric 0
                xnetwork 10.17.128.18/32 metric 0
                xnetwork 10.17.128.19/32 metric 0
                xnetwork 10.17.128.20/32 metric 0
                xnetwork 10.17.128.21/32 metric 0
                xnetwork 10.17.128.32/32 metric 5
                xnetwork 10.17.128.33/32 metric 5
                xnetwork 10.17.128.34/32 metric 5
                xnetwork 10.17.128.35/32 metric 5
                xnetwork 10.17.128.49/32 metric 5
                xnetwork 10.17.128.50/32 metric 5
                xnetwork 10.17.128.51/32 metric 5
                xnetwork 10.17.128.52/32 metric 5
                xnetwork 10.17.128.53/32 metric 5
                xnetwork 10.17.128.64/32 metric 5
                xnetwork 10.17.128.65/32 metric 5
                xnetwork 10.17.128.66/32 metric 5
                xnetwork 10.17.128.67/32 metric 5
                xnetwork 10.17.128.68/32 metric 5
                xnetwork 10.17.128.69/32 metric 5
                xnetwork 10.17.128.96/32 metric 5
                xnetwork 10.17.128.97/32 metric 5
                xrouter 10.17.255.2 metric 5
                xrouter 10.17.255.5 metric 5
                xrouter 10.17.255.7 metric 5
                external 10.255.1.1/32 metric2 10000
                external 10.17.128.2/32 metric2 10000
                external 10.17.128.32/32 metric2 10000
                external 10.17.128.50/32 metric2 10000
                external 10.17.128.64/32 metric2 10000
                external 10.17.128.81/32 metric2 10000
                external 10.17.128.96/32 metric2 10000`
