<template>
    <div class="columns is-multiline is-mobile heading has-text-centered">
      <div class="column is-4">
        <p class="lane-title">Lane</p>
      </div>
      <div class="column is-4">
        <p class="name-title">{{name}}</p>
      </div>
      <div class="column is-4">
        <p class="par-title">Par</p>
      </div>

      <div class="column is-4 lane">
        <p class="title is-5">{{active}}</p>
      </div>
      <div class="column is-4 name">
        <a @click="handleExit">
          <p class="id">{{id}}</p>
        </a>
      </div>
      <div class="column is-4 right">
        <p class="title is-5">
          <span class="par par-minus" @click="handleDec">-</span>
          {{par}}
          <span class="par par-plus" @click="handleInc">+</span>
        </p>
      </div>
    </div>
</template>

<script>
export default {
  props: ["par", "active", "name", "id"],
  methods: {
    handleExit() {
      console.log("EXIT");
    },
    handleInc() {
      if (this.par < 5) {
        this.$emit('incPar');
      }
    },
    handleDec() {
      if (this.par > 2) {
        this.$emit('decPar');
      }
    }
  }
};
</script>

<style scoped>
p {
  color: #fff;
}

.par-minus, .par-plus {
  color: #fff !important;
}

.column {
  padding-top: 0;
  padding-bottom: 0.25rem;
}

.heading {
  padding-top: 1rem;
}

.lane-title, .name-title, .par-title {
  color: #e2ff4b;
}

.left, .middle, .right {
  color :#fff;
}

.id {
  padding-top: 0.25rem;
}
</style>
