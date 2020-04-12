
const CREATE_GAME = '/test_create';
const EDIT_GAME = '/test_edit';

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
		},
		"sasas": {
			"sadasd": {
				"name": 10
			}
		}
	}
};



let GAME = {};



/*
**
**	CRUD FUNCTIONS
**
*/
function isUniq(name, arr) {
	arr.forEach(player => {
		if (name.toLowerCase() === player.name.toLowerCase()) {
			return false;
		}
	});
	return true;
}

function addPlayer(e) {
	e.preventDefault();
	if (this.player.length < 1) {
		this.errors.add = 'At least one character needed';
		return;
	}

	else if (this.player.length > 16) {
		this.errors.add = 'Max length is 16';
		return;
	}

	else if (!isUniq(this.player, this.selectedPlayers)) {
		this.errors.add = 'Player already exists';
		return;
	}

	this.selectedPlayers.push({ name: this.player, selected: true });
	this.player = '';
}

function deletePlayer(name) {
	this.selectedPlayers = this.selectedPlayers.filter(player => {
		return player.name != name;
	});
}

function deleteGame() {
	if (!confirm('This will delete whole game. Are you sure?'))
		return;
	localStorage.removeItem('id');
	this.course = {
		active: 0
	};
}


/*
**
**	UI RELATED
**
*/

function toggleSelected(player) {
	player.selected = !player.selected
}

function incScore(player) {
	this.course.baskets[this.course.active].scores[player].score++;
}

function decScore(player) {
	if (this.course.baskets[this.course.active].scores[player].score > 1) {
		this.course.baskets[this.course.active].scores[player].score--;
	}
}

function incPar() {
	this.course.baskets[this.course.active].par++;
}

function decPar() {
	this.course.baskets[this.course.active].par--;
}

function prev() {
	if (this.course.active > 1)
		this.course.active--;
}

function next() {
	if (this.course.active < this.course.basketCount)
	{
		this.course.active++;
	}
}

function nextToActive(course) {
	if (course.active < course.basketCount) {
		course.active++;
	}
}

/*
**
**	DATA
**
*/
function start() {
	this.playersArr = [];
	this.selectedPlayers.forEach(player => {
		if (player.selected) {
			this.playersArr.push(player.name);
		}
	});

	if (!this.playersArr.length) {
		this.errors.start = "At least one player must be selected"
		return;
	}

	else if (this.playersArr.length > 5) {
		this.errors.start = "Max 5 players"
		return;
	}

	const query = {
		players: this.playersArr,
		basketCount: 3
	};

	postData(CREATE_GAME, query).then((data) => {
		console.log(data.status);
		console.log(data);
		this.course = data;
		localStorage.setItem('id', this.course.id);
		// window.location.pathname = 'games/' + this.course.id + '/' + this.course.active;
	});
}

function sendData() {
	let jee = {};
	console.log('REQUEST WITH', jee, this.course);

	if (!confirm('This will save your current state to server so others can take lead.'))
		return ;

	postData(EDIT_GAME, this.course).then((data) => {
		console.log('FROM SERVER:', data);
		this.course = data;
		nextToActive(this.course);
		// if (this.course.active < this.course.basketCount) {
		// 	this.course.active++;
		// }
	});
}

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


function join(e) {
	e.preventDefault();

	if (!this.gameID.length)
		return;

	fetch('/games/' + this.gameID)
		.then((response) => {
			console.log(response.status);
			if (response.status != 200) {
				this.locked++;
				this.gameID = '';
				if (this.locked >= 3) {
					this.isDisabled = true;
				}
			}

			return response.json();
		})
		.then((data) => {
			this.course = data;
			console.log(data);
			localStorage.setItem('id', this.course.id);
		});
}



// TODO: Hide from user
var app = new Vue({
	el: '#app',
	data: {
		errors: {
			start: '',
			add: ''
		},
		gameID: '',
		locked: 0,
		isDisabled: false,
		player: '',
		// TODO: Get from server
		selectedPlayers: [
			{ name: 'Miikka', selected: true },
			{ name: 'Sande', selected: true },
			{ name: 'Pasi', selected: true },
			{ name: 'Joni', selected: false },
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
		sendData: sendData,
		incScore: incScore,
		decScore: decScore,
		incPar: incPar,
		decPar: decPar,
		deleteGame: deleteGame,
		join: join,
		prev: prev,
		next: next,
		testi: function(name) {
			// let prev = 0;
			let total = 0
			for (let i = 1; i <= this.course.basketCount; i++) {
				total += this.course.baskets[i].scores[name].score
			}
			// for (let key in this.course.baskets[1].scores) {
			// 	console.log(key);
			// 	this.course.baskets[1].scores[key].total = this.course.baskets[1].scores[key].score - this.course.baskets[1].par;
			// }

			// if (this.course.active > 1) {
			// 	prev += this.course.baskets[this.course.active - 1].scores[name].total;
			// }
			// return this.course.baskets[this.course.active].scores[name].score - this.course.baskets[this.course.active].par + prev;
			return total;
		},
		pars: function() {
			let total = 0
			for (let i = 1; i <= this.course.basketCount; i++) {
				total += this.course.baskets[i].par
			}
			return (total);
		}
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
	},
	created: function () {
		getLocation();
		const id = localStorage.getItem('id');
		if (id == null)
			return;
		const URL = `/games/${id}`;
		console.log('COOKIE', id);
		fetch(URL)
			.then((response) => {
				return response.json();
			})
			.then((data) => {
				this.course = data;
				// this.$forceUpdate();
				console.log(this.course);
			});
	}
});

function getLocation() {
  if (navigator.geolocation) {
    navigator.geolocation.getCurrentPosition(cb);
  } else {
	  console.log('Not supported');
  }
}

function cb(pos) {
	console.log(pos.coords.latitude);
	console.log(pos.coords.longitude);
}
