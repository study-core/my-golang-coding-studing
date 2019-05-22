package main 

import(
	b64 "encoding/base64"
	"fmt"
)

func main() {



	//decode
	deStr := `
MIIDKDCCAhCgAwIBAgIUOG1dK8XDu4Ei2pVPH1K2FBE80qcwDQYJKoZIhvcNAQEL
BQAwEjEQMA4GA1UEAwwHa3ViZS1jYTAeFw0xNzA5MTgwODQzMDBaFw0xODA5MTgw
ODQzMDBaMBMxETAPBgNVBAMTCHRlcnJlbmNlMIIBIjANBgkqhkiG9w0BAQEFAAOC
AQ8AMIIBCgKCAQEA3OhXR0g1rvwwWRax0E268uiPSnVvm+F5fI/v/AwkcemTTl9L
c+RpWND+VmDtlSIMMfboWcX3zaf84OeC8HA5+yWeKjrbmPjhGDFdeHIm8GubCZzC
0QPO9+lxFK9MmNFC7HxKNfhdgD0WEzzdNlFjIpRI1jpo7hikgSeDIJMob+SCPAOM
3uc5wekdLat3ECYAGHMaiVYJyqefCQ6oSs/D3lRcfGtIoRH5Qi2xBEsGzMOoD9es
wuCzDRzf2KqaxST+GZOl9JIHY/OipLhOeEtENAL/oEdMvjYgc+ltFngcv5iBvQsI
wQAei2efM/pqNZXOLEBg5JbEkRYw3TTKyuD0mQIDAQABo3UwczAOBgNVHQ8BAf8E
BAMCBaAwEwYDVR0lBAwwCgYIKwYBBQUHAwIwDAYDVR0TAQH/BAIwADAdBgNVHQ4E
FgQUchTPLG8isuFzurErTeic3IfKEuswHwYDVR0jBBgwFoAUIGXIViiGYmG1y6jM
7MOlnWxq9hwwDQYJKoZIhvcNAQELBQADggEBAIU/vZaMXM0w9wxTbugaiqVUOokB
HJGAThocEFWHKLnqcV4X7cxTsL50e/C08sYvljw8A3V7Q80GNiVyH3DJ2Yjrn523
VP3MNo92p4r23DuAU3U39hEpybOetflHvRekO2YexFrcP4G3pvKCYhCqZiI3SMlj
UlDo0W+pAoLoSX3i7k97v+6yPtJBpkEg22+6AVwdu36jETmqAsg9czsZLp3Y7HOK
DfVL4lvOvD1Kod/czB2Fuxb3Vw46O6H4torenTF+pk0SvZ16K9ybtFFq/k3l2ZnB
690sCIe2hqBH79BhXXn4xE9g6RTyW1jtrMd3HQmst7DXCmLzM80XGDgPBAc=
`

	sDec2, _ := b64.StdEncoding.DecodeString(deStr)
	fmt.Println("请求头解码结果:", string(sDec2))
}