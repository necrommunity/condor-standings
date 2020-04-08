function changeEvent(eventName, prettyName, userName) {
    $.get("/api/user/"+eventName+"/"+userName, function(data, status){
      //$('#eventName').html('<h2>'+ data.eventName.trim() +'</h2>');
      
      var totalTime = 0;
      var totalWins = 0;
      var totalLoss = 0;
      var totalRaces = 0;
      var fastestTime = 0;

      if (data == null) {
        $('tbody#info').html("");
        $('tbody#races').html("");
      }
      for (race of data){
        let winner = "racer"+race["raceWinner"]+"Name";
        totalRaces += 1;

        if (race[winner] == userName){
          totalWins += 1;
          totalTime += race["raceTime"]
          if (fastestTime == 0 || fastestTime > race["raceTime"]){
            fastestTime = race["raceTime"]
          }
          } else {
            totalLoss +=1;
        }
        // tabledata = tabledata + '<tr><td><a href="/user/' + race["racer1Name"] + '">'+ race["racer1Name"] + '</a></td></tr>'
      }
      let winsdata = '<tr><td width="" id="wins-td"> Wins: ' + totalWins + ' (' + Math.round((totalWins/totalRaces)*100) + '%) | Losses: '+ totalLoss + ' (' + Math.round((totalLoss/totalRaces)*100) + '%) </td><td id="fastest-win">Fastest Win: ' + convertToFormatTime(fastestTime) + ' | Average Win: ' + convertToFormatTime(totalTime/totalWins) + '</td><td width="" style="text-align: right"> <a href="https://twitch.tv/'+userName+'" target="_blank"><img src="/public/images/Glitch_White_RGB.png"  height="24px" width="24px"/></a></td></tr>'
      let winsbody = ''
      for (race of data) {
        winsbody += '<tr>'
        winsbody += '<td class="results" id="racer1"><a href="/user/'+race["racer1Name"]+'">'+race["racer1Name"] +'</a>' 
        if (race["raceWinner"]==1) {
          winsbody += ' 🎉 '
        }
        winsbody += '</td>'
        winsbody += '<td class="results" id="racer2"><a href="/user/'+race["racer2Name"]+'">'+race["racer2Name"] + '</a>'
        if(race["raceWinner"]==2) {
          winsbody += ' 🎉 '} 
        winsbody += '</td>'
        winsbody += '<td class="results" id="race-time">' + race["raceTimeF"] + '</td>'
        winsbody += '<td class="results" id="race-seed">'+ race["raceSeed"] + '</td>'
        winsbody += '<td class="results" id="race-type">'

        if (race["isAutoGen"]) {
          winsbody +=  'Autogen'
        } else {
          winsbody += 'Challenge'
        }
        winsbody += '</td>'
        winsbody += '<td class="results" id="race-vod">'
        if (race["raceVod"]["Valid"] == true){
          winsbody += '<a href="' + race["raceVod"]["String"] +'" target="_blank">🎥</a>'
        }
        winsbody += '</td>'
        winsbody += '</tr>'

      // console.log(totalWins +" "+totalTime+ " "+fastestTime+" "+convertToFormatTime(fastestTime));
      }
      $('tbody#info').html(winsdata);
      $('tbody#races').html(winsbody);
    });

    $(document).prop('title', 'Match Info');
  }
  
  $( window ).on( "load", function() {
    var tUsername = $('#twitchUsername').prop('value');
    changeEvent($("#eventSelector")[0][1].value, $("#eventSelector")[0][1].text, tUsername); 
  });

  function convertToFormatTime(t) {

  let fTime="" 
	let w = parseInt(t/100)
	let ms = Math.abs(parseInt((((t/100) - w) * -100)))
	let s = parseInt(((t/100) % 60))
	let m = parseInt((t/(100*60)) % 60)
	let h = parseInt(((t/(100*60*60)) % 24))
	if (h > 0) {
		fTime += zeroPad(h, 2)+":"
	}
  fTime += zeroPad(m,2)
  fTime += ":"
  fTime += zeroPad(s,2)
  fTime +="."
  fTime +=zeroPad(ms,2)
	return fTime
}

function zeroPad(n,length){
  var s=n+"",needed=length-s.length;
  if (needed>0) s=(Math.pow(10,needed)+"").slice(1)+s;
  return s;
}