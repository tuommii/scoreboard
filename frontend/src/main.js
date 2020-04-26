import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false

const JOIN_GAME = 'joinGame';

new Vue({
  render: h => h(App)
}).$mount('#app')
