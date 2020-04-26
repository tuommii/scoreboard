module.exports = {
	lintOnSave: true,
	devServer: {
		proxy: 'http://localhost:8080'
	  },
	  pages: {
		index: {
		  entry: 'src/main.js',
		  title: 'Scoreboard'
		}
	  }
}
