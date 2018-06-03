let containerSpecs = {};
let cloneURL = ''
let containerHash = ''

$(document).ready(function () {

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
