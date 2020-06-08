[![Build Status](https://dev.azure.com/tuommii/tuommii/_apis/build/status/tuommii.scoreboard?branchName=master)](https://dev.azure.com/tuommii/tuommii/_build/latest?definitionId=1&branchName=master)
![build](https://github.com/tuommii/scoreboard/workflows/build/badge.svg?branch=master)

<h1 align="center">
    <img src="/assets/screenshot1.png" height="200" />
    <img src="/assets/screenshot2.png" height="200" />
    <br><br>Scoreboard
</h1>

<h4 align="center">Simple score tracker for golf and disc golf. Made with Vue & Go :heart:</h4>
<br>

## :question: Motivation

It seems no one wants to be an accountant when we are playing so I wrote this so that the accountant can be switched in the middle of the game.

Created for mobile use. For best usability [add app to homescreen](https://www.howtogeek.com/196087/how-to-add-websites-to-the-home-screen-on-any-smartphone-or-tablet/).

## :dart: Features

- :heavy_check_mark: [Live](https://games.miikka.xyz/)
- :heavy_check_mark: Create pars **automatically** if the user is close to a [supported](https://github.com/tuommii/scoreboard/blob/master/assets/courses.json) course
- :heavy_check_mark: Accountant can be switched in the middle of the game
- :heavy_check_mark: Easy to use

## :memo: Technical features

- :heavy_check_mark: Demonstrates Go's concurrency features nicely
- :heavy_check_mark: Server can [save and restore](https://github.com/tuommii/scoreboard/blob/master/cmd/scoreboard/main.go#L40) it's state
- :heavy_check_mark: Each game in memory has it's own mutex instead of one global for all games
- :heavy_check_mark: [Automatic deployment](https://github.com/tuommii/scoreboard/blob/master/.github/workflows/build.yml#L1) with Github Actions
- :heavy_check_mark: JSON structure is designed to be easily upgradable
- :heavy_check_mark: Azure Pipelines & Docker
- :heavy_check_mark: Good rendering time
- :heavy_check_mark: Hosted on DigitalOcean behind nginx
- :heavy_check_mark: nginx gzip & cache (may vary during development)


## :question: How To Use It

### Local development
Clone repo

```
git clone https://github.com/tuommii/scoreboard.git
```

Start server

```
cd scoreboard
make
```

Start dev-server in a new terminal window

```
cd frontend
npm i
npm run serve
```

Go to: http://localhost:8081

### :trophy: Deploy

Every time when code is pushed to [production](https://github.com/tuommii/scoreboard/tree/production) branch, new binary will be compiled and sended to server.

<h1 align="left">
    <img src="/assets/cd.png" height="200" />
</h1>


### Docker
Build

`docker build . -t scoreboard`

Run

`docker run -p 8080:8080 scoreboard`


## Things to consider
_Contributions are more than welcome_
- [x] Create MVP
- [x] Circular navigation
- [x] Split frontend to Vue-components
- [x] Create my most played courses automatically based on location
- [x] Light theme because of sunshine
- [ ] Let user choose basket count
- [x] Make UI suitable for desktop also
- [ ] Use forced style guide for .vue-files
- [ ] Create courses automatically in Helsinki/Finland/World
- [ ] Add a lot of tests
- [ ] Frontend router
- [x] Github actions
- [ ] More server-side validations
- [ ] If project grows much, refactor to vuex
- [ ] PWA (Offline support)
- [x] User specific friends
- [ ] Fix typos
- [ ] UI improvements
- [x] Show spinner when waiting response
- [ ] More organized CSS
- [ ] CSS Animations
- [ ] Random or selected avatars
- [ ] Secure cookie
- [ ] Sound effects
- [ ] Copy id to clipboard

### Finishing game
- [ ] Create some graph
- [ ] For some users add possibility store game to database (validate time spent)
- [ ] When stats is implemented add support for marking OB (out of bounds)

## Example json
```json
{
  "id": "jt1",
  "basketCount": 1,
  "active": 1,
  "hasBooker": true,
  "baskets": {
    "1": {
      "orderNum": 1,
      "par": 3,
      "scores": {
        "Jian Yang": {
          "score": 3,
          "total": 0,
          "ob": 0
        },
        "Tiger King": {
          "score": 3,
          "total": 0,
          "ob": 0
        }
      }
    }
  },
  "createdAt": "2020-05-01T20:06:42.283050923+03:00",
  "editedAt": "2020-05-01T20:06:42.283052041+03:00",
  "name": "Default"
}
```
