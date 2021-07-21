# Spotify Data Visualizer #
* Full stack website for Spotify data: https://go-spotify-data.herokuapp.com/
## Specifications: ##
* Go to https://www.spotify.com/us/account/privacy/ and download your data.
* Duration, hourly distribution, monthly distribution and monthly top 3 artists charts with Chart.js.
* History, play count and play time for tracks and artists tables with DataTables.
## Dependencies: ##
* Golang
* Python 3, NumPy and Pandas
* Docker (optional)
## Run: ##
* `go mod init spotify`, `go mod tidy` then `go run spotify.go` and go to http://localhost:8080/. 
* `docker build -t spotifyalpine .` then `docker run -p 8080:8080 -tid spotifyalpine` and go to http://localhost:8080/. 
## Images: ##
<table>
    <tr>
        <td align="center">
            <img src="https://github.com/ssduman/go-spotify-data/blob/master/img/charts1.png" alt="home-page" width="384" height="216">
            <br />
            <i> 1 </i>
        </td>
        <td align="center">
            <img src="https://github.com/ssduman/go-spotify-data/blob/master/img/charts2.png" alt="play-okey" width="384" height="216">
            <br />
            <i> 2 </i>
        </td>
    </tr>
</table>

### Bugs and Limitations: ###
* Due to free version of Heroku, the site waking up like in 30 seconds if sleeping for the first time.
