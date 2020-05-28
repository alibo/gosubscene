# gosubscene (v0.0.1)
A simple API wrapper for [subscene.com](subscene.com) witten in golang.

## Usage

Run the api server with docker:

```bash
docker run -p 3000:3000 alibo/gosubscene:0.0.1
```

### Search

To search a movie/tv-show: 

```json
$ curl -s 'localhost:3000/search?q=rick and morty' | jq
{
  "status": 200,
  "data": {
    "found": true,
    "categories": {
      "Close": [
        {
          "title": "Rick and Morty - Fourth season",
          "url": "/subtitles/rick-and-morty-fourth-season"
        },
        {
          "title": "Rick and Morty - Third Season",
          "url": "/subtitles/rick-and-morty-third-season"
        },
        // ...
      ]
    }
}
```

### List

To list of all subtitles for an item: (Currently, only Farsi/Persian subtitles are supported!)

```json
$ curl -s localhost:3000/subtitles/rick-and-morty-fourth-season | jq 
{
  "status": 200,
  "data": [
    {
      "lang": "Farsi/Persian",
      "release": "Rick.and.Morty.S04E08.The.Vat.of.Acid.Episode.WEB-DL.2CH.x265.HEVC-PSA",
      "user": "submen",
      "comment": "فقط تنظیم کردم. مترجم عرفان صدیقی",
      "url": "/subtitles/rick-and-morty-fourth-season/farsi_persian/2222495"
    },
    {
      "lang": "Farsi/Persian",
      "release": "Rick.and.Morty.S04E09.Childrick.of.Mort.WEB-DL.2CH.x265.HEVC-PSA",
      "user": "submen",
      "comment": "فقط با این نسخه تنظیم کردم. مترجم عرفان صدیقی",
      "url": "/subtitles/rick-and-morty-fourth-season/farsi_persian/2222492"
    },
    // ...
  ]
}
```

### Details

To get details of a subtitle and the download link:

```json
$ curl -s localhost:3000/subtitles/rick-and-morty-fourth-season/farsi_persian/2222145 | jq
{
  "status": 200,
  "data": {
    "releases": [
      "Rick.and.Morty.S04E09.WEB-DL",
      "Rick.and.Morty.S04E09.AMZN.WEB-DL",
      "Rick.and.Morty.S04E09.WEB-DL.HEVC.x265-RMTeam",
      "Rick.and.Morty.S04E09.AMZN.WEB-DL.6CH.x264-DLHA",
      "Rick.and.Morty.S04E09.WEB-DL.HEVC.x264-RMTeam",
      "Rick.and.Morty.S04E09.AMZN.WEB-DL.6CH.x265-DLHA",
      "Rick.and.Morty.S04E09.1080p.HEVC.x265-MeGusta",
      "Rick.and.Morty.S04E09.1080p.WEB-DL.6CH.x265.HEVC-PSA",
      "Rick.and.Morty.S04E09.WEB-DL.480p.x264-RMTeam"
    ],
    "description": "Horizon Group برگردان و همگام سازی توسط www.horizongroup.ir\nهماهنگ با نسخه های نسخه های بدون سانسور:WEB-DL - RMTeam - DLHA\n(متن ترجمه بدون سانسور می باشد)\n(استفاده کنید subscene در صورت ناهماهنگی از نسخه دیگر موجود در)",
    "user": "HorizonGroup",
    "userRating": 10,
    "datetime": "5/25/2020 10:32 PM",
    "files": 1,
    "rate": 10,
    "rateCount": 6,
    "downloads": 133,
    "downloadLink": "/download/?token=lYXObJiYQNjJJixPMjVm9lbsYqAlfttxVupArYAYsDpiQXwXp-0R0E6DK1JPmi8k2rFkfX6ZhEx59BoWjozSG4tuf9-lYUAhb71CElAm8Kz_3K4fmrR0HGfLelZY6HDi0",
    "comments": 0,
    "commentLink": "https://comments.jeded.com/comments/2222145"
  }
}
```

### Download

To download a subtitle:

```bash
$ curl -OJl 'http://localhost:3000/download/?token=lYXObJiYQNjJJixPMjVm9lbsYqAlfttxVupArYAYsDpiQXwXp-0R0E6DK1JPmi8k2rFkfX6ZhEx59BoWjozSG4tuf9-lYUAhb71CElAm8Kz_3K4fmrR0HGfLelZY6HDi0'
  % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
                                 Dload  Upload   Total   Spent    Left  Speed
100 17399  100 17399    0     0  64680      0 --:--:-- --:--:-- --:--:-- 64680
curl: Saved to filename 'rick-and-morty-fourth-season_farsi_persian-2222145.zip'
```
