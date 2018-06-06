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

  // Make the default base ubuntu:14.04
  if (containerSpecs['base'] === undefined) {
    containerSpecs['base'] = 'cod-ubuntu:14.04'
  }

  // Make a GET request to send container specs to API
  $('input#start').click(function (event) {
    event.preventDefault();
    $('.select-container').hide()
    $('.nav').hide()
    $('#loader').show()

    // Validate github repo
    let val = $('input#repo-input').val()
    if (val !== '') {
      $('.message p').text('Validating repository...')
      let len = 'https://github.com/'.length
      let repo = val.slice(len)
      let url = 'https://api.github.com/repos/' + repo
      $.ajax ({
        type: 'GET',
        url: url,
        statusCode: {

          // If repo is not valid...
          404: function() {
            $('#repoFlash').delay(500).fadeIn('normal', function() {
              $('#loader').hide()
              $('.select-container').show()
            })
          }
        },

        // If repo is valid...
        success: function (data) {

          // Add repo to containerSpecs
          $('.message p').text('GitHub reposority validated.')
          cloneURL = data.clone_url
          // Replacing '/' with '|' to pass as url query string
          containerSpecs['cloneURL'] = cloneURL.replace(/\//g, '|')
          containerSpecs['name'] = data.name.toLowerCase()
          containerSpecs['id'] = String(data.id)

          // Start build...
          $('.select-container').hide()
          $('.nav').hide()
          $('#loader').show()
          startContainer();
        }
      })
    } else {
      startContainer();
    }

  })
})
