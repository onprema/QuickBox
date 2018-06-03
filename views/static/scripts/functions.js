// Verify Public Key (TODO)
function verifyKey(key) {
  if (key.length < 20) {
    return null
  } else {
    creds['key'] = key
  } return 1
}

function verifyGithubUsername() {
  let username = $('input#username').val()
  let url = 'https://api.github.com/users/' + username
  // Validate github username
  $.ajax ({
    type: 'GET',
    url: url,
    statusCode: {
      404: function() {
        $('div#userFlash').delay(500).fadeIn('normal', function() {
          $(this).delay(5000).fadeOut()
        })
      }
    },
    success: function (data) {
      // Add username to creds object
      creds['username'] = username;
      $('form#registration').hide()
      $('span#username').html(username)
      $('div#userSuccess').delay(200).fadeIn('normal', function() {
        $(this).show()
      })
    }
  })
 
  // Check is user has public keys linked to the github account
  $.ajax ({
    type: 'GET',
    url: url + '/keys',
    success: function (data) {
      if (data.length > 0) {
        let keys = []
        for (let i = 0; i < data.length; i++) {
          keys.push(data[i]['key'])
        }

      // Add keys to creds object
      creds['keys'] = keys;

      // If they have more than one key, they will be "chained together"
      $('textarea#keySuccess').delay(500).fadeIn('normal', function() {
        $(this).text(creds.keys.join(',').replace(',', '\n', -1))
      })

      // Display they github keys and confirm they want to use them
      $('div#keyMsgSuccess').delay(500).fadeIn('normal', function() {
        $(this).show()
      })

      // If they click yes, send key to backend and display service
      $('button#keyConfirmYes').click(function (event) {
        event.preventDefault()
        // send creds to backend
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
      })
      } else {
        $('div#keyMsgFailure').delay(500).fadeIn('normal', function() {
          $(this).show()
        })
        $('button#pk').click( function(event) {
          event.preventDefault()
          let key = $('textarea#pk').val()
          let result = verifyKey(key);
          if (result != null) {
            creds['keys'] = key
            $('div#keyMsgFailure').hide()
            $('div#keySuccess').delay(500).fadeIn('normal', function() {
              $(this).text('Key added successfully. Let\'s go!')
            })
          } else {
            $('div#badKeyFlash').delay(500).fadeIn('normal', function() {
              $(this).text('Invalid key. Please verify and try again.')
            })
          }
        })
      }
   }
  })
}
