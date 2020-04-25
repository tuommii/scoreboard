
const CREATE_GAME = '/test_create';
const EDIT_GAME = '/test_edit';

let LAT = 0;
let LON = 0;

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
	if (!confirm('The games remain on the server for a few hours. You can still come back with ID.'))
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
	if (this.course.baskets[this.course.active].scores[player].score < 42) {
		this.course.baskets[this.course.active].scores[player].score++;
	}
}

function decScore(player) {
	if (this.course.baskets[this.course.active].scores[player].score > 1) {
		this.course.baskets[this.course.active].scores[player].score--;
	}
}

function incPar() {
	if (this.course.baskets[this.course.active].par < 5)
		this.course.baskets[this.course.active].par++;
}

function decPar() {
	if (this.course.baskets[this.course.active].par > 2)
		this.course.baskets[this.course.active].par--;
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
		basketCount: 18,
		lat: LAT,
		lon: LON
	};

	navigator.geolocation.getCurrentPosition((position) => {
			console.log('Geolocation permissions granted');
			query.lat = position.coords.latitude;
			query.lon = position.coords.longitude;
			postData(CREATE_GAME, query).then((data) => {
				this.course = data;
				localStorage.setItem('id', this.course.id);
				window.scrollTo({
					top: 0
				});
			});
	}, (err) => {
		console.log('Geolocation not supported. Creating Par 3 course.');
		postData(CREATE_GAME, query).then((data) => {
			this.course = data;
			localStorage.setItem('id', this.course.id);
			window.scrollTo({
				top: 0
			});
		});
	});
}

function saveAndNext() {
	this.course.action = "next";
	now = new Date().toJSON();
	this.course.editedAt = now;
	postData(EDIT_GAME, this.course).then((data) => {
		console.log('FROM SERVER:', data);
		this.course = data;
		window.scrollTo({
			top: 0
		});
		this.course.action = "";
	});
}


function saveAndPrev() {
	if (this.course.active === 1)
		return;
	this.course.action = "back";
	now = new Date().toJSON();
	this.course.editedAt = now;
	postData(EDIT_GAME, this.course).then((data) => {
		console.log('FROM SERVER:', data);
		this.course = data;
		window.scrollTo({
			top: 0
		});
		this.course.action = "";
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
				this.errors.join = 'ID Not Found'
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
			this.gameID = '';
		});
}

function getLocation() {
	if (navigator.geolocation) {
		navigator.geolocation.getCurrentPosition(getPosition, locationError);
	} else {
		console.log('Not supported');
	}
}

function getPosition(pos) {
	LAT = pos.coords.latitude;
	LON = pos.coords.longitude;
	console.log(LAT);
	console.log(LON);
}

function locationError(error) {
	console.log(error);
}

(function () {

	// TODO: Hide from user
	var app = new Vue({
		el: '#app',
		data: {
			errors: {
				start: '',
				add: '',
				join: ''
			},
			lat: 0,
			lon: 0,
			gameID: '',
			locked: 0,
			isDisabled: false,
			player: '',
			selectedPlayers: [
				{ name: 'Miikka', selected: true },
				{ name: 'Sande', selected: true },
				{ name: 'Pasi', selected: true },
				{ name: 'Joni', selected: false },
				{ name: 'Random', selected: false },
			],
			playersArr: [],
			course: {}
		},
		methods: {
			addPlayer: addPlayer,
			toggleSelected: toggleSelected,
			deletePlayer: deletePlayer,
			start: start,
			incScore: incScore,
			decScore: decScore,
			incPar: incPar,
			decPar: decPar,
			deleteGame: deleteGame,
			join: join,
			saveAndNext: saveAndNext,
			saveAndPrev: saveAndPrev,
			totalScore: function (name) {
				let total = 0
				for (let i = 1; i <= this.course.basketCount; i++) {
					total += this.course.baskets[i].scores[name].score
				}
				return total;
			},
			totalPars: function () {
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
		mounted: function () {
			// getLocation();
			const id = localStorage.getItem('id');
			if (id == null)
				return;
			const URL = `/games/${id}`;
			fetch(URL)
				.then((response) => {
					return response.json();
				})
				.then((data) => {
					this.course = data;
				});
		}
	});
})();
