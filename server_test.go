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

func responseForEndpoint(endpoint string) []byte{
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

	expectedEncodedMarkup := "IyBDb3ZlcmFnZQoKWyoqTGlxdWlwZWRpYSoqXShodHRwczovL2xpcXVpcGVkaWEubmV0L3JvY2tldGxlYWd1ZS9Sb2NrZXRfTGVhZ3VlX0NoYW1waW9uc2hpcF9TZXJpZXMvMjAyMi0yMy9GYWxsL05vcnRoX0FtZXJpY2EvQ3VwKSAqKi8gLyoqIFsqKk9jdGFuZS5nZyoqXShodHRwczovL29jdGFuZS5nZy9ldmVudHMvYzAzNS1ybGNzLTIwMjItMjMtZmFsbC1ub3J0aC1hbWVyaWNhLXJlZ2lvbmFsLTIpICoqLyAvKiogWyoqU3RhcnQuZ2cqKl0oaHR0cHM6Ly93d3cuc3RhcnQuZ2cvdG91cm5hbWVudC9ybGNzLTIwMjItMjMtZmFsbC1jdXAtbm9ydGgtYW1lcmljYS9ldmVudC9tYWluLWV2ZW50KQ=="
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

	expectedEncodedMarkup = "fHwqKkRheSoqfCoqVVRDKip8fAp8Oi18Oi18Oi18Oi18CnxEYXkgMXxGcmlkYXl8WyoqMTc6MDAqKl0oaHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS9zZWFyY2g/cT0xNzowMCtHTVQpfHwKfERheSAyfFNhdHVyZGF5fFsqKjE3OjAwKipdKGh0dHBzOi8vd3d3Lmdvb2dsZS5jb20vc2VhcmNoP3E9MTc6MDArR01UKXx8CnxEYXkgM3xTdW5kYXl8WyoqMTc6MDAqKl0oaHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS9zZWFyY2g/cT0xNzowMCtHTVQpfCoqVG9kYXkqKnw="
	assert.Equal(t, http.StatusOK, 200)
    assert.Equal(t, expectedEncodedMarkup, string(responseData))
}

