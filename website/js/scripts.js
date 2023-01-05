
const selectedIata = document.getElementById("selectedIata");
let temperatureChartObject;
let pressureChartObject;
let windChartObject;
const temperatureChart = document.getElementById('temperatureChart');
const windChart = document.getElementById('windChart');
const pressureChart = document.getElementById('pressureChart');

function addData(chart, label, data) {
    chart.data.labels.push(label);
    chart.data.datasets.forEach((dataset) => {
        dataset.data.push(data);
    });
    chart.update();
}

function removeData(chart) {
    chart.data.labels.splice(0, 1);
    chart.data.datasets.forEach((dataset) => {
        dataset.data.splice(0, 1);
    });
    chart.update();
}

function loadChart() {
    getDataByAirport(selectedIata.value).then(data => {
        temperatureChartObject = new Chart(temperatureChart, {
            type: 'line',
            data: {
                labels: data.temperature.map(function (data) {
                    return data.date
                }),
                datasets: [{
                    label: 'Temperature',
                    data: data.temperature.map(function (data) {
                        return parseFloat(data.value)
                    }),
                    borderWidth: 2,
                    pointRadius: 4,
                    tension: 0.3
                },
                ]
            },
            options: {
                plugins: {
                    legend: {
                        display: false,
                        labels: {
                            color: '#FF000',
                        }
                    }
                }
            }
        });
        windChartObject = new Chart(windChart, {
            type: 'line',
            data: {
                labels: data.wind.map(function (data) {
                    return data.date
                }),
                datasets: [{
                    label: 'Wind',
                    data: data.wind.map(function (data) {
                        return parseFloat(data.value)
                    }),
                    borderWidth: 2,
                    borderColor: '#1ef800',
                    backgroundColor: '#1ef800',
                    color: '#1ef800',
                    pointRadius: 4,
                    tension: 0.3
                },
                ],
            },
            options: {
            }
        });
        pressureChartObject = new Chart(pressureChart, {
            type: 'line',
            data: {
                labels: data.pressure.map(function (data) {
                    return data.date
                }),
                datasets: [{
                    label: 'Pressure',
                    data: data.pressure.map(function (data) {
                        return parseFloat(data.value)
                    }),
                    borderWidth: 2,
                    borderColor: '#f80000',
                    backgroundColor: '#f80000',
                    color: '#f80000',
                    pointRadius: 4,
                    tension: 0.3
                },
                ],
            },
        });

    });

}
async function getDataByAirport(iata) {
    let request = new Request(`http://localhost:8080/get/data-by-iata-code-and-number/${iata}/10`, {
        method: 'GET',
        headers: new Headers()
    });
    return await fetch(request).then(response => response.json()).then(data => { return { pressure: data.pressure, wind: data.wind, temperature: data.temperature } });
}

$(function () {

    class GaugeChart {
        constructor(element, params) {
            this._element = element;
            this._minValue = params.minValue;
            this._initialValue = params.initialValue;
            this._higherValue = params.higherValue;
            this._title = params.title;
            this._subtitle = params.subtitle;
            this._customTicks = params.customTicks;
        }

        _buildConfig() {
            let element = this._element;

            return {
                value: this._initialValue,
                valueIndicator: {
                    color: '#fff'
                },
                geometry: {
                    startAngle: 180,
                    endAngle: 360
                },
                scale: {
                    startValue: this._minValue,
                    endValue: this._higherValue,
                    customTicks: this._customTicks,
                    tick: {
                        length: 8
                    },
                    label: {
                        font: {
                            color: '#87959f',
                            size: 9,
                            family: '"Open Sans", sans-serif'
                        }
                    }
                },
                title: {
                    verticalAlignment: 'bottom',
                    text: this._title,
                    font: {
                        family: '"Open Sans", sans-serif',
                        color: '#fff',
                        size: 10
                    },
                    subtitle: {
                        text: this._subtitle,
                        font: {
                            family: '"Open Sans", sans-serif',
                            color: '#fff',
                            weight: 700,
                            size: 28
                        }
                    }
                },
                onInitialized: function () {
                    let currentGauge = $(element);
                    let circle = currentGauge.find('.dxg-spindle-hole').clone();
                    let border = currentGauge.find('.dxg-spindle-border').clone();

                    currentGauge.find('.dxg-title text').first().attr('y', 48);
                    currentGauge.find('.dxg-title text').last().attr('y', 28);
                    currentGauge.find('.dxg-value-indicator').append(border, circle);
                }

            }
        }

        init() {
            $(this._element).dxCircularGauge(this._buildConfig());
        }
    }

    $(document).ready(function () {
        function printAirportData() {
            const selectedIata = document.getElementById("selectedIata");
            getDataByAirport(selectedIata.value).then(value => {
                // Init chart data
                let titles = ["Temperature", "Pressure", "Wind"];
                let mesures = ["ºC", "hPa", "km/h"];
                let minValue = [-15, 950, 0];
                let maxValue = [45, 1030, 250];
                let initialValue = [value.temperature[value.temperature.length - 1].value, value.pressure[value.pressure.length - 1].value, value.wind[value.wind.length - 1].value];
                let customTicks = [[-15, -10, -5, 0, 5, 10, 20, 30, 45], [950, 960, 970, 980, 990, 1000, 1010, 1020, 1030], [0, 50, 100, 150, 200, 250]];

                $('.gauge').each(function (index, item) {
                    let params = {
                        initialValue: initialValue[index],
                        minValue: minValue[index],
                        higherValue: maxValue[index],
                        title: titles[index],
                        customTicks: customTicks[index],
                        subtitle: `${initialValue[index]} ${mesures[index]}`
                    };

                    let gauge = new GaugeChart(item, params);
                    gauge.init();
                });
                function updateData() {
                    getDataByAirport(selectedIata.value).then(value => {
                        if (temperatureChartObject.data.labels.length >= 9) {
                            removeData(temperatureChartObject);
                        }
                        addData(temperatureChartObject, value.temperature[value.temperature.length - 1].date, value.temperature[value.temperature.length - 1].value);
                        if (windChartObject.data.labels.length >= 9) {
                            removeData(windChartObject);
                        }
                        addData(windChartObject, value.wind[value.wind.length - 1].date, value.wind[value.wind.length - 1].value);
                        if (pressureChartObject.data.labels.length >= 9) {
                            removeData(pressureChartObject);
                        }
                        addData(pressureChartObject, value.pressure[value.pressure.length - 1].date, value.pressure[value.pressure.length - 1].value);
                        $('.gauge').each(function (index, item) {
                            let gauge = $(item).dxCircularGauge('instance');
                            let newValue;
                            switch (index) {
                                case 0:
                                    newValue = value.temperature[value.temperature.length - 1].value;
                                    break;
                                case 1:
                                    newValue = value.pressure[value.pressure.length - 1].value;
                                    break;
                                case 2:
                                    newValue = value.wind[value.wind.length - 1].value;
                                    break;
                            }
                            let gaugeElement = $(gauge._$element[0]);

                            gaugeElement.find('.dxg-title text').last().html(`${newValue} ${mesures[index]}`);
                            gauge.value(newValue);
                        });
                    });
                }
                $('#random').click(function () {
                    updateData();
                });
                loadChart();
                setInterval(() => { updateData() }, 5000);
            });
        }
        printAirportData();
        selectedIata.addEventListener("change", printAirportData);
    });


});