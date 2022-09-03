import { useState } from "react";
import Table from "react-bootstrap/Table";
import Container from "react-bootstrap/esm/Container";
import Button from "react-bootstrap/esm/Button";
import OffCanvasTable from "./BottomProcess";

function TableProcess({ data }) {
  const columns = data[0] && Object.keys(data[0]);
  const [show, setShow] = useState(false);
  const [process, setProcess] = useState({});

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
      <Container style={{ marginTop: "3rem" }}>
        <Table striped bordered hover>
          <thead>
            <tr>
              {data[0] && columns.map((heading) => <th>{heading}</th>)}
              <th>Accion</th>
            </tr>
          </thead>
          <tbody>
            {data.map((row) => (
              <tr>
                {columns.map((column) => (
                  <td>{row[column]}</td>
                ))}
                <td>
                  <Button
                    size="sm"
                    variant="primary"
                    onClick={() => {
                      setProcess(row);
                      setShow(true);
                    }}
                  >
                    ver hijos
                  </Button>
                </td>
              </tr>
            ))}
          </tbody>
        </Table>
      </Container>
      <OffCanvasTable
        show={show}
        onHide={() => setShow(false)}
        data={process}
      />
    </>
  );
}

export default TableProcess;
