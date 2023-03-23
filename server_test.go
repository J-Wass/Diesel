package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	utils "diesel/utils"

	"github.com/stretchr/testify/assert"
)

func responseForEndpoint(endpoint string) []byte {
	req, _ := http.NewRequest("GET", endpoint, nil)
	httpWriter := httptest.NewRecorder()

	r := setupRouter()
	r.ServeHTTP(httpWriter, req)
	responseData, _ := ioutil.ReadAll(httpWriter.Body)
	return responseData
}

func TestSwiss(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Fall/North_America/Cup"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/swiss/%s", encodedUrl)

	responseData := responseForEndpoint(endpoint)

	expectedEncodedMarkup := "fCoqIyoqfCoqVGVhbXMqKnwqKlctTCoqfCoqUm91bmQgMSoqfCoqUm91bmQgMioqfCoqUm91bmQgMyoqfCoqUm91bmQgNCoqfCoqUm91bmQgNSoqfAp8Oi18Oi18Oi18Oi18Oi18Oi18Oi18Oi18CjEgfFsqKkZBWkUqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvRmFaZV9DbGFuKXwqKjMtMCoqfOKclO+4jyAzLTAgR0d84pyU77iPIDMtMCBBWExFfOKclO+4jyAzLTEgQ09MfHwKMiB8WyoqRlVSKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL0ZVUklBX0VzcG9ydHMpfCoqMy0wKip84pyU77iPIDMtMSBOUkd84pyU77iPIDMtMCBTU0d84pyU77iPIDMtMiBHMnx8CjMgfFsqKk5SRyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9OUkcpfCoqMy0xKip84p2MIDEtMyBGVVJ84pyU77iPIDMtMSBSR0V84pyU77iPIDMtMCBHRU5HfOKclO+4jyAzLTEgRElHfAo0IHxbKipDT0wqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvQ29tcGxleGl0eV9HYW1pbmcpfCoqMy0xKip84pyU77iPIDMtMiBPR3zinJTvuI8gMy0wIEdFTkd84p2MIDEtMyBGQVpFfOKclO+4jyAzLTIgVjF8CjQgfFsqKlNTRyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9TcGFjZXN0YXRpb25fR2FtaW5nKXwqKjMtMSoqfOKclO+4jyAzLTEgUkdFfOKdjCAwLTMgRlVSfOKclO+4jyAzLTAgU1J84pyU77iPIDMtMiBHMnwKNiB8WyoqT0cqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvT3BUaWNfR2FtaW5nKXwqKjMtMioqfOKdjCAyLTMgQ09MfOKclO+4jyAzLTAgRk9CfOKdjCAxLTMgVjF84pyU77iPIDMtMiAyNlJ84pyU77iPIDMtMSBESUcKNyB8WyoqVjEqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvVmVyc2lvbjEpfCoqMy0yKip84pyU77iPIDMtMiBESUd84p2MIDAtMyBHMnzinJTvuI8gMy0xIE9HfOKdjCAyLTMgQ09MfOKclO+4jyAzLTAgU1IKOCB8WyoqR0VORyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9HZW4uR19Nb2JpbDFfUmFjaW5nKXwqKjMtMioqfOKclO+4jyAzLTAgRk9CfOKdjCAwLTMgQ09MfOKdjCAwLTMgTlJHfOKclO+4jyAzLTIgQVhMRXzinJTvuI8gMy0xIEcyCnwtfC0gLSAtIC0gLXwtIC0gLXx8fHx8fAo5IHxbKipHMioqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9HMl9Fc3BvcnRzKXwqKjItMyoqfOKclO+4jyAzLTAgMjZSfOKclO+4jyAzLTAgVjF84p2MIDItMyBGVVJ84p2MIDItMyBTU0d84p2MIDEtMyBHRU5HCjEwIHxbKipESUcqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvRGlnbml0YXMpfCoqMi0zKip84p2MIDItMyBWMXzinJTvuI8gMy0wIDI2UnzinJTvuI8gMy0wIEFYTEV84p2MIDEtMyBOUkd84p2MIDEtMyBPRwoxMSB8WyoqU1IqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvU2hvcGlmeV9SZWJlbGxpb24pfCoqMi0zKip84p2MIDItMyBBWExFfOKclO+4jyAzLTIgR0d84p2MIDAtMyBTU0d84pyU77iPIDMtMSBGT0J84p2MIDAtMyBWMQoxMiB8WyoqMjZSKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlLzI2X1JJU0lORyl8KioxLTMqKnzinYwgMC0zIEcyfOKdjCAwLTMgRElHfOKclO+4jyAzLTEgR0d84p2MIDItMyBPR3wKMTMgfFsqKkZPQioqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9GbGFzaGVzX29mX0JyaWxsaWFuY2UpfCoqMS0zKip84p2MIDAtMyBHRU5HfOKdjCAwLTMgT0d84pyU77iPIDMtMSBSR0V84p2MIDEtMyBTUnwKMTMgfFsqKkFYTEUqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvVGVhbV9BWExFKXwqKjEtMyoqfOKclO+4jyAzLTIgU1J84p2MIDAtMyBGQVpFfOKdjCAwLTMgRElHfOKdjCAyLTMgR0VOR3wKMTUgfFsqKkdHKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL0dob3N0X0dhbWluZyl8KiowLTMqKnzinYwgMC0zIEZBWkV84p2MIDItMyBTUnzinYwgMS0zIDI2Unx8CjE1IHxbKipSR0UqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvUm9ndWUpfCoqMC0zKip84p2MIDEtMyBTU0d84p2MIDEtMyBOUkd84p2MIDEtMyBGT0J8fA=="
	assert.Equal(t, http.StatusOK, 200)
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestBroadcast(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Winter/South_America_Tiebreaker"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/broadcast/%s", encodedUrl)

	responseData := responseForEndpoint(endpoint)

	expectedEncodedMarkup := "fCoqUGxhdGZvcm1zKip8KipMaW5rKip8Cnw6LXw6LXwKfFR3aXRjaHxbKipSb2NrZXRsZWFndWVzYW0qKl0oaHR0cHM6Ly93d3cudHdpdGNoLnR2L3JvY2tldGxlYWd1ZXNhbSl8CnxUd2l0Y2h8WyoqUm9ja2V0c3RyZWV0bGl2ZSoqXShodHRwczovL3d3dy50d2l0Y2gudHYvcm9ja2V0c3RyZWV0bGl2ZSl8CnxZb3V0dWJlfFsqKlJsZXNwb3J0cyoqXShodHRwczovL3d3dy55b3V0dWJlLmNvbS9Acmxlc3BvcnRzKXwK"
	assert.Equal(t, http.StatusOK, 200)
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestBracketWithLeadingZero(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Winter/Middle_East_and_North_Africa/Open"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/bracket/%s/day/3", encodedUrl)

	responseData := responseForEndpoint(endpoint)

	expectedEncodedMarkup := "fCpFTElNSU5BVElPTip8fFsqKkxpcXVpcGVkaWEgQnJhY2tldCoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9Sb2NrZXRfTGVhZ3VlX0NoYW1waW9uc2hpcF9TZXJpZXMvMjAyMi0yMy9XaW50ZXIvTWlkZGxlX0Vhc3RfYW5kX05vcnRoX0FmcmljYS9PcGVuI1Jlc3VsdHMpfAp8Oi18Oi18Oi18CnwqKlJ1bGUgT25lKip8Kio0IC0gMCoqfFZpc2lvbiBFc3BvcnRzfAp8KipUZWFtIEZhbGNvbnMqKnwqKjQgLSAzKip8Q29sYXwKfCoqVGVhbSBGYWxjb25zKip8Kio0IC0gMCoqfFJ1bGUgT25lfA=="
	assert.Equal(t, http.StatusOK, 200)
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestBracket(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Fall/North_America/Cup"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/bracket/%s/day/3", encodedUrl)

	responseData := responseForEndpoint(endpoint)

	expectedEncodedMarkup := "fCpFTElNSU5BVElPTip8fFsqKkxpcXVpcGVkaWEgQnJhY2tldCoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9Sb2NrZXRfTGVhZ3VlX0NoYW1waW9uc2hpcF9TZXJpZXMvMjAyMi0yMy9GYWxsL05vcnRoX0FtZXJpY2EvQ3VwI1Jlc3VsdHMpfAp8Oi18Oi18Oi18CnwqKlZlcnNpb24xKip8Kio0IC0gMSoqfE5SR3wKfCoqR2VuLkcgTW9iaWwxIFJhY2luZyoqfCoqNCAtIDEqKnxTcGFjZXN0YXRpb258CnxHZW4uRyBNb2JpbDEgUmFjaW5nfCoqMyAtIDQqKnwqKlZlcnNpb24xKip8"
	assert.Equal(t, http.StatusOK, 200)
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestTitle(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Fall/North_America/Cup"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/title/%s", encodedUrl)

	responseData := responseForEndpoint(endpoint)

	expectedEncodedMarkup := "UkxDUyAyMDIyLTIzIC0gRmFsbDogTkEgUmVnaW9uYWwgMiAtIEZhbGwgQ3Vw"
	assert.Equal(t, http.StatusOK, 200)
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestMakeThread(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Winter/Asia-Pacific/Cup"
	encodedUrl := utils.EncodedBase64(url)
	template := utils.EncodedBase64("groups-bracketrd1-streams")
	endpoint := fmt.Sprintf("/makethread/%s/template/%s/day/1", encodedUrl, template)

	responseData := responseForEndpoint(endpoint)

	expectedEncodedMarkup := "IyBSTENTIDIwMjItMjMgLSBXaW50ZXI6IEFQQUMgUmVnaW9uYWwgMiAtIFdpbnRlciBDdXANCnx8KipEYXkqKnwqKlVUQyoqfHwKfDotfDotfDotfDotfAp8RGF5IDF8U2F0dXJkYXl8WyoqMTA6MDAqKl0oaHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS9zZWFyY2g/cT0xMDowMCtHTVQpfCoqVG9kYXkqKnwKfERheSAyfFN1bmRheXxbKioxMDowMCoqXShodHRwczovL3d3dy5nb29nbGUuY29tL3NlYXJjaD9xPTEwOjAwK0dNVCl8fA0KDQojIENvdmVyYWdlDQpbKipMaXF1aXBlZGlhKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1JvY2tldF9MZWFndWVfQ2hhbXBpb25zaGlwX1Nlcmllcy8yMDIyLTIzL1dpbnRlci9Bc2lhLVBhY2lmaWMvQ3VwKSAqKi8gLyoqIFsqKlN0YXJ0LmdnKipdKGh0dHBzOi8vd3d3LnN0YXJ0LmdnL3RvdXJuYW1lbnQvcmxjcy0yMDIyLTIzLXdpbnRlci1jdXAtYXBhYy9ldmVudC9tYWluLWV2ZW50KQ0KDQojIFN0cmVhbXMNCnwqKlBsYXRmb3JtcyoqfCoqTGluayoqfAp8Oi18Oi18CnxUd2l0Y2h8WyoqUm9ja2V0bGVhZ3VlYXBhYyoqXShodHRwczovL3d3dy50d2l0Y2gudHYvcm9ja2V0bGVhZ3VlYXBhYyl8CnxZb3V0dWJlfFsqKlJsZXNwb3J0cyoqXShodHRwczovL3d3dy55b3V0dWJlLmNvbS9Acmxlc3BvcnRzKXwKDQpUaGUgbWFpbiBjaGFubmVsIHdpbGwgc2hvdyBhIGZlYXR1cmVkIG1hdGNoIGZyb20gZWFjaCByb3VuZCwgd2hpbGUgdGhlIHBhcnRpY2lwYXRpbmcgdGVhbXMgbWF5IGJlIHJ1bm5pbmcgYSBzdHJlYW0gZm9yIHRoZWlyIG93biBtYXRjaGVzLiBCZWxvdyBpcyBhIExpc3Qgb2YgVGVhbSBCcm9hZGNhc3RzOg0KDQp8fHx8fAp8Oi18Oi18Oi18Oi18CnxbRGFyayBSaWZ0IEVzcG9ydHNdKGh0dHBzOi8vd3d3LnR3aXRjaC50di9EYXJrUmlmdEVzcG9ydHMpfFtFbGV2YXRlXShodHRwczovL3d3dy50d2l0Y2gudHYvZWxldmF0ZWdnKXxbR2FpbWluIEdsYWRpYXRvcnNdKGh0dHBzOi8vd3d3LnR3aXRjaC50di9nYWltaW5nbGFkaWF0b3JzdHYpfFtMb3R1cyBHYW1pbmddKGh0dHBzOi8vd3d3LnR3aXRjaC50di9sb3R1cyBnYW1pbmcxKQp8W0x5Y3VzIEVtcGlyZV0oaHR0cHM6Ly93d3cudHdpdGNoLnR2L0x5Y3VzRW1waXJlKXxbTmltbXQ1NV0oaHR0cHM6Ly93d3cudHdpdGNoLnR2L0hhenphUkwpfFtVSFVIXShodHRwczovL3d3dy50d2l0Y2gudHYvdGFudHJpa2VzcG9ydHMpfFtWYWluc2hvZXNdKGh0dHBzOi8vd3d3LnR3aXRjaC50di9hbWF0ZWxsbCkNCg0KIyBSZXN1bHRzDQp8fHx8fAp8Oi18Oi18Oi18Oi18CnwqKiMqKnwqKkdyb3VwIEEqKiAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyB8KipNYXRjaGVzKiogfCoqR2FtZSBEaWZmKiogfAp8MXxbKipFbGV2YXRlKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL0VsZXZhdGUpfDMtMHwrOXwKfDJ8WyoqVUhVSCoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9VSFVIKXwyLTF8KzN8CnwzfFsqKkFUSyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9BVEtfKEphcGFuZXNlX1RlYW0pKXwxLTJ8LTN8Cnw0fFsqKlRlYW0gR2FuRGVyUyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9UZWFtX0dhbkRlclMpfDAtM3wtOXwKCiYjeDIwMEI7Cgp8fHx8fAp8Oi18Oi18Oi18Oi18CnwqKiMqKnwqKkdyb3VwIEIqKiAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyB8KipNYXRjaGVzKiogfCoqR2FtZSBEaWZmKiogfAp8MXxbKipEZVRvTmF0b3IqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvRGVUb05hdG9yKXwzLTB8Kzl8CnwyfFsqKkJyb2tvbGkqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvQnJva29saSl8Mi0xfCsyfAp8M3xbKipBbW9uZ1oqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvQW1vbmdaKXwxLTJ8LTN8Cnw0fFsqKlJlemkncyBNaW5pb25zKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1JlemklMjdzX01pbmlvbnMpfDAtM3wtOHwKCiYjeDIwMEI7Cgp8fHx8fAp8Oi18Oi18Oi18Oi18CnwqKiMqKnwqKkdyb3VwIEMqKiAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyB8KipNYXRjaGVzKiogfCoqR2FtZSBEaWZmKiogfAp8MXxbKipOaW1tdDU1KipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL05pbW10NTUpfDMtMHwrNnwKfDJ8WyoqTG90dXMgR2FtaW5nKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL0xvdHVzX0dhbWluZyl8Mi0xfCszfAp8M3xbKipMeWN1cyBFbXBpcmUqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvTHljdXNfRW1waXJlKXwxLTJ8MHwKfDR8WyoqR29kYWxpb25zKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL0dvZGFsaW9ucyl8MC0zfC05fAoKJiN4MjAwQjsKCnx8fHx8Cnw6LXw6LXw6LXw6LXwKfCoqIyoqfCoqR3JvdXAgRCoqICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7IHwqKk1hdGNoZXMqKiB8KipHYW1lIERpZmYqKiB8CnwxfFsqKkRlY2ltYXRlIEdhbWluZyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9EZWNpbWF0ZV9HYW1pbmcpfDMtMHwrOHwKfDJ8WyoqR2FpbWluIEdsYWRpYXRvcnMqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvR2FpbWluX0dsYWRpYXRvcnMpfDItMXwrM3wKfDN8WyoqVmFpbnNob2VzKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1ZhaW5zaG9lcyl8MS0yfC00fAp8NHxbKipEYXJrIFJpZnQgRXNwb3J0cyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9EYXJrX1JpZnRfRXNwb3J0cyl8MC0zfC03fAoKJiN4MjAwQjsKCg0KDQojIFBsYXlvZmZzIC0gUm91bmQgMQ0KfCpFTElNSU5BVElPTip8fFsqKkxpcXVpcGVkaWEgQnJhY2tldCoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9Sb2NrZXRfTGVhZ3VlX0NoYW1waW9uc2hpcF9TZXJpZXMvMjAyMi0yMy9XaW50ZXIvQXNpYS1QYWNpZmljL0N1cCNSZXN1bHRzKXwKfDotfDotfDotfAp8KipCcm9rb2xpKip8KiozIC0gMCoqfFZhaW5zaG9lc3wKfCoqTG90dXMgR2FtaW5nKip8KiozIC0gMSoqfEFUS3wKfFVIVUh8KiowIC0gMyoqfCoqTHljdXMgRW1waXJlKip8CnwqKkdsYWRpYXRvcnMqKnwqKjMgLSAwKip8QW1vbmdafA=="
	assert.Equal(t, http.StatusOK, 200)
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestTemplates(t *testing.T) {
	endpoint := "/templates"

	responseData := responseForEndpoint(endpoint)

	expectedEncodedMarkup := "YnJhY2tldCwgYnJhY2tldC1wcml6ZXBvb2wsIGJyYWNrZXQtc3RyZWFtcywgZ3JvdXBzLCBncm91cHMtYnJhY2tldHJkMSwgZ3JvdXBzLWJyYWNrZXRyZDEtc3RyZWFtcywgZ3JvdXBzLXByaXplcG9vbCwgZ3JvdXBzLXN0cmVhbXMsIHN3aXNzLCBzd2lzcy1icmFja2V0cmQxLCBzd2lzcy1icmFja2V0cmQxLXN0cmVhbXMsIHN3aXNzLXByaXplcG9vbCwgc3dpc3Mtc3RyZWFtcw=="
	assert.Equal(t, http.StatusOK, 200)
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestPrizepool(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Winter/Middle_East_and_North_Africa/Open"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/prizepool/%s", encodedUrl)

	responseData := responseForEndpoint(endpoint)

	expectedEncodedMarkup := "fCoqUGxhY2UqKnwqKlByaXplKip8KipUZWFtKip8KipSTENTIFBvaW50cyoqfAp8Oi18Oi18Oi18Oi18CnwqKjFzdCoqfCQ5LDAwMHxUZWFtIEZhbGNvbnN8KzIwfAp8KioybmQqKnwkNiwwMDB8UnVsZSBPbmV8KzE2fAp8KiozcmQtNHRoKip8JDMsOTAwfENvbGF8KzEyfAp8KiozcmQtNHRoKip8JDMsOTAwfFZpc2lvbiBFc3BvcnRzfCsxMnwKfCoqNXRoLTh0aCoqfCQxLDgwMHxBcnJvd3N8Kzh8CnwqKjV0aC04dGgqKnwkMSw4MDB8RU1QVFl8Kzh8CnwqKjV0aC04dGgqKnwkMSw4MDB8VHJvdWJsZXN8Kzh8CnwqKjV0aC04dGgqKnwkMSw4MDB8VHdpc3RlZCBNaW5kc3wrOHw="
	assert.Equal(t, http.StatusOK, 200)
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestMVP(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Fall/North_America/Cup"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/mvp_candidates/%s/teams_allowed/4", encodedUrl)

	responseData := responseForEndpoint(endpoint)

	expectedEncodedMarkup := "dG9ybWVudCAoVmVyc2lvbjEpCmNvbW0gKFZlcnNpb24xKQpCZWFzdE1vZGUgKFZlcnNpb24xKQpBcHBhcmVudGx5SmFjayAoR2VuLkcgTW9iaWwxIFJhY2luZykKQ2hyb25pYyAoR2VuLkcgTW9iaWwxIFJhY2luZykKbm9seSAoR2VuLkcgTW9iaWwxIFJhY2luZykKQXJzZW5hbCAoU3BhY2VzdGF0aW9uIEdhbWluZykKRGFuaWVsIChTcGFjZXN0YXRpb24gR2FtaW5nKQpMaiAoU3BhY2VzdGF0aW9uIEdhbWluZykKR2FycmV0dEcgKE5SRykKanVzdGluLiAoTlJHKQpTcXVpc2h5IChOUkcp"
	assert.Equal(t, http.StatusOK, 200)
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestStreams(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Fall/North_America/Cup"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/streams/%s", encodedUrl)

	responseData := responseForEndpoint(endpoint)

	expectedEncodedMarkup := "fHx8fHwKfDotfDotfDotfDotfAp8WzI2IFJJU0lOR10oaHR0cHM6Ly93d3cudHdpdGNoLnR2LzI2cmlzaW5nKXxbVGVhbSBBWExFXShodHRwczovL3d3dy50d2l0Y2gudHYvdGVhbWF4bGVyOCl8W0NvbXBsZXhpdHkgR2FtaW5nXShodHRwczovL3d3dy50d2l0Y2gudHYvY29tcGxleGl0eSl8W0RpZ25pdGFzXShodHRwczovL3d3dy50d2l0Y2gudHYvRGlnbml0YXMpCnxbRmFaZSBDbGFuXShodHRwczovL3d3dy50d2l0Y2gudHYvQ2l6em9yeil8W0ZsYXNoZXMgb2YgQnJpbGxpYW5jZV0oaHR0cHM6Ly93d3cudHdpdGNoLnR2L01hclRoZUdhbWVyTW9tKXxbRlVSSUEgRXNwb3J0c10oaHR0cHM6Ly93d3cudHdpdGNoLnR2L0ZVUklBdHYpfFtHMiBFc3BvcnRzXShodHRwczovL3d3dy50d2l0Y2gudHYvZzJlc3BvcnRzKQp8W0dlbi5HIE1vYmlsMSBSYWNpbmddKGh0dHBzOi8vd3d3LnR3aXRjaC50di93aWRvdyl8W0dob3N0IEdhbWluZ10oaHR0cHM6Ly93d3cudHdpdGNoLnR2L3RlbmFjaXR5dHYpfFtOUkddKGh0dHBzOi8vd3d3LnR3aXRjaC50di9ucmdnZyl8W09wVGljIEdhbWluZ10oaHR0cHM6Ly93d3cudHdpdGNoLnR2L2hpdGNoYXJpaWRlKQp8W1JvZ3VlXShodHRwczovL3d3dy50d2l0Y2gudHYvUm9ndWUpfFtTaG9waWZ5IFJlYmVsbGlvbl0oaHR0cHM6Ly93d3cudHdpdGNoLnR2L3Nob3BpZnlyZWJlbGxpb24pfFtTcGFjZXN0YXRpb24gR2FtaW5nXShodHRwczovL3d3dy50d2l0Y2gudHYvc3BhY2VzdGF0aW9uKXxbVmVyc2lvbjFdKGh0dHBzOi8vd3d3LnR3aXRjaC50di92ZXJzaW9uMSk="
	assert.Equal(t, http.StatusOK, 200)
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestCoverage(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Fall/North_America/Cup"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/coverage/%s", encodedUrl)

	responseData := responseForEndpoint(endpoint)

	expectedEncodedMarkup := "WyoqTGlxdWlwZWRpYSoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9Sb2NrZXRfTGVhZ3VlX0NoYW1waW9uc2hpcF9TZXJpZXMvMjAyMi0yMy9GYWxsL05vcnRoX0FtZXJpY2EvQ3VwKSAqKi8gLyoqIFsqKlN0YXJ0LmdnKipdKGh0dHBzOi8vd3d3LnN0YXJ0LmdnL3RvdXJuYW1lbnQvcmxjcy0yMDIyLTIzLWZhbGwtY3VwLW5vcnRoLWFtZXJpY2EvZXZlbnQvbWFpbi1ldmVudCk="
	assert.Equal(t, http.StatusOK, 200)
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestGroups(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2021-22/Winter/Middle_East_and_North_Africa/2"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/groups/%s", encodedUrl)

	responseData := responseForEndpoint(endpoint)

	expectedEncodedMarkup := "fHx8fHwKfDotfDotfDotfDotfAp8KiojKip8KipHcm91cCBBKiogJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgfCoqTWF0Y2hlcyoqIHwqKkdhbWUgRGlmZioqIHwKfDF8WyoqU2FuZHJvY2sgR2FtaW5nKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1NhbmRyb2NrX0dhbWluZyl8My0wfCs5fAp8MnxbKipUdWNoZWwqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvVHVjaGVsKXwyLTF8MHwKfDN8WyoqRm94IEdhbWluZyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9Gb3hfR2FtaW5nKXwxLTJ8LTF8Cnw0fFsqKjI1ZVNwb3J0cyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS8yNWVTcG9ydHMpfDAtM3wtOHwKCiYjeDIwMEI7Cgp8fHx8fAp8Oi18Oi18Oi18Oi18CnwqKiMqKnwqKkdyb3VwIEIqKiAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyB8KipNYXRjaGVzKiogfCoqR2FtZSBEaWZmKiogfAp8MXxbKipUZWFtIEZhbGNvbnMqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvVGVhbV9GYWxjb25zKXwzLTB8Kzh8CnwyfFsqKkZvcmVzdCBIdW50ZXJzKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL0ZvcmVzdF9IdW50ZXJzKXwyLTF8KzJ8CnwzfFsqKlRoZSBFdmlsIGVTcG9ydHMqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvVGhlX0V2aWxfZVNwb3J0cyl8MS0yfC0yfAp8NHxbKipZVCBJbW1vcnRhbHMqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvaW5kZXgucGhwP3RpdGxlPVlvdSUyN3JlX1Rocm93aW5nX0ltbW9ydGFscyZhY3Rpb249ZWRpdCZyZWRsaW5rPTEpfDAtM3wtOHwKCiYjeDIwMEI7Cgp8fHx8fAp8Oi18Oi18Oi18Oi18CnwqKiMqKnwqKkdyb3VwIEMqKiAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyB8KipNYXRjaGVzKiogfCoqR2FtZSBEaWZmKiogfAp8MXxbKipUaGUgVWx0aW1hdGVzKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1RoZV9VbHRpbWF0ZXMpfDItMXwrM3wKfDJ8WyoqUmV2b2x1dGlvbioqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9SZXZvbHV0aW9uKXwyLTF8KzR8CnwzfFsqKlQySyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9UMkspfDItMXwtMXwKfDR8WyoqV29sdmVzKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1dvbHZlc18oTWlkZGxlX0Vhc3Rlcm5fVGVhbSkpfDAtM3wtNnwKCiYjeDIwMEI7Cgp8fHx8fAp8Oi18Oi18Oi18Oi18CnwqKiMqKnwqKkdyb3VwIEQqKiAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyB8KipNYXRjaGVzKiogfCoqR2FtZSBEaWZmKiogfAp8MXxbKipTQ1lURVMqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvU0NZVEVTKXwzLTB8KzZ8CnwyfFsqKkFOS0FBIEVzcG9ydHMqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvQU5LQUFfRXNwb3J0cyl8Mi0xfCs1fAp8M3xbKipLSU5HUyBFc3BvcnRzKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL0tJTkdTX0VzcG9ydHMpfDEtMnwtM3wKfDR8WyoqTmlnaHRtYXJlKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL05pZ2h0bWFyZSl8MC0zfC04fAoKJiN4MjAwQjsKCg=="
	assert.Equal(t, http.StatusOK, 200)
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestSchedule(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Fall/North_America/Cup"
	encodedUrl := utils.EncodedBase64(url)

	// Day 1
	endpoint := fmt.Sprintf("/schedule/%s/day/1", encodedUrl)
	responseData := responseForEndpoint(endpoint)

	expectedEncodedMarkup := "fHwqKkRheSoqfCoqVVRDKip8fAp8Oi18Oi18Oi18Oi18CnxEYXkgMXxGcmlkYXl8WyoqMTc6MDAqKl0oaHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS9zZWFyY2g/cT0xNzowMCtHTVQpfCoqVG9kYXkqKnwKfERheSAyfFNhdHVyZGF5fFsqKjE3OjAwKipdKGh0dHBzOi8vd3d3Lmdvb2dsZS5jb20vc2VhcmNoP3E9MTc6MDArR01UKXx8CnxEYXkgM3xTdW5kYXl8WyoqMTc6MDAqKl0oaHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS9zZWFyY2g/cT0xNzowMCtHTVQpfHw="
	assert.Equal(t, http.StatusOK, 200)
	assert.Equal(t, expectedEncodedMarkup, string(responseData))

	// Day 3
	endpoint = fmt.Sprintf("/schedule/%s/day/3", encodedUrl)
	responseData = responseForEndpoint(endpoint)

	expectedEncodedMarkup = "fHwqKkRheSoqfCoqVVRDKip8fAp8Oi18Oi18Oi18Oi18Cnx+fkRheSAxfn58fn5GcmlkYXl+fnx+flsqKjE3OjAwKipdKGh0dHBzOi8vd3d3Lmdvb2dsZS5jb20vc2VhcmNoP3E9MTc6MDArR01UKX5+fHwKfH5+RGF5IDJ+fnx+flNhdHVyZGF5fn58fn5bKioxNzowMCoqXShodHRwczovL3d3dy5nb29nbGUuY29tL3NlYXJjaD9xPTE3OjAwK0dNVCl+fnx8CnxEYXkgM3xTdW5kYXl8WyoqMTc6MDAqKl0oaHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS9zZWFyY2g/cT0xNzowMCtHTVQpfCoqVG9kYXkqKnw="
	assert.Equal(t, http.StatusOK, 200)
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}
