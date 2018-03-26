$(function($) {

});

function changeEvent(eventName) {
  console.log(eventName);
  $.get("api/"+eventName, function(data, status){
    $('#eventName').html(data.eventName.trim());
    var tabledata = "";
    for (var key in data.Participants) {
      console.log(tabledata);
      tabledata = tabledata + '<tr><td>' + data.Participants[key]['discordUsername'] + '</td><td>' + data.Participants[key]['eventPoints'] + '</td><td>' + data.Participants[key]['eventPlayed'] + '</td><td>' + data.Participants[key]['groupName'] + '</td></tr>';
    }
    $('#eventData').html('<table border="1">'+tabledata+'</table>');
  });

}
