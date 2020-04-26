<template>
  <div class="form">

    <div class="field" v-for="player in players" v-bind:key="player.name">
      <div class="control">
        <button class="button is-medium width-160" v-bind:class="{ 'is-primary': player.selected}"
          @click="toggleSelected(player)">{{player.name}}
        </button>
        <a class="delete" @click="deletePlayer(player.name)"></a>
      </div>
    </div>

    <div class="field has-addons">
      <div class="control">
        <input type="text" class="input" placeholder="Name" v-model="name" />
        <p class="help err">{{err}}</p>
      </div>
      <div class="control">
        <button class="button is-link lila" @click="handleAdd">Add</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      players: [
        {name: 'Miikka', selected: true},
        {name: 'Pasi', selected: true},
        {name: 'Sande', selected: false},
        {name: 'Joni', selected: false},
      ],
      name: '',
      err: '',
    }
  },
  methods: {
    toggleSelected(player) {
      player.selected = !player.selected;
    },
    deletePlayer(name) {
	    this.players = this.players.filter(player => {
		    return player.name != name;
      });
    },
    handleAdd() {
      // setTimeout(() => {
      //   this.err = '';
      // }, 3000);
      this.players.push({ name: this.name, selected: true });
	    this.name = '';
      console.log('ADD');
    }
  },
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
.width-160 {
  min-width: 160px;
}
.delete {
  margin-left: 10px;
}
</style>
