package rosospf

import (
	"bytes"
	"encoding/gob"
	"github.com/BaiMeow/NetworkMonitor/graph/entity"
	"log"
	"reflect"
	"regexp"
	"strconv"

	"github.com/BaiMeow/NetworkMonitor/graph/parse"
	"github.com/go-routeros/routeros/proto"
)

func init() {
	gob.Register(&proto.Sentence{})

	parse.Register("ros-ospf", func(m map[string]any) (parse.Parser[*entity.OSPF], error) {
		return &RosOSPF{}, nil
	})
}

var (
	ros6BodyPtpReg = regexp.MustCompile(`Point-To-Point ((?:[0-9]{1,3}\.){3}[0-9]{1,3}) (?:[0-9]{1,3}\.){3}[0-9]{1,3} (\d+)`)
	ros7BodyPtpReg = regexp.MustCompile(`type=p2p id=((?:[0-9]{1,3}\.){3}[0-9]{1,3}) data=(?:[0-9]{1,3}\.){3}[0-9]{1,3} metric=(\d+)`)
)

var _ parse.Parser[*entity.OSPF] = (*RosOSPF)(nil)

type RosOSPF struct {
	parse.Base[*entity.OSPF]
}

func (p *RosOSPF) Parse(input any) (*entity.OSPF, error) {
	raw, ok := input.([]byte)
	if !ok {
		log.Fatalf("invalid input data type: %s\n", reflect.TypeOf(input).Elem())
	}

	var graph entity.OSPF

	var sentences []*proto.Sentence
	gob.NewDecoder(bytes.NewReader(raw)).Decode(&sentences) // 这里本来应该在初始化就直接处理了，但是因为Init没有抛异常，所以这一步在这里做

	for _, sentence := range sentences {
		if sentence.Word == "!done" { // 这个判断可有可无 因为fetcher已经做了处理
			break
		}
		if sentence.Map["type"] != "router" || sentence.Map["area"] == "" || sentence.Map["id"] == "" {
			continue
		}
		graph.GetArea(sentence.Map["area"]).AddRouter(sentence.Map["id"])

		ptp := ros6BodyPtpReg.FindAllStringSubmatch(sentence.Map["body"], -1)
		for _, field := range ptp {
			if len(field) != 3 {
				continue
			}
			cost, err := strconv.Atoi(field[2])
			if err != nil {
				continue
			}
			graph.AddLink(sentence.Map["area"], sentence.Map["id"], field[1], cost)
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
			graph.AddLink(sentence.Map["area"], sentence.Map["id"], field[1], cost)
		}
	}
	return &graph, nil
}
