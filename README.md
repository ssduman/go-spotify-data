# Spotify Data Visualizer #
* Full stack website for Spotify data.
## Specifications: ##
* Go to https://www.spotify.com/us/account/privacy/ and download your data.
* No data is kept, uploads and auxilary files are deleted instantly.
* For testing, try `sample/StreamingHistorySample.json`.
* Duration, hourly distribution, monthly distribution and monthly top 3 artists charts with Chart.js.
* History, non-stop play time and play count and play time for tracks and artists tables with DataTables.
## Dependencies: ##
* Golang
* Python 3, NumPy and Pandas
* Docker (optional)
## Run: ##
* `go run spotify.go` and go to http://localhost:8080/. 
* `docker build -t spotifyalpine .` then `docker run -p 8080:8080 -tid spotifyalpine` and go to http://localhost:8080/. 
## Images: ##
<table>
    <tr>
        <td align="center">
            <img src="https://github.com/ssduman/go-spotify-data/blob/master/img/upload.jpg" alt="home-page" width="384" height="216">
            <br />
            <i> upload page </i>
        </td>
        <td align="center">
            <img src="https://github.com/ssduman/go-spotify-data/blob/master/img/charts1.jpg" alt="play-okey" width="384" height="216">
            <br />
            <i> distributions </i>
        </td>
    </tr>
    <tr>
        <td align="center">
            <img src="https://github.com/ssduman/go-spotify-data/blob/master/img/charts2.jpg" alt="home-page" width="384" height="216">
            <br />
            <i> charts </i>
        </td>
        <td align="center">
            <img src="https://github.com/ssduman/go-spotify-data/blob/master/img/charts3.jpg" alt="play-okey" width="384" height="216">
            <br />
            <i> top tracks tables </i>
        </td>
    </tr>
</table>
