const pool = require("../conection.js");

async function getCPU(req, res) {
  const connection = await pool.getConnection();
  const cpu = await connection.query("SELECT * FROM cpu");
  connection.release();
  res.json(cpu);
}

async function getLastCPU(req, res) {
  const connection = await pool.getConnection();
  const cpu = await connection.query(
    "SELECT * FROM cpu ORDER BY id DESC LIMIT 1"
  );
  connection.release();
  res.json(cpu);
}

module.exports = {
    getCPU,
    getLastCPU,
};