'use strict'

const express = require('express');
const app = express();
const port = 8001;

var count = 0;

function version(req, res){
    count++;

    res.send(`NODE HTTP echo reply v1 -- ${count}`);
} 

app.get('/', version);
app.get('/version', version);

app.listen(port, () => {
    console.log(`Example app listening on port ${port}!`);
});