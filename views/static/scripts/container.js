
$(document).ready(function () {

  let containerHash = ''
  let containerSpecs = {};

  // Collect the base image for the container to be ran (required)
  $("input#c-base").on('click', function () {
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
    $.get(url, function (data, status) {
      data = JSON.parse(data)
      containerId = data.Id
      containerHash = data.Id
      if (data.Id) {
        $("ul#running-containers").append(
          '<li id="' + containerId + '">' + String(data.Id) + 
          '<button id="' + containerId + '">Destroy</button>' + '</li>')
        let removeButton = 'button#' + containerHash
        console.log('removeButton: ', removeButton)
        $(removeButton).click(function (event) {
          event.preventDefault()
          console.log('Removing container: ', containerHash.slice(8))
          let url = '/api/v1/remove/' + containerHash
          let containerDisplay = 'li#' + containerHash
          $.get(url, function (data, status) {
            console.log(status)
            $(containerDisplay).remove()
          })
        })
      }
    })
  })
})
