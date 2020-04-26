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
            <!-- <div class="column is-2">
              <span class="sign" v-if="(totalScore(player) - totalPars()) >= 0">+</span>
              <span class="total">{{totalScore(player) - totalPars()}}</span>
            </div> -->
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
      console.log('INC', player);
    },
    decScore(player) {
      console.log('DEC', player);
    }
  }
};
</script>

<style>
</style>
