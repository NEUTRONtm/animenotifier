global.arn = require('./lib')
global.app = require('aero')()
global.HTTP = require('http-status-codes')
arn.maintenance = true
app.on('database ready', db => {
	global.db = db
})

// For POST requests
app.use(require('body-parser').json())

// Start the server
app.run()