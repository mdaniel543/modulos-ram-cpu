import React, { useState } from "react";
import Button from "react-bootstrap/Button";
import Offcanvas from "react-bootstrap/Offcanvas";

import TableProcess from "./TableProcess";

function OffCanvasTable({ show, onHide, data }) {
  return (
    <Offcanvas style={{height:"32rem"}} show={show} onHide={onHide} placement={"bottom"}>
      <Offcanvas.Header closeButton>
        <Offcanvas.Title>Hijos de {data.id}</Offcanvas.Title>
      </Offcanvas.Header>
      <Offcanvas.Body>
        <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTVlmWueUsfk-7eDhFEOyKZAUkQVLj2fpWqSw&usqp=CAU"></img>
      </Offcanvas.Body>
    </Offcanvas>
  );
}

export default OffCanvasTable;