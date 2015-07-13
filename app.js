var validator = require('is-my-json-valid');
var fs = require("fs");



// Load the content of a file to a string.
function loadStringFromFile(fileName) {
	var res = fs.readFileSync(fileName);
	if (res) {
		// console.log("" + res);
		return "" + res;
	}
}

function parseJson(json) {
	var obj;
	try {
		obj = JSON.parse(loadStringFromFile("./schemas/schema.json"));
	} catch(err) {
		obj = nil
	}
	return obj;
}


// Load the json-schema schema.
var sSch = JSON.parse(loadStringFromFile("./schemas/schema.json"));
 
var validate = validator(
	sSch,
	{verbose: true, greedy: true});
	
// console.log(validate.errors);

var s = loadStringFromFile("./schemas/moven-dataservice-schema.json");
console.log('schema loaded');	

var sObj = JSON.parse(s);	
console.log('schema parsed');	
	
validate(sObj);

console.log(validate.errors);


