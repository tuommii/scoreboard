console.log('Hello!');

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

// Here we send startingData
// button.addEventListener('click', (event) => {
// 	postData('/test_edit', edit).then((data) => {
// 		console.log(data);
// 	});
// });


function toggleSelected(player) {
	player.selected = !player.selected
}

var app = new Vue({
	el: '#app',
	data: {
		active: 0,
		// TODO: Get from server
		selectedPlayers: [
			{name: 'Miikka', selected: true},
			{name: 'Sande', selected: false},
			{name: 'Pesukarhu', selected: true},
		],
		// Game object
		course: {}
	},
	methods: {
		toggleSelected: toggleSelected
	}
});
