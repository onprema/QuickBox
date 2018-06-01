let containerSpecs = {};
let cloneURL = ''

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
    event.preventDefault();
    let query = '';
    for (let key in containerSpecs) {
      let val = containerSpecs[key];
      query += key + '=' + val + '&'
    }
    let url = '/api/v1/start/' + query
    console.log('URL: ', url)
    $.get(url, function (data, status) {

      data = JSON.parse(data) // {'Id': 'adf9a8dc09', 'Port': '23435', Name: repoName}
      console.log('Container data: ', data)
      port = data.Port
      containerId = data.Id

      if (data.Id) {

        displaySwitch(data);

        // When user clicks Destroy
        $('button#destroy').click(function (event) {
          event.preventDefault()
          let query = '';
          for (let key in containerSpecs) {
            let val = containerSpecs[key];
            query += key + '=' + val + '&'
          }
          removeURL = '/api/v1/remove/' + containerId
          let containerDisplay = 'li#' + containerId
          $.get(removeURL, function (data, status) {
            console.log('destroy', status)
            $(containerDisplay).remove()
            $('.select-container').css('display', 'block')
          })

          // Hide the running containers div
          $('.running-container').css('display', 'none')
        })
      }
    })
  })
})
