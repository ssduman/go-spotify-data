import os
import glob
import json
import argparse
import datetime
import numpy as np
import pandas as pd

def get_args():
    parser = argparse.ArgumentParser(description="spotify")
    parser.add_argument("--path", type=str, default="upload", help="input files")
    parser.add_argument("--timeZone", type=str, default="UTC", help="time zone")
    parser.add_argument("--files", nargs="+")
    return parser.parse_args()

def ms_to_hour(ms, hour=False):
    seconds = (ms / 1000) % 60
    minutes = (ms / (1000 * 60)) % 60
    hours = (ms / (1000 * 60 * 60)) % 24

    if hour:
        return "%02d:%02d:%02d" % (hours, minutes, seconds)
    return "%02d:%02d" % (minutes, seconds)

def top3(spotify_df):
    top_monthly_df = pd.DataFrame(spotify_df)
    top_monthly_df.set_index("endTime", drop=True, inplace=True)
    top_monthly_df = top_monthly_df.resample("ME").artistName.apply(lambda x: x.value_counts().head(3)).reset_index()
    top_monthly_df = top_monthly_df.rename(columns={"artistName": "count", "level_1": "artistName"})
    top_monthly_df["endTime"] = top_monthly_df.index.tolist()
    top_monthly_df = top_monthly_df.reset_index(drop=True)
    top_monthly_df.to_csv("df/top_monthly_df.csv", index=False)
    
def intervals(spotify_df):
    interval = spotify_df["msPlayed"].tolist()
    interval_df = pd.DataFrame(interval, columns=["Duration"])
    interval_range = np.arange(0, max(interval) + 120000, 120000)
    interval_df = interval_df.groupby(pd.cut(interval_df["Duration"], interval_range), observed=True).count()
    interval_new_df = pd.DataFrame({
        "intervat": [str(i) for i in interval_df.index.tolist()],
        "count": interval_df["Duration"].tolist()
    })
    interval_new_df.to_csv("df/interval_df.csv", index=False)

def make_tracks(spotify_df):
    artists = spotify_df["artistName"].tolist()
    tracks = spotify_df["trackName"].tolist()
    artists_tracks = [artists, tracks]
    artists_tracks_list = list(map(" - ".join, zip(*artists_tracks)))
    artists_tracks_counts = dict()
    for artists_track in artists_tracks_list:
        artists_tracks_counts[artists_track] = artists_tracks_counts.get(artists_track, 0) + 1
    top_artists_tracks = np.array(sorted(artists_tracks_counts.items(), key=lambda item: item[1], reverse=True))
    top_artists_tracks_df = spotify_df.groupby(["artistName", "trackName"]).size().reset_index(name="playCount")
    top_artists_tracks_df = top_artists_tracks_df.sort_values(by=["playCount"], ascending=False)
    top_artists_tracks_df = top_artists_tracks_df.reset_index(drop=True)
    top_artists_tracks_df.to_csv("df/top_artists_tracks_count_df.csv", index=False)

    top_artists_tracks_df = spotify_df.groupby(["artistName", "trackName"]).sum(numeric_only=True).reset_index()
    top_artists_tracks_df["msPlayed"] = top_artists_tracks_df["msPlayed"].apply(lambda x: ms_to_hour(x, hour=True))
    top_artists_tracks_df = top_artists_tracks_df.sort_values(by=["msPlayed"], ascending=False)
    top_artists_tracks_df = top_artists_tracks_df.rename(columns={"msPlayed": "playTime"})
    top_artists_tracks_df = top_artists_tracks_df.reset_index(drop=True)
    top_artists_tracks_df.to_csv("df/top_artists_tracks_playtime_df.csv", index=False)

def make_artists(spotify_df):
    artists = spotify_df["artistName"].tolist()
    artists_counts = dict()
    for artist in artists:
        artists_counts[artist] = artists_counts.get(artist, 0) + 1
    top_artist = np.array(sorted(artists_counts.items(), key=lambda item: item[1], reverse=True))
    top_artist_df = pd.DataFrame(top_artist, columns=["Artist", "Count"])
    top_artist_df.to_csv("df/top_artist_df.csv", index=False)

    top_artist_time_df = spotify_df[["artistName", "msPlayed"]].copy().groupby(["artistName"]).sum(numeric_only=True).reset_index()
    top_artist_time_df["msPlayed"] = top_artist_time_df["msPlayed"].apply(lambda x: ms_to_hour(x, hour=True))
    top_artist_time_df = top_artist_time_df.sort_values(by=["msPlayed"], ascending=False)
    top_artist_time_df = top_artist_time_df.rename(columns={"msPlayed": "playTime"})
    top_artist_time_df = top_artist_time_df.reset_index(drop=True)
    top_artist_time_df.to_csv("df/top_artist_time_df.csv", index=False)

def distributions(spotify_df):
    time_lists = [str(m).split(" ")[1][:-9] for m in spotify_df["endTime"].tolist()]
    spotify_hour_df = pd.DataFrame(data={"hour": time_lists})
    spotify_hour_df['hour'] = pd.to_datetime(spotify_hour_df['hour'], format='%H:%M')
    spotify_hour_df.set_index('hour', drop=False, inplace=True)
    spotify_hour_df = spotify_hour_df['hour'].groupby(pd.Grouper(freq='60Min')).count()
    spotify_hour_new_df = pd.DataFrame({
        "hours": [str(h).split(" ")[1][:-3] for h in spotify_hour_df.index],
        "count": spotify_hour_df.tolist()
    })
    spotify_hour_new_df.to_csv("df/spotify_hour_df.csv", index=False)

    monthly_df = pd.DataFrame(spotify_df["endTime"], columns=["endTime"])
    monthly_df.set_index("endTime", drop=False, inplace=True)
    monthly_df = monthly_df.resample('ME').count().dropna()
    xs = [str(x).split("T")[0] for x in monthly_df.index.values]
    monthly_df.index = xs
    monthly_new_df = pd.DataFrame({
        "month": monthly_df.index.tolist(),
        "count": monthly_df["endTime"].tolist()
    })
    monthly_new_df.to_csv("df/monthly_df.csv", index=False)

def nonstop_playing(spotify_df):
    end_times = [str(t)[:-9] for t in spotify_df["endTime"].tolist()]
    play_times = [ms_to_hour(t, True) for t in spotify_df["msPlayed"].tolist()]
    nonstop_records = []
    nonstop = False

    for i in range(len(end_times)):
        end_datetime = datetime.datetime.strptime(end_times[i], "%Y-%m-%d %H:%M")
        play_datetime = datetime.datetime.strptime(play_times[i], "%H:%M:%S")
        
        play_delta = datetime.timedelta(hours=play_datetime.hour, minutes=play_datetime.minute, seconds=play_datetime.second)
        
        start_datetime = end_datetime - play_delta
        
        start_time1 = "{:%Y-%m-%d %H:%M}".format(start_datetime - datetime.timedelta(minutes=1))
        start_time2 = "{:%Y-%m-%d %H:%M}".format(start_datetime)
        start_time3 = "{:%Y-%m-%d %H:%M}".format(start_datetime + datetime.timedelta(minutes=1))
        start_times = [start_time1, start_time2, start_time3]
        
        if i == 0:
            nonstop_records.append([start_time2, end_times[i], play_datetime, 1])
        elif end_times[i - 1] in start_times:
            if nonstop:
                nonstop_records[-1][1] = end_times[i]
                nonstop_records[-1][2] += + play_delta
                nonstop_records[-1][3] += 1
            nonstop = True
        else:
            nonstop_records.append([start_time2, end_times[i], play_datetime, 1])
            nonstop = False

    for i in range(len(nonstop_records)):
        nonstop_records[i][2] = "{:%H:%M:%S}".format(nonstop_records[i][2])
    nonstop_records = np.array(nonstop_records)
    nonstop_play_df = pd.DataFrame({
        "startTime": nonstop_records[:, 0],
        "endTime": nonstop_records[:, 1],
        "playTime": nonstop_records[:, 2],
        "trackCount": nonstop_records[:, 3]
    })
    nonstop_play_df.to_csv("df/nonstop_play_df.csv", index=False)

def main():
    args = get_args()
    os.makedirs("df", exist_ok=True)

    spotify = []
    try:
        for history_file in sorted(args.files):
            with open(args.path + history_file, encoding="utf8") as f:
                spotify.extend(json.load(f))
        # print(f"[python] history data count: {len(spotify)}, sample: {spotify[0]}")
    except Exception as e:
        print(f"[python] exception occurred while reading history: {e}")

    try:
        spotify_df = pd.DataFrame.from_records(spotify)
        spotify_df = spotify_df[["artistName", "trackName", "endTime", "msPlayed"]]
        spotify_df["endTime"] = pd.to_datetime(spotify_df["endTime"])
        spotify_df["endTime"] = spotify_df["endTime"].dt.tz_localize(tz="UTC")
        spotify_df["endTime"] = spotify_df["endTime"].dt.tz_convert(args.timeZone)
        spotify_df.to_csv("df/spotify_df.csv", index=False)
    except Exception as e:
        print(f"[python] exception occurred while creating df: {e}")

    try:
        top3(spotify_df)
        intervals(spotify_df)
        make_tracks(spotify_df)
        make_artists(spotify_df)
        distributions(spotify_df)
        nonstop_playing(spotify_df)
    except Exception as e:
        print(f"[python] exception occurred while parsing df: {e}")

if __name__ == "__main__":
    main()
