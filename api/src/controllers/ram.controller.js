const pool = require("../conection.js");

async function getRam(req, res) {
  const connection = await pool.getConnection();
  const ram = await connection.query("SELECT * FROM ram");
  connection.release();
  res.json(ram);
}

async function getLastRam(req, res) {
  const connection = await pool.getConnection();
  const ram = await connection.query(
    "SELECT * FROM ram ORDER BY id DESC LIMIT 1"
  );
  connection.release();
  res.json(ram);
}

module.exports = {
  getRam,
  getLastRam,
};
