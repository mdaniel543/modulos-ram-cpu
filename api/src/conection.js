const mysql = require('promise-mysql');

const config = require('./config');

const pool = mysql.createPool(config.database);

pool.getConnection().then(connection => {
    pool.releaseConnection(connection);
    console.log('DB is connected');
});

module.exports = pool;
