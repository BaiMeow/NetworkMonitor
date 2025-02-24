package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOSPF_Merge(t *testing.T) {
	var testcases = []struct {
		name   string
		input  [2]OSPF
		output OSPF
	}{
		{
			name: "split",
			input: [2]OSPF{
				{
					{
						AreaId: "0",
						Router: []Router{
							*newRouter("10.0.0.1"),
							*newRouter("10.0.0.2"),
						},
						Links: []Link{
							newLink("10.0.0.1", "10.0.0.2", 5),
							newLink("10.0.0.2", "10.0.0.1", 5),
						},
					},
				},
				{
					{
						AreaId: "0",
						Router: []Router{
							*newRouter("10.0.0.3"),
							*newRouter("10.0.0.4"),
						},
						Links: []Link{
							newLink("10.0.0.3", "10.0.0.4", 5),
							newLink("10.0.0.3", "10.0.0.4", 5),
						},
					},
				},
			},
			output: OSPF{
				{
					AreaId: "0",
					Router: []Router{
						*newRouter("10.0.0.1"),
						*newRouter("10.0.0.2"),
						*newRouter("10.0.0.3"),
						*newRouter("10.0.0.4"),
					},
					Links: []Link{
						newLink("10.0.0.1", "10.0.0.2", 5),
						newLink("10.0.0.2", "10.0.0.1", 5),
						newLink("10.0.0.3", "10.0.0.4", 5),
						newLink("10.0.0.3", "10.0.0.4", 5),
					},
				},
			},
		},
		{
			name: "SameAndSkip",
			input: [2]OSPF{
				{
					{
						AreaId: "0",
						Router: []Router{
							*newRouter("10.0.0.1"),
							*newRouter("10.0.0.2"),
						},
						Links: []Link{
							newLink("10.0.0.1", "10.0.0.2", 5),
							newLink("10.0.0.2", "10.0.0.1", 5),
						},
					},
				},
				{
					{
						AreaId: "0",
						Router: []Router{
							*newRouter("10.0.0.1"),
							*newRouter("10.0.0.2"),
						},
						Links: []Link{
							newLink("10.0.0.1", "10.0.0.2", 5),
							newLink("10.0.0.2", "10.0.0.1", 5),
						},
					},
				},
			},
			output: OSPF{
				{
					AreaId: "0",
					Router: []Router{
						*newRouter("10.0.0.1"),
						*newRouter("10.0.0.2"),
					},
					Links: []Link{
						newLink("10.0.0.1", "10.0.0.2", 5),
						newLink("10.0.0.2", "10.0.0.1", 5),
					},
				},
			},
		},
		{
			// should never happen
			name: "LittleDiffAndSkip",
			input: [2]OSPF{
				{
					{
						AreaId: "0",
						Router: []Router{
							*newRouter("10.0.0.1"),
							*newRouter("10.0.0.2"),
						},
						Links: []Link{
							newLink("10.0.0.1", "10.0.0.2", 4),
							newLink("10.0.0.2", "10.0.0.1", 5),
						},
					},
				},
				{
					{
						AreaId: "0",
						Router: []Router{
							*newRouter("10.0.0.1"),
							*newRouter("10.0.0.2"),
						},
						Links: []Link{
							newLink("10.0.0.1", "10.0.0.2", 5),
							newLink("10.0.0.2", "10.0.0.1", 5),
						},
					},
				},
			},
			output: OSPF{
				{
					AreaId: "0",
					Router: []Router{
						*newRouter("10.0.0.1"),
						*newRouter("10.0.0.2"),
					},
					Links: []Link{
						newLink("10.0.0.1", "10.0.0.2", 4),
						newLink("10.0.0.2", "10.0.0.1", 5),
					},
				},
			},
		},
		{
			// should never happen
			name: "FromZero",
			input: [2]OSPF{
				{
					{AreaId: "0"},
				},
				{
					{
						AreaId: "0",
						Router: []Router{
							*newRouter("10.0.0.1"),
							*newRouter("10.0.0.2"),
						},
						Links: []Link{
							newLink("10.0.0.1", "10.0.0.2", 5),
							newLink("10.0.0.2", "10.0.0.1", 5),
						},
					},
				},
			},
			output: OSPF{
				{
					AreaId: "0",
					Router: []Router{
						*newRouter("10.0.0.1"),
						*newRouter("10.0.0.2"),
					},
					Links: []Link{
						newLink("10.0.0.1", "10.0.0.2", 5),
						newLink("10.0.0.2", "10.0.0.1", 5),
					},
				},
			},
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			res := &testcase.input[0]
			res.Merge(&testcase.input[1])
			assert.Equal(t, *res, testcase.output)
		})
	}
}

func TestOSPF_Equal(t *testing.T) {
	var testcases = []struct {
		name   string
		input  [2]OSPF
		output bool
	}{
		{
			name: "split",
			input: [2]OSPF{
				{
					{
						AreaId: "0",
						Router: []Router{
							*newRouter("10.0.0.1"),
							*newRouter("10.0.0.2"),
						},
						Links: []Link{
							newLink("10.0.0.1", "10.0.0.2", 5),
							newLink("10.0.0.2", "10.0.0.1", 5),
						},
					},
				},
				{
					{
						AreaId: "0",
						Router: []Router{
							*newRouter("10.0.0.1"),
							*newRouter("10.0.0.2"),
						},
						Links: []Link{
							newLink("10.0.0.2", "10.0.0.1", 5),
							newLink("10.0.0.1", "10.0.0.2", 5),
						},
					},
				},
			},
			output: true,
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			assert.Equal(t, testcase.output, testcase.input[0].Equal(&testcase.input[1]))
		})
	}
}
