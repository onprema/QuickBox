
let containerSpecs = {};
let cloneURL = ''

function checkBuild() {
  let checkURL = '/api/v1/builds'
  let out = $.ajax({type: 'GET', url: checkURL, async: false}).responseText
  console.log(out)
  setTimeout(checkBuild, 1000)
}
function verifyKey() {

  // Collect the user's public key
  let key = $("textarea#pk").val();

  let registerURL = '/api/v1/register'
  $.ajax({
    method: "POST",
    url: registerURL,
    data: {
      'username': 'theo',
      'password': 123,
      'key': key,
    },
    headers: {
      'it_is': 'you\'re boy',
    }
    .done(function(msg) {
      console.log(msg)
    })
  })
  console.log(registerURL)
}

$(document).ready(function () {

  let containerHash = ''

  // Collect the base image for the container to be ran (required)
  $("input#base").on('click', function () {
    if ($(this).prop('checked')) {
      containerSpecs[$(this).attr('data-key')] = $(this).attr('data-val')
    } else {
      delete containerSpecs[$(this).attr('data-key')]
    }
  })


  // Make a GET request to send container specs to API
  $('input#start').click(function (event) {
    verifyKey();
    event.preventDefault();
    let query = '';
    if ($('input#base').prop('checked') == false) {
      $('#baseFlash').delay(500).fadeIn('normal', function () {
        $(this).delay(2500).fadeOut()
      })
    } else {
      $('.select-container').hide()
      $('#loader').show()
      for (let key in containerSpecs) {
        let val = containerSpecs[key];
        query += key + '=' + val + '&'
      }

      //checkBuild()

      let url = '/api/v1/start/' + query
      $.get(url, function (data, status) {
        data = JSON.parse(data)
        console.log('Container: ', data)
        port = data.Port
        containerId = data.Id

        if (data.Id) {

          displaySwitch(data);

          // When user clicks Destroy
          $('button#destroy').click(function (event) {
            event.preventDefault()
            $('#loader').show()
            removeURL = '/api/v1/remove/' + containerId
            $.get(removeURL, function (data, status) {
              console.log('destroy', status)
            })
            $('#loader').hide()
            location.reload()
          })
        }
      })
    }
  })
})
