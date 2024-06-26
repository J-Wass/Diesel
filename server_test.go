package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	utils "diesel/utils"

	"github.com/stretchr/testify/assert"
)

// Send a GET request to the specified endpoint, waits for the request to finish, and returns the response body.
func responseForEndpoint(t *testing.T, endpoint string) string {
	router := setupRouter()
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", endpoint, nil)
	done := make(chan bool)
	go func() {
		router.ServeHTTP(w, r)
		done <- true
	}()
	<-done
	buf := new(bytes.Buffer)
	io.Copy(buf, w.Body)
	
	assert.Equal(t, w.Result().StatusCode, 200)
	return buf.String()
}

func TestSwiss(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2024/Major_1/North_America/Open_Qualifier_1"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/swiss/%s", encodedUrl)

	responseData := responseForEndpoint(t, endpoint)

	expectedEncodedMarkup := "fCoqIyoqfCoqVGVhbXMqKnwqKlctTCoqfCoqUm91bmQgMSoqfCoqUm91bmQgMioqfCoqUm91bmQgMyoqfCoqUm91bmQgNCoqfCoqUm91bmQgNSoqfAp8Oi18Oi18Oi18Oi18Oi18Oi18Oi18Oi18CjEgfFsqKk04MCoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9NODApfCoqMy0wKip84pyU77iPIDMtMCBNVUZ84pyU77iPIDMtMiBUU0184pyU77iPIDMtMSBHRU5HfHwKMiB8WyoqRzIgU3RyaWRlKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL0cyX0VzcG9ydHMpfCoqMy0wKip84pyU77iPIDMtMiBTTk9XfOKclO+4jyAzLTIgTEd84pyU77iPIDMtMSBTU0d8fAozIHxbKipESUcqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvRGlnbml0YXMpfCoqMy0xKip84p2MIDAtMyBTUnzinJTvuI8gMy0wIE9HfOKclO+4jyAzLTAgVFNNfOKclO+4jyAzLTAgTEd8CjQgfFsqKkdFTkcqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvR2VuLkdfTW9iaWwxX1JhY2luZyl8KiozLTEqKnzinJTvuI8gMy0xIE9HfOKclO+4jyAzLTEgU1J84p2MIDEtMyBNODB84pyU77iPIDMtMCBZV1N8CjUgfFsqKlNTRyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9TcGFjZXN0YXRpb25fR2FtaW5nKXwqKjMtMSoqfOKclO+4jyAzLTAgUE9BQnzinJTvuI8gMy0xIFlXU3zinYwgMS0zIEcyIFN0cmlkZXzinJTvuI8gMy0yIFNSfAo2IHxbKipQT0FCKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1BpcmF0ZXNfb25fYV9Cb2F0KXwqKjMtMioqfOKdjCAwLTMgU1NHfOKdjCAyLTMgTlJHfOKclO+4jyAzLTAgT0d84pyU77iPIDMtMSBUU0184pyU77iPIDMtMCBZV1MKNyB8WyoqTEcqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvTHVtaW5vc2l0eV9HYW1pbmcpfCoqMy0yKip84pyU77iPIDMtMCBQVHzinYwgMi0zIEcyIFN0cmlkZXzinJTvuI8gMy0yIE1VRnzinYwgMC0zIERJR3zinJTvuI8gMy0wIE5SRwo4IHxbKipTUioqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9TaG9waWZ5X1JlYmVsbGlvbil8KiozLTIqKnzinJTvuI8gMy0wIERJR3zinYwgMS0zIEdFTkd84pyU77iPIDMtMiBOUkd84p2MIDItMyBTU0d84pyU77iPIDMtMiBTTk9XCnwtfC0gLSAtIC0gLXwtIC0gLXx8fHx8fAo5IHxbKipTTk9XKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1RoZV9Tbm93bWVuKXwqKjItMyoqfOKdjCAyLTMgRzIgU3RyaWRlfOKclO+4jyAzLTAgUFR84p2MIDItMyBZV1N84pyU77iPIDMtMiBYRHzinYwgMi0zIFNSCjEwIHxbKipOUkcqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvTlJHKXwqKjItMyoqfOKdjCAyLTMgWVdTfOKclO+4jyAzLTIgUE9BQnzinYwgMi0zIFNSfOKclO+4jyAzLTAgTVVGfOKdjCAwLTMgTEcKMTEgfFsqKllXUyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9Zb3VuZ19XaGlwcGVyc25hcHBlcnMpfCoqMi0zKip84pyU77iPIDMtMiBOUkd84p2MIDEtMyBTU0d84pyU77iPIDMtMiBTTk9XfOKdjCAwLTMgR0VOR3zinYwgMC0zIFBPQUIKMTIgfFsqKlhEKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL0RFTEVURURfWEQpfCoqMS0zKip84p2MIDItMyBUU0184p2MIDItMyBNVUZ84pyU77iPIDMtMiBQVHzinYwgMi0zIFNOT1d8CjEzIHxbKipUU00qKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvVFNNKXwqKjEtMyoqfOKclO+4jyAzLTIgWER84p2MIDItMyBNODB84p2MIDAtMyBESUd84p2MIDEtMyBQT0FCfAoxNCB8WyoqTVVGKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1RoZV9NdWZmaW5fTWVuKXwqKjEtMyoqfOKdjCAwLTMgTTgwfOKclO+4jyAzLTIgWER84p2MIDItMyBMR3zinYwgMC0zIE5SR3wKMTUgfFsqKlBUKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1Bsb3RfVHdpc3QpfCoqMC0zKip84p2MIDAtMyBMR3zinYwgMC0zIFNOT1d84p2MIDItMyBYRHx8CjE2IHxbKipPRyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9PRyl8KiowLTMqKnzinYwgMS0zIEdFTkd84p2MIDAtMyBESUd84p2MIDAtMyBQT0FCfHw="
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestBroadcast(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Winter/South_America_Tiebreaker"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/broadcast/%s", encodedUrl)

	responseData := responseForEndpoint(t, endpoint)

	expectedEncodedMarkup := "fCoqUGxhdGZvcm1zKip8KipMaW5rKip8Cnw6LXw6LXwKfFR3aXRjaHxbKipSb2NrZXRsZWFndWVzYW0qKl0oaHR0cHM6Ly93d3cudHdpdGNoLnR2L3JvY2tldGxlYWd1ZXNhbSl8CnxUd2l0Y2h8WyoqUm9ja2V0c3RyZWV0bGl2ZSoqXShodHRwczovL3d3dy50d2l0Y2gudHYvcm9ja2V0c3RyZWV0bGl2ZSl8CnxZb3V0dWJlfFsqKlJsZXNwb3J0cyoqXShodHRwczovL3d3dy55b3V0dWJlLmNvbS9Acmxlc3BvcnRzKXwK"
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestBracketWithLeadingZero(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Winter/Middle_East_and_North_Africa/Open"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/bracket/%s/day/3", encodedUrl)

	responseData := responseForEndpoint(t, endpoint)

	expectedEncodedMarkup := "fCpFTElNSU5BVElPTip8fFsqKkxpcXVpcGVkaWEgQnJhY2tldCoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9Sb2NrZXRfTGVhZ3VlX0NoYW1waW9uc2hpcF9TZXJpZXMvMjAyMi0yMy9XaW50ZXIvTWlkZGxlX0Vhc3RfYW5kX05vcnRoX0FmcmljYS9PcGVuI1Jlc3VsdHMpfAp8Oi18Oi18Oi18CnwqKlJ1bGUgT25lKip8Kio0IC0gMCoqfFZpc2lvbiBFc3BvcnRzfAp8KipUZWFtIEZhbGNvbnMqKnwqKjQgLSAzKip8Q29sYXwKfCoqVGVhbSBGYWxjb25zKip8Kio0IC0gMCoqfFJ1bGUgT25lfA=="
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestBracket(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Fall/North_America/Cup"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/bracket/%s/day/3", encodedUrl)

	responseData := responseForEndpoint(t, endpoint)

	expectedEncodedMarkup := "fCpFTElNSU5BVElPTip8fFsqKkxpcXVpcGVkaWEgQnJhY2tldCoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9Sb2NrZXRfTGVhZ3VlX0NoYW1waW9uc2hpcF9TZXJpZXMvMjAyMi0yMy9GYWxsL05vcnRoX0FtZXJpY2EvQ3VwI1Jlc3VsdHMpfAp8Oi18Oi18Oi18CnwqKlZlcnNpb24xKip8Kio0IC0gMSoqfE5SR3wKfCoqR2VuLkcgTW9iaWwxIFJhY2luZyoqfCoqNCAtIDEqKnxTcGFjZXN0YXRpb258CnxHZW4uRyBNb2JpbDEgUmFjaW5nfCoqMyAtIDQqKnwqKlZlcnNpb24xKip8"
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestTitle(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Fall/North_America/Cup"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/title/%s", encodedUrl)

	responseData := responseForEndpoint(t, endpoint)

	expectedEncodedMarkup := "UkxDUyAyMDIyLTIzIC0gRmFsbDogTkEgUmVnaW9uYWwgMiAtIEZhbGwgQ3Vw"
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestMakeThread(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Winter/Asia-Pacific/Cup"
	encodedUrl := utils.EncodedBase64(url)
	template := utils.EncodedBase64("groups-bracketrd1-streams")
	endpoint := fmt.Sprintf("/makethread/%s/template/%s/day/1", encodedUrl, template)

	responseData := responseForEndpoint(t, endpoint)

	expectedEncodedMarkup := "IyBSTENTIDIwMjItMjMgLSBXaW50ZXI6IEFQQUMgUmVnaW9uYWwgMiAtIFdpbnRlciBDdXANCnx8KipEYXkqKnwqKlVUQyoqfHwKfDotfDotfDotfDotfAp8RGF5IDF8U2F0dXJkYXl8WyoqMTA6MDAqKl0oaHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS9zZWFyY2g/cT0xMDowMCtHTVQpfCoqVG9kYXkqKnwKfERheSAyfFN1bmRheXxbKioxMDowMCoqXShodHRwczovL3d3dy5nb29nbGUuY29tL3NlYXJjaD9xPTEwOjAwK0dNVCl8fA0KDQojIENvdmVyYWdlDQpbKipMaXF1aXBlZGlhKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1JvY2tldF9MZWFndWVfQ2hhbXBpb25zaGlwX1Nlcmllcy8yMDIyLTIzL1dpbnRlci9Bc2lhLVBhY2lmaWMvQ3VwKSAqKi8gLyoqIFsqKlN0YXJ0LmdnKipdKGh0dHBzOi8vd3d3LnN0YXJ0LmdnL3RvdXJuYW1lbnQvcmxjcy0yMDIyLTIzLXdpbnRlci1jdXAtYXBhYy9ldmVudC9tYWluLWV2ZW50KSAqKi8gLyoqIFsqKlBpY2tzdG9wLmdnKipdKGh0dHBzOi8vcGlja3N0b3AuZ2cvcmwpDQoNCiMgU3RyZWFtcw0KfCoqUGxhdGZvcm1zKip8KipMaW5rKip8Cnw6LXw6LXwKfFR3aXRjaHxbKipSb2NrZXRsZWFndWVhcGFjKipdKGh0dHBzOi8vd3d3LnR3aXRjaC50di9yb2NrZXRsZWFndWVhcGFjKXwKfFlvdXR1YmV8WyoqUmxlc3BvcnRzKipdKGh0dHBzOi8vd3d3LnlvdXR1YmUuY29tL0BybGVzcG9ydHMpfAoNClRoZSBtYWluIGNoYW5uZWwgd2lsbCBzaG93IGEgZmVhdHVyZWQgbWF0Y2ggZnJvbSBlYWNoIHJvdW5kLCB3aGlsZSB0aGUgcGFydGljaXBhdGluZyB0ZWFtcyBtYXkgYmUgcnVubmluZyBhIHN0cmVhbSBmb3IgdGhlaXIgb3duIG1hdGNoZXMuIEJlbG93IGlzIGEgTGlzdCBvZiBUZWFtIEJyb2FkY2FzdHM6DQoNCnx8fHx8Cnw6LXw6LXw6LXw6LXwKfFtEYXJrIFJpZnQgRXNwb3J0c10oaHR0cHM6Ly93d3cudHdpdGNoLnR2L0RhcmtSaWZ0RXNwb3J0cyl8W0VsZXZhdGVdKGh0dHBzOi8vd3d3LnR3aXRjaC50di9lbGV2YXRlZ2cpfFtHYWltaW4gR2xhZGlhdG9yc10oaHR0cHM6Ly93d3cudHdpdGNoLnR2L2dhaW1pbmdsYWRpYXRvcnN0dil8W0xvdHVzIEdhbWluZ10oaHR0cHM6Ly93d3cudHdpdGNoLnR2L2xvdHVzIGdhbWluZzEpCnxbTHljdXMgRW1waXJlXShodHRwczovL3d3dy50d2l0Y2gudHYvTHljdXNFbXBpcmUpfFtOaW1tdDU1XShodHRwczovL3d3dy50d2l0Y2gudHYvSGF6emFSTCl8W1VIVUhdKGh0dHBzOi8vd3d3LnR3aXRjaC50di90YW50cmlrZXNwb3J0cyl8W1ZhaW5zaG9lc10oaHR0cHM6Ly93d3cudHdpdGNoLnR2L2FtYXRlbGxsKQ0KDQojIFJlc3VsdHMNCnx8fHx8Cnw6LXw6LXw6LXw6LXwKfCoqIyoqfCoqR3JvdXAgQSoqICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7IHwqKk1hdGNoZXMqKiB8KipHYW1lIERpZmYqKiB8CnwxfFsqKkVsZXZhdGUqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvRWxldmF0ZSl8My0wfCs5fAp8MnxbKipVSFVIKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1VIVUgpfDItMXwrM3wKfDN8WyoqQVRLKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL0FUS18oSmFwYW5lc2VfVGVhbSkpfDEtMnwtM3wKfDR8WyoqVGVhbSBHYW5EZXJTKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1RlYW1fR2FuRGVyUyl8MC0zfC05fAoKJiN4MjAwQjsKCnx8fHx8Cnw6LXw6LXw6LXw6LXwKfCoqIyoqfCoqR3JvdXAgQioqICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7IHwqKk1hdGNoZXMqKiB8KipHYW1lIERpZmYqKiB8CnwxfFsqKkRlVG9OYXRvcioqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9EZVRvTmF0b3IpfDMtMHwrOXwKfDJ8WyoqQnJva29saSoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9Ccm9rb2xpKXwyLTF8KzJ8CnwzfFsqKkFtb25nWioqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9BbW9uZ1opfDEtMnwtM3wKfDR8WyoqUmV6aSdzIE1pbmlvbnMqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvUmV6aSUyN3NfTWluaW9ucyl8MC0zfC04fAoKJiN4MjAwQjsKCnx8fHx8Cnw6LXw6LXw6LXw6LXwKfCoqIyoqfCoqR3JvdXAgQyoqICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7ICYjeDIwMEI7IHwqKk1hdGNoZXMqKiB8KipHYW1lIERpZmYqKiB8CnwxfFsqKk5pbW10NTUqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvTmltbXQ1NSl8My0wfCs2fAp8MnxbKipMb3R1cyBHYW1pbmcqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvTG90dXNfR2FtaW5nKXwyLTF8KzN8CnwzfFsqKkx5Y3VzIEVtcGlyZSoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9MeWN1c19FbXBpcmUpfDEtMnwwfAp8NHxbKipHb2RhbGlvbnMqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvR29kYWxpb25zKXwwLTN8LTl8CgomI3gyMDBCOwoKfHx8fHwKfDotfDotfDotfDotfAp8KiojKip8KipHcm91cCBEKiogJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgfCoqTWF0Y2hlcyoqIHwqKkdhbWUgRGlmZioqIHwKfDF8WyoqRGVjaW1hdGUgR2FtaW5nKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL0RlY2ltYXRlX0dhbWluZyl8My0wfCs4fAp8MnxbKipHYWltaW4gR2xhZGlhdG9ycyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9HYWltaW5fR2xhZGlhdG9ycyl8Mi0xfCszfAp8M3xbKipWYWluc2hvZXMqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvVmFpbnNob2VzKXwxLTJ8LTR8Cnw0fFsqKkRhcmsgUmlmdCBFc3BvcnRzKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL0RhcmtfUmlmdF9Fc3BvcnRzKXwwLTN8LTd8CgomI3gyMDBCOwoKDQoNCiMgUGxheW9mZnMgLSBSb3VuZCAxDQp8KkVMSU1JTkFUSU9OKnx8WyoqTGlxdWlwZWRpYSBCcmFja2V0KipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1JvY2tldF9MZWFndWVfQ2hhbXBpb25zaGlwX1Nlcmllcy8yMDIyLTIzL1dpbnRlci9Bc2lhLVBhY2lmaWMvQ3VwI1Jlc3VsdHMpfAp8Oi18Oi18Oi18CnwqKkJyb2tvbGkqKnwqKjMgLSAwKip8VmFpbnNob2VzfAp8KipMb3R1cyBHYW1pbmcqKnwqKjMgLSAxKip8QVRLfAp8VUhVSHwqKjAgLSAzKip8KipMeWN1cyBFbXBpcmUqKnwKfCoqR2xhZGlhdG9ycyoqfCoqMyAtIDAqKnxBbW9uZ1p8"
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestTemplates(t *testing.T) {
	endpoint := "/templates"

	responseMarkup := string(responseForEndpoint(t, endpoint))

	expectedEncodedMarkup := "YnJhY2tldCwgYnJhY2tldC1wcml6ZXBvb2wsIGJyYWNrZXQtc3RyZWFtcywgZ3JvdXBzLCBncm91cHMtYnJhY2tldHJkMSwgZ3JvdXBzLWJyYWNrZXRyZDEtc3RyZWFtcywgZ3JvdXBzLXByaXplcG9vbCwgZ3JvdXBzLXN0cmVhbXMsIHN3aXNzLCBzd2lzcy1icmFja2V0cmQxLCBzd2lzcy1icmFja2V0cmQxLXN0cmVhbXMsIHN3aXNzLXByaXplcG9vbCwgc3dpc3Mtc3RyZWFtcw=="
	assert.Equal(t, expectedEncodedMarkup, responseMarkup)
}

func TestPrizepool(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Winter/Middle_East_and_North_Africa/Open"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/prizepool/%s", encodedUrl)

	responseData := responseForEndpoint(t, endpoint)

	expectedEncodedMarkup := "fCoqUGxhY2UqKnwqKlByaXplKip8KipUZWFtKip8KipSTENTIFBvaW50cyoqfAp8Oi18Oi18Oi18Oi18CnwqKjFzdCoqfCQ5LDAwMHxUZWFtIEZhbGNvbnN8MjB8CnwqKjJuZCoqfCQ2LDAwMHxSdWxlIE9uZXwxNnwKfCoqM3JkLTR0aCoqfCQzLDkwMHxDb2xhfDEyfAp8KiozcmQtNHRoKip8JDMsOTAwfFZpc2lvbiBFc3BvcnRzfDEyfAp8Kio1dGgtOHRoKip8JDEsODAwfEFycm93c3w4fAp8Kio1dGgtOHRoKip8JDEsODAwfEVNUFRZfDh8CnwqKjV0aC04dGgqKnwkMSw4MDB8VHJvdWJsZXN8OHwKfCoqNXRoLTh0aCoqfCQxLDgwMHxUd2lzdGVkIE1pbmRzfDh8"
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestMVP(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Fall/North_America/Cup"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/mvp_candidates/%s/teams_allowed/4", encodedUrl)

	responseData := responseForEndpoint(t, endpoint)

	expectedEncodedMarkup := "dG9ybWVudCAoVmVyc2lvbjEpCmNvbW0gKFZlcnNpb24xKQpCZWFzdE1vZGUgKFZlcnNpb24xKQpBcHBhcmVudGx5SmFjayAoR2VuLkcgTW9iaWwxIFJhY2luZykKQ2hyb25pYyAoR2VuLkcgTW9iaWwxIFJhY2luZykKbm9seSAoR2VuLkcgTW9iaWwxIFJhY2luZykKQXJzZW5hbCAoU3BhY2VzdGF0aW9uIEdhbWluZykKRGFuaWVsIChTcGFjZXN0YXRpb24gR2FtaW5nKQpMaiAoU3BhY2VzdGF0aW9uIEdhbWluZykKR2FycmV0dEcgKE5SRykKanVzdGluLiAoTlJHKQpTcXVpc2h5IChOUkcp"
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestStreams(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2024/Major_1/North_America/Open_Qualifier_1"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/streams/%s", encodedUrl)

	responseData := responseForEndpoint(t, endpoint)

	expectedEncodedMarkup := "fHx8fHwKfDotfDotfDotfDotfAp8W0RFTEVURUQgWERdKGh0dHBzOi8vd3d3LnR3aXRjaC50di9EZWxldGVkWERHYW1pbmcpfFtEaWduaXRhc10oaHR0cHM6Ly93d3cudHdpdGNoLnR2L0RpZ25pdGFzKXxbRzIgU3RyaWRlXShodHRwczovL3d3dy50d2l0Y2gudHYvZzJlc3BvcnRzKXxbR2VuLkcgTW9iaWwxIFJhY2luZ10oaHR0cHM6Ly93d3cudHdpdGNoLnR2L3dpZG93KQp8W0x1bWlub3NpdHkgR2FtaW5nXShodHRwczovL3d3dy50d2l0Y2gudHYvbHVtaW5vc2l0eWdhbWluZyl8W004MF0oaHR0cHM6Ly93d3cudHdpdGNoLnR2L0ZlZXIpfFtUaGUgTXVmZmluIE1lbl0oaHR0cHM6Ly93d3cudHdpdGNoLnR2L0NKQ0opfFtOUkddKGh0dHBzOi8vd3d3LnR3aXRjaC50di9OUkcpCnxbT0ddKGh0dHBzOi8vd3d3LnR3aXRjaC50di90aGVkYW5nZXJ0YWNvKXxbUGlyYXRlcyBvbiBhIEJvYXRdKGh0dHBzOi8vd3d3LnR3aXRjaC50di9zcG9vZGFoKXxbUGxvdCBUd2lzdF0oaHR0cHM6Ly93d3cudHdpdGNoLnR2L3B5cm9qKXxbU2hvcGlmeSBSZWJlbGxpb25dKGh0dHBzOi8vd3d3LnR3aXRjaC50di9zaG9waWZ5cmViZWxsaW9uKQp8W1RoZSBTbm93bWVuXShodHRwczovL3d3dy50d2l0Y2gudHYvRnJvc3R5KXxbU3BhY2VzdGF0aW9uIEdhbWluZ10oaHR0cHM6Ly93d3cudHdpdGNoLnR2L2NoaWVmYmVlZiBybCl8W1RTTV0oaHR0cHM6Ly93d3cudHdpdGNoLnR2L0phbWFpY2FuQ29jb251dCl8W1lvdW5nIFdoaXBwZXJzbmFwcGVyc10oaHR0cHM6Ly93d3cudHdpdGNoLnR2L1BOREgp"
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestCoverage(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Fall/North_America/Cup"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/coverage/%s", encodedUrl)

	responseData := responseForEndpoint(t, endpoint)

	expectedEncodedMarkup := "WyoqTGlxdWlwZWRpYSoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9Sb2NrZXRfTGVhZ3VlX0NoYW1waW9uc2hpcF9TZXJpZXMvMjAyMi0yMy9GYWxsL05vcnRoX0FtZXJpY2EvQ3VwKSAqKi8gLyoqIFsqKlN0YXJ0LmdnKipdKGh0dHBzOi8vd3d3LnN0YXJ0LmdnL3RvdXJuYW1lbnQvcmxjcy0yMDIyLTIzLWZhbGwtY3VwLW5vcnRoLWFtZXJpY2EvZXZlbnQvbWFpbi1ldmVudCkgKiovIC8qKiBbKipQaWNrc3RvcC5nZyoqXShodHRwczovL3BpY2tzdG9wLmdnL3JsKQ=="
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestGroups(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2021-22/Winter/Middle_East_and_North_Africa/2"
	encodedUrl := utils.EncodedBase64(url)
	endpoint := fmt.Sprintf("/groups/%s", encodedUrl)

	responseData := responseForEndpoint(t, endpoint)

	expectedEncodedMarkup := "fHx8fHwKfDotfDotfDotfDotfAp8KiojKip8KipHcm91cCBBKiogJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgJiN4MjAwQjsgfCoqTWF0Y2hlcyoqIHwqKkdhbWUgRGlmZioqIHwKfDF8WyoqU2FuZHJvY2sgR2FtaW5nKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1NhbmRyb2NrX0dhbWluZyl8My0wfCs5fAp8MnxbKipUdWNoZWwqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvVHVjaGVsKXwyLTF8MHwKfDN8WyoqRm94IEdhbWluZyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9Gb3hfR2FtaW5nKXwxLTJ8LTF8Cnw0fFsqKjI1ZVNwb3J0cyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS8yNWVTcG9ydHMpfDAtM3wtOHwKCiYjeDIwMEI7Cgp8fHx8fAp8Oi18Oi18Oi18Oi18CnwqKiMqKnwqKkdyb3VwIEIqKiAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyB8KipNYXRjaGVzKiogfCoqR2FtZSBEaWZmKiogfAp8MXxbKipUZWFtIEZhbGNvbnMqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvVGVhbV9GYWxjb25zKXwzLTB8Kzh8CnwyfFsqKkZvcmVzdCBIdW50ZXJzKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL0ZvcmVzdF9IdW50ZXJzKXwyLTF8KzJ8CnwzfFsqKlRoZSBFdmlsIGVTcG9ydHMqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvVGhlX0V2aWxfZVNwb3J0cyl8MS0yfC0yfAp8NHxbKipZVCBJbW1vcnRhbHMqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvaW5kZXgucGhwP3RpdGxlPVlvdSUyN3JlX1Rocm93aW5nX0ltbW9ydGFscyZhY3Rpb249ZWRpdCZyZWRsaW5rPTEpfDAtM3wtOHwKCiYjeDIwMEI7Cgp8fHx8fAp8Oi18Oi18Oi18Oi18CnwqKiMqKnwqKkdyb3VwIEMqKiAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyB8KipNYXRjaGVzKiogfCoqR2FtZSBEaWZmKiogfAp8MXxbKipUaGUgVWx0aW1hdGVzKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1RoZV9VbHRpbWF0ZXMpfDItMXwrM3wKfDJ8WyoqUmV2b2x1dGlvbioqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9SZXZvbHV0aW9uKXwyLTF8KzR8CnwzfFsqKlQySyoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9UMkspfDItMXwtMXwKfDR8WyoqV29sdmVzKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL1dvbHZlc18oTWlkZGxlX0Vhc3Rlcm5fVGVhbSkpfDAtM3wtNnwKCiYjeDIwMEI7Cgp8fHx8fAp8Oi18Oi18Oi18Oi18CnwqKiMqKnwqKkdyb3VwIEQqKiAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyAmI3gyMDBCOyB8KipNYXRjaGVzKiogfCoqR2FtZSBEaWZmKiogfAp8MXxbKipTQ1lURVMqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvU0NZVEVTKXwzLTB8KzZ8CnwyfFsqKkFOS0FBIEVzcG9ydHMqKl0oaHR0cHM6Ly9saXF1aXBlZGlhLm5ldC9yb2NrZXRsZWFndWUvQU5LQUFfRXNwb3J0cyl8Mi0xfCs1fAp8M3xbKipLSU5HUyBFc3BvcnRzKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL0tJTkdTX0VzcG9ydHMpfDEtMnwtM3wKfDR8WyoqTmlnaHRtYXJlKipdKGh0dHBzOi8vbGlxdWlwZWRpYS5uZXQvcm9ja2V0bGVhZ3VlL05pZ2h0bWFyZSl8MC0zfC04fAoKJiN4MjAwQjsKCg=="
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

func TestSchedule(t *testing.T) {
	url := "https://liquipedia.net/rocketleague/Rocket_League_Championship_Series/2022-23/Fall/North_America/Cup"
	encodedUrl := utils.EncodedBase64(url)

	// Day 1
	endpoint := fmt.Sprintf("/schedule/%s/day/1", encodedUrl)
	responseData := responseForEndpoint(t, endpoint)

	expectedEncodedMarkup := "fHwqKkRheSoqfCoqVVRDKip8fAp8Oi18Oi18Oi18Oi18CnxEYXkgMXxGcmlkYXl8WyoqMTc6MDAqKl0oaHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS9zZWFyY2g/cT0xNzowMCtHTVQpfCoqVG9kYXkqKnwKfERheSAyfFNhdHVyZGF5fFsqKjE3OjAwKipdKGh0dHBzOi8vd3d3Lmdvb2dsZS5jb20vc2VhcmNoP3E9MTc6MDArR01UKXx8CnxEYXkgM3xTdW5kYXl8WyoqMTc6MDAqKl0oaHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS9zZWFyY2g/cT0xNzowMCtHTVQpfHw="
	assert.Equal(t, expectedEncodedMarkup, string(responseData))

	// Day 3
	endpoint = fmt.Sprintf("/schedule/%s/day/3", encodedUrl)
	responseData = responseForEndpoint(t, endpoint)

	expectedEncodedMarkup = "fHwqKkRheSoqfCoqVVRDKip8fAp8Oi18Oi18Oi18Oi18Cnx+fkRheSAxfn58fn5GcmlkYXl+fnx+flsqKjE3OjAwKipdKGh0dHBzOi8vd3d3Lmdvb2dsZS5jb20vc2VhcmNoP3E9MTc6MDArR01UKX5+fHwKfH5+RGF5IDJ+fnx+flNhdHVyZGF5fn58fn5bKioxNzowMCoqXShodHRwczovL3d3dy5nb29nbGUuY29tL3NlYXJjaD9xPTE3OjAwK0dNVCl+fnx8CnxEYXkgM3xTdW5kYXl8WyoqMTc6MDAqKl0oaHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS9zZWFyY2g/cT0xNzowMCtHTVQpfCoqVG9kYXkqKnw="
	assert.Equal(t, expectedEncodedMarkup, string(responseData))
}
