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
      changeEvent(event);
    }
  });


function changeEvent(eventName) {
  $.get("api/event/"+eventName, function(data, status){
    //$('#eventName').html('<h2>'+ data.eventName.trim() +'</h2>');
    var tabledata = "";
    var thereIsAGroup = false;
    var thereIsATier = false;
    for (var key in data.Participants) {
      tabledata = tabledata + '<tr><td><a href="https://www.twitch.tv/' + data.Participants[key]['twitchUsername'] + '" target="_blank" >' + data.Participants[key]['twitchUsername'] + '</td><td>' + data.Participants[key]['eventPoints'] + '</td><td>' + data.Participants[key]['eventPlayed'] + '</td>';
      if (data.Participants[key]['groupName']) {
        tabledata = tabledata + '<td>' + data.Participants[key]['groupName'] + '</td>';
        thereIsAGroup = true;
      } else if (eventName == "season_7") {
        tabledata = tabledata + '<td>' + data.Participants[key]['tierName'] + '</td>';
        thereIsATier = true;
      }
      tabledata = tabledata + '</tr>';
    }
    if (thereIsAGroup) {
      $('#eventData').html('<table class="alt" ><thead><tr><th>Name</th><th>Points</th><th>Played</th><th>Group</th></tr></thead><tbody>'+tabledata+'</tbody></table>');
    } else if (thereIsATier) {
      $('#eventData').html('<table class="alt" ><thead><tr><th>Name</th><th>Points</th><th>Played</th><th>Tier</th></tr></thead><tbody>'+tabledata+'</tbody></table>');
    } else {
      $('#eventData').html('<table class="alt" ><thead><tr><th>Name</th><th>Points</th><th>Played</th></tr></thead><tbody>'+tabledata+'</tbody></table>');
    }
  });
}
