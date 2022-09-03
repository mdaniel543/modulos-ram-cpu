const pool = require('../conection.js');

async function getProcess(req, res) {
    const { id } = req.params;
    const connection = await pool.getConnection();
    const process = await connection.query('SELECT * FROM process WHERE id = ?', [id]);
    connection.release();
    res.json(process);
}

async function getProcesses(req, res) {
    const connection = await pool.getConnection();
    const processes = await connection.query('SELECT * FROM process');
    connection.release();
    res.json(processes);
}

async function getProcessSon(req, res) {
    const { id } = req.params;
    const connection = await pool.getConnection();
    const process = await connection.query('SELECT * FROM process WHERE pid_padre = ?', [id]);
    connection.release();
    res.json(process);
}

async function getLastProcess(req, res) {
    const connection = await pool.getConnection();
    const process = await connection.query('SELECT * FROM process ORDER BY id DESC LIMIT 1');
    connection.release();
    res.json(process);
}

module.exports = {
    getProcess,
    getProcesses,
    getProcessSon,
    getLastProcess
}