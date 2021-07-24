var ctx1 = document.getElementById("chart1");
var chart1 = new Chart(ctx1, {
    type: "bar",
    data: {
        labels: intervalDataLabel,
        datasets: [
            {
                label: "Duration",
                data: intervalData,
                backgroundColor: Array(intervalData.length).fill("rgba(255, 99, 132, 0.2)"),
                borderWidth: 1,
            },
        ],
    },
    options: {
        scales: {
            y: {
                beginAtZero: true,
            },
        },
        title: {
            display: true,
            text: "Duration Distribution",
        }
    },
});

var ctx2 = document.getElementById("chart2");
var chart2 = new Chart(ctx2, {
    type: "bar",
    data: {
        labels: monthlyDataLabel,
        datasets: [
            {
                label: "Month",
                data: monthlyData,
                backgroundColor: Array(intervalData.length).fill("rgba(54, 162, 235, 0.2)"),
                borderWidth: 1,
            },
        ],
    },
    options: {
        scales: {
            y: {
                beginAtZero: true,
            },
        },
        title: {
            display: true,
            text: "Monthly Distribution",
        }
    },
});

var ctx3 = document.getElementById("chart3");
var chart3 = new Chart(ctx3, {
    type: "bar",
    data: {
        labels: hourDataLabel,
        datasets: [
            {
                label: "Hour",
                data: hourData,
                backgroundColor: Array(hourDataLabel.length).fill("rgba(255, 206, 86, 0.2)"),
                borderWidth: 1,
            },
        ],
    },
    options: {
        scaleShowValues: true,
        scales: {
            y: {
                beginAtZero: true,
            },
            xAxes: [{
                ticks: {
                    autoSkip: false
                }
            }]
        },
        title: {
            display: true,
            text: "Hourly Distribution",
        },
        plugins: {
            datalabels: {
                display: false
            }
        }
    },
});

var ctx4 = document.getElementById("top_month");
var chart4 = new Chart(ctx4, {
    type: "bar",
    data: {
        labels: monthTopMonths,
        datasets: [
            {
                label: "Top 1",
                data: monthTop1Count,
                backgroundColor: Array(intervalData.length).fill("rgba(75, 192, 192, 0.2)"),
                borderWidth: 1,
            },
            {
                label: "Top 2",
                data: monthTop2Count,
                backgroundColor: Array(intervalData.length).fill("rgba(153, 102, 255, 0.2)"),
                borderWidth: 1,
            },
            {
                label: "Top 3",
                data: monthTop3Count,
                backgroundColor: Array(intervalData.length).fill("rgba(255, 159, 64, 0.2)"),
                borderWidth: 1,
            },
        ],
    },
    options: {
        scales: {
            y: {
                beginAtZero: true,
            },
        },
        title: {
            display: true,
            text: "Monthly Top 3 Artists",
        },
        plugins: {
            tooltip: {
                callbacks: {
                    label: function (context) {
                        var d = context.dataIndex
                        var s = context.datasetIndex

                        var label = ""
                        if (s == 0) {
                            label = monthTop1Artis[d]
                        }
                        if (s == 1) {
                            label = monthTop2Artis[d]
                        }
                        if (s == 2) {
                            label = monthTop3Artis[d]
                        }

                        return label;
                    }
                }
            },
            datalabels: {
                align: 'end',
                anchor: 'end',
                rotation: -90,
                formatter: function (value, context) {
                    var d = context.dataIndex
                    var s = context.datasetIndex

                    var label = ""
                    if (s == 0) {
                        label = monthTop1Artis[d]
                    }
                    if (s == 1) {
                        label = monthTop2Artis[d]
                    }
                    if (s == 2) {
                        label = monthTop3Artis[d]
                    }

                    return label;
                }
            }
        }
    },
});
