function timeChart(id, ...data){
  var config = { 
      xaxis:{ 
        mode: "time",
        minTickSize: [1, "second"],
        autoScale: "none",
        timeBase: "milliseconds"
      },
      series: {
      lines: {
        show: true
      },
      points: {
        show: true
      },
    },
    grid: {
      hoverable: true,
      clickable: true
    },
  }
  $.plot($(id),  data , config);
$(window).resize(function(){
  $.plot($(id),  data , config);
 })
}

function useTooltip(chartId, tooltipId){
  $(chartId).bind("plothover",function(event, pos, item){
    if(item){
      $(tooltipId).html(item.datapoint[1].toFixed(2)).show().css({top: item.pageY+5,left: item.pageX+5})
      console.log(item)
    }else{
      $(tooltipId).hide()
    }
  })
}