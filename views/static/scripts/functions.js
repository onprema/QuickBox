// Verify Public Key (TODO)
function verifyKey(key) {
  if (key.length < 20) {
    return null
  } else {
    creds['key'] = key
  } return 1
}


function startContainer() {
  $('.message p').text('Building container...')
  let query = '';
  for (let key in containerSpecs) {
    let val = containerSpecs[key];
    query += key + '=' + val + '&'
  }

  console.log('[' + query +']')
  let url = '/api/v1/start/' + query
  $.get(url, function (data, status) {
    data = JSON.parse(data)
    console.log('Container: ', data)
    port = data.Port
    containerId = data.Id
    if (data.Id) {
      $('.message p').text('Done!')
      displaySwitch(data);
      // When user clicks Destroy
      $('button#destroy').click(function (event) {
        event.preventDefault()
        $('#loader').show()
        removeURL = '/api/v1/remove/' + containerId
        $.get(removeURL, function (data, status) {
          console.log('destroying container: ', data)
        })
        $('#loader').hide()
        location.reload()
      })
    }
  })
}

function verifyGithubUsername() {
  let username = $('input#username').val()
  let url = 'https://api.github.com/users/' + username
  // Validate github username
  $('div#loader').show()
  $('.message p').text('Verifying GitHub username...')
  $.ajax ({
    type: 'GET',
    url: url,
    statusCode: {
      404: function() {
        $('div#loader').hide()
        $('.message p').text('Invalid GitHub username. Please try again')
        $('.register').show()
        $('.register p').text('')
      }
    },
    success: function (data) {
      // Add username to creds object
      creds['username'] = username;
      $('form#registration').hide()
      $('.message p').text('Welcome, ' + username + '!')
    }
  })
 
  // Check if user has public keys linked to the github account
  $.ajax ({
    type: 'GET',
    url: url + '/keys',
    success: function (data) {
      $('div#loader').hide()
      if (data.length > 0) {
        let keys = []
        for (let i = 0; i < data.length; i++) {
          keys.push(data[i]['key'])
        }

        // Add keys to creds object
        creds['keys'] = keys;

        // Display the github key and confirm they want to use them
        $('div#keyMsgSuccess').delay(500).fadeIn('normal', function() {
          $(this).show()
        })

        // If they click yes, send key to backend and display service
        $('button#keyConfirmYes').click(function (event) {
          event.preventDefault()

          $('.message p').text('Adding SSH keys to server...')
          // send public key to backend
          let credURL = '/api/v1/register'
          $.ajax ({
            type: 'GET',
            url: credURL,
            headers: {
              creds: JSON.stringify(creds['keys'])
            },
            success: function(data) {
              $('.message p').text('SSH key added successfully!')
              console.log(data)
            }
          })

          // Display the service
          $('div#loader').hide()
          $('div#keyMsgSuccess').hide()
          $('.service').show()
          $('.message p').text('Welcome, ' + username + '!')

          // Set cookie?
					Cookies.set(username, true, { expires: .01 })
          console.log(document.cookie)
        })

        // Alternatively, provide an SSH key
        $('button#customKey').click( function(event) {
          event.preventDefault()
          let key = $('textarea#customKey').val()
          let result = verifyKey(key);
          if (result != null) {
            $('div#loader').hide()
            $('.message p').text('SSH key added successfully!')
            creds['keys'] = key
            $('div#keyMsg').hide()

            // send public key to backend
            let credURL = '/api/v1/register'
            $.ajax ({
              type: 'GET',
              url: credURL,
              headers: {
                creds: JSON.stringify(creds['keys'])
              },
              success: function(data) {
                console.log(data)
              }
            })
            $('div#keySuccess').hide()
            $('.service').show()
          } else {
            $('div#badCustomKeyFlash').delay(500).fadeIn('normal', function() {
              $(this).text('Invalid key. Please verify and try again.')
              setTimeout(function () {
                location.reload()
              }, 5000)
            })
          }
        })
      } else {
        $('div#keyMsg').delay(500).fadeIn('normal', function() {
          $('div#loader').hide()
          $(this).show()
        })
        }
      }
   })
  }
