
var curlcon = require("curlconverter");

ret = curlcon.toPython("curl --request POST --url https://open.workec.com/auth/accesstoken --header 'cache-control: no-cache' --header 'content-type: application/json' --data '{ 'appId': appId, 'appSecret': 'appSecret'}'")

console.log(ret)