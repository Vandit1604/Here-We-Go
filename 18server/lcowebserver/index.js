/*
Part of exercise file for go lang course at
https://web.learncodeonline.in
*/

const express = require('express')
const app = express()
const port = 8000

app.use(express.json()); 
app.use(express.urlencoded({extended: true}));

app.get('/', (req, res) => {
  res.status(200).send("Hi Vandit is Learning Golang and is aiming to become a better developer.")
})

app.get('/get', (req, res) => {
    res.status(200).json({message: "Hello from localhost:8000"})
  })


app.post('/post', (req, res) => {
    let myJson = req.body;      // your JSON
	
	res.status(200).send(myJson);
})

app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
})
  

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})