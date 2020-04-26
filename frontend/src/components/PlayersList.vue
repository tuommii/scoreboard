<template>
  <div class="form">
      <div class="field" v-for="player in players" v-bind:key="player.name">
        <div class="control">
          <button
            class="button is-medium width-160"
            v-bind:class="{'is-primary': player.selected}"
            @click="toggleSelected(player)">
              {{player.name}}
          </button>
          <a class="delete" @click="deletePlayer(player.name)"></a>
        </div>
      </div>

      <div class="field">
        <div class="control">
          <p class="help selected-count">{{selectedCount}} of {{players.length}} selected</p>
          <button class="button is-link is-medium width-160" @click="handleStart">Start</button>
          <p class="help err">{{errors.start}}</p>
        </div>
      </div>

      <div class="field has-addons">
        <div class="control">
          <input type="text" class="input" placeholder="Name" v-model="name" />
          <p class="help err">{{errors.add}}</p>
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
      name: "",
      players: [
        { name: "Miikka", selected: true },
        { name: "Pasi", selected: true },
        { name: "Sande", selected: false },
        { name: "Joni", selected: false }
      ],
      errors: {
        start: "",
        add: ""
      }
    };
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
      if (this.name.length < 1) {
        this.showErr("add", "At least one character needed");
        return;
      } else if (this.name.length > 8) {
        this.showErr("add", "Max length is 8");
        return;
      }
      if (isUniq(this.name, this.players) === false) {
        this.showErr("add", "Player already exists");
        return;
      }
      this.players.push({ name: this.name, selected: true });
      this.name = "";
    },
    handleStart() {
      let selected = [];
      this.players.forEach(player => {
        if (player.selected) {
          selected.push(player.name);
        }
      });

      if (!selected.length) {
        this.showErr("start", "At least one player must be selected");
        return ;
      }
      else if (selected.length > 5) {
        this.showErr("start", "Max 5 players");
        return ;
      }
      this.$emit('startGame', selected);
    },
    showErr(field, msg) {
      this.errors[field] = msg;
      setTimeout(() => {
        this.errors[field] = "";
      }, 3000);
    }
  },
  computed: {
    selectedCount() {
      let count = 0;
      this.players.forEach(player => {
        if (player.selected) {
          count++;
        }
      });
      return count;
    }
  }
};

function isUniq(name, arr) {
  let uniq = true;
  arr.forEach(player => {
    if (name.toLowerCase() === player.name.toLowerCase()) {
      uniq = false;
      return;
    }
  });
  return uniq;
}
</script>

<style>

.width-160 {
  min-width: 160px;
}
.delete {
  margin-left: 10px;
}
</style>
