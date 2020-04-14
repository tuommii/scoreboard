<template>
  <div class="form">
    <div class="control">
      ssasa
      <input type="text" class="input" placeholder="ID" v-model="gameID">
      <p class="help err">{{errors.join}}</p>
    </div>
    <div class="control">
      <button class="button" @click="handleJoin">Join</button>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      gameID: 'SAsas',
      errors: {
        join: 'Game not found'
      }
    }
  },
  methods: {
    handleJoin() {
      console.log('Join!');
    }
  }
}
</script>
