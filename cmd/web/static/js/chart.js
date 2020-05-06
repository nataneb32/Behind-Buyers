
function createAnalyticsChart(ctx){
  const chart = new Chart(ctx,{
    type: 'line',
        data: {
          datasets: [{
            label: "Access",
          borderColor: 'blue',
          data: data
          }]
        },
        options: {
          maintainAspectRatio: false,
          scales: {
            xAxes: [{
              type: 'time',
              time: {
                unit: 'minute',
                round: "second",
              },
              distribution: 'linear'
            }],
            yAxes: [{
            ticks: {
                beginAtZero: true
              }
          }]
          }
        }
  })
  return chart
}