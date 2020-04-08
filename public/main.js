console.log('Hello!');


const button = document.querySelector('#start');


// Example POST method implementation:
async function postData(url = '', data = {}) {
	// Default options are marked with *
	const response = await fetch(url, {
		method: 'POST', // *GET, POST, PUT, DELETE, etc.
		mode: 'cors', // no-cors, *cors, same-origin
		cache: 'no-cache', // *default, no-cache, reload, force-cache, only-if-cached
		credentials: 'same-origin', // include, *same-origin, omit
		headers: {
			'Content-Type': 'application/json'
			// 'Content-Type': 'application/x-www-form-urlencoded',
		},
		redirect: 'follow', // manual, *follow, error
		referrerPolicy: 'no-referrer', // no-referrer, *client
		body: JSON.stringify(data) // body data type must match "Content-Type" header
	});
	return response.json(); // parses JSON response into native JavaScript objects
}

let players = ["Miikka", "Joni", "Pasi"];

/**
 * Creates a new course object
 *
 * @param  {string[]}  players
 * @param  {number}  lanes
 *
 * @return {Object} Course object
*/
function createCourse(players, lanes) {
  let course = {

  };
  course["startedAt"] = Date.now();

  let i = 1;
  while (i <= lanes) {
    let j = 0;
    let score = {
    };
    while (j < players.length) {
      score[players[j]] = {
        score: 3,
        ob: false
      }
      j++;
      course[i] = score;
    }
    i++;
  }
  return course;
}

function calcPlayerScore(course, players, lanes) {
  let i = 1;
  // let total = 0;
  let stats = {};
  let j = 0;
  while (j < players.length) {
    stats[players[j]] = 0;
    j++;
  }
  while (i <= lanes) {
    j = 0;
    while (j < players.length) {
      // stats[players[j]].total = 0;
      stats[players[j]] += course[i][players[j]].score;
      j++;
    }
    // total += course[i][name].score;
    i++;
  }
  return stats;
}

const course = createCourse(players, 18);

console.log(course);
console.log(calcPlayerScore(course, players, 18));

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


const startingData = {
	basketCount: 3,
	players: ['Tiger King', 'Ying Jang']
};

button.addEventListener('click', (event) => {
	postData('/test_edit', edit).then((data) => {
		console.log(data);
	});
});
