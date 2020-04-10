
const CREATE_GAME = '/test_create';

// We test with this
const edit = {
	"id": "1mp",
	"basketCount": 3,
	"active": 42,
	"baskets": {
		"1": {
			"orderNum": 1,
			"par": 3,
			"scores": {
				"Miikka": {
					"score": 0,
					"ob": 0
				},
				"Player 2": {
					"score": 0,
					"ob": 0
				}
			}
		},
		"2": {
			"orderNum": 2,
			"par": 3,
			"scores": {
				"Miikka": {
					"score": 0,
					"ob": 0
				},
				"Player 2": {
					"score": 0,
					"ob": 0
				}
			}
		},
		"3": {
			"orderNum": 3,
			"par": 3,
			"scores": {
				"Miikka": {
					"score": 0,
					"ob": 0
				},
				"Player 2": {
					"score": 0,
					"ob": 0
				}
			}
		}
	}
};


// With this we can create game
const startingData = {
	basketCount: 3,
	players: ['Tiger King', 'Ying Jang']
};

// TODO: ignore case
// TODO: Show error if same
function isUniq(name, arr) {
	code = true;
	arr.forEach(player => {
		console.log(name, player.name);

		if (name === player.name)
		{
			code = false;
			return code;
		}
	});
	return code;
}

function addPlayer(e) {
	e.preventDefault();
	// this.playersArr.push(this.player);
	if (this.player.length < 1) {
		this.errors.add = 'At least one character needed';
		return ;
	}

	else if (this.player.length > 16) {
		this.errors.add = 'Max length is 16';
		return ;
	}

	else if (!isUniq(this.player, this.selectedPlayers)) {
		this.errors.add = 'Player already exists';
		return ;
	}

	this.selectedPlayers.push({name: this.player, selected: true});
	this.player = '';
}

function toggleSelected(player) {
	player.selected = !player.selected
}

function deletePlayer(name) {
	this.selectedPlayers = this.selectedPlayers.filter(player => {
		console.log(player);
		return player.name != name;
	})
	// console.log(player, i);
}

function start() {
	this.playersArr = [];
	this.selectedPlayers.forEach(player => {
		if (player.selected) {
			this.playersArr.push(player.name);
		}
	});

	if (!this.playersArr.length) {
		this.errors.start = "At least one player must be selected"
		return ;
	}

	else if (this.playersArr.length > 5) {
		this.errors.start = "Max 5 players"
		return ;
	}

	const query = {
		players: this.playersArr,
		basketCount: 3
	};

	postData(CREATE_GAME, query).then((data) => {
		console.log(data);
		this.course = data;
		this.active = 1;
		// window.location.pathname = 'games/' + this.course.id + '/' + this.course.active;
	});
}

function sendData() {
	postData('/test_edit', edit).then((data) => {
		console.log('FROM SERVER:', data);
		this.course.active += 1;
		this.course = data;
	});
}

function join(e) {
	e.preventDefault();
}


var app = new Vue({
	el: '#app',
	data: {
		errors: {
			start: '',
			add: ''
		},
		active: 0,
		player: '',
		// TODO: Get from server
		selectedPlayers: [
			{name: 'Miikka', selected: true},
			{name: 'Player 2', selected: true},
			{name: 'Player 3', selected: false},
		],
		playersArr: [],
		// Game object
		course: {}
	},
	methods: {
		addPlayer: addPlayer,
		toggleSelected: toggleSelected,
		deletePlayer: deletePlayer,
		start: start,
		sendData: sendData
	},
	computed: {
		selectedCount() {
			let count = 0;
			this.selectedPlayers.forEach(player => {
				if (player.selected) {
					count++;
				}
			});
			return count;
		}
	}
});


async function postData(url = '', data = {}) {
	const response = await fetch(url, {
		method: 'POST',
		mode: 'cors',
		cache: 'no-cache',
		credentials: 'same-origin',
		headers: {
			'Content-Type': 'application/json'
		},
		redirect: 'follow',
		referrerPolicy: 'no-referrer',
		body: JSON.stringify(data)
	});
	return response.json();
}
