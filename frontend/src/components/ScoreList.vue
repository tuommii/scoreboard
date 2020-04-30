<template>
<div>
    <!-- <div class="card" v-for="(i, player) in course.baskets[1].scores" v-bind:key="player">
      <div class="card-content">
        <div class="content">
              <span class="name">{{player}}</span>
                <button class="button round" @click="decScore(player)">-</button>
                <span class="current">{{course.baskets[course.active].scores[player].score}}</span>
                <button class="button round" @click="incScore(player)">+</button>
                <span class="sign" v-if="(totalScore(player) - totalPars()) >= 0">+</span>
                <span class="total">{{totalScore(player) - totalPars()}}</span> -->
<div class="columns is-mobile score-card" v-for="(i, player) in course.baskets[1].scores" v-bind:key="player">
  <div class="column name-col is-4">
    <span class="name-val">{{player}}</span>
  </div>
  <div class="column cur-col">
    <div class="columns is-mobile">
      <div class="column minus-col">
        <button class="button round" @click="decScore(player)">-</button>
      </div>
      <div class="column score-col">
        <span class="current">{{course.baskets[course.active].scores[player].score}}</span>
      </div>
      <div class="column plus-col">
        <button class="button round" @click="incScore(player)">+</button>
      </div>
    </div>
  </div>
  <div class="column is-2 total-col">
    <span class="sign" v-if="(totalScore(player) - totalPars()) >= 0">+</span>
    <span class="total">{{totalScore(player) - totalPars()}}</span>
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
.score-card {
  /* background: #eee; */
  color: #000 !important;
  background: #fff;
  background-image: linear-gradient(#f1f2f3, #eee) !important;
  border-radius: 4px;
  border-top: 1px solid #fff;
  border-bottom: 1px solid rgba(0, 0, 0, 0.8);
  margin-top: 1rem !important;
}

.name-col, .score-col, .total-col {
  margin-top: calc(.5em - 1px);
}

.name-col {
  background: #eee;
  border-right: 1px solid #ccc;

}

.total-col {
  text-align: right;
}

.score-col {
  text-align: center;
}

.score {
  text-align: center;
  border: 0px solid #000;
}

.current {
  line-height: 1rem;
  font-size: 1.5rem;
  font-weight: 700;
}

.round {
    border-radius: 50% !important;
    min-width: 41px;
    padding-top: 0;
    margin-top: 0;
    /* border: 1px solid rgba(0,0,0,0.9); */
    /* background: linear-gradient(135deg, #62a8ff, #5146ff); */
    background: #1199a2 !important;
    border: 0 !important;
    color: #fff;
}
</style>
