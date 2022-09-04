import React from "react";
import {
  Chart as ChartJS,
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Filler,
  Legend,
} from "chart.js";
import { Line } from "react-chartjs-2";

ChartJS.register(
  CategoryScale,
  LinearScale,
  PointElement,
  LineElement,
  Title,
  Tooltip,
  Filler,
  Legend
);

function AreaChart({ data }) {
  const options = {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        position: "bottom",
      },
      title: {
        display: true,
        text: "Uso de RAM",
      },
    },
    scales: {
      
      y: {
        min: 0,
        max: 100,
        stepSize: 5,
      },
    },
  };

  const labels = data.map((item) => item.id);

  const datas = {
    labels,
    datasets: [
      {
        fill: true,
        label: "Porcentaje de uso de RAM",
        data: data.map((item) => item.percentage),
        borderColor: "rgb(53, 162, 235)",
        backgroundColor: "rgba(53, 162, 235, 0.5)",
      },
    ],
  };
  return (
    <div
      style={{
        marginTop: "rem",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        height: "40rem",
        width: "100%",
      }}
    >
      <Line options={options} data={datas} height="500%" />
    </div>
  );
}

export default AreaChart;
