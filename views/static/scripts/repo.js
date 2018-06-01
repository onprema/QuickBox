$(document).ready(function () {
  
  // Get GitHub repo (optional)
  $('button#repo').on('click', function () {
    event.preventDefault();
    let val = $('input#repo-input').val()
    let len = 'https://github.com/'.length
    let repo = val.slice(len)
    let url = 'https://api.github.com/repos/' + repo
    $.ajax ({
      type: 'GET',
      url: url,
      statusCode: {
        404: function() {
          $('#repoFlash').delay(500).fadeIn('normal', function() {
            $(this).delay(5000).fadeOut()
          })
        }
      },
      success: function (data) {
        $('#repoSuccess').delay(500).fadeIn('normal', function() {
          $(this).css('display', 'block')
        })
        cloneURL = data.clone_url
        // Replacing '/' with '|' to pass as url query string
        containerSpecs['cloneURL'] = cloneURL.replace(/\//g, '|')
        containerSpecs['name'] = data.name.toLowerCase()
        containerSpecs['id'] = String(data.id)
        console.log('FROM GITHUB API: ', data)
      }
    })
	})
})
