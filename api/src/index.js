const express = require('express');
const cors = require('cors');
const morgan = require('morgan');
const routes = require('./routes/routes');
const app = express();

app.use(morgan('dev'));
app.use(cors());
app.use(express.json());
app.use(express.urlencoded({ extended: false }));

app.get("/", (req,res)=>{
    res.json({"message":"ok"});
    console.log("this is a test");
});

app.use("/api", routes);

app.listen(3000,()=>{
    console.log("app listening on PORT:3000");
})