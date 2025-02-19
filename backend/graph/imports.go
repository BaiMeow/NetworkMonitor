package graph

import (
	// fetcher
	_ "github.com/BaiMeow/NetworkMonitor/graph/fetch/bgp"
	_ "github.com/BaiMeow/NetworkMonitor/graph/fetch/birdlggo"
	_ "github.com/BaiMeow/NetworkMonitor/graph/fetch/cmd"
	_ "github.com/BaiMeow/NetworkMonitor/graph/fetch/http"
	_ "github.com/BaiMeow/NetworkMonitor/graph/fetch/ros"
	_ "github.com/BaiMeow/NetworkMonitor/graph/fetch/sftp"
	_ "github.com/BaiMeow/NetworkMonitor/graph/fetch/ssh"
	_ "github.com/BaiMeow/NetworkMonitor/graph/fetch/tcp"

	// parser
	_ "github.com/BaiMeow/NetworkMonitor/graph/parse/bgp"
	_ "github.com/BaiMeow/NetworkMonitor/graph/parse/birdospf"
	_ "github.com/BaiMeow/NetworkMonitor/graph/parse/mtrbgp"
	_ "github.com/BaiMeow/NetworkMonitor/graph/parse/rosospf"
)
