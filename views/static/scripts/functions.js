// Makes hash for root password
function hash() {
  let hash = '';
  let range = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
  for (let i = 0; i < 16; i++) {
    hash += range.charAt(Math.floor(Math.random() * range.length))
  }
  return hash;
}

function startContainer () {
  $('.message p').text('Building container...');
  let query = '';
  for (let key in containerSpecs) {
    let val = containerSpecs[key];
    query += key + '=' + val + '&';
  }
  let url = '/api/v1/start/' + query;
  $.get(url, function (data, status) {
    if data.Id === undefined {
        alert('invalid request');
        return;
    }
    data = JSON.parse(data);
    port = data.Port;
    containerId = data.Id;
    if (data.Id) {
      $('.message p').text('Done!');
      displaySwitch(data);
      $('button#destroy').click(function (event) {
        event.preventDefault();
        removeURL = '/api/v1/remove/' + containerId;
        $.get(removeURL, function (data, status) {});
        $('.running-container').hide()
        $('.footer').hide()
        $('.message p').text('Your container has been destroyed. Returning to home page.')
        function reload() { location.reload() };
        setTimeout(reload, 3000);
      });
    }
  });
}

function displaySwitch (data) {
  $('#loader').hide();
  $('.running-container').show();
  $('.cmd').append('ssh root@138.68.8.138 -p ' + String(data.Port));
  $('.pw').append('<p>Password: <strong>' + String(data.Password) + '</strong></p>')
  $('.running-container').append('<button id="destroy">Destroy</button>');
  $('.running-container').append('<p><span style="color: maroon">This container will automatically be destroyed in 5 minutes.</p></span>');
  $('.footer').show()
}
