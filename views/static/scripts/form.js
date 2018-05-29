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

  // Get GitHub repo
  $('button#repo').on('click', function () {
    event.preventDefault();
    let val = $('#repo-input').val()
    let len = 'https://github.com/'.length
    let repo = val.slice(len)
    let url = 'https://api.github.com/repos/' + repo
    $.ajax ({
      type: 'GET',
      url: url,
      success: function (data) {
        console.log('success', data)
      }
    })

    // Parse repo URL to remove slashes
    val = val.split('/') // ["https:", "", "github.com", "user", "repo"]
    val = val.slice(3)   // ["user", "repo"]
    val = val.join(':')  // ["user:repo"]
    containerSpecs['repo'] = val
	})

  // Make a GET request to check API status
  $('button#status-check').click(function () {
    let url = '/api/v1/status'
	$.get(url, function(data, status) {
		console.log(data, status)
	})
 })

  // Make a POST request to send container specs to API
  $('input#start').click(function (event) {
    event.preventDefault();
    let query = '';
    for (let key in containerSpecs) {
      let val = containerSpecs[key];
      query += key + '=' + val + '&'
    }
    let url = '/api/v1/start/' + query
    $.get(url, function (data, status) {
      console.log(data)
    })
  })
})
