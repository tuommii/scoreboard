<template>
  <div class="form">
    <div class="control">
      <input type="text" class="input" placeholder="ID" v-model="gameID">
      <!-- <p class="help err">{{}}</p> -->
    </div>
    <div class="control">
      <button class="button" @click="handleJoin">Join</button>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    handleJoin() {
      if (!this.$store.state.gameID) {
        return ;
      }
      console.log(this.$store.state.gameID);
      this.$store.commit('updateGameID', '');
    }
  },
  computed: {
    gameID: {
      get() {
        return this.$store.state.gameID;
      },
      set(value) {
        this.$store.commit('updateGameID', value);
      }
    }
  },
  mounted() {
    console.log(this.$store.state.test.id);
  }
}
</script>
