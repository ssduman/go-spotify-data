// https://gist.github.com/olegpolukhin/3a4379a400c2c928f2d23059a78f1b82
// https://stackoverflow.com/questions/25087960/json-unmarshal-time-that-isnt-in-rfc-3339-format

package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

type CustomTime struct {
	time.Time
}

type StreamingHistory struct {
	EndTime    CustomTime `json:"endTime"`
	ArtistName string     `json:"artistName"`
	TrackName  string     `json:"trackName"`
	MsPlayed   int        `json:"msPlayed"`
	hhmmss     string
}

const ctLayout = "2006-01-02 15:04"

var streamingHistory []StreamingHistory

var templateDF = make(map[string][][]string)
var templateData = make(map[string][]string)

func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), "\"")
	ct.Time, err = time.Parse(ctLayout, s)
	return
}

func (ct *CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", ct.Time.Format(ctLayout))), nil
}

func (ct *CustomTime) String() string {
	return string(ct.Time.Format(ctLayout))
}

func readFile(filename string) []byte {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Printf("failed to open json file: %s, error: %v\n", filename, err)
		return nil
	}

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v\n", err)
		return nil
	}

	return data
}

func convertUTC(ct CustomTime, timeZone string) time.Time {
	location, err := time.LoadLocation(timeZone)
	if err != nil {
		panic(err)
	}

	converted := ct.In(location).Format(ctLayout)

	convertedTime, err := time.Parse(ctLayout, converted)
	if err != nil {
		panic(err)
	}
	return convertedTime
}

func readCSV(path string) ([][]string, error) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("Unable to read input file "+path, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		fmt.Printf("Unable to parse file as CSV for "+path, err)
		return nil, err
	}

	return records, nil
}

func sliceColumn(matrix [][]string, index int) ([]string, error) {
	colSize := len(matrix[0])
	if index >= colSize {
		return nil, errors.New("out of index")
	}

	column := make([]string, 0)
	for _, row := range matrix {
		column = append(column, row[index])
	}
	return column, nil
}

func sliceStep(arr []string, first int, step int) ([]string, error) {
	size := len(arr)
	if first >= size {
		return nil, errors.New("out of index")
	}

	sliced := make([]string, 0)
	for index := first; index < size; index += step {
		sliced = append(sliced, arr[index])
	}
	return sliced, nil
}

func ms_to_hour(ms int, hour bool) string {
	seconds := (ms / 1000) % 60
	minutes := (ms / (1000 * 60)) % 60
	hours := (ms / (1000 * 60 * 60)) % 24

	if hour {
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
	}
	return fmt.Sprintf("%02d:%02d", minutes, seconds)
}

func ms_to_hour_map(arr []string, f func(int, bool) string) []string {
	newArr := make([]string, len(arr))
	for i := 0; i < len(arr); i++ {
		result := strings.Split(arr[i], " ")
		v1, _ := strconv.Atoi(result[0][1 : len(result[0])-1])
		v2, _ := strconv.Atoi(result[1][:len(result[1])-1])
		newArr[i] = f(v1, false) + " - " + f(v2, false)
	}
	return newArr
}

func ms_to_hour_map_concurent(arr []string, f func(int, bool) string, c_ms chan []string) {
	newArr := make([]string, len(arr))
	for i := 0; i < len(arr); i++ {
		result := strings.Split(arr[i], " ")
		v1, _ := strconv.Atoi(result[0][1 : len(result[0])-1])
		v2, _ := strconv.Atoi(result[1][:len(result[1])-1])
		newArr[i] = f(v1, false) + " - " + f(v2, false)
	}
	c_ms <- newArr
}

func readHistory(filesPath string, filesNameArray []string) []StreamingHistory {
	var streamingHistory []StreamingHistory
	var tempHistory []StreamingHistory

	for _, v := range filesNameArray {
		streamingHistoryData := readFile(filesPath + v)
		err := json.Unmarshal(streamingHistoryData, &tempHistory)
		if err != nil {
			fmt.Printf("failed to Unmarshal json file, error: %v\n", err)
		}
		streamingHistory = append(streamingHistory, tempHistory...)
	}

	return streamingHistory
}

func createData(filesPath string, filesNameArray []string, timeZone string) error {
	streamingHistory = readHistory(filesPath, filesNameArray)

	for i := 0; i < len(streamingHistory); i++ {
		streamingHistory[i].hhmmss = ms_to_hour(streamingHistory[0].MsPlayed, true)
		streamingHistory[i].EndTime = CustomTime{convertUTC(streamingHistory[i].EndTime, timeZone)}
	}

	hour_df, err := readCSV("df/spotify_hour_df.csv")
	monthly_df, err := readCSV("df/monthly_df.csv")
	interval_df, err := readCSV("df/interval_df.csv")
	top_monthly_df, err := readCSV("df/top_monthly_df.csv")
	templateDF["top_artist_df"], err = readCSV("df/top_artist_df.csv")
	templateDF["nonstop_play_df"], err = readCSV("df/nonstop_play_df.csv")
	templateDF["top_artist_time_df"], err = readCSV("df/top_artist_time_df.csv")
	templateDF["top_artists_tracks_count_df"], err = readCSV("df/top_artists_tracks_count_df.csv")
	templateDF["top_artists_tracks_playtime_df"], err = readCSV("df/top_artists_tracks_playtime_df.csv")

	if err != nil {
		return err
	}

	templateData["hourData"], _ = sliceColumn(hour_df, 1)
	templateData["hourDataLabel"], _ = sliceColumn(hour_df, 0)

	templateData["monthlyData"], _ = sliceColumn(monthly_df, 1)
	templateData["monthlyDataLabel"], _ = sliceColumn(monthly_df, 0)

	templateData["intervalData"], _ = sliceColumn(interval_df, 1)
	intervalDataLabel, _ := sliceColumn(interval_df, 0)
	// templateData["intervalDataLabel"] = ms_to_hour_map(intervalDataLabel[1:], ms_to_hour)
	c_ms := make(chan []string)
	go ms_to_hour_map_concurent(intervalDataLabel[1:], ms_to_hour, c_ms)
	templateData["intervalDataLabel"] = <-c_ms

	monthArtist, _ := sliceColumn(top_monthly_df, 0)
	monthCount, _ := sliceColumn(top_monthly_df, 1)
	monthMonth, _ := sliceColumn(top_monthly_df, 2)

	templateData["monthTop1Artis"], _ = sliceStep(monthArtist[1:], 0, 3)
	templateData["monthTop1Count"], _ = sliceStep(monthCount[1:], 0, 3)

	templateData["monthTop2Artis"], _ = sliceStep(monthArtist[1:], 1, 3)
	templateData["monthTop2Count"], _ = sliceStep(monthCount[1:], 1, 3)

	templateData["monthTop3Artis"], _ = sliceStep(monthArtist[1:], 2, 3)
	templateData["monthTop3Count"], _ = sliceStep(monthCount[1:], 2, 3)

	templateData["monthTopMonths"], _ = sliceStep(monthMonth[1:], 0, 3)

	return nil
}

func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func showCharts(c *gin.Context) {
	if _, err := os.Stat("df"); os.IsNotExist(err) {
		c.Redirect(http.StatusFound, "/upload")
	} else {
		os.RemoveAll("df")
		c.HTML(http.StatusOK, "charts.html", gin.H{
			"streamingHistory":  streamingHistory,
			"hourData":          templateData["hourData"][1:],
			"hourDataLabel":     templateData["hourDataLabel"][1:],
			"monthlyData":       templateData["monthlyData"][1:],
			"monthlyDataLabel":  templateData["monthlyDataLabel"][1:],
			"intervalData":      templateData["intervalData"][1:],
			"intervalDataLabel": templateData["intervalDataLabel"],

			"monthTopMonths": templateData["monthTopMonths"],
			"monthTop1Artis": templateData["monthTop1Artis"],
			"monthTop1Count": templateData["monthTop1Count"],
			"monthTop2Artis": templateData["monthTop2Artis"],
			"monthTop2Count": templateData["monthTop2Count"],
			"monthTop3Artis": templateData["monthTop3Artis"],
			"monthTop3Count": templateData["monthTop3Count"],

			"top_artist":                  templateDF["top_artist_df"][1:],
			"nonstop_play":                templateDF["nonstop_play_df"][1:],
			"top_artist_time":             templateDF["top_artist_time_df"][1:],
			"top_artists_tracks_count":    templateDF["top_artists_tracks_count_df"][1:],
			"top_artists_tracks_playtime": templateDF["top_artists_tracks_playtime_df"][1:],
		})
	}
}

func uploadFileGet(c *gin.Context) {
	if _, err := os.Stat("upload"); os.IsNotExist(err) {
		os.Mkdir("upload", 0777)
	}

	errorMessage, err := c.Cookie("errorMessage")
	if err != nil {
		errorMessage = ""
	}

	c.HTML(http.StatusOK, "upload.html", gin.H{"errorMessage": errorMessage})
}

func uploadFilePost(c *gin.Context) {
	if _, err := os.Stat("upload"); os.IsNotExist(err) {
		os.Mkdir("upload", 0777)
	}

	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	isSample := c.PostForm("isSample")

	files := form.File["files"]
	filesPath := "upload/"
	filesNameArray := make([]string, 0)

	if len(files) == 0 && isSample == "" {
		c.SetCookie("errorMessage", "No file uploaded", 10, "/", c.Request.URL.Hostname(), false, true)
		c.Redirect(http.StatusFound, "/upload")
		return
	} else if len(files) != 0 && isSample == "" {
		for _, file := range files {
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, filesPath+filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
			filesNameArray = append(filesNameArray, filename)
		}
		sort.Strings(filesNameArray)
	} else if isSample != "" {
		filesNameArray = append(filesNameArray, "StreamingHistorySample.json")
		filesPath = "sample/"
	}

	pythonVer := "python"
	if os.Getenv("IS_DOCKER") != "" {
		pythonVer = "python3"
	}

	timeZone := c.PostForm("timeZone")
	_, err = time.LoadLocation(timeZone)
	if err != nil {
		c.SetCookie("errorMessage", "Time zone not found", 10, "/", c.Request.URL.Hostname(), false, true)
		c.Redirect(http.StatusFound, "/upload")
		return
	}

	command := []string{pythonVer, "script/spotify.py", "--path", filesPath, "--timeZone", timeZone, "--files"}
	command = append(command, filesNameArray...)

	cmd := exec.Command(command[0], command[1:]...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		panic(err)
	}
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	go copyOutput(stdout)
	go copyOutput(stderr)

	cmd.Wait()

	err = createData(filesPath, filesNameArray, timeZone)
	if err != nil {
		if _, err := os.Stat("upload"); !os.IsNotExist(err) {
			os.RemoveAll("upload")
		}
		if _, err := os.Stat("df"); !os.IsNotExist(err) {
			os.RemoveAll("df")
		}
		c.SetCookie("errorMessage", "Wrong data format, https://github.com/ssduman/go-spotify-data/blob/master/sample/StreamingHistorySample.json", 10, "/", c.Request.URL.Hostname(), false, true)
		c.Redirect(http.StatusFound, "/upload")
		return
	}

	c.Redirect(http.StatusFound, "/charts")

	os.RemoveAll("upload")
}

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFile("/favicon.png", "./assets/img/favicon.png")
	router.LoadHTMLGlob("templates/*")
	router.NoRoute(func(c *gin.Context) { c.JSON(http.StatusNotFound, gin.H{"m": "not found"}) })

	router.GET("/", showCharts)
	router.GET("/charts", showCharts)
	router.GET("/upload", uploadFileGet)
	router.POST("/upload", uploadFilePost)

	port := os.Getenv("PORT")
	if port != "" {
		router.Run(":" + port)
	} else {
		router.Run(":8080")
	}
}
