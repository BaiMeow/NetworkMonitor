package rosospf

import (
	"context"
	"encoding/gob"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"github.com/BaiMeow/NetworkMonitor/trace"
	"log"
	"reflect"
	"regexp"
	"strconv"

	"github.com/BaiMeow/NetworkMonitor/graph/parse"
	"github.com/go-routeros/routeros/v3/proto"
)

func init() {
	gob.Register(&proto.Sentence{})

	parse.Register("ros-ospf", func(m map[string]any) (parse.Parser[*entity.OSPF], error) {
		return &RosOSPF{}, nil
	})
}

var (
	ros6BodyPtpReg     = regexp.MustCompile(`Point-To-Point ((?:[0-9]{1,3}\.){3}[0-9]{1,3}) (?:[0-9]{1,3}\.){3}[0-9]{1,3} (\d+)`)
	ros7BodyPtpReg     = regexp.MustCompile(`type=p2p id=((?:[0-9]{1,3}\.){3}[0-9]{1,3}) data=(?:[0-9]{1,3}\.){3}[0-9]{1,3} metric=(\d+)`)
	ros7BodyNetworkReg = regexp.MustCompile(`type=network id=((?:[0-9]{1,3}\.){3}[0-9]{1,3}) data=(?:[0-9]{1,3}\.){3}[0-9]{1,3} metric=(\d+)`)
)

var _ parse.Parser[*entity.OSPF] = (*RosOSPF)(nil)

type RosOSPF struct {
	parse.Base[*entity.OSPF]
}

type network struct {
	cost    int
	area    string
	routers []string
}

func (p *RosOSPF) Parse(ctx context.Context, input any) (*entity.OSPF, error) {
	ctx, span := trace.Tracer.Start(ctx,
		"parse/rosospf/RosOSPF.Parse",
	)
	defer span.End()

	sentences, ok := input.([2][]*proto.Sentence)
	if !ok {
		log.Fatalf("invalid input data type: %s\n", reflect.TypeOf(input).Elem())
	}

	areaSentences := sentences[1]
	areaIdMap := make(map[string]string)
	for _, v := range areaSentences {
		areaIdMap[v.Map["name"]] = v.Map["area-id"]
	}

	networks := make(map[string]*network)
	lsaSentences := sentences[0]
	var graph entity.OSPF
	for _, sentence := range lsaSentences {
		if sentence.Word == "!done" { // 这个判断可有可无 因为fetcher已经做了处理
			break
		}
		if sentence.Map["type"] != "router" || sentence.Map["area"] == "" || sentence.Map["id"] == "" {
			continue
		}
		graph.GetArea(areaIdMap[sentence.Map["area"]]).AddRouter(sentence.Map["id"])

		ptp := ros6BodyPtpReg.FindAllStringSubmatch(sentence.Map["body"], -1)
		for _, field := range ptp {
			if len(field) != 3 {
				continue
			}
			cost, err := strconv.Atoi(field[2])
			if err != nil {
				continue
			}
			graph.AddLink(areaIdMap[sentence.Map["area"]], sentence.Map["id"], field[1], cost)
		}
		ptp = ros7BodyPtpReg.FindAllStringSubmatch(sentence.Map["body"], -1)
		for _, field := range ptp {
			if len(field) != 3 {
				continue
			}
			cost, err := strconv.Atoi(field[2])
			if err != nil {
				continue
			}
			graph.AddLink(areaIdMap[sentence.Map["area"]], sentence.Map["id"], field[1], cost)
		}
		nw := ros7BodyNetworkReg.FindAllStringSubmatch(sentence.Map["body"], -1)
		for _, field := range nw {
			if len(field) != 3 {
				continue
			}
			cost, err := strconv.Atoi(field[2])
			if err != nil {
				continue
			}
			nw := networks[field[1]]
			if nw == nil {
				nw = &network{
					cost: cost,
					area: areaIdMap[sentence.Map["area"]],
				}
				networks[field[1]] = nw
			}
			nw.routers = append(nw.routers, sentence.Map["id"])
		}
	}

	for _, network := range networks {
		area := graph.GetArea(network.area)
		for i := 0; i < len(network.routers); i++ {
			for j := i + 1; j < len(network.routers); j++ {
				area.AddLink(network.routers[i], network.routers[j], network.cost)
				area.AddLink(network.routers[j], network.routers[i], network.cost)
			}
		}
	}

	return &graph, nil
}
