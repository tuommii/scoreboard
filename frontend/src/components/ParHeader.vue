<template>
  <div>
    <div class="columns is-multiline is-mobile has-text-centered heading">
      <div class="column is-4">
        <p class="heading">Lane</p>
      </div>
      <div class="column is-4">
        <p class="heading">{{name}}: {{id}}</p>
      </div>
      <div class="column is-4">
        <p class="heading">Par</p>
      </div>

      <div class="column is-4">
        <p class="title is-5 light">{{active}}</p>
      </div>
      <div class="column is-4">
        <a @click="handleExit">
          <span class="icon exit">
            <i class="fas fa-times-circle"></i>
          </span>
        </a>
      </div>
      <div class="column is-4">
        <p class="title is-5 light">
          <span class="par par-minus" @click="handleDec">-</span>
          {{par}}
          <span class="par par-plus" @click="handleInc">+</span>
        </p>
      </div>
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

<style>
</style>
