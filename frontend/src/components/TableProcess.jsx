import { useState } from "react";
import Table from "react-bootstrap/Table";
import Container from "react-bootstrap/esm/Container";
import Button from "react-bootstrap/esm/Button";
import axios from "axios";


function TableProcess({ data, accion, show, setSons }) {
  const columns = data[0] && Object.keys(data[0]);

  const showHijos = (pid) => {
    console.log(pid);
    const resp = axios.get(`http://34.125.81.117:3000/api/son/process/${pid}`);
    resp.then((res) => {
      setSons(res.data);
      show();
    });
  };

  const style = {
    top: 0,
    left: 0,
    zIndex: 10,
    height: "2.5rem",
    position: "sticky",
    color: "#fff",
    backgroundColor: "#24242c ",
  };

  return (
    <>
      <Container style={{ marginTop: "2.5rem" }}>
        <div style={{ overflowY: "auto", height: "40rem" }}>
          <Table striped bordered hover>
            <thead>
              <tr>
                {data[0] &&
                  columns.map((heading) => <th style={style}>{heading}</th>)}
                {accion && <th style={style}>Accion</th>}
              </tr>
            </thead>
            <tbody>
              {data.map((row) => (
                <tr>
                  {columns.map((column) => (
                    <td>
                      {(() => {
                        if (column === "state") {
                          return(
                          
                            row[column] === 1 || row[column] === 1026 ? "suspendido" : 
                            row[column] === 0 || row[column] === 2 ? "activo" : 
                            row[column] === 5 ? "detenidos" : "zombie" 
                          );
                        } else {
                          return row[column];
                        }
                      })()}
                    </td>
                  ))}
                  {accion && (
                    <td>
                      <Button
                        size="sm"
                        variant="dark"
                        onClick={() => 
                          showHijos(row.pid)
                        }
                      >
                        ver hijos
                      </Button>
                    </td>
                  )}
                </tr>
              ))}
            </tbody>
          </Table>
        </div>
      </Container>
    </>
  );
}

export default TableProcess;
