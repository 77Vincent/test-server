import express from "express"
import data from "./data.js"

const app = express()

app.get('/', (req, res) => {
    res.render('index', data.index)
})

export default app
