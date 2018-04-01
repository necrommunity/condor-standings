$(function($) {
  $("#reset").on("click", function () {
      $("#eventSelector").val($("#eventSelector").data("default-value"));
      $("#eventName").html('');
      $("#eventData").html('');

    });
  });

  $("#refresh").on("click", function () {
    changeEvent($("#eventSelector").val());
  });


function changeEvent(eventName) {
  $.get("api/"+eventName, function(data, status){
    //$('#eventName').html('<h2>'+ data.eventName.trim() +'</h2>');
    var tabledata = "";
    for (var key in data.Participants) {
      tabledata = tabledata + '<tr><td>' + data.Participants[key]['discordUsername'] + '</td><td>' + data.Participants[key]['eventPoints'] + '</td><td>' + data.Participants[key]['eventPlayed'] + '</td><td>' + data.Participants[key]['groupName'] + '</td></tr>';
    }
    $('#eventData').html('<table class="alt" border="1"><thead><tr><th>Name</th><th>Points</th><th>Played</th><th>Group</th></tr></thead><tbody>'+tabledata+'</tbody></table>');
  });
}
