$(document).ready(function () {

  // Collect the specs for the container to be ran
  let containerSpecs = {};
  $("input#c-base").on('click', function () {
    if ($(this).prop('checked')) {
      containerSpecs[$(this).attr('data-key')] = $(this).attr('data-val')
    } else {
      delete containerSpecs[$(this).attr('data-key')]
    }
  })

  // Make a GET request to check API status
  $('button#status-check').click(function () {
    let url = 'http://0.0.0.0:5001/api/v1/status'
    $.get(url, function (data, status) {
      console.log(data)
    })
  })

  // Make a POST request to send container specs to API
  $('input#start').click(function (event) {
    event.preventDefault();
    let query = '';
    for (let key in containerSpecs) {
      let val = containerSpecs[key];
      query += key + '=' + val
    }
    let url = 'http://0.0.0.0:5001/api/v1/start/' + query
    $.get(url, function (data, status) {
      console.log(data)
    })
  })
})
