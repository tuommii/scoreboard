<template>
  <div id="app">
    <div class="section">
      <!-- HOME -->
      <div v-if="!course.active">
        <JoinGame @joinGame="joinGame($event)" />
        <PlayersList @startGame="startGame($event)" />
      </div>

      <!-- SCOREBOARD -->
      <div class="scoreboard" v-else>
        <ParHeader
          @incPar="incPar"
          @decPar="decPar"
          :par="course.baskets[course.active].par"
          :active="course.active"
          :id="course.id"
          :name="course.name" />
        <ScoreList
        @incScore="incScore"
        @decScore="decScore"
        :course="course"/>
        <Navigation @navigate="navigate" :active="course.active" :basketCount="course.basketCount" />
      </div>
    </div>
  </div>
</template>

<script>
import JoinGame from "./components/JoinGame.vue";
import PlayersList from "./components/PlayersList.vue";
import ParHeader from "./components/ParHeader.vue";
import ScoreList from "./components/ScoreList.vue";
import Navigation from "./components/Navigation.vue";
import "../node_modules/bulma/css/bulma.min.css";

const CREATE_GAME = "/test_create";
// const EDIT_GAME = "/test_edit";

export default {
  name: "App",
  data() {
    return {
      course: {}
    };
  },
  components: {
    JoinGame,
    PlayersList,
    ParHeader,
    ScoreList,
    Navigation
  },
  methods: {
    joinGame(id) {
      console.log("PARENT:", id);
    },
    createGame(query) {
      postData(CREATE_GAME, query).then(data => {
        this.course = data;
        // localStorage.setItem('id', this.course.id);
        console.log(data);
        window.scrollTo({
          top: 0
        });
      });
    },
    startGame(players) {
      console.log(players);
      const query = {
        players: players,
        basketCount: 18,
        lat: 0,
        lon: 0
      };

      navigator.geolocation.getCurrentPosition(
        (pos) => {
          query.lat = pos.coords.latitude;
          query.lon = pos.coords.longitude;
          this.createGame(query);
        },
        (err) => {
          console.log(err);
          this.createGame(query);
        }
      );
    },
    navigate(num) {
      console.log('Navigate', num);
      this.course.active = num;
    },
    incPar() {
      this.course.baskets[this.course.active].par++;
    },
    decPar() {
      this.course.baskets[this.course.active].par--;
    },
    incScore(player) {
      this.course.baskets[this.course.active].scores[player].score++;
    },
    decScore(player) {
      this.course.baskets[this.course.active].scores[player].score--;
    }
  }
};

async function postData(url = "", data = {}) {
  const response = await fetch(url, {
    method: "POST",
    mode: "cors",
    cache: "no-cache",
    credentials: "same-origin",
    headers: {
      "Content-Type": "application/json"
    },
    redirect: "follow",
    referrerPolicy: "no-referrer",
    body: JSON.stringify(data)
  });
  return response.json();
}
</script>

<style>
#app {
  font-family: -apple-system, BlinkMacSystemFont, Ubuntu, "Segoe UI", Roboto, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  /* margin-top: 60px; */
  background: #222;
  background: radial-gradient(#1e6a84, #033148);
  min-height: 100vh;
}
.section {
  padding-top: 1rem;
}
</style>
