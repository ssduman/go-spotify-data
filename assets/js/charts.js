var ctx1 = document.getElementById("chart1");
var chart1 = new Chart(ctx1, {
    type: "bar",
    data: {
        labels: intervalDataLabel,
        datasets: [
            {
                label: "Duration",
                data: intervalData,
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
                borderWidth: 1,
            },
            {
                label: "Top 2",
                data: monthTop2Count,
                borderWidth: 1,
            },
            {
                label: "Top 3",
                data: monthTop3Count,
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
