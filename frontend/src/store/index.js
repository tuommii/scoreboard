import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
	state: {
		counter: 0,
		gameID: 'gameID from store',
		errors: {
			join: 'Game not found'
		},
		test: {
			id: 'JEEE!'
		}
	},
	getters: {},
	mutations: {
		inc: state => state.counter++,
		updateGameID(state, id) {
			state.gameID = id;
		}
	},
	actions: {}
});
