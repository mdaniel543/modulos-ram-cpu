const express = require("express");
const router = express.Router();

const process = require("../controllers/process.controller");
const ram = require("../controllers/ram.controller");

router.get("/ram/last", ram.getLastRam);

router.get("/ram", ram.getRam);

router.get("/process/last", process.getLastProcess);

router.get("/process/:id", process.getProcess);

router.get("/process/:id/son", process.getProcessSon);

router.get("/process", process.getProcesses);

module.exports = router;