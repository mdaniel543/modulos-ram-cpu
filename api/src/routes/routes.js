const express = require("express");
const router = express.Router();

const process = require("../controllers/process.controller");
const ram = require("../controllers/ram.controller");
const cpu = require("../controllers/cpu.controller");

router.get("/ram/last", ram.getLastRam);

router.get("/ram", ram.getRam);

//--------------------------------------------

router.get("/last/process", process.getLastProcess);

router.get("/process/:id", process.getProcess);

router.get("/son/process/:id", process.getProcessSon);

router.get("/process", process.getProcesses);

router.get("/count/process", process.countProcess);

//--------------------------------------------

router.get("/cpu/last", cpu.getLastCPU);

router.get("/cpu", cpu.getCPU);

module.exports = router;