<template>
  <div id="app" v-bind:class="{ 'dark': course.active}" v-cloak>
    <loading :active.sync="isLoading"
        :can-cancel="true"
        :is-full-page="fullPage"></loading>
    <div class="section">
      <div class="container">
        <!-- HOME -->
        <div v-if="!course.active">
          <a href="/" class="logo">
            <img src="prize-144.png" width="60" height="60" />
            <h1 class="title is-5">Scoreboard</h1>
          </a>
          <JoinGame @joinGame="joinGame($event)" />
          <PlayersList @startGame="startGame($event)" />
          <p class="help">Source code on <a href="https://github.com/tuommii/scoreboard">Github</a></p>
        </div>

        <!-- SCOREBOARD -->
        <div class="scoreboard" v-else>
          <ParHeader
            @incPar="incPar"
            @decPar="decPar"
            @exit="exit"
            :par="course.baskets[course.active].par"
            :active="course.active"
            :id="course.id"
            :name="course.name"
          />
          <ScoreList @incScore="incScore" @decScore="decScore" :course="course" />
          <Navigation
            @navigate="navigate"
            :active="course.active"
            :basketCount="course.basketCount"
          />
        </div>
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

import Loading from 'vue-loading-overlay';
import 'vue-loading-overlay/dist/vue-loading.css';

import "../node_modules/bulma/css/bulma.min.css";

const BASE = "/games/";
const CREATE_GAME = BASE + "create";
const EDIT_GAME = BASE + "edit";
const EXIT_GAME = "/exit/"

export default {
  name: "App",
  data() {
    return {
      isLoading: false,
      course: {
      }
    };
  },
  components: {
    JoinGame,
    PlayersList,
    ParHeader,
    ScoreList,
    Navigation,
    Loading
  },
  methods: {
    joinGame(id) {
      if (!id.length) return;

      fetch(BASE + id)
        .then(response => {
          return response.json();
        })
        .then(data => {
          if (data.hasBooker) {
            return;
          }
          this.course = data;
          // TODO: Small window when someone can join
          this.course.hasBooker = true;
          localStorage.setItem("id", this.course.id);
        })
        .catch(() => {
        });
    },
    createGame(query) {
      this.isLoading = true;
      postData(CREATE_GAME, query).then(data => {
        this.course = data;
        this.isLoading = false;
        localStorage.setItem("id", this.course.id);
        window.scrollTo({
          top: 0
        });
      });
    },
    startGame(players) {
      const query = {
        players: players,
        basketCount: 18,
        lat: 0,
        lon: 0
      };

      navigator.geolocation.getCurrentPosition(
        pos => {
          query.lat = pos.coords.latitude;
          query.lon = pos.coords.longitude;
          this.createGame(query);
        },
        () => {
          this.createGame(query);
        }
      );
    },
    navigate(num) {

      // Trying prevent screen flashing. Only show loader after x time
      let gotResponse = false;
      setTimeout(() => {
        if (!gotResponse) {
          this.isLoading = true;
        }
      }, 100);

      this.course.editedAt = new Date().toJSON();
      postData(EDIT_GAME, this.course).then(data => {
        gotResponse = true;
        this.course = data;
        this.course.active = num;
        this.isLoading = false;
        window.scrollTo({ top: 0 });
      });
    },
    incPar() {
      this.course.baskets[this.course.active].par++;
    },
    decPar() {
      this.course.baskets[this.course.active].par--;
    },
    incScore(player) {
      if (this.course.baskets[this.course.active].scores[player].score < 42) {
        this.course.baskets[this.course.active].scores[player].score++;
      }
    },
    decScore(player) {
      if (this.course.baskets[this.course.active].scores[player].score > 1) {
        this.course.baskets[this.course.active].scores[player].score--;
      }
    },
    exit() {
      if (!confirm("The games remain on the server for a few hours. You can still come back with ID.")) {
        return;
      }

      let gotResponse = false;
      setTimeout(() => {
        if (!gotResponse) {
          this.isLoading = true;
        }
      }, 100);

      fetch(EXIT_GAME + this.course.id)
        .then(response => {
          return response.json();
        })
        .then(() => {
          gotResponse = true;
          this.isLoading = false;
          localStorage.removeItem("id");
          this.course = {
            active: 0,
          };
        });
    }
  },
  mounted() {
    const id = localStorage.getItem("id");
    if (id == null) return;
    const URL = `${BASE}${id}`;
    fetch(URL)
      .then(response => {
        return response.json();
      })
      .then(data => {
        this.course = data;
      });
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
  font-family: -apple-system, BlinkMacSystemFont, Ubuntu, "Segoe UI", Roboto,
    Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: #2c3e50;
  /* margin-top: 60px; */
  background: #eee;
  min-height: 100vh;
  max-width: 480px;
}

h1 {
  padding-left: 1rem;
}

a.logo {
  display: flex;
  align-items: center;
}

.dark {
  color: #2c3e50;
  /* margin-top: 60px; */
  background: #033148 !important;
  background: radial-gradient(#1e6a84, #033148) !important;
  min-height: 100vh;
}
.section {
  padding-top: 1rem;
}

@media only screen and (min-width: 481px) {
  #app {
    /* width: 480px; */
    margin: 0 auto;
    /* background: #eee; */
  }
  .section {
    width: 480px;
    margin: 0 auto;
  }
}
</style>
