<template>
  <div id="app">
    <div class="section">
      <JoinGame @joinGame="joinGame($event)" />
      <PlayersList />
    </div>
  </div>
</template>

<script>
import JoinGame from './components/JoinGame.vue'
import PlayersList from './components/PlayersList.vue'
import '../node_modules/bulma/css/bulma.min.css'

export default {
  name: 'App',
  components: {
    JoinGame,
    PlayersList
  },
  methods: {
    joinGame(id) {
      console.log('PARENT:', id);
    }
  }
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
