$(document).ready(function () {

  // Get GitHub repo (optional)
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
    val = val.split('/')
    val = val.slice(3)
    val = val.join(':')
    containerSpecs['repo'] = val
	})

})
