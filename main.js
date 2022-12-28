import express from "express"
import router from "./src/router.js"
import migrate from "./src/db.js"

const PORT = 8080
const HOST = '0.0.0.0'

// database setup
migrate()

const app = express();
app.set('views', 'views')
app.set('view engine', 'pug')
app.use(router)

app.listen(PORT, HOST, () => {
    console.log(`Running on http://${HOST}:${PORT}`);
});