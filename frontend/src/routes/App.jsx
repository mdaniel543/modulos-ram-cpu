import Tab from "react-bootstrap/Tab";
import Tabs from "react-bootstrap/Tabs";
import Container from "react-bootstrap/Container";
import DoughnutChart from "../components/DoughnutChart";
import AreaChart from "../components/AreaChart";
import axios from "axios";
import TableProcess from "../components/TableProcess";
import {useState, useEffect} from 'react'

function App() {
  const [ram, setRam] = useState({
    free: 0,
    used: 0,
    percentage: 0,
  })

  const [rams, setRams] = useState([])

  useEffect(() => {
    const resp = axios.get('http://localhost:3000/api/ram/last')
    resp.then((response) => {
      setRam(response.data[0])
      console.log(response.data[0])
    })
    const resp2 = axios.get('http://localhost:3000/api/ram')
    resp2.then((response) => {
      setRams(response.data)
    })
  }, [])


  return (
    <Container style={{ marginTop: "3rem" }}>
      <h1 style={{ textAlign: "center" }}> Monitoreo de recursos</h1>
      <Tabs
        defaultActiveKey="ram"
        id="fill-tab-example"
        className="mb-3"
        fill
        style={{ marginTop: "3rem" }}
      >
        <Tab eventKey="ram" title="RAM">
          <DoughnutChart data={ram} title={"RAM"}/>
          <h2 style={{ textAlign: "center", marginTop: "20px" }}> {ram.percentage}% </h2>
          <AreaChart data={rams}/>
        </Tab>
        <Tab eventKey="cpu" title="CPU"></Tab>
        <Tab eventKey="process" title="PROCESOS">
          <TableProcess data={rams}/>
        </Tab>
      </Tabs>
    </Container>
  );
}

export default App;
