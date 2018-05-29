$(document).ready(function () {
  $('button#status-check').click(function () {
    let url = '/api/v1/status'
	$.get(url, function(data, status) {
		console.log(data, status)
	})
 })
})
