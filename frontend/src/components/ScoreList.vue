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
  /* background-image: linear-gradient(#f1f2f3, #eee) !important; */
  background: #fff;
  border-radius: 4px;
  border-top: 1px solid #fff;
  border-bottom: 1px solid rgba(0, 0, 0, 0.5);
  margin-bottom: 1rem;
}

.card-content {
  padding: 0.5rem 0;
  margin: 0;
}

.round {
    border-radius: 50% !important;
    min-width: 41px;
    /* border: 1px solid rgba(0,0,0,0.9); */
    /* background: linear-gradient(135deg, #62a8ff, #5146ff); */
    background: #1199a2;
    border: 0 !important;
    color: #fff;
}

</style>
