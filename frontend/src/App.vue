<template>
  <div id="app">
    <div class="section">

      <!-- HOME -->
      <div v-if="!course.active">
        <JoinGame @joinGame="joinGame($event)" />
        <PlayersList @startGame="startGame($event)" />
      </div>

      <!-- SCOREBOARD -->
      <div v-else>
        <h1>Started</h1>
      </div>

    </div>
  </div>
</template>

<script>
import JoinGame from "./components/JoinGame.vue";
import PlayersList from "./components/PlayersList.vue";
import "../node_modules/bulma/css/bulma.min.css";

const CREATE_GAME = "/test_create";
// const EDIT_GAME = "/test_edit";

export default {
  name: "App",
  data() {
    return {
      active: 0,
      course: {}
    };
  },
  components: {
    JoinGame,
    PlayersList
  },
  methods: {
    joinGame(id) {
      console.log("PARENT:", id);
    },
    startGame(players) {
      console.log(players);
      const query = {
        players: players,
        basketCount: 18,
        lat: 0,
        lon: 0
      };
      postData(CREATE_GAME, query).then(data => {
        this.course = data;
        // localStorage.setItem('id', this.course.id);
        console.log(data);
        window.scrollTo({
          top: 0
        });
      });
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
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
