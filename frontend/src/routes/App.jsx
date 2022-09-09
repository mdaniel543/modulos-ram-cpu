import Col from "react-bootstrap/Col";
import Nav from "react-bootstrap/Nav";
import Row from "react-bootstrap/Row";
import Tab from "react-bootstrap/Tab";
import Button from "react-bootstrap/esm/Button";
import DoughnutChart from "../components/DoughnutChart";
import AreaChart from "../components/AreaChart";
import axios from "axios";
import TableProcess from "../components/TableProcess";
import { useState, useEffect } from "react";
import "../styles/index.css";

function App() {
  const [ram, setRam] = useState({
    free: 0,
    used: 0,
    percentage: 0,
  });
  const [rams, setRams] = useState([]);
  const [cpu, setCpu] = useState({
    free: 0,
    used: 0,
    percentage: 0,
  });
  const [cpus, setCpus] = useState([]);
  const [process, setProcess] = useState([]);
  const [countProcess, setCountProcess] = useState([]);
  const [processSon, setProcessSon] = useState([]);

  const [tab_key, setTab_key] = useState("cero");
  const [showHijos, setShowHijos] = useState(false);

  const [manage_ram, setManage_ram] = useState(false);
  const [manage_cpu, setManage_cpu] = useState(false);
  const [manage_process, setManage_process] = useState(false);

  useEffect(() => {
    const resp = axios.get("http://localhost:3000/api/ram/last");
    resp.then((response) => {
      setRam(response.data[0]);
      console.log(response.data[0]);
    });
    const resp2 = axios.get("http://localhost:3000/api/ram");
    resp2.then((response) => {
      setRams(response.data);
    });
    const resp3 = axios.get("http://localhost:3000/api/count/process");
    resp3.then((response) => {
      console.log(response.data);
      setCountProcess(response.data);
    });
    const resp4 = axios.get("http://localhost:3000/api/process");
    resp4.then((response) => {
      setProcess(response.data);
    });
  }, []);

  return (
    <>
      <div style={{ marginRight: "2rem" }}>
        <Tab.Container
          bg="dark"
          variant="dark"
          id="left-tabs-example"
          activeKey={tab_key}
        >
          <Row>
            <Col sm={3}>
              <div className="Bar" style={{ backgroundColor: "#fff" }}>
                <Nav
                  variant="pills"
                  className="flex-column"
                  style={{ marginLeft: "5px" }}
                >
                  <Nav.Item style={{ marginTop: "10px" }}>
                    <Nav.Link
                      style={{ fontSize: 22 }}
                      onClick={() => {
                        if (
                          tab_key === "first" ||
                          (tab_key === "first-2" && manage_ram)
                        ) {
                          setTab_key("five");
                        }
                        setManage_ram(!manage_ram);
                      }}
                    >
                      🧮 RAM{" "}
                      {manage_ram ? (
                        <text style={{ fontSize: 16 }}>❌</text>
                      ) : (
                        <text style={{ fontSize: 16 }}></text>
                      )}
                    </Nav.Link>
                    {manage_ram ? (
                      <>
                        <Nav.Link
                          style={{ marginLeft: "20px", fontSize: 20 }}
                          eventKey="first"
                          onClick={() => setTab_key("first")}
                        >
                          ⚠️ Porcentaje
                        </Nav.Link>
                        <Nav.Link
                          style={{ marginLeft: "20px", fontSize: 20 }}
                          eventKey="first-2"
                          onClick={() => setTab_key("first-2")}
                        >
                          📊 Grafica de uso
                        </Nav.Link>{" "}
                      </>
                    ) : (
                      <></>
                    )}
                  </Nav.Item>
                  <Nav.Item style={{ marginTop: "10px" }}>
                    <Nav.Link
                      style={{ fontSize: 22 }}
                      onClick={() => {
                        if (
                          tab_key === "second" ||
                          (tab_key === "second-2" && manage_cpu)
                        ) {
                          setTab_key("five");
                        }
                        setManage_cpu(!manage_cpu);
                      }}
                    >
                      🚀 CPU{" "}
                      {manage_cpu ? (
                        <text style={{ fontSize: 16 }}>❌</text>
                      ) : (
                        <text style={{ fontSize: 16 }}></text>
                      )}
                    </Nav.Link>
                    {manage_cpu ? (
                      <>
                        <Nav.Link
                          style={{ marginLeft: "20px", fontSize: 20 }}
                          eventKey="second"
                          onClick={() => setTab_key("second")}
                        >
                          ⚠️ Porcentaje
                        </Nav.Link>
                        <Nav.Link
                          style={{ marginLeft: "20px", fontSize: 20 }}
                          eventKey="second-2"
                          onClick={() => setTab_key("second-2")}
                        >
                          📊 Grafica de uso
                        </Nav.Link>{" "}
                      </>
                    ) : (
                      <></>
                    )}
                  </Nav.Item>
                  <Nav.Item style={{ marginTop: "10px" }}>
                    <Nav.Link
                      style={{ fontSize: 22 }}
                      onClick={() => {
                        if (
                          tab_key === "third" ||
                          (tab_key === "third-2" && manage_process)
                        ) {
                          setTab_key("five");
                        }
                        setManage_process(!manage_process);
                      }}
                    >
                      🗃 Procesos{" "}
                      {manage_process ? (
                        <text style={{ fontSize: 16 }}>❌</text>
                      ) : (
                        <text style={{ fontSize: 16 }}></text>
                      )}
                    </Nav.Link>
                    {manage_process ? (
                      <>
                        <Nav.Link
                          style={{ marginLeft: "20px", fontSize: 20 }}
                          eventKey="third"
                          onClick={() => setTab_key("third")}
                        >
                          📊 Conteo de procesos
                        </Nav.Link>
                        <Nav.Link
                          style={{ marginLeft: "20px", fontSize: 20 }}
                          eventKey="third-2"
                          onClick={() => setTab_key("third-2")}
                        >
                          📈 Proceso Padres
                        </Nav.Link>
                        {showHijos ? (
                          <Nav.Link
                            style={{ marginLeft: "20px", fontSize: 20 }}
                            eventKey="third-3"
                            onClick={() => setTab_key("third-3")}
                          >
                            📉 Proceso Hijos
                          </Nav.Link>
                        ) : (
                          <></>
                        )}
                      </>
                    ) : (
                      <></>
                    )}
                  </Nav.Item>
                </Nav>
              </div>
            </Col>

            <Col sm={9}>
              <Tab.Content style={{ marginLeft: "-4rem" }}>
                <Tab.Pane eventKey="cero">
                  <div
                    style={{
                      display: "flex",
                      alignItems: "center",
                      justifyContent: "center",
                      marginTop: "18rem",
                    }}
                  >
                    <h1>Bienvenido 👋</h1>
                  </div>
                </Tab.Pane>
                <Tab.Pane eventKey="first">
                  {ram && (
                    <>
                      <DoughnutChart data={ram} title={"RAM"} />
                      <h2 style={{ textAlign: "center", marginTop: "20px" }}>
                        {" "}
                        {ram.percentage}%{" "}
                      </h2>
                    </>
                  )}
                </Tab.Pane>
                <Tab.Pane eventKey="first-2">
                  {rams && <AreaChart data={rams} />}
                </Tab.Pane>
                <Tab.Pane eventKey="second">{cpu && <></>}</Tab.Pane>
                <Tab.Pane eventKey="second-2">{cpus && <></>}</Tab.Pane>
                <Tab.Pane eventKey="third">
                  {countProcess && (
                    <>
                      <TableProcess data={countProcess} accion={false} />
                    </>
                  )}
                </Tab.Pane>
                <Tab.Pane eventKey="third-2">
                  {process && (
                    <>
                      <TableProcess
                        data={process}
                        accion={true}
                        show={() => {
                          setTab_key("third-3");
                          setShowHijos(true);
                        }}
                        setSons={setProcessSon}
                      />
                    </>
                  )}
                </Tab.Pane>
                <Tab.Pane eventKey="third-3">
                  {showHijos && (
                    <>
                      <Button
                        size="xl"
                        variant="outline-danger"
                        aria-label="Hide"
                        style={{ marginLeft: "95%", marginTop: "2%" }}
                        onClick={() => {
                          setTab_key("third-2");
                          setShowHijos(false);
                          setProcessSon([]);
                        }}
                      >
                        ❌
                      </Button>
                      {processSon.length > 0 ? (
                        <TableProcess data={processSon} accion={false} />
                      ) : (
                        <h1 style={{ marginTop: "15rem", textAlign: "center" }}>
                          No hay procesos hijos
                        </h1>
                      )}
                    </>
                  )}
                </Tab.Pane>
                <Tab.Pane eventKey="five">
                  <div
                    style={{
                      display: "flex",
                      alignItems: "center",
                      justifyContent: "center",
                      marginTop: "18rem",
                    }}
                  >
                    <h1>👈 Seleccione una opcion</h1>
                  </div>
                </Tab.Pane>
              </Tab.Content>
            </Col>
          </Row>
        </Tab.Container>
      </div>
    </>
  );
}

export default App;
