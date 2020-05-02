# Scoreboard
Simple score tracker for golf and discgolf. Made with **Go** & **Vue.js**

Live demo: https://games.miikka.xyz/

Created for mobile use. For best usability [add app to homescreen](https://www.howtogeek.com/196087/how-to-add-websites-to-the-home-screen-on-any-smartphone-or-tablet/).

## Things to consider
_Contributions are more than welcome_
- [x] Create MVP
- [x] Circular navigation
- [x] Split frontend to Vue-components
- [x] Create most played courses automatically based on location
- [x] Light theme because of sunshine
- [ ] Make UI suitable for desktop also
- [ ] Use forced style guide for .vue-files
- [ ] Create courses automatically in Helsinki/Findland/World
- [ ] Add a lot of tests
- [ ] Frontend router
- [ ] Github actions
- [ ] More server-side validations
- [ ] If project grows much, refactor to vuex
- [ ] PWA (Offline support)
- [ ] User specific friends
- [ ] Fix typos
- [ ] UI improvements
- [ ] Show spinner when waiting response
- [ ] More organized CSS
- [ ] CSS Animations
- [ ] Random or selected avatars
- [ ] Secure cookie
- [ ] Sound effects
- [ ] Copy id to clipboard

### Finishing game
- [ ] Create some graph
- [ ] For some users add possibility store game to database
- [ ] When stats is implemented add support for marking OB (out of bounds)

## Dev
**Clone repo**

```
git clone https://github.com/tuommii/scoreboard.git
```

**Start server**

```
cd scoreboard
make
```

**Start dev-server in a new terminal window**

```
cd frontend
npm i
npm run serve
```

Go to: http://localhost:8081

## Example json
```json
{
  "id": "jt1",
  "basketCount": 1,
  "active": 1,
  "baskets": {
    "1": {
      "orderNum": 1,
      "par": 3,
      "scores": {
        "Jing Yang": {
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
