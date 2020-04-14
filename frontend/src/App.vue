<template>
  <div id="app">
    <img alt="Vue logo" src="./assets/logo.png">
    <HelloWorld msg="Haloo ei ole nero!"/>
    <button class="button">Klikkaa</button>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import '../node_modules/bulma/css/bulma.min.css'

export default {
  name: 'App',
  components: {
    HelloWorld
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
