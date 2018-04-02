$(function($) {
  $("#reset").on("click", function () {
      $("#eventSelector").val($("#eventSelector").data("default-value"));
      $("#eventName").html('');
      $("#eventData").html('');

    });
  });

  $("#refresh").on("click", function () {
    var event = $("#eventSelector").val();
    console.log(event);
    if (event !== null ){
      changeEvent();
    }
  });


function changeEvent(eventName) {
  $.get("api/event/"+eventName, function(data, status){
    //$('#eventName').html('<h2>'+ data.eventName.trim() +'</h2>');
    var tabledata = "";
    var thereIsAGroup = false;
    for (var key in data.Participants) {
      tabledata = tabledata + '<tr><td>' + data.Participants[key]['discordUsername'] + '</td><td>' + data.Participants[key]['eventPoints'] + '</td><td>' + data.Participants[key]['eventPlayed'] + '</td>';
      if (data.Participants[key]['groupName']) {
        tabledata = tabledata + '<td>' + data.Participants[key]['groupName'] + '</td>';
        thereIsAGroup = true;
      }
      tabledata = tabledata + '</tr>';
    }
    if (thereIsAGroup) {
      $('#eventData').html('<table class="alt" ><thead><tr><th>Name</th><th>Points</th><th>Played</th><th>Group</th></tr></thead><tbody>'+tabledata+'</tbody></table>');
    } else {
      $('#eventData').html('<table class="alt" ><thead><tr><th>Name</th><th>Points</th><th>Played</th></tr></thead><tbody>'+tabledata+'</tbody></table>');
    }
  });
}