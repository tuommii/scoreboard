
const CREATE_GAME = '/test_create';

// We test with this
const edit = {
	"id": "tigeri-jing-5",
	"basketCount": 3,
	"1": {
		"par": 3,
		"Tigerking": {
			"score": 6,
			"ob": 1
		},
		"Jing Jang": {
			"score": 1,
			"ob": 0
		}
	},
	"2": {
		"par": 4,
		"Tigerking": {
			"score": 7,
			"ob": 1
		},
		"Jing Jang": {
			"score": 1,
			"ob": 0
		}
	},
	"3": {
		"par": 5,
		"Tigerking": {
			"score": 8,
			"ob": 1
		},
		"Jing Jang": {
			"score": 1,
			"ob": 0
		}
	}
}

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
			console.log('SAMA');
			code = false;
			return code;
		}
	});
	return code;
}

function addPlayer(e) {
	e.preventDefault();
	// this.playersArr.push(this.player);
	if (this.player.length > 0 && isUniq(this.player, this.selectedPlayers)) {
		this.selectedPlayers.push({name: this.player, selected: true});
		this.player = '';
	}
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

var app = new Vue({
	el: '#app',
	data: {
		active: 0,
		player: '',
		// TODO: Get from server
		selectedPlayers: [
			{name: 'Miikka', selected: true},
			{name: 'Sande', selected: false},
			{name: 'Pesukarhu', selected: true},
		],
		playersArr: [],
		// Game object
		course: {}
	},
	methods: {
		addPlayer: addPlayer,
		toggleSelected: toggleSelected,
		deletePlayer: deletePlayer,
		start: start
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
