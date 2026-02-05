package bgp

import (
	"fmt"
	"testing"
	"time"

	"github.com/BaiMeow/NetworkMonitor/graph/fetch"
	apipb "github.com/osrg/gobgp/v4/api"
)

func TestGoBGPWatch_GetData(t *testing.T) {
	fetcher, err := fetch.Spawn["gobgp-watch"](map[string]any{
		"target": "172.16.4.66:50051",
		"mtls": map[string]any{
			"ca": `-----BEGIN CERTIFICATE-----
MIIBuDCCAV2gAwIBAgIUVnQe5CvJtb13tvG6dmHxyta32pMwCgYIKoZIzj0EAwIw
MDELMAkGA1UEBhMCQ04xDTALBgNVBAoMBEROMTExEjAQBgNVBAMMCWNvbGxlY3Rv
cjAgFw0yNTEwMjExNDQ3NDVaGA8yMTI1MDkyNzE0NDc0NVowMDELMAkGA1UEBhMC
Q04xDTALBgNVBAoMBEROMTExEjAQBgNVBAMMCWNvbGxlY3RvcjBZMBMGByqGSM49
AgEGCCqGSM49AwEHA0IABOPKxK68PPoSu7W+agqsTgHcZzOLMady/s8x9FbRsuMq
lZsy3eQqvIjh7uHpKtRzuvyUUt3tAgwYu/TNnxlNVcejUzBRMB0GA1UdDgQWBBT5
KzSkcMvJEjvKT+aUSpEUPt9m4DAfBgNVHSMEGDAWgBT5KzSkcMvJEjvKT+aUSpEU
Pt9m4DAPBgNVHRMBAf8EBTADAQH/MAoGCCqGSM49BAMCA0kAMEYCIQCYNWsFtmAW
DafzbqMChrJjwBTdkVAY19UDM3Q8DGnYgAIhALpoCK53OkeGDdAGbPpAUASBx7ql
qKafFRy1y81AFcbP
-----END CERTIFICATE-----`,
			"cert": `-----BEGIN CERTIFICATE-----
MIIB5jCCAY2gAwIBAgIURuGYNBPEohEv1T4lR2Geftae+aEwCgYIKoZIzj0EAwIw
MDELMAkGA1UEBhMCQ04xDTALBgNVBAoMBEROMTExEjAQBgNVBAMMCWNvbGxlY3Rv
cjAeFw0yNTEwMjExNTUxMDNaFw0zNTEwMTkxNTUxMDNaMD8xCzAJBgNVBAYTAmNu
MQ0wCwYDVQQLDARkbjExMQ0wCwYDVQQKDARkbjExMRIwEAYDVQQDDAlsb2NhbGhv
c3QwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATsxJFfIKwRRsgXmJ8nge59Cv8/
d0PQ/zr3aMBjCFXWFaWvB+BHZxV8YVe2famqx3omTJ0n4A0pkRJfSyXa7+OQo3Yw
dDAJBgNVHRMEAjAAMAsGA1UdDwQEAwIF4DAaBgNVHREEEzARgglsb2NhbGhvc3SH
BKwQBEIwHQYDVR0OBBYEFK4QVjtP36+VVkaSfsCERzvcohzTMB8GA1UdIwQYMBaA
FPkrNKRwy8kSO8pP5pRKkRQ+32bgMAoGCCqGSM49BAMCA0cAMEQCIB6QQU1q0ZF7
85NWvm1StV2zJeWCpoZc12Q5PJikmRAnAiA8tDT3PPZb0mQAWzmGiUDSWNlrk2vQ
/4O4G+BVhvJaTQ==
-----END CERTIFICATE-----`,
			"key": `-----BEGIN EC PRIVATE KEY-----
MHcCAQEEIEsMV40ImIQpTe5tzAixnbHKwP+YEFXCCL5vc4j0yLlUoAoGCCqGSM49
AwEHoUQDQgAE7MSRXyCsEUbIF5ifJ4HufQr/P3dD0P8692jAYwhV1hWlrwfgR2cV
fGFXtn2pqsd6JkydJ+ANKZESX0sl2u/jkA==
-----END EC PRIVATE KEY-----`,
		},
	})
	if err != nil {
		t.Error(err)
		return
	}
	time.Sleep(time.Second * 10)
	data, err := fetcher.GetData(t.Context())
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(len(data.([]*apipb.Path)))
}
