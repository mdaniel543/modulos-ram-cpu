import { Chart as ChartJS, ArcElement, Tooltip, Legend, Title, } from "chart.js";

import { Doughnut } from "react-chartjs-2";

ChartJS.register(ArcElement, Tooltip, Legend, Title);

function DoughnutChart({data , title}) {
  const datas = {
    labels: ["usado", "libre"],
    datasets: [
      {
        label: "Uso de " + title,
        data: [data.used, data.free],
        backgroundColor: ["rgba(75, 192, 192, 0.2)", "rgba(255, 206, 86, 0.2)"],
        borderColor: ["rgba(75, 192, 192, 1)", "rgba(255, 206, 86, 1)"],
        borderWidth: 1,
      },
    ],
  };

  const options = {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
      legend: {
        position: "bottom",
      },
      title: {
        display: true,
        text: "Porcentaje de uso de " + title + " " + data.percentage + "%",
      },
    },
  };

  return (
    <div
      style={{
        marginTop: "2rem",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        height: "100%",
        width: "100%",
      }}
    >
      <Doughnut data={datas} options={options} height="400%" />
    </div>
  );
}

export default DoughnutChart;
