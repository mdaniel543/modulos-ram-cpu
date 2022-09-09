const pool = require("../conection.js");

async function getProcess(req, res) {
  const { id } = req.params;
  const connection = await pool.getConnection();
  const process = await connection.query("SELECT * FROM process WHERE id = ?", [
    id,
  ]);
  connection.release();
  res.json(process);
}

async function getProcesses(req, res) {
  const connection = await pool.getConnection();
  const processes = await connection.query(
    "SELECT pid, name, user, state, memory FROM process WHERE pid_padre IS NULL and pid != 0"
  );
  connection.release();
  res.json(processes);
}

async function getProcessSon(req, res) {
  const { id } = req.params;
  const connection = await pool.getConnection();
  const process = await connection.query(
    "SELECT * FROM process WHERE pid_padre = ?",
    [id]
  );
  connection.release();
  res.json(process);
}

async function getLastProcess(req, res) {
  const connection = await pool.getConnection();
  const process = await connection.query(
    "SELECT * FROM process ORDER BY id DESC LIMIT 1"
  );
  connection.release();
  res.json(process);
}

async function countProcess(req, res) {
  const connection = await pool.getConnection();
  const process = await connection.query(`
    SELECT 'Procesos en ejecución' as 'Tipo de procesos', count(*) as Conteo FROM process WHERE (state = 0 or state = 2) and pid != 0
    UNION ALL
    SELECT 'Procesos suspendidos', count(*) FROM process WHERE state = 1 or state = 1026
    UNION ALL
    SELECT 'Procesos detenidos', count(*) FROM process WHERE state = 5
    UNION ALL
    SELECT 'Procesos zombie', count(*) FROM process WHERE state = 4
    UNION ALL
    SELECT 'Total de procesos', count(*) FROM process
    `);
  connection.release();
  res.json(process);
}

module.exports = {
  getProcess,
  getProcesses,
  getProcessSon,
  getLastProcess,
  countProcess,
};
