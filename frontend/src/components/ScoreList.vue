<template>
  <div class>
    <div class="card" v-for="(i, player) in course.baskets[1].scores" v-bind:key="player">
      <div class="card-content">
        <div class="content">
          <div class="columns is-mobile data">
            <div class="column">
              <span class="name">{{player}}</span>
            </div>
            <div class="column is-6">
              <button class="button round" @click="decScore(player)">-</button>
              <span class="current">{{course.baskets[course.active].scores[player].score}}</span>
              <button class="button round" @click="incScore(player)">+</button>
            </div>
            <div class="column is-2">
              <span class="sign" v-if="(totalScore(player) - totalPars()) >= 0">+</span>
              <span class="total">{{totalScore(player) - totalPars()}}</span>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  props: ["course"],
  data() {
    return {};
  },
  methods: {
    incScore(player) {
      this.$emit('incScore', player);
    },
    decScore(player) {
      this.$emit('decScore', player);
    },
    totalScore: function (name) {
      let total = 0
      for (let i = 1; i <= this.course.basketCount; i++) {
        total += this.course.baskets[i].scores[name].score
      }
      return total;
		},
    totalPars: function () {
      let total = 0
      for (let i = 1; i <= this.course.basketCount; i++) {
        total += this.course.baskets[i].par
      }
      return (total);
		}
  }
};
</script>

<style>
  .card {
    /* background: #eee; */
    color: #000 !important;
    background-image: linear-gradient(#fff, #eee) !important;
  }
</style>
