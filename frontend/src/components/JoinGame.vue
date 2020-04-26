<template>
  <div class="form">
    <div class="field has-addons">
      <div class="control">
        <input type="text" class="input" placeholder="ID" v-model="id" />
        <p class="help err">{{err}}</p>
        <p class="help">Your turn to keep the score.</p>
      </div>
      <div class="control">
        <button class="button is-link lila" @click="handleJoin">Join</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      id: '',
      err: '',
    }
  },
  methods: {
    handleJoin() {
      if (!this.id) {
        this.err = `ID can't be empty`;
        setTimeout(() => {
          this.err = '';
        }, 3000);
        return;
      }
      this.$emit('joinGame', this.id);
      this.id = '';
    }
  }
};
</script>

<style>
.err {
  color: red;
}
.lila {
    border-color: #6257c4 !important;
    background-color: #6257c4 !important;
    box-shadow: 0 4px 12px -2px rgba(98,87,196,.4) !important;
    font-weight: 700;
}
</style>
