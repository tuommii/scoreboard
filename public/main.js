console.log('Hello!');


const button = document.querySelector('#start');


// Example POST method implementation:
async function postData(url = '', data = {}) {
	// Default options are marked with *
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
			"score": 5,
			"ob": 0
		},
		"Jing Jang": {
			"score": 6,
			"ob": 1
		}
	},
	"2": {
		"par": 4,
		"Tigerking": {
			"score": 3,
			"ob": 0
		},
		"Jing Jang": {
			"score": 4,
			"ob": 1
		}
	},
	"3": {
		"par": 5,
		"Tigerking": {
			"score": 3,
			"ob": 0
		},
		"Jing Jang": {
			"score": 4,
			"ob": 1
		}
	}
}

// With this we can create game
const startingData = {
	basketCount: 3,
	players: ['Tiger King', 'Ying Jang']
};

// Here we send startingData
button.addEventListener('click', (event) => {
	postData('/test_edit', edit).then((data) => {
		console.log(data);
	});
});
